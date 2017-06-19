package analysisSigDim

import (
	"fmt"
	"time"

	"github.com/lunny/tango"

	"pccqcpa.com.cn/app/rpm/api/models/bi/analysis"
	"pccqcpa.com.cn/app/rpm/api/services/bi/analysisService"
	"pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/components/zlog"
)

var analysisSigDimService analysisService.AnalysisSigDimService

type AnalysisSigDimAction struct {
	tango.Json
	tango.Ctx
}

// 分页查询
// by author Jason
// by time 2016-11-15 16:10:56
func (this *AnalysisSigDimAction) List() util.RstMsg {
	paramMap := util.GetPageMsg(&this.Ctx)
	pageData, err := analysisSigDimService.List(paramMap)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessListMsg("分页查询信息成功", pageData)
}

// 查询日期
// by author Jason
// by time 2016-11-15 16:10:56
func (this AnalysisSigDimAction) ListDate() util.RstMsg {
	rst, err := analysisSigDimService.FindDistinctDate()
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("分页查询信息成功", rst)
}

// 多参数查询
// by author Jason
// by time 2016-11-15 16:10:56
func (this *AnalysisSigDimAction) Get() util.RstMsg {
	paramMap, err := util.GetParmFromRouter(&this.Ctx)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	if date, ok := paramMap["as_of_date"]; ok {
		asOfDate, err := time.Parse("2006-01-02", date.(string))
		if nil != err {
			return util.ErrorMsg("转换日期格式出错", err)
		}
		paramMap["as_of_date"] = asOfDate
	}
	if util.IsPaginQuery(&this.Ctx) {
		util.GetPageMsg(&this.Ctx, paramMap)
		pageData, err := analysisSigDimService.List(paramMap)
		if nil != err {
			return util.ErrorMsg(err.Error(), err)
		}
		return util.SuccessListMsg("分页查询信息成功", pageData)
	}

	rst, err := analysisSigDimService.Find(paramMap)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("查询信息成功", rst)
}

// 新增信息
// by author Jason
// by time 2016-11-15 16:10:56
// func (this *AnalysisSigDimAction) Post() util.RstMsg {
// 	var one = new(analysis.AnalysisSigDim)
// 	err := this.DecodeJson(one)
// 	if nil != err {
// 		er := fmt.Errorf("json转换实体出错")
// 		zlog.Error(er.Error(), err)
// 		return util.ErrorMsg(er.Error(), err)
// 	}
// 	err = analysisSigDimService.Add(one)
// 	if nil != err {
// 		return util.ErrorMsg(err.Error(), err)
// 	}
// 	return util.SuccessMsg("新增信息成功", one)
// }

// 执行存储过程，占用post方法
func (this *AnalysisSigDimAction) Post() util.RstMsg {
	asOfDate := this.Form("asOfDate")
	date, err := time.Parse("2006-01-02", asOfDate)
	if nil != err {
		er := fmt.Errorf("转换日期格式出错【%v】", asOfDate)
		zlog.Error(er.Error(), err)
		return util.ErrorMsg(er.Error(), err)
	}
	err = analysisSigDimService.ExecProcedure(date)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("执行贷款定价水平分析单一维度存储过程成功", nil)
}

// 更新信息
// by author Jason
// by time 2016-11-15 16:10:56
func (this *AnalysisSigDimAction) Put() util.RstMsg {
	var one = new(analysis.AnalysisSigDim)
	err := this.DecodeJson(one)
	if nil != err {
		er := fmt.Errorf("json转换实体出错")
		zlog.Error(er.Error(), err)
		return util.ErrorMsg(er.Error(), err)
	}
	err = analysisSigDimService.Update(one)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("更新信息成功", one)
}

// 删除信息
// by author Jason
// by time 2016-11-15 16:10:56
func (this *AnalysisSigDimAction) Delete() util.RstMsg {
	var one = new(analysis.AnalysisSigDim)
	err := this.DecodeJson(one)
	if nil != err {
		er := fmt.Errorf("json转换实体出错")
		zlog.Error(er.Error(), err)
		return util.ErrorMsg(er.Error(), err)
	}
	err = analysisSigDimService.Delete(one)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("删除信息成功", one)
}
