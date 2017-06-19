package dim

import (
	"database/sql"
	"fmt"

	modelsUtil "pccqcpa.com.cn/app/rpm/api/models/util"
	"pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/components/zlog"
	"platform/dbobj"
)

// 结构体对应数据库表【RPM_BI_DIM_PRODUCT】贷款方案分析产品维度表
// by author Jason
// by time 2016-10-31 15:38:16
type DimProduct struct {
	UUID              string // 主键
	ProductCodeAnaly  string // 定价分析产品编码  PRODUCT_CODE_ANALY
	ProductNameAnaly  string // 定价分析产品名称  PRODUCT_NAME_ANALY
	ProductCodeCredit string // 信贷产品编码     PRODUCT_CODE_CREDIT
	ProductNameCredit string // 信贷产品名称     PRODUCT_NAME_CREDIT
}

// 结构体scan
// by author Jason
// by time 2016-10-31 15:38:16
func (this DimProduct) scan(rows *sql.Rows) (*DimProduct, error) {
	var one = new(DimProduct)
	values := []interface{}{
		&one.UUID,
		&one.ProductCodeAnaly,
		&one.ProductNameAnaly,
		&one.ProductCodeCredit,
		&one.ProductNameCredit,
	}
	err := util.OracleScan(rows, values)
	return one, err
}

// 分页操作
// by author Jason
// by time 2016-10-31 15:38:16
func (this DimProduct) List(param ...map[string]interface{}) (*util.PageData, error) {
	pageData, rows, err := modelsUtil.List(dimProductTabales, dimProductCols, dimProductColsSort, param...)
	defer rows.Close()
	if nil != err {
		er := fmt.Errorf("分页查询贷款方案分析产品维度表信息出错")
		zlog.Error(er.Error(), err)
		return nil, er
	}

	var rst []*DimProduct
	for rows.Next() {
		one, err := this.scan(rows)
		if nil != err {
			er := fmt.Errorf("分页查询贷款方案分析产品维度表信息row.Scan()出错")
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
// by time 2016-10-31 15:38:16
func (this DimProduct) Find(param ...map[string]interface{}) ([]*DimProduct, error) {
	rows, err := modelsUtil.FindRows(dimProductTabales, dimProductCols, dimProductColsSort, param...)
	defer rows.Close()
	if nil != err {
		er := fmt.Errorf("多参数查询贷款方案分析产品维度表信息出错")
		zlog.Error(er.Error(), err)
		return nil, er
	}
	defer rows.Close()
	var rst []*DimProduct
	for rows.Next() {
		one, err := this.scan(rows)
		if nil != err {
			er := fmt.Errorf("多参数查询贷款方案分析产品维度表信息row.Scan()出错")
			zlog.Error(er.Error(), err)
			return nil, er
		}
		rst = append(rst, one)
	}
	return rst, nil
}

// 新增贷款方案分析产品维度表信息
// by author Jason
// by time 2016-10-31 15:38:16
func (this DimProduct) Add() error {
	paramMap := map[string]interface{}{
		"product_code_analy":  this.ProductCodeAnaly,
		"product_name_analy":  this.ProductNameAnaly,
		"product_code_credit": this.ProductCodeCredit,
		"product_name_credit": this.ProductNameCredit,
	}
	err := util.OracleAdd("RPM_BI_DIM_PRODUCT", paramMap)
	if nil != err {
		er := fmt.Errorf("新增贷款方案分析产品维度表信息出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

// 新增贷款业务分析基准利率降息维度表信息
// by author Jason
// by time 2016-10-31 15:34:54
func (this DimProduct) BatchAdd(models []DimProduct) (sql.Result, error) {
	sqlTx, err := dbobj.Default.Begin()
	var tableName = "RPM_BI_DIM_PRODUCT"
	var deleteSQL = "DELETE FROM " + tableName
	var msg = "[贷款方案分析产品维度表信息-" + tableName + "]"
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
			"product_code_analy":  model.ProductCodeAnaly,
			"product_name_analy":  model.ProductNameAnaly,
			"product_code_credit": model.ProductCodeCredit,
			"product_name_credit": model.ProductNameCredit,
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

// 更新贷款方案分析产品维度表信息
// by author Jason
// by time 2016-10-31 15:38:16
func (this DimProduct) Update() error {
	paramMap := map[string]interface{}{
		"product_code_analy":  this.ProductCodeAnaly,
		"product_name_analy":  this.ProductNameAnaly,
		"product_code_credit": this.ProductCodeCredit,
		"product_name_credit": this.ProductNameCredit,
	}
	whereParam := map[string]interface{}{
		"uuid": this.UUID,
	}
	err := util.OracleUpdate("RPM_BI_DIM_PRODUCT", paramMap, whereParam)
	if nil != err {
		er := fmt.Errorf("更新贷款方案分析产品维度表信息出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

// 删除贷款方案分析产品维度表信息
// by author Jason
// by time 2016-10-31 15:38:16
func (this DimProduct) Delete() error {
	whereParam := map[string]interface{}{
		"uuid": this.UUID,
	}
	err := util.OracleDelete("RPM_BI_DIM_PRODUCT", whereParam)
	if nil != err {
		er := fmt.Errorf("删除贷款方案分析产品维度表信息出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

var dimProductTabales string = `RPM_BI_DIM_PRODUCT T`

var dimProductCols map[string]string = map[string]string{
	"UUID":                "' '",
	"PRODUCT_CODE_ANALY":  "' '",
	"PRODUCT_NAME_ANALY":  "' '",
	"PRODUCT_CODE_CREDIT": "' '",
	"PRODUCT_NAME_CREDIT": "' '",
}

var dimProductColsSort = []string{
	"UUID",
	"PRODUCT_CODE_ANALY",
	"PRODUCT_NAME_ANALY",
	"PRODUCT_CODE_CREDIT",
	"PRODUCT_NAME_CREDIT",
}
