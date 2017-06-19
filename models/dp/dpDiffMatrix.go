package dp

import (
	"database/sql"
	"fmt"
	"time"

	"pccqcpa.com.cn/app/rpm/api/models/dim"
	"pccqcpa.com.cn/app/rpm/api/models/sys"
	modelsUtil "pccqcpa.com.cn/app/rpm/api/models/util"
	"pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/components/zlog"
)

// 结构体对应数据库表【RPM_BIZ_DP_DIFF_MATRIX】存款差异化矩阵表
// by author Jason
// by time 2017-02-07 16:48:18
type DpDiffMatrix struct {
	UUID           string      // 唯一键
	Product        dim.Product // 产品
	Term           float64     // 期限
	Name           string      // 名称
	Branch         sys.Organ   // 机构
	CustGrade      string      // 客户级别
	AmountGrade    string      // 金额区间
	Channel        string      // 渠道
	Rate           float64     // 利率
	RateUpperLimit float64     // 利率上限
	CreateTime     time.Time   // 创建时间
}

// 结构体scan
// by author Jason
// by time 2017-02-07 16:48:18
func (this DpDiffMatrix) scan(rows *sql.Rows) (*DpDiffMatrix, error) {
	var one = new(DpDiffMatrix)
	values := []interface{}{
		&one.UUID,
		&one.Term,
		&one.Name,
		&one.CustGrade,
		&one.AmountGrade,
		&one.Channel,
		&one.Rate,
		&one.RateUpperLimit,
		&one.CreateTime,

		&one.Product.ProductCode,
		&one.Product.ProductName,

		&one.Branch.OrganCode,
		&one.Branch.OrganName,
	}
	err := util.OracleScan(rows, values)
	return one, err
}

// 分页操作
// by author Jason
// by time 2017-02-07 16:48:18
func (this DpDiffMatrix) List(param ...map[string]interface{}) (*util.PageData, error) {
	pageData, rows, err := modelsUtil.List(dpDiffMatrixTabales, dpDiffMatrixCols, dpDiffMatrixColsSort, param...)
	defer rows.Close()
	if nil != err {
		er := fmt.Errorf("分页查询存款差异化矩阵表信息出错")
		zlog.Error(er.Error(), err)
		return nil, er
	}

	var rst []*DpDiffMatrix
	for rows.Next() {
		one, err := this.scan(rows)
		if nil != err {
			er := fmt.Errorf("分页查询存款差异化矩阵表信息row.Scan()出错")
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
// by time 2017-02-07 16:48:18
func (this DpDiffMatrix) Find(param ...map[string]interface{}) ([]*DpDiffMatrix, error) {
	rows, err := modelsUtil.FindRows(dpDiffMatrixTabales, dpDiffMatrixCols, dpDiffMatrixColsSort, param...)
	defer rows.Close()
	if nil != err {
		er := fmt.Errorf("多参数查询存款差异化矩阵表信息出错")
		zlog.Error(er.Error(), err)
		return nil, er
	}
	var rst []*DpDiffMatrix
	for rows.Next() {
		one, err := this.scan(rows)
		if nil != err {
			er := fmt.Errorf("多参数查询存款差异化矩阵表信息row.Scan()出错")
			zlog.Error(er.Error(), err)
			return nil, er
		}
		rst = append(rst, one)
	}
	return rst, nil
}

// 新增存款差异化矩阵表信息
// by author Jason
// by time 2017-02-07 16:48:18
func (this DpDiffMatrix) Add() error {
	paramMap := map[string]interface{}{
		"product":      this.Product,
		"term":         this.Term,
		"name":         this.Name,
		"branch":       this.Branch,
		"cust_grade":   this.CustGrade,
		"amount_grade": this.AmountGrade,
		"channel":      this.Channel,
		"rate":         this.Rate,
		"create_time":  util.GetCurrentTime(),
	}
	err := util.OracleAdd("RPM_BIZ_DP_DIFF_MATRIX", paramMap)
	if nil != err {
		er := fmt.Errorf("新增存款差异化矩阵表信息出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

// 更新存款差异化矩阵表信息
// by author Jason
// by time 2017-02-07 16:48:18
func (this DpDiffMatrix) Update() error {
	paramMap := map[string]interface{}{
		"product":      this.Product,
		"term":         this.Term,
		"name":         this.Name,
		"branch":       this.Branch,
		"cust_grade":   this.CustGrade,
		"amount_grade": this.AmountGrade,
		"channel":      this.Channel,
		"rate":         this.Rate,
	}
	whereParam := map[string]interface{}{
		"uuid": this.UUID,
	}
	err := util.OracleUpdate("RPM_BIZ_DP_DIFF_MATRIX", paramMap, whereParam)
	if nil != err {
		er := fmt.Errorf("更新存款差异化矩阵表信息出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

// 删除存款差异化矩阵表信息
// by author Jason
// by time 2017-02-07 16:48:18
func (this DpDiffMatrix) Delete() error {
	whereParam := map[string]interface{}{
		"uuid": this.UUID,
	}
	err := util.OracleDelete("RPM_BIZ_DP_DIFF_MATRIX", whereParam)
	if nil != err {
		er := fmt.Errorf("删除存款差异化矩阵表信息出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

var dpDiffMatrixTabales string = `
		 RPM_BIZ_DP_DIFF_MATRIX T
	LEFT JOIN RPM_DIM_PRODUCT PRODUCT
	  ON (T.PRODUCT = PRODUCT.PRODUCT_CODE)
	LEFT JOIN SYS_SEC_ORGAN BRANCH
	  ON (T.BRANCH = BRANCH.ORGAN_CODE)`

var dpDiffMatrixCols map[string]string = map[string]string{
	"T.UUID":               "' '",
	"T.TERM":               "0",
	"T.NAME":               "' '",
	"T.CUST_GRADE":         "' '",
	"T.AMOUNT_GRADE":       "' '",
	"T.CHANNEL":            "' '",
	"T.RATE":               "0",
	"T.CREATE_TIME":        "sysdate",
	"T.PRODUCT":            "' '",
	"PRODUCT.PRODUCT_NAME": "' '",
	"T.BRANCH":             "' '",
	"BRANCH.ORGAN_NAME":    "' '",
	"T.RATE_UPPER_LIMIT":   "0",
}

var dpDiffMatrixColsSort = []string{
	"T.UUID",
	"T.TERM",
	"T.NAME",
	"T.CUST_GRADE",
	"T.AMOUNT_GRADE",
	"T.CHANNEL",
	"T.RATE",
	"T.RATE_UPPER_LIMIT",
	"T.CREATE_TIME",
	"T.PRODUCT",
	"PRODUCT.PRODUCT_NAME",
	"T.BRANCH",
	"BRANCH.ORGAN_NAME",
}
