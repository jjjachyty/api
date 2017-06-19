package functionIndicators

import (
	"fmt"

	"pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/components/zlog"
)

type AddRate struct {
}

// 查询增值税
func (this *AddRate) Calulate(paramMap map[string]interface{}) (float64, error) {
	paramMap["flag"] = util.FLAG_TRUE
	paramMap["tax_type"] = util.ADD_TAX
	// 加入生效期限
	err := util.GetStartTimeParam(paramMap)
	if nil != err {
		return 0, err
	}
	taxs, err := taxService.SelectEcByParmas(paramMap)
	if nil != err {
		var er = fmt.Errorf("查询增值税出错")
		zlog.Error(er.Error(), err)
		return -1, er
	}
	switch len(taxs) {
	case 0:
		var er = fmt.Errorf("查询增值税记录为空")
		zlog.Error(er.Error(), err)
		return -1, er
	case 1:
		return taxs[0].TaxRate, nil
	default:
		var er = fmt.Errorf("查询增值税记录有多条")
		zlog.Error(er.Error(), err)
		return -1, er
	}
	return -1, fmt.Errorf("查询增值税出错")
}
