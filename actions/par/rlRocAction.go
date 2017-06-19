package par

import (
	"fmt"

	"github.com/lunny/tango"

	"pccqcpa.com.cn/app/rpm/api/models/par"

	"pccqcpa.com.cn/app/rpm/api/services/parService"

	"pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/components/zlog"
)

var rlRocService parService.RlRocService

type RlRocAction struct {
	tango.Json
	tango.Ctx
}

// 分页查询
// by author Yeqc
// by time 2016-12-30 11:31:37
/**
 * @api {get} /rlRoc 分页查询零售资本回报率
 * @apiName ListRlRoc
 * @apiGroup RlRoc
 *
 * @apiVersion 1.0.0
 *
 * @apiUse RlRoc
 *
 */
func (this *RlRocAction) List() util.RstMsg {
	paramMap := util.GetPageMsg(&this.Ctx)
	pageData, err := rlRocService.List(paramMap)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessListMsg("分页查询信息成功", pageData)
}

// 多参数查询
// by author Yeqc
// by time 2016-12-30 11:31:37
/**
 * @api {get} /rlRoc/(*param) 多参数查询零售资本回报率
 * @apiName GetRlRoc
 * @apiGroup RlRoc
 *
 * @apiVersion 1.0.0
 *
 * @apiUse RlRoc
 *
 */
func (this *RlRocAction) Get() util.RstMsg {
	paramMap, err := util.GetParmFromRouter(&this.Ctx)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	if util.IsPaginQuery(&this.Ctx) {
		util.GetPageMsg(&this.Ctx, paramMap)
		pageData, err := rlRocService.List(paramMap)
		if nil != err {
			return util.ErrorMsg(err.Error(), err)
		}
		return util.SuccessListMsg("分页查询信息成功", pageData)
	}

	rst, err := rlRocService.Find(paramMap)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("查询信息成功", rst)
}

// 新增信息
// by author Yeqc
// by time 2016-12-30 11:31:37
/**
 * @api {post} /rlRoc 新增零售资本回报率
 * @apiName PostRlRoc
 * @apiGroup RlRoc
 *
 * @apiVersion 1.0.0
 *
 * @apiUse RlRoc
 *
 */
func (this *RlRocAction) Post() util.RstMsg {
	var one = new(par.RlRoc)
	err := this.DecodeJson(one)
	if nil != err {
		er := fmt.Errorf("json转换实体出错")
		zlog.Error(er.Error(), err)
		return util.ErrorMsg(er.Error(), err)
	}
	err = rlRocService.Add(one)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("新增信息成功", one)
}

// 更新信息
// by author Yeqc
// by time 2016-12-30 11:31:37
/**
 * @api {put} /rlRoc 更新零售资本回报率
 * @apiName PutRlRoc
 * @apiGroup RlRoc
 *
 * @apiVersion 1.0.0
 *
 * @apiUse RlRoc
 *
 */
func (this *RlRocAction) Put() util.RstMsg {
	var one = new(par.RlRoc)
	err := this.DecodeJson(one)
	if nil != err {
		er := fmt.Errorf("json转换实体出错")
		zlog.Error(er.Error(), err)
		return util.ErrorMsg(er.Error(), err)
	}
	err = rlRocService.Update(one)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("更新信息成功", one)
}

// 删除信息
// by author Yeqc
// by time 2016-12-30 11:31:37
/**
 * @api {delete} /rlRoc 删除零售资本回报率
 * @apiName DeleteRlRoc
 * @apiGroup RlRoc
 *
 * @apiVersion 1.0.0
 *
 * @apiUse RlRoc
 *
 */
func (this *RlRocAction) Delete() util.RstMsg {
	var one = new(par.RlRoc)
	err := this.DecodeJson(one)
	if nil != err {
		er := fmt.Errorf("json转换实体出错")
		zlog.Error(er.Error(), err)
		return util.ErrorMsg(er.Error(), err)
	}
	err = rlRocService.Delete(one)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("删除信息成功", one)
}
