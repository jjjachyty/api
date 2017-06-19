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

var dpTargetRateService parService.DpTargetRateService

type DpTargetRateAction struct {
	tango.Json
	tango.Ctx
}

// 分页查询
// by author Jason
// by time 2017-02-15 10:19:49
func(this *DpTargetRateAction) List() util.RstMsg {
	paramMap := util.GetPageMsg(&this.Ctx)
	pageData, err := dpTargetRateService.List(paramMap)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessListMsg("分页查询信息成功", pageData)
}

// 多参数查询
// by author Jason
// by time 2017-02-15 10:19:49
func(this *DpTargetRateAction) Get() util.RstMsg {
	paramMap, err := util.GetParmFromRouter(&this.Ctx)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	if util.IsPaginQuery(&this.Ctx) {
		util.GetPageMsg(&this.Ctx, paramMap)
		pageData,err := dpTargetRateService.List(paramMap)
		if nil != err {
			return util.ErrorMsg(err.Error(), err)
		}
		return util.SuccessListMsg("分页查询信息成功", pageData)
	}

	rst, err := dpTargetRateService.Find(paramMap)
	if nil != err {
			return util.ErrorMsg(err.Error(), err)
		}
	return util.SuccessMsg("查询信息成功", rst)
}

// 新增信息
// by author Jason
// by time 2017-02-15 10:19:49
func(this *DpTargetRateAction) Post() util.RstMsg {
	var one = new(par.DpTargetRate)
	err := currentMsg.DecodeJson(this.Ctx, one)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	err = dpTargetRateService.Add(one)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("新增信息成功",one)
}

// 更新信息
// by author Jason
// by time 2017-02-15 10:19:49
func(this *DpTargetRateAction) Put() util.RstMsg {
	var one = new(par.DpTargetRate)
	err := currentMsg.DecodeJson(this.Ctx, one)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	err = dpTargetRateService.Update(one)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("更新信息成功",one)
}

// 删除信息
// by author Jason
// by time 2017-02-15 10:19:49
func(this *DpTargetRateAction) Delete() util.RstMsg {
	var one = new(par.DpTargetRate)
	err := this.DecodeJson(one)
	if nil != err {
		er := fmt.Errorf("json转换实体出错")
		zlog.Error(er.Error(), err)
		return util.ErrorMsg(er.Error(), err)
	}
	err = dpTargetRateService.Delete(one)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("删除信息成功",one)
}

