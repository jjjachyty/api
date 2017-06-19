package dim

import (
	"database/sql"
	"fmt"

	"platform/dbobj"

	modelsUtil "pccqcpa.com.cn/app/rpm/api/models/util"
	"pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/components/zlog"
)

// 结构体对应数据库表【RPM_BI_DIM_Amount_RANGE】贷款业务方案分析金额区间维度
// by author Jason
// by time 2016-10-31 16:46:57
type DimAmountRange struct {
	UUID      string  // 主键
	Range     string  // 金额区间
	AmountL   float64 // 金额下限
	AmountU   float64 // 金额上限
	RangeName string  // 金额翻译
	Sort      int     //排序字段
}

// 结构体scan
// by author Jason
// by time 2016-10-31 16:46:57
func (this DimAmountRange) scan(rows *sql.Rows) (*DimAmountRange, error) {
	var one = new(DimAmountRange)
	values := []interface{}{
		&one.UUID,
		&one.Range,
		&one.AmountL,
		&one.AmountU,
		&one.RangeName,
		&one.Sort,
	}
	err := util.OracleScan(rows, values)
	return one, err
}

// 分页操作
// by author Jason
// by time 2016-10-31 16:46:57
func (this DimAmountRange) List(param ...map[string]interface{}) (*util.PageData, error) {
	pageData, rows, err := modelsUtil.List(dimAmountRangeTabales, dimAmountRangeCols, dimAmountRangeColsSort, param...)
	defer rows.Close()
	if nil != err {
		er := fmt.Errorf("分页查询贷款业务方案分析金额区间维度信息出错")
		zlog.Error(er.Error(), err)
		return nil, er
	}

	var rst []*DimAmountRange
	for rows.Next() {
		one, err := this.scan(rows)
		if nil != err {
			er := fmt.Errorf("分页查询贷款业务方案分析金额区间维度信息row.Scan()出错")
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
// by time 2016-10-31 16:46:57
func (this DimAmountRange) Find(param ...map[string]interface{}) ([]*DimAmountRange, error) {
	rows, err := modelsUtil.FindRows(dimAmountRangeTabales, dimAmountRangeCols, dimAmountRangeColsSort, param...)
	defer rows.Close()
	if nil != err {
		er := fmt.Errorf("多参数查询贷款业务方案分析金额区间维度信息出错")
		zlog.Error(er.Error(), err)
		return nil, er
	}
	defer rows.Close()
	var rst []*DimAmountRange
	for rows.Next() {
		one, err := this.scan(rows)
		if nil != err {
			er := fmt.Errorf("多参数查询贷款业务方案分析金额区间维度信息row.Scan()出错")
			zlog.Error(er.Error(), err)
			return nil, er
		}
		rst = append(rst, one)
	}
	return rst, nil
}

// 新增贷款业务方案分析金额区间维度信息
// by author Jason
// by time 2016-10-31 16:46:57
func (this DimAmountRange) Add() error {
	paramMap := map[string]interface{}{
		"range":      this.Range,
		"amount_l":   this.AmountL,
		"amount_u":   this.AmountU,
		"range_name": this.RangeName,
		"sort":       this.Sort,
	}
	err := util.OracleAdd("RPM_BI_DIM_Amount_RANGE", paramMap)
	if nil != err {
		er := fmt.Errorf("新增贷款业务方案分析金额区间维度信息出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

// 新增贷款业务分析基准利率降息维度表信息
// by author Jason
// by time 2016-10-31 15:34:54
func (this DimAmountRange) BatchAdd(models []DimAmountRange) (sql.Result, error) {
	sqlTx, err := dbobj.Default.Begin()
	var tableName = "RPM_BI_DIM_AMOUNT_RANGE"
	var deleteSQL = "DELETE FROM " + tableName
	var msg = "[贷款业务方案分析金额区间维度信息-" + tableName + "]"
	if nil != err {
		zlog.Error(msg+"导入,事物开启失败", nil)
		return nil, err
	}
	//先清除表数据
	stmt, err := sqlTx.Prepare(deleteSQL)
	if nil != err {
		zlog.Error(msg+"导入,获取Stmt失败", nil)
		return nil, err
	}
	result, err := stmt.Exec()
	zlog.Infof(msg+"导入,DELETE-SQL:%s", nil, deleteSQL)
	if nil != err {
		zlog.Error(msg+"导入,DELETE出错", err)
		sqlTx.Rollback()
		return result, err
	}
	//关闭连接
	stmt.Close()

	for i, model := range models {
		zlog.Debugf(msg+"导入第%d条", nil, i+1)

		paramMap := map[string]interface{}{
			"range":      model.Range,
			"amount_l":   model.AmountL,
			"amount_u":   model.AmountU,
			"range_name": model.RangeName,
			"sort":       model.Sort,
		}
		result, err = util.OracleBatchStmtAdd(tableName, paramMap, sqlTx)
		if nil != err {
			zlog.Errorf(msg+"导入第%d行出错", err, i+1)
			err = stmt.Close()
			if nil != err {
				zlog.Error(msg+"链接关闭出错", nil)
				sqlTx.Rollback()
				return nil, err
			}
			sqlTx.Rollback()
			return result, fmt.Errorf(msg+"导入第%d行出错,%s", i+1, err)
		}
	}

	//关闭连接
	err = stmt.Close()
	if nil != err {
		zlog.Error(msg+"链接关闭出错", nil)

		sqlTx.Rollback()
		return nil, err
	}
	zlog.Info(msg+"导入成功", nil)
	sqlTx.Commit()

	return nil, nil

}

// 更新贷款业务方案分析金额区间维度信息
// by author Jason
// by time 2016-10-31 16:46:57
func (this DimAmountRange) Update() error {
	paramMap := map[string]interface{}{
		"range":      this.Range,
		"amount_l":   this.AmountL,
		"amount_u":   this.AmountU,
		"range_name": this.RangeName,
		"sort":       this.Sort,
	}
	whereParam := map[string]interface{}{
		"uuid": this.UUID,
	}
	err := util.OracleUpdate("RPM_BI_DIM_Amount_RANGE", paramMap, whereParam)
	if nil != err {
		er := fmt.Errorf("更新贷款业务方案分析金额区间维度信息出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

// 删除贷款业务方案分析金额区间维度信息
// by author Jason
// by time 2016-10-31 16:46:57
func (this DimAmountRange) Delete() error {
	whereParam := map[string]interface{}{
		"uuid": this.UUID,
	}
	err := util.OracleDelete("RPM_BI_DIM_Amount_RANGE", whereParam)
	if nil != err {
		er := fmt.Errorf("删除贷款业务方案分析金额区间维度信息出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

var dimAmountRangeTabales string = `RPM_BI_DIM_Amount_RANGE`

var dimAmountRangeCols map[string]string = map[string]string{
	"UUID":       "' '",
	"RANGE":      "' '",
	"Amount_L":   "0",
	"Amount_U":   "0",
	"RANGE_NAME": "' '",
	"SORT":       "0",
}

var dimAmountRangeColsSort = []string{
	"UUID",
	"RANGE",
	"AMOUNT_L",
	"AMOUNT_U",
	"RANGE_NAME",
	"SORT",
}
