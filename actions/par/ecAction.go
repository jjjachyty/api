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

var ecService parService.EcService

type EcAction struct {
	tango.Json
	tango.Ctx
}

// 分页查询
// by author Yeqc
// by time 2016-12-19 10:21:15
/**
 * @api {get} /ec 分页查询经济资本占用率
 * @apiName ListEc
 * @apiGroup Ec
 *
 * @apiVersion 1.0.0
 *
 * @apiUse Ec
 *
 */
func (this *EcAction) List() util.RstMsg {
	paramMap := util.GetPageMsg(&this.Ctx)
	pageData, err := ecService.List(paramMap)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessListMsg("分页查询信息成功", pageData)
}

// 多参数查询
// by author Yeqc
// by time 2016-12-19 10:21:15
/**
 * @api {get} /ec/(*param) 多参数查询经济资本占用率
 * @apiName GetEc
 * @apiGroup Ec
 *
 * @apiVersion 1.0.0
 *
 * @apiUse Ec
 *
 */
func (this *EcAction) Get() util.RstMsg {
	paramMap, err := util.GetParmFromRouter(&this.Ctx)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	if util.IsPaginQuery(&this.Ctx) {
		util.GetPageMsg(&this.Ctx, paramMap)
		pageData, err := ecService.List(paramMap)
		if nil != err {
			return util.ErrorMsg(err.Error(), err)
		}
		return util.SuccessListMsg("分页查询信息成功", pageData)
	}

	rst, err := ecService.Find(paramMap)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("查询信息成功", rst)
}

// 新增信息
// by author Yeqc
// by time 2016-12-19 10:21:15
/**
 * @api {post} /ec 新增经济资本占用率
 * @apiName PostEc
 * @apiGroup Ec
 *
 * @apiVersion 1.0.0
 *
 * @apiUse Ec
 *
 */
func (this *EcAction) Post() util.RstMsg {
	var one = new(par.Ec)
	err := currentMsg.DecodeJson(this.Ctx, one)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	err = ecService.Add(one)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("新增信息成功", one)
}

// 更新信息
// by author Yeqc
// by time 2016-12-19 10:21:15
/**
 * @api {put} /ec 更新经济资本占用率
 * @apiName PutEc
 * @apiGroup Ec
 *
 * @apiVersion 1.0.0
 *
 * @apiUse Ec
 *
 */
func (this *EcAction) Put() util.RstMsg {
	var one = new(par.Ec)
	err := currentMsg.DecodeJson(this.Ctx, one)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	err = ecService.Update(one)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("更新信息成功", one)
}

// 删除信息
// by author Yeqc
// by time 2016-12-19 10:21:15
/**
 * @api {delete} /ec 删除经济资本占用率
 * @apiName DeleteEc
 * @apiGroup Ec
 *
 * @apiVersion 1.0.0
 *
 * @apiUse Ec
 *
 */
func (this *EcAction) Delete() util.RstMsg {
	var one = new(par.Ec)
	err := this.DecodeJson(one)
	if nil != err {
		er := fmt.Errorf("json转换实体出错")
		zlog.Error(er.Error(), err)
		return util.ErrorMsg(er.Error(), err)
	}
	err = ecService.Delete(one)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("删除信息成功", one)
}
