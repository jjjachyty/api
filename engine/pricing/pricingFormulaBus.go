package pricing

import (
	"fmt"
	// "time"

	"pccqcpa.com.cn/app/rpm/api/models/biz/ln"
	"pccqcpa.com.cn/app/rpm/api/services/loanService"
	"pccqcpa.com.cn/app/rpm/api/services/parService"
	"pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/components/zlog"

	funct "pccqcpa.com.cn/app/rpm/api/engine/functionIndicators"
)

type PricingFormulaBus struct {
}

var sceneDpService loanService.SceneDpService
var sceneItdService loanService.SceneItdService
var baseRate funct.BaseRate
var evaYieldService parService.EvaYieldService
var sceneDiscountAdjService parService.SceneDiscountAdjService
var sceneItdYieldService parService.SceneItdYieldService

// var sceneDiscountService parService.SceneDiscountService

// var ftpRate funct.FtpRate

// var pricingBus PricingBus

// 计算派生优惠净利润
// 查询存款派生业务和中间派生业务
// 循环获取存款业务和中间业务的净利润求和
func (p PricingFormulaBus) SceneNetProfitCalculate(pricing *ln.LnPricing) (interface{}, error) {
	fmt.Println("开始计算派生优惠净利润")
	paramMap := map[string]interface{}{
		"business_code": pricing.BusinessCode,
	}
	sceneDps, err := sceneDpService.Find(paramMap)
	if nil != err {
		er := fmt.Errorf("查询存款派生信息出错")
		return nil, er
	}
	sceneItds, err := sceneItdService.Find(paramMap)
	if nil != err {
		er := fmt.Errorf("查询中间派生信息出错")
		return nil, er
	}

	// 循环计算存款派生净利润
	var sumSceneNetProfit float64 = 0
	for _, sceneDp := range sceneDps {
		principal := sceneDp.Value    // 存款金额
		term := float64(sceneDp.Term) // 存款期限
		intRate := sceneDp.Rate       // 对客利率
		// 获取存款FTP利率
		dpFtpParamMap := map[string]interface{}{
			"organ":      pricing.LnBusiness.Organ.OrganCode,
			"term":       sceneDp.Term,
			"product":    sceneDp.Product.ProductCode,
			"currency":   sceneDp.Currency,
			"param_type": util.DP_BUSINESS,
		}
		rate, err := ftpRate.Calulate(dpFtpParamMap)
		if nil != err {
			er := fmt.Errorf("查询存款FTP出错")
			zlog.Error(er.Error(), err)
			return nil, er
		}

		// 获取存款运营费用率ocRate
		ocParamMap := map[string]interface{}{
			"organ":   pricing.LnBusiness.Organ.OrganCode,
			"product": sceneDp.Product.ProductCode,
			//"cust_size": pricing.LnBusiness.Cust.CustSize,
		}
		ocRateValue, err := ocRate.Calulate(ocParamMap)
		if nil != err {
			return nil, err
		}

		// 处理存款期限
		// 超过一年取一年
		// 小于一年取实际期限
		if 0 != sceneDp.Term && 360 > sceneDp.Term { //不为活期
			term = (float64(sceneDp.Term) / 360)
		} else {
			term = float64(1)
		}

		expr := fmt.Sprintf("%f", principal) + "*" + fmt.Sprint(term) + "*(" + (fmt.Sprintf("%f", rate) + "-" + fmt.Sprintf("%f", intRate) + "-" + fmt.Sprintf("%f", ocRateValue)) + ")"
		zlog.Infof("计算派生存款业务贡献(年)【%s】", nil, expr)
		netProfit, err := util.Calculate(expr)
		if nil != err {
			return nil, err
		}

		sumSceneNetProfit += netProfit
	}

	// 循环计算中间业务派生净利润
	for _, sceneItd := range sceneItds {
		// 获取中间收益率,全行统一，查询一次
		sceneItdYieldParamMap := map[string]interface{}{
			"flag":    util.FLAG_TRUE,
			"product": sceneItd.Product.ProductCode,
		}
		sceneItdYield, err := sceneItdYieldService.FindOne(sceneItdYieldParamMap)
		if nil != err {
			return nil, err
		}
		if nil == sceneItdYield {
			er := fmt.Errorf("查询中间业务收益率为空")
			zlog.Error(er.Error(), er)
			return nil, err
		}

		incomeValue := sceneItd.Value // 中间业务毛收入
		expr := fmt.Sprintf("%f", incomeValue) + "*" + fmt.Sprintf("%f", sceneItdYield.ItdYield)
		zlog.Infof("计算派生中间业务贡献(年)【%s】", nil, expr)
		netProfit, err := util.Calculate(expr)
		if nil != err {
			return nil, err
		}
		sumSceneNetProfit += netProfit
	}
	zlog.Infof("派生业务贡献(年)-净利润(元)=%f", nil, sumSceneNetProfit)
	pricing.SceneNetPorfit = sumSceneNetProfit
	return sumSceneNetProfit, nil
}

func (p PricingFormulaBus) BaseRateCalculate(pricing *ln.LnPricing) (interface{}, error) {

	term := pricingBus.getTermDay(pricing.LnBusiness.Term, pricing.LnBusiness.TermMult)

	// fmt.Println("\n\n\n\n期限\n\n\n\n", term, pricing.L nBusiness.Term, pricing.LnBusiness.TermMult)
	paramMap := map[string]interface{}{
		"base_rate_type": pricing.LnBusiness.BaseRateType,
		"term":           term,
		"flag":           util.FLAG_TRUE,
	}
	rate, err := baseRate.Calulate(paramMap)
	if nil != err {
		return nil, err
	}
	pricing.BaseRate = rate
	return rate, nil
}

// 派生中间业务EVA
func (p *PricingFormulaBus) SceneItdEvaCalculate(pricing *ln.LnPricing) (interface{}, error) {

	paramMap := map[string]interface{}{
		"business_code": pricing.LnBusiness.BusinessCode,
	}
	sceneItds, err := sceneItdService.Find(paramMap)
	if nil != err {
		return nil, err
	}

	var sceneItdEva float64
	for _, sceneItd := range sceneItds {
		sceneItdYieldValue, err := p.getSceneIbYield(sceneItd.Product.ProductCode)
		if nil != err {
			return nil, err
		}

		income := sceneItd.Value // 中收收入
		// 派生中间业务EVA
		sceneItdEva += income * sceneItdYieldValue
	}

	pricing.DerviedIBEVA = sceneItdEva
	return sceneItdEva, nil
}

// 派生中间业务优惠点数
// 中间收入＊中收收益率／（贷款金额＊期限）＊折让系数＊调节系数
func (p *PricingFormulaBus) SceneItdDiscountCalculate(pricing *ln.LnPricing) (interface{}, error) {
	principal := pricing.LnBusiness.Principal // 贷款金额
	term := pricingBus.getTermDay(pricing.LnBusiness.Term, pricing.LnBusiness.TermMult)

	// 获取调节系数
	paramMap := map[string]interface{}{
		"biz_type":       util.DISCOUNT_DERIVED,
		"flag":           util.FLAG_TRUE,
		"gap_proportion": pricing.LnBusiness.Cust.GapProportion, // 实际EVA比例
	}
	deriveAdjs, err := sceneDiscountAdjService.Find(paramMap)
	if nil != err {
		return nil, err
	}
	var deriveAdjValue float64
	if 0 == len(deriveAdjs) {
		deriveAdjValue = 1
	} else if 1 < len(deriveAdjs) {
		er := fmt.Errorf("查询派生调节系数有多条")
		return nil, er
	} else {
		deriveAdjValue = deriveAdjs[0].AdjValue
	}

	// 查询中收数组
	paramMap = map[string]interface{}{
		"business_code": pricing.LnBusiness.BusinessCode,
	}
	sceneItds, err := sceneItdService.Find(paramMap)
	if nil != err {
		return nil, err
	}

	var sceneItdDisctoun float64
	// time.Sleep(time.Second * 2)
	for _, sceneItd := range sceneItds {

		// 获取中收收益率

		sceneItdYieldValue, err := p.getSceneIbYield(sceneItd.Product.ProductCode)
		if nil != err {
			return nil, err
		}

		// 获取折让系数
		sceneDiscount, err := sceneDiscountService.GetDerivedDiscount(
			pricing.LnBusiness.Cust.CustImplvl,
			pricing.LnBusiness.Cust.CustSize,
			sceneItd.Product.ProductCode)
		if nil != err {
			return nil, err
		}
		sceneDiscountValue := sceneDiscount.Rate.Float64

		income := sceneItd.Value // 中收收入
		expr := fmt.Sprintf("%f", income) + "*" +
			fmt.Sprintf("%f", sceneItdYieldValue) + "/(" +
			fmt.Sprintf("%f", principal) + "*" +
			fmt.Sprintf("%v", term) + "/360)*" +
			fmt.Sprintf("%f", sceneDiscountValue) + "*" +
			fmt.Sprintf("%f", deriveAdjValue)
		var value float64
		value, err = util.Calculate(expr)
		if nil != err {
			er := fmt.Errorf("计算派生中间业务出错")
			zlog.Error(er.Error(), err)
			return nil, er
		}
		// fmt.Println("－－－－－－－－－－－－－－－", value)
		sceneItdDisctoun += value

	}
	// fmt.Println("+++++++++++++++", sceneItdDisctoun)
	pricing.DerviedIBPOints = sceneItdDisctoun
	return sceneItdDisctoun, nil
}

// 获取中收收益率
func (p PricingFormulaBus) getSceneIbYield(productCode string) (float64, error) {

	paramMap := map[string]interface{}{
		"product": productCode,
		"flag":    util.FLAG_TRUE,
	}
	sceneItdYields, err := sceneItdYieldService.Find(paramMap)
	if nil != err {
		return -1, err
	}
	var sceneItdYieldValue float64
	if 0 == len(sceneItdYields) {
		// sceneItdYieldValue = 1
		er := fmt.Errorf("对公贷款派生中间业务EVA收益率查询为空")
		zlog.Error(er.Error(), er)
		return -1, er
	} else if 1 < len(sceneItdYields) {
		er := fmt.Errorf("查询中收收益率有多条")
		return -1, er
	} else {
		sceneItdYieldValue = sceneItdYields[0].ItdYield
	}
	return sceneItdYieldValue, nil
}
