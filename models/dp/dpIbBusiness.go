package dp

import (
	"database/sql"
	"fmt"
	"time"

	"pccqcpa.com.cn/app/rpm/api/models/dim"
	modelsUtil "pccqcpa.com.cn/app/rpm/api/models/util"
	"pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/components/zlog"
)

// 结构体对应数据库表【RPM_BIZ_DP_IB_BUSINESS】一对一存款中间收入表
// by author Jason
// by time 2016-12-06 16:59:09
type DpIbBusiness struct {
	UUID         string      // 主键 默认值sys_guid()
	BusinessCode string      // 业务单号
	Product      dim.Product // 产品
	Amount       float64     // 金额
	Status       string      // 状态
	CreateTime   time.Time   // 创建时间
	UpdateTime   time.Time   // 更新时间
	CreateUser   string      // 创建人
	UpdateUser   string      // 更新人
}

// 结构体scan
// by author Jason
// by time 2016-12-06 16:59:09
func (this DpIbBusiness) scan(rows *sql.Rows) (*DpIbBusiness, error) {
	var one = new(DpIbBusiness)
	values := []interface{}{
		&one.UUID,
		&one.BusinessCode,
		&one.Product.ProductCode,
		&one.Amount,
		&one.Status,
		&one.CreateTime,
		&one.CreateUser,
		&one.UpdateTime,
		&one.UpdateUser,
		&one.Product.ProductName,
	}
	err := util.OracleScan(rows, values)
	return one, err
}

// 分页操作
// by author Jason
// by time 2016-12-06 16:59:09
func (this DpIbBusiness) List(param ...map[string]interface{}) (*util.PageData, error) {
	pageData, rows, err := modelsUtil.List(dpIbBusinessTabales, dpIbBusinessCols, dpIbBusinessColsSort, param...)
	defer rows.Close()
	if nil != err {
		er := fmt.Errorf("分页查询一对一存款中间收入表信息出错")
		zlog.Error(er.Error(), err)
		return nil, er
	}

	var rst []*DpIbBusiness
	for rows.Next() {
		one, err := this.scan(rows)
		if nil != err {
			er := fmt.Errorf("分页查询一对一存款中间收入表信息row.Scan()出错")
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
// by time 2016-12-06 16:59:09
func (this DpIbBusiness) Find(param ...map[string]interface{}) ([]*DpIbBusiness, error) {
	rows, err := modelsUtil.FindRows(dpIbBusinessTabales, dpIbBusinessCols, dpIbBusinessColsSort, param...)
	defer rows.Close()
	if nil != err {
		er := fmt.Errorf("多参数查询一对一存款中间收入表信息出错")
		zlog.Error(er.Error(), err)
		return nil, er
	}
	var rst []*DpIbBusiness
	for rows.Next() {
		one, err := this.scan(rows)
		if nil != err {
			er := fmt.Errorf("多参数查询一对一存款中间收入表信息row.Scan()出错")
			zlog.Error(er.Error(), err)
			return nil, er
		}
		rst = append(rst, one)
	}
	return rst, nil
}

// 新增一对一存款中间收入表信息
// by author Jason
// by time 2016-12-06 16:59:09
func (this DpIbBusiness) Add() error {
	paramMap := map[string]interface{}{
		"business_code": this.BusinessCode,
		"product":       this.Product.ProductCode,
		"amount":        this.Amount,
		"status":        this.Status,
		"create_time":   util.GetCurrentTime(),
		"create_user":   this.CreateUser,
		"update_time":   util.GetCurrentTime(),
		"update_user":   this.UpdateUser,
	}
	err := util.OracleAdd("RPM_BIZ_DP_IB_BUSINESS", paramMap)
	if nil != err {
		er := fmt.Errorf("新增一对一存款中间收入表信息出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

// 更新一对一存款中间收入表信息
// by author Jason
// by time 2016-12-06 16:59:09
func (this DpIbBusiness) Update() error {
	paramMap := map[string]interface{}{
		"product":     this.Product.ProductCode,
		"amount":      this.Amount,
		"status":      this.Status,
		"update_time": util.GetCurrentTime(),
		"update_user": this.UpdateUser,
	}
	whereParam := map[string]interface{}{
		"uuid": this.UUID,
	}
	err := util.OracleUpdate("RPM_BIZ_DP_IB_BUSINESS", paramMap, whereParam)
	if nil != err {
		er := fmt.Errorf("更新一对一存款中间收入表信息出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

// 删除一对一存款中间收入表信息
// by author Jason
// by time 2016-12-06 16:59:09
func (this DpIbBusiness) Delete() error {
	whereParam := map[string]interface{}{
		"uuid": this.UUID,
	}
	err := util.OracleDelete("RPM_BIZ_DP_IB_BUSINESS", whereParam)
	if nil != err {
		er := fmt.Errorf("删除一对一存款中间收入表信息出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

var dpIbBusinessTabales string = `
		 RPM_BIZ_DP_IB_BUSINESS T 
	LEFT JOIN RPM_DIM_PRODUCT PRODUCT 
	  ON (T.PRODUCT = PRODUCT.PRODUCT_CODE)`

var dpIbBusinessCols map[string]string = map[string]string{
	"T.BUSINESS_CODE":      "' '",
	"T.UUID":               "' '",
	"T.PRODUCT":            "' '",
	"T.AMOUNT":             "0",
	"T.STATUS":             "' '",
	"T.CREATE_TIME":        "sysdate",
	"T.CREATE_USER":        "' '",
	"T.UPDATE_TIME":        "sysdate",
	"T.UPDATE_USER":        "' '",
	"PRODUCT.PRODUCT_NAME": "' '",
}

var dpIbBusinessColsSort = []string{
	"T.UUID",
	"T.BUSINESS_CODE",
	"T.PRODUCT",
	"T.AMOUNT",
	"T.STATUS",
	"T.CREATE_TIME",
	"T.CREATE_USER",
	"T.UPDATE_TIME",
	"T.UPDATE_USER",
	"PRODUCT.PRODUCT_NAME",
}
