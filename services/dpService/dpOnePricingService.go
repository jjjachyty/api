package dpService

import (
	"errors"
	"fmt"

	"pccqcpa.com.cn/app/rpm/api/models/amData"
	"pccqcpa.com.cn/app/rpm/api/models/biz/ln"
	"pccqcpa.com.cn/app/rpm/api/models/dp"
	"pccqcpa.com.cn/app/rpm/api/services/commonService"
	"pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/components/zlog"
)

type DpOnePricingService struct{}

var dpOnePricing dp.DpOnePricing

// 分页操作
// by author Jason
// by time 2016-12-12 10:44:29
func (d DpOnePricingService) List(param ...map[string]interface{}) (*util.PageData, error) {
	var paramMap map[string]interface{}
	var handleParamStrs = []string{
		"business_code", "cust.cust_code", "cust.cust_name",
	}
	var utilBase = util.GetUtilBase()
	if 0 < len(param) {

		paramMap = utilBase.DeepCopy(param[0]).(map[string]interface{})
	}
	// 判断是否有非法字符
	flag := util.CheckQueryParams(paramMap)
	if flag {
		return nil, fmt.Errorf("查询条件中不允许出现[^$.*+?{}()[]|'`~\\]特殊字符")
	}
	if _, ok := paramMap[util.RPM_CUST_OWNER]; !ok {
		utilBase.HandleParamToLike(paramMap, handleParamStrs...)
		return dpOnePricing.List(paramMap)
	}
	var currentUser = paramMap[util.RPM_CUST_OWNER].(amData.SidUser)
	var owner = currentUser.UserId
	delete(paramMap, util.RPM_CUST_OWNER)
	if util.RPM_CM == currentUser.RoleIds {
		utilBase.HandleParamToLike(paramMap, handleParamStrs...)
		paramMap["cust.owner"] = owner
		return dpOnePricing.List(paramMap)
	} else {

		dpOnePricingListMap := utilBase.DeepCopy(paramMap).(map[string]interface{})
		var vMap = map[string]interface{}{}
		util.SplitMap(dpOnePricingListMap, vMap, []string{"start", "length", "order", "sort"})

		simMap := utilBase.DeepCopy(dpOnePricingListMap).(map[string]interface{})
		realMap := utilBase.DeepCopy(dpOnePricingListMap).(map[string]interface{})
		delete(simMap, util.SEARCH_LIKE)
		utilBase.HandleParamToLike(simMap, handleParamStrs...)
		utilBase.HandleParamToLike(realMap, handleParamStrs...)
		simMap["CUST.OWNER"] = owner
		simMap["CUST.STATUS"] = util.SIMULATION_CUST

		realMap["CUST.STATUS"] = util.REAL_CUST

		return dpOnePricing.UnionList(vMap, simMap, realMap)
	}
}

// 多参数查询返回多条纪录
// by author Jason
// by time 2016-12-12 10:44:29
func (DpOnePricingService) Find(param ...map[string]interface{}) ([]*dp.DpOnePricing, error) {
	if 0 < len(param) {
		delete(param[0], util.SEARCH_LIKE)
	}
	return dpOnePricing.Find(param...)
}

// 多参数查询返回一条纪录
// by author Jason
// by time 2016-12-12 10:44:29
func (this DpOnePricingService) FindOne(paramMap map[string]interface{}) (*dp.DpOnePricing, error) {
	models, err := this.Find(paramMap)
	if nil != err {
		return nil, err
	}
	switch len(models) {
	case 0:
		return nil, nil
	case 1:
		return models[0], nil
	}
	er := errors.New("查询有多条返回纪录")
	zlog.Error(er.Error(), er)
	return nil, er
}

// 新增纪录
// by author Jason
// by time 2016-12-12 10:44:29
func (DpOnePricingService) Add(model *dp.DpOnePricing) error {
	return model.Add()
}

// 更新纪录
// by author Jason
// by time 2016-12-12 10:44:29
func (DpOnePricingService) Update(model *dp.DpOnePricing) error {
	return model.Update()
}

// 部分更新
func (d DpOnePricingService) Patch(paramMap map[string]interface{}) error {
	return dpOnePricing.Patch(paramMap)
}

// 删除纪录
// by author Jason
// by time 2016-12-12 10:44:29
func (DpOnePricingService) Delete(model *dp.DpOnePricing) error {
	return model.Delete()
}

// 删除纪录
// by author Jason
// by time 2016-12-12 10:44:29
func (d *DpOnePricingService) BegainPricing(dpOnePricing *dp.DpOnePricing, cust ln.CustInfo) error {
	// 获取存款业务编号
	businessCode := commonService.CodeService{}.GetOneDpBusinessByTime()

	dpOnePricing.Cust = cust
	dpOnePricing.BusinessCode = businessCode
	dpOnePricing.Status = util.DP_ONE_PRICING_STATUS_UNFINISHED

	err := dpOnePricing.Add()
	if nil != err {
		return err
	}
	return nil
}

func (d DpOnePricingService) callBack(slice1, slice2 interface{}) interface{} {
	return append(slice1.([]*dp.DpOnePricing), slice2.([]*dp.DpOnePricing)...)
}
