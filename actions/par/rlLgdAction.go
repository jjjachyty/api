package par

import (
	"fmt"

	"github.com/lunny/tango"

	"pccqcpa.com.cn/app/rpm/api/models/par"

	"pccqcpa.com.cn/app/rpm/api/services/parService"

	"pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/components/zlog"
)

var rlLgdService parService.RlLgdService

type RlLgdAction struct {
	tango.Json
	tango.Ctx
}

// 分页查询
// by author Yeqc
// by time 2016-12-30 10:06:36
/**
 * @api {get} /rlLgd 分页查询零售违约损失率
 * @apiName ListRlLgd
 * @apiGroup RlLgd
 *
 * @apiVersion 1.0.0
 *
 * @apiUse RlLgd
 *
 */
func (this *RlLgdAction) List() util.RstMsg {
	paramMap := util.GetPageMsg(&this.Ctx)
	pageData, err := rlLgdService.List(paramMap)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessListMsg("分页查询信息成功", pageData)
}

// 多参数查询
// by author Yeqc
// by time 2016-12-30 10:06:36
/**
 * @api {get} /rlLgd/(*param) 多参数查询零售违约损失率
 * @apiName GetRlLgd
 * @apiGroup RlLgd
 *
 * @apiVersion 1.0.0
 *
 * @apiUse RlLgd
 *
 */
func (this *RlLgdAction) Get() util.RstMsg {
	paramMap, err := util.GetParmFromRouter(&this.Ctx)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	if util.IsPaginQuery(&this.Ctx) {
		util.GetPageMsg(&this.Ctx, paramMap)
		pageData, err := rlLgdService.List(paramMap)
		if nil != err {
			return util.ErrorMsg(err.Error(), err)
		}
		return util.SuccessListMsg("分页查询信息成功", pageData)
	}

	rst, err := rlLgdService.Find(paramMap)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("查询信息成功", rst)
}

// 新增信息
// by author Yeqc
// by time 2016-12-30 10:06:36
/**
 * @api {post} /rlLgd 新增零售违约损失率
 * @apiName PostRlLgd
 * @apiGroup RlLgd
 *
 * @apiVersion 1.0.0
 *
 * @apiUse RlLgd
 *
 */
func (this *RlLgdAction) Post() util.RstMsg {
	var one = new(par.RlLgd)
	err := this.DecodeJson(one)
	if nil != err {
		er := fmt.Errorf("json转换实体出错")
		zlog.Error(er.Error(), err)
		return util.ErrorMsg(err.Error(), err)
	}
	err = rlLgdService.Add(one)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("新增信息成功", one)
}

// 更新信息
// by author Yeqc
// by time 2016-12-30 10:06:36
/**
 * @api {put} /rlLgd 更新零售违约损失率
 * @apiName PutRlLgd
 * @apiGroup RlLgd
 *
 * @apiVersion 1.0.0
 *
 * @apiUse RlLgd
 *
 */
func (this *RlLgdAction) Put() util.RstMsg {
	var one = new(par.RlLgd)
	err := this.DecodeJson(one)
	if nil != err {
		er := fmt.Errorf("json转换实体出错")
		zlog.Error(er.Error(), err)
		return util.ErrorMsg(err.Error(), err)
	}
	err = rlLgdService.Update(one)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("更新信息成功", one)
}

// 删除信息
// by author Yeqc
// by time 2016-12-30 10:06:36
/**
 * @api {delete} /rlLgd 删除零售违约损失率
 * @apiName DeleteRlLgd
 * @apiGroup RlLgd
 *
 * @apiVersion 1.0.0
 *
 * @apiUse RlLgd
 *
 */
func (this *RlLgdAction) Delete() util.RstMsg {
	var one = new(par.RlLgd)
	err := this.DecodeJson(one)
	if nil != err {
		er := fmt.Errorf("json转换实体出错")
		zlog.Error(er.Error(), err)
		return util.ErrorMsg(er.Error(), err)
	}
	err = rlLgdService.Delete(one)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("删除信息成功", one)
}
