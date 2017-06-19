package util

import (
	"database/sql"
	// "fmt"
	"strconv"
	"strings"

	apiUtil "pccqcpa.com.cn/app/rpm/api/util"
)

type modelUtil struct {
	ParamMap    map[string]interface{} //传入参数
	tables      string
	colSort     []string
	cols        map[string]string
	sqlParamMap []interface{} // sql 参数
}

func (m modelUtil) GetWhereSql(paramLen ...int) (string, []interface{}) {
	var whereSql string
	var params []interface{}
	if nil != paramLen {
		whereSql, params = getWhereSql(paramLen[0], m.ParamMap)
	} else {
		whereSql, params = getWhereSql(len(m.sqlParamMap), m.ParamMap)
	}
	whereSql = handleWhereSql(whereSql)
	return whereSql, params
}

func NewModelUtil(param ...map[string]interface{}) *modelUtil {
	if 0 < len(param) {
		return &modelUtil{
			ParamMap: param[0],
		}
	} else {
		return &modelUtil{}
	}
}

func (m *modelUtil) SetTableMsg(tables string, cols map[string]string, colSort []string) {
	m.cols = cols
	m.colSort = colSort
	m.tables = tables
}

// 拼接查询字符串
func (m modelUtil) UnionAll(sqls ...string) (unionSql string) {
	for index, sql := range sqls {
		if 0 == index {
			unionSql += sql
		} else {
			unionSql += " UNION " + sql
		}
	}
	return
}

// 获取查询所需列
func (m modelUtil) GetSelectSql() (selectSql string) {
	selectSql = "SELECT "
	length := len(m.colSort)
	for index, col := range m.colSort {
		switch index {
		case length - 1:
			selectSql += strings.ToUpper(col) + " "
		default:
			selectSql += strings.ToUpper(col) + ","
		}
	}
	selectSql += " FROM " + m.tables + " "
	return
}

// 拼接select 和 where
func (m *modelUtil) GetSelectWhereSql() (string, []interface{}) {
	whereSql, param := m.GetWhereSql()
	m.sqlParamMap = append(m.sqlParamMap, param...)
	return m.GetSelectSql() + " WHERE " + whereSql, param
}

func (m modelUtil) GetSelectCoalesceSql(tables string) (coalesceSql string) {
	coalesceSql = m.GetSelectOut()
	coalesceSql += " FROM (" + tables + ") T "
	return
}

func (m modelUtil) GetSelectOut() string {
	selectOutSql := "SELECT "
	length := len(m.colSort)
	for index, col := range m.colSort {
		colStrs := strings.Split(col, ".")
		selectOutSql += "COALESCE(" + strings.ToUpper(colStrs[len(colStrs)-1]) + "," + strings.ToUpper(m.cols[col]) + ")"
		switch index {
		case length - 1:
			selectOutSql += " "
		default:
			selectOutSql += ","
		}
	}
	return selectOutSql
}

// 获取排序sql
func (m modelUtil) GetSortSql(paramMap map[string]interface{}) string {
	paramMap["sort"] = m.handleSortParam(paramMap["sort"])
	sortSql := getOrderBy("T.", paramMap)
	delete(paramMap, "sort")
	delete(paramMap, "order")
	return sortSql
}

func (m modelUtil) handleSortParam(sort interface{}) string {
	if nil == sort {
		return ""
	} else {
		var rst string
		sortString := sort.(string)
		sortStrs := strings.Split(sortString, ",")
		for index, sortStr := range sortStrs {
			strs := strings.Split(sortStr, ".")
			rst += strs[len(strs)-1]
			if index < len(sortStrs)-1 {
				rst += ","
			}
		}
		return rst
	}
}

// 获取完成的分页sql
func (m modelUtil) GetPageData(unionSql string, tables string, paramMap map[string]interface{}) (*apiUtil.PageData, *sql.Rows, error) {
	m.ParamMap = paramMap
	sortSql := m.GetSortSql(m.ParamMap)
	selectSqlOut := m.GetSelectOut() + " FROM ( "
	selectSqlMid := `SELECT T.*,ROWNUM AS NUM FROM (SELECT * FROM ` //分页查询sql
	selectSqlIn := "(" + unionSql + ") T "
	// selectSqlSort := `SELECT * FROM (`
	// whereSqlSort := `)`
	//获取分页数据
	start, length := getPageCount(m.ParamMap)
	whereSql, sqlParam := m.GetWhereSql(0)
	if "" != strings.TrimSpace(whereSql) {
		whereSql = " WHERE " + whereSql
	}
	m.sqlParamMap = append(m.sqlParamMap, sqlParam...)

	whereSqlMid := ") T WHERE ROWNUM <= " + strconv.Itoa(start+length)
	whereSqlOut := ") WHERE NUM > " + strconv.Itoa(start)
	// pageDataSql := selectSqlOut + selectSqlMid + selectSqlIn + whereSqlMid + whereSql + sortSql + whereSqlOut
	pageDataSql := selectSqlOut + selectSqlMid + selectSqlIn + whereSql + sortSql + whereSqlMid + whereSqlOut
	rows, err := getPageRow(pageDataSql, m.sqlParamMap)
	if nil != err {
		return nil, nil, err
	}
	pageData := getpageData("from "+selectSqlIn+whereSql, m.sqlParamMap)
	return pageData, rows, nil
}

type modelEngine struct {
	sql string
}

func NewModelEngine() *modelEngine {
	return new(modelEngine)
}

func (m *modelEngine) InitSql() *modelEngine {
	return m
}
