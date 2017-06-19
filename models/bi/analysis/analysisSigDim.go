package analysis

import (
	"database/sql"
	"fmt"
	"time"

	modelsUtil "pccqcpa.com.cn/app/rpm/api/models/util"
	"pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/components/zlog"
	"platform/dbobj"
)

// 结构体对应数据库表【RPM_BI_ANALYSIS_SIG_DIM】贷款业务分析方案单维度分析
// by author Jason
// by time 2016-11-15 16:10:56
type AnalysisSigDim struct {
	UUID          string    // UIUID
	DimCode       string    // 维度代码
	DimName       string    // 维度名称
	DimType       string    // 维度类型
	CustNum       float64   // 客户数
	BusNum        float64   // 业务笔数
	AmountTotal   float64   // 发放金额
	AmountPercent float64   // 金额占比金额占比
	BusPercent    float64   // 笔数占比
	AmountAvg     float64   // 户均金额
	IntRate       float64   // 执行利率
	FloatingLevel float64   // 浮动水平
	TermAvg       float64   // 平均期限
	AsOfDate      time.Time // 数据日期
	Sort          int       // 排序
}

// 结构体scan
// by author Jason
// by time 2016-11-15 16:10:56
func (this AnalysisSigDim) scan(rows *sql.Rows) (*AnalysisSigDim, error) {
	var one = new(AnalysisSigDim)
	values := []interface{}{
		&one.UUID,
		&one.DimCode,
		&one.DimName,
		&one.DimType,
		&one.CustNum,
		&one.BusNum,
		&one.AmountTotal,
		&one.AmountPercent,
		&one.BusPercent,
		&one.AmountAvg,
		&one.IntRate,
		&one.FloatingLevel,
		&one.TermAvg,
		&one.AsOfDate,
	}
	err := util.OracleScan(rows, values)
	return one, err
}

// 结构体scan
// by author Jason
// by time 2016-11-15 16:10:56
func (this AnalysisSigDim) scanToSort(rows *sql.Rows) (*AnalysisSigDim, error) {
	var one = new(AnalysisSigDim)
	values := []interface{}{
		&one.UUID,
		&one.DimCode,
		&one.DimName,
		&one.DimType,
		&one.CustNum,
		&one.BusNum,
		&one.AmountTotal,
		&one.AmountPercent,
		&one.BusPercent,
		&one.AmountAvg,
		&one.IntRate,
		&one.FloatingLevel,
		&one.TermAvg,
		&one.AsOfDate,
		&one.Sort,
	}
	err := util.OracleScan(rows, values)
	return one, err
}

// 分页操作
// by author Jason
// by time 2016-11-15 16:10:56
func (this AnalysisSigDim) List(param ...map[string]interface{}) (*util.PageData, error) {
	pageData, rows, err := modelsUtil.List(analysisSigDimTabales, analysisSigDimCols, analysisSigDimColsSort, param...)
	defer rows.Close()
	if nil != err {
		er := fmt.Errorf("分页查询贷款业务分析方案单维度分析信息出错")
		zlog.Error(er.Error(), err)
		return nil, er
	}

	var rst []*AnalysisSigDim
	for rows.Next() {
		one, err := this.scan(rows)
		if nil != err {
			er := fmt.Errorf("分页查询贷款业务分析方案单维度分析信息row.Scan()出错")
			zlog.Error(er.Error(), err)
			return nil, er
		}
		rst = append(rst, one)
	}
	pageData.Rows = rst
	return pageData, nil
}

// 多参数查询
// by author Jason
// by time 2016-11-15 16:10:56
func (this AnalysisSigDim) Find(param ...map[string]interface{}) ([]*AnalysisSigDim, error) {
	rows, err := modelsUtil.FindRows(analysisSigDimTabales, analysisSigDimCols, analysisSigDimColsSort, param...)
	defer rows.Close()
	if nil != err {
		er := fmt.Errorf("多参数查询贷款业务分析方案单维度分析信息出错")
		zlog.Error(er.Error(), err)
		return nil, er
	}
	defer rows.Close()
	var rst []*AnalysisSigDim
	for rows.Next() {
		one, err := this.scan(rows)
		if nil != err {
			er := fmt.Errorf("多参数查询贷款业务分析方案单维度分析信息row.Scan()出错")
			zlog.Error(er.Error(), err)
			return nil, er
		}
		rst = append(rst, one)
	}
	return rst, nil
}

// 查询日期
// by author Jason
// by time 2016-11-15 16:10:56
func (this AnalysisSigDim) FindDistinctDate() ([]string, error) {
	sql := "select distinct as_of_date from rpm_bi_analysis_sig_dim T"
	rows, err := dbobj.Default.Query(sql)
	defer rows.Close()
	if nil != err {
		er := fmt.Errorf("查询贷款业务分析方案单维度分析日期信息出错")
		zlog.Error(er.Error(), err)
		return nil, er
	}
	var rst []string
	for rows.Next() {
		var one time.Time
		err := rows.Scan(&one)
		if nil != err {
			er := fmt.Errorf("查询贷款业务分析方案单维度分析日期信息row.Scan()出错")
			zlog.Error(er.Error(), err)
			return nil, er
		}
		rst = append(rst, one.Format("2006-01-02"))
	}
	return rst, nil
}

// 执行存储过程
// by author Jason
// by time 2016-11-15 16:10:56
func (this AnalysisSigDim) ExecProcedure(p_as_of_date time.Time) error {
	var retFlg, retMsg *string
	err := util.OracleExecProcedure("PROC_PRICING_ANALYSIS_BUS", p_as_of_date, retFlg, retMsg)
	if nil != err {
		er := fmt.Errorf("执行存储过程出错【PROC_PRICING_ANALYSIS_BUS】")
		zlog.Error(err.Error(), err)
		return er
	}
	return nil
}

// 多参数查询排序供前台出图表
// by author Jason
// by time 2016-11-15 16:10:56
func (this AnalysisSigDim) FindToSort(leftJoinTable, leftJoinCode string, param ...map[string]interface{}) ([]*AnalysisSigDim, error) {
	tables := "RPM_BI_ANALYSIS_SIG_DIM T" + " left join " + leftJoinTable + " on (dim_code = " + leftJoinCode + ")"
	closSort := append(analysisSigDimColsSort, "T1.sort")
	rows, err := modelsUtil.FindRows(tables, analysisSigDimColsSortMap, closSort, param...)
	defer rows.Close()
	if nil != err {
		er := fmt.Errorf("多参数查询贷款业务分析方案单维度分析信息出错")
		zlog.Error(er.Error(), err)
		return nil, er
	}
	var rst []*AnalysisSigDim
	for rows.Next() {
		one, err := this.scanToSort(rows)
		if nil != err {
			er := fmt.Errorf("多参数查询贷款业务分析方案单维度分析信息row.Scan()出错")
			zlog.Error(er.Error(), err)
			return nil, er
		}
		rst = append(rst, one)
	}
	return rst, nil
}

// 新增贷款业务分析方案单维度分析信息
// by author Jason
// by time 2016-11-15 16:10:56
func (this AnalysisSigDim) Add() error {
	paramMap := map[string]interface{}{
		"dim_code":       this.DimCode,
		"dim_name":       this.DimName,
		"dim_type":       this.DimType,
		"cust_num":       this.CustNum,
		"bus_num":        this.BusNum,
		"amount_total":   this.AmountTotal,
		"amount_percent": this.AmountPercent,
		"bus_percent":    this.BusPercent,
		"amount_avg":     this.AmountAvg,
		"int_rate":       this.IntRate,
		"floating_level": this.FloatingLevel,
		"term_avg":       this.TermAvg,
		"as_of_date":     this.AsOfDate,
	}
	err := util.OracleAdd("RPM_BI_ANALYSIS_SIG_DIM", paramMap)
	if nil != err {
		er := fmt.Errorf("新增贷款业务分析方案单维度分析信息出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

// 更新贷款业务分析方案单维度分析信息
// by author Jason
// by time 2016-11-15 16:10:56
func (this AnalysisSigDim) Update() error {
	paramMap := map[string]interface{}{
		"dim_code":       this.DimCode,
		"dim_name":       this.DimName,
		"dim_type":       this.DimType,
		"cust_num":       this.CustNum,
		"bus_num":        this.BusNum,
		"amount_total":   this.AmountTotal,
		"amount_percent": this.AmountPercent,
		"bus_percent":    this.BusPercent,
		"amount_avg":     this.AmountAvg,
		"int_rate":       this.IntRate,
		"floating_level": this.FloatingLevel,
		"term_avg":       this.TermAvg,
		"as_of_date":     this.AsOfDate,
	}
	whereParam := map[string]interface{}{
		"uuid": this.UUID,
	}
	err := util.OracleUpdate("RPM_BI_ANALYSIS_SIG_DIM", paramMap, whereParam)
	if nil != err {
		er := fmt.Errorf("更新贷款业务分析方案单维度分析信息出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

// 删除贷款业务分析方案单维度分析信息
// by author Jason
// by time 2016-11-15 16:10:56
func (this AnalysisSigDim) Delete() error {
	whereParam := map[string]interface{}{
		"uuid": this.UUID,
	}
	err := util.OracleDelete("RPM_BI_ANALYSIS_SIG_DIM", whereParam)
	if nil != err {
		er := fmt.Errorf("删除贷款业务分析方案单维度分析信息出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

var analysisSigDimTabales string = `RPM_BI_ANALYSIS_SIG_DIM T`

var analysisSigDimCols map[string]string = map[string]string{
	"UUID":           "' '",
	"DIM_CODE":       "' '",
	"DIM_NAME":       "' '",
	"DIM_TYPE":       "' '",
	"CUST_NUM":       "0",
	"BUS_NUM":        "0",
	"AMOUNT_TOTAL":   "0",
	"AMOUNT_PERCENT": "0",
	"BUS_PERCENT":    "0",
	"AMOUNT_AVG":     "0",
	"INT_RATE":       "0",
	"FLOATING_LEVEL": "0",
	"TERM_AVG":       "0",
	"AS_OF_DATE":     "sysdate",
}

var analysisSigDimColsSortMap map[string]string = map[string]string{
	"UUID":           "' '",
	"DIM_CODE":       "' '",
	"DIM_NAME":       "' '",
	"DIM_TYPE":       "' '",
	"CUST_NUM":       "0",
	"BUS_NUM":        "0",
	"AMOUNT_TOTAL":   "0",
	"AMOUNT_PERCENT": "0",
	"BUS_PERCENT":    "0",
	"AMOUNT_AVG":     "0",
	"INT_RATE":       "0",
	"FLOATING_LEVEL": "0",
	"TERM_AVG":       "0",
	"AS_OF_DATE":     "sysdate",
	"T1.sort":        "0",
}

var analysisSigDimColsSort = []string{
	"UUID",
	"DIM_CODE",
	"DIM_NAME",
	"DIM_TYPE",
	"CUST_NUM",
	"BUS_NUM",
	"AMOUNT_TOTAL",
	"AMOUNT_PERCENT",
	"BUS_PERCENT",
	"AMOUNT_AVG",
	"INT_RATE",
	"FLOATING_LEVEL",
	"TERM_AVG",
	"AS_OF_DATE",
}
