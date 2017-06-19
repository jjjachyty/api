package par

import (
	"database/sql"
	"fmt"
	"time"

	modelsUtil "pccqcpa.com.cn/app/rpm/api/models/util"
	"pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/components/zlog"
)

// 结构体对应数据库表【RPM_PAR_BASE_RATE】基准利率
// by author Yeqc
// by time 2016-12-19 10:49:29
/**
 * @apiDefine BaseRate
 * @apiSuccess {string}     UUID            主键默认值sys_guid()
 * @apiSuccess {string}   	BaseRateType    利率类型
 * @apiSuccess {string}   	Term      		期限
 * @apiSuccess {string}   	BaseRateName    基准利率名称
 * @apiSuccess {string}   	Rate      		基准利率
 * @apiSuccess {string}   	Flag      		生效标志
 * @apiSuccess {time.Time}	CreateTime      创建时间
 * @apiSuccess {string}   	CreateUser      创建人
 * @apiSuccess {time.Time}	UpdateTime      更新时间
 * @apiSuccess {string}   	UpdateUser      更新人
 */
type BaseRate struct {
	UUID         string    // 主键
	BaseRateType string    // 利率类型
	Term         float64   // 期限
	BaseRateName string    // 基准利率名称
	Rate         float64   // 基准利率
	Flag         string    // 生效标志
	CreateTime   time.Time // 创建时间
	UpdateTime   time.Time // 更新时间
	CreateUser   string    // 创建人
	UpdateUser   string    // 更新人
}

// 结构体scan
// by author Yeqc
// by time 2016-12-19 10:49:29
func (this BaseRate) scan(rows *sql.Rows) (*BaseRate, error) {
	var one = new(BaseRate)
	values := []interface{}{
		&one.UUID,
		&one.BaseRateType,
		&one.Term,
		&one.BaseRateName,
		&one.Rate,
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
// by time 2016-12-19 10:49:29
func (this BaseRate) List(param ...map[string]interface{}) (*util.PageData, error) {
	pageData, rows, err := modelsUtil.List(baseRateTabales, baseRateCols, baseRateColsSort, param...)
	defer rows.Close()
	if nil != err {
		er := fmt.Errorf("分页查询基准利率信息出错")
		zlog.Error(er.Error(), err)
		return nil, er
	}

	var rst []*BaseRate
	for rows.Next() {
		one, err := this.scan(rows)
		if nil != err {
			er := fmt.Errorf("分页查询基准利率信息row.Scan()出错")
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
// by time 2016-12-19 10:49:29
func (this BaseRate) Find(param ...map[string]interface{}) ([]*BaseRate, error) {
	rows, err := modelsUtil.FindRows(baseRateTabales, baseRateCols, baseRateColsSort, param...)
	defer rows.Close()
	if nil != err {
		er := fmt.Errorf("多参数查询基准利率信息出错")
		zlog.Error(er.Error(), err)
		return nil, er
	}
	var rst []*BaseRate
	for rows.Next() {
		one, err := this.scan(rows)
		if nil != err {
			er := fmt.Errorf("多参数查询基准利率信息row.Scan()出错")
			zlog.Error(er.Error(), err)
			return nil, er
		}
		rst = append(rst, one)
	}
	return rst, nil
}

// 新增基准利率信息
// by author Yeqc
// by time 2016-12-19 10:49:29
func (this BaseRate) Add() error {
	paramMap := map[string]interface{}{
		"base_rate_type": this.BaseRateType,
		"term":           this.Term,
		"base_rate_name": this.BaseRateName,
		"rate":           this.Rate,
		"flag":           this.Flag,
		"create_time":    util.GetCurrentTime(),
		"update_time":    util.GetCurrentTime(),
		"create_user":    this.CreateUser,
		"update_user":    this.UpdateUser,
	}
	err := util.OracleAdd("RPM_PAR_BASE_RATE", paramMap)
	if nil != err {
		er := fmt.Errorf("新增基准利率信息出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

// 更新基准利率信息
// by author Yeqc
// by time 2016-12-19 10:49:29
func (this BaseRate) Update() error {
	paramMap := map[string]interface{}{
		"base_rate_type": this.BaseRateType,
		"term":           this.Term,
		"base_rate_name": this.BaseRateName,
		"rate":           this.Rate,
		"flag":           this.Flag,
		"update_time":    util.GetCurrentTime(),
		"update_user":    this.UpdateUser,
	}
	whereParam := map[string]interface{}{
		"uuid": this.UUID,
	}
	err := util.OracleUpdate("RPM_PAR_BASE_RATE", paramMap, whereParam)
	if nil != err {
		er := fmt.Errorf("更新基准利率信息出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

// 删除基准利率信息
// by author Yeqc
// by time 2016-12-19 10:49:29
func (this BaseRate) Delete() error {
	whereParam := map[string]interface{}{
		"uuid": this.UUID,
	}
	err := util.OracleDelete("RPM_PAR_BASE_RATE", whereParam)
	if nil != err {
		er := fmt.Errorf("删除基准利率信息出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

var baseRateTableName string = "RPM_PAR_BASE_RATE"

var baseRateTabales string = `
	RPM_PAR_BASE_RATE T
	`

var baseRateCols map[string]string = map[string]string{
	"T.UUID":           "' '",
	"T.BASE_RATE_TYPE": "' '",
	"T.TERM":           "0",
	"T.BASE_RATE_NAME": "' '",
	"T.RATE":           "0",
	"T.FLAG":           "' '",
	"T.CREATE_TIME":    "sysdate",
	"T.UPDATE_TIME":    "sysdate",
	"T.CREATE_USER":    "' '",
	"T.UPDATE_USER":    "' '",
}

var baseRateColsSort = []string{
	"T.UUID",
	"T.BASE_RATE_TYPE",
	"T.TERM",
	"T.BASE_RATE_NAME",
	"T.RATE",
	"T.FLAG",
	"T.CREATE_TIME",
	"T.UPDATE_TIME",
	"T.CREATE_USER",
	"T.UPDATE_USER",
}
