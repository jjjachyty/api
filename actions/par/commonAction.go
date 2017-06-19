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

var commonService parService.CommonService

type CommonAction struct {
	tango.Json
	tango.Ctx
}

/**
 * @api {get} /common 分页查询通用参数
 * @apiName ListCommon
 * @apiGroup Common
 *
 * @apiVersion 1.0.0
 *
 * @apiUse Common
 *
 */
func (this *CommonAction) List() util.RstMsg {
	paramMap := util.GetPageMsg(&this.Ctx)
	pageData, err := commonService.List(paramMap)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessListMsg("分页查询信息成功", pageData)
}

// 多参数查询
// by author Yeqc
// by time 2016-12-20 10:58:09
/**
 * @api {get} /common/(*param) 多参数查询通用参数
 * @apiName GetCommon
 * @apiGroup Common
 *
 * @apiVersion 1.0.0
 *
 * @apiUse Common
 *
 */
func (this *CommonAction) Get() util.RstMsg {
	paramMap, err := util.GetParmFromRouter(&this.Ctx)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	if util.IsPaginQuery(&this.Ctx) {
		util.GetPageMsg(&this.Ctx, paramMap)
		pageData, err := commonService.List(paramMap)
		if nil != err {
			return util.ErrorMsg(err.Error(), err)
		}
		return util.SuccessListMsg("分页查询信息成功", pageData)
	}

	rst, err := commonService.Find(paramMap)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("查询信息成功", rst)
}

// 新增信息
// by author Yeqc
// by time 2016-12-20 10:58:09
/**
 * @api {post} /common 新增通用参数
 * @apiName PostCommon
 * @apiGroup Common
 *
 * @apiVersion 1.0.0
 *
 * @apiUse Common
 *
 */
func (this *CommonAction) Post() util.RstMsg {
	var one = new(par.Common)
	err := currentMsg.DecodeJson(this.Ctx, one)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	err = commonService.Add(one)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("新增信息成功", one)
}

// 更新信息
// by author Yeqc
// by time 2016-12-20 10:58:09
/**
 * @api {put} /common 更新通用参数
 * @apiName PutCommon
 * @apiGroup Common
 *
 * @apiVersion 1.0.0
 *
 * @apiUse Common
 *
 */
func (this *CommonAction) Put() util.RstMsg {
	var one = new(par.Common)
	err := currentMsg.DecodeJson(this.Ctx, one)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	err = commonService.Update(one)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("更新信息成功", one)
}

// 删除信息
// by author Yeqc
// by time 2016-12-20 10:58:09
/**
 * @api {delete} /common 删除通用参数
 * @apiName DeleteCommon
 * @apiGroup Common
 *
 * @apiVersion 1.0.0
 *
 * @apiUse Common
 *
 */
func (this *CommonAction) Delete() util.RstMsg {
	var one = new(par.Common)
	err := this.DecodeJson(one)
	if nil != err {
		er := fmt.Errorf("json转换实体出错")
		zlog.Error(er.Error(), err)
		return util.ErrorMsg(er.Error(), err)
	}
	err = commonService.Delete(one)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("删除信息成功", one)
}
