package functionIndicators

import (
	"pccqcpa.com.cn/app/rpm/api/custom"
)

var InitFunctionIndicators map[string]interface{}

// 初始化各个函数指标
func init() {
	InitFunctionIndicators = map[string]interface{}{
		"FtpRate":       new(FtpRate),       //资金成本率
		"FtpRateSocket": new(FtpRateSocket), //资金成本率－接口
		"EcRate":        new(EcRate),        //经济资本占用率
		"IncomeTax":     new(IncomeRate),    //所得税
		"AddTax":        new(AddRate),       //增值税
		"CapCostRate":   new(CapCostRate),   //资本成本率
		"CapPftRate":    new(CapPftRate),    //资本利润率
		"LgdRate":       new(LgdRate),       //预期损失率
		"OcRate":        new(OcRate),        //运营费用率
		"PdRate":        new(PdRate),        //违约概率

		"TermDay": new(TermDay), //贷款实际天数
	}

	// 调用客户化开发Init函数，覆盖函数追包Map达到客户化开发目的
	custom.Init(InitFunctionIndicators)
}
