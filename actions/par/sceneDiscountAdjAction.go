package par

import (
	"github.com/lunny/tango"

	"pccqcpa.com.cn/app/rpm/api/models/par"
	"pccqcpa.com.cn/app/rpm/api/services/parService"
	"pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/app/rpm/api/util/currentMsg"
	"pccqcpa.com.cn/components/zlog"
)

type SceneDiscountAdjAction struct {
	tango.Ctx
	tango.Json
}

var sceneDiscountAdjService parService.SceneDiscountAdjService

/**
 * @api {get} /scenediscountadj 分页查询派生调节系数
 * @apiName ListSceneDiscountAdj
 * @apiGroup SceneDiscountAdj
 *
 * @apiVersion 1.0.0
 *
 * @apiUse SceneDiscountAdj
 *
 */
func (s *SceneDiscountAdjAction) List() util.RstMsg {
	paramMap := util.GetPageMsg(&s.Ctx)
	pageData, err := sceneDiscountAdjService.List(paramMap)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessListMsg("查询派生调节系数分页成功", pageData)
}

/**
 * @api {get} /scenediscountadj/(*param) 多参数查询派生调节系数
 * @apiName GetSceneDiscountAdj
 * @apiGroup SceneDiscountAdj
 *
 * @apiVersion 1.0.0
 *
 * @apiUse SceneDiscountAdj
 *
 */
func (s *SceneDiscountAdjAction) Get() util.RstMsg {
	paramMap, err := util.GetParmFromRouter(&s.Ctx)
	if nil != err {
		zlog.Error(err.Error(), err)
		return util.ErrorMsg(err.Error(), err)
	}
	// 判断是否为分页查询
	if util.IsPaginQuery(&s.Ctx) {
		util.GetPageMsg(&s.Ctx, paramMap)
		pageData, err := sceneDiscountAdjService.List(paramMap)
		if nil != err {
			return util.ErrorMsg(err.Error(), err)
		}
		return util.SuccessListMsg("多参数查询派生调节系数分页成功", pageData)
	}

	// 多参数查询不分页
	sceneDiscountAdjs, err := sceneDiscountAdjService.Find(paramMap)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("多参数查询指标成功", sceneDiscountAdjs)
}

/**
 * @api {post} /scenediscountadj 新增派生调节系数
 * @apiName PostSceneDiscountAdj
 * @apiGroup SceneDiscountAdj
 *
 * @apiVersion 1.0.0
 *
 * @apiUse SceneDiscountAdj
 *
 */
func (s *SceneDiscountAdjAction) Post() util.RstMsg {
	var SceneDiscountAdj = new(par.SceneDiscountAdj)
	err := currentMsg.DecodeJson(s.Ctx, SceneDiscountAdj)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	err = sceneDiscountAdjService.Add(SceneDiscountAdj)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("新增派生调节系数成功", SceneDiscountAdj)
}

/**
 * @api {put} /scenediscountadj 更新派生调节系数
 * @apiName PutSceneDiscountAdj
 * @apiGroup SceneDiscountAdj
 *
 * @apiVersion 1.0.0
 *
 * @apiUse SceneDiscountAdj
 *
 */
func (s *SceneDiscountAdjAction) Put() util.RstMsg {
	var SceneDiscountAdj = new(par.SceneDiscountAdj)
	err := currentMsg.DecodeJson(s.Ctx, SceneDiscountAdj)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	err = sceneDiscountAdjService.Update(SceneDiscountAdj)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("更新派生调节系数成功", SceneDiscountAdj)
}

/**
 * @api {delete} /scenediscountadj 删除派生调节系数
 * @apiName DeleteSceneDiscountAdj
 * @apiGroup SceneDiscountAdj
 *
 * @apiVersion 1.0.0
 *
 * @apiUse SceneDiscountAdj
 *
 */
func (s *SceneDiscountAdjAction) Delete() util.RstMsg {
	var SceneDiscountAdj = new(par.SceneDiscountAdj)
	err := s.DecodeJson(SceneDiscountAdj)
	if nil != err {
		return util.ErrorMsg("不能转换为实体", err)
	}
	err = sceneDiscountAdjService.Delete(SceneDiscountAdj)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("删除派生调节系数成功", SceneDiscountAdj)
}
