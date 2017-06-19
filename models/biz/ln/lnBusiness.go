package ln

import (
	"database/sql"
	"fmt"
	"time"

	"pccqcpa.com.cn/app/rpm/api/models/dim"
	"pccqcpa.com.cn/app/rpm/api/models/sys"
	modelsUtil "pccqcpa.com.cn/app/rpm/api/models/util"
	"pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/components/zlog"
)

type LnBusiness struct {
	UUID         string      `xorm:"varchar(44) pk notnull unique 'UUID'"`
	BusinessCode string      `xorm:"varchar(44)  notnull unique 'BUSINESS_CODE'"`
	Cust         CustInfo    `xorm:extends`
	Organ        sys.Organ   `xorm:extends`
	Product      dim.Product `xorm:extends`
	Currency     string      `xorm:"VARCHAR2(4) 'TERM'"`
	// LnPricing        LnPricing   `xorm:"VARCHART2(44)"`
	IntRate          float64
	LnPricingStatus  string
	Term             int       `xorm:"VARCHAR2(4) 'TERM'"`
	TermMult         string    `xorm:"VARCHAR2(2) 'TERM_MULT'"`
	RateType         string    `xorm:"VARCHAR2(2) 'RATE_TYPE'"`
	RpymType         string    `xorm:"VARCHAR2(44) 'REPRICE_FREQ'"`
	RepriceFreq      int       `xorm:"NUMBER(4,0) 'REPRICE_FREQ'"`
	RpymInterestFreq int       `xorm:"NUMBER(4,0) 'RPYM_INTEREST_FREQ'"`
	RpymCapitalFreq  int       `xorm:"NUMBER(4,0) 'RPYM_CAPITAL_FREQ'"`
	Principal        float64   `xorm:"NUMBER(32,6)  'PRINCIPAL'"`
	BaseRateType     string    `xorm:"VARCHAR2(2)  'BASE_RATE_TYPE'"`
	MainMortgageType string    `xorm:"VARCHAR2(2)   'MAIN_MORTGAGE_TYPE'"`
	CreateTime       time.Time `xorm:"DATE created  'CREATE_TIME'"`
	CreateUser       string    `xorm:"VARCHAR2(44)   'CREATE_USER'"`
	UpdateTime       time.Time `xorm:"DATE updated  'UPDATE_TIME'"`
	UpdateUser       string    `xorm:"VARCHAR2(44)   'UPDATE_USER'"`
	Flag             string    `xorm:"VARCHAR2(2)   'FLAG'"`
	Status           string    `xorm:"VARCHAR2(2)   'STATUS'"`
	ResChr1          string    `xorm:"VARCHAR2(44)   'RES_CHR1'"`
	ResChr2          string    `xorm:"VARCHAR2(44)   'RES_CHR2'"`
	ResChr3          string    `xorm:"VARCHAR2(44)   'RES_CHR3'"`
	ResChr4          string    `xorm:"VARCHAR2(44)   'RES_CHR4'"`
	ResChr5          string    `xorm:"VARCHAR2(44)   'RES_CHR5'"`
	ResChr6          string    `xorm:"VARCHAR2(44)   'RES_CHR6'"`
	ResChr7          string    `xorm:"VARCHAR2(44)   'RES_CHR7'"`
	ResChr8          string    `xorm:"VARCHAR2(44)   'RES_CHR8'"`
	ResChr9          string    `xorm:"VARCHAR2(44)   'RES_CHR9'"`
	ResChr10         string    `xorm:"VARCHAR2(44)   'RES_CHR10'"`
	ResNum1          float64   `xorm:"NUMBER(32,6) 'RES_NUM1'"`
	ResNum2          float64   `xorm:"NUMBER(32,6) 'RES_NUM2'"`
	ResNum3          float64   `xorm:"NUMBER(32,6) 'RES_NUM3'"`
	ResNum4          float64   `xorm:"NUMBER(32,6) 'RES_NUM4'"`
	ResNum5          float64   `xorm:"NUMBER(32,6) 'RES_NUM5'"`
	ResNum6          float64   `xorm:"NUMBER(32,6) 'RES_NUM6'"`
	ResNum7          float64   `xorm:"NUMBER(32,6) 'RES_NUM7'"`
	ResNum8          float64   `xorm:"NUMBER(32,6) 'RES_NUM8'"`
	ResNum9          float64   `xorm:"NUMBER(32,6) 'RES_NUM9'"`
	ResNum10         float64   `xorm:"NUMBER(32,6) 'RES_NUM10'"`
}

type NullLnBusiness struct {
	UUID             sql.NullString
	BusinessCode     sql.NullString
	Cust             CustInfo
	Organ            sys.NullOrgan
	Product          dim.NullProduct
	Currency         sql.NullString
	Term             sql.NullInt64
	TermMult         sql.NullString
	RateType         sql.NullString
	RpymType         sql.NullString
	RepriceFreq      sql.NullInt64
	RpymInterestFreq sql.NullInt64
	RpymCapitalFreq  sql.NullInt64
	Principal        sql.NullFloat64
	BaseRateType     sql.NullString
	MainMortgageType sql.NullString
	CreateTime       time.Time
	CreateUser       sql.NullString
	UpdateTime       time.Time
	UpdateUser       sql.NullString
	Flag             sql.NullString
	Status           sql.NullString
	ResChr1          sql.NullString
	ResChr2          sql.NullString
	ResChr3          sql.NullString
	ResChr4          sql.NullString
	ResChr5          sql.NullString
	ResChr6          sql.NullString
	ResChr7          sql.NullString
	ResChr8          sql.NullString
	ResChr9          sql.NullString
	ResChr10         sql.NullString
	ResNum1          sql.NullFloat64
	ResNum2          sql.NullFloat64
	ResNum3          sql.NullFloat64
	ResNum4          sql.NullFloat64
	ResNum5          sql.NullFloat64
	ResNum6          sql.NullFloat64
	ResNum7          sql.NullFloat64
	ResNum8          sql.NullFloat64
	ResNum9          sql.NullFloat64
	ResNum10         sql.NullFloat64
}

func (l *LnBusiness) scan(rows *sql.Rows) (*LnBusiness, error) {
	var lnBusiness = new(LnBusiness)
	values := []interface{}{
		&lnBusiness.UUID,
		&lnBusiness.BusinessCode,
		&lnBusiness.Currency,
		&lnBusiness.Term,
		&lnBusiness.TermMult,
		&lnBusiness.RateType,
		&lnBusiness.RpymType,
		&lnBusiness.RepriceFreq,
		&lnBusiness.RpymInterestFreq,
		&lnBusiness.RpymCapitalFreq,
		&lnBusiness.Principal,
		&lnBusiness.BaseRateType,
		&lnBusiness.MainMortgageType,
		&lnBusiness.CreateTime,
		&lnBusiness.UpdateTime,
		&lnBusiness.CreateUser,
		&lnBusiness.UpdateUser,
		&lnBusiness.Flag,
		// &lnBusiness.Status,

		&lnBusiness.Cust.CustCode,
		&lnBusiness.Cust.CustName,
		&lnBusiness.Cust.Organization,
		&lnBusiness.Cust.CustType,
		&lnBusiness.Cust.CustImplvl,
		&lnBusiness.Cust.CustCredit,
		&lnBusiness.Cust.Branch.OrganCode,
		&lnBusiness.Cust.Industry.IndustryCode,
		&lnBusiness.Cust.CustSize,
		&lnBusiness.Cust.CustCapital,
		&lnBusiness.Cust.GapProportion,
		&lnBusiness.Cust.StockContribute,
		&lnBusiness.Cust.StockFreeze,
		&lnBusiness.Cust.StockUsage,
		&lnBusiness.Cust.UseProduct,
		&lnBusiness.Cust.CooperationPeriod,

		&lnBusiness.Organ.OrganCode,
		&lnBusiness.Organ.OrganName,
		&lnBusiness.Organ.OrganLevel,
		&lnBusiness.Organ.ParentOrgan,
		&lnBusiness.Organ.LeafFlag,

		&lnBusiness.Product.ProductCode,
		&lnBusiness.Product.ProductName,
		&lnBusiness.Product.ProductType,
		&lnBusiness.Product.ProductTypeDesc,
		&lnBusiness.Product.ProductLevel,

		&lnBusiness.Cust.Industry.IndustryCode,
		&lnBusiness.Cust.Industry.IndustryName,

		&lnBusiness.IntRate,
		&lnBusiness.Status,
	}
	err := util.OracleScan(rows, values)
	// fmt.Println("--我是查询后的business信息--", lnBusiness.Term, lnBusiness.TermMult)
	return lnBusiness, err

}

//查询对公贷款业务信息并分页
func (l *LnBusiness) List(param ...map[string]interface{}) (*util.PageData, error) {
	zlog.AppOperateLog("", "SelectLnBusinessByPage", zlog.SELECT, nil, nil, "业务信息分页查询操作")

	pageData, rows, err := modelsUtil.List(lnBusinessTables, lnBusinessCols, lnBusinessColsSort, param...)
	defer rows.Close()
	if nil != err {
		er := fmt.Errorf("分页查询对公业务信息出错")
		zlog.Error(er.Error(), err)
		return nil, er
	}

	lns, err := l.handleRows(rows)
	if nil != err {
		return nil, err
	}
	pageData.Rows = lns
	return pageData, nil
}

func (l LnBusiness) UnionList(paramMap map[string]interface{}, unionParam ...map[string]interface{}) (*util.PageData, error) {
	modelUtil := modelsUtil.NewModelUtil()
	modelUtil.SetTableMsg(lnBusinessTables, lnBusinessCols, lnBusinessColsSort)
	var sqls = make([]string, 0)
	var sqlParam = make([]interface{}, 0)
	for _, param := range unionParam {
		modelUtil.ParamMap = param
		sql, param := modelUtil.GetSelectWhereSql()
		sqlParam = append(sqlParam, param...)
		sqls = append(sqls, sql)
	}
	unionSql := modelUtil.UnionAll(sqls...)
	pageData, rows, err := modelUtil.GetPageData(unionSql, "", paramMap)
	if nil != err {
		rows.Close()
		return nil, err
	}
	defer rows.Close()

	lns, err := l.handleRows(rows)
	if nil != err {
		return nil, err
	}
	pageData.Rows = lns
	return pageData, nil
}

func (l LnBusiness) handleRows(rows *sql.Rows) ([]*LnBusiness, error) {
	var lns []*LnBusiness
	for rows.Next() {
		ln, err := l.scan(rows)
		if nil != err {
			er := fmt.Errorf("查询客户分页row.Scan()出错")
			zlog.Error(er.Error(), err)
			return nil, er
		}
		lns = append(lns, ln)
	}

	return lns, nil
}

func (l *LnBusiness) Find(param ...map[string]interface{}) ([]*LnBusiness, error) {
	rows, err := modelsUtil.FindRows(lnBusinessTables, lnBusinessCols, lnBusinessColsSort, param...)
	defer rows.Close()
	if nil != err {
		er := fmt.Errorf("查询对公业务信息出错")
		zlog.Error(er.Error(), err)
		return nil, er
	}
	var lnBusiness []*LnBusiness
	for rows.Next() {
		business, err := l.scan(rows)
		if nil != err {
			er := fmt.Errorf("查询对公业务信息rows.Scan()出错")
			zlog.Error(er.Error(), err)
			return nil, er
		}
		lnBusiness = append(lnBusiness, business)
	}
	return lnBusiness, nil
}

func (l *LnBusiness) Add() error {
	param := map[string]interface{}{
		"business_code":      l.BusinessCode,
		"cust":               l.Cust.CustCode,
		"product":            l.Product.ProductCode,
		"organ":              l.Organ.OrganCode,
		"currency":           l.Currency,
		"term":               l.Term,
		"term_mult":          l.TermMult,
		"rate_type":          l.RateType,
		"rpym_type":          l.RpymType,
		"reprice_freq":       l.RepriceFreq,
		"rpym_interest_freq": l.RpymInterestFreq,
		"rpym_capital_freq":  l.RpymCapitalFreq,
		"principal":          l.Principal,
		"base_rate_type":     l.BaseRateType,
		"main_mortgage_type": l.MainMortgageType,
		"create_time":        util.GetCurrentTime(),
		"update_time":        util.GetCurrentTime(),
		"create_user":        l.CreateUser,
		"update_user":        l.UpdateUser,
		"flag":               l.Flag,
		"status":             l.Status,
	}
	err := util.OracleAdd(lnBusinessTableName, param)
	if nil != err {
		er := fmt.Errorf("新增对公业务信息出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

func (l *LnBusiness) Update() error {
	param := map[string]interface{}{
		"business_code":      l.BusinessCode,
		"cust":               l.Cust.CustCode,
		"product":            l.Product.ProductCode,
		"organ":              l.Organ.OrganCode,
		"currency":           l.Currency,
		"term":               l.Term,
		"term_mult":          l.TermMult,
		"rate_type":          l.RateType,
		"rpym_type":          l.RpymType,
		"reprice_freq":       l.RepriceFreq,
		"rpym_interest_freq": l.RpymInterestFreq,
		"rpym_capital_freq":  l.RpymCapitalFreq,
		"principal":          l.Principal,
		"base_rate_type":     l.BaseRateType,
		"main_mortgage_type": l.MainMortgageType,
		"update_time":        util.GetCurrentTime(),
		"update_user":        l.UpdateUser,
		"flag":               l.Flag,
		"status":             l.Status,
	}
	whereParam := map[string]interface{}{
		"business_code": l.BusinessCode,
		// "uuid": l.UUID,
	}
	err := util.OracleUpdate(lnBusinessTableName, param, whereParam)
	if nil != err {
		er := fmt.Errorf("更新对公业务信息出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

func (l LnBusiness) Patch(param map[string]interface{}, businessCode string) error {
	whereParam := map[string]interface{}{
		"business_code": businessCode,
	}
	param["update_time"] = util.GetCurrentTime()
	err := util.OracleUpdate(lnBusinessTableName, param, whereParam)
	if nil != err {
		er := fmt.Errorf("更新对公业务部分信息出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

func (l *LnBusiness) Delete() error {
	whereParam := map[string]interface{}{
		"business_code": l.BusinessCode,
	}
	err := util.OracleDelete(lnBusinessTableName, whereParam)
	if nil != err {
		er := fmt.Errorf("删除对公业务信息出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

var lnBusinessTableName = "rpm_biz_ln_coltd_business"

var lnBusinessTables string = `
	       RPM_BIZ_LN_COLTD_BUSINESS T
	  LEFT JOIN RPM_BIZ_CUST_INFO CUST
	    ON (T.CUST = CUST.CUST_CODE)
	  LEFT JOIN SYS_SEC_ORGAN ORGAN
	    ON (T.ORGAN = ORGAN.ORGAN_CODE)
	  LEFT JOIN RPM_DIM_PRODUCT PRODUCT
	    ON (T.PRODUCT = PRODUCT.PRODUCT_CODE)
	  LEFT JOIN RPM_DIM_INDUSTRY INDUSTRY
	    ON (CUST.INDUSTRY = INDUSTRY.INDUSTRY_CODE)
	  LEFT JOIN RPM_BIZ_LN_COLTD_PRICING LNPRICING
	    ON (T.BUSINESS_CODE = LNPRICING.BUSINESS_CODE)
`

var lnBusinessCols = map[string]string{
	"T.UUID":                    "''",
	"T.BUSINESS_CODE":           "' '",
	"T.CURRENCY":                "' '",
	"T.TERM":                    "0",
	"T.TERM_MULT":               "' '",
	"T.RATE_TYPE":               "' '",
	"T.RPYM_TYPE":               "' '",
	"T.REPRICE_FREQ":            "0",
	"T.RPYM_INTEREST_FREQ":      "0",
	"T.RPYM_CAPITAL_FREQ":       "0",
	"T.PRINCIPAL":               "0",
	"T.BASE_RATE_TYPE":          "' '",
	"T.MAIN_MORTGAGE_TYPE":      "' '",
	"T.CREATE_TIME":             "sysdate",
	"T.UPDATE_TIME":             "sysdate",
	"T.CREATE_USER":             "' '",
	"T.UPDATE_USER":             "' '",
	"T.FLAG":                    "' '",
	"T.STATUS":                  "' '",
	"CUST.CUST_CODE":            "' '",
	"CUST.CUST_NAME":            "' '",
	"CUST.ORGANIZATION":         "' '",
	"CUST.CUST_TYPE":            "' '",
	"CUST.CUST_IMPLVL":          "' '",
	"CUST.CUST_CREDIT":          "' '",
	"CUST.BRANCH":               "' '",
	"CUST.INDUSTRY":             "' '",
	"CUST.CUST_SIZE":            "' '",
	"CUST.CUST_CAPITAL":         "0",
	"CUST.GAP_PROPORTION":       "0",
	"CUST.STOCK_CONTRIBUTE":     "0",
	"CUST.STOCK_FREEZE":         "0",
	"CUST.STOCK_USAGE":          "0",
	"CUST.USE_PRODUCT":          "0",
	"CUST.COOPERATION_PERIOD":   "0",
	"ORGAN.ORGAN_CODE":          "' '",
	"ORGAN.ORGAN_NAME":          "' '",
	"ORGAN.ORGAN_LEVEL":         "' '",
	"ORGAN.PARENT_ORGAN":        "' '",
	"ORGAN.LEAF_FLAG":           "'1'",
	"PRODUCT.PRODUCT_CODE":      "' '",
	"PRODUCT.PRODUCT_NAME":      "' '",
	"PRODUCT.PRODUCT_TYPE":      "' '",
	"PRODUCT.PRODUCT_TYPE_DESC": "' '",
	"PRODUCT.PRODUCT_LEVEL":     "' '",
	"INDUSTRY.INDUSTRY_CODE":    "' '",
	"INDUSTRY.INDUSTRY_NAME":    "' '",

	"LNPRICING.INT_RATE": "0",
	"LNPRICING.STATUS":   "'0'",
}

var lnBusinessColsSort = []string{
	"T.UUID",
	"T.BUSINESS_CODE",
	"T.CURRENCY",
	"T.TERM",
	"T.TERM_MULT",
	"T.RATE_TYPE",
	"T.RPYM_TYPE",
	"T.REPRICE_FREQ",
	"T.RPYM_INTEREST_FREQ",
	"T.RPYM_CAPITAL_FREQ",
	"T.PRINCIPAL",
	"T.BASE_RATE_TYPE",
	"T.MAIN_MORTGAGE_TYPE",
	"T.CREATE_TIME",
	"T.UPDATE_TIME",
	"T.CREATE_USER",
	"T.UPDATE_USER",
	"T.FLAG",
	// "T.STATUS",
	"CUST.CUST_CODE",
	"CUST.CUST_NAME",
	"CUST.ORGANIZATION",
	"CUST.CUST_TYPE",
	"CUST.CUST_IMPLVL",
	"CUST.CUST_CREDIT",
	"CUST.BRANCH",
	"CUST.INDUSTRY",
	"CUST.CUST_SIZE",
	"CUST.CUST_CAPITAL",
	"CUST.GAP_PROPORTION",
	"CUST.STOCK_CONTRIBUTE",
	"CUST.STOCK_FREEZE",
	"CUST.STOCK_USAGE",
	"CUST.USE_PRODUCT",
	"CUST.COOPERATION_PERIOD",
	"ORGAN.ORGAN_CODE",
	"ORGAN.ORGAN_NAME",
	"ORGAN.ORGAN_LEVEL",
	"ORGAN.PARENT_ORGAN",
	"ORGAN.LEAF_FLAG",
	"PRODUCT.PRODUCT_CODE",
	"PRODUCT.PRODUCT_NAME",
	"PRODUCT.PRODUCT_TYPE",
	"PRODUCT.PRODUCT_TYPE_DESC",
	"PRODUCT.PRODUCT_LEVEL",
	"INDUSTRY.INDUSTRY_CODE",
	"INDUSTRY.INDUSTRY_NAME",

	"LNPRICING.INT_RATE",
	"LNPRICING.STATUS",
}
