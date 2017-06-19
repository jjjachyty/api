package par

import (
	"database/sql"
	"fmt"

	"pccqcpa.com.cn/app/rpm/api/models/dim"
	modelsUtil "pccqcpa.com.cn/app/rpm/api/models/util"
	"pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/components/zlog"
)

// 结构体对应数据库表【RPM_PAR_RL_PD】零售违约概率-零售贷款模块
// by author Yeqc
// by time 2016-12-30 11:03:34
/**
 * @apiDefine RlPd
 * @apiSuccess {string}     UUID            主键默认值sys_guid()
 * @apiSuccess {Product}   	ProductCode    	产品
 * @apiSuccess {float64}   	TotalAcount     总账户
 * @apiSuccess {float64}   	BreachAcount   	违约账户
 * @apiSuccess {float64}   	Pd    			违约概率 百分比
 */
type RlPd struct {
	UUID         string      // 默认值sys_guid()
	ProductCode  dim.Product // 产品
	TotalAcount  float64     // 总账户
	BreachAcount float64     // 违约账户
	Pd           float64     // 违约概率 百分比
}

// 结构体scan
// by author Yeqc
// by time 2016-12-30 11:03:34
func (this RlPd) scan(rows *sql.Rows) (*RlPd, error) {
	var one = new(RlPd)
	values := []interface{}{
		&one.UUID,
		// &one.ProductCode,
		&one.TotalAcount,
		&one.BreachAcount,
		&one.Pd,

		&one.ProductCode.ProductCode,
		&one.ProductCode.ProductName,
		&one.ProductCode.ProductType,
		&one.ProductCode.ProductTypeDesc,
		&one.ProductCode.ProductLevel,
	}
	err := util.OracleScan(rows, values)
	return one, err
}

// 分页操作
// by author Yeqc
// by time 2016-12-30 11:03:34
func (this RlPd) List(param ...map[string]interface{}) (*util.PageData, error) {
	pageData, rows, err := modelsUtil.List(rlPdTabales, rlPdCols, rlPdColsSort, param...)
	defer rows.Close()
	if nil != err {
		er := fmt.Errorf("分页查询零售违约概率-零售贷款模块信息出错")
		zlog.Error(er.Error(), err)
		return nil, er
	}

	var rst []*RlPd
	for rows.Next() {
		one, err := this.scan(rows)
		if nil != err {
			er := fmt.Errorf("分页查询零售违约概率-零售贷款模块信息row.Scan()出错")
			zlog.Error(er.Error(), err)
			return nil, er
		}
		rst = append(rst, one)
	}
	pageData.Rows = rst
	return pageData, nil
}

// 多参数查询
// by author Yeqc
// by time 2016-12-30 11:03:34
func (this RlPd) Find(param ...map[string]interface{}) ([]*RlPd, error) {
	rows, err := modelsUtil.FindRows(rlPdTabales, rlPdCols, rlPdColsSort, param...)
	defer rows.Close()
	if nil != err {
		er := fmt.Errorf("多参数查询零售违约概率-零售贷款模块信息出错")
		zlog.Error(er.Error(), err)
		return nil, er
	}
	var rst []*RlPd
	for rows.Next() {
		one, err := this.scan(rows)
		if nil != err {
			er := fmt.Errorf("多参数查询零售违约概率-零售贷款模块信息row.Scan()出错")
			zlog.Error(er.Error(), err)
			return nil, er
		}
		rst = append(rst, one)
	}
	return rst, nil
}

// 新增零售违约概率-零售贷款模块信息
// by author Yeqc
// by time 2016-12-30 11:03:34
func (this RlPd) Add() error {
	paramMap := map[string]interface{}{
		"product_code":  this.ProductCode.ProductCode,
		"total_acount":  this.TotalAcount,
		"breach_acount": this.BreachAcount,
		"pd":            this.Pd,
	}
	err := util.OracleAdd("RPM_PAR_RL_PD", paramMap)
	if nil != err {
		er := fmt.Errorf("新增零售违约概率-零售贷款模块信息出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

// 更新零售违约概率-零售贷款模块信息
// by author Yeqc
// by time 2016-12-30 11:03:34
func (this RlPd) Update() error {
	paramMap := map[string]interface{}{
		"product_code":  this.ProductCode.ProductCode,
		"total_acount":  this.TotalAcount,
		"breach_acount": this.BreachAcount,
		"pd":            this.Pd,
	}
	whereParam := map[string]interface{}{
		"uuid": this.UUID,
	}
	err := util.OracleUpdate("RPM_PAR_RL_PD", paramMap, whereParam)
	if nil != err {
		er := fmt.Errorf("更新零售违约概率-零售贷款模块信息出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

// 删除零售违约概率-零售贷款模块信息
// by author Yeqc
// by time 2016-12-30 11:03:34
func (this RlPd) Delete() error {
	whereParam := map[string]interface{}{
		"uuid": this.UUID,
	}
	err := util.OracleDelete("RPM_PAR_RL_PD", whereParam)
	if nil != err {
		er := fmt.Errorf("删除零售违约概率-零售贷款模块信息出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

var rlPdTabales string = `
	RPM_PAR_RL_PD T
	LEFT JOIN RPM_DIM_PRODUCT PRODUCTCODE
	ON (PRODUCTCODE.PRODUCT_CODE=T.PRODUCT_CODE)
	`

var rlPdCols map[string]string = map[string]string{
	"T.UUID": "' '",
	// "T.PRODUCT_CODE":  "' '",
	"T.TOTAL_ACOUNT":  "0",
	"T.BREACH_ACOUNT": "0",
	"T.PD":            "0",

	"PRODUCTCODE.product_code":      "' '",
	"PRODUCTCODE.product_name":      "' '",
	"PRODUCTCODE.product_type":      "' '",
	"PRODUCTCODE.product_type_desc": "' '",
	"PRODUCTCODE.product_level":     "' '",
}

var rlPdColsSort = []string{
	"T.UUID",
	// "T.PRODUCT_CODE",
	"T.TOTAL_ACOUNT",
	"T.BREACH_ACOUNT",
	"T.PD",

	"PRODUCTCODE.product_code",
	"PRODUCTCODE.product_name",
	"PRODUCTCODE.product_type",
	"PRODUCTCODE.product_type_desc",
	"PRODUCTCODE.product_level",
}
