package par

import (
	"database/sql"
	"fmt"
	"time"

	modelsUtil "pccqcpa.com.cn/app/rpm/api/models/util"
	"pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/components/zlog"
)

// 结构体对应数据库表【RPM_PAR_EVA_YIELD】EVA收益率表
// by author Yeqc
// by time 2016-12-19 15:10:36
/**
 * @apiDefine EvaYield
 * @apiSuccess {string}     UUID            主键默认值sys_guid()
 * @apiSuccess {float64}   	EvaYieldRate    EVA收益率
 * @apiSuccess {string}   	Flag      		生效标志
 * @apiSuccess {time.Time}	CreateTime      创建时间
 * @apiSuccess {string}   	CreateUser      创建人
 * @apiSuccess {time.Time}	UpdateTime      更新时间
 * @apiSuccess {string}   	UpdateUser      更新人
 */
type EvaYield struct {
	UUID         string    // 主键
	EvaYieldRate float64   // EVA收益率
	Flag         string    // 生效标志
	CreateTime   time.Time // 创建时间
	UpdateTime   time.Time // 更新时间
	CreateUser   string    // 创建人
	UpdateUser   string    // 更新人
}

// 结构体scan
// by author Yeqc
// by time 2016-12-19 15:10:36
func (this EvaYield) scan(rows *sql.Rows) (*EvaYield, error) {
	var one = new(EvaYield)
	values := []interface{}{
		&one.UUID,
		&one.EvaYieldRate,
		&one.Flag,
		&one.CreateTime,
		&one.UpdateTime,
		&one.CreateUser,
		&one.UpdateUser,
	}
	err := util.OracleScan(rows, values)
	return one, err
}

// 分页操作
// by author Yeqc
// by time 2016-12-19 15:10:36
func (this EvaYield) List(param ...map[string]interface{}) (*util.PageData, error) {
	pageData, rows, err := modelsUtil.List(evaYieldTabales, evaYieldCols, evaYieldColsSort, param...)
	defer rows.Close()
	if nil != err {
		er := fmt.Errorf("分页查询EVA收益率表信息出错")
		zlog.Error(er.Error(), err)
		return nil, er
	}

	var rst []*EvaYield
	for rows.Next() {
		one, err := this.scan(rows)
		if nil != err {
			er := fmt.Errorf("分页查询EVA收益率表信息row.Scan()出错")
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
// by time 2016-12-19 15:10:36
func (this EvaYield) Find(param ...map[string]interface{}) ([]*EvaYield, error) {
	rows, err := modelsUtil.FindRows(evaYieldTabales, evaYieldCols, evaYieldColsSort, param...)
	defer rows.Close()
	if nil != err {
		er := fmt.Errorf("多参数查询EVA收益率表信息出错")
		zlog.Error(er.Error(), err)
		return nil, er
	}
	var rst []*EvaYield
	for rows.Next() {
		one, err := this.scan(rows)
		if nil != err {
			er := fmt.Errorf("多参数查询EVA收益率表信息row.Scan()出错")
			zlog.Error(er.Error(), err)
			return nil, er
		}
		rst = append(rst, one)
	}
	return rst, nil
}

// 新增EVA收益率表信息
// by author Yeqc
// by time 2016-12-19 15:10:36
func (this EvaYield) Add() error {
	paramMap := map[string]interface{}{
		"eva_yield_rate": this.EvaYieldRate,
		"flag":           this.Flag,
		"create_time":    util.GetCurrentTime(),
		"update_time":    util.GetCurrentTime(),
		"create_user":    this.CreateUser,
		"update_user":    this.UpdateUser,
	}
	err := util.OracleAdd("RPM_PAR_EVA_YIELD", paramMap)
	if nil != err {
		er := fmt.Errorf("新增EVA收益率表信息出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

// 更新EVA收益率表信息
// by author Yeqc
// by time 2016-12-19 15:10:36
func (this EvaYield) Update() error {
	paramMap := map[string]interface{}{
		"eva_yield_rate": this.EvaYieldRate,
		"flag":           this.Flag,
		"update_time":    util.GetCurrentTime(),
		"update_user":    this.UpdateUser,
	}
	whereParam := map[string]interface{}{
		"uuid": this.UUID,
	}
	err := util.OracleUpdate("RPM_PAR_EVA_YIELD", paramMap, whereParam)
	if nil != err {
		er := fmt.Errorf("更新EVA收益率表信息出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

// 删除EVA收益率表信息
// by author Yeqc
// by time 2016-12-19 15:10:36
func (this EvaYield) Delete() error {
	whereParam := map[string]interface{}{
		"uuid": this.UUID,
	}
	err := util.OracleDelete("RPM_PAR_EVA_YIELD", whereParam)
	if nil != err {
		er := fmt.Errorf("删除EVA收益率表信息出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

var evaYieldTabales string = `RPM_PAR_EVA_YIELD T`

var evaYieldCols map[string]string = map[string]string{
	"T.UUID":           "' '",
	"T.EVA_YIELD_RATE": "0",
	"T.FLAG":           "' '",
	"T.CREATE_TIME":    "sysdate",
	"T.UPDATE_TIME":    "sysdate",
	"T.CREATE_USER":    "' '",
	"T.UPDATE_USER":    "' '",
}

var evaYieldColsSort = []string{
	"T.UUID",
	"T.EVA_YIELD_RATE",
	"T.FLAG",
	"T.CREATE_TIME",
	"T.UPDATE_TIME",
	"T.CREATE_USER",
	"T.UPDATE_USER",
}
