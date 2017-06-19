package dim

import (
	"fmt"

	"github.com/lunny/tango"

	"pccqcpa.com.cn/app/rpm/api/models/bi/dim"

	"pccqcpa.com.cn/app/rpm/api/services/bi/dimService"

	"pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/components/zlog"
)

var lineService dimService.LineService

type LineAction struct {
	tango.Json
	tango.Ctx
}

// excel 导入
// by author Jason
// by time 2016-10-31 15:34:54
func (this LineAction) ExcelImport() util.RstMsg {
	var models []dim.DimLine
	var formName string = "dimLine"
	var excelName string = util.GetExcelTmpPath() + formName + ".xlsx"

	err := this.SaveToFile("file", excelName)

	if nil != err {
		return util.ErrorMsg("导入失败", err)
	}

	err = util.GetValueFormExcel(&models, excelName)

	if nil != err {
		return util.ErrorMsg("导入失败", err)
	}

	rst, err := lineService.BatchAdd(models)

	if nil != err {
		return util.ErrorMsg("导入失败", err)
	}
	return util.SuccessMsg("导入成功", rst)
}

// 分页查询
// by author Jason
// by time 2016-10-31 15:35:57
func (this *LineAction) List() util.RstMsg {
	paramMap := util.GetPageMsg(&this.Ctx)
	pageData, err := lineService.List(paramMap)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessListMsg("分页查询信息成功", pageData)
}

// 多参数查询
// by author Jason
// by time 2016-10-31 15:35:57
func (this *LineAction) Get() util.RstMsg {
	paramMap, err := util.GetParmFromRouter(&this.Ctx)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	if util.IsPaginQuery(&this.Ctx) {
		util.GetPageMsg(&this.Ctx, paramMap)
		pageData, err := lineService.List(paramMap)
		if nil != err {
			return util.ErrorMsg(err.Error(), err)
		}
		return util.SuccessListMsg("分页查询信息成功", pageData)
	}

	rst, err := lineService.Find(paramMap)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("查询信息成功", rst)
}

// 新增信息
// by author Jason
// by time 2016-10-31 15:35:57
func (this *LineAction) Post() util.RstMsg {
	var one = new(dim.DimLine)
	err := this.DecodeJson(one)
	if nil != err {
		er := fmt.Errorf("json转换实体出错")
		zlog.Error(er.Error(), err)
		return util.ErrorMsg(er.Error(), err)
	}
	err = lineService.Add(one)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("新增信息成功", one)
}

// 更新信息
// by author Jason
// by time 2016-10-31 15:35:57
func (this *LineAction) Put() util.RstMsg {
	var one = new(dim.DimLine)
	err := this.DecodeJson(one)
	if nil != err {
		er := fmt.Errorf("json转换实体出错")
		zlog.Error(er.Error(), err)
		return util.ErrorMsg(er.Error(), err)
	}
	err = lineService.Update(one)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("更新信息成功", one)
}

// 删除信息
// by author Jason
// by time 2016-10-31 15:35:57
func (this *LineAction) Delete() util.RstMsg {
	var one = new(dim.DimLine)
	err := this.DecodeJson(one)
	if nil != err {
		er := fmt.Errorf("json转换实体出错")
		zlog.Error(er.Error(), err)
		return util.ErrorMsg(er.Error(), err)
	}
	err = lineService.Delete(one)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("删除信息成功", one)
}
