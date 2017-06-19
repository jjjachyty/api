package dim

import (
	"database/sql"
	"fmt"

	"platform/dbobj"

	modelsUtil "pccqcpa.com.cn/app/rpm/api/models/util"
	"pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/components/zlog"
)

// 结构体对应数据库表【RPM_BI_DIM_Term_RANGE】贷款业务方案分析期限区间维度
// by author Jason
// by time 2016-10-31 16:46:57
type DimTermRange struct {
	UUID     string // 主键
	TermL    string // 期限下限
	TermU    string // 期限上限
	TermCode string // 期限含义
	TermName string // 期限翻译
	Sort     int
}

// 结构体scan
// by author Jason
// by time 2016-10-31 16:46:57
func (this DimTermRange) scan(rows *sql.Rows) (*DimTermRange, error) {
	var one = new(DimTermRange)
	values := []interface{}{
		&one.UUID,
		&one.TermL,
		&one.TermU,
		&one.TermCode,
		&one.TermName,
		&one.Sort,
	}
	err := util.OracleScan(rows, values)
	return one, err
}

// 分页操作
// by author Jason
// by time 2016-10-31 16:46:57
func (this DimTermRange) List(param ...map[string]interface{}) (*util.PageData, error) {
	pageData, rows, err := modelsUtil.List(dimTermRangeTabales, dimTermRangeCols, dimTermRangeColsSort, param...)
	defer rows.Close()
	if nil != err {
		er := fmt.Errorf("分页查询贷款业务方案分析期限区间维度信息出错")
		zlog.Error(er.Error(), err)
		return nil, er
	}

	var rst []*DimTermRange
	for rows.Next() {
		one, err := this.scan(rows)
		if nil != err {
			er := fmt.Errorf("分页查询贷款业务方案分析期限区间维度信息row.Scan()出错")
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
func (this DimTermRange) Find(param ...map[string]interface{}) ([]*DimTermRange, error) {
	rows, err := modelsUtil.FindRows(dimTermRangeTabales, dimTermRangeCols, dimTermRangeColsSort, param...)
	defer rows.Close()
	if nil != err {
		er := fmt.Errorf("多参数查询贷款业务方案分析期限区间维度信息出错")
		zlog.Error(er.Error(), err)
		return nil, er
	}
	defer rows.Close()
	var rst []*DimTermRange
	for rows.Next() {
		one, err := this.scan(rows)
		if nil != err {
			er := fmt.Errorf("多参数查询贷款业务方案分析期限区间维度信息row.Scan()出错")
			zlog.Error(er.Error(), err)
			return nil, er
		}
		rst = append(rst, one)
	}
	return rst, nil
}

// 新增贷款业务方案分析期限区间维度信息
// by author Jason
// by time 2016-10-31 16:46:57
func (this DimTermRange) Add() error {
	paramMap := map[string]interface{}{
		"Term_l":    this.TermL,
		"Term_u":    this.TermU,
		"Term_code": this.TermCode,
		"Term_name": this.TermName,
		"Sort":      this.Sort,
	}
	err := util.OracleAdd("RPM_BI_DIM_Term_RANGE", paramMap)
	if nil != err {
		er := fmt.Errorf("新增贷款业务方案分析期限区间维度信息出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

// 新增贷款业务分析基准利率降息维度表信息
// by author Jason
// by time 2016-10-31 15:34:54
func (this DimTermRange) BatchAdd(models []DimTermRange) (sql.Result, error) {
	sqlTx, err := dbobj.Default.Begin()
	var tableName = "RPM_BI_DIM_Term_RANGE"
	var deleteSQL = "DELETE FROM " + tableName
	var msg = "[贷款业务方案分析期限区间维度信息-" + tableName + "]"
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
			"Term_l":    model.TermL,
			"Term_u":    model.TermU,
			"Term_code": model.TermCode,
			"Term_name": model.TermName,
			"Sort":      model.Sort,
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

// 更新贷款业务方案分析期限区间维度信息
// by author Jason
// by time 2016-10-31 16:46:57
func (this DimTermRange) Update() error {
	paramMap := map[string]interface{}{
		"Term_l":    this.TermL,
		"Term_u":    this.TermU,
		"Term_code": this.TermCode,
		"Term_name": this.TermName,
		"Sort":      this.Sort,
	}
	whereParam := map[string]interface{}{
		"uuid": this.UUID,
	}
	err := util.OracleUpdate("RPM_BI_DIM_Term_RANGE", paramMap, whereParam)
	if nil != err {
		er := fmt.Errorf("更新贷款业务方案分析期限区间维度信息出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

// 删除贷款业务方案分析期限区间维度信息
// by author Jason
// by time 2016-10-31 16:46:57
func (this DimTermRange) Delete() error {
	whereParam := map[string]interface{}{
		"uuid": this.UUID,
	}
	err := util.OracleDelete("RPM_BI_DIM_Term_RANGE", whereParam)
	if nil != err {
		er := fmt.Errorf("删除贷款业务方案分析期限区间维度信息出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

var dimTermRangeTabales string = `RPM_BI_DIM_Term_RANGE`

var dimTermRangeCols map[string]string = map[string]string{
	"UUID":      "' '",
	"Term_L":    "' '",
	"Term_U":    "' '",
	"Term_CODE": "' '",
	"Term_NAME": "' '",
	"Sort":      "0",
}

var dimTermRangeColsSort = []string{
	"UUID",
	"Term_L",
	"Term_U",
	"Term_CODE",
	"Term_NAME",
	"Sort",
}
