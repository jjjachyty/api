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

var rlEcService parService.RlEcService

type RlEcAction struct {
	tango.Json
	tango.Ctx
}

// 分页查询
// by author Yeqc
// by time 2016-12-30 09:32:23
/**
 * @api {get} /rlEc 分页查询零售资本成本率
 * @apiName ListRlEc
 * @apiGroup RlEc
 *
 * @apiVersion 1.0.0
 *
 * @apiUse RlEc
 *
 */
func (this *RlEcAction) List() util.RstMsg {
	paramMap := util.GetPageMsg(&this.Ctx)
	pageData, err := rlEcService.List(paramMap)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessListMsg("分页查询信息成功", pageData)
}

// 多参数查询
// by author Yeqc
// by time 2016-12-30 09:32:23
/**
 * @api {get} /rlEc/(*param) 多参数查询基准利率
 * @apiName GetRlEc
 * @apiGroup RlEc
 *
 * @apiVersion 1.0.0
 *
 * @apiUse RlEc
 *
 */
func (this *RlEcAction) Get() util.RstMsg {
	paramMap, err := util.GetParmFromRouter(&this.Ctx)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	if util.IsPaginQuery(&this.Ctx) {
		util.GetPageMsg(&this.Ctx, paramMap)
		pageData, err := rlEcService.List(paramMap)
		if nil != err {
			return util.ErrorMsg(err.Error(), err)
		}
		return util.SuccessListMsg("分页查询信息成功", pageData)
	}

	rst, err := rlEcService.Find(paramMap)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("查询信息成功", rst)
}

// 新增信息
// by author Yeqc
// by time 2016-12-30 09:32:23
/**
 * @api {post} /rlEc 新增基准利率
 * @apiName PostRlEc
 * @apiGroup RlEc
 *
 * @apiVersion 1.0.0
 *
 * @apiUse RlEc
 *
 */
func (this *RlEcAction) Post() util.RstMsg {
	var one = new(par.RlEc)
	err := currentMsg.DecodeJson(this.Ctx, one)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	err = rlEcService.Add(one)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("新增信息成功", one)
}

// 更新信息
// by author Yeqc
// by time 2016-12-30 09:32:23
/**
 * @api {put} /rlEc 更新基准利率
 * @apiName PutRlEc
 * @apiGroup RlEc
 *
 * @apiVersion 1.0.0
 *
 * @apiUse RlEc
 *
 */
func (this *RlEcAction) Put() util.RstMsg {
	var one = new(par.RlEc)
	err := currentMsg.DecodeJson(this.Ctx, one)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	err = rlEcService.Update(one)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("更新信息成功", one)
}

// 删除信息
// by author Yeqc
// by time 2016-12-30 09:32:23
/**
 * @api {delete} /rlEc 删除基准利率
 * @apiName DeleteRlEc
 * @apiGroup RlEc
 *
 * @apiVersion 1.0.0
 *
 * @apiUse RlEc
 *
 */
func (this *RlEcAction) Delete() util.RstMsg {
	var one = new(par.RlEc)
	err := this.DecodeJson(one)
	if nil != err {
		er := fmt.Errorf("json转换实体出错")
		zlog.Error(er.Error(), err)
		return util.ErrorMsg(er.Error(), err)
	}
	err = rlEcService.Delete(one)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("删除信息成功", one)
}
