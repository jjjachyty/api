package par

import (
	"database/sql"
	"fmt"

	modelsUtil "pccqcpa.com.cn/app/rpm/api/models/util"
	"pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/components/zlog"
	"time"
)

// 结构体对应数据库表【RPM_PAR_COMMON】通用参数设置
// by author Yeqc
// by time 2016-12-20 10:58:09
/**
 * @apiDefine Common
 * @apiSuccess {string}     UUID            主键默认值sys_guid()
 * @apiSuccess {string}   	ParMod    		参数模块
 * @apiSuccess {string}   	ParKey      	参数主键
 * @apiSuccess {string}   	ParValue   		参数值
 * @apiSuccess {string}   	Explain    		解释说明
 * @apiSuccess {string}   	Flag      		生效标志位
 * @apiSuccess {float64}   	Sort      		排序
 * @apiSuccess {time.Time}	CreateTime      创建时间
 * @apiSuccess {string}   	CreateUser      创建人
 * @apiSuccess {time.Time}	UpdateTime      更新时间
 * @apiSuccess {string}   	UpdateUser      更新人
 */
type Common struct {
	UUID       string    // 主键
	ParMod     string    // 参数模块
	ParKey     string    // 参数主键
	ParValue   string    // 参数值
	Explain    string    // 解释说明
	Flag       string    // 生效标志位
	CreateTime time.Time // 创建时间
	CreateUser string    // 创建人
	UpdateTime time.Time // 更新时间
	UpdateUser string    // 更新人
	Sort       float64   // 排序
}

// 结构体scan
// by author Yeqc
// by time 2016-12-20 10:58:09
func (this Common) scan(rows *sql.Rows) (*Common, error) {
	var one = new(Common)
	values := []interface{}{
		&one.UUID,
		&one.ParMod,
		&one.ParKey,
		&one.ParValue,
		&one.Explain,
		&one.Flag,
		&one.CreateTime,
		&one.CreateUser,
		&one.UpdateTime,
		&one.UpdateUser,
		&one.Sort,
	}
	err := util.OracleScan(rows, values)
	return one, err
}

// 分页操作
// by author Yeqc
// by time 2016-12-20 10:58:09
func (this Common) List(param ...map[string]interface{}) (*util.PageData, error) {
	pageData, rows, err := modelsUtil.List(commonTabales, commonCols, commonColsSort, param...)
	defer rows.Close()
	if nil != err {
		er := fmt.Errorf("分页查询通用参数设置信息出错")
		zlog.Error(er.Error(), err)
		return nil, er
	}

	var rst []*Common
	for rows.Next() {
		one, err := this.scan(rows)
		if nil != err {
			er := fmt.Errorf("分页查询通用参数设置信息row.Scan()出错")
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
// by time 2016-12-20 10:58:09
func (this Common) Find(param ...map[string]interface{}) ([]*Common, error) {
	rows, err := modelsUtil.FindRows(commonTabales, commonCols, commonColsSort, param...)
	defer rows.Close()
	if nil != err {
		er := fmt.Errorf("多参数查询通用参数设置信息出错")
		zlog.Error(er.Error(), err)
		return nil, er
	}
	var rst []*Common
	for rows.Next() {
		one, err := this.scan(rows)
		if nil != err {
			er := fmt.Errorf("多参数查询通用参数设置信息row.Scan()出错")
			zlog.Error(er.Error(), err)
			return nil, er
		}
		rst = append(rst, one)
	}
	return rst, nil
}

// 新增通用参数设置信息
// by author Yeqc
// by time 2016-12-20 10:58:09
func (this Common) Add() error {
	paramMap := map[string]interface{}{
		"par_mod":     this.ParMod,
		"par_key":     this.ParKey,
		"par_value":   this.ParValue,
		"explain":     this.Explain,
		"flag":        this.Flag,
		"create_time": util.GetCurrentTime(),
		"create_user": this.CreateUser,
		"update_time": util.GetCurrentTime(),
		"update_user": this.UpdateUser,
		"sort":        this.Sort,
	}
	err := util.OracleAdd("RPM_PAR_COMMON", paramMap)
	if nil != err {
		er := fmt.Errorf("新增通用参数设置信息出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

// 更新通用参数设置信息
// by author Yeqc
// by time 2016-12-20 10:58:09
func (this Common) Update() error {
	paramMap := map[string]interface{}{
		"par_mod":     this.ParMod,
		"par_key":     this.ParKey,
		"par_value":   this.ParValue,
		"explain":     this.Explain,
		"flag":        this.Flag,
		"update_time": util.GetCurrentTime(),
		"update_user": this.UpdateUser,
		"sort":        this.Sort,
	}
	whereParam := map[string]interface{}{
		"uuid": this.UUID,
	}
	err := util.OracleUpdate("RPM_PAR_COMMON", paramMap, whereParam)
	if nil != err {
		er := fmt.Errorf("更新通用参数设置信息出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

// 删除通用参数设置信息
// by author Yeqc
// by time 2016-12-20 10:58:09
func (this Common) Delete() error {
	whereParam := map[string]interface{}{
		"uuid": this.UUID,
	}
	err := util.OracleDelete("RPM_PAR_COMMON", whereParam)
	if nil != err {
		er := fmt.Errorf("删除通用参数设置信息出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

var commonTabales string = `RPM_PAR_COMMON T`

var commonCols map[string]string = map[string]string{
	"T.UUID":        "' '",
	"T.PAR_MOD":     "' '",
	"T.PAR_KEY":     "' '",
	"T.PAR_VALUE":   "' '",
	"T.EXPLAIN":     "' '",
	"T.FLAG":        "' '",
	"T.CREATE_TIME": "sysdate",
	"T.CREATE_USER": "' '",
	"T.UPDATE_TIME": "sysdate",
	"T.UPDATE_USER": "' '",
	"T.SORT":        "0",
}

var commonColsSort = []string{
	"T.UUID",
	"T.PAR_MOD",
	"T.PAR_KEY",
	"T.PAR_VALUE",
	"T.EXPLAIN",
	"T.FLAG",
	"T.CREATE_TIME",
	"T.CREATE_USER",
	"T.UPDATE_TIME",
	"T.UPDATE_USER",
	"T.SORT",
}
