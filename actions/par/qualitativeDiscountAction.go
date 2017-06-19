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

var qualitativeDiscountService parService.QualitativeDiscountService

type QualitativeDiscountAction struct {
	tango.Json
	tango.Ctx
}

// 分页查询
// by author Yeqc
// by time 2016-12-21 10:20:03
/**
 * @api {get} /qualitativeDiscount 分页查询定性优惠点数
 * @apiName ListQualitativeDiscount
 * @apiGroup QualitativeDiscount
 *
 * @apiVersion 1.0.0
 *
 * @apiUse QualitativeDiscount
 *
 */
func (this *QualitativeDiscountAction) List() util.RstMsg {
	paramMap := util.GetPageMsg(&this.Ctx)
	pageData, err := qualitativeDiscountService.List(paramMap)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessListMsg("分页查询信息成功", pageData)
}

// 多参数查询
// by author Yeqc
// by time 2016-12-21 10:20:03
/**
 * @api {get} /qualitativeDiscount/(*param) 多参数查询定性优惠点数
 * @apiName GetQualitativeDiscount
 * @apiGroup QualitativeDiscount
 *
 * @apiVersion 1.0.0
 *
 * @apiUse QualitativeDiscount
 *
 */
func (this *QualitativeDiscountAction) Get() util.RstMsg {
	paramMap, err := util.GetParmFromRouter(&this.Ctx)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	if util.IsPaginQuery(&this.Ctx) {
		util.GetPageMsg(&this.Ctx, paramMap)
		pageData, err := qualitativeDiscountService.List(paramMap)
		if nil != err {
			return util.ErrorMsg(err.Error(), err)
		}
		return util.SuccessListMsg("分页查询信息成功", pageData)
	}

	rst, err := qualitativeDiscountService.Find(paramMap)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("查询信息成功", rst)
}

// 新增信息
// by author Yeqc
// by time 2016-12-21 10:20:03
/**
 * @api {post} /qualitativeDiscount 新增定性优惠点数
 * @apiName PostQualitativeDiscount
 * @apiGroup QualitativeDiscount
 *
 * @apiVersion 1.0.0
 *
 * @apiUse QualitativeDiscount
 *
 */
func (this *QualitativeDiscountAction) Post() util.RstMsg {
	var one = new(par.QualitativeDiscount)
	err := currentMsg.DecodeJson(this.Ctx, one)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	err = qualitativeDiscountService.Add(one)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("新增信息成功", one)
}

// 更新信息
// by author Yeqc
// by time 2016-12-21 10:20:03
/**
 * @api {put} /qualitativeDiscount 更新定性优惠点数
 * @apiName PutQualitativeDiscount
 * @apiGroup QualitativeDiscount
 *
 * @apiVersion 1.0.0
 *
 * @apiUse QualitativeDiscount
 *
 */
func (this *QualitativeDiscountAction) Put() util.RstMsg {
	var one = new(par.QualitativeDiscount)
	err := currentMsg.DecodeJson(this.Ctx, one)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	err = qualitativeDiscountService.Update(one)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("更新信息成功", one)
}

// 删除信息
// by author Yeqc
// by time 2016-12-21 10:20:03
/**
 * @api {delete} /qualitativeDiscount 删除定性优惠点数
 * @apiName DeleteQualitativeDiscount
 * @apiGroup QualitativeDiscount
 *
 * @apiVersion 1.0.0
 *
 * @apiUse QualitativeDiscount
 *
 */
func (this *QualitativeDiscountAction) Delete() util.RstMsg {
	var one = new(par.QualitativeDiscount)
	err := this.DecodeJson(one)
	if nil != err {
		er := fmt.Errorf("json转换实体出错")
		zlog.Error(er.Error(), err)
		return util.ErrorMsg(er.Error(), err)
	}
	err = qualitativeDiscountService.Delete(one)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("删除信息成功", one)
}
