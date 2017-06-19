package util

import (
	"database/sql"
	"fmt"
	"platform/dbobj"
	"strconv"
	"strings"
	"time"

	"pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/components/zlog"
)

//分页查询
//参数1 前台传送过来的信息 包括 page：当前页 count：每页的条数 searchLike：模糊查询
//参数2 相关表 exp sys_sec_user t left join sys_sec_role t1 on (? = ?)
//参数3 key：列名 value：如果为空需要替换成的值 exp：｛"uuid":" ","crate_time":"sysdate"｝
func List(tables string, cols map[string]string, colsSort []string, param ...map[string]interface{}) (*util.PageData, *sql.Rows, error) {
	// 处理panic
	defer func() {
		if err := recover(); nil != err {
			fmt.Println("err", err)
		}
	}()

	var paramMap map[string]interface{}
	if 0 == len(param) {
		paramMap = make(map[string]interface{})
	} else {
		paramMap = param[0]
	}
	zlog.Info("开始分页查询", nil)
	var tableAlias string = ""

	//拼接sql
	selectSqlOut := `select `                         //外层查询sql
	selectSqlMid := `select t.*,rownum as num from (` //分页查询sql
	selectSqlIn := ` select  `
	for _, v := range colsSort {
		str := strings.SplitAfter(v, ".")
		if "uuid" == strings.ToLower(v) || (2 == len(str) && "uuid" == str[1]) {
			selectSqlOut += ` uuid, `
			if 2 == len(str) {
				tableAlias = str[0]
			}
		} else if 2 == len(str) {
			selectSqlOut += `coalesce(` + str[1] + `,` + cols[v] + `),`
		} else if 1 == len(str) {
			selectSqlOut += `coalesce(` + v + `,` + cols[v] + `),`
		}
		selectSqlIn += ` ` + v + `, `
	}
	selectSqlIn = strings.TrimRight(selectSqlIn, ", ")
	selectSqlOut = strings.TrimRight(selectSqlOut, ",")
	selectSqlOut += ` from (`

	//获取分页数据
	start, length := getPageCount(paramMap)

	whereSqlMid := ") t where rownum <= " + strconv.Itoa(start+length)
	whereSqlOut := ") where num > " + strconv.Itoa(start)

	//排序方式
	orderBy := getOrderBy(tableAlias, paramMap)
	delete(paramMap, "sort")
	delete(paramMap, "order")

	//拼接where条件
	whereSqlIn, params := getWhereSql(0, paramMap)
	whereSqlIn = handleWhereSql(whereSqlIn)

	if strings.Contains(strings.ToLower(tables), " where ") || "" == strings.TrimSpace(whereSqlIn) {
		whereSqlIn = " from " + tables + whereSqlIn
	} else {
		whereSqlIn = " from " + tables + " where " + whereSqlIn

	}

	// whereSqlIn = " from " + tables + whereSqlIn
	sql := selectSqlOut + selectSqlMid + selectSqlIn + whereSqlIn + orderBy + whereSqlMid + whereSqlOut

	rows, err := getPageRow(sql, params)
	if nil != err {
		return nil, nil, err
	}
	pageData := getpageData(whereSqlIn, params)
	return pageData, rows, nil
}

func handleWhereSql(whereSqlIn string) string {
	whereSqlIn = strings.TrimLeft(whereSqlIn, " ")

	whereSqlIn = strings.TrimPrefix(whereSqlIn, "or")
	whereSqlIn = strings.TrimPrefix(whereSqlIn, "and")
	return whereSqlIn
}

func getPageRow(sql string, params []interface{}) (*sql.Rows, error) {

	zlog.Infof("\n分页查询sql：%s\n参数\n%s", nil, sql, params)
	//查询sql
	rows, err := dbobj.Default.Query(sql, params...)
	if nil != err {
		defer rows.Close()
		zlog.Infof("\n分页查询sql：%s\n参数\n%s", nil, sql, params)
		zlog.Error("分页查询失败", err)
		return nil, err
	}
	return rows, nil
}

func getpageData(whereSqlIn string, params []interface{}) *util.PageData {
	//初始化返回结果
	var pageData = new(util.PageData)
	//查询总条数
	var total int64
	countSql := `select count(1) `
	sql := countSql + whereSqlIn
	countRows, err := dbobj.Default.Query(sql, params...)
	defer countRows.Close()
	if nil != err {
		zlog.Errorf("分页查询总数出错\nsql: %v\n参数\n%v", err, sql, params)
		return nil
	}
	for countRows.Next() {
		err = countRows.Scan(&total)
		countRows.Next()
		if nil != err {
			zlog.Error("查询分页总数row.Scan()报错", err)
			return nil
		}
	}
	pageData.Page.TotalRows = total
	return pageData
}

//处理分页数据
func getPageCount(paramMap map[string]interface{}) (int, int) {
	var start, length int
	_, ok := paramMap["start"]
	if ok {
		start, _ = strconv.Atoi(paramMap["start"].(string))
	} else {
		start = 0
	}

	_, ok = paramMap["length"]
	if ok {
		length, _ = strconv.Atoi(paramMap["length"].(string))
	} else {
		length = 20
	}

	delete(paramMap, "start")
	delete(paramMap, "length")
	return start, length
}

//获取排序方式
func getOrderBy(tableAlias string, paramMap map[string]interface{}) string {
	orderBy := ""
	sort, ok := paramMap["sort"]
	// fmt.Println("－－－－－－－－", ok, sort)
	if ok && "" != sort {
		var sorts, orders []string
		var index int = 0
		if order, ok := paramMap["order"]; ok {
			orders = strings.Split(order.(string), ",")
		} else {
			orders = append(orders, "ASC")
		}

		sorts = strings.Split(sort.(string), ",")

		// 判断是否为关联表排序
		for _, str := range sorts {
			var orderStrings = strings.Split(str, ".")
			var orderStringsLength = len(orderStrings)
			if 1 < orderStringsLength {
				orderBy += orderStrings[orderStringsLength-2] + "." + util.UperChange(orderStrings[orderStringsLength-1]) + " " + orders[index] + ", "
			} else {
				orderBy += util.UperChange(sort.(string)) + " " + orders[index] + ", "
			}
			if index < len(orders)-1 {
				index++
			}

		}
		orderBy = " order by " + orderBy
	} else {
		//默认排序方式
		orderBy += " ORDER BY " + tableAlias + "CREATE_TIME, "
		paramMap["order"] = "DESC"
	}
	order, ok := paramMap["order"]
	if ok {
		orderBy += tableAlias + "UUID " + order.(string)
	} else {
		orderBy += tableAlias + "UUID "
	}
	return orderBy
}

//拼接where条件
func getWhereSql(paramLen int, paramMap map[string]interface{}) (string, []interface{}) {

	whereSql := " "
	var i int = 1 //用作替换符
	if paramLen > 0 {
		i = paramLen + 1
	}
	var params = make([]interface{}, 0)
	for k, v := range paramMap {
		// fmt.Println("－－－－－－遍历map获取whereSql－－－－－", k, v, whereSql)
		if util.SEARCH_LIKE == k {
			if nil == v {
				continue
			}
			for _, v := range paramMap[util.SEARCH_LIKE].([]map[string]interface{}) {
				var key string
				//判断值是否为空，如果为空的话就不拼接sql
				if "or" != v["type"].(string) {
					//处理列名
					key = v["key"].(string)
					var nullInterface = interface{}("")
					if nil == v["value"] || nullInterface == v["value"] {
						continue
					}
				}

				colName := util.UperChange(key)

				switch v["type"].(string) {
				case "like":
					whereSql += ` and ` + handleColName(colName) + ` like '%%` + v["value"].(string) + `%%'`
					continue
				case "eq":
					// whereSql += ` and ` + colName + ` = '` + v["value"].(string) + `'`
					whereSql += ` and ` + handleColName(colName) + ` = :` + strconv.Itoa(i)
					params = append(params, v["value"].(string))
					i++
					continue
				case "gt":
					// whereSql += ` and ` + colName + ` > ` + v["value"].(string)
					whereSql += ` and ` + handleColName(colName) + ` > :` + strconv.Itoa(i)
					params = append(params, v["value"])
					i++
					continue
				case "lt":
					// whereSql += ` and ` + colName + ` < ` + v["value"].(string)
					whereSql += ` and ` + handleColName(colName) + ` < :` + strconv.Itoa(i)
					params = append(params, v["value"].(float64))
					i++
					continue
				case "ge":
					// whereSql += ` and ` + colName + ` >= ` + v["value"].(string)
					whereSql += ` and ` + handleColName(colName) + ` >= :` + strconv.Itoa(i)
					params = append(params, v["value"].(float64))
					i++
					continue
				case "le":
					// whereSql += ` and ` + colName + ` <= ` + v["value"].(string)
					whereSql += ` and ` + handleColName(colName) + ` <= :` + strconv.Itoa(i)
					params = append(params, v["value"])
					i++
					continue
				case "in":
					whereSql += ` and ` + handleColName(colName) + ` in (` + v["value"].(string) + `)`
					// whereSql += ` and ` + colName + ` in (:` + strconv.Itoa(i) + `)`
					// params = append(params, v["value"].(string))
					// i++
					continue
				case "or":
					keys := v["key"].([]interface{})
					values := v["value"].([]interface{})
					for j, v := range keys {
						key := util.UperChange(v.(string))
						if nil == values[j] {
							continue
						}
						value := fmt.Sprint(values[j])
						if j == 0 {
							whereSql += ` and (`
						} else {
							whereSql += ` or `
						}
						whereSql += `  upper(` + handleColName(key) + `) like upper('%%` + value + `%%')`
						if len(keys)-1 == j {
							whereSql += ` )`
							continue
						}
					}
					continue
				case "<>":
					whereSql += ` and ` + handleColName(colName) + ` <> :` + strconv.Itoa(i)
					params = append(params, v["value"].(string))
					i++
					continue

				}
			}
			// delete(paramMap, "searchLike")
			continue
		}

		// whereSql += ` and ` + k + ` = '` + v.(string) + `'`
		switch v.(type) {
		case time.Time:
			whereSql += ` and to_char(` + handleColName(k) + `,'yyyy-mm-dd') = :` + strconv.Itoa(i)
			params = append(params, v.(time.Time).Format("2006-01-02"))
		default:
			whereSql += ` and ` + handleColName(k) + ` = :` + strconv.Itoa(i)
			params = append(params, fmt.Sprint(v))
		}
		i++
	}
	return whereSql, params
}

func handleColName(colName string) string {
	zlog.Info("colName"+colName, nil)
	if "START_TIME" == strings.ToUpper(strings.Split(colName, ".")[len(strings.Split(colName, "."))-1]) {
		return "to_date(to_char(t.start_time,'yyyy-mm-dd'),'yyyy-mm-dd')"
	}
	if 1 == len(strings.Split(colName, ".")) {
		return "T." + colName
	}
	return colName
}
