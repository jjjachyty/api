package bi

import (
	// "strconv"
	"strings"

	"github.com/lunny/tango"
	"pccqcpa.com.cn/app/rpm/api/models/bi/ln"
	"pccqcpa.com.cn/app/rpm/api/services/biService"
	"pccqcpa.com.cn/app/rpm/api/util"
)

type ExcelImportAction struct {
	tango.Json
	tango.Ctx
}

var loanInfoService biService.LoanInfoService

//贷款基本信息导入
func (ctx ExcelImportAction) LoanInfo() util.RstMsg {

	var loaninfos []ln.LoanInfo

	err := ctx.SaveToFile("loanInfo", util.GetExcelTmpPath()+"loanInfo.xlsx")

	if err != nil {
		return util.ErrorMsg("贷款基础数据导入失败", err)
	}

	err = util.GetValueFormExcel(&loaninfos, util.GetExcelTmpPath()+"loanInfo.xlsx")

	if err != nil {
		return util.ErrorMsg("贷款基础数据导入失败", err)
	}

	r, er := loanInfoService.BatchInsert(loaninfos)

	if nil == er {
		return util.SuccessMsg("导入成功", r)
	}
	return util.ErrorMsg("导入失败", err)
}

func (e *ExcelImportAction) ListDate() util.RstMsg {
	rst, err := loanInfoService.ListDate()
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("查询贷款信息日期成功", rst)
}

func (e ExcelImportAction) Get() util.RstMsg {
	paramMap := util.GetPageMsg(&e.Ctx)
	asOfDate := ""
	if date, ok := paramMap["as_of_date"]; ok {
		asOfDate = strings.TrimSpace(date.(string))
	}
	rst, err := loanInfoService.FindAmount(asOfDate)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("查询贷款信息金额降序成功", rst)
}
