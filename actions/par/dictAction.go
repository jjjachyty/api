package par

import (
	"github.com/lunny/tango"
	"pccqcpa.com.cn/app/rpm/api/models/par"
	"pccqcpa.com.cn/app/rpm/api/services/parService"
	"pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/app/rpm/api/util/currentMsg"
	"pccqcpa.com.cn/components/zlog"
)

type DictAction struct {
	tango.Json
	tango.Ctx
}

var dictService parService.DictService

/**
 * @api {get} /dict 分页查询数据字典
 * @apiName ListDict
 * @apiGroup Dict
 *
 * @apiVersion 1.0.0
 *
 * @apiUse Dict
 *
 */
func (d *DictAction) List() util.RstMsg {
	paramMap := util.GetPageMsg(&d.Ctx)
	pageData, err := dictService.List(paramMap)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessListMsg("查询字典分页成功", pageData)
}

// url : /api/rpm/dict/(*param)
/**
 * @api {get} /dict/(*param) 多参数查询数据字典
 * @apiName GetDict
 * @apiGroup Dict
 *
 * @apiVersion 1.0.0
 *
 * @apiUse Dict
 *
 */
func (d *DictAction) Get() util.RstMsg {
	paramMap, err := util.GetParmFromRouter(&d.Ctx)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	if util.IsPaginQuery(&d.Ctx) {
		// 分页查询
		util.GetPageMsg(&d.Ctx, paramMap)
		pageData, err := dictService.List(paramMap)
		if nil != err {
			return util.ErrorMsg(err.Error(), err)
		}
		return util.SuccessListMsg("多参数查询字典分页成功", pageData)
	}
	dicts, err := dictService.Find(paramMap)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("查询数据字典成功", dicts)

}

/**
 * @api {post} /dict 新增数据字典
 * @apiName PostDict
 * @apiGroup Dict
 *
 * @apiVersion 1.0.0
 *
 * @apiUse Dict
 *
 */
func (d *DictAction) Post() util.RstMsg {
	var dict = new(par.Dict)
	err := currentMsg.DecodeJson(d.Ctx, dict)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	err = dictService.Add(dict)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("新增字典成功", dict)
}

/**
 * @api {put} /dict 更新数据字典
 * @apiName PutDict
 * @apiGroup Dict
 *
 * @apiVersion 1.0.0
 *
 * @apiUse Dict
 *
 */
func (d *DictAction) Put() util.RstMsg {
	var dict = new(par.Dict)
	err := currentMsg.DecodeJson(d.Ctx, dict)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	err = dictService.Update(dict)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("更新字典成功", dict)
}

/**
 * @api {delete} /dict 删除数据字典
 * @apiName DeleteDict
 * @apiGroup Dict
 *
 * @apiVersion 1.0.0
 *
 * @apiUse Dict
 *
 */
func (d *DictAction) Delete() util.RstMsg {
	var dict = new(par.Dict)
	err := d.DecodeJson(dict)
	if nil != err {
		zlog.Error("不能转换为实体", err)
		return util.ErrorMsg("不能转换为实体", err)
	}
	err = dictService.Delete(dict)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("删除字典成功", dict)
}
