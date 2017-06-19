package functionIndicators

import (
	"fmt"

	// "pccqcpa.com.cn/app/rpm/api/models/dim"
	"pccqcpa.com.cn/app/rpm/api/models/sys"
	"pccqcpa.com.cn/app/rpm/api/services/dimService"
	"pccqcpa.com.cn/app/rpm/api/services/parService"
	"pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/components/zlog"
)

type EcRate struct {
}

var ecService parService.EcService
var productService dimService.ProductService

var startProductcode string

// 经济资本占用率
// 递归查询－优先产品查询，然后机构递归查询
func (this *EcRate) Calulate(paramMap map[string]interface{}) (float64, error) {

	paramMap["flag"] = util.FLAG_TRUE
	// organ, ok := paramMap["organ"]
	// if !ok {
	// 	var er error = fmt.Errorf("计算经济资本占用率未传机构参数")
	// 	return -1, er
	// }
	// product, ok := paramMap["product"]
	// if !ok {
	// 	var er error = fmt.Errorf("计算经济资本占用率未传产品参数")
	// 	return -1, er
	// }
	// var organCode, productCode string = organ.(sys.Organ).OrganCode, product.(dim.Product).ProductCode
	startProductcode = paramMap["product"].(string)

	ecParamMap := map[string]interface{}{
		"organ":   paramMap["organ"],
		"product": paramMap["product"],
		"flag":    util.FLAG_TRUE,
	}

	// 加入生效期限
	err := util.GetStartTimeParam(ecParamMap)
	if nil != err {
		return 0, err
	}

	ecRate, err := this.getEcRate(ecParamMap)
	if nil != err {
		return -1, err
	}
	return ecRate, nil
}

// 递归查询－优先产品查询，然后机构递归查询
func (this *EcRate) getEcRate(paramMap map[string]interface{}) (float64, error) {
	ecs, err := ecService.SelectEcByParmas(paramMap)
	if nil != err {
		zlog.Error("计算经济资本占用率", err)
		return -1, err
	}
	switch len(ecs) {
	case 0:
		zlog.Infof("查询机构为[%s]产品为[%s]的经济资本占用率记录为空", nil, paramMap["organ"], paramMap["product"])
		// var er = fmt.Errorf("查询经济资本占用率记录为空")
		var productParamMap = map[string]interface{}{
			"product_code": paramMap["product"],
			"flag":         util.FLAG_TRUE,
		}
		products, err := productService.SelectPorductByParams(productParamMap)
		if nil != err {
			zlog.Error(err.Error(), err)
			return -1, err
		}
		if 1 == len(products) {
			//判断产品是否取到＃，如果已经取到了顶级产品还没有查到值，则机构向上，在循环产品
			if util.TREE_TOP_CODE == products[0].ParentProduct.ProductCode {
				fmt.Println("---------查询父级机构----------")
				organ, err := sys.SelectOrganByOrgCode(paramMap["organ"].(string))
				if nil != err {
					var er = fmt.Errorf("计算经济资本占用率时查询父级机构出错")
					zlog.Error(er.Error(), err)
					return -1, er
				}
				if util.TREE_TOP_CODE == organ.ParentOrgan {
					var er = fmt.Errorf("机构、产品都已经到达最上级，仍未查询到参数")
					zlog.Error(er.Error(), er)
					return -1, er
				}
				paramMap["product"] = startProductcode
				paramMap["organ"] = organ.ParentOrgan
			} else {
				paramMap["product"] = products[0].ParentProduct.ProductCode
			}
		}
		return this.getEcRate(paramMap)
	case 1:
		return ecs[0].EcRate, nil
	default:
		var er = fmt.Errorf("计算经济资本占用率时查询有多条经济资本占用率记录")
		return -1, er
	}

	return -1, fmt.Errorf("递归机构、产品查询经济资本占用率出错")
}
