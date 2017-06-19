package par

import (
	"github.com/lunny/tango"

	"pccqcpa.com.cn/app/rpm/api/models/par"
	"pccqcpa.com.cn/app/rpm/api/services/parService"
	"pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/app/rpm/api/util/currentMsg"
	"pccqcpa.com.cn/components/zlog"
)

type IndicatorAction struct {
	tango.Json
	tango.Ctx
}

var indicatorServer parService.IndicatorService

// 查询所有指标并分页
// url: /api/rpm/ind/indicator/
/**
 * @api {get} /indicator 分页查询指标
 * @apiName ListIndicator
 * @apiGroup Indicator
 *
 * @apiVersion 1.0.0
 *
 * @apiUse Indicator
 *
 */
func (ind *IndicatorAction) List() util.RstMsg {
	paramMap := util.GetPageMsg(&ind.Ctx)
	pageData, err := indicatorServer.List(paramMap)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessListMsg("查询所有指标分页成功", pageData)
}

// 带参数查询，判断header里面的start-row-number是否为空，如果为空，则不分页 放在service层判断
/**
 * @api {get} /indicator/(*param) 多参数查询基准利率
 * @apiName GetIndicator
 * @apiGroup Indicator
 *
 * @apiVersion 1.0.0
 *
 * @apiUse Indicator
 *
 */
func (ind *IndicatorAction) Get() util.RstMsg {

	paramMap, err := util.GetParmFromRouter(&ind.Ctx)
	if nil != err {
		zlog.Error(err.Error(), err)
		return util.ErrorMsg(err.Error(), err)
	}
	// 判断是否为分页查询
	if util.IsPaginQuery(&ind.Ctx) {
		util.GetPageMsg(&ind.Ctx, paramMap)
		pageData, err := indicatorServer.List(paramMap)
		if nil != err {
			return util.ErrorMsg(err.Error(), err)
		}
		return util.SuccessListMsg("多参数查询指标分页成功", pageData)
	}

	// 多参数查询不分页
	inds, err := indicatorServer.Find(paramMap)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("多参数查询指标成功", inds)
}

// 新增指标
// url:/api/rpm/indicator
/**
 * @api {post} /indicator 新增基准利率
 * @apiName PostIndicator
 * @apiGroup Indicator
 *
 * @apiVersion 1.0.0
 *
 * @apiUse Indicator
 *
 */
func (this *IndicatorAction) Post() interface{} {

	var indicator = new(par.Indicator)
	err := currentMsg.DecodeJson(this.Ctx, indicator)
	if nil != err {
		return util.ErrorMsg("不能转换为实体", err)
	}
	ind, err := indicatorServer.Add(indicator)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("新增指标成功", ind)
}

// 删除
// 先判断是否有指标依赖此指标，如果没有则删除，如果有则提示有其他指标依赖此指标，请先接触依赖关系
// url: /api/rpm/indicator
/**
 * @api {delete} /indicator 删除基准利率
 * @apiName DeleteIndicator
 * @apiGroup Indicator
 *
 * @apiVersion 1.0.0
 *
 * @apiUse Indicator
 *
 */
func (ind *IndicatorAction) Delete() interface{} {

	var indicator = new(par.Indicator)
	ind.DecodeJson(indicator)

	err := indicatorServer.Delete(indicator)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("删除指标成功", nil)
}

// 修改
// url: /api/rpm/indicator
/**
 * @api {put} /indicator 更新基准利率
 * @apiName PutIndicator
 * @apiGroup Indicator
 *
 * @apiVersion 1.0.0
 *
 * @apiUse Indicator
 *
 */
func (ind *IndicatorAction) Put() interface{} {

	var indicator = new(par.Indicator)
	err := currentMsg.DecodeJson(ind.Ctx, indicator)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}

	indi, err := indicatorServer.Update(indicator)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("修改指标成功", indi)
}
