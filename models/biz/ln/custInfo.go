package ln

import (
	"database/sql"
	"fmt"
	"time"

	"pccqcpa.com.cn/components/zlog"

	"pccqcpa.com.cn/app/rpm/api/models/dim"
	"pccqcpa.com.cn/app/rpm/api/models/sys"
	modelsUtil "pccqcpa.com.cn/app/rpm/api/models/util"
	"pccqcpa.com.cn/app/rpm/api/util"
)

type CustInfo struct {
	UUID              string       //主键
	CustCode          string       //客户编号 唯一
	CustName          string       //客户名称
	Organization      string       //组织机构代码
	CustType          string       //客户类型
	CustImplvl        string       //客户级别
	CustCredit        string       //客户信用
	Branch            sys.Organ    //客户机构
	Industry          dim.Industry //行业
	CustSize          string       //客户规模
	CustCapital       float64      //注册资本
	StockContribute   float64      //存量贡献
	StockUsage        float64      //存量贡献使用
	StockFreeze       float64      //存量优惠冻结
	UseProduct        int          //使用产品数
	CooperationPeriod int          // 合作年限数
	GapProportion     float64      // 实际EVA比例
	Dp020Threshold    string       // 一对一门槛
	EvaGap            float64      // 客户EVA缺口
	Status            string       //状态位
	Flag              string       //生效标志
	CreateTime        time.Time    //创建日期
	UpdateTime        time.Time    //更新日期
	CreateUser        string       //创建人
	UpdateUser        string       //更新人
	Owner             string       //所属人
}

// 客户信息rows.Scan()
func scan(rows *sql.Rows) (*CustInfo, error) {
	var custInfo = new(CustInfo)
	values := []interface{}{
		&custInfo.UUID,
		&custInfo.CustCode,
		&custInfo.CustName,
		&custInfo.Organization,
		&custInfo.CustType,
		&custInfo.CustImplvl,
		&custInfo.CustCredit,
		&custInfo.CustSize,
		&custInfo.CustCapital,
		&custInfo.StockContribute,
		&custInfo.StockUsage,
		&custInfo.StockFreeze,
		&custInfo.UseProduct,
		&custInfo.CooperationPeriod,
		&custInfo.GapProportion,
		&custInfo.Dp020Threshold,
		&custInfo.EvaGap,
		&custInfo.Status,
		&custInfo.Flag,
		&custInfo.CreateTime,
		&custInfo.UpdateTime,
		&custInfo.CreateUser,
		&custInfo.UpdateUser,
		&custInfo.Branch.OrganCode,
		&custInfo.Branch.OrganName,
		&custInfo.Industry.IndustryCode,
		&custInfo.Industry.IndustryName,
		&custInfo.Owner,
	}
	err := util.OracleScan(rows, values)
	return custInfo, err
}

// 查询客户分页操作
func (c *CustInfo) List(param ...map[string]interface{}) (*util.PageData, error) {

	var pageData = new(util.PageData)

	var tableName string = custTables

	pageData, rows, err := modelsUtil.List(tableName, custCols, custColsSort, param...)
	defer rows.Close()

	if nil != err {
		er := fmt.Errorf("查询客户分页出错")
		zlog.Error(er.Error(), err)
		return nil, er
	}
	custInfos, err := c.handleRows(rows)
	if nil != err {
		return nil, err
	}
	pageData.Rows = custInfos
	return pageData, nil
}

func (c CustInfo) UnionList(paramMap map[string]interface{}, unionParam ...map[string]interface{}) (*util.PageData, error) {
	modelUtil := modelsUtil.NewModelUtil()
	modelUtil.SetTableMsg(custTables, custCols, custColsSort)
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

	custInfos, err := c.handleRows(rows)
	if nil != err {
		return nil, err
	}
	pageData.Rows = custInfos
	return pageData, nil
}

func (c CustInfo) handleRows(rows *sql.Rows) ([]*CustInfo, error) {
	var custInfos []*CustInfo
	for rows.Next() {
		custInfo, err := scan(rows)
		if nil != err {
			er := fmt.Errorf("查询客户分页row.Scan()出错")
			zlog.Error(er.Error(), err)
			return nil, er
		}
		custInfos = append(custInfos, custInfo)
	}

	return custInfos, nil
}

func (c *CustInfo) Find(param ...map[string]interface{}) ([]*CustInfo, error) {
	rows, err := modelsUtil.FindRows(custTables, custCols, custColsSort, param...)
	if nil != err {
		er := fmt.Errorf("查询客户信息出错")
		zlog.Error(er.Error(), err)
		return nil, er
	}
	defer rows.Close()
	var custInfos []*CustInfo
	for rows.Next() {
		custInfo, err := scan(rows)
		if nil != err {
			er := fmt.Errorf("查询客户信息rows.Scan()出错")
			zlog.Error(er.Error(), err)
			return nil, er
		}
		custInfos = append(custInfos, custInfo)
	}
	return custInfos, nil
}

// 新增对公客户信息
func (c *CustInfo) Add() error {
	zlog.AppOperateLog(c.UpdateUser, "InsertCustInfo", zlog.ADD, c, nil, "新增模拟客户信息")

	param := map[string]interface{}{
		"cust_code":          c.CustCode,
		"cust_name":          c.CustName,
		"organization":       c.Organization,
		"cust_type":          c.CustType,
		"cust_implvl":        c.CustImplvl,
		"cust_credit":        c.CustCredit,
		"branch":             c.Branch.OrganCode,
		"industry":           c.Industry.IndustryCode,
		"cust_size":          c.CustSize,
		"cust_capital":       c.CustCapital,
		"STOCK_CONTRIBUTE":   c.StockContribute,
		"STOCK_USAGE":        c.StockUsage,
		"STOCK_FREEZE":       c.StockFreeze,
		"use_product":        c.UseProduct,
		"cooperation_period": c.CooperationPeriod,
		"gap_proportion":     c.GapProportion,
		"FLAG":               util.FLAG_TRUE,
		"status":             util.SIMULATION_CUST,
		"create_time":        util.GetCurrentTime(),
		"create_user":        c.CreateUser,
		"update_time":        util.GetCurrentTime(),
		"update_user":        c.UpdateUser,
		"owner":              c.Owner,
	}
	err := util.OracleAdd(custInfoTableName, param)
	if nil != err {
		er := fmt.Errorf("新增客户信息出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

// 更新客户信息
func (c *CustInfo) Update() error {
	zlog.AppOperateLog(c.UpdateUser, "UpdateCustInfo", zlog.UPDATE, c, nil, "更新客户信息")
	param := map[string]interface{}{
		"cust_name":          c.CustName,
		"organization":       c.Organization,
		"cust_type":          c.CustType,
		"cust_implvl":        c.CustImplvl,
		"cust_credit":        c.CustCredit,
		"branch":             c.Branch.OrganCode,
		"industry":           c.Industry.IndustryCode,
		"cust_size":          c.CustSize,
		"cust_capital":       c.CustCapital,
		"STOCK_CONTRIBUTE":   c.StockContribute,
		"STOCK_USAGE":        c.StockUsage,
		"STOCK_FREEZE":       c.StockFreeze,
		"use_product":        c.UseProduct,
		"cooperation_period": c.CooperationPeriod,
		"gap_proportion":     c.GapProportion,
		"status":             c.Status,
		"update_time":        util.GetCurrentTime(),
		"update_user":        c.UpdateUser,
		"owner":              c.Owner,
	}
	whereParam := map[string]interface{}{
		"uuid":      c.UUID,
		"cust_code": c.CustCode,
	}
	err := util.OracleUpdate(custInfoTableName, param, whereParam)
	if nil != err {
		er := fmt.Errorf("更新客户信息失败")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil

}

// 删除客户信息，先判断该用户是否为模拟客户，如果不是则不可以删除
// 如果是模拟客户，则判断是否有业务单子引用该客户，如果该客户有业务，则不可以删除
func (c *CustInfo) Delete() error {
	whereParam := map[string]interface{}{
		"uuid":      c.UUID,
		"cust_code": c.CustCode,
	}
	err := util.OracleDelete(custInfoTableName, whereParam)
	if nil != err {
		er := fmt.Errorf("删除客户信息失败")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

var custInfoTableName string = "rpm_biz_cust_info"

var custTables string = `
	     RPM_BIZ_CUST_INFO T
	LEFT JOIN SYS_SEC_ORGAN BRANCH
	  ON (T.BRANCH = BRANCH.ORGAN_CODE)
	LEFT JOIN RPM_DIM_INDUSTRY INDUSTRY
	  ON (T.INDUSTRY = INDUSTRY.INDUSTRY_CODE)
`

var cmCustTables string = custTables + `
	left join rpm_biz_user_cust_real t4
	  on (T.uuid = t4.cust_uuid)
`

var custCols map[string]string = map[string]string{
	"T.UUID":             "''",
	"T.CUST_CODE":        "' '",
	"T.CUST_NAME":        "' '",
	"T.ORGANIZATION":     "' '",
	"T.CUST_TYPE":        "' '",
	"T.CUST_IMPLVL":      "' '",
	"T.CUST_CREDIT":      "' '",
	"T.CUST_SIZE":        "' '",
	"T.CUST_CAPITAL":     "0",
	"T.STOCK_CONTRIBUTE": "0",
	"T.STOCK_USAGE":      "0",
	"T.STOCK_FREEZE":     "0",

	"T.USE_PRODUCT":        "0",
	"T.COOPERATION_PERIOD": "0",
	"T.GAP_PROPORTION":     "0",
	"T.DP_020_THRESHOLD":   "' '",
	"T.EVA_GAP":            "0",
	"T.STATUS":             "' '",
	"T.FLAG":               "' '",
	"T.CREATE_TIME":        "sysdate",
	"T.UPDATE_TIME":        "sysdate",
	"T.CREATE_USER":        "' '",
	"T.UPDATE_USER":        "' '",
	"T.BRANCH":             "' '",
	"BRANCH.ORGAN_NAME":    "' '",

	"T.INDUSTRY":             "' '",
	"INDUSTRY.INDUSTRY_NAME": "' '",
	"T.OWNER":                "' '",
}

var custColsSort = []string{
	"T.UUID",
	"T.CUST_CODE",
	"T.CUST_NAME",
	"T.ORGANIZATION",
	"T.CUST_TYPE",
	"T.CUST_IMPLVL",
	"T.CUST_CREDIT",
	"T.CUST_SIZE",
	"T.CUST_CAPITAL",
	"T.STOCK_CONTRIBUTE",
	"T.STOCK_USAGE",
	"T.STOCK_FREEZE",
	"T.USE_PRODUCT",
	"T.COOPERATION_PERIOD",
	"T.GAP_PROPORTION",
	"T.DP_020_THRESHOLD",
	"T.EVA_GAP",
	"T.STATUS",
	"T.FLAG",
	"T.CREATE_TIME",
	"T.UPDATE_TIME",
	"T.CREATE_USER",
	"T.UPDATE_USER",

	"T.BRANCH",
	"BRANCH.ORGAN_NAME",

	"T.INDUSTRY",
	"INDUSTRY.INDUSTRY_NAME",
	"T.OWNER",
}

var newWhereSQL = `
SELECT 
  rownum,
  COALESCE(T.UUID,''),
  COALESCE(T.CUST_CODE,' '),
  COALESCE(T.CUST_NAME,' '),
  COALESCE(T.ORGANIZATION,' '),
  COALESCE(T.CUST_TYPE,' '),
  COALESCE(T.CUST_IMPLVL,' '),
  COALESCE(T.CUST_CREDIT,' '),
  COALESCE(T.CUST_SIZE,' '),
  COALESCE(T.CUST_CAPITAL,0),
  COALESCE(T.STOCK_CONTRIBUTE,0),
  COALESCE(T.STOCK_USAGE,0),
  COALESCE(T.STOCK_FREEZE,0),
  COALESCE(T.USE_PRODUCT,0),
  COALESCE(T.COOPERATION_PERIOD,0),
  COALESCE(T.GAP_PROPORTION,0),
  COALESCE(T.DP_020_THRESHOLD,' '),
  COALESCE(T.EVA_GAP,0),
  COALESCE(T.STATUS,' '),
  COALESCE(T.FLAG,' '),
  COALESCE(T.CREATE_TIME,sysdate),
  COALESCE(T.UPDATE_TIME,sysdate),
  COALESCE(T.CREATE_USER,' '),
  COALESCE(T.UPDATE_USER,' '),
  COALESCE(T.BRANCH,' '),
  COALESCE(BRANCH.ORGAN_NAME,' '),
  COALESCE(T.INDUSTRY,' '),
  COALESCE(INDUSTRY.INDUSTRY_NAME,' ')
FROM
     RPM_BIZ_CUST_INFO t
    LEFT JOIN SYS_SEC_ORGAN BRANCH
    ON (T.BRANCH = BRANCH.ORGAN_CODE)
    LEFT JOIN RPM_DIM_INDUSTRY INDUSTRY
    ON (T.INDUSTRY = INDUSTRY.INDUSTRY_CODE)
    WHERE t.STATUS ='01' AND t.FLAG='1' `

var stkWhereSQL = `
    SELECT 
  rownum,
  COALESCE(T.UUID,''),
  COALESCE(T.CUST_CODE,' '),
  COALESCE(T.CUST_NAME,' '),
  COALESCE(T.ORGANIZATION,' '),
  COALESCE(T.CUST_TYPE,' '),
  COALESCE(T.CUST_IMPLVL,' '),
  COALESCE(T.CUST_CREDIT,' '),
  COALESCE(T.CUST_SIZE,' '),
  COALESCE(T.CUST_CAPITAL,0),
  COALESCE(T.STOCK_CONTRIBUTE,0),
  COALESCE(T.STOCK_USAGE,0),
  COALESCE(T.STOCK_FREEZE,0),
  COALESCE(T.USE_PRODUCT,0),
  COALESCE(T.COOPERATION_PERIOD,0),
  COALESCE(T.GAP_PROPORTION,0),
  COALESCE(T.DP_020_THRESHOLD,' '),
  COALESCE(T.EVA_GAP,0),
  COALESCE(T.STATUS,' '),
  COALESCE(T.FLAG,' '),
  COALESCE(T.CREATE_TIME,sysdate),
  COALESCE(T.UPDATE_TIME,sysdate),
  COALESCE(T.CREATE_USER,' '),
  COALESCE(T.UPDATE_USER,' '),
  COALESCE(T.BRANCH,' '),
  COALESCE(BRANCH.ORGAN_NAME,' '),
  COALESCE(T.INDUSTRY,' '),
  COALESCE(INDUSTRY.INDUSTRY_NAME,' ')
FROM
     RPM_BIZ_CUST_INFO t
    LEFT JOIN SYS_SEC_ORGAN BRANCH
    ON (T.BRANCH = BRANCH.ORGAN_CODE)
    LEFT JOIN RPM_DIM_INDUSTRY INDUSTRY
    ON (T.INDUSTRY = INDUSTRY.INDUSTRY_CODE)
    WHERE t.STATUS ='02' AND t.FLAG='1'
`
