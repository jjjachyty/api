package dpService

import (
	// "errors"
	// "fmt"
	"pccqcpa.com.cn/app/rpm/api/models/biz/dp"
	// "pccqcpa.com.cn/app/rpm/api/util"
)

type DpBasePricingService struct{}

var dpPricingModel dp.DpPricingModel

func (DpBasePricingService) GetDpPricing() ([]dp.DpPricing, error) {
	return dpPricingModel.List("")
}

func (DpBasePricingService) FindDpPricing(param map[string]interface{}) ([]dp.DpPricing, error) {
	var whereSql string = "WHERE 1=1 "
	for k, val := range param {
		switch k {
		case "date":
			whereSql += "AND THIS_.CREATE_TIME=date'" + val.(string) + "'"
		case "term":
			whereSql += "AND THIS_.TERM='" + val.(string) + "'"
		case "product":
			whereSql += "AND THIS_.PRODUCT='" + val.(string) + "'"
		}
	}
	return dpPricingModel.List(whereSql)
}
