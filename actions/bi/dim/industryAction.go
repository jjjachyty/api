package dim

import (
	"fmt"

	"github.com/lunny/tango"

	"pccqcpa.com.cn/app/rpm/api/models/bi/dim"

	"pccqcpa.com.cn/app/rpm/api/services/bi/dimService"

	"pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/components/zlog"
)

var industryService dimService.IndustryService

type IndustryAction struct {
	tango.Json
	tango.Ctx
}

// excel 导入
// by author Jason
// by time 2016-10-31 15:34:54
func (this IndustryAction) ExcelImport() util.RstMsg {
	var models []dim.DimIndustry
	var formName string = "dimIndustry"
	var excelName string = util.GetExcelTmpPath() + formName + ".xlsx"

	err := this.SaveToFile("file", excelName)

	if nil != err {
		return util.ErrorMsg("导入失败", err)
	}

	err = util.GetValueFormExcel(&models, excelName)

	if nil != err {
		return util.ErrorMsg("导入失败", err)
	}

	rst, err := industryService.BatchAdd(models)

	if nil != err {
		return util.ErrorMsg("导入失败", err)
	}
	return util.SuccessMsg("导入成功", rst)
}

// 分页查询
// by author Jason
// by time 2016-10-31 15:34:16
func (this *IndustryAction) List() util.RstMsg {
	paramMap := util.GetPageMsg(&this.Ctx)
	pageData, err := industryService.List(paramMap)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessListMsg("分页查询信息成功", pageData)
}

// 多参数查询
// by author Jason
// by time 2016-10-31 15:34:16
func (this *IndustryAction) Get() util.RstMsg {
	paramMap, err := util.GetParmFromRouter(&this.Ctx)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	if util.IsPaginQuery(&this.Ctx) {
		util.GetPageMsg(&this.Ctx, paramMap)
		pageData, err := industryService.List(paramMap)
		if nil != err {
			return util.ErrorMsg(err.Error(), err)
		}
		return util.SuccessListMsg("分页查询信息成功", pageData)
	}

	rst, err := industryService.Find(paramMap)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("查询信息成功", rst)
}

// 新增信息
// by author Jason
// by time 2016-10-31 15:34:16
func (this *IndustryAction) Post() util.RstMsg {
	var one = new(dim.DimIndustry)
	err := this.DecodeJson(one)
	if nil != err {
		er := fmt.Errorf("json转换实体出错")
		zlog.Error(er.Error(), err)
		return util.ErrorMsg(er.Error(), err)
	}
	err = industryService.Add(one)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("新增信息成功", one)
}

// 更新信息
// by author Jason
// by time 2016-10-31 15:34:16
func (this *IndustryAction) Put() util.RstMsg {
	var one = new(dim.DimIndustry)
	err := this.DecodeJson(one)
	if nil != err {
		er := fmt.Errorf("json转换实体出错")
		zlog.Error(er.Error(), err)
		return util.ErrorMsg(er.Error(), err)
	}
	err = industryService.Update(one)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("更新信息成功", one)
}

// 删除信息
// by author Jason
// by time 2016-10-31 15:34:16
func (this *IndustryAction) Delete() util.RstMsg {
	var one = new(dim.DimIndustry)
	err := this.DecodeJson(one)
	if nil != err {
		er := fmt.Errorf("json转换实体出错")
		zlog.Error(er.Error(), err)
		return util.ErrorMsg(er.Error(), err)
	}
	err = industryService.Delete(one)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("删除信息成功", one)
}
