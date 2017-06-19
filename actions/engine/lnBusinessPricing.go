package engine

import (
	"errors"
	// "fmt"
	"strings"

	"github.com/lunny/tango"

	"pccqcpa.com.cn/app/rpm/api/models/biz/ln"
	"pccqcpa.com.cn/app/rpm/api/services/pricingService"
	"pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/app/rpm/api/util/currentMsg"
)

var service pricingService.PricingService

type LnBusinessPricingAction struct {
	tango.Ctx
	tango.Json
}

// url: /api/rpm/pricing/lnbase/
func (l *LnBusinessPricingAction) Post() util.RstMsg {
	var lnPrincing = new(ln.LnPricing)
	var err error
	err = currentMsg.DecodeJsonUser(l.Ctx, lnPrincing)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	businessCode := l.Form("BusinessCode")
	StockUsage := l.FormFloat64("StockUsage")
	if "" == businessCode {
		return util.ErrorMsg("[BusinessCode.业务单号]为空,不能计算对公贷款定价", errors.New("[BusinessCode.业务单号]为空,不能计算对公贷款定价"))
	} else if "" == l.Form("StockUsage") {
		return util.ErrorMsg("[StockUsage.存量优惠使用]为空,不能计算对公贷款情景优惠", errors.New("[StockUsage.存量优惠使用]为空,不能计算对公贷款情景优惠"))
	} else {
		lnPrincing, err = service.LnBusinessPricing(businessCode, util.LN_BUSINESS+","+util.LN_SCENE, map[string]interface{}{"StockUsage": StockUsage})
		if nil != err {
			return util.ErrorMsg(err.Error(), err)
		}

	}
	lnPrincing.MarginType = strings.TrimSpace(lnPrincing.MarginType)

	return util.SuccessMsg("对公贷款定价成功", lnPrincing)

}

// url: /api/rpm/pricing/lnbase/
func (l *LnBusinessPricingAction) Post_back() util.RstMsg {
	var lnPrincing *ln.LnPricing
	var err error
	businessCode := l.Form("BusinessCode")
	if "" == businessCode {
		return util.ErrorMsg("BusinessCode为空,不能计算对公贷款基础定价", nil)
	} else {
		lnPrincing, err = service.LnBusinessPricing(businessCode, util.LN_BUSINESS)
		if nil != err {
			return util.ErrorMsg(err.Error(), err)
		}
	}

	return util.SuccessMsg("对公贷款基础定价成功", lnPrincing)

}
