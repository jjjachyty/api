package analysisService

import (
	"fmt"
	"strings"
	"time"

	"pccqcpa.com.cn/app/rpm/api/models/bi/analysis"
	"pccqcpa.com.cn/app/rpm/api/models/charts"
	"pccqcpa.com.cn/app/rpm/api/models/procedure"
	"pccqcpa.com.cn/app/rpm/api/services/parService"
	"pccqcpa.com.cn/app/rpm/api/util"
)

type AnalysisCrossDimService struct{}

var analysisCrossDim analysis.AnalysisCrossDim

// 分页操作
// by author Jason
// by time 2016-11-21 16:26:56
func (AnalysisCrossDimService) List(param ...map[string]interface{}) (*util.PageData, error) {
	return analysisCrossDim.List(param...)
}

// 多参数查询返回多条纪录
// 先执行交叉分析的存储过程，然后在返回结果数据
// by author Jason
// by time 2016-11-21 16:26:56
func (AnalysisCrossDimService) Find(param ...map[string]interface{}) (*charts.AnalysisCrossDimChart, error) {
	if 0 < len(param) {
		paramMap := param[0]
		dimOneCode := paramMap["dim_one_code"].(string)
		dimTwoCode := paramMap["dim_two_code"].(string)
		asOfDate, _ := time.Parse("2006-01-02", paramMap["as_of_date"].(string))
		paramMap["as_of_date"] = asOfDate

		// 处理NAME列
		var dimOneName, dimTwoName string
		var dictService parService.DictService
		dicts, err := dictService.Find(map[string]interface{}{"parent_dict": "DIM_ONE_NAME"})
		if nil != err {
			return nil, err
		}
		for _, dict := range dicts {
			switch strings.ToUpper(dict.DictCode) {
			case strings.ToUpper(dimOneCode):
				dimOneName = dict.DictName
			case strings.ToUpper(dimTwoCode):
				dimTwoName = dict.DictName
			}
		}

		proc := procedure.NewPricingAnalysisCorss(asOfDate).
			WithDimOne(dimOneCode, dimOneName).
			WithDimTwo(dimTwoCode, dimTwoName)
		err = proc.ExecProcess()
		if nil != err {
			return nil, err
		}
		return analysisCrossDim.Find(param...)
	}
	er := fmt.Errorf("未传dimOneCode,dimTowCode,asOfDate参数")
	return nil, er
}

// 新增纪录
// by author Jason
// by time 2016-11-21 16:26:56
func (AnalysisCrossDimService) Add(model *analysis.AnalysisCrossDim) error {
	return model.Add()
}

// 更新纪录
// by author Jason
// by time 2016-11-21 16:26:56
func (AnalysisCrossDimService) Update(model *analysis.AnalysisCrossDim) error {
	return model.Update()
}

// 删除纪录
// by author Jason
// by time 2016-11-21 16:26:56
func (AnalysisCrossDimService) Delete(model *analysis.AnalysisCrossDim) error {
	return model.Delete()
}
