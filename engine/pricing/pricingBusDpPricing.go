package pricing

import (
	"fmt"
	"strconv"
	"time"

	"pccqcpa.com.cn/app/rpm/api/models/dp"
	// "pccqcpa.com.cn/app/rpm/api/services/dpService"
	"pccqcpa.com.cn/app/rpm/api/engine/functionIndicators"
	"pccqcpa.com.cn/app/rpm/api/services/parService"
	"pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/components/zlog"
)

// 存款一对一派生当前业务EVA 资金成本率
func (PricingBus) DpFtpRateCalculate(dpBusiness *dp.DpOneBusiness, businessType string) (interface{}, error) {
	paramMap := map[string]interface{}{
		"currency":   dpBusiness.Ccy,
		"term":       dpBusiness.Term,
		"organ":      dpBusiness.Organ.OrganCode,
		"product":    dpBusiness.Product.ProductCode,
		"param_type": util.PARAM_DP,
	}
	rate, err := ftpRate.Calulate(paramMap)
	if nil != err {
		return nil, err
	}
	return rate, err
}

// 存款一对一 当前业务EVA 存款保险费率
// if 金额 <= 50W return val
// if 金额 >  50W reutrn (50W * val) / 金额
func (PricingBus) DpInsuranceRateCalculate(dpBusiness *dp.DpOneBusiness, businessType string) (interface{}, error) {
	paramMap := map[string]interface{}{
		"flag":    util.FLAG_TRUE,
		"par_mod": util.PARAM_DP,
		"par_key": util.PREMIUM_RATE,
	}
	commonVal, err := parService.CommonService{}.FindOne(paramMap)
	if nil != err {
		return nil, fmt.Errorf("查询公共参数存款保险费率出错")
	}
	if nil == commonVal {
		return nil, fmt.Errorf("查询公共参数存款保险费率为空")
	}
	val, err := strconv.ParseFloat(commonVal.ParValue, 64)
	if nil != err {
		return nil, fmt.Errorf("查询公共参数存款保险费率不是一个有效数字")
	}

	// 存款金额
	DpAmount := dpBusiness.Amount
	var AmountLimit float64 = 500000 // 金额额度
	if AmountLimit < DpAmount {
		expr := fmt.Sprint(AmountLimit) + "*" + fmt.Sprint(val) + "/" + fmt.Sprintf("%f", DpAmount)
		rate, err := util.Calculate(expr)
		if nil != err {
			return nil, fmt.Errorf("计算存款保险费出错【%s】", expr)
		}
		return rate, nil
	}

	return val, nil
}

// 操作风险费率
// 与机构产品挂钩，如果没有数据，则机构往上查询
func (PricingBus) DpOperationRiskRateCalculate(dpBusiness *dp.DpOneBusiness, businessType string) (interface{}, error) {

	var paramMap = map[string]interface{}{
		"product": dpBusiness.Product.ProductCode,
		"organ":   dpBusiness.Organ.OrganCode,
		"flag":    util.FLAG_TRUE,
	}
	rate, err := functionIndicators.DpOperationRiskRate{}.Calulate(paramMap)
	if nil != err {
		return nil, err
	}
	return rate, nil
}

// 存款一对一派生当前业务EVA 运营费用率
func (PricingBus) DpOcCalculate(dpBusiness *dp.DpOneBusiness, businessType string) (interface{}, error) {
	paramMap := map[string]interface{}{
		"organ":      dpBusiness.Organ.OrganCode,
		"product":    dpBusiness.Product.ProductCode,
		"param_type": util.PARAM_DP,
	}
	rate, err := ocRate.Calulate(paramMap)
	if nil != err {
		return nil, err
	}
	return rate, nil
}

// 存款一对一派生当前业务EVA 执行利率
func (PricingBus) DpInitRateCalculate(dpBusiness *dp.DpOneBusiness, businessType string) (interface{}, error) {
	return dpBusiness.Rate, nil
}

// 存款一对一派生当前业务EVA 存款金额
func (PricingBus) DpAmountCalculate(dpBusiness *dp.DpOneBusiness, businessType string) (interface{}, error) {
	return dpBusiness.Amount, nil
}

// 存款一对一派生当前业务EVA 存款期限
// 期限超过1年按1年计算，不足1年按实际期限计算
// RPM业务邓强 指示 使用实际期限 如果为活期则期限默认为1年
func (PricingBus) DpTermCalculate(dpBusiness *dp.DpOneBusiness, businessType string) (interface{}, error) {
	// if 360 <= dpBusiness.Term {
	// 	return 1, nil
	// }
	if 0 == dpBusiness.Term {
		return 1, nil
	}
	expr := fmt.Sprint(dpBusiness.Term) + "/" + fmt.Sprint(360)
	dpTerm, err := util.Calculate(expr)
	if nil != err {
		return nil, fmt.Errorf("处理存款日期出错【%s】", expr)
	}
	return dpTerm, nil
}

// 存款一对一派生中间业务收入EVA 中间业务收入
func (PricingBus) DpIbIncomeCalculate(ibBusiness *dp.DpIbBusiness, businessType string) (interface{}, error) {
	return ibBusiness.Amount, nil
}

// 存款一对一派生中间业务收入EVA 中间业务收益率
func (PricingBus) DpIbYieldCalculate(ibBusiness *dp.DpIbBusiness, businessType string) (interface{}, error) {
	ibYield, err := PricingFormulaBus{}.getSceneIbYield(ibBusiness.Product.ProductCode)
	if nil != err {
		return nil, err
	}
	return ibYield, nil
}

// 存款一对一存量-原执行利率
func (PricingBus) DpOneStockRateCalculate(dpOneStockPricing *dp.DpOneStockPricing, businessType string) (interface{}, error) {
	return dpOneStockPricing.RateOld, nil
}

// 存款一对一存量-客户要求执行利率
func (PricingBus) DpOneStockCustRateCalculate(dpOneStockPricing *dp.DpOneStockPricing, businessType string) (interface{}, error) {
	return dpOneStockPricing.RateNew, nil
}

// 存款一对一存量-存款金额
func (PricingBus) DpOneStockAmountCalculate(dpOneStockPricing *dp.DpOneStockPricing, businessType string) (interface{}, error) {
	return dpOneStockPricing.AmountOld, nil
}

// 存款一对一存量-后续期限
func (PricingBus) DpOneStockRestTermCalculate(dpOneStockPricing *dp.DpOneStockPricing, businessType string) (interface{}, error) {
	// 如果为活期则返回 360
	if 0 == dpOneStockPricing.RestTerm {
		zlog.Info("活期产品，期限取1年", nil)
		return 360, nil
	}
	// 到期日算法
	endTime, _ := time.Parse("2006-01-02", dpOneStockPricing.EndOfDate.Format("2006-01-02"))
	nowTime, _ := time.Parse("2006-01-02", time.Now().Format("2006-01-02"))
	val := int(endTime.Sub(nowTime) / time.Hour / 24)
	if val <= 0 {
		return 0, nil
	}
	return val, nil

	// 直接使用剩余期限
	// return dpOneStockPricing.RestTerm, nil
}

// 存款一对一存量-已存期限
func (PricingBus) DpOneStockSaveTermCalculate(dpOneStockPricing *dp.DpOneStockPricing, businessType string) (interface{}, error) {
	endTime, _ := time.Parse("2006-01-02", dpOneStockPricing.EndOfDate.Format("2006-01-02"))
	nowTime, _ := time.Parse("2006-01-02", time.Now().Format("2006-01-02"))
	startTime, _ := time.Parse("2006-01-02", dpOneStockPricing.StartOfDate.Format("2006-01-02"))
	val := endTime.Sub(nowTime) / time.Hour / 24
	if val <= 0 {
		return int(endTime.Sub(startTime) / time.Hour / 24), nil
	}
	return int(nowTime.Sub(startTime) / time.Hour / 24), nil
}

// 存款一对一存量-资金收益
func (PricingBus) DpOneStockFtpCalculate(dpOneStockPricing *dp.DpOneStockPricing, businessType string) (interface{}, error) {
	var term int = 0
	if 0 != dpOneStockPricing.RestTerm {
		endTime, _ := time.Parse("2006-01-02", dpOneStockPricing.EndOfDate.Format("2006-01-02"))
		startTime, _ := time.Parse("2006-01-02", dpOneStockPricing.StartOfDate.Format("2006-01-02"))
		term = int(endTime.Sub(startTime) / time.Hour / 24)
	}
	paramMap := map[string]interface{}{
		"currency":   dpOneStockPricing.Ccy,
		"term":       term,
		"organ":      dpOneStockPricing.Organ.OrganCode,
		"product":    dpOneStockPricing.Product.ProductCode,
		"param_type": util.PARAM_DP,
	}
	rate, err := ftpRate.Calulate(paramMap)
	if nil != err {
		return nil, err
	}
	return rate, err
	return 1, nil
}

// 一对一存款存量-存款保险费率
func (PricingBus) DpOneStockInsuranceRateCalculate(dpOneStockPricing *dp.DpOneStockPricing, businessType string) (interface{}, error) {
	paramMap := map[string]interface{}{
		"flag":    util.FLAG_TRUE,
		"par_mod": util.PARAM_DP,
		"par_key": util.PREMIUM_RATE,
	}
	commonVal, err := parService.CommonService{}.FindOne(paramMap)
	if nil != err {
		return nil, fmt.Errorf("查询公共参数存款保险费率出错")
	}
	if nil == commonVal {
		return nil, fmt.Errorf("查询公共参数存款保险费率为空")
	}
	val, err := strconv.ParseFloat(commonVal.ParValue, 64)
	if nil != err {
		return nil, fmt.Errorf("查询公共参数存款保险费率不是一个有效数字")
	}
	return val, nil

}

// 一对一存款存量-运营成本
func (PricingBus) DpOneStockOcRateCalculate(dpOneStockPricing *dp.DpOneStockPricing, businessType string) (interface{}, error) {
	paramMap := map[string]interface{}{
		"organ":      dpOneStockPricing.Organ.OrganCode,
		"product":    dpOneStockPricing.Product.ProductCode,
		"param_type": util.PARAM_DP,
	}
	rate, err := ocRate.Calulate(paramMap)
	if nil != err {
		return nil, err
	}
	return rate, nil
}

// 一对一存款存量-操作风险成本率
func (PricingBus) DpOneStockOperationRiskRateCalculate(dpOneStockPricing *dp.DpOneStockPricing, businessType string) (interface{}, error) {
	paramMap := map[string]interface{}{
		"organ":   dpOneStockPricing.Organ.OrganCode,
		"product": dpOneStockPricing.Product.ProductCode,
	}
	rate, err := functionIndicators.DpOperationRiskRate{}.Calulate(paramMap)
	if nil != err {
		er := fmt.Errorf("查询存款操作风险率为空:机构编码【%v】，产品编码【%v】", dpOneStockPricing.Organ.OrganCode,
			dpOneStockPricing.Product.ProductCode)
		zlog.Error(er.Error(), er)
		return -1, er
	} else {
		return rate, nil
	}
}

// 一对一存款存量-流失存款金额
func (PricingBus) DpOneStockLostAmountCalculate(dpOneStockPricing *dp.DpOneStockPricing, businessType string) (interface{}, error) {
	return dpOneStockPricing.AmountUse, nil
}

// 存款一对一存款存量-违约执行利率
func (PricingBus) DpOneStockPdInitRateCalculate(dpOneStockPricing *dp.DpOneStockPricing, businessType string) (interface{}, error) {

	common, err := parService.CommonService{}.FindOne(map[string]interface{}{
		"flag":    util.FLAG_TRUE,
		"par_key": "DP_ONE_PRICING_CURRENT_RATE",
	})
	if nil != err {
		return nil, err
	}
	if nil == common {
		er := fmt.Errorf("查询公共参数表存款一对一定价违约执行利率为空")
		zlog.Error(er.Error(), er)
		return nil, er
	}
	rate, err := strconv.ParseFloat(common.ParValue, 64)
	if nil != err {
		er := fmt.Errorf("公共参数表[存款一对一定价违约执行利率]参数为非数字")
		zlog.Error(er.Error(), err)
		return nil, er
	}
	return rate, nil
}

// 存款一对一存款存量威胁-当前余额
func (PricingBus) DpOneStockCurrentAmountCalculate(dpOneStockPricing *dp.DpOneStockPricing, businessType string) (interface{}, error) {
	return dpOneStockPricing.CurrentAmount, nil
}
