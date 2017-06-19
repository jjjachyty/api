package ln

import (
	"fmt"

	"github.com/lunny/tango"

	"pccqcpa.com.cn/app/rpm/api/models/biz/ln"
	"pccqcpa.com.cn/app/rpm/api/services/loanService"
	"pccqcpa.com.cn/app/rpm/api/services/pricingService"
	"pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/app/rpm/api/util/currentMsg"
	"pccqcpa.com.cn/components/zlog"
)

var ps pricingService.PricingService

var lnBusinessService loanService.LnBusinessService

type LnBusinessAction struct {
	tango.Json
	tango.Ctx
}

func (l *LnBusinessAction) List() util.RstMsg {
	paramMap := util.GetPageMsg(&l.Ctx)

	user, err := currentMsg.GetCurrentUser(l.Ctx)

	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}

	// if "rpmcm" == user.RoleIds {
	// 	paramMap["owner"] = user
	// }
	paramMap["owner"] = user

	pageData, err := lnBusinessService.List(paramMap)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessListMsg("分页查询对公业务信息成功", pageData)
}

func (l *LnBusinessAction) Get() util.RstMsg {

	paramMap, err := util.GetParmFromRouter(&l.Ctx)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}

	// 更换key
	if val, ok := paramMap["status"]; ok {
		paramMap["lnpricing.status"] = val
		delete(paramMap, "status")
	}

	if util.IsPaginQuery(&l.Ctx) {
		util.GetPageMsg(&l.Ctx, paramMap)
		user, err := currentMsg.GetCurrentUser(l.Ctx)

		if nil != err {
			return util.ErrorMsg(err.Error(), err)
		}
		paramMap["owner"] = user
		pageData, err := lnBusinessService.List(paramMap)
		if nil != err {
			return util.ErrorMsg(err.Error(), err)
		}
		return util.SuccessListMsg("分页查询对公业务信息成功", pageData)
	}

	lbs, err := lnBusinessService.Find(paramMap)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("查询对公业务信息成功", lbs)
}

// 新增订单业务
func (l *LnBusinessAction) Post() util.RstMsg {
	var lnBusiness = new(ln.LnBusiness)
	err := currentMsg.DecodeJson(l.Ctx, lnBusiness)
	// 如果没有机构信息，则取当前机构信息
	if "" == lnBusiness.Organ.OrganCode {
		organCode, err := currentMsg.GetCurrentUserBranchCode(l.Ctx)
		if nil != err {
			return util.ErrorMsg(err.Error(), err)
		}
		lnBusiness.Organ.OrganCode = organCode
	}
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	err = lnBusinessService.Add(lnBusiness)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("新增对公业务信息成功", lnBusiness)
}

// 修改对公业务信息
func (l *LnBusinessAction) Put() util.RstMsg {
	var lnBusiness = new(ln.LnBusiness)
	var lnPricing = new(ln.LnPricing)
	err := currentMsg.DecodeJson(l.Ctx, lnBusiness)
	// 如果没有机构信息，则取当前机构信息
	if "" == lnBusiness.Organ.OrganCode {
		organCode, err := currentMsg.GetCurrentUserBranchCode(l.Ctx)
		if nil != err {
			return util.ErrorMsg(err.Error(), err)
		}
		lnBusiness.Organ.OrganCode = organCode
	}
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	if nil != err {
		err := currentMsg.DecodeJson(l.Ctx, lnPricing)
		if nil != err {
			return util.ErrorMsg(err.Error(), err)
		} else {
			lnBusiness = &lnPricing.LnBusiness
		}
	}
	err = lnBusinessService.Update(lnBusiness)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("更新对公业务信息成功", lnBusiness)
}

// 删除对公业务信息，同时删除对公定价单、抵质押品、保证人信息
func (l *LnBusinessAction) Delete() util.RstMsg {
	var lnBusiness = new(ln.LnBusiness)
	err := l.DecodeJson(&lnBusiness)
	if nil != err {
		er := fmt.Errorf("json数据转换为对公业务信息出错")
		zlog.Errorf(er.Error(), err)
		return util.ErrorMsg(er.Error(), err)
	}
	err = lnBusinessService.Delete(lnBusiness)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("删除对公业务信息成功", nil)
}
