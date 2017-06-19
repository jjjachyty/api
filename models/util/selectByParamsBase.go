package util

import (
	"database/sql"
	"fmt"
	"platform/dbobj"
	"strings"

	"pccqcpa.com.cn/components/zlog"
)

// 多参数查询
// 参数1 相关表 exp sys_sec_user t left join sys_sec_role t1 on (? = ?)
// 参数2 key：列名 value：如果为空需要替换成的值 exp：｛"uuid":" ","crate_time":"sysdate"｝
// 参数3 列排序
// 参数4 前台传送过来的信息 包括 searchLike：模糊查询

func FindRows(tables string, cols map[string]string, colsSort []string, param ...map[string]interface{}) (*sql.Rows, error) {
	// 处理panic
	defer func() {
		if err := recover(); nil != err {
			fmt.Println("err", err)
		}
	}()

	// fmt.Println("paramMap", param)
	var paramMap map[string]interface{}
	if 0 == len(param) {
		paramMap = make(map[string]interface{})
	} else {
		paramMap = param[0]
	}
	// 获取主标别名
	var tableAlias string = ""
	var uuid = colsSort[0]
	uuids := strings.Split(uuid, ".")
	if 1 < len(uuids) {
		tableAlias = uuids[0] + "."
	}

	//拼接sql
	selectSqlOut := `select ` //外层查询sql
	for _, v := range colsSort {
		selectSqlOut += ` coalesce(` + v + `,` + cols[v] + `),`
	}
	selectSqlOut = strings.TrimRight(selectSqlOut, ",")
	selectSqlOut += ` from ` + tables + ` where 1=1 `

	//排序方式
	orderBy := getOrderBy(tableAlias, paramMap)
	delete(paramMap, "sort")
	delete(paramMap, "order")

	//拼接where条件
	whereSql, params := getWhereSql(0, paramMap)
	sql := selectSqlOut + whereSql + orderBy
	zlog.Infof("\nsql：%s\n参数\n%s", nil, sql, params)
	//查询sql
	// zlog.Debugf("SQL:%s", nil, sql)
	// zlog.Debugf("Args:%v", nil, params...)
	rows, err := dbobj.Default.Query(sql, params...)
	if nil != err {
		zlog.Infof("\nsql：%s\n参数\n%s", nil, sql, params)
	}
	return rows, err
}
