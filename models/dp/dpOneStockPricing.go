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

// 结构体对应数据库表【RPM_BIZ_DP_ONE_STOCK_PRICING】一对一存款存量威胁表-用于调整利率威胁使用
// by author Jason
// by time 2017-01-04 10:46:11
type DpOneStockPricing struct {
	UUID               string      // 主键 默认值SYS_GUID()
	BusinessCode       string      // 业务编号
	DpOneStockBusiness string      // 存量业务业务
	Cust               ln.CustInfo // 客户
	Organ              sys.Organ   // 机构
	Product            dim.Product // 产品
	Ccy                string      // 币种
	RateOld            float64     // 原利率
	RateNew            float64     // 修改利率
	AmountOld          float64     // 金额
	AmountUse          float64     // 威胁金额
	AmountLost         float64     // 流失金额
	CurrentAmount      float64     // 当前余额（直接取RPM_BIZ_DP_ONE_STOCK_BUSINESS表该字段值）
	RestTerm           int         // 剩余期限（直接取RPM_BIZ_DP_ONE_STOCK_BUSINESS表该字段值）
	StartOfDate        time.Time   // 起息日
	EndOfDate          time.Time   // 到期日
	CreateTime         time.Time   // 创建时间
	UpdateTime         time.Time   // 更新时间
	CreateUser         string      // 创建人
	UpdateUser         string      // 更新人
}

// 结构体scan
// by author Jason
// by time 2017-01-04 10:46:11
func (this DpOneStockPricing) scan(rows *sql.Rows) (*DpOneStockPricing, error) {
	var one = new(DpOneStockPricing)
	values := []interface{}{
		&one.UUID,
		&one.BusinessCode,
		&one.DpOneStockBusiness,
		&one.Cust.CustCode,
		&one.Organ.OrganCode,
		&one.Ccy,
		&one.RateOld,
		&one.RateNew,
		&one.AmountOld,
		&one.AmountUse,
		&one.CurrentAmount,
		&one.RestTerm,
		&one.StartOfDate,
		&one.EndOfDate,
		&one.CreateTime,
		&one.UpdateTime,
		&one.CreateUser,
		&one.UpdateUser,

		&one.Product.ProductCode,
		&one.Product.ProductName,
	}
	err := util.OracleScan(rows, values)
	return one, err
}

// 分页操作
// by author Jason
// by time 2017-01-04 10:46:11
func (this DpOneStockPricing) List(param ...map[string]interface{}) (*util.PageData, error) {
	pageData, rows, err := modelsUtil.List(dpOneStockPricingTabales, dpOneStockPricingCols, dpOneStockPricingColsSort, param...)
	defer rows.Close()
	if nil != err {
		er := fmt.Errorf("分页查询一对一存款存量威胁表-用于调整利率威胁使用信息出错")
		zlog.Error(er.Error(), err)
		return nil, er
	}

	var rst []*DpOneStockPricing
	for rows.Next() {
		one, err := this.scan(rows)
		if nil != err {
			er := fmt.Errorf("分页查询一对一存款存量威胁表-用于调整利率威胁使用信息row.Scan()出错")
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
func (this DpOneStockPricing) Find(param ...map[string]interface{}) ([]*DpOneStockPricing, error) {
	rows, err := modelsUtil.FindRows(dpOneStockPricingTabales, dpOneStockPricingCols, dpOneStockPricingColsSort, param...)
	defer rows.Close()
	if nil != err {
		er := fmt.Errorf("多参数查询一对一存款存量威胁表-用于调整利率威胁使用信息出错")
		zlog.Error(er.Error(), err)
		return nil, er
	}
	var rst []*DpOneStockPricing
	for rows.Next() {
		one, err := this.scan(rows)
		if nil != err {
			er := fmt.Errorf("多参数查询一对一存款存量威胁表-用于调整利率威胁使用信息row.Scan()出错")
			zlog.Error(er.Error(), err)
			return nil, er
		}
		rst = append(rst, one)
	}
	return rst, nil
}

// 新增一对一存款存量威胁表-用于调整利率威胁使用信息
// by author Jason
// by time 2017-01-04 10:46:11
func (this DpOneStockPricing) Add() error {
	paramMap := map[string]interface{}{
		"business_code":         this.BusinessCode,
		"dp_one_stock_business": this.DpOneStockBusiness,
		"cust":                  this.Cust.CustCode,
		"organ":                 this.Organ.OrganCode,
		"product":               this.Product.ProductCode,
		"ccy":                   this.Ccy,
		// "save_term":     this.SaveTerm,
		// "rest_term":     this.RestTerm,
		"rate_old":       this.RateOld,
		"rate_new":       this.RateNew,
		"amount_old":     this.AmountOld,
		"amount_use":     this.AmountUse,
		"current_amount": this.CurrentAmount,
		"rest_term":      this.RestTerm,
		"start_of_date":  this.StartOfDate,
		"end_of_date":    this.EndOfDate,
		"create_time":    util.GetCurrentTime(),
		"update_time":    util.GetCurrentTime(),
		"create_user":    this.CreateUser,
		"update_user":    this.UpdateUser,
	}
	err := util.OracleAdd("RPM_BIZ_DP_ONE_STOCK_PRICING", paramMap)
	if nil != err {
		er := fmt.Errorf("新增一对一存款存量威胁表-用于调整利率威胁使用信息出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

// 更新一对一存款存量威胁表-用于调整利率威胁使用信息
// by author Jason
// by time 2017-01-04 10:46:11
func (this DpOneStockPricing) Update() error {
	paramMap := map[string]interface{}{
		"business_code":         this.BusinessCode,
		"dp_one_stock_business": this.DpOneStockBusiness,
		"cust":                  this.Cust.CustCode,
		"organ":                 this.Organ.OrganCode,
		"product":               this.Product.ProductCode,
		"ccy":                   this.Ccy,
		// "save_term":     this.SaveTerm,
		// "rest_term":     this.RestTerm,
		"rate_old":       this.RateOld,
		"rate_new":       this.RateNew,
		"amount_old":     this.AmountOld,
		"amount_use":     this.AmountUse,
		"current_amount": this.CurrentAmount,
		"rest_term":      this.RestTerm,
		"start_of_date":  this.StartOfDate,
		"end_of_date":    this.EndOfDate,
		"update_time":    util.GetCurrentTime(),
		"update_user":    this.UpdateUser,
	}
	whereParam := map[string]interface{}{
		"uuid": this.UUID,
	}
	err := util.OracleUpdate("RPM_BIZ_DP_ONE_STOCK_PRICING", paramMap, whereParam)
	if nil != err {
		er := fmt.Errorf("更新一对一存款存量威胁表-用于调整利率威胁使用信息出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

// 删除一对一存款存量威胁表-用于调整利率威胁使用信息
// by author Jason
// by time 2017-01-04 10:46:11
func (this DpOneStockPricing) Delete() error {
	whereParam := map[string]interface{}{
		"uuid": this.UUID,
	}
	err := util.OracleDelete("RPM_BIZ_DP_ONE_STOCK_PRICING", whereParam)
	if nil != err {
		er := fmt.Errorf("删除一对一存款存量威胁表-用于调整利率威胁使用信息出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

var dpOneStockPricingTabales string = `
			RPM_BIZ_DP_ONE_STOCK_PRICING T
	   LEFT JOIN RPM_BIZ_CUST_INFO CUST
	     ON (T.CUST = CUST.CUST_CODE)
	   LEFT JOIN RPM_DIM_PRODUCT PRODUCT
	     ON (T.PRODUCT = PRODUCT.PRODUCT_CODE)
	   LEFT JOIN SYS_SEC_ORGAN ORGAN
	     ON (T.ORGAN = ORGAN.ORGAN_CODE)`

var dpOneStockPricingCols map[string]string = map[string]string{
	"T.UUID":                  "' '",
	"T.BUSINESS_CODE":         "' '",
	"T.DP_ONE_STOCK_BUSINESS": "' '",
	"T.CUST":                  "' '",
	"T.ORGAN":                 "' '",
	"T.CCY":                   "' '",
	// "T.SAVE_TERM":     "0",
	// "T.REST_TERM":     "0",
	"T.RATE_OLD":       "0",
	"T.RATE_NEW":       "0",
	"T.AMOUNT_OLD":     "0",
	"T.AMOUNT_USE":     "0",
	"T.CURRENT_AMOUNT": "0",
	"T.REST_TERM":      "0",
	"T.START_OF_DATE":  "sysdate",
	"T.END_OF_DATE":    "sysdate",
	"T.CREATE_TIME":    "sysdate",
	"T.UPDATE_TIME":    "sysdate",
	"T.CREATE_USER":    "' '",
	"T.UPDATE_USER":    "' '",

	"T.PRODUCT":            "' '",
	"PRODUCT.PRODUCT_NAME": "' '",
}

var dpOneStockPricingColsSort = []string{
	"T.UUID",
	"T.BUSINESS_CODE",
	"T.DP_ONE_STOCK_BUSINESS",
	"T.CUST",
	"T.ORGAN",
	"T.CCY",
	// "T.SAVE_TERM",
	// "T.REST_TERM",
	"T.RATE_OLD",
	"T.RATE_NEW",
	"T.AMOUNT_OLD",
	"T.AMOUNT_USE",
	"T.CURRENT_AMOUNT",
	"T.REST_TERM",
	"T.START_OF_DATE",
	"T.END_OF_DATE",
	"T.CREATE_TIME",
	"T.UPDATE_TIME",
	"T.CREATE_USER",
	"T.UPDATE_USER",

	"T.PRODUCT",
	"PRODUCT.PRODUCT_NAME",
}
