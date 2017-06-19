package functionIndicators

import (
	"fmt"

	"pccqcpa.com.cn/app/rpm/api/services/parService"
	"pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/components/zlog"
)

type CooperationPeriodDiscount struct {
}

var qualitativeDiscountService parService.QualitativeDiscountService

// 客户合作年限优惠点数
func (q *CooperationPeriodDiscount) Calulate(paramMap map[string]interface{}) (float64, error) {
	var discount float64
	paramMap["flag"] = "1"
	cooperationPeriod := paramMap["cooperation_period"]
	paramMap["stock_scene_type"] = util.COOPERATION_PERIOD
	delete(paramMap, "cooperation_period")
	var whereInSql = `(select max(stock_scene_val)  stock_scene_val  from rpm_par_qualitative_discount t  
					    where flag = 1
					      and stock_scene_type = '` + util.COOPERATION_PERIOD + `'
	                      and stock_scene_val <=` + fmt.Sprint(cooperationPeriod) + `)`
	paramMap["searchLike"] = []map[string]interface{}{
		map[string]interface{}{
			"type":  "in",
			"key":   "stock_scene_val",
			"value": whereInSql,
		},
	}

	qualitativeDiscounts, err := qualitativeDiscountService.Find(paramMap)
	if nil != err {
		er := fmt.Errorf("查询合作年限优惠点数出错")
		zlog.Error(er.Error(), err)
		return -1, er
	}

	switch len(qualitativeDiscounts) {
	case 0:
		zlog.Info("查询合作年限优惠点数纪录数为空，默认返回0", nil)
		return 0, nil
	case 1:
		discount = qualitativeDiscounts[0].Discount
	default:
		er := fmt.Errorf("查询合作年限优惠点数纪录数有多条")
		zlog.Error(er.Error(), er)
		return -1, er
	}

	return discount, nil
}
