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

var dpDifferService parService.DpDifferService

type DpDifferAction struct {
	tango.Json
	tango.Ctx
}

// 分页查询
// by author Jason
// by time 2016-11-30 09:31:42
/**
 * @api {get} /dp/differ 分页查询存款差异化定价参数
 * @apiName ListDpDiffer
 * @apiGroup DpDiffer
 *
 * @apiVersion 1.0.0
 *
 * @apiUse DpDiffer
 *
 */
func (this *DpDifferAction) List() util.RstMsg {
	paramMap := util.GetPageMsg(&this.Ctx)
	pageData, err := dpDifferService.List(paramMap)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessListMsg("分页查询信息成功", pageData)
}

// 多参数查询
// by author Jason
// by time 2016-11-30 09:31:42
/**
 * @api {get} /dp/differ/(*param) 多参数查询基准利率
 * @apiName GetDpDiffer
 * @apiGroup DpDiffer
 *
 * @apiVersion 1.0.0
 *
 * @apiUse DpDiffer
 *
 */
func (this *DpDifferAction) Get() util.RstMsg {
	paramMap, err := util.GetParmFromRouter(&this.Ctx)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	if util.IsPaginQuery(&this.Ctx) {
		util.GetPageMsg(&this.Ctx, paramMap)
		pageData, err := dpDifferService.List(paramMap)
		if nil != err {
			return util.ErrorMsg(err.Error(), err)
		}
		return util.SuccessListMsg("分页查询信息成功", pageData)
	}

	rst, err := dpDifferService.Find(paramMap)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("查询信息成功", rst)
}

// 新增信息
// by author Jason
// by time 2016-11-30 09:31:42
/**
 * @api {post} /dp/differ 新增基准利率
 * @apiName PostDpDiffer
 * @apiGroup DpDiffer
 *
 * @apiVersion 1.0.0
 *
 * @apiUse DpDiffer
 *
 */
func (this *DpDifferAction) Post() util.RstMsg {
	var one = new(par.DpDiffer)
	err := currentMsg.DecodeJson(this.Ctx, one)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	err = dpDifferService.Add(one)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("新增信息成功", one)
}

// 更新信息
// by author Jason
// by time 2016-11-30 09:31:42
/**
 * @api {put} /dp/differ 更新基准利率
 * @apiName PutDpDiffer
 * @apiGroup DpDiffer
 *
 * @apiVersion 1.0.0
 *
 * @apiUse DpDiffer
 *
 */
func (this *DpDifferAction) Put() util.RstMsg {
	var one = new(par.DpDiffer)
	err := currentMsg.DecodeJson(this.Ctx, one)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	err = dpDifferService.Update(one)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("更新信息成功", one)
}

// 删除信息
// by author Jason
// by time 2016-11-30 09:31:42
/**
 * @api {delete} /dp/differ 删除基准利率
 * @apiName DeleteDpDiffer
 * @apiGroup DpDiffer
 *
 * @apiVersion 1.0.0
 *
 * @apiUse DpDiffer
 *
 */
func (this *DpDifferAction) Delete() util.RstMsg {
	var one = new(par.DpDiffer)
	err := this.DecodeJson(one)
	if nil != err {
		er := fmt.Errorf("json转换实体出错")
		zlog.Error(er.Error(), err)
		return util.ErrorMsg(er.Error(), err)
	}
	err = dpDifferService.Delete(one)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("删除信息成功", one)
}
