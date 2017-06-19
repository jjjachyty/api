package util

import (
	"database/sql"
	"fmt"
	"platform/dbobj"
	"strconv"
	"strings"

	"pccqcpa.com.cn/components/zlog"
)

func OracleScan(rows *sql.Rows, values []interface{}) error {
	err := rows.Scan(values...)
	if nil != err {
		zlog.Error("查询数据rows.Scan()错误", err)
		return err
	}
	return nil
}

func OracleExec(sql string, values []interface{}) error {
	var err error
	if nil == values {
		err = dbobj.Default.Exec(sql)
	} else {
		err = dbobj.Default.Exec(sql, values...)
	}

	zlog.Infof("执行SQL：\nsql:[%s]\n参数:[%v]", nil, sql, values)

	if nil != err {
		zlog.Errorf("执行Sql失败：\nsql:[%v]\n参数:[%v]", err, sql, values)
		return err
	}
	return nil
}

// 新增数据
func OracleAdd(tableName string, param map[string]interface{}) error {
	var i = 1
	sql := `insert into ` + tableName + `(uuid `
	valueSql := ` values(sys_guid() `
	var values []interface{}
	for k, v := range param {
		if "uuid" == strings.ToLower(k) {
			continue
		}
		sql += `,` + k
		num := strconv.Itoa(i)
		switch strings.ToLower(k) {
		case "create_time", "update_time":
			valueSql += `, to_date(:` + num + `, 'YYYY-MM-DD hh24:mi:ss')`
		default:
			valueSql += `, :` + num
		}
		values = append(values, v)
		i++
	}
	sql += `) `
	valueSql += `)`
	err := OracleExec(sql+valueSql, values)
	zlog.Infof("新增表[%v]的数据:\nsql:[%v]\nvalues:[%v]", err, tableName, sql+valueSql, values)
	if nil != err {
		zlog.Errorf("新增表[%v]的数据失败:\nsql:[%v]\nvalues:[%v]", err, tableName, sql+valueSql, values)
	}
	return err
}

// 新增数据带事务
func OracleBatchAdd(tableName string, param map[string]interface{}, sqlTx *sql.Tx) (sql.Result, error) {
	var i = 1
	sql := `insert into ` + tableName + `(uuid `
	valueSql := ` values(sys_guid() `
	var values []interface{}
	for k, v := range param {
		if "uuid" == strings.ToLower(k) {
			continue
		}
		sql += `,` + k
		num := strconv.Itoa(i)
		switch strings.ToLower(k) {
		case "create_time", "update_time":
			valueSql += `, to_date(:` + num + `, 'YYYY-MM-DD hh24:mi:ss')`
		default:
			valueSql += `, :` + num
		}
		values = append(values, v)
		i++
	}
	sql += `) `
	valueSql += `)`
	result, err := sqlTx.Exec(sql+valueSql, values...)
	if nil != err {
		zlog.Errorf("新增表[%v]的数据失败:\nsql:[%v]\nvalues:[%v]", err, tableName, sql+valueSql, values)
	}
	return result, err
}

// 执行存储过程
func OracleExecProcedure(procedureName string, param ...interface{}) error {
	var procedureSql string = "call " + procedureName + "("
	for i := 0; i < len(param); i++ {
		switch i {
		case 0:
			procedureSql += ":" + strconv.Itoa(i+1)
		default:
			procedureSql += ", :" + strconv.Itoa(i+1)
		}
	}
	procedureSql += ")"
	err := OracleExec(procedureSql, param)
	return err
}

func OracleBatchStmtAdd(tableName string, param map[string]interface{}, sqlTx *sql.Tx) (sql.Result, error) {
	var i = 1
	sql := `insert into ` + tableName + `(`
	valueSql := ` values(`
	var values []interface{}
	for k, v := range param {
		if "uuid" == strings.ToLower(k) {
			sql += k + `,`
			valueSql += `sys_guid(), `
			continue
		}
		sql += k + `,`
		num := strconv.Itoa(i)
		switch strings.ToLower(k) {
		case "create_time", "update_time":
			valueSql += `, to_date(:` + num + `, 'YYYY-MM-DD hh24:mi:ss')`
		default:
			valueSql += `:` + num + `, `
		}
		values = append(values, v)
		i++
	}
	sql = sql[:len(sql)-1] + `) `
	valueSql = valueSql[:len(valueSql)-2] + `)`
	stmt, err := sqlTx.Prepare(sql + valueSql)
	if nil != err {
		zlog.Error("数据导入,获取Stmt失败", nil)
		return nil, err
	}
	result, err := stmt.Exec(values...)
	if nil != err {
		zlog.Errorf("新增表[%v]的数据失败:\nsql:[%v]\nvalues:[%v]", err, tableName, sql+valueSql, values)
	}
	// zlog.Errorf("sql【%s】\n参数【%s】",nil, )
	zlog.Infof("\nsql:[%v]\nvalues:[%v]", err, sql+valueSql, values)
	return result, err
}

// 更新数据
func OracleUpdate(tableName string, param map[string]interface{}, whereParam map[string]interface{}) error {
	var i = 1
	sql := `update ` + tableName + ` set `
	whereSql := ` where 1=1 `
	var values []interface{}
	for k, v := range param {
		num := strconv.Itoa(i)
		switch strings.ToLower(k) {
		case "create_time", "update_time":
			sql += k + ` = to_date(:` + num + `, 'YYYY-MM-DD hh24:mi:ss'),`
		default:
			sql += k + ` = :` + num + `,`
		}
		values = append(values, v)
		i++
	}
	for k, v := range whereParam {
		whereSql += ` and ` + k + ` = :` + strconv.Itoa(i)
		i++
		values = append(values, v)
	}
	sql = strings.TrimRight(sql, ",")
	err := OracleExec(sql+whereSql, values)

	zlog.Infof("更新表[%v]的数据失败:\nsql:[%v]\nvalues:[%v]", err, tableName, sql+whereSql, values)

	if nil != err {
		zlog.Errorf("更新表[%v]的数据失败:\nsql:[%v]\nvalues:[%v]", err, tableName, sql+whereSql, values)
	}
	return err
}

// 删除数据
func OracleDelete(tableName string, wehreParam map[string]interface{}) error {
	sql := `delete from ` + tableName + ` where 1=1 `
	var i = 1
	var values []interface{}
	for k, v := range wehreParam {
		sql += ` and ` + k + ` = :` + strconv.Itoa(i)
		values = append(values, v)
		i++
	}
	err := OracleExec(sql, values)
	if nil != err {
		zlog.Errorf("删除表[%v]的数据失败:\nsql:[%v]\nvalues:[%v]", err, tableName, sql, values)
	}
	return err
}

func OracleQuery(sql string, param ...interface{}) (*sql.Rows, error) {
	rows, err := dbobj.Default.Query(sql, param...)

	if nil != err {
		er := fmt.Errorf("执行sql查询出错", sql, param)
		zlog.Errorf(er.Error()+",sql[%v],param[%#v]", err, sql, param)
		return nil, er
	}
	return rows, nil
}
