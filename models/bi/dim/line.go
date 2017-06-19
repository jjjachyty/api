package dim

import (
	"database/sql"
	"fmt"

	modelsUtil "pccqcpa.com.cn/app/rpm/api/models/util"
	"pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/components/zlog"
	"platform/dbobj"
)

// 结构体对应数据库表【RPM_BI_DIM_LINE】贷款业务方案分析业务条线表
// by author Jason
// by time 2016-10-31 15:33:31
type DimLine struct {
	UUID        string // 主键
	LineCode    string // 业务条线编吗
	LineName    string // 业务条线名称
	ProductType string // 产品类型
}

// 结构体scan
// by author Jason
// by time 2016-10-31 15:33:31
func (this DimLine) scan(rows *sql.Rows) (*DimLine, error) {
	var one = new(DimLine)
	values := []interface{}{
		&one.UUID,
		&one.LineCode,
		&one.LineName,
		&one.ProductType,
	}
	err := util.OracleScan(rows, values)
	return one, err
}

// 分页操作
// by author Jason
// by time 2016-10-31 15:33:31
func (this DimLine) List(param ...map[string]interface{}) (*util.PageData, error) {
	pageData, rows, err := modelsUtil.List(dimLineTabales, dimLineCols, dimLineColsSort, param...)
	defer rows.Close()
	if nil != err {
		er := fmt.Errorf("分页查询贷款业务方案分析业务条线表信息出错")
		zlog.Error(er.Error(), err)
		return nil, er
	}

	var rst []*DimLine
	for rows.Next() {
		one, err := this.scan(rows)
		if nil != err {
			er := fmt.Errorf("分页查询贷款业务方案分析业务条线表信息row.Scan()出错")
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
// by time 2016-10-31 15:33:31
func (this DimLine) Find(param ...map[string]interface{}) ([]*DimLine, error) {
	rows, err := modelsUtil.FindRows(dimLineTabales, dimLineCols, dimLineColsSort, param...)
	defer rows.Close()
	if nil != err {
		er := fmt.Errorf("多参数查询贷款业务方案分析业务条线表信息出错")
		zlog.Error(er.Error(), err)
		return nil, er
	}
	defer rows.Close()
	var rst []*DimLine
	for rows.Next() {
		one, err := this.scan(rows)
		if nil != err {
			er := fmt.Errorf("多参数查询贷款业务方案分析业务条线表信息row.Scan()出错")
			zlog.Error(er.Error(), err)
			return nil, er
		}
		rst = append(rst, one)
	}
	return rst, nil
}

// 新增贷款业务方案分析业务条线表信息
// by author Jason
// by time 2016-10-31 15:33:31
func (this DimLine) Add() error {
	paramMap := map[string]interface{}{
		"line_code":    this.LineCode,
		"line_name":    this.LineName,
		"product_type": this.ProductType,
	}
	err := util.OracleAdd("RPM_BI_DIM_LINE", paramMap)
	if nil != err {
		er := fmt.Errorf("新增贷款业务方案分析业务条线表信息出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

// 新增贷款业务分析基准利率降息维度表信息
// by author Jason
// by time 2016-10-31 15:34:54
func (this DimLine) BatchAdd(models []DimLine) (sql.Result, error) {
	sqlTx, err := dbobj.Default.Begin()
	var tableName = "RPM_BI_DIM_LINE"
	var deleteSQL = "DELETE FROM " + tableName
	var msg = "[贷款业务方案分析业务条线表信息-" + tableName + "]"
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
			"line_code":    model.LineCode,
			"line_name":    model.LineName,
			"product_type": model.ProductType,
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

// 更新贷款业务方案分析业务条线表信息
// by author Jason
// by time 2016-10-31 15:33:31
func (this DimLine) Update() error {
	paramMap := map[string]interface{}{
		"line_code":    this.LineCode,
		"line_name":    this.LineName,
		"product_type": this.ProductType,
	}
	whereParam := map[string]interface{}{
		"uuid": this.UUID,
	}
	err := util.OracleUpdate("RPM_BI_DIM_LINE", paramMap, whereParam)
	if nil != err {
		er := fmt.Errorf("更新贷款业务方案分析业务条线表信息出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

// 删除贷款业务方案分析业务条线表信息
// by author Jason
// by time 2016-10-31 15:33:31
func (this DimLine) Delete() error {
	whereParam := map[string]interface{}{
		"uuid": this.UUID,
	}
	err := util.OracleDelete("RPM_BI_DIM_LINE", whereParam)
	if nil != err {
		er := fmt.Errorf("删除贷款业务方案分析业务条线表信息出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

var dimLineTabales string = `RPM_BI_DIM_LINE`

var dimLineCols map[string]string = map[string]string{
	"UUID":         "' '",
	"LINE_CODE":    "' '",
	"LINE_NAME":    "' '",
	"PRODUCT_TYPE": "' '",
}

var dimLineColsSort = []string{
	"UUID",
	"LINE_CODE",
	"LINE_NAME",
	"PRODUCT_TYPE",
}
