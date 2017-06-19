package dp

import (
	"fmt"

	"github.com/lunny/tango"

	"pccqcpa.com.cn/app/rpm/api/models/dp"
	"pccqcpa.com.cn/app/rpm/api/services/dpService"
	"pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/app/rpm/api/util/currentMsg"
	"pccqcpa.com.cn/components/zlog"
)

var dpOnePricingService dpService.DpOnePricingService

type DpOnePricingAction struct {
	tango.Json
	tango.Ctx
}

// 分页查询
// by author Jason
// by time 2016-12-12 10:44:29
func (this *DpOnePricingAction) List() util.RstMsg {
	paramMap := util.GetPageMsg(&this.Ctx)
	err := this.addOwner(paramMap)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	user, err := currentMsg.GetCurrentUser(this.Ctx)

	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	paramMap["owner"] = user

	pageData, err := dpOnePricingService.List(paramMap)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessListMsg("分页查询信息成功", pageData)
}

// 多参数查询
// by author Jason
// by time 2016-12-12 10:44:29
func (this *DpOnePricingAction) Get() util.RstMsg {
	paramMap, err := util.GetParmFromRouter(&this.Ctx)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	err = this.addOwner(paramMap)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	if util.IsPaginQuery(&this.Ctx) {
		util.GetPageMsg(&this.Ctx, paramMap)
		user, err := currentMsg.GetCurrentUser(this.Ctx)

		if nil != err {
			return util.ErrorMsg(err.Error(), err)
		}
		paramMap["owner"] = user
		pageData, err := dpOnePricingService.List(paramMap)
		if nil != err {
			return util.ErrorMsg(err.Error(), err)
		}
		return util.SuccessListMsg("分页查询信息成功", pageData)
	}

	rst, err := dpOnePricingService.Find(paramMap)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("查询信息成功", rst)
}

// 新增信息
// by author Jason
// by time 2016-12-12 10:44:29
func (this *DpOnePricingAction) Post() util.RstMsg {
	var one = new(dp.DpOnePricing)
	err := currentMsg.DecodeJson(this.Ctx, one)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	err = dpOnePricingService.Add(one)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("新增信息成功", one)
}

// 更新信息
// by author Jason
// by time 2016-12-12 10:44:29
func (this *DpOnePricingAction) Put() util.RstMsg {
	var one = new(dp.DpOnePricing)
	err := currentMsg.DecodeJson(this.Ctx, one)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	err = dpOnePricingService.Update(one)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("更新信息成功", one)
}

// 删除信息
// by author Jason
// by time 2016-12-12 10:44:29
func (this *DpOnePricingAction) Delete() util.RstMsg {
	var one = new(dp.DpOnePricing)
	err := this.DecodeJson(one)
	if nil != err {
		er := fmt.Errorf("json转换实体出错")
		zlog.Error(er.Error(), err)
		return util.ErrorMsg(er.Error(), err)
	}
	err = dpOnePricingService.Delete(one)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("删除信息成功", one)
}

func (d DpOnePricingAction) addOwner(paramMap map[string]interface{}) error {
	user, err := currentMsg.GetCurrentUser(d.Ctx)

	if nil != err {
		return err
	}

	if "rpmcm" == user.RoleIds {
		paramMap["cust.owner"] = user.UserId
	}
	return nil
}
