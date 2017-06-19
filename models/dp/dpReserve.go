package dp

import (
	"database/sql"
	"fmt"
	"time"

	modelsUtil "pccqcpa.com.cn/app/rpm/api/models/util"
	"pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/components/zlog"
)

// 结构体对应数据库表【RPM_PAR_DP_RESERVE】存款准备金
// by author Jason
// by time 2016-12-06 17:00:38
type DpReserve struct {
	UUID       string    // 主键 默认值 SYS_GUID()
	Reserve    float64   // 保证金百分比
	ParamType  string    // 参数类型
	Flag       string    // 生效标志（1：生效 0:失效）
	CreateTime time.Time // 创建时间
	CreateUser string    // 创建人
	UpdateTime time.Time // 更新时间
	UpdateUser string    // 更新人
}

// 结构体scan
// by author Jason
// by time 2016-12-06 17:00:38
func (this DpReserve) scan(rows *sql.Rows) (*DpReserve, error) {
	var one = new(DpReserve)
	values := []interface{}{
		&one.UUID,
		&one.Reserve,
		&one.ParamType,
		&one.Flag,
		&one.CreateTime,
		&one.CreateUser,
		&one.UpdateTime,
		&one.UpdateUser,
	}
	err := util.OracleScan(rows, values)
	return one, err
}

// 分页操作
// by author Jason
// by time 2016-12-06 17:00:38
func (this DpReserve) List(param ...map[string]interface{}) (*util.PageData, error) {
	pageData, rows, err := modelsUtil.List(dpReserveTabales, dpReserveCols, dpReserveColsSort, param...)
	defer rows.Close()
	if nil != err {
		er := fmt.Errorf("分页查询存款准备金信息出错")
		zlog.Error(er.Error(), err)
		return nil, er
	}

	var rst []*DpReserve
	for rows.Next() {
		one, err := this.scan(rows)
		if nil != err {
			er := fmt.Errorf("分页查询存款准备金信息row.Scan()出错")
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
// by time 2016-12-06 17:00:38
func (this DpReserve) Find(param ...map[string]interface{}) ([]*DpReserve, error) {
	rows, err := modelsUtil.FindRows(dpReserveTabales, dpReserveCols, dpReserveColsSort, param...)
	defer rows.Close()
	if nil != err {
		er := fmt.Errorf("多参数查询存款准备金信息出错")
		zlog.Error(er.Error(), err)
		return nil, er
	}
	var rst []*DpReserve
	for rows.Next() {
		one, err := this.scan(rows)
		if nil != err {
			er := fmt.Errorf("多参数查询存款准备金信息row.Scan()出错")
			zlog.Error(er.Error(), err)
			return nil, er
		}
		rst = append(rst, one)
	}
	return rst, nil
}

// 新增存款准备金信息
// by author Jason
// by time 2016-12-06 17:00:38
func (this DpReserve) Add() error {
	paramMap := map[string]interface{}{
		"reserve":     this.Reserve,
		"param_type":  this.ParamType,
		"flag":        this.Flag,
		"create_time": util.GetCurrentTime(),
		"create_user": this.CreateUser,
		"update_time": util.GetCurrentTime(),
		"update_user": this.UpdateUser,
	}
	err := util.OracleAdd("RPM_PAR_DP_RESERVE", paramMap)
	if nil != err {
		er := fmt.Errorf("新增存款准备金信息出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

// 更新存款准备金信息
// by author Jason
// by time 2016-12-06 17:00:38
func (this DpReserve) Update() error {
	paramMap := map[string]interface{}{
		"reserve":     this.Reserve,
		"param_type":  this.ParamType,
		"flag":        this.Flag,
		"update_time": util.GetCurrentTime(),
		"update_user": this.UpdateUser,
	}
	whereParam := map[string]interface{}{
		"uuid": this.UUID,
	}
	err := util.OracleUpdate("RPM_PAR_DP_RESERVE", paramMap, whereParam)
	if nil != err {
		er := fmt.Errorf("更新存款准备金信息出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

// 删除存款准备金信息
// by author Jason
// by time 2016-12-06 17:00:38
func (this DpReserve) Delete() error {
	whereParam := map[string]interface{}{
		"uuid": this.UUID,
	}
	err := util.OracleDelete("RPM_PAR_DP_RESERVE", whereParam)
	if nil != err {
		er := fmt.Errorf("删除存款准备金信息出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

var dpReserveTabales string = `RPM_PAR_DP_RESERVE T`

var dpReserveCols map[string]string = map[string]string{
	"T.UUID":        "' '",
	"T.RESERVE":     "0",
	"T.PARAM_TYPE":  "' '",
	"T.FLAG":        "' '",
	"T.CREATE_TIME": "sysdate",
	"T.CREATE_USER": "' '",
	"T.UPDATE_TIME": "sysdate",
	"T.UPDATE_USER": "' '",
}

var dpReserveColsSort = []string{
	"T.UUID",
	"T.RESERVE",
	"T.PARAM_TYPE",
	"T.FLAG",
	"T.CREATE_TIME",
	"T.CREATE_USER",
	"T.UPDATE_TIME",
	"T.UPDATE_USER",
}
