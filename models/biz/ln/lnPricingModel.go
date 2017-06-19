package ln

import (
	"database/sql"
	"errors"
	"strings"

	"time"

	"platform/dbobj"

	"pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/components/zlog"

	"fmt"

	modelsUtil "pccqcpa.com.cn/app/rpm/api/models/util"
)

//LnPricingModel type 定价单数据库模型类
type LnPricingModel struct {
}

//LnPricing type 对公贷款定价单实体
type LnPricing struct {
	UUID                      string  `xorm:"varchar(44) pk notnull unique 'UUID'"`
	BusinessCode              string  `xorm:"varchar(44)   'BUSINESS_CODE'"`
	ContractCode              string  `xorm:"varchar(44)   'CONTRACT_CODE'"`
	PlnCode                   string  `xorm:"varchar(44)   'PLN_CODE'"`
	CustCode                  string  `xorm:"varchar(44)   'CUST_CODE'"`
	CustName                  string  `xorm:"varchar(44)   'CUST_NAME'"`
	CustType                  string  `xorm:"varchar(44)   'CUST_TYPE'"`
	CustImplvl                string  `xorm:"varchar(44)   'CUST_IMPLVL'"`
	CustCredit                string  `xorm:"varchar(44)   'CUST_CREDIT'"`
	BranchCode                string  `xorm:"varchar(44)   'BRANCH_CODE'"`
	BranchName                string  `xorm:"varchar(44)   'BRANCH_NAME'"`
	IndustryCode              string  `xorm:"varchar(44)   'INDUSTRY_NAME'"`
	IndustryName              string  `xorm:"varchar(44)   'INDUSTRY_NAME'"`
	FtpRate                   float64 `xorm:"NUMBER(32,6) 'FTP_RATE'"`
	OcRate                    float64 `xorm:"NUMBER(32,6) 'OC_RATE'"`
	PdRate                    float64 `xorm:"NUMBER(32,6) 'PD_RATE'"`
	LgdRate                   float64 `xorm:"NUMBER(32,6) 'LGD_RATE'"`
	ElRate                    float64 `xorm:"NUMBER(32,6) 'EL_RATE'"`
	EcRate                    float64 `xorm:"NUMBER(32,6) 'EC_RATE'"`
	CapCostRate               float64 `xorm:"NUMBER(32,6) 'CAP_COST_RATE'"`
	CapPftRate                float64 `xorm:"NUMBER(32,6) 'CAP_PFT_RATE'"`
	IncomeTax                 float64 `xorm:"NUMBER(32,6) 'INCOME_TAX'"`
	SalesTax                  float64 `xorm:"NUMBER(32,6) 'SALES_TAX'"`
	AddTax                    float64 `xorm:"NUMBER(32,6) 'ADD_TAX'"`
	StockUsage                float64
	StockPoints               float64
	DerviedDPEVA              float64
	DerviedDPPoints           float64
	DerviedIBEVA              float64
	DerviedIBPOints           float64
	UseProduct                int        // 使用产品数
	CooperationPeriod         int        // 合作年限数
	CooperationPeriodDiscount float64    // 使用产品数优惠点数
	UseProductDiscount        float64    // 合作年限优惠点数
	QualitativeDiscount       float64    // 定性优惠点数
	Discount                  float64    `xorm:"NUMBER(32,6) 'DISCOUNT'"`
	BaseRate                  float64    `xorm:"NUMBER(32,6) 'BASE_RATE'"` // 基准利率
	BottomRate                float64    `xorm:"NUMBER(32,6) 'BOTTOM_RATE'"`
	SceneRate                 float64    `xorm:"NUMBER(32,6) 'SCENE_RATE'"`
	TgtRate                   float64    `xorm:"NUMBER(32,6) 'TGT_RATE'"`
	IntRate                   float64    `xorm:"NUMBER(32,6) 'INT_RATE'"`   // 执行利率
	MarginType                string     `xorm:"VARCHAR2(1) 'MARGIN_TYPE'"` // 浮动类型
	MarginInt                 float64    `xorm:"NUMBER(32,6) 'MARGIN_INT'"` // 浮动制
	MarginBottom              float64    `xorm:"NUMBER(32,6) 'MARGIN_BOTTOM'"`
	MarginScene               float64    `xorm:"NUMBER(32,6) 'MARGIN_SCENE'"`
	MarginTgt                 float64    `xorm:"NUMBER(32,6) 'MARGIN_TGT'"`
	BottomCon                 float64    `xorm:"NUMBER(32,6) 'BOTTOM_CON'"`
	BottomEva                 float64    `xorm:"NUMBER(32,6) 'BOTTOM_EVA'"`
	BottomRaroc               float64    `xorm:"NUMBER(32,6) 'BOTTOM_RAROC'"`
	TgtCon                    float64    `xorm:"NUMBER(32,6) 'TGT_CON'"`
	TgtEva                    float64    `xorm:"NUMBER(32,6) 'TGT_EVA'"`
	TgtRaroc                  float64    `xorm:"NUMBER(32,6) 'TGT_RAROC'"`
	IntCon                    float64    `xorm:"NUMBER(32,6) 'INT_CON'"`
	IntEva                    float64    `xorm:"NUMBER(32,6) 'INT_EVA'"`
	IntRaroc                  float64    `xorm:"NUMBER(32,6) 'INT_RAROC'"`
	ExchgRate                 float64    `xorm:"NUMBER(32,6) 'EXCHG_RATE'"`
	OneLnNetProfit            float64    //单笔贷款税后净利润
	OneLnYearEva              float64    //单笔贷款年经济增加值
	OneLnRaroc                float64    //单笔贷款RAROC
	SceneNetPorfit            float64    //派生业务净利润
	SceneYearEva              float64    //派生业务年经济增加值
	SumNetProfit              float64    //综合业务贡献
	SumEva                    float64    //综合业务年经济增加值
	SumRaroc                  float64    //综合业务RAROC
	Remark                    string     `xorm:"VARCHAR2(100)   'REMARK'"`
	Flag                      string     `xorm:"VARCHAR2(1)   'FLAG'"`
	Status                    string     `xorm:"VARCHAR2(2)   'STATUS'"`
	CreateTime                time.Time  `xorm:"DATE created  'CREATE_TIME'"`
	CreateUser                string     `xorm:"VARCHAR2(44)   'CREATE_USER'"`
	UpdateTime                time.Time  `xorm:"DATE updated  'UPDATE_TIME'"`
	UpdateUser                string     `xorm:"VARCHAR2(44)   'UPDATE_USER'"`
	ResChr1                   string     `xorm:"VARCHAR2(44)   'RES_CHR1'"`
	ResChr2                   string     `xorm:"VARCHAR2(44)   'RES_CHR2'"`
	ResChr3                   string     `xorm:"VARCHAR2(44)   'RES_CHR3'"`
	ResChr4                   string     `xorm:"VARCHAR2(44)   'RES_CHR4'"`
	ResChr5                   string     `xorm:"VARCHAR2(44)   'RES_CHR5'"`
	ResChr6                   string     `xorm:"VARCHAR2(44)   'RES_CHR6'"`
	ResChr7                   string     `xorm:"VARCHAR2(44)   'RES_CHR7'"`
	ResChr8                   string     `xorm:"VARCHAR2(44)   'RES_CHR8'"`
	ResChr9                   string     `xorm:"VARCHAR2(44)   'RES_CHR9'"`
	ResChr10                  string     `xorm:"VARCHAR2(44)   'RES_CHR10'"`
	ResNum1                   float64    `xorm:"NUMBER(32,6) 'RES_NUM1'"`
	ResNum2                   float64    `xorm:"NUMBER(32,6) 'RES_NUM2'"`
	ResNum3                   float64    `xorm:"NUMBER(32,6) 'RES_NUM3'"`
	ResNum4                   float64    `xorm:"NUMBER(32,6) 'RES_NUM4'"`
	ResNum5                   float64    `xorm:"NUMBER(32,6) 'RES_NUM5'"`
	ResNum6                   float64    `xorm:"NUMBER(32,6) 'RES_NUM6'"`
	ResNum7                   float64    `xorm:"NUMBER(32,6) 'RES_NUM7'"`
	ResNum8                   float64    `xorm:"NUMBER(32,6) 'RES_NUM8'"`
	ResNum9                   float64    `xorm:"NUMBER(32,6) 'RES_NUM9'"`
	ResNum10                  float64    `xorm:"NUMBER(32,6) 'RES_NUM10'"`
	LnBusiness                LnBusiness `xorm:"-"`
	RowNumber                 float64
}

//LnPricing type 对公贷款定价单实体
// type NullLnPricing struct {
// 	UUID            sql.NullString  `xorm:"varchar(44) pk notnull unique 'UUID'"`
// 	BusinessCode    sql.NullString  `xorm:"varchar(44)   'BUSINESS_CODE'"`
// 	ContractCode    sql.NullString  `xorm:"varchar(44)   'CONTRACT_CODE'"`
// 	PlnCode         sql.NullString  `xorm:"varchar(44)   'PLN_CODE'"`
// 	CustCode        sql.NullString  `xorm:"varchar(44)   'CUST_CODE'"`
// 	CustName        sql.NullString  `xorm:"varchar(44)   'CUST_NAME'"`
// 	CustType        sql.NullString  `xorm:"varchar(44)   'CUST_TYPE'"`
// 	CustImplvl      sql.NullString  `xorm:"varchar(44)   'CUST_IMPLVL'"`
// 	CustCredit      sql.NullString  `xorm:"varchar(44)   'CUST_CREDIT'"`
// 	BranchCode      sql.NullString  `xorm:"varchar(44)   'BRANCH_CODE'"`
// 	BranchName      sql.NullString  `xorm:"varchar(44)   'BRANCH_NAME'"`
// 	IndustryCode    sql.NullString  `xorm:"varchar(44)   'INDUSTRY_NAME'"`
// 	IndustryName    sql.NullString  `xorm:"varchar(44)   'INDUSTRY_NAME'"`
// 	FtpRate         sql.NullFloat64 `xorm:"NUMBER(32,6) 'FTP_RATE'"`
// 	OcRate          sql.NullFloat64 `xorm:"NUMBER(32,6) 'OC_RATE'"`
// 	PdRate          sql.NullFloat64 `xorm:"NUMBER(32,6) 'PD_RATE'"`
// 	LgdRate         sql.NullFloat64 `xorm:"NUMBER(32,6) 'LGD_RATE'"`
// 	ElRate          sql.NullFloat64 `xorm:"NUMBER(32,6) 'EL_RATE'"`
// 	EcRate          sql.NullFloat64 `xorm:"NUMBER(32,6) 'EC_RATE'"`
// 	CapCostRate     sql.NullFloat64 `xorm:"NUMBER(32,6) 'CAP_COST_RATE'"`
// 	CapPftRate      sql.NullFloat64 `xorm:"NUMBER(32,6) 'CAP_PFT_RATE'"`
// 	IncomeTax       sql.NullFloat64 `xorm:"NUMBER(32,6) 'INCOME_TAX'"`
// 	SalesTax        sql.NullFloat64 `xorm:"NUMBER(32,6) 'SALES_TAX'"`
// 	AddTax          sql.NullFloat64 `xorm:"NUMBER(32,6) 'ADD_TAX'"`
// 	StockUsage      sql.NullFloat64
// 	StockPoints     sql.NullFloat64
// 	DerviedDPEVE    sql.NullFloat64
// 	DerviedDPPoints sql.NullFloat64
// 	DerviedIBEVA    sql.NullFloat64
// 	DerviedIBPOints sql.NullFloat64
// 	Discount        sql.NullFloat64 `xorm:"NUMBER(32,6) 'DISCOUNT'"`
// 	BaseRate        sql.NullFloat64 `xorm:"NUMBER(32,6) 'BASE_RATE'"`
// 	BottomRate      sql.NullFloat64 `xorm:"NUMBER(32,6) 'BOTTOM_RATE'"`
// 	SceneRate       sql.NullFloat64 `xorm:"NUMBER(32,6) 'SCENE_RATE'"`
// 	TgtRate         sql.NullFloat64 `xorm:"NUMBER(32,6) 'TGT_RATE'"`
// 	IntRate         sql.NullFloat64 `xorm:"NUMBER(32,6) 'INT_RATE'"`
// 	MarginType      sql.NullString  `xorm:"VARCHAR2(1) 'MARGIN_TYPE'"`
// 	MarginInt       sql.NullFloat64 `xorm:"NUMBER(32,6) 'MARGIN_INT'"`
// 	MarginBottom    sql.NullFloat64 `xorm:"NUMBER(32,6) 'MARGIN_BOTTOM'"`
// 	MarginScene     sql.NullFloat64 `xorm:"NUMBER(32,6) 'MARGIN_SCENE'"`
// 	MarginTgt       sql.NullFloat64 `xorm:"NUMBER(32,6) 'MARGIN_TGT'"`
// 	BottomCon       sql.NullFloat64 `xorm:"NUMBER(32,6) 'BOTTOM_CON'"`
// 	BottomEva       sql.NullFloat64 `xorm:"NUMBER(32,6) 'BOTTOM_EVA'"`
// 	BottomRaroc     sql.NullFloat64 `xorm:"NUMBER(32,6) 'BOTTOM_RAROC'"`
// 	TgtCon          sql.NullFloat64 `xorm:"NUMBER(32,6) 'TGT_CON'"`
// 	TgtEva          sql.NullFloat64 `xorm:"NUMBER(32,6) 'TGT_EVA'"`
// 	TgtRaroc        sql.NullFloat64 `xorm:"NUMBER(32,6) 'TGT_RAROC'"`
// 	IntCon          sql.NullFloat64 `xorm:"NUMBER(32,6) 'INT_CON'"`
// 	IntEva          sql.NullFloat64 `xorm:"NUMBER(32,6) 'INT_EVA'"`
// 	IntRaroc        sql.NullFloat64 `xorm:"NUMBER(32,6) 'INT_RAROC'"`
// 	ExchgRate       sql.NullFloat64 `xorm:"NUMBER(32,6) 'EXCHG_RATE'"`
// 	Remark          sql.NullString  `xorm:"VARCHAR2(100)   'REMARK'"`
// 	Flag            sql.NullString  `xorm:"VARCHAR2(1)   'FLAG'"`
// 	Status          sql.NullString  `xorm:"VARCHAR2(2)   'STATUS'"`
// 	CreateTime      time.Time       `xorm:"DATE created  'CREATE_TIME'"`
// 	CreateUser      sql.NullString  `xorm:"VARCHAR2(44)   'CREATE_USER'"`
// 	UpdateTime      time.Time       `xorm:"DATE updated  'UPDATE_TIME'"`
// 	UpdateUser      sql.NullString  `xorm:"VARCHAR2(44)   'UPDATE_USER'"`
// 	ResChr1         sql.NullString  `xorm:"VARCHAR2(44)   'RES_CHR1'"`
// 	ResChr2         sql.NullString  `xorm:"VARCHAR2(44)   'RES_CHR2'"`
// 	ResChr3         sql.NullString  `xorm:"VARCHAR2(44)   'RES_CHR3'"`
// 	ResChr4         sql.NullString  `xorm:"VARCHAR2(44)   'RES_CHR4'"`
// 	ResChr5         sql.NullString  `xorm:"VARCHAR2(44)   'RES_CHR5'"`
// 	ResChr6         sql.NullString  `xorm:"VARCHAR2(44)   'RES_CHR6'"`
// 	ResChr7         sql.NullString  `xorm:"VARCHAR2(44)   'RES_CHR7'"`
// 	ResChr8         sql.NullString  `xorm:"VARCHAR2(44)   'RES_CHR8'"`
// 	ResChr9         sql.NullString  `xorm:"VARCHAR2(44)   'RES_CHR9'"`
// 	ResChr10        sql.NullString  `xorm:"VARCHAR2(44)   'RES_CHR10'"`
// 	ResNum1         sql.NullFloat64 `xorm:"NUMBER(32,6) 'RES_NUM1'"`
// 	ResNum2         sql.NullFloat64 `xorm:"NUMBER(32,6) 'RES_NUM2'"`
// 	ResNum3         sql.NullFloat64 `xorm:"NUMBER(32,6) 'RES_NUM3'"`
// 	ResNum4         sql.NullFloat64 `xorm:"NUMBER(32,6) 'RES_NUM4'"`
// 	ResNum5         sql.NullFloat64 `xorm:"NUMBER(32,6) 'RES_NUM5'"`
// 	ResNum6         sql.NullFloat64 `xorm:"NUMBER(32,6) 'RES_NUM6'"`
// 	ResNum7         sql.NullFloat64 `xorm:"NUMBER(32,6) 'RES_NUM7'"`
// 	ResNum8         sql.NullFloat64 `xorm:"NUMBER(32,6) 'RES_NUM8'"`
// 	ResNum9         sql.NullFloat64 `xorm:"NUMBER(32,6) 'RES_NUM9'"`
// 	ResNum10        sql.NullFloat64 `xorm:"NUMBER(32,6) 'RES_NUM10'"`
// 	LnBusiness      NullLnBusiness  `xorm:"-"`
//  	RowNumber       sql.NullFloat64,
//  }

//QueryPricingListByPageAndOrder func 查询所有的定价单
func (LnPricingModel) QueryPricingListByPageAndOrderWithCustCodeName(startRowNumber int64, pageSize int64, orderAttr string, orderType util.OrderType, parms map[string]interface{}) ([]LnPricing, util.Page, error) {
	var page util.Page
	var lnps []LnPricing
	//var args []interface{}
	var orderSQLStr string
	var whereSQLStr string
	var pageSQL string
	// var querySQL = "select * from RPM_BIZ_LN_COLTD_PRICING THIS_ left join RPM_BIZ_LN_COLTD_BUSINESS LnBusiness on THIS_.BUSINESS_CODE = LnBusiness.BUSINESS_CODE LEFT JOIN SYS_SEC_ORGAN Organ ON Organ.ORGAN_CODE = LnBusiness.ORGAN LEFT JOIN RPM_DIM_PRODUCT PRODUCT ON PRODUCT.PRODUCT_CODE=LnBusiness.PRODUCT where THIS_.FLAG!='0'"

	var querySQL = `SELECT LNBUSINESS.BUSINESS_CODE,CUST.CUST_CODE,CUST.CUST_NAME,ORGAN.ORGAN_NAME,PRODUCT.PRODUCT_NAME,LNBUSINESS.CURRENCY,LNBUSINESS.TERM,LNBUSINESS.TERM_MULT,coalesce(THIS_.INT_RATE,0),LNBUSINESS.RATE_TYPE,LNBUSINESS.RPYM_TYPE,LNBUSINESS.REPRICE_FREQ,LNBUSINESS.PRINCIPAL,coalesce(this_.STATUS,'-1')
FROM RPM_BIZ_LN_COLTD_BUSINESS LNBUSINESS
LEFT JOIN RPM_BIZ_LN_COLTD_PRICING THIS_ 
ON LNBUSINESS.BUSINESS_CODE=THIS_.BUSINESS_CODE
LEFT JOIN RPM_BIZ_CUST_INFO CUST ON LNBUSINESS.CUST=CUST.CUST_CODE
LEFT JOIN SYS_SEC_ORGAN ORGAN ON ORGAN.ORGAN_CODE=LNBUSINESS.ORGAN
LEFT JOIN RPM_DIM_PRODUCT PRODUCT ON PRODUCT.PRODUCT_CODE=LNBUSINESS.PRODUCT
 `

	if orderAttr != "" {
		orderSQLStr = "ORDER BY "
		if strings.Contains(orderAttr, ".") {
			orderAttrs := strings.Split(orderAttr, ".")
			if 3 == len(orderAttrs) {
				orderSQLStr = orderSQLStr + orderAttrs[1] + "." + util.UperChange(orderAttrs[2]) + " "
			} else if 2 == len(orderAttrs) {
				orderSQLStr = orderSQLStr + orderAttrs[0] + "." + util.UperChange(orderAttrs[1]) + " "
			}

		} else {
			orderSQLStr = orderSQLStr + "THIS_." + util.UperChange(orderAttr) + " "
		}
		// if orderType == util.ASC {
		// 	orderSQLStr = orderSQLStr + " ASC "
		// } else if orderType == util.DESC {
		// 	orderSQLStr = orderSQLStr + " DESC "
		// }
		if orderType == util.ASC {
			orderSQLStr += string(orderType) + " NULLS FIRST "
		} else {
			orderSQLStr += string(orderType) + " NULLS LAST "
		}

	}

	if len(parms) != 0 {

		for k, v := range parms {

			cols := queryCodtion[k]
			if "" != cols {
				if "" != whereSQLStr {
					whereSQLStr += " and "
				}
				if "status" == k && "-1" == v {
					whereSQLStr = whereSQLStr + cols + " is null"
				} else if "searchLike" == k {
					for _, sv := range v.([]map[string]interface{}) {
						switch sv["type"] {
						case "in":
							whereSQLStr += sv["key"].(string) + " in (" + sv["value"].(string) + ") "
						default:
							er := fmt.Errorf("贷款定价单数据查询不支持【%s】查询", sv["type"].(string))
							zlog.Error(er.Error(), er)
						}

					}
				} else {
					//args = append(args, v)
					whereSQLStr = whereSQLStr + cols + " Like '%" + v.(string) + "%'"
				}
			} else { //不支持的查询
				return lnps, page, errors.New("暂不支持[" + k + "]的查询,只支持按照[BusinessCode/CustCode/CustName]查询")
			}

		}
	}
	if "" != whereSQLStr {
		whereSQLStr = " WHERE (" + whereSQLStr + ")"
	}

	querySQL = querySQL + whereSQLStr + orderSQLStr

	page = util.GetPage(startRowNumber, pageSize, querySQL)
	pageSQL = "SELECT PAGE_2.* FROM (SELECT PAGE_1.*,ROWNUM AS rowno FROM ( " + querySQL + ") PAGE_1  WHERE ROWNUM <=:0) PAGE_2 WHERE PAGE_2.rowno >:1"
	zlog.Debugf("SQL:%s\n[=============%#v]", nil, pageSQL, parms)
	//fmt.Println("开始分也行数,结束分页行数", page.EndRowNumber, page.StartRowNumber)
	rows, err := dbobj.Default.Query(pageSQL, page.EndRowNumber, page.StartRowNumber)
	if err == nil {
		for rows.Next() {
			var lnp LnPricing
			rows.Scan(&lnp.BusinessCode, &lnp.LnBusiness.Cust.CustCode, &lnp.LnBusiness.Cust.CustName, &lnp.LnBusiness.Organ.OrganName, &lnp.LnBusiness.Product.ProductName, &lnp.LnBusiness.Currency, &lnp.LnBusiness.Term, &lnp.LnBusiness.TermMult, &lnp.IntRate, &lnp.LnBusiness.RateType, &lnp.LnBusiness.RpymType, &lnp.LnBusiness.RepriceFreq, &lnp.LnBusiness.Principal, &lnp.Status, &lnp.RowNumber)
			lnps = append(lnps, lnp)
		}
		return lnps, page, nil
	}
	return lnps, page, err

}

//QueryPricingListWithPID func 查询单一定价单
func (LnPricingModel) QueryPricingListWithPID(pid string) ([]LnPricing, error) {
	// var querySQL = "select THIS_.*,LnBusiness.*,Organ.*,PRODUCT.*,ROWNUM from RPM_BIZ_LN_COLTD_PRICING THIS_ left join RPM_BIZ_LN_COLTD_BUSINESS LnBusiness on THIS_.BUSINESS_CODE = LnBusiness.BUSINESS_CODE LEFT JOIN SYS_SEC_ORGAN Organ ON Organ.ORGAN_CODE = LnBusiness.ORGAN LEFT JOIN RPM_DIM_PRODUCT PRODUCT ON PRODUCT.PRODUCT_CODE=LnBusiness.PRODUCT LEFT JOIN RPM_BIZ_CUST_INFO where THIS_.FLAG!='0' AND THIS_.BUSINESS_CODE=:0"
	var querySQL = "select " + selectcols + ",ROWNUM from RPM_BIZ_LN_COLTD_PRICING THIS_  where THIS_.BUSINESS_CODE=:0"

	zlog.AppOperateLog("", "PricingListService.LnPringList", zlog.SELECT, nil, nil, "查询对公贷款定价单数据")
	rows, err := dbobj.Default.Query(querySQL, pid)
	lnp := getResult(rows)
	return lnp, err
}

func (LnPricingModel) Delete(nlnp LnPricing) error {

	err := dbobj.Default.Exec("DELETE FROM RPM_BIZ_LN_COLTD_BUSINESS THIS_ WHERE THIS_.BUSINESS_CODE=:0", nlnp.BusinessCode)
	return err
}

func getResult(rows *sql.Rows) []LnPricing {
	var lnps []LnPricing

	defer rows.Close()

	for rows.Next() {
		var lnp LnPricing

		rows.Scan(&lnp.UUID,
			&lnp.BusinessCode,
			&lnp.ContractCode,
			&lnp.PlnCode,
			&lnp.CustCode,
			&lnp.CustName,
			&lnp.CustType,
			&lnp.CustImplvl,
			&lnp.CustCredit,
			&lnp.BranchCode,
			&lnp.BranchName,
			&lnp.IndustryCode,
			&lnp.IndustryName,
			&lnp.FtpRate,
			&lnp.OcRate,
			&lnp.PdRate,
			&lnp.LgdRate,
			&lnp.ElRate,
			&lnp.EcRate,
			&lnp.CapCostRate,
			&lnp.CapPftRate,
			&lnp.IncomeTax,
			&lnp.SalesTax,
			&lnp.AddTax,
			&lnp.StockUsage,
			&lnp.StockPoints,
			&lnp.DerviedDPEVA,
			&lnp.DerviedDPPoints,
			&lnp.DerviedIBEVA,
			&lnp.DerviedIBPOints,
			&lnp.UseProduct,
			&lnp.CooperationPeriod,
			&lnp.CooperationPeriodDiscount,
			&lnp.UseProductDiscount,
			&lnp.QualitativeDiscount,
			&lnp.Discount,
			&lnp.BaseRate,
			&lnp.BottomRate,
			&lnp.SceneRate,
			&lnp.TgtRate,
			&lnp.IntRate,
			&lnp.MarginType,
			&lnp.MarginInt,
			&lnp.MarginBottom,
			&lnp.MarginScene,
			&lnp.MarginTgt,
			&lnp.BottomCon,
			&lnp.BottomEva,
			&lnp.BottomRaroc,
			&lnp.TgtCon,
			&lnp.TgtEva,
			&lnp.TgtRaroc,
			&lnp.IntCon,
			&lnp.IntEva,
			&lnp.IntRaroc,
			&lnp.ExchgRate,

			&lnp.OneLnNetProfit,
			&lnp.OneLnYearEva,
			&lnp.OneLnRaroc,
			&lnp.SceneNetPorfit,
			&lnp.SceneYearEva,
			&lnp.SumNetProfit,
			&lnp.SumEva,
			&lnp.SumRaroc,
			&lnp.Remark,
			&lnp.Flag,
			&lnp.Status,
			&lnp.CreateTime,
			&lnp.CreateUser,
			&lnp.UpdateTime,
			&lnp.UpdateUser,
			&lnp.ResChr1,
			&lnp.ResChr2,
			&lnp.ResChr3,
			&lnp.ResChr4,
			&lnp.ResChr5,
			&lnp.ResChr6,
			&lnp.ResChr7,
			&lnp.ResChr8,
			&lnp.ResChr9,
			&lnp.ResChr10,
			&lnp.ResNum1,
			&lnp.ResNum2,
			&lnp.ResNum3,
			&lnp.ResNum4,
			&lnp.ResNum5,
			&lnp.ResNum6,
			&lnp.ResNum7,
			&lnp.ResNum8,
			&lnp.ResNum9,
			&lnp.ResNum10,
			// &lnp.LnBusiness.UUID,
			// &lnp.LnBusiness.BusinessCode,
			// &lnp.LnBusiness.Cust.CustCode,
			// &lnp.LnBusiness.Organ.OrganCode,
			// &lnp.LnBusiness.Product.ProductCode,
			// &lnp.LnBusiness.Currency,
			// &lnp.LnBusiness.Term,
			// &lnp.LnBusiness.TermMult,
			// &lnp.LnBusiness.RateType,
			// &lnp.LnBusiness.RpymType,
			// &lnp.LnBusiness.RepriceFreq,
			// &lnp.LnBusiness.RpymInterestFreq,
			// &lnp.LnBusiness.RpymCapitalFreq,
			// &lnp.LnBusiness.Principal,
			// &lnp.LnBusiness.BaseRateType,
			// &lnp.LnBusiness.MainMortgageType,
			// &lnp.LnBusiness.CreateTime,
			// &lnp.LnBusiness.UpdateTime,
			// &lnp.LnBusiness.CreateUser,
			// &lnp.LnBusiness.UpdateUser,
			// &lnp.LnBusiness.Flag,
			// &lnp.LnBusiness.Status,
			// &lnp.LnBusiness.ResChr1,
			// &lnp.LnBusiness.ResChr2,
			// &lnp.LnBusiness.ResChr3,
			// &lnp.LnBusiness.ResChr4,
			// &lnp.LnBusiness.ResChr5,
			// &lnp.LnBusiness.ResChr6,
			// &lnp.LnBusiness.ResChr7,
			// &lnp.LnBusiness.ResChr8,
			// &lnp.LnBusiness.ResChr9,
			// &lnp.LnBusiness.ResChr10,
			// &lnp.LnBusiness.ResNum1,
			// &lnp.LnBusiness.ResNum2,
			// &lnp.LnBusiness.ResNum3,
			// &lnp.LnBusiness.ResNum4,
			// &lnp.LnBusiness.ResNum5,
			// &lnp.LnBusiness.ResNum6,
			// &lnp.LnBusiness.ResNum7,
			// &lnp.LnBusiness.ResNum8,
			// &lnp.LnBusiness.ResNum9,
			// &lnp.LnBusiness.ResNum10,
			// &lnp.LnBusiness.Organ.UUID,
			// &lnp.LnBusiness.Organ.OrganCode,
			// &lnp.LnBusiness.Organ.OrganName,
			// &lnp.LnBusiness.Organ.OrganLevel,
			// &lnp.LnBusiness.Organ.ParentOrgan,
			// &lnp.LnBusiness.Organ.LeafFlag,
			// &lnp.LnBusiness.Organ.Remark,
			// &lnp.LnBusiness.Organ.Flag,
			// &lnp.LnBusiness.Organ.Status,
			// &lnp.LnBusiness.Organ.CreateTime,
			// &lnp.LnBusiness.Organ.CreateUser,
			// &lnp.LnBusiness.Organ.UpdateTime,
			// &lnp.LnBusiness.Organ.UpdateUser,
			// &lnp.LnBusiness.Organ.ResChr1,
			// &lnp.LnBusiness.Organ.ResChr2,
			// &lnp.LnBusiness.Organ.ResChr3,
			// &lnp.LnBusiness.Product.UUID,
			// &lnp.LnBusiness.Product.ProductCode,
			// &lnp.LnBusiness.Product.ProductName,
			// &lnp.LnBusiness.Product.ProductType,
			// &lnp.LnBusiness.Product.ProductTypeDesc,
			// &lnp.LnBusiness.Product.ProductLevel,
			// &lnp.LnBusiness.Product.ParentProduct,
			// &lnp.LnBusiness.Product.CreateTime,
			// &lnp.LnBusiness.Product.CreateUser,
			// &lnp.LnBusiness.Product.UpdateTime,
			// &lnp.LnBusiness.Product.UpdateUser,
			// &lnp.LnBusiness.Product.LeafFlag,
			// &lnp.LnBusiness.Product.Flag,
			&lnp.RowNumber,
		)

		lnps = append(lnps, lnp)
	}
	return lnps
}

// ------------ by Jason ----------

func (l *LnPricing) scan(rows *sql.Rows) (*LnPricing, error) {
	var lnPricing = new(LnPricing)
	values := []interface{}{
		&lnPricing.UUID,
		&lnPricing.BusinessCode,
		&lnPricing.ContractCode,
		&lnPricing.PlnCode,
		&lnPricing.CustCode,
		&lnPricing.CustName,
		&lnPricing.CustType,
		&lnPricing.CustImplvl,
		&lnPricing.CustCredit,
		&lnPricing.BranchCode,
		&lnPricing.BranchName,
		&lnPricing.IndustryCode,
		&lnPricing.IndustryName,
		&lnPricing.FtpRate,
		&lnPricing.OcRate,
		&lnPricing.PdRate,
		&lnPricing.LgdRate,
		&lnPricing.ElRate,
		&lnPricing.EcRate,
		&lnPricing.CapCostRate,
		&lnPricing.CapPftRate,
		&lnPricing.IncomeTax,
		&lnPricing.SalesTax,
		&lnPricing.AddTax,
		&lnPricing.StockUsage,
		&lnPricing.StockPoints,
		&lnPricing.DerviedDPEVA,
		&lnPricing.DerviedDPPoints,
		&lnPricing.DerviedIBEVA,
		&lnPricing.DerviedIBPOints,
		&lnPricing.UseProduct,
		&lnPricing.CooperationPeriod,
		&lnPricing.CooperationPeriodDiscount,
		&lnPricing.UseProductDiscount,
		&lnPricing.QualitativeDiscount,
		&lnPricing.Discount,
		&lnPricing.BaseRate,
		&lnPricing.BottomRate,
		&lnPricing.SceneRate,
		&lnPricing.TgtRate,
		&lnPricing.IntRate,
		&lnPricing.MarginType,
		&lnPricing.MarginInt,
		&lnPricing.MarginBottom,
		&lnPricing.MarginScene,
		&lnPricing.MarginTgt,
		&lnPricing.BottomCon,
		&lnPricing.BottomEva,
		&lnPricing.BottomRaroc,
		&lnPricing.TgtCon,
		&lnPricing.TgtEva,
		&lnPricing.TgtRaroc,
		&lnPricing.IntCon,
		&lnPricing.IntEva,
		&lnPricing.IntRaroc,
		&lnPricing.ExchgRate,
		&lnPricing.OneLnNetProfit,
		&lnPricing.OneLnYearEva,
		&lnPricing.OneLnRaroc,
		&lnPricing.SceneNetPorfit,
		&lnPricing.SceneYearEva,
		&lnPricing.SumNetProfit,
		&lnPricing.SumEva,
		&lnPricing.SumRaroc,
		&lnPricing.Remark,
		&lnPricing.Flag,
		&lnPricing.Status,
		&lnPricing.CreateTime,
		&lnPricing.CreateUser,
		&lnPricing.UpdateTime,
		&lnPricing.UpdateUser,
	}
	err := util.OracleScan(rows, values)
	return lnPricing, err
}

func (l *LnPricing) Find(param ...map[string]interface{}) ([]*LnPricing, error) {

	var lnbusiness = new(LnBusiness)
	lb, err := lnbusiness.Find(param...)
	if nil != err {
		return nil, err
	}
	if 0 == len(lb) {
		er := fmt.Errorf("未查询到业务编号的业务信息")
		zlog.Error(er.Error(), er)
		return nil, er
	}
	lnbusiness = lb[0]

	rows, err := modelsUtil.FindRows(lnPricingTables, lnPricingCols, lnPricingColsSort, param...)
	defer rows.Close()
	if nil != err {
		er := fmt.Errorf("查询定价单出错")
		zlog.Error(er.Error(), err)
		return nil, er
	}
	var lnPricings []*LnPricing
	for rows.Next() {
		lnpricing, err := l.scan(rows)
		if nil != err {
			er := fmt.Errorf("查询定价单rows.Scan()出错")
			zlog.Error(er.Error(), err)
			return nil, er
		}
		lnpricing.LnBusiness = *lnbusiness
		lnPricings = append(lnPricings, lnpricing)
	}
	if 0 == len(lnPricings) {
		var lnp = new(LnPricing)
		lnp.LnBusiness = *lnbusiness
		lnp.BusinessCode = lnbusiness.BusinessCode
		lnPricings = append(lnPricings, lnp)
	}
	// fmt.Println("==========", lnPricings)
	return lnPricings, nil
}

// SavePricing func LnPricing 保存定价单
// author Jason
func (l *LnPricing) Add() error {
	param := map[string]interface{}{
		"business_code":               l.BusinessCode,
		"contract_code":               l.ContractCode,
		"pln_code":                    l.PlnCode,
		"cust_code":                   l.CustCode,
		"cust_name":                   l.CustName,
		"cust_type":                   l.CustType,
		"cust_implvl":                 l.CustImplvl,
		"cust_credit":                 l.CustCredit,
		"branch_code":                 l.BranchCode,
		"branch_name":                 l.BranchName,
		"industry_code":               l.IndustryCode,
		"industry_name":               l.IndustryName,
		"ftp_rate":                    l.FtpRate,
		"oc_rate":                     l.OcRate,
		"pd_rate":                     l.PdRate,
		"lgd_rate":                    l.LgdRate,
		"el_rate":                     l.ElRate,
		"ec_rate":                     l.EcRate,
		"cap_cost_rate":               l.CapCostRate,
		"cap_pft_rate":                l.CapPftRate,
		"income_tax":                  l.IncomeTax,
		"sales_tax":                   l.SalesTax,
		"add_tax":                     l.AddTax,
		"stock_usage":                 l.StockUsage,
		"stock_points":                l.StockPoints,
		"dervied_dp_eva":              l.DerviedDPEVA,
		"dervied_dp_points":           l.DerviedDPPoints,
		"dervied_ib_eva":              l.DerviedIBEVA,
		"dervied_ib_points":           l.DerviedIBPOints,
		"use_product":                 l.UseProduct,
		"cooperation_period":          l.CooperationPeriod,
		"cooperation_period_discount": l.CooperationPeriodDiscount,
		"use_product_discount":        l.UseProductDiscount,
		"qualitative_discount":        l.QualitativeDiscount,
		"discount":                    l.Discount,
		"base_rate":                   l.BaseRate,
		"bottom_rate":                 l.BottomRate,
		"scene_rate":                  l.SceneRate,
		"tgt_rate":                    l.TgtRate,
		"int_rate":                    l.IntRate,
		"margin_type":                 l.MarginType,
		"margin_int":                  l.MarginInt,
		"margin_bottom":               l.MarginBottom,
		"margin_scene":                l.MarginScene,
		"margin_tgt":                  l.MarginTgt,
		"bottom_con":                  l.BottomCon,
		"bottom_eva":                  l.BottomEva,
		"bottom_raroc":                l.BottomRaroc,
		"tgt_con":                     l.TgtCon,
		"tgt_eva":                     l.TgtEva,
		"tgt_raroc":                   l.TgtRaroc,
		"int_con":                     l.IntCon,
		"int_eva":                     l.IntEva,
		"int_raroc":                   l.IntRaroc,
		"exchg_rate":                  l.ExchgRate,
		"one_ln_net_profit":           l.OneLnNetProfit,
		"one_ln_year_eva":             l.OneLnYearEva,
		"one_ln_raroc":                l.OneLnRaroc,
		"scene_net_porfit":            l.SceneNetPorfit,
		"scene_year_eva":              l.SceneYearEva,
		"sum_net_profit":              l.SumNetProfit,
		"sum_eva":                     l.SumEva,
		"sum_raroc":                   l.SumRaroc,
		"remark":                      l.Remark,
		"flag":                        l.Flag,
		"status":                      l.Status,
		"create_time":                 util.GetCurrentTime(),
		"create_user":                 l.CreateUser,
		"update_time":                 util.GetCurrentTime(),
		"update_user":                 l.UpdateUser,
	}
	err := util.OracleAdd(lnPricingTables, param)
	if nil != err {
		er := fmt.Errorf("新增定价单出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

func (l *LnPricing) Update() error {
	param := map[string]interface{}{
		"business_code":               l.BusinessCode,
		"contract_code":               l.ContractCode,
		"pln_code":                    l.PlnCode,
		"cust_code":                   l.CustCode,
		"cust_name":                   l.CustName,
		"cust_type":                   l.CustType,
		"cust_implvl":                 l.CustImplvl,
		"cust_credit":                 l.CustCredit,
		"branch_code":                 l.BranchCode,
		"branch_name":                 l.BranchName,
		"industry_code":               l.IndustryCode,
		"industry_name":               l.IndustryName,
		"ftp_rate":                    l.FtpRate,
		"oc_rate":                     l.OcRate,
		"pd_rate":                     l.PdRate,
		"lgd_rate":                    l.LgdRate,
		"el_rate":                     l.ElRate,
		"ec_rate":                     l.EcRate,
		"cap_cost_rate":               l.CapCostRate,
		"cap_pft_rate":                l.CapPftRate,
		"income_tax":                  l.IncomeTax,
		"sales_tax":                   l.SalesTax,
		"add_tax":                     l.AddTax,
		"stock_usage":                 l.StockUsage,
		"stock_points":                l.StockPoints,
		"dervied_dp_eva":              l.DerviedDPEVA,
		"dervied_dp_points":           l.DerviedDPPoints,
		"dervied_ib_eva":              l.DerviedIBEVA,
		"dervied_ib_points":           l.DerviedIBPOints,
		"use_product":                 l.UseProduct,
		"cooperation_period":          l.CooperationPeriod,
		"cooperation_period_discount": l.CooperationPeriodDiscount,
		"use_product_discount":        l.UseProductDiscount,
		"qualitative_discount":        l.QualitativeDiscount,
		"discount":                    l.Discount,
		"base_rate":                   l.BaseRate,
		"bottom_rate":                 l.BottomRate,
		"scene_rate":                  l.SceneRate,
		"tgt_rate":                    l.TgtRate,
		"int_rate":                    l.IntRate,
		"margin_type":                 l.MarginType,
		"margin_int":                  l.MarginInt,
		"margin_bottom":               l.MarginBottom,
		"margin_scene":                l.MarginScene,
		"margin_tgt":                  l.MarginTgt,
		"bottom_con":                  l.BottomCon,
		"bottom_eva":                  l.BottomEva,
		"bottom_raroc":                l.BottomRaroc,
		"tgt_con":                     l.TgtCon,
		"tgt_eva":                     l.TgtEva,
		"tgt_raroc":                   l.TgtRaroc,
		"int_con":                     l.IntCon,
		"int_eva":                     l.IntEva,
		"int_raroc":                   l.IntRaroc,
		"exchg_rate":                  l.ExchgRate,
		"one_ln_net_profit":           l.OneLnNetProfit,
		"one_ln_year_eva":             l.OneLnYearEva,
		"one_ln_raroc":                l.OneLnRaroc,
		"scene_net_porfit":            l.SceneNetPorfit,
		"scene_year_eva":              l.SceneYearEva,
		"sum_net_profit":              l.SumNetProfit,
		"sum_eva":                     l.SumEva,
		"sum_raroc":                   l.SumRaroc,
		"remark":                      l.Remark,
		"flag":                        l.Flag,
		"status":                      l.Status,
		"update_time":                 util.GetCurrentTime(),
		"update_user":                 l.UpdateUser,
	}
	whereParam := map[string]interface{}{
		"business_code": l.BusinessCode,
	}
	err := util.OracleUpdate(lnPricingTables, param, whereParam)
	if nil != err {
		er := fmt.Errorf("更新定价单出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

func (l LnPricingModel) Patch(paramMap, PKMap map[string]interface{}) error {
	whereParam := map[string]interface{}{
		"business_code": PKMap["business_code"],
	}
	err := util.OracleUpdate(lnPricingTables, paramMap, whereParam)
	if nil != err {
		er := fmt.Errorf("更新定价单出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

func (l *LnPricing) Delete() error {
	whereParam := map[string]interface{}{
		"uuid": l.UUID,
	}
	err := util.OracleDelete(lnPricingTables, whereParam)
	if nil != err {
		er := fmt.Errorf("删除定价单出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

func (l *LnPricing) DeleteByBusinessCode(businessCode string) error {
	whereParam := map[string]interface{}{"business_code": l.BusinessCode}
	err := util.OracleDelete(lnPricingTables, whereParam)
	if nil != err {
		er := fmt.Errorf("业务单号删除定价单出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

var lnPricingTables string = "RPM_BIZ_LN_COLTD_PRICING T"

var lnPricingCols = map[string]string{
	"T.UUID":                        "' '",
	"T.BUSINESS_CODE":               "' '",
	"T.CONTRACT_CODE":               "' '",
	"T.PLN_CODE":                    "' '",
	"T.CUST_CODE":                   "' '",
	"T.CUST_NAME":                   "' '",
	"T.CUST_TYPE":                   "' '",
	"T.CUST_IMPLVL":                 "' '",
	"T.CUST_CREDIT":                 "' '",
	"T.BRANCH_CODE":                 "' '",
	"T.BRANCH_NAME":                 "' '",
	"T.INDUSTRY_CODE":               "' '",
	"T.INDUSTRY_NAME":               "' '",
	"T.FTP_RATE":                    "0",
	"T.OC_RATE":                     "0",
	"T.PD_RATE":                     "0",
	"T.LGD_RATE":                    "0",
	"T.EL_RATE":                     "0",
	"T.EC_RATE":                     "0",
	"T.CAP_COST_RATE":               "0",
	"T.CAP_PFT_RATE":                "0",
	"T.INCOME_TAX":                  "0",
	"T.SALES_TAX":                   "0",
	"T.ADD_TAX":                     "0",
	"T.STOCK_USAGE":                 "0",
	"T.STOCK_POINTS":                "0",
	"T.DERVIED_DP_EVA":              "0",
	"T.DERVIED_DP_POINTS":           "0",
	"T.DERVIED_IB_EVA":              "0",
	"T.DERVIED_IB_POINTS":           "0",
	"T.USE_PRODUCT":                 "0",
	"T.COOPERATION_PERIOD":          "0",
	"T.COOPERATION_PERIOD_DISCOUNT": "0",
	"T.USE_PRODUCT_DISCOUNT":        "0",
	"T.QUALITATIVE_DISCOUNT":        "0",
	"T.DISCOUNT":                    "0",
	"T.BASE_RATE":                   "0",
	"T.BOTTOM_RATE":                 "0",
	"T.SCENE_RATE":                  "0",
	"T.TGT_RATE":                    "0",
	"T.INT_RATE":                    "0",
	"T.MARGIN_TYPE":                 "' '",
	"T.MARGIN_INT":                  "0",
	"T.MARGIN_BOTTOM":               "0",
	"T.MARGIN_SCENE":                "0",
	"T.MARGIN_TGT":                  "0",
	"T.BOTTOM_CON":                  "0",
	"T.BOTTOM_EVA":                  "0",
	"T.BOTTOM_RAROC":                "0",
	"T.TGT_CON":                     "0",
	"T.TGT_EVA":                     "0",
	"T.TGT_RAROC":                   "0",
	"T.INT_CON":                     "0",
	"T.INT_EVA":                     "0",
	"T.INT_RAROC":                   "0",
	"T.EXCHG_RATE":                  "0",
	"T.ONE_LN_NET_PROFIT":           "0",
	"T.ONE_LN_YEAR_EVA":             "0",
	"T.ONE_LN_RAROC":                "0",
	"T.SCENE_NET_PORFIT":            "0",
	"T.SCENE_YEAR_EVA":              "0",
	"T.SUM_NET_PROFIT":              "0",
	"T.SUM_EVA":                     "0",
	"T.SUM_RAROC":                   "0",
	"T.REMARK":                      "' '",
	"T.FLAG":                        "' '",
	"T.STATUS":                      "' '",
	"T.CREATE_TIME":                 "sysdate",
	"T.CREATE_USER":                 "' '",
	"T.UPDATE_TIME":                 "sysdate",
	"T.UPDATE_USER":                 "' '",
}

var lnPricingColsSort = []string{
	"T.UUID",
	"T.BUSINESS_CODE",
	"T.CONTRACT_CODE",
	"T.PLN_CODE",
	"T.CUST_CODE",
	"T.CUST_NAME",
	"T.CUST_TYPE",
	"T.CUST_IMPLVL",
	"T.CUST_CREDIT",
	"T.BRANCH_CODE",
	"T.BRANCH_NAME",
	"T.INDUSTRY_CODE",
	"T.INDUSTRY_NAME",
	"T.FTP_RATE",
	"T.OC_RATE",
	"T.PD_RATE",
	"T.LGD_RATE",
	"T.EL_RATE",
	"T.EC_RATE",
	"T.CAP_COST_RATE",
	"T.CAP_PFT_RATE",
	"T.INCOME_TAX",
	"T.SALES_TAX",
	"T.ADD_TAX",
	"T.STOCK_USAGE",
	"T.STOCK_POINTS",
	"T.DERVIED_DP_EVA",
	"T.DERVIED_DP_POINTS",
	"T.DERVIED_IB_EVA",
	"T.DERVIED_IB_POINTS",
	"T.USE_PRODUCT",
	"T.COOPERATION_PERIOD",
	"T.COOPERATION_PERIOD_DISCOUNT",
	"T.USE_PRODUCT_DISCOUNT",
	"T.QUALITATIVE_DISCOUNT",
	"T.DISCOUNT",
	"T.BASE_RATE",
	"T.BOTTOM_RATE",
	"T.SCENE_RATE",
	"T.TGT_RATE",
	"T.INT_RATE",
	"T.MARGIN_TYPE",
	"T.MARGIN_INT",
	"T.MARGIN_BOTTOM",
	"T.MARGIN_SCENE",
	"T.MARGIN_TGT",
	"T.BOTTOM_CON",
	"T.BOTTOM_EVA",
	"T.BOTTOM_RAROC",
	"T.TGT_CON",
	"T.TGT_EVA",
	"T.TGT_RAROC",
	"T.INT_CON",
	"T.INT_EVA",
	"T.INT_RAROC",
	"T.EXCHG_RATE",
	"T.ONE_LN_NET_PROFIT",
	"T.ONE_LN_YEAR_EVA",
	"T.ONE_LN_RAROC",
	"T.SCENE_NET_PORFIT",
	"T.SCENE_YEAR_EVA",
	"T.SUM_NET_PROFIT",
	"T.SUM_EVA",
	"T.SUM_RAROC",
	"T.REMARK",
	"T.FLAG",
	"T.STATUS",
	"T.CREATE_TIME",
	"T.CREATE_USER",
	"T.UPDATE_TIME",
	"T.UPDATE_USER",
}

var selectcols = `coalesce(UUID,' '),
coalesce(BUSINESS_CODE,' '),
coalesce(CONTRACT_CODE,' '),
coalesce(PLN_CODE,' '),
coalesce(CUST_CODE,' '),
coalesce(CUST_NAME,' '),
coalesce(CUST_TYPE,' '),
coalesce(CUST_IMPLVL,' '),
coalesce(CUST_CREDIT,' '),
coalesce(BRANCH_CODE,' '),
coalesce(BRANCH_NAME,' '),
coalesce(INDUSTRY_CODE,' '),
coalesce(INDUSTRY_NAME,' '),
coalesce(FTP_RATE,0),
coalesce(OC_RATE,0),
coalesce(PD_RATE,0),
coalesce(LGD_RATE,0),
coalesce(EL_RATE,0),
coalesce(EC_RATE,0),
coalesce(CAP_COST_RATE,0),
coalesce(CAP_PFT_RATE,0),
coalesce(INCOME_TAX,0),
coalesce(SALES_TAX,0),
coalesce(ADD_TAX,0),
coalesce(STOCK_USAGE,0),
coalesce(STOCK_POINTS,0),
coalesce(DERVIED_DP_EVA,0),
coalesce(DERVIED_DP_POINTS,0),
coalesce(DERVIED_IB_EVA,0),
coalesce(DERVIED_IB_POINTS,0),
coalesce(USE_PRODUCT,0),
coalesce(COOPERATION_PERIOD,0),
coalesce(COOPERATION_PERIOD_DISCOUNT,0),
coalesce(USE_PRODUCT_DISCOUNT,0),
coalesce(QUALITATIVE_DISCOUNT,0),
coalesce(DISCOUNT,0),
coalesce(BASE_RATE,0),
coalesce(BOTTOM_RATE,0),
coalesce(SCENE_RATE,0),
coalesce(TGT_RATE,0),
coalesce(INT_RATE,0),
coalesce(MARGIN_TYPE,' '),
coalesce(MARGIN_INT,0),
coalesce(MARGIN_BOTTOM,0),
coalesce(MARGIN_SCENE,0),
coalesce(MARGIN_TGT,0),
coalesce(BOTTOM_CON,0),
coalesce(BOTTOM_EVA,0),
coalesce(BOTTOM_RAROC,0),
coalesce(TGT_CON,0),
coalesce(TGT_EVA,0),
coalesce(TGT_RAROC,0),
coalesce(INT_CON,0),
coalesce(INT_EVA,0),
coalesce(INT_RAROC,0),
coalesce(EXCHG_RATE,0),
coalesce(ONE_LN_NET_PROFIT,0),
coalesce(ONE_LN_YEAR_EVA,0),
coalesce(ONE_LN_RAROC,0),
coalesce(SCENE_NET_PORFIT,0),
coalesce(SCENE_YEAR_EVA,0),
coalesce(SUM_NET_PROFIT,0),
coalesce(SUM_EVA,0),
coalesce(SUM_RAROC,0),
coalesce(REMARK,' '),
coalesce(FLAG,' '),
coalesce(STATUS,' '),
coalesce(CREATE_TIME,sysdate),
coalesce(CREATE_USER,' '),
coalesce(UPDATE_TIME,sysdate),
coalesce(UPDATE_USER,' '),
coalesce(RES_CHR1,' '),
coalesce(RES_CHR2,' '),
coalesce(RES_CHR3,' '),
coalesce(RES_CHR4,' '),
coalesce(RES_CHR5,' '),
coalesce(RES_CHR6,' '),
coalesce(RES_CHR7,' '),
coalesce(RES_CHR8,' '),
coalesce(RES_CHR9,' '),
coalesce(RES_CHR10,' '),
coalesce(RES_NUM1,0),
coalesce(RES_NUM2,0),
coalesce(RES_NUM3,0),
coalesce(RES_NUM4,0),
coalesce(RES_NUM5,0),
coalesce(RES_NUM6,0),
coalesce(RES_NUM7,0),
coalesce(RES_NUM8,0),
coalesce(RES_NUM9,0),
coalesce(RES_NUM10,0)`

var queryCodtion = map[string]string{
	"cust_name":     "CUST.CUST_NAME",
	"cust_code":     "CUST.CUST_CODE",
	"business_code": "LNBUSINESS.BUSINESS_CODE",
	"status":        "this_.STATUS",
	"searchLike":    "searchLike",
	"organ_name":    "ORGAN.ORGAN_NAME",
	"owner":         "cust.owner",
}
