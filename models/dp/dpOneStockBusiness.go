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

// 结构体对应数据库表【RPM_BIZ_DP_ONE_STOCK_BUSINESS】一对一存款存量表-每日跑批更新
// by author Jason
// by time 2017-01-04 10:46:11
type DpOneStockBusiness struct {
	UUID          string      // 主键 默认值SYS_GUID()
	BusinessCode  string      // 业务编号
	Cust          ln.CustInfo // 客户
	Organ         sys.Organ   // 机构
	Product       dim.Product // 产品
	Ccy           string      // 币种
	Term          float64     // 期限
	StartOfDate   time.Time   // 起息日
	EndOfDate     time.Time   // 到期日
	Amount        float64     // 金额
	CurrentAmount float64     // 当前余额（考虑有提前支取部分的情况）
	RestTerm      int         // 剩余期限（天）
	CreateTime    time.Time   // 创建时间
	UpdateTime    time.Time   // 更新时间
	CreateUser    string      // 创建人
	UpdateUser    string      // 更新人
	Rate          float64     // 执行利率
}

// 结构体scan
// by author Jason
// by time 2017-01-04 10:46:11
func (this DpOneStockBusiness) scan(rows *sql.Rows) (*DpOneStockBusiness, error) {
	var one = new(DpOneStockBusiness)
	values := []interface{}{
		&one.UUID,
		&one.BusinessCode,
		&one.Cust.CustCode,
		&one.Organ.OrganCode,
		&one.Ccy,
		&one.Term,
		&one.StartOfDate,
		&one.EndOfDate,
		&one.Amount,
		&one.CurrentAmount,
		&one.RestTerm,
		&one.CreateTime,
		&one.UpdateTime,
		&one.CreateUser,
		&one.UpdateUser,
		&one.Rate,

		&one.Product.ProductCode,
		&one.Product.ProductName,
	}
	err := util.OracleScan(rows, values)
	return one, err
}

// 分页操作
// by author Jason
// by time 2017-01-04 10:46:11
func (this DpOneStockBusiness) List(param ...map[string]interface{}) (*util.PageData, error) {
	pageData, rows, err := modelsUtil.List(dpOneStockBusinessTabales, dpOneStockBusinessCols, dpOneStockBusinessColsSort, param...)
	defer rows.Close()
	if nil != err {
		er := fmt.Errorf("分页查询一对一存款存量表-每日跑批更新信息出错")
		zlog.Error(er.Error(), err)
		return nil, er
	}

	var rst []*DpOneStockBusiness
	for rows.Next() {
		one, err := this.scan(rows)
		if nil != err {
			er := fmt.Errorf("分页查询一对一存款存量表-每日跑批更新信息row.Scan()出错")
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
// by time 2017-01-04 10:46:11
func (this DpOneStockBusiness) Find(param ...map[string]interface{}) ([]*DpOneStockBusiness, error) {
	rows, err := modelsUtil.FindRows(dpOneStockBusinessTabales, dpOneStockBusinessCols, dpOneStockBusinessColsSort, param...)
	defer rows.Close()
	if nil != err {
		er := fmt.Errorf("多参数查询一对一存款存量表-每日跑批更新信息出错")
		zlog.Error(er.Error(), err)
		return nil, er
	}
	var rst []*DpOneStockBusiness
	for rows.Next() {
		one, err := this.scan(rows)
		if nil != err {
			er := fmt.Errorf("多参数查询一对一存款存量表-每日跑批更新信息row.Scan()出错")
			zlog.Error(er.Error(), err)
			return nil, er
		}
		rst = append(rst, one)
	}
	return rst, nil
}

// 新增一对一存款存量表-每日跑批更新信息
// by author Jason
// by time 2017-01-04 10:46:11
func (this DpOneStockBusiness) Add() error {
	paramMap := map[string]interface{}{
		"business_code":  this.BusinessCode,
		"cust":           this.Cust,
		"organ":          this.Organ,
		"product":        this.Product.ProductCode,
		"ccy":            this.Ccy,
		"term":           this.Term,
		"start_of_date":  this.StartOfDate,
		"end_of_date":    this.EndOfDate,
		"amount":         this.Amount,
		"current_amount": this.CurrentAmount,
		"rest_term":      this.RestTerm,
		"create_time":    util.GetCurrentTime(),
		"update_time":    util.GetCurrentTime(),
		"create_user":    this.CreateUser,
		"update_user":    this.UpdateUser,
		"rate":           this.Rate,
	}
	err := util.OracleAdd("RPM_BIZ_DP_ONE_STOCK_BUSINESS", paramMap)
	if nil != err {
		er := fmt.Errorf("新增一对一存款存量表-每日跑批更新信息出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

// 更新一对一存款存量表-每日跑批更新信息
// by author Jason
// by time 2017-01-04 10:46:11
func (this DpOneStockBusiness) Update() error {
	paramMap := map[string]interface{}{
		"business_code":  this.BusinessCode,
		"cust":           this.Cust,
		"organ":          this.Organ,
		"product":        this.Product.ProductCode,
		"ccy":            this.Ccy,
		"start_of_date":  this.StartOfDate,
		"term":           this.Term,
		"end_of_date":    this.EndOfDate,
		"amount":         this.Amount,
		"current_amount": this.CurrentAmount,
		"rest_term":      this.RestTerm,
		"update_time":    util.GetCurrentTime(),
		"update_user":    this.UpdateUser,
		"rate":           this.Rate,
	}
	whereParam := map[string]interface{}{
		"uuid": this.UUID,
	}
	err := util.OracleUpdate("RPM_BIZ_DP_ONE_STOCK_BUSINESS", paramMap, whereParam)
	if nil != err {
		er := fmt.Errorf("更新一对一存款存量表-每日跑批更新信息出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

// 删除一对一存款存量表-每日跑批更新信息
// by author Jason
// by time 2017-01-04 10:46:11
func (this DpOneStockBusiness) Delete() error {
	whereParam := map[string]interface{}{
		"uuid": this.UUID,
	}
	err := util.OracleDelete("RPM_BIZ_DP_ONE_STOCK_BUSINESS", whereParam)
	if nil != err {
		er := fmt.Errorf("删除一对一存款存量表-每日跑批更新信息出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

var dpOneStockBusinessTabales string = `
				RPM_BIZ_DP_ONE_STOCK_BUSINESS T
		   LEFT JOIN RPM_DIM_PRODUCT PRODUCT
		     ON (T.PRODUCT = PRODUCT.PRODUCT_CODE)`

var dpOneStockBusinessCols map[string]string = map[string]string{
	"T.UUID":           "' '",
	"T.BUSINESS_CODE":  "' '",
	"T.CUST":           "' '",
	"T.ORGAN":          "' '",
	"T.CCY":            "' '",
	"T.TERM":           "0",
	"T.START_OF_DATE":  "sysdate",
	"T.END_OF_DATE":    "sysdate",
	"T.AMOUNT":         "0",
	"T.CURRENT_AMOUNT": "0",
	"T.REST_TERM":      "0",
	"T.CREATE_TIME":    "sysdate",
	"T.UPDATE_TIME":    "sysdate",
	"T.CREATE_USER":    "' '",
	"T.UPDATE_USER":    "' '",
	"T.RATE":           "0",

	"T.PRODUCT":            "' '",
	"PRODUCT.PRODUCT_NAME": "' '",
}

var dpOneStockBusinessColsSort = []string{
	"T.UUID",
	"T.BUSINESS_CODE",
	"T.CUST",
	"T.ORGAN",
	"T.CCY",
	"T.TERM",
	"T.START_OF_DATE",
	"T.END_OF_DATE",
	"T.AMOUNT",
	"T.CURRENT_AMOUNT",
	"T.REST_TERM",
	"T.CREATE_TIME",
	"T.UPDATE_TIME",
	"T.CREATE_USER",
	"T.UPDATE_USER",
	"T.RATE",

	"T.PRODUCT",
	"PRODUCT.PRODUCT_NAME",
}
