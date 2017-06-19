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

type SceneltdYieldAction struct {
	tango.Ctx
	tango.Json
}

var sceneltdYieldService parService.SceneItdYieldService

// 分页查询
// by author Yeqc
// by time 2016-12-21 10:30:46
/**
 * @api {get} /scene 分页查询派生中间收入收益率
 * @apiName ListSceneItdYield
 * @apiGroup SceneItdYield
 *
 * @apiVersion 1.0.0
 *
 * @apiUse SceneItdYield
 *
 */
func (this *SceneltdYieldAction) List() util.RstMsg {
	paramMap := util.GetPageMsg(&this.Ctx)
	pageData, err := sceneltdYieldService.List(paramMap)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessListMsg("分页查询信息成功", pageData)
}

/**
 * @api {get} /scene/(*param) 多参数查询派生中间收入收益率
 * @apiName GetSceneItdYield
 * @apiGroup SceneItdYield
 *
 * @apiVersion 1.0.0
 *
 * @apiUse SceneItdYield
 *
 */
func (s *SceneltdYieldAction) Get() util.RstMsg {
	paramMap, err := util.GetParmFromRouter(&s.Ctx)
	if nil != err {
		zlog.Error(err.Error(), err)
		return util.ErrorMsg(err.Error(), err)
	}
	if util.IsPaginQuery(&s.Ctx) {
		util.GetPageMsg(&s.Ctx, paramMap)
		pageData, err := sceneltdYieldService.List(paramMap)
		if nil != err {
			return util.ErrorMsg(err.Error(), err)
		}
		return util.SuccessListMsg("分页查询信息成功", pageData)
	}
	//多参数查询
	sceneltdYields, err := sceneltdYieldService.Find(paramMap)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("多参数查询指标成功", sceneltdYields)
}

// 新增信息
// by author Yeqc
// by time 2016-12-21 10:30:46
/**
 * @api {post} /scene 新增派生中间收入收益率
 * @apiName PostSceneItdYield
 * @apiGroup SceneItdYield
 *
 * @apiVersion 1.0.0
 *
 * @apiUse SceneItdYield
 *
 */
func (this *SceneltdYieldAction) Post() util.RstMsg {
	var one = new(par.SceneItdYield)
	err := currentMsg.DecodeJson(this.Ctx, one)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	err = sceneltdYieldService.Add(one)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("新增信息成功", one)
}

// 更新信息
// by author Yeqc
// by time 2016-12-21 10:30:46
/**
 * @api {put} /scene 更新派生中间收入收益率
 * @apiName PutSceneItdYield
 * @apiGroup SceneItdYield
 *
 * @apiVersion 1.0.0
 *
 * @apiUse SceneItdYield
 *
 */
func (this *SceneltdYieldAction) Put() util.RstMsg {
	var one = new(par.SceneItdYield)
	err := currentMsg.DecodeJson(this.Ctx, one)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	err = sceneltdYieldService.Update(one)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("更新信息成功", one)
}

// 删除信息
// by author Yeqc
// by time 2016-12-21 10:30:46
/**
 * @api {delete} /scene 删除派生中间收入收益率
 * @apiName DeleteSceneItdYield
 * @apiGroup SceneItdYield
 *
 * @apiVersion 1.0.0
 *
 * @apiUse SceneItdYield
 *
 */
func (this *SceneltdYieldAction) Delete() util.RstMsg {
	var one = new(par.SceneItdYield)
	err := this.DecodeJson(one)
	if nil != err {
		er := fmt.Errorf("json转换实体出错")
		zlog.Error(er.Error(), err)
		return util.ErrorMsg(er.Error(), err)
	}
	err = sceneltdYieldService.Delete(one)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("删除信息成功", one)
}
