package dp

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"pccqcpa.com.cn/app/rpm/api/models/biz/ln"
	"pccqcpa.com.cn/app/rpm/api/models/sys"
	modelsUtil "pccqcpa.com.cn/app/rpm/api/models/util"
	"pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/components/zlog"
)

// 结构体对应数据库表【RPM_BIZ_DP_ONE_PRICING】存款一对一结果表
// by author Jason
// by time 2016-12-12 10:44:29
type DpOnePricing struct {
	UUID          string      // 主键默认值sys_guid()
	Cust          ln.CustInfo // 存款客户
	BusinessCode  string      // 业务单号
	Organ         sys.Organ   // 业务办理机构
	CurrentEva    float64     // 当前EVA
	IbEva         float64     // 派生EVA
	SumDp         float64     // 存款总额
	SumIb         float64     // 中间业务总额
	SumBreak      float64     // 总盈亏
	LostCost      float64     // 流失机会成本
	NextEva       float64     // 未来EVA （当前EVA+派生EVA）
	CurrentDpLost float64     // 办理当前存款损失
	StockDpLost   float64     // 存量存款业务流失
	Status        string      // 状态
	CreateUser    string      // 创建人
	UpdateUser    string      // 更新人
	CreateTime    time.Time   // 创建时间
	UpdateTime    time.Time   // 更新时间
}

// 结构体scan
// by author Jason
// by time 2016-12-12 10:44:29
func (this DpOnePricing) scan(rows *sql.Rows) (*DpOnePricing, error) {
	var one = new(DpOnePricing)
	values := []interface{}{
		&one.UUID,
		&one.BusinessCode,
		&one.CurrentEva,
		&one.IbEva,
		&one.SumDp,
		&one.SumIb,
		&one.SumBreak,
		&one.LostCost,
		&one.NextEva,
		&one.CurrentDpLost,
		&one.StockDpLost,
		&one.Status,
		&one.CreateUser,
		&one.UpdateUser,
		&one.CreateTime,
		&one.UpdateTime,
		&one.Cust.CustCode,
		&one.Cust.CustName,
		&one.Cust.EvaGap,

		&one.Cust.CustSize,
		&one.Cust.CustCredit,
		&one.Cust.CustImplvl,
		&one.Cust.Industry.IndustryName,

		&one.Organ.OrganCode,
		&one.Organ.OrganName,
	}
	err := util.OracleScan(rows, values)
	return one, err
}

// 分页操作
// by author Jason
// by time 2016-12-12 10:44:29
func (this DpOnePricing) List(param ...map[string]interface{}) (*util.PageData, error) {
	pageData, rows, err := modelsUtil.List(dpOnePricingTabales, dpOnePricingCols, dpOnePricingColsSort, param...)
	defer rows.Close()
	if nil != err {
		er := fmt.Errorf("分页查询存款一对一结果表信息出错")
		zlog.Error(er.Error(), err)
		return nil, er
	}

	var rst []*DpOnePricing
	for rows.Next() {
		one, err := this.scan(rows)
		if nil != err {
			er := fmt.Errorf("分页查询存款一对一结果表信息row.Scan()出错")
			zlog.Error(er.Error(), err)
			return nil, er
		}
		rst = append(rst, one)
	}
	pageData.Rows = rst
	return pageData, nil
}

func (d DpOnePricing) UnionList(paramMap map[string]interface{}, unionParam ...map[string]interface{}) (*util.PageData, error) {
	modelUtil := modelsUtil.NewModelUtil()
	modelUtil.SetTableMsg(dpOnePricingTabales, dpOnePricingCols, dpOnePricingColsSort)
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

	dops, err := d.handleRows(rows)
	if nil != err {
		return nil, err
	}
	pageData.Rows = dops
	return pageData, nil
}

func (d DpOnePricing) handleRows(rows *sql.Rows) ([]*DpOnePricing, error) {
	var dops []*DpOnePricing
	for rows.Next() {
		dop, err := d.scan(rows)
		if nil != err {
			er := fmt.Errorf("查询客户分页row.Scan()出错")
			zlog.Error(er.Error(), err)
			return nil, er
		}
		dops = append(dops, dop)
	}
	return dops, nil
}

// 多参数查询
// by author Jason
// by time 2016-12-12 10:44:29
func (this DpOnePricing) Find(param ...map[string]interface{}) ([]*DpOnePricing, error) {
	rows, err := modelsUtil.FindRows(dpOnePricingTabales, dpOnePricingCols, dpOnePricingColsSort, param...)
	defer rows.Close()
	if nil != err {
		er := fmt.Errorf("多参数查询存款一对一结果表信息出错")
		zlog.Error(er.Error(), err)
		return nil, er
	}
	var rst []*DpOnePricing
	for rows.Next() {
		one, err := this.scan(rows)
		if nil != err {
			er := fmt.Errorf("多参数查询存款一对一结果表信息row.Scan()出错")
			zlog.Error(er.Error(), err)
			return nil, er
		}
		rst = append(rst, one)
	}
	return rst, nil
}

// 新增存款一对一结果表信息
// by author Jason
// by time 2016-12-12 10:44:29
func (this DpOnePricing) Add() error {
	paramMap := map[string]interface{}{
		"business_code":   this.BusinessCode,
		"cust":            this.Cust.CustCode,
		"organ":           this.Organ.OrganCode,
		"current_eva":     this.CurrentEva,
		"ib_eva":          this.IbEva,
		"sum_dp":          this.SumDp,
		"sum_ib":          this.SumIb,
		"sum_break":       this.SumBreak,
		"lost_cost":       this.LostCost,
		"next_eva":        this.NextEva,
		"current_dp_lost": this.CurrentDpLost,
		"stock_dp_lost":   this.StockDpLost,
		"status":          this.Status,
		"create_user":     this.CreateUser,
		"update_user":     this.UpdateUser,
		"create_time":     util.GetCurrentTime(),
		"update_time":     util.GetCurrentTime(),
	}
	err := util.OracleAdd("RPM_BIZ_DP_ONE_PRICING", paramMap)
	if nil != err {
		er := fmt.Errorf("新增存款一对一结果表信息出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

// 更新存款一对一结果表信息
// by author Jason
// by time 2016-12-12 10:44:29
func (this DpOnePricing) Update() error {
	paramMap := map[string]interface{}{
		"organ":           this.Organ.OrganCode,
		"current_eva":     this.CurrentEva,
		"ib_eva":          this.IbEva,
		"sum_dp":          this.SumDp,
		"sum_ib":          this.SumIb,
		"sum_break":       this.SumBreak,
		"lost_cost":       this.LostCost,
		"next_eva":        this.NextEva,
		"current_dp_lost": this.CurrentDpLost,
		"stock_dp_lost":   this.StockDpLost,
		"status":          this.Status,
		"update_user":     this.UpdateUser,
		"update_time":     util.GetCurrentTime(),
	}
	whereParam := map[string]interface{}{
		"business_code": this.BusinessCode,
	}
	err := util.OracleUpdate("RPM_BIZ_DP_ONE_PRICING", paramMap, whereParam)
	if nil != err {
		er := fmt.Errorf("更新存款一对一结果表信息出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

// 更新存款一对一结果表信息 部分更新
// by author Jason
// by time 2016-12-12 10:44:29
func (d DpOnePricing) Patch(paramMap map[string]interface{}) error {
	var businessCode string = ""
	if val, ok := paramMap["business_code"]; ok {
		businessCode = val.(string)
	} else if "" != strings.TrimSpace(d.BusinessCode) {
		businessCode = d.BusinessCode
	} else {
		er := fmt.Errorf("更新存款一对一定价单未传主键【business_code】")
		zlog.Error(er.Error(), er)
		return er
	}
	whereParam := map[string]interface{}{
		"business_code": businessCode,
	}
	err := util.OracleUpdate("RPM_BIZ_DP_ONE_PRICING", paramMap, whereParam)
	if nil != err {
		er := fmt.Errorf("更新存款一对一结果表信息出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

// 删除存款一对一结果表信息
// by author Jason
// by time 2016-12-12 10:44:29
func (this DpOnePricing) Delete() error {
	whereParam := map[string]interface{}{
		"uuid": this.UUID,
	}
	err := util.OracleDelete("RPM_BIZ_DP_ONE_PRICING", whereParam)
	if nil != err {
		er := fmt.Errorf("删除存款一对一结果表信息出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

var dpOnePricingTabales string = `
							 RPM_BIZ_DP_ONE_PRICING T 
						LEFT JOIN RPM_BIZ_CUST_INFO CUST
						  ON (T.CUST = CUST.CUST_CODE)
						LEFT JOIN RPM_DIM_INDUSTRY INDUSTRY
						  ON (CUST.INDUSTRY = INDUSTRY.INDUSTRY_CODE)
						LEFT JOIN SYS_SEC_ORGAN ORGAN
						  ON (T.ORGAN = ORGAN.ORGAN_CODE)`

var dpOnePricingCols map[string]string = map[string]string{
	"T.UUID":            "' '",
	"T.BUSINESS_CODE":   "' '",
	"T.CURRENT_EVA":     "0",
	"T.IB_EVA":          "0",
	"T.SUM_DP":          "0",
	"T.SUM_IB":          "0",
	"T.SUM_BREAK":       "0",
	"T.LOST_COST":       "0",
	"T.NEXT_EVA":        "0",
	"T.CURRENT_DP_LOST": "0",
	"T.STOCK_DP_LOST":   "0",
	"T.STATUS":          "' '",
	"T.CREATE_USER":     "' '",
	"T.UPDATE_USER":     "' '",
	"T.CREATE_TIME":     "sysdate",
	"T.UPDATE_TIME":     "sysdate",

	"CUST.CUST_CODE":   "' '",
	"CUST.CUST_NAME":   "' '",
	"CUST.EVA_GAP":     "0",
	"CUST.CUST_SIZE":   "' '",
	"CUST.CUST_CREDIT": "' '",
	"CUST.CUST_IMPLVL": "' '",

	"INDUSTRY.INDUSTRY_NAME": "' '",

	"T.ORGAN":          "' '",
	"ORGAN.ORGAN_NAME": "' '",
}

var dpOnePricingColsSort = []string{
	"T.UUID",
	"T.BUSINESS_CODE",
	"T.CURRENT_EVA",
	"T.IB_EVA",
	"T.SUM_DP",
	"T.SUM_IB",
	"T.SUM_BREAK",
	"T.LOST_COST",
	"T.NEXT_EVA",
	"T.CURRENT_DP_LOST",
	"T.STOCK_DP_LOST",
	"T.STATUS",
	"T.CREATE_USER",
	"T.UPDATE_USER",
	"T.CREATE_TIME",
	"T.UPDATE_TIME",

	"CUST.CUST_CODE",
	"CUST.CUST_NAME",
	"CUST.EVA_GAP",
	"CUST.CUST_SIZE",
	"CUST.CUST_CREDIT",
	"CUST.CUST_IMPLVL",

	"INDUSTRY.INDUSTRY_NAME",

	"T.ORGAN",
	"ORGAN.ORGAN_NAME",
}
