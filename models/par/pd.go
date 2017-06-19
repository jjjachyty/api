package par

import (
	"database/sql"
	"time"

	"fmt"
	modelsUtil "pccqcpa.com.cn/app/rpm/api/models/util"
	"pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/components/zlog"
)

// 结构体对应数据库表【RPM_PAR_PD】违约概率
/**
 * @apiDefine Pd
 * @apiSuccess {string}     	UUID            主键默认值sys_guid()
 * @apiSuccess {string}   		CustCredit    	信用等级
 * @apiSuccess {float64}   		PdRate      	违约概率
 * @apiSuccess {string}   		Flag   			生效标志
 * @apiSuccess {time.Time }   	StartTime    	生效日期
 * @apiSuccess {time.Time}		CreateTime      创建时间
 * @apiSuccess {string}   		CreateUser      创建人
 * @apiSuccess {time.Time}		UpdateTime      更新时间
 * @apiSuccess {string}   		UpdateUser      更新人
 */
type Pd struct {
	UUID       string    //主键
	CustCredit string    //信用等级
	PdRate     float64   //违约概率
	Flag       string    //生效标志
	StartTime  time.Time //生效日期
	CreateTime time.Time //创建日期
	CreateUser string    //创建人
	UpdateTime time.Time //更新时间
	UpdateUser string    //更新人
}

var pdTables string = `
	RPM_PAR_PD T
`

var pdCols = map[string]string{
	"T.UUID":        "' '",
	"T.CUST_CREDIT": "' '",
	"T.PD_RATE":     "0",
	"T.FLAG":        "' '",
	"T.START_TIME":  "sysdate",
	"T.CREATE_TIME": "sysdate",
	"T.CREATE_USER": "' '",
	"T.UPDATE_TIME": "sysdate",
	"T.UPDATE_USER": "' '",
}

var pdColsSort = []string{
	"T.UUID",
	"T.CUST_CREDIT",
	"T.PD_RATE",
	"T.FLAG",
	"T.START_TIME",
	"T.CREATE_TIME",
	"T.CREATE_USER",
	"T.UPDATE_TIME",
	"T.UPDATE_USER",
}

//结构体scan
func (this *Pd) scanPd(rows *sql.Rows) (*Pd, error) {
	var one = new(Pd)
	var values = []interface{}{
		&one.UUID,
		&one.CustCredit,
		&one.PdRate,
		&one.Flag,
		&one.StartTime,
		&one.CreateTime,
		&one.CreateUser,
		&one.UpdateTime,
		&one.UpdateUser,
	}
	err := util.OracleScan(rows, values)
	if nil != err {
		return nil, err
	}
	return one, nil
}

func (this *Pd) SelectPdByParams(paramMap map[string]interface{}) ([]*Pd, error) {
	rows, err := modelsUtil.FindRows(pdTables, pdCols, pdColsSort, paramMap)
	if nil != err {
		return nil, err
	}
	defer rows.Close()
	var pds []*Pd
	for rows.Next() {
		pd, err := this.scanPd(rows)
		if nil != err {
			return nil, err
		}
		pds = append(pds, pd)
	}
	return pds, nil
}

// 分页操作
// by author Yeqc
// by time 2016-12-21 09:52:26
func (this Pd) List(param ...map[string]interface{}) (*util.PageData, error) {
	pageData, rows, err := modelsUtil.List(pdTables, pdCols, pdColsSort, param...)
	defer rows.Close()
	if nil != err {
		er := fmt.Errorf("分页查询违约概率信息出错")
		zlog.Error(er.Error(), err)
		return nil, er
	}

	var rst []*Pd
	for rows.Next() {
		one, err := this.scanPd(rows)
		if nil != err {
			er := fmt.Errorf("分页查询违约概率信息row.Scan()出错")
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
// by time 2016-12-21 09:52:26
func (this Pd) Find(param ...map[string]interface{}) ([]*Pd, error) {
	rows, err := modelsUtil.FindRows(pdTables, pdCols, pdColsSort, param...)
	defer rows.Close()
	if nil != err {
		er := fmt.Errorf("多参数查询违约概率信息出错")
		zlog.Error(er.Error(), err)
		return nil, er
	}
	var rst []*Pd
	for rows.Next() {
		one, err := this.scanPd(rows)
		if nil != err {
			er := fmt.Errorf("多参数查询违约概率信息row.Scan()出错")
			zlog.Error(er.Error(), err)
			return nil, er
		}
		rst = append(rst, one)
	}
	return rst, nil
}

// 新增违约概率信息
// by author Yeqc
// by time 2016-12-21 09:52:26
func (this Pd) Add() error {
	paramMap := map[string]interface{}{
		"cust_credit": this.CustCredit,
		"pd_rate":     this.PdRate,
		"flag":        this.Flag,
		"start_time":  this.StartTime,
		"create_time": util.GetCurrentTime(),
		"create_user": this.CreateUser,
		"update_time": util.GetCurrentTime(),
		"update_user": this.UpdateUser,
	}
	err := util.OracleAdd("RPM_PAR_PD", paramMap)
	if nil != err {
		er := fmt.Errorf("新增违约概率信息出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

// 更新违约概率信息
// by author Yeqc
// by time 2016-12-21 09:52:26
func (this Pd) Update() error {
	paramMap := map[string]interface{}{
		"cust_credit": this.CustCredit,
		"pd_rate":     this.PdRate,
		"flag":        this.Flag,
		"start_time":  this.StartTime,
		"update_time": util.GetCurrentTime(),
		"update_user": this.UpdateUser,
	}
	whereParam := map[string]interface{}{
		"uuid": this.UUID,
	}
	err := util.OracleUpdate("RPM_PAR_PD", paramMap, whereParam)
	if nil != err {
		er := fmt.Errorf("更新违约概率信息出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

// 删除违约概率信息
// by author Yeqc
// by time 2016-12-21 09:52:26
func (this Pd) Delete() error {
	whereParam := map[string]interface{}{
		"uuid": this.UUID,
	}
	err := util.OracleDelete("RPM_PAR_PD", whereParam)
	if nil != err {
		er := fmt.Errorf("删除违约概率信息出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}
