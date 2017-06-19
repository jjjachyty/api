package analysisCrossDim

import (
	"fmt"

	"github.com/lunny/tango"

	"pccqcpa.com.cn/app/rpm/api/models/bi/analysis"

	"pccqcpa.com.cn/app/rpm/api/services/bi/analysisService"

	"pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/app/rpm/api/util/currentMsg"
	"pccqcpa.com.cn/components/zlog"
)

var analysisCrossDimService analysisService.AnalysisCrossDimService

type AnalysisCrossDimAction struct {
	tango.Json
	tango.Ctx
}

/**
 * @api {get} /analysisCrossDim 分页查询贷款交叉分析-多维度
 * @apiName ListAnalysisCrossDim
 * @apiGroup AnalysisCrossDim
 *
 * @apiVersion 1.0.0
 *
 * @apiUse AnalysisCrossDim
 *
 */
func (this *AnalysisCrossDimAction) List() util.RstMsg {
	paramMap := util.GetPageMsg(&this.Ctx)
	pageData, err := analysisCrossDimService.List(paramMap)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessListMsg("分页查询信息成功", pageData)
}

// 多参数查询
// by author Jason
// by time 2016-11-21 16:26:56
/**
 * @api {get} /analysisCrossDim/(*param) 多参数查询贷款交叉分析-多维度
 * @apiName GetAnalysisCrossDim
 * @apiGroup AnalysisCrossDim
 *
 * @apiVersion 1.0.0
 *
 * @apiUse AnalysisCrossDim
 *
 */
func (this *AnalysisCrossDimAction) Get() util.RstMsg {
	paramMap, err := util.GetParmFromRouter(&this.Ctx)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	if util.IsPaginQuery(&this.Ctx) {
		util.GetPageMsg(&this.Ctx, paramMap)
		pageData, err := analysisCrossDimService.List(paramMap)
		if nil != err {
			return util.ErrorMsg(err.Error(), err)
		}
		return util.SuccessListMsg("分页查询信息成功", pageData)
	}

	rst, err := analysisCrossDimService.Find(paramMap)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("查询信息成功", rst)
}

// 新增信息
// by author Jason
// by time 2016-11-21 16:26:56
/**
 * @api {post} /analysisCrossDim 新增贷款交叉分析-多维度
 * @apiName PostAnalysisCrossDim
 * @apiGroup AnalysisCrossDim
 *
 * @apiVersion 1.0.0
 *
 * @apiUse AnalysisCrossDim
 * * @apiSuccessExample {url} Success-Response:
 *   www.baidu.com
 *
 * @api {get} /user/:id
 * @apiExample {curl} Example usage:
 *     curl -i http://localhost/user/4711
 */
func (this *AnalysisCrossDimAction) Post() util.RstMsg {
	var one = new(analysis.AnalysisCrossDim)
	err := currentMsg.DecodeJson(this.Ctx, one)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	err = analysisCrossDimService.Add(one)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("新增信息成功", one)
}

// 更新信息
// by author Jason
// by time 2016-11-21 16:26:56
/**
 * @api {put} /analysisCrossDim 更新贷款交叉分析-多维度
 * @apiName PutAnalysisCrossDim
 * @apiGroup AnalysisCrossDim
 *
 * @apiVersion 1.0.0
 *
 * @apiUse AnalysisCrossDim
 *
 */
func (this *AnalysisCrossDimAction) Put() util.RstMsg {
	var one = new(analysis.AnalysisCrossDim)
	err := currentMsg.DecodeJson(this.Ctx, one)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	err = analysisCrossDimService.Update(one)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("更新信息成功", one)
}

// 删除信息
// by author Jason
// by time 2016-11-21 16:26:56
/**
 * @api {delete} /analysisCrossDim 删除贷款交叉分析-多维度
 * @apiName DeleteAnalysisCrossDim
 * @apiGroup AnalysisCrossDim
 *
 * @apiVersion 1.0.0
 *
 * @apiUse AnalysisCrossDim
 *
 */
func (this *AnalysisCrossDimAction) Delete() util.RstMsg {
	var one = new(analysis.AnalysisCrossDim)
	err := this.DecodeJson(one)
	if nil != err {
		er := fmt.Errorf("json转换实体出错")
		zlog.Error(er.Error(), err)
		return util.ErrorMsg(er.Error(), err)
	}
	err = analysisCrossDimService.Delete(one)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("删除信息成功", one)
}
