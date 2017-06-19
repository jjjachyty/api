package dp

import (
	"github.com/lunny/tango"
	// "pccqcpa.com.cn/app/rpm/api/models/biz/dp"
	"pccqcpa.com.cn/app/rpm/api/services/dpService"
	"pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/components/zlog"
)

var dpPricingService dpService.DpBasePricingService

type DpPricingAction struct {
	tango.Json
	tango.Ctx
}

//api/rpm/dpPricing
func (DpPricingAction) List() util.RstMsg {

	var PageData util.PageData
	dpPricing, err := dpPricingService.GetDpPricing()
	if nil == err {
		PageData.Rows = dpPricing
		return util.ReturnSuccess("存款标准化定价查询成功", PageData)
	}
	return util.ErrorMsg("存款标准化定价查询失败", err)
}

//api/rpm/dpPricing
func (dpp *DpPricingAction) Get() util.RstMsg {
	paramMap, err := util.GetParmFromRouter(&dpp.Ctx)
	if nil != err {
		zlog.Error(err.Error(), err)
		return util.ErrorMsg(err.Error(), err)
	}
	// var dpPricing dp.DpPricing
	dpPricing, err := dpPricingService.FindDpPricing(paramMap)
	if nil != err {
		return util.ErrorMsg("存款标准化定价条件查询失败", err)
	}
	return util.SuccessMsg("存款标准化定价条件查询成功", dpPricing)
}
