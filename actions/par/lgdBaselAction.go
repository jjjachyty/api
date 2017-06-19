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

var lgdBaselService parService.LgdBaselService

type LgdBaselAction struct {
	tango.Json
	tango.Ctx
}

// 分页查询
// by author Yeqc
// by time 2016-12-19 10:16:29
/**
 * @api {get} /lgdBasel 分页查询违约损失率
 * @apiName ListLgdBasel
 * @apiGroup LgdBasel
 *
 * @apiVersion 1.0.0
 *
 * @apiUse LgdBasel
 *
 */
func (this *LgdBaselAction) List() util.RstMsg {
	paramMap := util.GetPageMsg(&this.Ctx)
	pageData, err := lgdBaselService.List(paramMap)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessListMsg("分页查询信息成功", pageData)
}

// 多参数查询
// by author Yeqc
// by time 2016-12-19 10:16:29
/**
 * @api {get} /lgdBasel/(*param) 多参数查询违约损失率
 * @apiName GetLgdBasel
 * @apiGroup LgdBasel
 *
 * @apiVersion 1.0.0
 *
 * @apiUse LgdBasel
 *
 */
func (this *LgdBaselAction) Get() util.RstMsg {
	paramMap, err := util.GetParmFromRouter(&this.Ctx)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	if util.IsPaginQuery(&this.Ctx) {
		util.GetPageMsg(&this.Ctx, paramMap)
		pageData, err := lgdBaselService.List(paramMap)
		if nil != err {
			return util.ErrorMsg(err.Error(), err)
		}
		return util.SuccessListMsg("分页查询信息成功", pageData)
	}

	rst, err := lgdBaselService.Find(paramMap)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("查询信息成功", rst)
}

// 新增信息
// by author Yeqc
// by time 2016-12-19 10:16:29
/**
 * @api {post} /lgdBasel 新增违约损失率
 * @apiName PostLgdBasel
 * @apiGroup LgdBasel
 *
 * @apiVersion 1.0.0
 *
 * @apiUse LgdBasel
 *
 */
func (this *LgdBaselAction) Post() util.RstMsg {
	var one = new(par.LgdBasel)
	err := currentMsg.DecodeJson(this.Ctx, one)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	err = lgdBaselService.Add(one)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("新增信息成功", one)
}

// 更新信息
// by author Yeqc
// by time 2016-12-19 10:16:29
/**
 * @api {put} /lgdBasel 更新违约损失率
 * @apiName PutLgdBasel
 * @apiGroup LgdBasel
 *
 * @apiVersion 1.0.0
 *
 * @apiUse LgdBasel
 *
 */
func (this *LgdBaselAction) Put() util.RstMsg {
	var one = new(par.LgdBasel)
	err := currentMsg.DecodeJson(this.Ctx, one)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	err = lgdBaselService.Update(one)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("更新信息成功", one)
}

// 删除信息
// by author Yeqc
// by time 2016-12-19 10:16:29
/**
 * @api {delete} /lgdBasel 删除违约损失率
 * @apiName DeleteLgdBasel
 * @apiGroup LgdBasel
 *
 * @apiVersion 1.0.0
 *
 * @apiUse LgdBasel
 *
 */
func (this *LgdBaselAction) Delete() util.RstMsg {
	var one = new(par.LgdBasel)
	err := this.DecodeJson(one)
	if nil != err {
		er := fmt.Errorf("json转换实体出错")
		zlog.Error(er.Error(), err)
		return util.ErrorMsg(er.Error(), err)
	}
	err = lgdBaselService.Delete(one)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("删除信息成功", one)
}
