package ln

//贷款业务分析基础数据 Model类
//@auth by janly
import (
	"database/sql"
	"time"

	"platform/dbobj"

	"fmt"

	"pccqcpa.com.cn/components/zlog"
)

type LoanInfo struct {
	SerialNumber         string    //	贷款借据号
	ContractNumber       string    //	合同号
	SubjectCode          string    //	科目号
	SubjectName          string    //	科目名称
	CustCode             string    //	客户号
	CustName             string    //   客户名称
	CustCodeCredit       string    //	信贷客户号
	CustCodeCore         string    //	核心客户号
	ProductCode          string    //	产品号
	ProductName          string    //	产品名称
	LoanType             string    //	贷款类型(对公/零售)
	CustType             string    //	客户类型
	OrgCodeCredit        string    //	信贷机构编码
	OrgNameCredit        string    //	信贷机构名称
	OrgCodeCore          string    //	核心机构代码
	OrgNameCore          string    //	核心机构名称
	OrgCodeMc            string    //	管理会计机构编码
	OrgNameMc            string    //	管理会计机构名称
	Area                 string    //	地区
	IndustryCode         string    //	行业代码
	IndustryName         string    //	行业名称
	Amount               float64   //	贷款金额
	AmountRange          string    //	金额区间
	BaseRate             float64   //	基准利率
	BaseRateMc           float64   //	管快基准利率
	ReferenceAmount      float64   //	基准*金额
	InitRate             float64   //	执行利率
	ExecutiveAmount      float64   //	执行*金额
	ContractStartDate    time.Time //	合同起始日
	ContractStartMonth   int       //	合同起始日（月）
	RateCutDay           time.Time //	降息日
	ContractEndDate      time.Time //	合同到期日
	TermDay              int       //	期限（日）
	TermMonth            int       //	期限（月）
	TermMonthAmount      float64   //	期限（月）*金额
	TermYear             float64   //	期限（年）
	TermInterval         string    //	期限区间
	RateAdjTypeCode      string    //	利率调整类型编码
	RateAdjTypeName      string    //	利率调整类型
	RePricingFreq        int       //	重定价频率
	RePricingFreqUnit    string    //	重定价频率单位
	RePayFreq            int       //	还款频率
	RePayFreqUnit        string    //	还款频率单位
	RePayMethodCode      string    //	还款方式代码
	RePayMethodName      string    //	还款方式
	GuaranteeWayCode     string    //	担保方式编码
	GuaranteeWayName     string    //	担保方式
	MainGuaranteeWayCode string    //	主担保方式编码
	MainGuaranteeWayName string    //	主担保方式
	Balance              float64   //	余额
	CustScaleCodeGb      string    //	客户规模编码（国标）
	CustScaleNameGb      string    //	客户规模名称（国标）
	CustScaleCodeLine    string    //	客户规模（行标）
	CustScaleNameLine    string    //	客户规模（行标）
	CreditRating         string    //	信用评级
	AsOfDate             time.Time //数据日期
}

type LoanInfoModel struct{}

func (LoanInfoModel) Insert(loanInfo LoanInfo, sqlTx *sql.Tx) (sql.Result, error) {

	var truncateSQL = "TRUNCATE TABLE RPM_BI_LN_INFO"
	zlog.Debugf("贷款原始数据删除-SQL:%s", nil, truncateSQL)
	result, err := sqlTx.Exec(truncateSQL)
	zlog.Debugf("贷款原始数据导入-SQL:%s", nil, insertSQL)

	result, err = sqlTx.Exec(insertSQL,
		loanInfo.SerialNumber,
		loanInfo.ContractNumber,
		loanInfo.SubjectCode,
		loanInfo.SubjectName,
		loanInfo.CustCode,
		loanInfo.CustName,
		loanInfo.CustCodeCredit,
		loanInfo.CustCodeCore,
		loanInfo.ProductCode,
		loanInfo.ProductName,
		loanInfo.LoanType,
		loanInfo.CustType,
		loanInfo.OrgCodeCredit,
		loanInfo.OrgNameCredit,
		loanInfo.OrgCodeCore,
		loanInfo.OrgNameCore,
		loanInfo.OrgCodeMc,
		loanInfo.OrgNameMc,
		loanInfo.Area,
		loanInfo.IndustryCode,
		loanInfo.IndustryName,
		loanInfo.Amount,
		loanInfo.AmountRange,
		loanInfo.BaseRate,
		loanInfo.BaseRateMc,
		loanInfo.ReferenceAmount,
		loanInfo.InitRate,
		loanInfo.ExecutiveAmount,
		loanInfo.ContractStartDate,
		loanInfo.ContractStartMonth,
		loanInfo.RateCutDay,
		loanInfo.ContractEndDate,
		loanInfo.TermDay,
		loanInfo.TermMonth,
		loanInfo.TermMonthAmount,
		loanInfo.TermYear,
		loanInfo.TermInterval,
		loanInfo.RateAdjTypeCode,
		loanInfo.RateAdjTypeName,
		loanInfo.RePricingFreq,
		loanInfo.RePricingFreqUnit,
		loanInfo.RePayFreq,
		loanInfo.RePayFreqUnit,
		loanInfo.RePayMethodCode,
		loanInfo.RePayMethodName,
		loanInfo.GuaranteeWayCode,
		loanInfo.GuaranteeWayName,
		loanInfo.MainGuaranteeWayCode,
		loanInfo.MainGuaranteeWayName,
		loanInfo.Balance,
		loanInfo.CustScaleCodeGb,
		loanInfo.CustScaleNameGb,
		loanInfo.CustScaleCodeLine,
		loanInfo.CustScaleNameLine,
		loanInfo.CreditRating,
		loanInfo.AsOfDate)

	return result, err
}

func (LoanInfoModel) BatchInsert(loanInfos []LoanInfo) (sql.Result, error) {
	sqlTx, err := dbobj.Default.Begin()
	if nil != err {
		zlog.Error("[贷款基础信息-RPM_BI_LN_INFO]导入,事物开启失败", nil)
		return nil, err
	}
	//先清除表数据
	stmt, err := sqlTx.Prepare(truncateSQL)
	if nil != err {
		zlog.Error("[贷款基础信息-RPM_BI_LN_INFO]导入,获取Stmt失败", nil)
		return nil, err
	}
	result, err := stmt.Exec()
	zlog.Infof("[贷款基础信息-RPM_BI_LN_INFO]导入,TRUNCATE-SQL:%s", nil, truncateSQL)
	if nil != err {
		zlog.Error("[贷款基础信息-RPM_BI_LN_INFO]导入,TRUNCATE出错", err)
		sqlTx.Rollback()
		return result, err
	}
	//关闭连接
	stmt.Close()
	//插入表数据
	stmt, err = sqlTx.Prepare(insertSQL)
	if nil != err {
		zlog.Error("[贷款基础信息-RPM_BI_LN_INFO]导入,获取Stmt失败", nil)
		return nil, err
	}

	for k, loanInfo := range loanInfos {
		zlog.Debugf("[贷款基础信息-RPM_BI_LN_INFO]导入第%d条", nil, k+1)

		result, err = stmt.Exec(loanInfo.SerialNumber,
			loanInfo.ContractNumber,
			loanInfo.SubjectCode,
			loanInfo.SubjectName,
			loanInfo.CustCode,
			loanInfo.CustName,
			loanInfo.CustCodeCredit,
			loanInfo.CustCodeCore,
			loanInfo.ProductCode,
			loanInfo.ProductName,
			loanInfo.LoanType,
			loanInfo.CustType,
			loanInfo.OrgCodeCredit,
			loanInfo.OrgNameCredit,
			loanInfo.OrgCodeCore,
			loanInfo.OrgNameCore,
			loanInfo.OrgCodeMc,
			loanInfo.OrgNameMc,
			loanInfo.Area,
			loanInfo.IndustryCode,
			loanInfo.IndustryName,
			loanInfo.Amount,
			loanInfo.AmountRange,
			loanInfo.BaseRate,
			loanInfo.BaseRateMc,
			loanInfo.ReferenceAmount,
			loanInfo.InitRate,
			loanInfo.ExecutiveAmount,
			loanInfo.ContractStartDate,
			loanInfo.ContractStartMonth,
			loanInfo.RateCutDay,
			loanInfo.ContractEndDate,
			loanInfo.TermDay,
			loanInfo.TermMonth,
			loanInfo.TermMonthAmount,
			loanInfo.TermYear,
			loanInfo.TermInterval,
			loanInfo.RateAdjTypeCode,
			loanInfo.RateAdjTypeName,
			loanInfo.RePricingFreq,
			loanInfo.RePricingFreqUnit,
			loanInfo.RePayFreq,
			loanInfo.RePayFreqUnit,
			loanInfo.RePayMethodCode,
			loanInfo.RePayMethodName,
			loanInfo.GuaranteeWayCode,
			loanInfo.GuaranteeWayName,
			loanInfo.MainGuaranteeWayCode,
			loanInfo.MainGuaranteeWayName,
			loanInfo.Balance,
			loanInfo.CustScaleCodeGb,
			loanInfo.CustScaleNameGb,
			loanInfo.CustScaleCodeLine,
			loanInfo.CustScaleNameLine,
			loanInfo.CreditRating,
			loanInfo.AsOfDate)

		if nil != err {
			zlog.Errorf("[贷款基础信息-RPM_BI_LN_INFO]导入第%d行出错", err, k+1)
			//关闭连接
			err = stmt.Close()
			if nil != err {
				zlog.Error("[贷款基础信息-RPM_BI_LN_INFO]链接关闭出错", nil)
				sqlTx.Rollback()
				return nil, err
			}
			sqlTx.Rollback()
			return result, fmt.Errorf("[贷款基础信息-RPM_BI_LN_INFO]导入第%d行出错,%s", k+1, err.Error())
		}

	}
	//关闭连接
	err = stmt.Close()
	if nil != err {
		zlog.Error("[贷款基础信息-RPM_BI_LN_INFO]链接关闭出错", nil)
		sqlTx.Rollback()
		return nil, err
	}
	zlog.Info("[贷款基础信息-RPM_BI_LN_INFO]导入成功", nil)
	sqlTx.Commit()

	return nil, nil
}

// 查询金额根据金额倒序排
func (LoanInfoModel) FindAmount(asOfDate string) ([][]float64, error) {
	sqlText := `SELECT AMOUNT FROM RPM_BI_LN_INFO  T ` + whereSql
	if "" != asOfDate {
		sqlText += `and to_char(T.as_of_date,'yyyy-mm-dd') = :1`
	}
	sqlText += `
		 ORDER BY AMOUNT DESC
	`
	zlog.Infof("查询金额根据金额排序sql【%v】", nil, sqlText)
	var rows *sql.Rows
	var err error = nil
	if "" != asOfDate {
		rows, err = dbobj.Default.Query(sqlText, asOfDate)
	} else {
		rows, err = dbobj.Default.Query(sqlText, asOfDate)
	}
	defer rows.Close()
	if nil != err {
		er := fmt.Errorf("查询贷款金额降序失败")
		zlog.Error(er.Error(), err)
		return nil, er
	}
	var i float64 = 0
	var rst [][]float64
	for rows.Next() {
		var one []float64
		var amount float64
		err := rows.Scan(&amount)
		if nil != err {
			er := fmt.Errorf("查询贷款金额降序scan()失败")
			zlog.Error(er.Error(), err)
			return nil, er
		}
		i++
		one = append(one, i)
		one = append(one, amount)
		rst = append(rst, one)
	}
	return rst, nil
}

var whereSql string = ` where loan_type = 'GS' AND AMOUNT < 100000000`

func (LoanInfoModel) ListDate() ([]string, error) {
	sql := "select distinct as_of_date from RPM_BI_LN_INFO T " + whereSql
	rows, err := dbobj.Default.Query(sql)
	defer rows.Close()
	if nil != err {
		er := fmt.Errorf("查询贷款业务基础信【RPM_BI_LN_INFO】息日期信息出错")
		zlog.Error(er.Error(), err)
		return nil, er
	}
	var rst []string
	for rows.Next() {
		var one time.Time
		err := rows.Scan(&one)
		if nil != err {
			er := fmt.Errorf("查询贷款业务基础信【RPM_BI_LN_INFO】息日期信息row.Scan()出错")
			zlog.Error(er.Error(), err)
			return nil, er
		}
		rst = append(rst, one.Format("2006-01-02"))
	}
	return rst, nil
}

var insertSQL = `
  			INSERT
INTO RPM_BI_LN_INFO
  (
    SERIAL_NUMBER,
    CONTRACT_NUMBER,
    SUBJECT_CODE,
    SUBJECT_NAME,
    CUST_CODE,
	CUST_NAME,
    CUST_CODE_CREDIT,
    CUST_CODE_CORE,
    PRODUCT_CODE,
    PRODUCT_NAME,
    LOAN_TYPE,
    CUST_TYPE,
    ORG_CODE_CREDIT,
    ORG_NAME_CREDIT,
    ORG_CODE_CORE,
    ORG_NAME_CORE,
    ORG_CODE_MC,
    ORG_NAME_MC,
    AREA,
    INDUSTRY_CODE,
    INDUSTRY_NAME,
    AMOUNT,
    AMOUNT_RANGE,
    BASE_RATE,
    BASE_RATE_MC,
    REFERENCE_AMOUNT,
    INIT_RATE,
    EXECUTIVE_AMOUNT,
    CONTRACT_START_DATE,
    CONTRACT_START_MONTH,
    RATE_CUT_DAY,
    CONTRACT_END_DATE,
    TERM_DAY,
    TERM_MONTH,
    TERM_MONTH_AMOUNT,
    TERM_YEAR,
    TERM_INTERVAL,
    RATE_ADJ_TYPE_CODE,
    RATE_ADJ_TYPE_NAME,
    RE_PRICING_FREQ,
    RE_PRICING_FREQ_UNIT,
    RE_PAY_FREQ,
    RE_PAY_FREQ_UNIT,
    RE_PAY_METHOD_CODE,
    RE_PAY_METHOD_NAME,
    GUARANTEE_WAY_CODE,
    GUARANTEE_WAY_NAME,
    MAIN_GUARANTEE_WAY_CODE,
    MAIN_GUARANTEE_WAY_NAME,
    BALANCE,
    CUST_SCALE_CODE_GB,
    CUST_SCALE_NAME_GB,
    CUST_SCALE_CODE_LINE,
    CUST_SCALE_NAME_LINE,
    CREDIT_RATING,
	AS_OF_DATE
  )
  VALUES
  (
    :0,
    :1,
    :2,
    :3,
    :4,
    :5,
    :6,
    :7,
    :8,
    :9,
    :10,
    :11,
    :12,
    :13,
    :14,
    :15,
    :16,
    :17,
    :18,
    :19,
    :20,
    :21,
    :22,
    :23,
    :24,
    :25,
    :26,
    :27,
    :28,
    :29,
    :30,
    :31,
    :32,
    :33,
    :34,
    :35,
    :36,
    :37,
    :38,
    :39,
    :40,
    :41,
    :42,
    :43,
    :44,
    :45,
    :46,
    :47,
    :48,
    :49,
    :50,
    :51,
    :52,
    :53,
	:54,
	:55
  )
  `
var truncateSQL = "DELETE FROM RPM_BI_LN_INFO"
