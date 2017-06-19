package analysisService

import (
	"fmt"

	"pccqcpa.com.cn/app/rpm/api/models/bi/analysis"
	"pccqcpa.com.cn/app/rpm/api/util"
)

type SigDimAdvice struct{}

func (sig SigDimAdvice) CustScaleGB(pageData *util.PageData) string {
	var returnMsg string
	var flag bool
	anSigDim := pageData.Rows.([]*analysis.AnalysisSigDim)
	length := len(anSigDim)
	fmt.Println(flag, length)
	return returnMsg
}
