package functionIndicators

import (
	"fmt"

	// "pccqcpa.com.cn/app/rpm/api/models/dim"
	"pccqcpa.com.cn/app/rpm/api/models/par"
	"pccqcpa.com.cn/app/rpm/api/models/sys"
	"pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/components/zlog"
)

type OcRate struct {
}

var oc par.Oc

var extremeOc float64

// OCRATE
// 获取运营费用率
// 参数1:机构
// 参数2:产品
// 参数3:客户规模
// 返回值:运营费用率
// 逻辑: 根据参数查询记录，如果没有查询到记录则根据机构查询父级 如果取到值了，优先取手工值，如果没有手工值取auto_oc
// if 分支行费用率 > 总行费用率（优先手工值）* 极值
//		＝ 总行费用率
// else
// 		= 分支行费用率
func (this *OcRate) Calulate(paramMap map[string]interface{}) (float64, error) {
	//产品、机构是结构体，获取码值
	// if organ, ok := paramMap["organ"]; ok {
	// 	paramMap["organ"] = organ.(sys.Organ).OrganCode
	// }
	// if product, ok := paramMap["product"]; ok {
	// 	paramMap["product"] = product.(dim.Product).ProductCode
	// }
	// 获取总行运营费用率与oc极值
	topOc, err := this.getTopOrganOcRate(paramMap)
	if nil != err {
		return -1, err
	}

	// 加入生效期限
	err = util.GetStartTimeParam(paramMap)
	if nil != err {
		return 0, err
	}

	// 获取当前机构的运用费用率，如果没有则递归机构向上查询
	ocRate, err := this.getOcRate(topOc, paramMap)
	if nil != err {
		zlog.Error("递归机构获取运营费用率出错", err)
		return -1, err
	}
	zlog.Infof("计算运营费用率为[%f]", nil, ocRate)
	return ocRate, nil
}

// 获取总行运营费用率
func (this *OcRate) getTopOrganOcRate(paramMap map[string]interface{}) (float64, error) {
	// 获取总行机构码值
	var organParam = map[string]interface{}{"parent_organ": "#"}
	organs, err := sys.SelectOrganByParams(organParam)
	if nil != err {
		er := fmt.Errorf("查询总行机构出错")
		zlog.Error("查询总行机构出错", err)
		return -1, er
	}

	businessOrgan := paramMap["organ"]
	defer func() { paramMap["organ"] = businessOrgan }()
	paramMap["organ"] = organs[0].OrganCode
	paramMap["flag"] = util.FLAG_TRUE
	ocs, err := oc.Find(paramMap)
	if nil != err {
		zlog.Error("查询总行运营费用率错误", err)
		return -1, err
	}
	switch len(ocs) {
	case 0:
		err := fmt.Errorf("获取总行[%v]产品为[%v]运营费用率为空", paramMap["organ"], paramMap["product"])
		zlog.Errorf(err.Error(), err)
		return -1, err
	case 1:
		if 0 != ocs[0].ManualOc {
			return ocs[0].ManualOc, nil
		} else {
			return ocs[0].AutoOc, nil
		}
	default:
		err := fmt.Errorf("查询总行运营费用率有多条")
		zlog.Error(err.Error(), err)
		return -1, err
	}
}

// 递归查询运用费用率
func (this *OcRate) getOcRate(topOc float64, paramMap map[string]interface{}) (float64, error) {
	organCode := paramMap["organ"].(string)

	paramMap["flag"] = util.FLAG_TRUE
	ocs, err := oc.Find(paramMap)
	if nil != err {
		zlog.Errorf("查询机构[%s]的运用费用率错误", nil, organCode)
		return -1, err
	}
	switch len(ocs) {
	case 0:
		zlog.Infof("查询机构[%s]运营费用率记录为空，向上查询", nil, organCode)
		organ, err := sys.SelectOrganByOrgCode(organCode)
		if nil != err {
			zlog.Errorf("查询机构码值为[%s]的机构错误", err, organCode)
			return -1, err
		}
		paramMap["organ"] = organ.ParentOrgan
		return this.getOcRate(topOc, paramMap)
	case 1:
		if 0 != ocs[0].ManualOc {
			// return this.compareOcRate(ocs[0].ManualOc, topOc), nil
			return ocs[0].ManualOc, nil
		}
		return this.compareOcRate(ocs[0].AutoOc, topOc), nil
	default:
		err := fmt.Errorf("查询机构[%s]运营费用率有多条", organCode)
		zlog.Error(err.Error(), err)
		return -1, err
	}
	return -1, fmt.Errorf("递归查询运营费用率错误")
}

func (this *OcRate) compareOcRate(ocRate, topOc float64) float64 {
	if ocRate >= topOc*extremeOc {
		return topOc * extremeOc
	}
	return ocRate
}

// 获取总行机构运营费用率
// 获取运营费用率极值
func init() {
	value, err := getOcExtreme()
	if nil != err {
		zlog.Error("查询Oc极值错误", err)
	}
	extremeOc = value
}

// 获取oc极值
func getOcExtreme() (float64, error) {
	return 1.4, nil
}
