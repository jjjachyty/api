package workFlow

import (
	"github.com/lunny/tango"
	"pccqcpa.com.cn/app/rpm/api/services/loanService"
	"pccqcpa.com.cn/app/rpm/api/util"
)

//PricingListAction type 对公贷款定价单控制实体
type WorkFlowLnPricingAction struct {
	tango.Json
	tango.Ctx
}

var pls loanService.PricingListService

func (pla *WorkFlowLnPricingAction) Post() util.RstMsg {
	businessCode := pla.Form("businessKey")
	var status string
	switch pla.Form("status") {
	case "end":
		status = "4"
	case "reback":
		status = "5"
	default:
		status = "undefine"
	}
	paramMap := map[string]interface{}{
		"business_code": businessCode,
		"status":        status,
	}
	err := pls.Patch(paramMap)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("修改对公贷款结果信息成功", nil)
}
