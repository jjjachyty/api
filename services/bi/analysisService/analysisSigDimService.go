package analysisService

import (
	"errors"
	"strings"
	"time"
	// "sync"

	"pccqcpa.com.cn/app/rpm/api/models/bi/analysis"

	"pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/components/zlog"
)

type AnalysisSigDimService struct{}

var analysisSigDim analysis.AnalysisSigDim

// var findMutex sync.Mutex

// 分页操作
// by author Jason
// by time 2016-11-15 16:10:56
func (AnalysisSigDimService) List(param ...map[string]interface{}) (*util.PageData, error) {
	return analysisSigDim.List(param...)
}

// 多参数查询返回多条纪录
// by author Jason
// by time 2016-11-15 16:10:56
func (AnalysisSigDimService) Find(param ...map[string]interface{}) ([]*analysis.AnalysisSigDim, error) {
	// findMutex.Lock()
	// defer findMutex.Unlock()
	if 0 != len(param) {
		paramMap := param[0]
		if dimType, ok := paramMap["dim_type"]; ok {
			paramMap["dim_type"] = strings.ToUpper(dimType.(string))
			switch strings.ToUpper(dimType.(string)) {
			case "ORGAN", "AREA_TWO", "AREA_ONE", "PRODUCT", "INDUSTRY_CODE":
				paramMap["sort"] = "floating_level"
				paramMap["order"] = "DESC"
				return analysisSigDim.Find(paramMap)
			case "CUST_SCALE_CODE_LINE": // 客户规模行标
				paramMap["sort"] = "T1.sort"
				return analysisSigDim.FindToSort("(select code,sort from rpm_bi_dim_scale where type = '行标') T1", "T1.code", paramMap)
			case "CUST_SCALE_CODE_GB": //客户规模国标
				paramMap["sort"] = "T1.sort"
				return analysisSigDim.FindToSort("(select code,sort from rpm_bi_dim_scale where type = '国标') T1", "T1.code", paramMap)
			case "TERM_INTERVAL": //期限区间
				paramMap["sort"] = "T1.sort"
				return analysisSigDim.FindToSort("(select term_code,sort from RPM_BI_DIM_TERM_RANGE) T1", "T1.TERM_CODE", paramMap)
			case "AMOUNT_RANGE":
				paramMap["sort"] = "T1.sort"
				return analysisSigDim.FindToSort("(select range,sort from RPM_BI_DIM_AMOUNT_RANGE) T1", "T1.RANGE", paramMap)
			case "CREDIT_RATING": //信用等级
				paramMap["sort"] = "T1.sort"
				return analysisSigDim.FindToSort("(select rating_code,sort from RPM_BI_DIM_CREDIT_RATING) T1", "T1.RATING_CODE", paramMap)
			}
		}
	}
	return analysisSigDim.Find(param...)
}

// 多参数查询返回一条纪录
// by author Jason
// by time 2016-11-15 16:10:56
func (this AnalysisSigDimService) FindOne(paramMap map[string]interface{}) (*analysis.AnalysisSigDim, error) {
	models, err := this.Find(paramMap)
	if nil != err {
		return nil, err
	}
	switch len(models) {
	case 0:
		return nil, nil
	case 1:
		return models[0], nil
	}
	er := errors.New("查询【RPM_BI_ANALYSIS_SIG_DIM】贷款业务分析方案单维度分析有多条返回纪录")
	zlog.Error(er.Error(), er)
	return nil, er
}

// 执行存储过程
// by author Jason
// by time 2016-11-15 16:10:56
func (this AnalysisSigDimService) ExecProcedure(p_as_of_date time.Time) error {
	return analysisSigDim.ExecProcedure(p_as_of_date)
}

// 查询日期
// by author Jason
// by time 2016-11-15 16:10:56
func (this AnalysisSigDimService) FindDistinctDate() ([]string, error) {
	return analysisSigDim.FindDistinctDate()
}

// 新增纪录
// by author Jason
// by time 2016-11-15 16:10:56
func (AnalysisSigDimService) Add(model *analysis.AnalysisSigDim) error {
	return model.Add()
}

// 更新纪录
// by author Jason
// by time 2016-11-15 16:10:56
func (AnalysisSigDimService) Update(model *analysis.AnalysisSigDim) error {
	return model.Update()
}

// 删除纪录
// by author Jason
// by time 2016-11-15 16:10:56
func (AnalysisSigDimService) Delete(model *analysis.AnalysisSigDim) error {
	return model.Delete()
}
