package procedure

import (
	"fmt"
	"time"

	"pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/components/zlog"
)

// 贷款交叉分析-多维度 存储过程
// by author Jason
// by time 2016-11-21 16:26:56
type pricingAnalysisCross struct {
	procedureName string    // 存储过程名称
	asOfDate      time.Time // 数据日期
	dimOneCode    string    // 维度一
	dimOneName    string
	dimTwoCode    string // 维度二
	dimTwoName    string
}

func NewPricingAnalysisCorss(date time.Time) pricingAnalysisCross {
	return pricingAnalysisCross{
		procedureName: "PROC_PRICING_ANALYSIS_CRO",
		asOfDate:      date,
	}
}

func (p pricingAnalysisCross) WithDimOne(dimOneCode, dimOneName string) pricingAnalysisCross {
	p.dimOneCode = dimOneCode
	p.dimOneName = dimOneName
	return p
}

func (p pricingAnalysisCross) WithDimTwo(dimTwoCode, dimTwoName string) pricingAnalysisCross {
	p.dimTwoCode = dimTwoCode
	p.dimTwoName = dimTwoName
	return p
}

func (p pricingAnalysisCross) ExecProcess() error {
	var retFlg, retMsg *string
	err := util.OracleExecProcedure(
		p.procedureName,
		p.asOfDate,
		p.dimOneCode,
		p.dimOneName,
		p.dimTwoCode,
		p.dimTwoName,
		retFlg,
		retMsg,
	)
	zlog.Infof("执行存储过程与参数【%#v】", nil, p)
	if nil != err {
		er := fmt.Errorf("执行存储过程出错【%s】", p.procedureName)
		zlog.Error(err.Error(), err)
		return er
	}
	return nil
}
