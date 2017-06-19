package functionIndicators

import (
	"fmt"

	// "pccqcpa.com.cn/app/rpm/api/models/dim"
	"pccqcpa.com.cn/app/rpm/api/models/par"
	"pccqcpa.com.cn/app/rpm/api/models/sys"
	"pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/components/zlog"
)

// var taxService parService.TaxService

type CapCostRate struct {
	StartProductcode interface{}
}

var roc par.Roc

// 资本成本率
func (capCost *CapCostRate) Calulate(paramMap map[string]interface{}) (float64, error) {
	capCost.StartProductcode = paramMap["product"]
	rocOne, err := capCost.GetCapRate(paramMap)
	if nil != err {
		return -1, fmt.Errorf("查询资本成本率出错")
	}
	return rocOne.CapitalCost, nil
}

// 循环递归获取资本成本率或资本回报率
func (capCost *CapCostRate) GetCapRate(paramMap map[string]interface{}) (*par.Roc, error) {
	paramMap["flag"] = util.FLAG_TRUE

	// 加入生效期限
	err := util.GetStartTimeParam(paramMap)
	if nil != err {
		zlog.Error(err.Error(), err)
	}

	rocs, err := roc.SelectRocByParams(paramMap)
	if nil != err {
		return nil, err
	}
	rocsLength := len(rocs)
	switch rocsLength {
	case 0:
		//递归循环获取 优先循环产品、然后机构
		zlog.Infof("查询机构为[%s]产品为[%s]的资本成本率记录为空\n", nil, paramMap["organ"], paramMap["product"])
		// var er = fmt.Errorf("查询经济资本占用率记录为空")
		var productParamMap = map[string]interface{}{
			"product_code": paramMap["product"],
			"flag":         util.FLAG_TRUE,
		}
		products, err := productService.SelectPorductByParams(productParamMap)
		if nil != err {
			zlog.Error(err.Error(), err)
			return nil, err
		}
		if 1 == len(products) {
			//判断产品是否取到＃，如果已经取到了顶级产品还没有查到值，则机构向上，在循环产品
			if util.TREE_TOP_CODE == products[0].ParentProduct.ProductCode {
				organ, err := sys.SelectOrganByOrgCode(paramMap["organ"].(string))
				if nil != err {
					var er = fmt.Errorf("计算资本成本率时查询父级机构出错")
					zlog.Error(er.Error(), err)
					return nil, er
				}
				if util.TREE_TOP_CODE == organ.ParentOrgan {
					var er = fmt.Errorf("机构、产品都已经到达最上级，仍未查询到资本成本率或资本回报率参数")
					zlog.Error(er.Error(), er)
					return nil, er
				}
				paramMap["product"] = capCost.StartProductcode
				paramMap["organ"] = organ.ParentOrgan
			} else {
				paramMap["product"] = products[0].ParentProduct.ProductCode
			}
		}
		return capCost.GetCapRate(paramMap)
	case 1:
		return rocs[0], nil
	default:
		er := fmt.Errorf("查询资本回报率有多条数据")
		zlog.Error(er.Error(), er)
		return nil, er
	}
}
