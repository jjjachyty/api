package par

import (
	"fmt"

	"github.com/lunny/tango"

	"pccqcpa.com.cn/app/rpm/api/models/par"

	"pccqcpa.com.cn/app/rpm/api/services/parService"

	"pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/components/zlog"
)

var rlPdService parService.RlPdService

type RlPdAction struct {
	tango.Json
	tango.Ctx
}

// 分页查询
// by author Yeqc
// by time 2016-12-30 11:03:34
/**
 * @api {get} /rlPd 分页查询零售违约概率
 * @apiName ListRlPd
 * @apiGroup RlPd
 *
 * @apiVersion 1.0.0
 *
 * @apiUse RlPd
 *
 */
func (this *RlPdAction) List() util.RstMsg {
	paramMap := util.GetPageMsg(&this.Ctx)
	pageData, err := rlPdService.List(paramMap)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessListMsg("分页查询信息成功", pageData)
}

// 多参数查询
// by author Yeqc
// by time 2016-12-30 11:03:34
/**
 * @api {get} /rlPd/(*param) 多参数查询零售违约概率
 * @apiName GetRlPd
 * @apiGroup RlPd
 *
 * @apiVersion 1.0.0
 *
 * @apiUse RlPd
 *
 */
func (this *RlPdAction) Get() util.RstMsg {
	paramMap, err := util.GetParmFromRouter(&this.Ctx)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	if util.IsPaginQuery(&this.Ctx) {
		util.GetPageMsg(&this.Ctx, paramMap)
		pageData, err := rlPdService.List(paramMap)
		if nil != err {
			return util.ErrorMsg(err.Error(), err)
		}
		return util.SuccessListMsg("分页查询信息成功", pageData)
	}

	rst, err := rlPdService.Find(paramMap)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("查询信息成功", rst)
}

// 新增信息
// by author Yeqc
// by time 2016-12-30 11:03:34
/**
 * @api {post} /rlPd 新增零售违约概率
 * @apiName PostRlPd
 * @apiGroup RlPd
 *
 * @apiVersion 1.0.0
 *
 * @apiUse RlPd
 *
 */
func (this *RlPdAction) Post() util.RstMsg {
	var one = new(par.RlPd)
	err := this.DecodeJson(one)
	if nil != err {
		er := fmt.Errorf("json转换实体出错")
		zlog.Error(er.Error(), err)
		return util.ErrorMsg(er.Error(), err)
	}
	err = rlPdService.Add(one)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("新增信息成功", one)
}

// 更新信息
// by author Yeqc
// by time 2016-12-30 11:03:34
/**
 * @api {put} /rlPd 更新零售违约概率
 * @apiName PutRlPd
 * @apiGroup RlPd
 *
 * @apiVersion 1.0.0
 *
 * @apiUse RlPd
 *
 */
func (this *RlPdAction) Put() util.RstMsg {
	var one = new(par.RlPd)
	err := this.DecodeJson(one)
	if nil != err {
		er := fmt.Errorf("json转换实体出错")
		zlog.Error(er.Error(), err)
		return util.ErrorMsg(er.Error(), err)
	}
	err = rlPdService.Update(one)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("更新信息成功", one)
}

// 删除信息
// by author Yeqc
// by time 2016-12-30 11:03:34
/**
 * @api {delete} /rlPd 删除零售违约概率
 * @apiName DeleteRlPd
 * @apiGroup RlPd
 *
 * @apiVersion 1.0.0
 *
 * @apiUse RlPd
 *
 */
func (this *RlPdAction) Delete() util.RstMsg {
	var one = new(par.RlPd)
	err := this.DecodeJson(one)
	if nil != err {
		er := fmt.Errorf("json转换实体出错")
		zlog.Error(er.Error(), err)
		return util.ErrorMsg(er.Error(), err)
	}
	err = rlPdService.Delete(one)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("删除信息成功", one)
}
