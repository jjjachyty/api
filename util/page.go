package util

//分页查询公共类
//@auth By Janly
import (
	"database/sql"
	"errors"
	"fmt"
	"math"
	"net/http"
	"platform/dbobj"

	"strconv"

	"pccqcpa.com.cn/components/zlog"
)

// Page type 分页信息
//@auth By Janly11
type Page struct {
	PreviousPage   int64 //上一页
	CurrentPage    int64 //当前页
	NextPage       int64 //下一页
	PageSize       int64 //分页大小
	TotalRows      int64 //总记录数
	TotalPage      int64 //总页数
	StartRowNumber int64 //开始行数
	EndRowNumber   int64 //结束行数
}

func GetPage(stratRowNumber int64, pageSize int64, querStr string, params ...[]interface{}) Page {
	var page Page
	var count int64
	var rows *sql.Rows
	var err error
	querStr = "SELECT COUNT(1) FROM (" + querStr + ")"
	if nil != params && 0 < len(params[0]) {
		rows, err = dbobj.Default.Query(querStr, params[0]...)
		zlog.Debugf("分页查询SQL：%s\nparam[%v]", err, querStr, params[0])
	} else {
		rows, err = dbobj.Default.Query(querStr)
		zlog.Debugf("分页查询SQL：%s", err, querStr)
	}
	if nil != err {
		zlog.Error("分页查询失败：", err)
	}
	defer rows.Close()
	fmt.Println(rows.Next())
	rows.Scan(&count)
	zlog.Debugf("查询数据总条数为:%d", nil, count)
	if nil == err {
		page.TotalRows = count
	} else {
		page.TotalRows = 0
	}
	page.setPageAndSize(pageSize)
	page.setStartAndEndRows(stratRowNumber)
	page.setPreviousPage()
	page.setNextPage()
	page.setTotalPage()

	// repHeader.Set("Total-Page", strconv.FormatInt(page.TotalPage, 10))
	// repHeader.Set("Total-Rows", strconv.FormatInt(page.TotalRows, 10))
	return page
}

// func GetPage(reqHeader http.Header, bean interface{}, tableName string, queryStr string, queryArgs ...interface{}) Page {

// 	var page Page
// 	var session = Engine.NewSession()
// 	if "" != tableName {
// 		session.Table(tableName)
// 	}
// 	if queryStr != "" {
// 		session.Where(queryStr, queryArgs...)
// 	}
// 	count, err := session.Count(bean)
// 	zlog.Debugf("查询数据总条数为:%d", nil, count)
// 	if nil == err {
// 		page.TotalRows = count
// 	} else {
// 		page.TotalRows = 0
// 	}

// 	page.setPageAndSize(reqHeader)
// 	page.setPreviousPage()
// 	page.setNextPage()
// 	page.setTotalPage()
// 	page.setStartAndEndRows(reqHeader)

// 	// repHeader.Set("Total-Page", strconv.FormatInt(page.TotalPage, 10))
// 	// repHeader.Set("Total-Rows", strconv.FormatInt(page.TotalRows, 10))
// 	return page

// }

func (page *Page) setPageAndSize(pageSize int64) {

	if pageSize == 0 {
		zlog.Warning("分页查询[分页大小]对应的值为0,默认查询全部数据", nil)
		page.PageSize = page.TotalRows
		page.CurrentPage = 1
	} else {
		page.PageSize = pageSize
	}
}

func (page *Page) setPreviousPage() {
	if page.CurrentPage > 1 {
		page.PreviousPage = page.CurrentPage - 1
	} else {
		page.PreviousPage = page.CurrentPage
	}
}

func (page *Page) setNextPage() {
	if page.CurrentPage < page.TotalPage {
		page.NextPage = page.CurrentPage + 1
	} else {
		page.NextPage = page.TotalPage
	}
}

//设置总页数
//@auth by Janly
func (page *Page) setTotalPage() {
	if 0 == page.PageSize {
		page.TotalPage = 1
	} else {
		pageSizeVal := page.TotalRows / page.PageSize
		pageLeft := page.TotalRows % page.PageSize

		if pageLeft > 0 { //有余数
			page.TotalPage = int64(math.Trunc(float64(pageSizeVal))) + 1
		} else {
			page.TotalPage = int64(math.Trunc(float64(pageSizeVal)))
		}

	}
}

func (page *Page) setStartAndEndRows(startRowNumber int64) {
	var endRowNumber int64
	fmt.Println("开始行数", startRowNumber)
	if startRowNumber < 0 {
		startRowNumber = 0
	} else if startRowNumber > page.TotalRows {
		startRowNumber = page.TotalRows
	}
	page.StartRowNumber = startRowNumber
	if startRowNumber == 0 {
		endRowNumber = page.PageSize
	} else {
		endRowNumber = startRowNumber + page.PageSize
	}

	if endRowNumber > page.TotalRows {
		page.EndRowNumber = page.TotalRows
	} else {
		page.EndRowNumber = endRowNumber
	}
	if 0 == page.PageSize {
		page.CurrentPage = 1
	} else {
		currentPageMod := page.StartRowNumber % page.PageSize
		currentPage := page.StartRowNumber / page.PageSize
		if currentPageMod > 0 { //有余数
			page.CurrentPage = int64(math.Trunc(float64(currentPage))) + 1
		} else {
			page.CurrentPage = currentPage
		}
	}
}

//GetPageAndOrder func
//用于Action 分页查询返回startRowNumber，pageSize，orderAttr,orderType，error
func GetPageAndOrder(header http.Header) (int64, int64, string, OrderType, error) {
	var startRowNumberInt int64
	var pageSizeInt int64
	var orderType OrderType
	var err error
	startRowNumberStr := header.Get("start-row-number")
	pageSizeStr := header.Get("page-size")
	orderAttr := header.Get("Order-Attr")
	orderTypeStr := header.Get("Order-Type")

	switch orderTypeStr {
	case "ASC":
		orderType = ASC
	default:
		orderType = DESC
	}

	if startRowNumberStr != "" {
		startRowNumberInt, err = strconv.ParseInt(startRowNumberStr, 10, 32)
		if err != nil {
			return 0, 0, orderAttr, orderType, errors.New(" Start-Row-Number 不是正整数")
		}

	} else {
		return 0, 0, orderAttr, orderType, errors.New("无法从头中获取 Start-Row-Number ")
	}

	if pageSizeStr != "" {
		pageSizeInt, err = strconv.ParseInt(pageSizeStr, 10, 32)
		if err != nil {
			return startRowNumberInt, 0, orderAttr, orderType, errors.New(" Page-Size 不是正整数")
		}

	} else {
		return startRowNumberInt, 0, orderAttr, orderType, errors.New("无法从头中获取 Page-Size ")
	}

	// if orderTypeStr == "DESC" {
	// 	orderType = DESC
	// } else if orderTypeStr == "ASC" {
	// 	orderType = ASC
	// }

	return startRowNumberInt, pageSizeInt, orderAttr, orderType, nil
}
