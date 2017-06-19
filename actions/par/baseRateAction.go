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

var baseRateService parService.BaseRateService

type BaseRateAction struct {
	tango.Json
	tango.Ctx
}

/**
 * @api {get} /baseRate 分页查询基准利率
 * @apiName ListBaseRate
 * @apiGroup BaseRate
 *
 * @apiVersion 1.0.0
 *
 * @apiUse BaseRate
 *
 */
func (this *BaseRateAction) List() util.RstMsg {
	paramMap := util.GetPageMsg(&this.Ctx)
	pageData, err := baseRateService.List(paramMap)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessListMsg("分页查询信息成功", pageData)
}

// 多参数查询
// by author Yeqc
// by time 2016-12-19 10:49:29
/**
 * @api {get} /baseRate/(*param) 多参数查询基准利率
 * @apiName GetBaseRate
 * @apiGroup BaseRate
 *
 * @apiVersion 1.0.0
 *
 * @apiUse BaseRate
 *
 */
func (this *BaseRateAction) Get() util.RstMsg {
	paramMap, err := util.GetParmFromRouter(&this.Ctx)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	if util.IsPaginQuery(&this.Ctx) {
		util.GetPageMsg(&this.Ctx, paramMap)
		pageData, err := baseRateService.List(paramMap)
		if nil != err {
			return util.ErrorMsg(err.Error(), err)
		}
		return util.SuccessListMsg("分页查询信息成功", pageData)
	}

	rst, err := baseRateService.Find(paramMap)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("查询信息成功", rst)
}

// 新增信息
// by author Yeqc
// by time 2016-12-19 10:49:29
/**
 * @api {post} /baseRate 新增基准利率
 * @apiName PostBaseRate
 * @apiGroup BaseRate
 *
 * @apiVersion 1.0.0
 *
 * @apiUse BaseRate
 *
 */
func (this *BaseRateAction) Post() util.RstMsg {
	var one = new(par.BaseRate)
	err := currentMsg.DecodeJson(this.Ctx, one)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	err = baseRateService.Add(one)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("新增信息成功", one)
}

// 更新信息
// by author Yeqc
// by time 2016-12-19 10:49:29
/**
 * @api {put} /baseRate 更新基准利率
 * @apiName PutBaseRate
 * @apiGroup BaseRate
 *
 * @apiVersion 1.0.0
 *
 * @apiUse BaseRate
 *
 */
func (this *BaseRateAction) Put() util.RstMsg {
	var one = new(par.BaseRate)
	err := currentMsg.DecodeJson(this.Ctx, one)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	err = baseRateService.Update(one)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("更新信息成功", one)
}

// 删除信息
// by author Yeqc
// by time 2016-12-19 10:49:29
/**
 * @api {delete} /baseRate 删除基准利率
 * @apiName DeleteBaseRate
 * @apiGroup BaseRate
 *
 * @apiVersion 1.0.0
 *
 * @apiUse BaseRate
 *
 */
func (this *BaseRateAction) Delete() util.RstMsg {
	var one = new(par.BaseRate)
	err := this.DecodeJson(one)
	if nil != err {
		er := fmt.Errorf("json转换实体出错")
		zlog.Error(er.Error(), err)
		return util.ErrorMsg(er.Error(), err)
	}
	err = baseRateService.Delete(one)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("删除信息成功", one)
}
