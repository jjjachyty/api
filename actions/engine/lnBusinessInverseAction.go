package engine

import (
	"fmt"
	"github.com/lunny/tango"

	// "pccqcpa.com.cn/app/rpm/api/services/pricingService"
	"pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/components/zlog"
)

// var service pricingService.PricingService

type LnBusinessInverseAction struct {
	tango.Ctx
	tango.Json
}

// url: /api/rpm/lnbusinessinverse/businessCode/xxx
func (l *LnBusinessInverseAction) Post() util.RstMsg {

	paramMap := make(map[string]interface{})
	paramMap["IntRate"] = l.FormFloat64("IntRate")
	paramMap["MarginType"] = l.Form("MarginType")
	paramMap["MarginInt"] = l.FormFloat64("MarginInt")

	if "" != l.Form("BusinessCode") {
		pricingRstMsg, err := service.LnBusinessPricing(l.Form("BusinessCode"), util.LN_INVERSE, paramMap)
		if nil != err {
			return util.ErrorMsg(err.Error(), err)
		}
		return util.SuccessMsg("利率反算成功", pricingRstMsg)
	}
	er := fmt.Errorf("调用定价引擎接口出错，未传BusinessCode值")
	zlog.Error(er.Error(), er)
	return util.ErrorMsg(er.Error(), er)

}
