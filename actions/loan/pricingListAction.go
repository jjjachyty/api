package loan

import (
	"errors"
	"fmt"

	"github.com/lunny/tango"
	"pccqcpa.com.cn/app/rpm/api/models/biz/ln"
	"pccqcpa.com.cn/app/rpm/api/services/loanService"
	"pccqcpa.com.cn/app/rpm/api/util"
	// "pccqcpa.com.cn/app/rpm/api/util/currentMsg"
	"pccqcpa.com.cn/components/zlog"
)

//PricingListAction type 对公贷款定价单控制实体
type PricingListAction struct {
	tango.Json
	tango.Ctx
}

var pls loanService.PricingListService

// List PricingListAction.Get func 获取所有的定价单
//URL: https://xxx.xxx.xxx.xxx:xxxx/api/rpm/ln/pricelist
func (pla *PricingListAction) List() util.RstMsg {

	zlog.AppOperateLog("", "PricingListAction.List", zlog.SELECT, nil, nil, "查询对公贷款定价单信息")

	startRowNumber, pageSize, OrderAttr, OrderType, err := util.GetPageAndOrder(pla.Req().Header)

	params, err := util.GetParmFromRouter(&pla.Ctx)

	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}

	// user, err := currentMsg.GetCurrentUser(pla.Ctx)
	// if "rpmcm" == user.RoleIds {
	// 	params["owner"] = user.UserId
	// }
	flag := util.CheckQueryParams(params)

	if flag {
		return util.ErrorMsg("查询条件中不允许出现[^$.*+?{}()[]|'`~\\]特殊字符", errors.New("查询条件中不允许出现[^$.*+?{}()[]|'`~\\]特殊字符"))
	}
	if err == nil {
		pageData, err := pls.GetLnPringListByPageWithCustCodeName(startRowNumber, pageSize, OrderAttr, OrderType, params)
		if err == nil {
			return util.ReturnSuccess("获取对公贷款定价单成功", pageData)
		} else {
			return util.ErrorMsg("获取对公贷款定价单失败", err)
		}
	}

	return util.ErrorMsg("获取对公贷款定价单失败", err)
}

func (pla *PricingListAction) Get() util.RstMsg {

	pid := pla.Param(":businesscode")

	zlog.AppOperateLog("", "PricingListAction.List", zlog.SELECT, nil, nil, "查询对公贷款定价单信息")
	pageData, err := pls.GetLnPringWithPID(pid)
	if err == nil {
		return util.ReturnSuccess("获取对公贷款定价单"+pid+"成功", pageData)
	}
	return util.ErrorMsg(err.Error(), err)
}

func (pla *PricingListAction) Put() util.RstMsg {

	var lnp = new(ln.LnPricing)
	err := pla.DecodeJson(&lnp)
	if nil != err {
		er := fmt.Errorf("json数据转换为对公贷款定价单结构体数据出错")
		zlog.Error(er.Error(), er)
		return util.ErrorMsg(er.Error(), err)
	}
	err = pls.Update(lnp)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("修改对公贷款结果信息成功", nil)
}

func (pla *PricingListAction) Patch() util.RstMsg {

	var pMap map[string]interface{}
	err := pla.DecodeJson(&pMap)
	if nil != err {
		er := fmt.Errorf("json数据转换[MAP]对象数据出错")
		zlog.Error(er.Error(), er)
		return util.ErrorMsg(er.Error(), err)
	}
	// 获取业务单号
	businessCode := pla.Param(":businesscode")
	pMap["business_code"] = businessCode
	err = pls.Patch(pMap)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("修改对公贷款结果信息成功", nil)
}

func (pla *PricingListAction) Delete() util.RstMsg {
	var lnp = new(ln.LnPricing)
	err := pla.DecodeJson(&lnp)
	if nil != err {
		er := fmt.Errorf("json数据转换为对公贷款定价单结构体数据出错")
		zlog.Error(er.Error(), er)
		return util.ErrorMsg(er.Error(), err)
	}
	err = pls.Delete(lnp)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("删除对公贷款信息成功", nil)
}
