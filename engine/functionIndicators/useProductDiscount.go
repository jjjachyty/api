package functionIndicators

import (
	"fmt"

	"pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/components/zlog"
)

type UseProductDiscount struct {
}

// 客户使用产品数优惠点数
func (q *UseProductDiscount) Calulate(paramMap map[string]interface{}) (float64, error) {
	var discount float64
	paramMap["flag"] = "1"
	useProduct := paramMap["use_product"]
	paramMap["stock_scene_type"] = util.USE_PRODUCT
	delete(paramMap, "use_product")
	var whereInSql = `(select max(stock_scene_val)  stock_scene_val  from rpm_par_qualitative_discount t  
	                    where flag = 1
	                      and stock_scene_type = '` + util.USE_PRODUCT + `'
	                      and stock_scene_val <=` + fmt.Sprint(useProduct) + `)`
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
		zlog.Info("查询合作年限优惠点数纪录数为空,默认返回0", nil)
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
