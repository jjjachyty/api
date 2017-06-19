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

var dpOpService parService.DpOpService

type DpOpAction struct {
	tango.Json
	tango.Ctx
}

// 分页查询
// by author Yeqc
// by time 2016-12-28 10:37:13
/**
 * @api {get} /dpOp 分页查询存款操作风险率
 * @apiName ListDpOp
 * @apiGroup DpOp
 *
 * @apiVersion 1.0.0
 *
 * @apiUse DpOp
 *
 */
func (this *DpOpAction) List() util.RstMsg {
	paramMap := util.GetPageMsg(&this.Ctx)
	pageData, err := dpOpService.List(paramMap)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessListMsg("分页查询信息成功", pageData)
}

// 多参数查询
// by author Yeqc
// by time 2016-12-28 10:37:13
/**
 * @api {get} /dpOp/(*param) 多参数查询存款操作风险率
 * @apiName GetDpOp
 * @apiGroup DpOp
 *
 * @apiVersion 1.0.0
 *
 * @apiUse DpOp
 *
 */
func (this *DpOpAction) Get() util.RstMsg {
	paramMap, err := util.GetParmFromRouter(&this.Ctx)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	if util.IsPaginQuery(&this.Ctx) {
		util.GetPageMsg(&this.Ctx, paramMap)
		pageData, err := dpOpService.List(paramMap)
		if nil != err {
			return util.ErrorMsg(err.Error(), err)
		}
		return util.SuccessListMsg("分页查询信息成功", pageData)
	}

	rst, err := dpOpService.Find(paramMap)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("查询信息成功", rst)
}

// 新增信息
// by author Yeqc
// by time 2016-12-28 10:37:13
/**
 * @api {post} /dpOp 新增存款操作风险率
 * @apiName PostDpOp
 * @apiGroup DpOp
 *
 * @apiVersion 1.0.0
 *
 * @apiUse DpOp
 *
 */
func (this *DpOpAction) Post() util.RstMsg {
	var one = new(par.DpOp)
	err := currentMsg.DecodeJson(this.Ctx, one)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	err = dpOpService.Add(one)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("新增信息成功", one)
}

// 更新信息
// by author Yeqc
// by time 2016-12-28 10:37:13
/**
 * @api {put} /dpOp 更新存款操作风险率
 * @apiName PutDpOp
 * @apiGroup DpOp
 *
 * @apiVersion 1.0.0
 *
 * @apiUse DpOp
 *
 */
func (this *DpOpAction) Put() util.RstMsg {
	var one = new(par.DpOp)
	err := currentMsg.DecodeJson(this.Ctx, one)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	err = dpOpService.Update(one)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("更新信息成功", one)
}

// 删除信息
// by author Yeqc
// by time 2016-12-28 10:37:13
/**
 * @api {delete} /dpOp 删除存款操作风险率
 * @apiName DeleteDpOp
 * @apiGroup DpOp
 *
 * @apiVersion 1.0.0
 *
 * @apiUse DpOp
 *
 */
func (this *DpOpAction) Delete() util.RstMsg {
	var one = new(par.DpOp)
	err := this.DecodeJson(one)
	if nil != err {
		er := fmt.Errorf("json转换实体出错")
		zlog.Error(er.Error(), err)
		return util.ErrorMsg(er.Error(), err)
	}
	err = dpOpService.Delete(one)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("删除信息成功", one)
}
