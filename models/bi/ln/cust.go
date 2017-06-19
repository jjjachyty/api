package ln

import (
	"database/sql"
	"fmt"
	"time"

	modelsUtil "pccqcpa.com.cn/app/rpm/api/models/util"
	"pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/components/zlog"
)

// 结构体对应数据库表【RPM_BI_LN_CUST】贷款业务分析客户维度表
// by author Jason
// by time 2016-10-31 15:44:32
type LnCust struct {
	PAsOfDate        time.Time // 数据日期
	OrgCodeA         string    // 一级分行号编码
	OrgNameA         string    // 一级分行号名称
	OrgCodeF         string    // 最末级机构号名称
	OrgNameF         string    // 最末级机构号名称
	CustCode         string    // 整合客户号
	CustCodeCredit   string    // 信贷系统客户号
	CustCodeCore     string    // 核心系统客户号
	CustName         string    // 整合客户名称
	CustType         string    // 客户类型(对公零售)
	CustSizeInt      string    // 客户规模（国标）
	CustSizeLine     string    // 客户规模（行标）
	CustNature       string    // 客户性质（零售）农户/个体工商户/城镇居民
	CustCreditRating string    // 客户信用等级
	CustGradeCrm     string    // 客户CRM评级
	CustGradeRl      string    // 客户等级（零售）
	CustIndustryCode string    // 行业编码
	CustIndustryName string    // 行业名称
}

// 结构体scan
// by author Jason
// by time 2016-10-31 15:44:32
func (this LnCust) scan(rows *sql.Rows) (*LnCust, error) {
	var one = new(LnCust)
	values := []interface{}{
		&one.PAsOfDate,
		&one.OrgCodeA,
		&one.OrgNameA,
		&one.OrgCodeF,
		&one.OrgNameF,
		&one.CustCode,
		&one.CustCodeCredit,
		&one.CustCodeCore,
		&one.CustName,
		&one.CustType,
		&one.CustSizeInt,
		&one.CustSizeLine,
		&one.CustNature,
		&one.CustCreditRating,
		&one.CustGradeCrm,
		&one.CustGradeRl,
		&one.CustIndustryCode,
		&one.CustIndustryName,
	}
	err := util.OracleScan(rows, values)
	return one, err
}

// 分页操作
// by author Jason
// by time 2016-10-31 15:44:32
func (this LnCust) List(param ...map[string]interface{}) (*util.PageData, error) {
	pageData, rows, err := modelsUtil.List(lnCustTabales, lnCustCols, lnCustColsSort, param...)
	defer rows.Close()
	if nil != err {
		er := fmt.Errorf("分页查询贷款业务分析客户维度表信息出错")
		zlog.Error(er.Error(), err)
		return nil, er
	}

	var rst []*LnCust
	for rows.Next() {
		one, err := this.scan(rows)
		if nil != err {
			er := fmt.Errorf("分页查询贷款业务分析客户维度表信息row.Scan()出错")
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
// by time 2016-10-31 15:44:32
func (this LnCust) Find(param ...map[string]interface{}) ([]*LnCust, error) {
	rows, err := modelsUtil.FindRows(lnCustTabales, lnCustCols, lnCustColsSort, param...)
	defer rows.Close()
	if nil != err {
		er := fmt.Errorf("多参数查询贷款业务分析客户维度表信息出错")
		zlog.Error(er.Error(), err)
		return nil, er
	}
	defer rows.Close()
	var rst []*LnCust
	for rows.Next() {
		one, err := this.scan(rows)
		if nil != err {
			er := fmt.Errorf("多参数查询贷款业务分析客户维度表信息row.Scan()出错")
			zlog.Error(er.Error(), err)
			return nil, er
		}
		rst = append(rst, one)
	}
	return rst, nil
}

// 新增贷款业务分析客户维度表信息
// by author Jason
// by time 2016-10-31 15:44:32
func (this LnCust) Add() error {
	paramMap := map[string]interface{}{
		"p_as_of_date":       this.PAsOfDate,
		"org_code_a":         this.OrgCodeA,
		"org_name_a":         this.OrgNameA,
		"org_code_f":         this.OrgCodeF,
		"org_name_f":         this.OrgNameF,
		"cust_code":          this.CustCode,
		"cust_code_credit":   this.CustCodeCredit,
		"cust_code_core":     this.CustCodeCore,
		"cust_name":          this.CustName,
		"cust_type":          this.CustType,
		"cust_size_int":      this.CustSizeInt,
		"cust_size_line":     this.CustSizeLine,
		"cust_nature":        this.CustNature,
		"cust_credit_rating": this.CustCreditRating,
		"cust_grade_crm":     this.CustGradeCrm,
		"cust_grade_rl":      this.CustGradeRl,
		"cust_industry_code": this.CustIndustryCode,
		"cust_industry_name": this.CustIndustryName,
	}
	err := util.OracleAdd("RPM_BI_LN_CUST", paramMap)
	if nil != err {
		er := fmt.Errorf("新增贷款业务分析客户维度表信息出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

// 新增贷款业务分析基准利率降息维度表信息
// by author Jason
// by time 2016-10-31 15:34:54
func (this LnCust) BatchAdd(sqlTx *sql.Tx) (sql.Result, error) {
	paramMap := map[string]interface{}{
		"p_as_of_date":       this.PAsOfDate,
		"org_code_a":         this.OrgCodeA,
		"org_name_a":         this.OrgNameA,
		"org_code_f":         this.OrgCodeF,
		"org_name_f":         this.OrgNameF,
		"cust_code":          this.CustCode,
		"cust_code_credit":   this.CustCodeCredit,
		"cust_code_core":     this.CustCodeCore,
		"cust_name":          this.CustName,
		"cust_type":          this.CustType,
		"cust_size_int":      this.CustSizeInt,
		"cust_size_line":     this.CustSizeLine,
		"cust_nature":        this.CustNature,
		"cust_credit_rating": this.CustCreditRating,
		"cust_grade_crm":     this.CustGradeCrm,
		"cust_grade_rl":      this.CustGradeRl,
		"cust_industry_code": this.CustIndustryCode,
		"cust_industry_name": this.CustIndustryName,
	}
	result, err := util.OracleBatchAdd("RPM_BI_LN_CUST", paramMap, sqlTx)
	return result, err
}

// 更新贷款业务分析客户维度表信息
// by author Jason
// by time 2016-10-31 15:44:32
func (this LnCust) Update() error {
	paramMap := map[string]interface{}{
		"p_as_of_date":       this.PAsOfDate,
		"org_code_a":         this.OrgCodeA,
		"org_name_a":         this.OrgNameA,
		"org_code_f":         this.OrgCodeF,
		"org_name_f":         this.OrgNameF,
		"cust_code":          this.CustCode,
		"cust_code_credit":   this.CustCodeCredit,
		"cust_code_core":     this.CustCodeCore,
		"cust_name":          this.CustName,
		"cust_type":          this.CustType,
		"cust_size_int":      this.CustSizeInt,
		"cust_size_line":     this.CustSizeLine,
		"cust_nature":        this.CustNature,
		"cust_credit_rating": this.CustCreditRating,
		"cust_grade_crm":     this.CustGradeCrm,
		"cust_grade_rl":      this.CustGradeRl,
		"cust_industry_code": this.CustIndustryCode,
		"cust_industry_name": this.CustIndustryName,
	}
	whereParam := map[string]interface{}{
		"cust_code": this.CustCode,
	}
	err := util.OracleUpdate("RPM_BI_LN_CUST", paramMap, whereParam)
	if nil != err {
		er := fmt.Errorf("更新贷款业务分析客户维度表信息出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

// 删除贷款业务分析客户维度表信息
// by author Jason
// by time 2016-10-31 15:44:32
func (this LnCust) Delete() error {
	whereParam := map[string]interface{}{
		"cust_code": this.CustCode,
	}
	err := util.OracleDelete("RPM_BI_LN_CUST", whereParam)
	if nil != err {
		er := fmt.Errorf("删除贷款业务分析客户维度表信息出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

var lnCustTabales string = `RPM_BI_LN_CUST`

var lnCustCols map[string]string = map[string]string{
	"P_AS_OF_DATE":       "sysdate",
	"ORG_CODE_A":         "' '",
	"ORG_NAME_A":         "' '",
	"ORG_CODE_F":         "' '",
	"ORG_NAME_F":         "' '",
	"CUST_CODE":          "' '",
	"CUST_CODE_CREDIT":   "' '",
	"CUST_CODE_CORE":     "' '",
	"CUST_NAME":          "' '",
	"CUST_TYPE":          "' '",
	"CUST_SIZE_INT":      "' '",
	"CUST_SIZE_LINE":     "' '",
	"CUST_NATURE":        "' '",
	"CUST_CREDIT_RATING": "' '",
	"CUST_GRADE_CRM":     "' '",
	"CUST_GRADE_RL":      "' '",
	"CUST_INDUSTRY_CODE": "' '",
	"CUST_INDUSTRY_NAME": "' '",
}

var lnCustColsSort = []string{
	"P_AS_OF_DATE",
	"ORG_CODE_A",
	"ORG_NAME_A",
	"ORG_CODE_F",
	"ORG_NAME_F",
	"CUST_CODE",
	"CUST_CODE_CREDIT",
	"CUST_CODE_CORE",
	"CUST_NAME",
	"CUST_TYPE",
	"CUST_SIZE_INT",
	"CUST_SIZE_LINE",
	"CUST_NATURE",
	"CUST_CREDIT_RATING",
	"CUST_GRADE_CRM",
	"CUST_GRADE_RL",
	"CUST_INDUSTRY_CODE",
	"CUST_INDUSTRY_NAME",
}
