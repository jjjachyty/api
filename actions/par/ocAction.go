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

var ocService parService.OcService

type OcAction struct {
	tango.Json
	tango.Ctx
}

// 分页查询
// by author Yeqc
// by time 2016-12-19 10:21:15
/**
 * @api {get} /oc 分页查询运营成本率
 * @apiName ListOc
 * @apiGroup Oc
 *
 * @apiVersion 1.0.0
 *
 * @apiUse Oc
 *
 */
func (this *OcAction) List() util.RstMsg {
	paramMap := util.GetPageMsg(&this.Ctx)
	pageData, err := ocService.List(paramMap)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessListMsg("分页查询信息成功", pageData)
}

// 多参数查询
// by author Yeqc
// by time 2016-12-19 10:21:15
/**
 * @api {get} /oc/(*param) 多参数查询运营成本率
 * @apiName GetOc
 * @apiGroup Oc
 *
 * @apiVersion 1.0.0
 *
 * @apiUse Oc
 *
 */
func (this *OcAction) Get() util.RstMsg {
	paramMap, err := util.GetParmFromRouter(&this.Ctx)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	if util.IsPaginQuery(&this.Ctx) {
		util.GetPageMsg(&this.Ctx, paramMap)
		pageData, err := ocService.List(paramMap)
		if nil != err {
			return util.ErrorMsg(err.Error(), err)
		}
		return util.SuccessListMsg("分页查询信息成功", pageData)
	}

	rst, err := ocService.Find(paramMap)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("查询信息成功", rst)
}

// 新增信息
// by author Yeqc
// by time 2016-12-19 10:21:15
/**
 * @api {post} /oc 新增运营成本率
 * @apiName PostOc
 * @apiGroup Oc
 *
 * @apiVersion 1.0.0
 *
 * @apiUse Oc
 *
 */
func (this *OcAction) Post() util.RstMsg {
	var one = new(par.Oc)
	err := currentMsg.DecodeJson(this.Ctx, one)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	err = ocService.Add(one)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("新增信息成功", one)
}

// 更新信息
// by author Yeqc
// by time 2016-12-19 10:21:15
/**
 * @api {put} /oc 更新运营成本率
 * @apiName PutOc
 * @apiGroup Oc
 *
 * @apiVersion 1.0.0
 *
 * @apiUse Oc
 *
 */
func (this *OcAction) Put() util.RstMsg {
	var one = new(par.Oc)
	err := currentMsg.DecodeJson(this.Ctx, one)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	err = ocService.Update(one)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("更新信息成功", one)
}

// 删除信息
// by author Yeqc
// by time 2016-12-19 10:21:15
/**
 * @api {delete} /oc 删除运营成本率
 * @apiName DeleteOc
 * @apiGroup Oc
 *
 * @apiVersion 1.0.0
 *
 * @apiUse Oc
 *
 */
func (this *OcAction) Delete() util.RstMsg {
	var one = new(par.Oc)
	err := this.DecodeJson(one)
	if nil != err {
		er := fmt.Errorf("json转换实体出错")
		zlog.Error(er.Error(), err)
		return util.ErrorMsg(er.Error(), err)
	}
	err = ocService.Delete(one)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("删除信息成功", one)
}
