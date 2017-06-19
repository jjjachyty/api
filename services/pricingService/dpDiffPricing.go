package pricingService

import (
	"fmt"

	"pccqcpa.com.cn/app/rpm/api/models/dp"
	"pccqcpa.com.cn/app/rpm/api/services/dpService"
	"pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/components/zlog"
)

type DpDiffPricingService struct{}

func (d DpDiffPricingService) pricing(model *dp.DpPricing) error {
	pricingRst, err := rpmEngine.StartPricing(model, util.ONE_DP)
	if nil != err {
		return err
	}
	model.DpRate = pricingRst["DpRate"].(float64)
	if "" != model.UUID {
		return dpService.DpPricingService{}.Update(model)
	}

	err = model.Add()
	if nil != err {
		return err
	}
	paramMap := map[string]interface{}{
		"business_code": model.BusinessCode,
	}
	dpPricingModel, err := dpService.DpPricingService{}.FindOne(paramMap)
	if nil != err {
		return err
	} else if nil != dpPricingModel {
		model.UUID = dpPricingModel.UUID
		return nil
	} else {
		er := fmt.Errorf("查询存款定价单为空")
		zlog.Error(er.Error(), er)
		return er
	}
}
