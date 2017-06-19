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

var pdService parService.PdService

type PdAction struct {
	tango.Json
	tango.Ctx
}

// 分页查询
// by author Yeqc
// by time 2016-12-21 09:52:26
/**
 * @api {get} /pd 分页查询违约概率
 * @apiName ListPd
 * @apiGroup Pd
 *
 * @apiVersion 1.0.0
 *
 * @apiUse Pd
 *
 */
func (this *PdAction) List() util.RstMsg {
	paramMap := util.GetPageMsg(&this.Ctx)
	pageData, err := pdService.List(paramMap)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessListMsg("分页查询信息成功", pageData)
}

// 多参数查询
// by author Yeqc
// by time 2016-12-21 09:52:26
/**
 * @api {get} /pd/(*param) 多参数查询违约概率
 * @apiName GetPd
 * @apiGroup Pd
 *
 * @apiVersion 1.0.0
 *
 * @apiUse Pd
 *
 */
func (this *PdAction) Get() util.RstMsg {
	paramMap, err := util.GetParmFromRouter(&this.Ctx)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	if util.IsPaginQuery(&this.Ctx) {
		util.GetPageMsg(&this.Ctx, paramMap)
		pageData, err := pdService.List(paramMap)
		if nil != err {
			return util.ErrorMsg(err.Error(), err)
		}
		return util.SuccessListMsg("分页查询信息成功", pageData)
	}

	rst, err := pdService.Find(paramMap)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("查询信息成功", rst)
}

// 新增信息
// by author Yeqc
// by time 2016-12-21 09:52:26
/**
 * @api {post} /pd 新增违约概率
 * @apiName PostPd
 * @apiGroup Pd
 *
 * @apiVersion 1.0.0
 *
 * @apiUse Pd
 *
 */
func (this *PdAction) Post() util.RstMsg {
	var one = new(par.Pd)
	err := currentMsg.DecodeJson(this.Ctx, one)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	err = pdService.Add(one)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("新增信息成功", one)
}

// 更新信息
// by author Yeqc
// by time 2016-12-21 09:52:26
/**
 * @api {put} /pd 更新违约概率
 * @apiName PutPd
 * @apiGroup Pd
 *
 * @apiVersion 1.0.0
 *
 * @apiUse Pd
 *
 */
func (this *PdAction) Put() util.RstMsg {
	var one = new(par.Pd)
	err := currentMsg.DecodeJson(this.Ctx, one)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	err = pdService.Update(one)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("更新信息成功", one)
}

// 删除信息
// by author Yeqc
// by time 2016-12-21 09:52:26
/**
 * @api {delete} /pd 删除违约概率
 * @apiName DeletePd
 * @apiGroup Pd
 *
 * @apiVersion 1.0.0
 *
 * @apiUse Pd
 *
 */
func (this *PdAction) Delete() util.RstMsg {
	var one = new(par.Pd)
	err := this.DecodeJson(one)
	if nil != err {
		er := fmt.Errorf("json转换实体出错")
		zlog.Error(er.Error(), err)
		return util.ErrorMsg(er.Error(), err)
	}
	err = pdService.Delete(one)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("删除信息成功", one)
}
