package dim

import (
	"database/sql"
	"fmt"
	"time"

	modelsUtil "pccqcpa.com.cn/app/rpm/api/models/util"
	"pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/components/zlog"

	"platform/dbobj"
)

// 结构体对应数据库表【RPM_BI_DIM_BASE_RATE_REDUCTION】贷款业务分析基准利率降息维度表
// by author Jason
// by time 2016-10-31 15:34:54
type DimBaseRateReduction struct {
	UUID         string    // 主键
	TermRange    string    // 期限区间
	ReductionDay time.Time // 降息日
	BaseRate     float64   // 基准利率
}

// 结构体scan
// by author Jason
// by time 2016-10-31 15:34:54
func (this DimBaseRateReduction) scan(rows *sql.Rows) (*DimBaseRateReduction, error) {
	var one = new(DimBaseRateReduction)
	values := []interface{}{
		&one.UUID,
		&one.TermRange,
		&one.ReductionDay,
		&one.BaseRate,
	}
	err := util.OracleScan(rows, values)
	return one, err
}

// 分页操作
// by author Jason
// by time 2016-10-31 15:34:54
func (this DimBaseRateReduction) List(param ...map[string]interface{}) (*util.PageData, error) {
	pageData, rows, err := modelsUtil.List(dimBaseRateReductionTabales, dimBaseRateReductionCols, dimBaseRateReductionColsSort, param...)
	defer rows.Close()
	if nil != err {
		er := fmt.Errorf("分页查询贷款业务分析基准利率降息维度表信息出错")
		zlog.Error(er.Error(), err)
		return nil, er
	}

	var rst []*DimBaseRateReduction
	for rows.Next() {
		one, err := this.scan(rows)
		if nil != err {
			er := fmt.Errorf("分页查询贷款业务分析基准利率降息维度表信息row.Scan()出错")
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
// by time 2016-10-31 15:34:54
func (this DimBaseRateReduction) Find(param ...map[string]interface{}) ([]*DimBaseRateReduction, error) {
	rows, err := modelsUtil.FindRows(dimBaseRateReductionTabales, dimBaseRateReductionCols, dimBaseRateReductionColsSort, param...)
	defer rows.Close()
	if nil != err {
		er := fmt.Errorf("多参数查询贷款业务分析基准利率降息维度表信息出错")
		zlog.Error(er.Error(), err)
		return nil, er
	}
	defer rows.Close()
	var rst []*DimBaseRateReduction
	for rows.Next() {
		one, err := this.scan(rows)
		if nil != err {
			er := fmt.Errorf("多参数查询贷款业务分析基准利率降息维度表信息row.Scan()出错")
			zlog.Error(er.Error(), err)
			return nil, er
		}
		rst = append(rst, one)
	}
	return rst, nil
}

// 新增贷款业务分析基准利率降息维度表信息
// by author Jason
// by time 2016-10-31 15:34:54
func (this DimBaseRateReduction) Add() error {
	paramMap := map[string]interface{}{
		"Term_range":    this.TermRange,
		"reduction_day": this.ReductionDay,
		"base_rate":     this.BaseRate,
	}
	err := util.OracleAdd("RPM_BI_DIM_BASE_RATE_REDUCTION", paramMap)
	if nil != err {
		er := fmt.Errorf("新增贷款业务分析基准利率降息维度表信息出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

// 新增贷款业务分析基准利率降息维度表信息
// by author Jason
// by time 2016-10-31 15:34:54
func (this DimBaseRateReduction) BatchAdd(models []DimBaseRateReduction) (sql.Result, error) {
	sqlTx, err := dbobj.Default.Begin()
	if nil != err {
		zlog.Error("[贷款业务分析基准利率降息维度表-RPM_BI_DIM_BASE_RATE_REDUCTION]导入,事物开启失败", nil)
		return nil, err
	}
	//先清除表数据
	stmt, err := sqlTx.Prepare(deleteSQL)
	if nil != err {
		zlog.Error("[贷款业务分析基准利率降息维度表-RPM_BI_DIM_BASE_RATE_REDUCTION]导入,获取Stmt失败", nil)
		return nil, err
	}
	result, err := stmt.Exec()
	zlog.Infof("[贷款业务分析基准利率降息维度表-RPM_BI_DIM_BASE_RATE_REDUCTION]导入,TRUNCATE-SQL:%s", nil, deleteSQL)
	if nil != err {
		zlog.Error("[贷款业务分析基准利率降息维度表-RPM_BI_DIM_BASE_RATE_REDUCTION]导入,TRUNCATE出错", err)
		sqlTx.Rollback()
		return result, err
	}
	//关闭连接
	stmt.Close()

	for i, model := range models {
		zlog.Debugf("[贷款业务分析基准利率降息维度表-RPM_BI_DIM_BASE_RATE_REDUCTION]导入第%d条", nil, i+1)

		paramMap := map[string]interface{}{
			"Term_range":    model.TermRange,
			"reduction_day": model.ReductionDay,
			"base_rate":     model.BaseRate,
		}
		result, err = util.OracleBatchStmtAdd("RPM_BI_DIM_BASE_RATE_REDUCTION", paramMap, sqlTx)
		if nil != err {
			zlog.Errorf("[贷款业务分析基准利率降息维度表-RPM_BI_DIM_BASE_RATE_REDUCTION]导入第%d行出错", err, i+1)
			sqlTx.Rollback()
			return result, fmt.Errorf("[贷款业务分析基准利率降息维度表-RPM_BI_DIM_BASE_RATE_REDUCTION]导入第%d行出错,%s", i+1, err)
		}
	}

	//关闭连接
	err = stmt.Close()
	if nil != err {
		zlog.Error("[贷款业务分析基准利率降息维度表-RPM_BI_DIM_BASE_RATE_REDUCTION]链接关闭出错", nil)
		sqlTx.Rollback()
		return nil, err
	}
	zlog.Info("[贷款业务分析基准利率降息维度表-RPM_BI_DIM_BASE_RATE_REDUCTION]导入成功", nil)
	sqlTx.Commit()

	return nil, nil

}

// 更新贷款业务分析基准利率降息维度表信息
// by author Jason
// by time 2016-10-31 15:34:54
func (this DimBaseRateReduction) Update() error {
	paramMap := map[string]interface{}{
		"Term_range":    this.TermRange,
		"reduction_day": this.ReductionDay,
		"base_rate":     this.BaseRate,
	}
	whereParam := map[string]interface{}{
		"uuid": this.UUID,
	}
	err := util.OracleUpdate("RPM_BI_DIM_BASE_RATE_REDUCTION", paramMap, whereParam)
	if nil != err {
		er := fmt.Errorf("更新贷款业务分析基准利率降息维度表信息出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

// 删除贷款业务分析基准利率降息维度表信息
// by author Jason
// by time 2016-10-31 15:34:54
func (this DimBaseRateReduction) Delete() error {
	whereParam := map[string]interface{}{
		"uuid": this.UUID,
	}
	err := util.OracleDelete("RPM_BI_DIM_BASE_RATE_REDUCTION", whereParam)
	if nil != err {
		er := fmt.Errorf("删除贷款业务分析基准利率降息维度表信息出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

var dimBaseRateReductionTabales string = `RPM_BI_DIM_BASE_RATE_REDUCTION`

var dimBaseRateReductionCols map[string]string = map[string]string{
	"UUID":          "' '",
	"TERM_RANGE":    "' '",
	"REDUCTION_DAY": "sysdate",
	"BASE_RATE":     "0",
}

var dimBaseRateReductionColsSort = []string{
	"UUID",
	"TERM_RANGE",
	"REDUCTION_DAY",
	"BASE_RATE",
}

var deleteSQL = "DELETE FROM RPM_BI_DIM_BASE_RATE_REDUCTION"
