package engine

import (
	"errors"
	"strings"

	"github.com/lunny/tango"

	"pccqcpa.com.cn/app/rpm/api/models/dp"
	"pccqcpa.com.cn/app/rpm/api/services/pricingService"
	"pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/app/rpm/api/util/currentMsg"
)

type DpOnePricing struct {
	tango.Ctx
	tango.Json
}

//URL POST:/api/rpm/pricing/dpone/
func (d DpOnePricing) Post() util.RstMsg {
	businessCode := d.Form("businessCode")
	if "" == strings.TrimSpace(businessCode) {
		er := errors.New("存款一对一定价未传businessCode")
		return util.ErrorMsg(er.Error(), er)
	}
	var dpOnePricing = new(dp.DpOnePricing)

	err := currentMsg.DecodeJsonUser(d.Ctx, dpOnePricing)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	organCode, err := currentMsg.GetCurrentUserBranchCode(d.Ctx)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	dpOnePricing.BusinessCode = businessCode
	dpOnePricing.Organ.OrganCode = organCode
	rst, err := pricingService.NewDpOnePricing().Pricing(dpOnePricing)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("存款一对一定价成功", rst)

}
