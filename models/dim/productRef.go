package dim

import (
	"database/sql"
	"fmt"

	modelsUtil "pccqcpa.com.cn/app/rpm/api/models/util"
	"pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/components/zlog"
)

// 结构体对应数据库表【RPM_DIM_PRODUCT_REF】产品映射关系表
// by author Jason
// by time 2017-03-30 13:51:39
type ProductRef struct {
	UUID        string  // 主键
	ProductCode Product `json:"product"` // 产品
	ProductName string  // 产品名字
	ProductRef  Product // 引用产品
	Mod         string  // 用途模块
	Describe    string  // 描述
	CreateTime  string  // 创建时间
	CreateUser  string  // 创建用户
	UpdateTime  string  // 更新时间
	UpdateUser  string  // 更新用户
	Flag        string  // 状态
}

// 结构体scan
// by author Jason
// by time 2017-03-30 13:51:39
func (this ProductRef) scan(rows *sql.Rows) (*ProductRef, error) {
	var one = new(ProductRef)
	values := []interface{}{
		&one.UUID,
		&one.ProductCode.ProductCode,
		&one.ProductName,
		&one.ProductRef.ProductCode,
		&one.Mod,
		&one.Describe,
		&one.CreateTime,
		&one.CreateUser,
		&one.UpdateTime,
		&one.UpdateUser,
		&one.Flag,

		&one.ProductCode.ProductName,
		&one.ProductRef.ProductName,
	}
	err := util.OracleScan(rows, values)
	return one, err
}

// 分页操作
// by author Jason
// by time 2017-03-30 13:51:39
func (this ProductRef) List(param ...map[string]interface{}) (*util.PageData, error) {
	pageData, rows, err := modelsUtil.List(productRefTabales, productRefCols, productRefColsSort, param...)
	defer rows.Close()
	if nil != err {
		er := fmt.Errorf("分页查询产品映射关系表信息出错")
		zlog.Error(er.Error(), err)
		return nil, er
	}

	var rst []*ProductRef
	for rows.Next() {
		one, err := this.scan(rows)
		if nil != err {
			er := fmt.Errorf("分页查询产品映射关系表信息row.Scan()出错")
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
// by time 2017-03-30 13:51:39
func (this ProductRef) Find(param ...map[string]interface{}) ([]*ProductRef, error) {
	rows, err := modelsUtil.FindRows(productRefTabales, productRefCols, productRefColsSort, param...)
	defer rows.Close()
	if nil != err {
		er := fmt.Errorf("多参数查询产品映射关系表信息出错")
		zlog.Error(er.Error(), err)
		return nil, er
	}
	var rst []*ProductRef
	for rows.Next() {
		one, err := this.scan(rows)
		if nil != err {
			er := fmt.Errorf("多参数查询产品映射关系表信息row.Scan()出错")
			zlog.Error(er.Error(), err)
			return nil, er
		}
		rst = append(rst, one)
	}
	return rst, nil
}

// 新增产品映射关系表信息
// by author Jason
// by time 2017-03-30 13:51:39
func (this ProductRef) Add() error {
	paramMap := map[string]interface{}{
		"product_code": this.ProductCode.ProductCode,
		"product_name": this.ProductName,
		"product_ref":  this.ProductRef.ProductCode,
		"mod":          this.Mod,
		"describe":     this.Describe,
		"create_time":  util.GetCurrentTime(),
		"create_user":  this.CreateUser,
		"update_time":  util.GetCurrentTime(),
		"update_user":  this.UpdateUser,
		"flag":         this.Flag,
	}
	err := util.OracleAdd("RPM_DIM_PRODUCT_REF", paramMap)
	if nil != err {
		er := fmt.Errorf("新增产品映射关系表信息出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

// 更新产品映射关系表信息
// by author Jason
// by time 2017-03-30 13:51:39
func (this ProductRef) Update() error {
	paramMap := map[string]interface{}{
		"product_code": this.ProductCode.ProductCode,
		"product_name": this.ProductName,
		"product_ref":  this.ProductRef.ProductCode,
		"mod":          this.Mod,
		"describe":     this.Describe,
		"create_user":  this.CreateUser,
		"update_time":  util.GetCurrentTime(),
		"update_user":  this.UpdateUser,
		"flag":         this.Flag,
	}
	whereParam := map[string]interface{}{
		"uuid": this.UUID,
	}
	err := util.OracleUpdate("RPM_DIM_PRODUCT_REF", paramMap, whereParam)
	if nil != err {
		er := fmt.Errorf("更新产品映射关系表信息出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

// 删除产品映射关系表信息
// by author Jason
// by time 2017-03-30 13:51:39
func (this ProductRef) Delete() error {
	whereParam := map[string]interface{}{
		"uuid": this.UUID,
	}
	err := util.OracleDelete("RPM_DIM_PRODUCT_REF", whereParam)
	if nil != err {
		er := fmt.Errorf("删除产品映射关系表信息出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

var productRefTabales string = `
		  RPM_DIM_PRODUCT_REF T
	 LEFT JOIN RPM_DIM_PRODUCT PRODUCT 
	   ON (T.PRODUCT_CODE = PRODUCT.PRODUCT_CODE)
	 LEFT JOIN RPM_DIM_PRODUCT PRODUCTREF
	   ON (T.PRODUCT_REF = PRODUCTREF.PRODUCT_CODE)`

var productRefCols map[string]string = map[string]string{
	"T.UUID":         "' '",
	"T.PRODUCT_CODE": "' '",
	"T.PRODUCT_NAME": "' '",
	"T.PRODUCT_REF":  "' '",
	"T.MOD":          "' '",
	"T.DESCRIBE":     "' '",
	"T.CREATE_TIME":  "' '",
	"T.CREATE_USER":  "' '",
	"T.UPDATE_TIME":  "' '",
	"T.UPDATE_USER":  "' '",
	"T.FLAG":         "' '",

	"PRODUCT.PRODUCT_NAME":    "' '",
	"PRODUCTREF.PRODUCT_NAME": "' '",
}

var productRefColsSort = []string{
	"T.UUID",
	"T.PRODUCT_CODE",
	"T.PRODUCT_NAME",
	"T.PRODUCT_REF",
	"T.MOD",
	"T.DESCRIBE",
	"T.CREATE_TIME",
	"T.CREATE_USER",
	"T.UPDATE_TIME",
	"T.UPDATE_USER",
	"T.FLAG",

	"PRODUCT.PRODUCT_NAME",
	"PRODUCTREF.PRODUCT_NAME",
}
