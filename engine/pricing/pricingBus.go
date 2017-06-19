package pricing

import (
	"fmt"
	"reflect"

	"pccqcpa.com.cn/app/rpm/api/models/biz/ln"
	"pccqcpa.com.cn/app/rpm/api/models/par"
	"pccqcpa.com.cn/app/rpm/api/util"

	funct "pccqcpa.com.cn/app/rpm/api/engine/functionIndicators"
)

type PricingBus struct {
}

var ftpRate funct.FtpRate
var capPftRate funct.CapPftRate
var termDay funct.TermDay
var ocRate funct.OcRate
var lgdRate funct.LgdRate
var ecRate funct.EcRate
var addTax funct.AddRate
var capCostRate funct.CapCostRate
var pdRate funct.PdRate
var incomeTax funct.IncomeRate

var cooperationPeriodDiscount funct.CooperationPeriodDiscount
var useProductDiscount funct.UseProductDiscount

// 初始化map key：业务类型businessType value：true／false
// true：直接反射取值
// false：调用函数指标
var isFunction = map[string]bool{
	util.LN_BUSINESS: false,
	util.LN_SCENE:    true,
	util.LN_INVERSE:  true,
}

// 计算资金成本率
func (p PricingBus) FtpRateCalculate(pricing *ln.LnPricing, businessType string) (interface{}, error) {
	fmt.Println("计算资金成本率")
	if isFunction[businessType] {
		return pricing.FtpRate, nil
	}

	term := p.getTermDay(pricing.LnBusiness.Term, pricing.LnBusiness.TermMult)
	paramMap := map[string]interface{}{
		"rate_type":          pricing.LnBusiness.RateType,
		"currency":           pricing.LnBusiness.Currency,
		"term":               term,
		"organ":              pricing.LnBusiness.Organ.OrganCode,
		"rpym_interest_freq": pricing.LnBusiness.RpymInterestFreq,
		"rpym_capital_freq":  pricing.LnBusiness.RpymCapitalFreq,
		"rpym_type":          pricing.LnBusiness.RpymType,
		"reprice_freq":       pricing.LnBusiness.RepriceFreq,
		"product":            pricing.LnBusiness.Product.ProductCode,
		"param_type":         util.LN_BUSINESS,
	}
	rate, err := ftpRate.Calulate(paramMap)
	if nil != err {
		return nil, fmt.Errorf("贷款：" + err.Error())
	}
	pricing.FtpRate = rate
	return rate, nil
}

func (PricingBus) CapPftRateCalculate(pricing *ln.LnPricing, businessType string) (interface{}, error) {
	if isFunction[businessType] {
		return pricing.CapPftRate, nil
	}

	paramMap := map[string]interface{}{
		"organ":     pricing.LnBusiness.Organ.OrganCode,
		"product":   pricing.LnBusiness.Product.ProductCode,
		"cust_size": pricing.LnBusiness.Cust.CustSize,
	}
	rate, err := capPftRate.Calulate(paramMap)
	if nil != err {
		return nil, err
	}
	pricing.CapPftRate = rate
	return rate, nil
}

func (p PricingBus) TermDayCalculate(pricing *ln.LnPricing, businessType string) (interface{}, error) {
	termDay := p.getTermDay(pricing.LnBusiness.Term, pricing.LnBusiness.TermMult)
	return termDay, nil
}

func (PricingBus) OcRateCalculate(pricing *ln.LnPricing, businessType string) (interface{}, error) {
	if isFunction[businessType] {
		return pricing.OcRate, nil
	}
	paramMap := map[string]interface{}{
		"organ":   pricing.LnBusiness.Organ.OrganCode,
		"product": pricing.LnBusiness.Product.ProductCode,
		//"cust_size":  pricing.LnBusiness.Cust.CustSize,
		"param_type": util.PARAM_LN,
	}
	rate, err := ocRate.Calulate(paramMap)
	if nil != err {
		return nil, err
	}
	pricing.OcRate = rate
	return rate, nil
}

func (PricingBus) LgdRateCalculate(pricing *ln.LnPricing, businessType string) (interface{}, error) {
	if isFunction[businessType] {
		return pricing.LgdRate, nil
	}

	// 判断主担保方式是否为信用
	// 如果为信用，则返回 1
	if util.MAIN_MORTGAGE_TYPE_CREDIT == pricing.LnBusiness.MainMortgageType {
		pricing.LgdRate = util.CREDIT_LOW_LGD
		return util.CREDIT_LOW_LGD, nil
	}

	paramMap := map[string]interface{}{
		"business_code": pricing.LnBusiness.BusinessCode,
		"principal":     pricing.LnBusiness.Principal,
	}
	rate, err := lgdRate.Calulate(paramMap)
	if nil != err {
		return nil, err
	}
	pricing.LgdRate = rate
	return rate, nil
}

func (PricingBus) EcRateCalculate(pricing *ln.LnPricing, businessType string) (interface{}, error) {
	if isFunction[businessType] {
		return pricing.EcRate, nil
	}

	paramMap := map[string]interface{}{
		"organ":   pricing.LnBusiness.Organ.OrganCode,
		"product": pricing.LnBusiness.Product.ProductCode,
	}
	rate, err := ecRate.Calulate(paramMap)
	if nil != err {
		return nil, err
	}
	pricing.EcRate = rate
	return rate, nil
}

func (PricingBus) AddTaxCalculate(pricing *ln.LnPricing, businessType string) (interface{}, error) {
	if isFunction[businessType] {
		return pricing.AddTax, nil
	}

	paramMap := map[string]interface{}{}
	rate, err := addTax.Calulate(paramMap)
	if nil != err {
		return nil, err
	}
	pricing.AddTax = rate
	return rate, nil
}

func (PricingBus) CapCostRateCalculate(pricing *ln.LnPricing, businessType string) (interface{}, error) {
	if isFunction[businessType] {
		return pricing.CapCostRate, nil
	}
	paramMap := map[string]interface{}{
		"organ":     pricing.LnBusiness.Organ.OrganCode,
		"product":   pricing.LnBusiness.Product.ProductCode,
		"cust_size": pricing.LnBusiness.Cust.CustSize,
	}
	rate, err := capCostRate.Calulate(paramMap)
	if nil != err {
		return nil, err
	}
	pricing.CapCostRate = rate
	return rate, nil
}

func (PricingBus) PdRateCalculate(pricing *ln.LnPricing, businessType string) (interface{}, error) {
	if isFunction[businessType] {
		return pricing.PdRate, nil
	}

	paramMap := map[string]interface{}{
		"cust_credit": pricing.LnBusiness.Cust.CustCredit,
	}
	rate, err := pdRate.Calulate(paramMap)
	if nil != err {
		return nil, err
	}
	pricing.PdRate = rate
	return rate, nil
}

func (PricingBus) IncomeTaxCalculate(pricing *ln.LnPricing, businessType string) (interface{}, error) {
	if isFunction[businessType] {
		return pricing.IncomeTax, nil
	}

	paramMap := map[string]interface{}{}
	rate, err := incomeTax.Calulate(paramMap)
	if nil != err {
		return nil, err
	}
	pricing.IncomeTax = rate
	return rate, nil
}

// 执行利率
func (PricingBus) IntRateCalculate(pricing *ln.LnPricing, businessType string) (interface{}, error) {
	return pricing.IntRate, nil
}

// 反算期限
// 如果期限大于一年取一年，小于一年取实际期限
func (p PricingBus) InverseTermYearCalculate(pricing *ln.LnPricing, businessType string) (interface{}, error) {
	term := p.getTermDay(pricing.LnBusiness.Term, pricing.LnBusiness.TermMult)
	if term >= 360 {
		return 1, nil
	}
	return float64(term) / 360.0, nil
}

// 合作年限优惠参数
func (PricingBus) CooperationPeriodDiscountCalculate(pricing *ln.LnPricing, businessType string) (interface{}, error) {

	paramMap := map[string]interface{}{
		"cooperation_period": pricing.LnBusiness.Cust.CooperationPeriod,
	}
	pricing.CooperationPeriod = pricing.LnBusiness.Cust.CooperationPeriod
	discount, err := cooperationPeriodDiscount.Calulate(paramMap)
	if nil != err {
		return nil, err
	}
	pricing.CooperationPeriodDiscount = discount
	return discount, nil
}

// 实用产品优惠点数
func (PricingBus) UseProductDiscountCalculate(pricing *ln.LnPricing, businessType string) (interface{}, error) {
	paramMap := map[string]interface{}{
		"use_product": pricing.LnBusiness.Cust.UseProduct,
	}
	pricing.UseProduct = pricing.LnBusiness.Cust.UseProduct
	discount, err := useProductDiscount.Calulate(paramMap)
	if nil != err {
		return nil, err
	}
	pricing.UseProductDiscount = discount
	return discount, nil
}

func (PricingBus) getTermDay(term int, termMult string) int {
	switch termMult {
	case "Y":
		return term * 360
	case "M":
		return term * 30
	case "D":
		return term
	}
	return 0
}

func init() {
	PricingBus{}.Init()
}

func (p PricingBus) Init() {
	pricingBusType := reflect.TypeOf(new(PricingBus))
	var dicts = make([]*par.Dict, 0)
	for i := 0; i < pricingBusType.NumMethod(); i++ {
		var dict = par.Dict{
			DictCode: pricingBusType.Method(i).Name,
			DictName: pricingBusType.Method(i).Name,
			Sort:     i,
			Flag:     util.FLAG_TRUE,
		}
		dicts = append(dicts, &dict)
	}
	util.PutCacheByCacheName(util.RPM_PARENT_DICT_CACHE, util.PRICING_BUS_DICT, dicts, 0)
}
