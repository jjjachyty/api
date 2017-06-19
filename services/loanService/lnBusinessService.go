package loanService

import (
	"fmt"
	"github.com/lunny/tango"

	"pccqcpa.com.cn/app/rpm/api/models/amData"
	"pccqcpa.com.cn/app/rpm/api/models/biz/ln"
	"pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/components/zlog"
)

//LnBusinessService type 定价单实体
type LnBusinessService struct {
	tango.Ctx
}

var lnBusinessModel ln.LnBusiness

func (l *LnBusinessService) List(param ...map[string]interface{}) (*util.PageData, error) {
	var paramMap map[string]interface{}
	var hendleParamLikeStrs = l.getHanleParamLikeStrs()
	if 0 < len(param) {
		paramMap = utilBase.DeepCopy(param[0]).(map[string]interface{})
	}
	// 判断是否有非法字符
	flag := util.CheckQueryParams(paramMap)
	if flag {
		return nil, fmt.Errorf("查询条件中不允许出现[^$.*+?{}()[]|'`~\\]特殊字符")
	}

	var currentUser = paramMap["owner"].(amData.SidUser)
	var owner = currentUser.UserId
	delete(paramMap, "owner")
	if util.RPM_CM == currentUser.RoleIds {
		utilBase.HandleParamToLike(paramMap, hendleParamLikeStrs...)
		paramMap["cust.owner"] = owner
		return lnBusinessModel.List(paramMap)
	} else {

		lnbusinessListMap := utilBase.DeepCopy(paramMap).(map[string]interface{})
		var vMap = map[string]interface{}{}
		util.SplitMap(lnbusinessListMap, vMap, []string{"start", "length", "order", "sort"})

		simMap := utilBase.DeepCopy(lnbusinessListMap).(map[string]interface{})
		realMap := utilBase.DeepCopy(lnbusinessListMap).(map[string]interface{})
		delete(simMap, util.SEARCH_LIKE)
		utilBase.HandleParamToLike(simMap, hendleParamLikeStrs...)
		utilBase.HandleParamToLike(realMap, hendleParamLikeStrs...)
		simMap["CUST.OWNER"] = owner
		simMap["CUST.STATUS"] = util.SIMULATION_CUST

		realMap["CUST.STATUS"] = util.REAL_CUST

		return lnBusinessModel.UnionList(vMap, simMap, realMap)
	}
}

func (l LnBusinessService) getHanleParamLikeStrs() []string {
	return []string{
		"business_code", "cust.cust_code", "cust.cust_name", "organ.organ_name",
	}
}

func (l *LnBusinessService) Find(param ...map[string]interface{}) ([]*ln.LnBusiness, error) {
	return lnBusinessModel.Find(param...)
}

// 先判断是否存在，如果存在就更新
func (l *LnBusinessService) Add(lnBusiness *ln.LnBusiness) error {
	paramMap := map[string]interface{}{
		"business_code": lnBusiness.BusinessCode,
	}
	lbs, err := lnBusiness.Find(paramMap)
	if nil != err {
		return err
	}
	if 1 == len(lbs) {
		lnBusiness.UUID = lbs[0].UUID
		return l.Update(lnBusiness)
	}
	if 1 < len(lbs) {
		er := fmt.Errorf("业务编号为[%v]的业务记录数有[%v]条", lnBusiness.BusinessCode, len(lbs))
		zlog.Error(er.Error(), err)
		return er
	}
	err = lnBusiness.Add()
	if nil != err {
		return err
	} else {
		return l.DeleteMortgage(lnBusiness)
	}
}

// 判断贷款客户是否和保证人一样
func (l *LnBusinessService) Update(lnBusiness *ln.LnBusiness) error {
	if util.MAIN_MORTGAGE_TYPE_CREDIT != lnBusiness.MainMortgageType {
		paramMap := map[string]interface{}{
			"business_code": lnBusiness.BusinessCode,
			"guarante":      lnBusiness.Cust.CustCode,
		}
		guarantes, err := LnGuaranteService{}.Find(paramMap)
		if nil != err {
			er := fmt.Errorf("更新业务信息时查询保证人出错")
			zlog.Error(er.Error(), err)
			return er
		}
		if 0 < len(guarantes) {
			er := fmt.Errorf("贷款客户与保证人相同，请先删除保证人")
			zlog.Error(er.Error(), er)
			return er
		}
	}
	err := lnBusiness.Update()
	if nil != err {
		return err
	} else {
		return l.DeleteMortgage(lnBusiness)
	}
}

// 保存业务信息之后，如果担保方式为信用，则删除抵押品信息
func (l LnBusinessService) DeleteMortgage(lnBusiness *ln.LnBusiness) error {
	if util.MAIN_MORTGAGE_TYPE_CREDIT == lnBusiness.MainMortgageType {
		err := LnMortService{}.DeleteByBusinessCode(lnBusiness.BusinessCode)
		if nil != err {
			return err
		} else {
			return LnGuaranteService{}.DeleteByBusinessCode(lnBusiness.BusinessCode)
		}
	} else {
		return nil
	}
}

// 删除定价结果、保证人、抵质押、派生等记录
func (l LnBusinessService) Delete(lnBusiness *ln.LnBusiness) error {
	var mort ln.LnMort
	var guarante ln.LnGuarante
	// 删除抵质押
	err := mort.DeleteByBusinessCode(lnBusiness.BusinessCode)
	if nil != err {
		return err
	}
	// 删除保证人
	err = guarante.DeleteByBusinessCode(lnBusiness.BusinessCode)
	if nil != err {
		return err
	}

	// 删除业务信息
	return lnBusiness.Delete()
}

// BUG(Jason): #1: 未实现事务
// BUG(Jason): #2: 未删除派生信息

func (l LnBusinessService) callBack(slice1, slice2 interface{}) interface{} {
	return append(slice1.([]*ln.LnBusiness), slice2.([]*ln.LnBusiness)...)
}
