package functionIndicators

import (
	"fmt"

	"pccqcpa.com.cn/app/rpm/api/models/par"
	// "pccqcpa.com.cn/app/rpm/api/models/sys"
	"pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/components/zlog"
)

type PdRate struct {
}

var pd par.Pd

func (this *PdRate) Calulate(paramMap map[string]interface{}) (float64, error) {
	paramMap["flag"] = util.FLAG_TRUE

	// 加入生效期限
	err := util.GetStartTimeParam(paramMap)
	if nil != err {
		return 0, err
	}

	pds, err := pd.SelectPdByParams(paramMap)
	if nil != err {
		zlog.Error("查询违约概率失败", err)
		return -1, err
	}
	switch len(pds) {
	case 0:
		err := fmt.Errorf("查询违约概率为空")
		zlog.Error(err.Error(), err)
		return -1, err
	case 1:
		return pds[0].PdRate, nil
	default:
		err := fmt.Errorf("查询违约概率有多条")
		zlog.Error(err.Error(), err)
		return -1, err
	}
	return -1, fmt.Errorf("查询违约出错")
}
