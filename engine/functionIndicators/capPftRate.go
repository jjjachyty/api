package functionIndicators

import (
	"fmt"
	// "pccqcpa.com.cn/app/rpm/api/services/parService"
	// "pccqcpa.com.cn/app/rpm/api/util"
	// "pccqcpa.com.cn/components/zlog"
	// "pccqcpa.com.cn/app/rpm/api/models/dim"
	// "pccqcpa.com.cn/app/rpm/api/models/sys"
)

// var taxService parService.TaxService

type CapPftRate struct {
}

var capCost CapCostRate

// 资本利润率
func (this *CapPftRate) Calulate(paramMap map[string]interface{}) (float64, error) {
	capCost.StartProductcode = paramMap["product"]
	rocOne, err := capCost.GetCapRate(paramMap)
	if nil != err {
		return -1, fmt.Errorf("查询资本回报率出错")
	}
	return rocOne.CapitalProfit, nil
}
