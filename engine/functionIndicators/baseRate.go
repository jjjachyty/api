package functionIndicators

import (
	"fmt"

	"pccqcpa.com.cn/app/rpm/api/services/parService"
	"pccqcpa.com.cn/app/rpm/api/util"
)

var baseRateService parService.BaseRateService

type BaseRate struct {
}

// 查询所得税
func (b BaseRate) Calulate(paramMap map[string]interface{}) (float64, error) {
	var term int
	// fmt.Println("\n\n\n\n\n基准利率参数", paramMap)
	value, ok := paramMap["term"]
	if ok {
		term = value.(int)
	}
	delete(paramMap, "term")
	var whereInSql = `
		select max(term)  term  from rpm_par_base_rate t  
		 where flag = 1 
		   and base_rate_type = ` + paramMap["base_rate_type"].(string) + ` 
		   and term <` + fmt.Sprint(term)
	// + `   union all
	// select min(term)  term  from rpm_par_base_rate t  where term >=` + fmt.Sprint(term)
	paramMap["searchLike"] = []map[string]interface{}{
		map[string]interface{}{
			"type":  "in",
			"key":   "term",
			"value": whereInSql,
		},
	}
	paramMap["flag"] = util.FLAG_TRUE
	baseRate, err := baseRateService.Find(paramMap)
	if nil != err {
		er := fmt.Errorf("查询基准利率出错")
		return -1, er
	}
	if 0 == len(baseRate) {
		er := fmt.Errorf("查询基准利率为空")
		return -1, er
	}
	if 1 < len(baseRate) {
		er := fmt.Errorf("查询基准利率有多条")
		return -1, er
	}
	// fmt.Printf("\n\n\n\n基准利率[%v]\n\n\n\n", baseRate[0].Rate)
	return baseRate[0].Rate, nil
}
