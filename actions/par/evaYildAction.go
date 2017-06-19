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

var evaYieldService parService.EvaYieldService

type EvaYieldAction struct {
	tango.Json
	tango.Ctx
}

// 分页查询
// by author Yeqc
// by time 2016-12-19 15:10:36\
/**
 * @api {get} /evaYield 分页查询eva收益率
 * @apiName ListEvaYield
 * @apiGroup EvaYield
 *
 * @apiVersion 1.0.0
 *
 * @apiUse EvaYield
 *
 */
func (this *EvaYieldAction) List() util.RstMsg {
	paramMap := util.GetPageMsg(&this.Ctx)
	pageData, err := evaYieldService.List(paramMap)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessListMsg("分页查询信息成功", pageData)
}

// 多参数查询
// by author Yeqc
// by time 2016-12-19 15:10:36
/**
 * @api {get} /evaYield/(*param) 多参数查询eva收益率
 * @apiName GetEvaYield
 * @apiGroup EvaYield
 *
 * @apiVersion 1.0.0
 *
 * @apiUse EvaYield
 *
 */
func (this *EvaYieldAction) Get() util.RstMsg {
	paramMap, err := util.GetParmFromRouter(&this.Ctx)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	if util.IsPaginQuery(&this.Ctx) {
		util.GetPageMsg(&this.Ctx, paramMap)
		pageData, err := evaYieldService.List(paramMap)
		if nil != err {
			return util.ErrorMsg(err.Error(), err)
		}
		return util.SuccessListMsg("分页查询信息成功", pageData)
	}

	rst, err := evaYieldService.Find(paramMap)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("查询信息成功", rst)
}

// 新增信息
// by author Yeqc
// by time 2016-12-19 15:10:36
/**
 * @api {post} /evaYield 新增eva收益率
 * @apiName PostEvaYield
 * @apiGroup EvaYield
 *
 * @apiVersion 1.0.0
 *
 * @apiUse EvaYield
 *
 */
func (this *EvaYieldAction) Post() util.RstMsg {
	var one = new(par.EvaYield)
	err := currentMsg.DecodeJson(this.Ctx, one)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	err = evaYieldService.Add(one)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("新增信息成功", one)
}

// 更新信息
// by author Yeqc
// by time 2016-12-19 15:10:36
/**
 * @api {put} /evaYield 更新eva收益率
 * @apiName PutEvaYield
 * @apiGroup EvaYield
 *
 * @apiVersion 1.0.0
 *
 * @apiUse EvaYield
 *
 */
func (this *EvaYieldAction) Put() util.RstMsg {
	var one = new(par.EvaYield)
	err := currentMsg.DecodeJson(this.Ctx, one)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	err = evaYieldService.Update(one)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("更新信息成功", one)
}

// 删除信息
// by author Yeqc
// by time 2016-12-19 15:10:36
/**
 * @api {delete} /evaYield 删除eva收益率
 * @apiName DeleteEvaYield
 * @apiGroup EvaYield
 *
 * @apiVersion 1.0.0
 *
 * @apiUse EvaYield
 *
 */
func (this *EvaYieldAction) Delete() util.RstMsg {
	var one = new(par.EvaYield)
	err := this.DecodeJson(one)
	if nil != err {
		er := fmt.Errorf("json转换实体出错")
		zlog.Error(er.Error(), err)
		return util.ErrorMsg(er.Error(), err)
	}
	err = evaYieldService.Delete(one)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("删除信息成功", one)
}
