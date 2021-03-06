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

var dpOneStockBusinessService dpService.DpOneStockBusinessService

type DpOneStockBusinessAction struct {
	tango.Json
	tango.Ctx
}

// 分页查询
// by author Jason
// by time 2017-01-04 10:46:11
func (this *DpOneStockBusinessAction) List() util.RstMsg {
	paramMap := util.GetPageMsg(&this.Ctx)
	pageData, err := dpOneStockBusinessService.List(paramMap)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessListMsg("分页查询信息成功", pageData)
}

// 多参数查询
// by author Jason
// by time 2017-01-04 10:46:11
func (this *DpOneStockBusinessAction) Get() util.RstMsg {
	paramMap, err := util.GetParmFromRouter(&this.Ctx)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	if util.IsPaginQuery(&this.Ctx) {
		util.GetPageMsg(&this.Ctx, paramMap)
		pageData, err := dpOneStockBusinessService.List(paramMap)
		if nil != err {
			return util.ErrorMsg(err.Error(), err)
		}
		return util.SuccessListMsg("分页查询信息成功", pageData)
	}

	rst, err := dpOneStockBusinessService.Find(paramMap)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("查询信息成功", rst)
}

// 新增信息
// by author Jason
// by time 2017-01-04 10:46:11
func (this *DpOneStockBusinessAction) Post() util.RstMsg {
	var one = new(dp.DpOneStockBusiness)
	err := currentMsg.DecodeJson(this.Ctx, one)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	err = dpOneStockBusinessService.Add(one)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("新增信息成功", one)
}

// 更新信息
// by author Jason
// by time 2017-01-04 10:46:11
func (this *DpOneStockBusinessAction) Put() util.RstMsg {
	var one = new(dp.DpOneStockBusiness)
	err := currentMsg.DecodeJson(this.Ctx, one)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	err = dpOneStockBusinessService.Update(one)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("更新信息成功", one)
}

// 删除信息
// by author Jason
// by time 2017-01-04 10:46:11
func (this *DpOneStockBusinessAction) Delete() util.RstMsg {
	var one = new(dp.DpOneStockBusiness)
	err := this.DecodeJson(one)
	if nil != err {
		er := fmt.Errorf("json转换实体出错")
		zlog.Error(er.Error(), err)
		return util.ErrorMsg(er.Error(), err)
	}
	err = dpOneStockBusinessService.Delete(one)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("删除信息成功", one)
}
