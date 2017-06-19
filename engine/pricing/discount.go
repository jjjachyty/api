package pricing

import (
	"fmt"

	"pccqcpa.com.cn/app/rpm/api/models/biz/ln"
	"pccqcpa.com.cn/app/rpm/api/services/parService"
	"pccqcpa.com.cn/app/rpm/api/util"
)

var sceneDiscountService parService.SceneDiscountService

// 存量贡献折让率计算
func (PricingBus) StockDiscountCalculate(pricing *ln.LnPricing, businessType string) (interface{}, error) {
	if 0 != pricing.StockUsage { //存量使用为0 则直接返回0
		sceneDiscount, err := sceneDiscountService.GetStockDiscount(pricing.LnBusiness.Cust.CustImplvl, pricing.LnBusiness.Cust.CustSize)
		if nil == err {
			return sceneDiscount.Rate.Float64, nil
		}

		return 0.0, err
	}
	return 0.0, nil
}
func (PricingBus) DerivedDiscountCalculate(pricing *ln.LnPricing, businessType string) (interface{}, error) {

	sceneDiscount, err := sceneDiscountService.GetDerivedDiscount(pricing.LnBusiness.Cust.CustImplvl, pricing.LnBusiness.Cust.CustSize, pricing.LnBusiness.Product.ProductCode)
	if nil == err {
		return sceneDiscount.Rate.Float64, nil
	}

	return 0.0, err
}

// 存量贡献折让系数计算
func (p PricingBus) StockDiscountCoefficientCalculate(pricing *ln.LnPricing, businessType string) (interface{}, error) { // 获取调节系数
	if 0 != pricing.StockUsage { //存量使用为0 则直接返回0
		paramMap := map[string]interface{}{
			"flag":           util.FLAG_TRUE,
			"biz_type":       util.DISCOUNT_STOCK,
			"gap_proportion": pricing.LnBusiness.Cust.GapProportion, // 实际EVA比例
		}
		deriveAdjs, err := sceneDiscountAdjService.Find(paramMap)
		if nil != err {
			return nil, err
		}
		length := len(deriveAdjs)
		var sceneDiscountAdj float64
		switch length {
		case 0:
			sceneDiscountAdj = 1.00
		case 1:
			sceneDiscountAdj = deriveAdjs[0].AdjValue
		default:
			return nil, fmt.Errorf("查询到存量优惠调节系数有多个,请检查表[RPM_PAR_SCENE_DISCOUNT_ADJ]配置%v", paramMap)
		}

		return sceneDiscountAdj, nil
	}
	return 0.0, nil
}
func (p PricingBus) DerivedDiscountCoefficientCalculate(pricing *ln.LnPricing, businessType string) (interface{}, error) { // 获取调节系数
	paramMap := map[string]interface{}{
		"flag":           util.FLAG_TRUE,
		"biz_type":       util.DISCOUNT_DERIVED,
		"gap_proportion": pricing.LnBusiness.Cust.GapProportion, // 实际EVA比例
	}
	deriveAdjs, err := sceneDiscountAdjService.Find(paramMap)
	if nil != err {
		return nil, err
	}
	length := len(deriveAdjs)
	var sceneDiscountAdj float64
	switch length {
	case 0:
		sceneDiscountAdj = 1.00
	case 1:
		sceneDiscountAdj = deriveAdjs[0].AdjValue
	default:
		return nil, fmt.Errorf("查询到存量优惠调节系数有多个,请检查表[RPM_PAR_SCENE_DISCOUNT_ADJ]配置%v", paramMap)
	}
	return sceneDiscountAdj, nil
}
