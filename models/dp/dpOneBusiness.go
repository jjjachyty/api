package dp

import (
	"database/sql"
	"fmt"
	"time"

	"pccqcpa.com.cn/app/rpm/api/models/biz/ln"
	"pccqcpa.com.cn/app/rpm/api/models/dim"
	"pccqcpa.com.cn/app/rpm/api/models/sys"
	modelsUtil "pccqcpa.com.cn/app/rpm/api/models/util"
	"pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/components/zlog"
)

// 结构体对应数据库表【RPM_BIZ_DP_ONE_BUSINESS】一对一存款业务表
// by author Jason
// by time 2016-12-06 16:59:09
type DpOneBusiness struct {
	UUID         string      // 主键 默认值sys_guid()
	BusinessCode string      // 业务单号
	Cust         ln.CustInfo // 存款客户
	Organ        sys.Organ   // 机构
	Product      dim.Product // 产品
	Term         int         // 期限
	Ccy          string      // 币种
	Rate         float64     // 利率
	Amount       float64     // 存款金额
	Status       string      // 状态
	CreateTime   time.Time   // 创建时间
	CreateUser   string      // 创建人
	UpdateTime   time.Time   // 更新时间
	UpdateUser   string      // 更新人
}

// 结构体scan
// by author Jason
// by time 2016-12-06 16:59:09
func (this DpOneBusiness) scan(rows *sql.Rows) (*DpOneBusiness, error) {
	var one = new(DpOneBusiness)
	values := []interface{}{
		&one.UUID,
		&one.BusinessCode,
		&one.Cust.CustCode,
		&one.Organ.OrganCode,
		&one.Product.ProductCode,
		&one.Term,
		&one.Ccy,
		&one.Rate,
		&one.Amount,
		&one.Status,
		&one.CreateTime,
		&one.CreateUser,
		&one.UpdateTime,
		&one.UpdateUser,

		&one.Product.ProductName,

		&one.Organ.OrganName,

		&one.Cust.CustName,
	}
	err := util.OracleScan(rows, values)
	return one, err
}

// 分页操作
// by author Jason
// by time 2016-12-06 16:59:09
func (this DpOneBusiness) List(param ...map[string]interface{}) (*util.PageData, error) {
	pageData, rows, err := modelsUtil.List(dpOneBusinessTabales, dpOneBusinessCols, dpOneBusinessColsSort, param...)
	defer rows.Close()
	if nil != err {
		er := fmt.Errorf("分页查询一对一存款业务表信息出错")
		zlog.Error(er.Error(), err)
		return nil, er
	}

	var rst []*DpOneBusiness
	for rows.Next() {
		one, err := this.scan(rows)
		if nil != err {
			er := fmt.Errorf("分页查询一对一存款业务表信息row.Scan()出错")
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
func (this DpOneBusiness) Find(param ...map[string]interface{}) ([]*DpOneBusiness, error) {
	rows, err := modelsUtil.FindRows(dpOneBusinessTabales, dpOneBusinessCols, dpOneBusinessColsSort, param...)
	defer rows.Close()
	if nil != err {
		er := fmt.Errorf("多参数查询一对一存款业务表信息出错")
		zlog.Error(er.Error(), err)
		return nil, er
	}
	var rst []*DpOneBusiness
	for rows.Next() {
		one, err := this.scan(rows)
		if nil != err {
			er := fmt.Errorf("多参数查询一对一存款业务表信息row.Scan()出错")
			zlog.Error(er.Error(), err)
			return nil, er
		}
		rst = append(rst, one)
	}
	return rst, nil
}

// 新增一对一存款业务表信息
// by author Jason
// by time 2016-12-06 16:59:09
func (this DpOneBusiness) Add() error {
	paramMap := map[string]interface{}{
		"business_code": this.BusinessCode,
		"cust":          this.Cust.CustCode,
		"organ":         this.Organ.OrganCode,
		"product":       this.Product.ProductCode,
		"term":          this.Term,
		"ccy":           this.Ccy,
		"rate":          this.Rate,
		"amount":        this.Amount,
		"status":        this.Status,
		"create_time":   util.GetCurrentTime(),
		"create_user":   this.CreateUser,
		"update_time":   util.GetCurrentTime(),
		"update_user":   this.UpdateUser,
	}
	err := util.OracleAdd("RPM_BIZ_DP_ONE_BUSINESS", paramMap)
	if nil != err {
		er := fmt.Errorf("新增一对一存款业务表信息出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

// 更新一对一存款业务表信息
// by author Jason
// by time 2016-12-06 16:59:09
func (this DpOneBusiness) Update() error {
	paramMap := map[string]interface{}{
		"cust":        this.Cust.CustCode,
		"organ":       this.Organ.OrganCode,
		"product":     this.Product.ProductCode,
		"term":        this.Term,
		"ccy":         this.Ccy,
		"rate":        this.Rate,
		"amount":      this.Amount,
		"status":      this.Status,
		"update_time": util.GetCurrentTime(),
		"update_user": this.UpdateUser,
	}
	whereParam := map[string]interface{}{
		"uuid": this.UUID,
	}
	err := util.OracleUpdate("RPM_BIZ_DP_ONE_BUSINESS", paramMap, whereParam)
	if nil != err {
		er := fmt.Errorf("更新一对一存款业务表信息出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

// 删除一对一存款业务表信息
// by author Jason
// by time 2016-12-06 16:59:09
func (this DpOneBusiness) Delete() error {
	whereParam := map[string]interface{}{
		"uuid": this.UUID,
	}
	err := util.OracleDelete("RPM_BIZ_DP_ONE_BUSINESS", whereParam)
	if nil != err {
		er := fmt.Errorf("删除一对一存款业务表信息出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

var dpOneBusinessTabales string = `
		 RPM_BIZ_DP_ONE_BUSINESS T
	LEFT JOIN RPM_DIM_PRODUCT PRODUCT
	  ON (T.PRODUCT = PRODUCT.PRODUCT_CODE)
	LEFT JOIN RPM_BIZ_CUST_INFO CUSTINFO
	  ON (T.CUST = CUSTINFO.CUST_CODE)
	LEFT JOIN SYS_SEC_ORGAN ORGAN
	  ON (T.ORGAN = ORGAN.ORGAN_CODE)`

var dpOneBusinessCols map[string]string = map[string]string{
	"T.UUID":               "' '",
	"T.BUSINESS_CODE":      "' '",
	"T.CUST":               "' '",
	"T.ORGAN":              "' '",
	"T.PRODUCT":            "' '",
	"T.TERM":               "0",
	"T.CCY":                "' '",
	"T.RATE":               "0",
	"T.AMOUNT":             "0",
	"T.STATUS":             "' '",
	"T.CREATE_TIME":        "sysdate",
	"T.CREATE_USER":        "' '",
	"T.UPDATE_TIME":        "sysdate",
	"T.UPDATE_USER":        "' '",
	"PRODUCT.PRODUCT_NAME": "' '",

	"ORGAN.ORGAN_NAME": "' '",

	"CUSTINFO.CUST_NAME": "' '",
}

var dpOneBusinessColsSort = []string{
	"T.UUID",
	"T.BUSINESS_CODE",
	"T.CUST",
	"T.ORGAN",
	"T.PRODUCT",
	"T.TERM",
	"T.CCY",
	"T.RATE",
	"T.AMOUNT",
	"T.STATUS",
	"T.CREATE_TIME",
	"T.CREATE_USER",
	"T.UPDATE_TIME",
	"T.UPDATE_USER",
	"PRODUCT.PRODUCT_NAME",

	"ORGAN.ORGAN_NAME",

	"CUSTINFO.CUST_NAME",
}
