package par

import (
	"fmt"

	"github.com/lunny/tango"

	"pccqcpa.com.cn/app/rpm/api/models/par"
	"pccqcpa.com.cn/app/rpm/api/services/parService"
	"pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/app/rpm/api/util/currentMsg"
	"pccqcpa.com.cn/components/zlog"
)

var dpBaseRateService parService.DpBaseRateService

type DpBaseRateAction struct {
	tango.Json
	tango.Ctx
}

// 分页查询
// by author Yeqc
// by time 2017-01-05 14:25:41
func (this *DpBaseRateAction) List() util.RstMsg {
	paramMap := util.GetPageMsg(&this.Ctx)
	pageData, err := dpBaseRateService.List(paramMap)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessListMsg("分页查询信息成功", pageData)
}

// 多参数查询
// by author Yeqc
// by time 2017-01-05 14:25:41
func (this *DpBaseRateAction) Get() util.RstMsg {
	paramMap, err := util.GetParmFromRouter(&this.Ctx)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	if util.IsPaginQuery(&this.Ctx) {
		util.GetPageMsg(&this.Ctx, paramMap)
		pageData, err := dpBaseRateService.List(paramMap)
		if nil != err {
			return util.ErrorMsg(err.Error(), err)
		}
		return util.SuccessListMsg("分页查询信息成功", pageData)
	}

	rst, err := dpBaseRateService.Find(paramMap)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("查询信息成功", rst)
}

// 新增信息
// by author Yeqc
// by time 2017-01-05 14:25:41
func (this *DpBaseRateAction) Post() util.RstMsg {
	var one = new(par.DpBaseRate)
	err := currentMsg.DecodeJson(this.Ctx, one)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	err = dpBaseRateService.Add(one)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("新增信息成功", one)
}

// 更新信息
// by author Yeqc
// by time 2017-01-05 14:25:41
func (this *DpBaseRateAction) Put() util.RstMsg {
	var one = new(par.DpBaseRate)
	err := currentMsg.DecodeJson(this.Ctx, one)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	err = dpBaseRateService.Update(one)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("更新信息成功", one)
}

// 删除信息
// by author Yeqc
// by time 2017-01-05 14:25:41
func (this *DpBaseRateAction) Delete() util.RstMsg {
	var one = new(par.DpBaseRate)
	err := this.DecodeJson(one)
	if nil != err {
		er := fmt.Errorf("json转换实体出错")
		zlog.Error(er.Error(), err)
		return util.ErrorMsg(er.Error(), err)
	}
	err = dpBaseRateService.Delete(one)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("删除信息成功", one)
}
