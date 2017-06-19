package par

import (
	// "encoding/json"
	"fmt"

	"github.com/lunny/tango"

	"pccqcpa.com.cn/app/rpm/api/models/par"

	"pccqcpa.com.cn/app/rpm/api/services/parService"

	"pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/app/rpm/api/util/currentMsg"
	"pccqcpa.com.cn/components/zlog"
)

var taxService parService.TaxService

type TaxAction struct {
	tango.Json
	tango.Ctx
}

// 分页查询
// by author Yeqc
// by time 2016-12-21 11:10:11
/**
 * @api {get} /tax 分页查询税率
 * @apiName ListTax
 * @apiGroup Tax
 *
 * @apiVersion 1.0.0
 *
 * @apiUse Tax
 *
 */
func (this *TaxAction) List() util.RstMsg {
	paramMap := util.GetPageMsg(&this.Ctx)
	pageData, err := taxService.List(paramMap)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessListMsg("分页查询信息成功", pageData)
}

// 多参数查询
// by author Yeqc
// by time 2016-12-21 11:10:11
/**
 * @api {get} /tax/(*param) 多参数查询税率
 * @apiName GetTax
 * @apiGroup Tax
 *
 * @apiVersion 1.0.0
 *
 * @apiUse Tax
 *
 */
func (this *TaxAction) Get() util.RstMsg {
	paramMap, err := util.GetParmFromRouter(&this.Ctx)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	if util.IsPaginQuery(&this.Ctx) {
		util.GetPageMsg(&this.Ctx, paramMap)
		pageData, err := taxService.List(paramMap)
		if nil != err {
			return util.ErrorMsg(err.Error(), err)
		}
		return util.SuccessListMsg("分页查询信息成功", pageData)
	}

	rst, err := taxService.Find(paramMap)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("查询信息成功", rst)
}

// 新增信息
// by author Yeqc
// by time 2016-12-21 11:10:11
/**
 * @api {post} /tax 新增税率
 * @apiName PostTax
 * @apiGroup Tax
 *
 * @apiVersion 1.0.0
 *
 * @apiUse Tax
 *
 */
func (this *TaxAction) Post() util.RstMsg {
	var one = new(par.Tax)
	err := currentMsg.DecodeJson(this.Ctx, one)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	err = taxService.Add(one)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("新增信息成功", one)
}

// 更新信息
// by author Yeqc
// by time 2016-12-21 11:10:11
/**
 * @api {put} /tax 更新税率
 * @apiName PutTax
 * @apiGroup Tax
 *
 * @apiVersion 1.0.0
 *
 * @apiUse Tax
 *
 */
func (this *TaxAction) Put() util.RstMsg {
	var one = new(par.Tax)
	err := currentMsg.DecodeJson(this.Ctx, one)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	err = taxService.Update(one)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("更新信息成功", one)
}

// 删除信息
// by author Yeqc
// by time 2016-12-21 11:10:11
/**
 * @api {delete} /tax 删除税率
 * @apiName DeleteTax
 * @apiGroup Tax
 *
 * @apiVersion 1.0.0
 *
 * @apiUse Tax
 *
 */
func (this *TaxAction) Delete() util.RstMsg {
	var one = new(par.Tax)
	err := this.DecodeJson(one)
	if nil != err {
		er := fmt.Errorf("json转换实体出错")
		zlog.Error(er.Error(), err)
		return util.ErrorMsg(er.Error(), err)
	}
	err = taxService.Delete(one)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("删除信息成功", one)
}
