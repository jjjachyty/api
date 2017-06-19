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

var rocService parService.RocService

type RocAction struct {
	tango.Json
	tango.Ctx
}

// 分页查询
// by author Yeqc
// by time 2016-12-21 13:39:33
/**
 * @api {get} /roc 分页查询资本回报率
 * @apiName ListRoc
 * @apiGroup Roc
 *
 * @apiVersion 1.0.0
 *
 * @apiUse Roc
 *
 */
func (this *RocAction) List() util.RstMsg {
	paramMap := util.GetPageMsg(&this.Ctx)
	pageData, err := rocService.List(paramMap)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessListMsg("分页查询信息成功", pageData)
}

// 多参数查询
// by author Yeqc
// by time 2016-12-21 13:39:33
/**
 * @api {get} /roc/(*param) 多参数查询资本回报率
 * @apiName GetRoc
 * @apiGroup Roc
 *
 * @apiVersion 1.0.0
 *
 * @apiUse Roc
 *
 */
func (this *RocAction) Get() util.RstMsg {
	paramMap, err := util.GetParmFromRouter(&this.Ctx)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	if util.IsPaginQuery(&this.Ctx) {
		util.GetPageMsg(&this.Ctx, paramMap)
		pageData, err := rocService.List(paramMap)
		if nil != err {
			return util.ErrorMsg(err.Error(), err)
		}
		return util.SuccessListMsg("分页查询信息成功", pageData)
	}

	rst, err := rocService.Find(paramMap)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("查询信息成功", rst)
}

// 新增信息
// by author Yeqc
// by time 2016-12-21 13:39:33
/**
 * @api {post} /roc 新增资本回报率
 * @apiName PostRoc
 * @apiGroup Roc
 *
 * @apiVersion 1.0.0
 *
 * @apiUse Roc
 *
 */
func (this *RocAction) Post() util.RstMsg {
	var one = new(par.Roc)
	err := currentMsg.DecodeJson(this.Ctx, one)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	err = rocService.Add(one)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("新增信息成功", one)
}

// 更新信息
// by author Yeqc
// by time 2016-12-21 13:39:33
/**
 * @api {put} /roc 更新资本回报率
 * @apiName PutRoc
 * @apiGroup Roc
 *
 * @apiVersion 1.0.0
 *
 * @apiUse Roc
 *
 */
func (this *RocAction) Put() util.RstMsg {
	var one = new(par.Roc)
	err := currentMsg.DecodeJson(this.Ctx, one)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	err = rocService.Update(one)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("更新信息成功", one)
}

// 删除信息
// by author Yeqc
// by time 2016-12-21 13:39:33
/**
 * @api {delete} /roc 删除资本回报率
 * @apiName DeleteRoc
 * @apiGroup Roc
 *
 * @apiVersion 1.0.0
 *
 * @apiUse Roc
 *
 */
func (this *RocAction) Delete() util.RstMsg {
	var one = new(par.Roc)
	err := this.DecodeJson(one)
	if nil != err {
		er := fmt.Errorf("json转换实体出错")
		zlog.Error(er.Error(), err)
		return util.ErrorMsg(er.Error(), err)
	}
	err = rocService.Delete(one)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("删除信息成功", one)
}
