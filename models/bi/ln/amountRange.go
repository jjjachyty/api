package ln

import (
	"database/sql"
	"fmt"
	"time"

	modelsUtil "pccqcpa.com.cn/app/rpm/api/models/util"
	"pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/components/zlog"
)

// 结构体对应数据库表【RPM_BI_LN_AMOUNT_RANGE】金额区间维度表
// by author Jason
// by time 2016-10-31 15:42:39
type LnAmountRange struct {
	UUID       string    // 主键
	Range      string    // 金额区间
	AmountL    string    // 金额上限
	AmountU    string    // 金额下限
	CreateTime time.Time // 创建时间
}

// 结构体scan
// by author Jason
// by time 2016-10-31 15:42:39
func (this LnAmountRange) scan(rows *sql.Rows) (*LnAmountRange, error) {
	var one = new(LnAmountRange)
	values := []interface{}{
		&one.UUID,
		&one.Range,
		&one.AmountL,
		&one.AmountU,
		&one.CreateTime,
	}
	err := util.OracleScan(rows, values)
	return one, err
}

// 分页操作
// by author Jason
// by time 2016-10-31 15:42:39
func (this LnAmountRange) List(param ...map[string]interface{}) (*util.PageData, error) {
	pageData, rows, err := modelsUtil.List(lnAmountRangeTabales, lnAmountRangeCols, lnAmountRangeColsSort, param...)
	defer rows.Close()
	if nil != err {
		er := fmt.Errorf("分页查询金额区间维度表信息出错")
		zlog.Error(er.Error(), err)
		return nil, er
	}

	var rst []*LnAmountRange
	for rows.Next() {
		one, err := this.scan(rows)
		if nil != err {
			er := fmt.Errorf("分页查询金额区间维度表信息row.Scan()出错")
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
// by time 2016-10-31 15:42:39
func (this LnAmountRange) Find(param ...map[string]interface{}) ([]*LnAmountRange, error) {
	rows, err := modelsUtil.FindRows(lnAmountRangeTabales, lnAmountRangeCols, lnAmountRangeColsSort, param...)
	defer rows.Close()
	if nil != err {
		er := fmt.Errorf("多参数查询金额区间维度表信息出错")
		zlog.Error(er.Error(), err)
		return nil, er
	}
	defer rows.Close()
	var rst []*LnAmountRange
	for rows.Next() {
		one, err := this.scan(rows)
		if nil != err {
			er := fmt.Errorf("多参数查询金额区间维度表信息row.Scan()出错")
			zlog.Error(er.Error(), err)
			return nil, er
		}
		rst = append(rst, one)
	}
	return rst, nil
}

// 新增金额区间维度表信息
// by author Jason
// by time 2016-10-31 15:42:39
func (this LnAmountRange) Add() error {
	paramMap := map[string]interface{}{
		"range":       this.Range,
		"amount_l":    this.AmountL,
		"amount_u":    this.AmountU,
		"create_time": util.GetCurrentTime(),
	}
	err := util.OracleAdd("RPM_BI_LN_AMOUNT_RANGE", paramMap)
	if nil != err {
		er := fmt.Errorf("新增金额区间维度表信息出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

// 新增贷款业务分析基准利率降息维度表信息
// by author Jason
// by time 2016-10-31 15:34:54
func (this LnAmountRange) BatchAdd(sqlTx *sql.Tx) (sql.Result, error) {
	paramMap := map[string]interface{}{
		"range":       this.Range,
		"amount_l":    this.AmountL,
		"amount_u":    this.AmountU,
		"create_time": util.GetCurrentTime(),
	}
	result, err := util.OracleBatchAdd("RPM_BI_LN_AMOUNT_RANGE", paramMap, sqlTx)
	return result, err
}

// 更新金额区间维度表信息
// by author Jason
// by time 2016-10-31 15:42:39
func (this LnAmountRange) Update() error {
	paramMap := map[string]interface{}{
		"range":    this.Range,
		"amount_l": this.AmountL,
		"amount_u": this.AmountU,
	}
	whereParam := map[string]interface{}{
		"uuid": this.UUID,
	}
	err := util.OracleUpdate("RPM_BI_LN_AMOUNT_RANGE", paramMap, whereParam)
	if nil != err {
		er := fmt.Errorf("更新金额区间维度表信息出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

// 删除金额区间维度表信息
// by author Jason
// by time 2016-10-31 15:42:39
func (this LnAmountRange) Delete() error {
	whereParam := map[string]interface{}{
		"uuid": this.UUID,
	}
	err := util.OracleDelete("RPM_BI_LN_AMOUNT_RANGE", whereParam)
	if nil != err {
		er := fmt.Errorf("删除金额区间维度表信息出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

var lnAmountRangeTabales string = `RPM_BI_LN_AMOUNT_RANGE`

var lnAmountRangeCols map[string]string = map[string]string{
	"UUID":        "' '",
	"RANGE":       "' '",
	"AMOUNT_L":    "' '",
	"AMOUNT_U":    "' '",
	"CREATE_TIME": "sysdate",
}

var lnAmountRangeColsSort = []string{
	"UUID",
	"RANGE",
	"AMOUNT_L",
	"AMOUNT_U",
	"CREATE_TIME",
}
