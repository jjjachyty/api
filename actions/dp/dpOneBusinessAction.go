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

var dpOneBusinessService dpService.DpOneBusinessService

type DpOneBusinessAction struct {
	tango.Json
	tango.Ctx
}

// 分页查询
// by author Jason
// by time 2016-12-06 16:59:09
func (this *DpOneBusinessAction) List() util.RstMsg {
	paramMap := util.GetPageMsg(&this.Ctx)
	pageData, err := dpOneBusinessService.List(paramMap)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessListMsg("分页查询信息成功", pageData)
}

// 多参数查询
// by author Jason
// by time 2016-12-06 16:59:09
func (this *DpOneBusinessAction) Get() util.RstMsg {
	paramMap, err := util.GetParmFromRouter(&this.Ctx)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	if util.IsPaginQuery(&this.Ctx) {
		util.GetPageMsg(&this.Ctx, paramMap)
		pageData, err := dpOneBusinessService.List(paramMap)
		if nil != err {
			return util.ErrorMsg(err.Error(), err)
		}
		return util.SuccessListMsg("分页查询存款一对一业务信息成功", pageData)
	}

	rst, err := dpOneBusinessService.Find(paramMap)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("查询存款一对一业务信息成功", rst)
}

// 新增信息
// by author Jason
// by time 2016-12-06  16:59:09
func (this *DpOneBusinessAction) Post() util.RstMsg {
	var one = new(dp.DpOneBusiness)
	err := currentMsg.DecodeJson(this.Ctx, one)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	dpOneBusinessService.Ctx = this.Ctx
	err = dpOneBusinessService.Add(one)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("新增存款一对一业务信息成功", one)
}

// 更新信息
// by author Jason
// by time 2016-12-06 16:59:09
func (this *DpOneBusinessAction) Put() util.RstMsg {
	var one = new(dp.DpOneBusiness)
	err := currentMsg.DecodeJson(this.Ctx, one)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	dpOneBusinessService.Ctx = this.Ctx
	err = dpOneBusinessService.Update(one)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("更新存款一对一业务信息成功", one)
}

// 删除信息
// by author Jason
// by time 2016-12-06 16:59:09
func (this *DpOneBusinessAction) Delete() util.RstMsg {
	var one = new(dp.DpOneBusiness)
	err := this.DecodeJson(one)
	if nil != err {
		er := fmt.Errorf("json转换实体出错")
		zlog.Error(er.Error(), err)
		return util.ErrorMsg(er.Error(), err)
	}
	err = dpOneBusinessService.Delete(one)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("删除存款一对一业务信息成功", one)
}
