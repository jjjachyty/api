package pricing

import (
	"fmt"

	"pccqcpa.com.cn/app/rpm/api/models/biz/ln"
	"pccqcpa.com.cn/app/rpm/api/services/loanService"
	"pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/components/zlog"
)

func (PricingBus) PrincipalCalculate(pricing *ln.LnPricing, businessType string) (interface{}, error) {
	return pricing.LnBusiness.Principal, nil
}
func (PricingBus) StockContributionCalculate(pricing *ln.LnPricing, businessType string) (interface{}, error) {
	return pricing.StockUsage, nil
}

//派生存款FTP利率
func (p PricingFormulaBus) DerviedDPEVACalculate(pricing *ln.LnPricing) (interface{}, error) {
	var sumeva float64
	var eva float64
	var err error
	var ftpRateValue float64
	var ocRateValue float64

	var businessCode = pricing.BusinessCode
	sceneDps, err := loanService.SceneDpService{}.Find(map[string]interface{}{"business_Code": businessCode})
	for _, sceneDp := range sceneDps {
		fmt.Println("－－－计算派生存款eva 期限1－－－", sceneDp.Term)
		//查询FTP利率
		paramMap := map[string]interface{}{
			"currency":   sceneDp.Currency,
			"term":       sceneDp.Term,
			"organ":      pricing.LnBusiness.Organ.OrganCode,
			"product":    sceneDp.Product.ProductCode,
			"param_type": util.DP_BUSINESS,
		}
		ftpRateValue, err = ftpRate.Calulate(paramMap)
		if nil != err {
			return nil, fmt.Errorf("派生存款EVA【%s】:"+err.Error(), sceneDp.Product.ProductName)
		}
		//查询运营成本率
		ocParamMap := map[string]interface{}{
			"organ":     pricing.LnBusiness.Organ.OrganCode,
			"product":   sceneDp.Product.ProductCode,
			//"cust_size": pricing.LnBusiness.Cust.CustSize,
		}
		ocRateValue, err = ocRate.Calulate(ocParamMap)
		if nil != err {
			return nil, err
		}

		var term = float64(1.00)
		if 0 != sceneDp.Term { //不为活期
			term = (float64(sceneDp.Term) / 360)
		}
		zlog.Infof("计算派生存款Eva(ftpRate-ocRate-sceneDp.Rate)*sceneDp.Value*term):%s", nil, fmt.Sprintf("(%f-%f-%f)*%f*%f)", ftpRateValue, ocRateValue, sceneDp.Rate, sceneDp.Value, term))

		eva, err = util.Calculate(fmt.Sprintf("(%f-%f-%f)*%f*%f)", ftpRateValue, ocRateValue, sceneDp.Rate, sceneDp.Value, term))
		if err == nil {
			zlog.Infof("计算派生存款Eva(ftpRate-ocRate-sceneDp.Rate)*sceneDp.Value*term)结果%f", nil, eva)
			sumeva = sumeva + eva
		}

	}
	pricing.DerviedDPEVA = sumeva

	return sumeva, err
}

//派生存款优惠点数
func (p PricingFormulaBus) DerviedDPPointsCalculate(pricing *ln.LnPricing) (interface{}, error) {
	var points float64
	var businessCode = pricing.BusinessCode
	lnTerm := PricingBus{}.getTermDay(pricing.LnBusiness.Term, pricing.LnBusiness.TermMult) / 360.0
	sceneDps, err := loanService.SceneDpService{}.Find(map[string]interface{}{"business_Code": businessCode})
	for _, sceneDp := range sceneDps {
		//查询FTP利率
		paramMap := map[string]interface{}{
			"currency":   sceneDp.Currency,
			"term":       sceneDp.Term,
			"organ":      pricing.LnBusiness.Organ.OrganCode,
			"product":    sceneDp.Product.ProductCode,
			"param_type": util.DP_BUSINESS,
		}
		ftpRate, err := ftpRate.Calulate(paramMap)
		if nil != err {
			zlog.Error("计算派生存款优惠点数出错", err)
			return nil, fmt.Errorf("派生存款优惠点数【%s】："+err.Error(), sceneDp.Product.ProductName)
		}
		//查询运营成本率
		ocParamMap := map[string]interface{}{
			"organ":     pricing.LnBusiness.Organ.OrganCode,
			"product":   sceneDp.Product.ProductCode,
			//"cust_size": pricing.LnBusiness.Cust.CustSize,
		}
		ocRate, err := ocRate.Calulate(ocParamMap)
		if nil != err {
			return nil, err
		}

		var term float64 = 1.0
		if 0 != sceneDp.Term { //不为活期
			term = (float64(sceneDp.Term) / float64(360.0))
		}

		eva := (ftpRate - ocRate - sceneDp.Rate) * sceneDp.Value * term
		fmt.Println(ftpRate, ocRate, sceneDp.Rate, sceneDp.Value, term)
		//派生业务折让率
		sceneDiscount, err := sceneDiscountService.GetDerivedDiscount(pricing.LnBusiness.Cust.CustImplvl, pricing.LnBusiness.Cust.CustSize, sceneDp.Product.ProductCode)
		if nil != err {
			return nil, err
		}
		//zlog.Debugf("计算派生存款优惠点数(eva / (pricing.LnBusiness.Principal * lnTerm) * sceneDiscount.Rate)%v", nil, "eva", eva, "Principal", pricing.LnBusiness.Principal, "lnTerm", lnTerm, "sceneDiscount.Rate", sceneDiscount.Rate.Float64)

		sceneDiscountParamMap := map[string]interface{}{
			"flag":           util.FLAG_TRUE,
			"biz_type":       util.DISCOUNT_DERIVED,
			"gap_proportion": pricing.LnBusiness.Cust.GapProportion, // 实际EVA比例
		}
		sceneDiscountAdjs, err := sceneDiscountAdjService.Find(sceneDiscountParamMap)
		if nil != err {
			return nil, err
		}
		length := len(sceneDiscountAdjs)
		var sceneDiscountAdj float64
		switch length {
		case 0:
			sceneDiscountAdj = 1.00
		case 1:
			sceneDiscountAdj = sceneDiscountAdjs[0].AdjValue
		default:
			return nil, fmt.Errorf("查询到存量优惠调节系数有多个,请检查表[RPM_PAR_SCENE_DISCOUNT_ADJ]配置%v", paramMap)
		}
		//points = points + (eva / (pricing.LnBusiness.Principal * float64(lnTerm)) * sceneDiscount.Rate.Float64 * sceneDiscountAdj)
		fmt.Println(fmt.Sprintf("%f/(%f*%f)*%f*%f", eva, pricing.LnBusiness.Principal, float64(lnTerm), sceneDiscount.Rate.Float64, sceneDiscountAdj))
		pointsTmp, err := util.Calculate(fmt.Sprintf("%f/(%f*%f)*%f*%f", eva, pricing.LnBusiness.Principal, float64(lnTerm), sceneDiscount.Rate.Float64, sceneDiscountAdj))
		if err == nil {
			points = points + pointsTmp
		}
		zlog.Debugf("DerviedDPPoints:%f", nil, points)
	}

	pricing.DerviedDPPoints = points
	zlog.Debugf("最终返回DerviedDPPoints:%f", nil, points)
	return points, err
}

//目标利率 TgtRateCalculate
func (p PricingFormulaBus) TgtRateCalculate(pricing *ln.LnPricing) (interface{}, error) {
	return pricing.TgtRate, nil
}
