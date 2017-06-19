package par

import (
	"database/sql"
	"fmt"

	modelsUtil "pccqcpa.com.cn/app/rpm/api/models/util"
	"pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/components/zlog"
)

// 结构体对应数据库表【RPM_PAR_RL_ROC】零售资本回报率-零售模块
// by author Yeqc
// by time 2016-12-30 11:31:37
/**
 * @apiDefine RlRoc
 * @apiSuccess {string}     UUID            		主键默认值sys_guid()
 * @apiSuccess {float64}   	Amount    				金额区间值
 * @apiSuccess {float64}   	CapitalCost      		资本成本率 百分比
 * @apiSuccess {float64}   	CapitalProfit   		资本回报率 百分比
 */
type RlRoc struct {
	UUID          string  // 主键默认值sys_guid()
	Amount        float64 // 金额区间值
	CapitalCost   float64 // 资本成本率 百分比
	CapitalProfit float64 // 资本回报率 百分比
}

// 结构体scan
// by author Yeqc
// by time 2016-12-30 11:31:37
func (this RlRoc) scan(rows *sql.Rows) (*RlRoc, error) {
	var one = new(RlRoc)
	values := []interface{}{
		&one.UUID,
		&one.Amount,
		&one.CapitalCost,
		&one.CapitalProfit,
	}
	err := util.OracleScan(rows, values)
	return one, err
}

// 分页操作
// by author Yeqc
// by time 2016-12-30 11:31:37
func (this RlRoc) List(param ...map[string]interface{}) (*util.PageData, error) {
	pageData, rows, err := modelsUtil.List(rlRocTabales, rlRocCols, rlRocColsSort, param...)
	defer rows.Close()
	if nil != err {
		er := fmt.Errorf("分页查询零售资本回报率-零售模块信息出错")
		zlog.Error(er.Error(), err)
		return nil, er
	}

	var rst []*RlRoc
	for rows.Next() {
		one, err := this.scan(rows)
		if nil != err {
			er := fmt.Errorf("分页查询零售资本回报率-零售模块信息row.Scan()出错")
			zlog.Error(er.Error(), err)
			return nil, er
		}
		rst = append(rst, one)
	}
	pageData.Rows = rst
	return pageData, nil
}

// 多参数查询
// by author Yeqc
// by time 2016-12-30 11:31:37
func (this RlRoc) Find(param ...map[string]interface{}) ([]*RlRoc, error) {
	rows, err := modelsUtil.FindRows(rlRocTabales, rlRocCols, rlRocColsSort, param...)
	defer rows.Close()
	if nil != err {
		er := fmt.Errorf("多参数查询零售资本回报率-零售模块信息出错")
		zlog.Error(er.Error(), err)
		return nil, er
	}
	var rst []*RlRoc
	for rows.Next() {
		one, err := this.scan(rows)
		if nil != err {
			er := fmt.Errorf("多参数查询零售资本回报率-零售模块信息row.Scan()出错")
			zlog.Error(er.Error(), err)
			return nil, er
		}
		rst = append(rst, one)
	}
	return rst, nil
}

// 新增零售资本回报率-零售模块信息
// by author Yeqc
// by time 2016-12-30 11:31:37
func (this RlRoc) Add() error {
	paramMap := map[string]interface{}{
		"amount":         this.Amount,
		"capital_cost":   this.CapitalCost,
		"capital_profit": this.CapitalProfit,
	}
	err := util.OracleAdd("RPM_PAR_RL_ROC", paramMap)
	if nil != err {
		er := fmt.Errorf("新增零售资本回报率-零售模块信息出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

// 更新零售资本回报率-零售模块信息
// by author Yeqc
// by time 2016-12-30 11:31:37
func (this RlRoc) Update() error {
	paramMap := map[string]interface{}{
		"amount":         this.Amount,
		"capital_cost":   this.CapitalCost,
		"capital_profit": this.CapitalProfit,
	}
	whereParam := map[string]interface{}{
		"uuid": this.UUID,
	}
	err := util.OracleUpdate("RPM_PAR_RL_ROC", paramMap, whereParam)
	if nil != err {
		er := fmt.Errorf("更新零售资本回报率-零售模块信息出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

// 删除零售资本回报率-零售模块信息
// by author Yeqc
// by time 2016-12-30 11:31:37
func (this RlRoc) Delete() error {
	whereParam := map[string]interface{}{
		"uuid": this.UUID,
	}
	err := util.OracleDelete("RPM_PAR_RL_ROC", whereParam)
	if nil != err {
		er := fmt.Errorf("删除零售资本回报率-零售模块信息出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

var rlRocTabales string = `RPM_PAR_RL_ROC T`

var rlRocCols map[string]string = map[string]string{
	"T.UUID":           "' '",
	"T.AMOUNT":         "0",
	"T.CAPITAL_COST":   "0",
	"T.CAPITAL_PROFIT": "0",
}

var rlRocColsSort = []string{
	"T.UUID",
	"T.AMOUNT",
	"T.CAPITAL_COST",
	"T.CAPITAL_PROFIT",
}
