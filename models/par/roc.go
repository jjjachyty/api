package par

import (
	"database/sql"
	"fmt"
	"time"

	"pccqcpa.com.cn/app/rpm/api/models/dim"
	"pccqcpa.com.cn/app/rpm/api/models/sys"
	"pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/components/zlog"

	modelsUtil "pccqcpa.com.cn/app/rpm/api/models/util"
)

// 目前采用dbobj查询
// 结构体对应数据库表【RPM_PAR_ROC】资本回报率
/**
 * @apiDefine Roc
 * @apiSuccess {string}     UUID            主键默认值sys_guid()
 * @apiSuccess {Organ}   	Organ    		机构
 * @apiSuccess {Product}   	Product      	产品
 * @apiSuccess {string}   	CustSize   		客户规模
 * @apiSuccess {float64}   	CapitalCost    	资本成本率
 * @apiSuccess {float64}   	CapitalProfit   资本回报率
 * @apiSuccess {time.Time}  StartTime   	开始日期
 * @apiSuccess {string}   	Flag      		生效标志
 * @apiSuccess {time.Time}	CreateTime      创建时间
 * @apiSuccess {string}   	CreateUser      创建人
 * @apiSuccess {time.Time}	UpdateTime      更新时间
 * @apiSuccess {string}   	UpdateUser      更新人
 */
type Roc struct {
	UUID          string      `xorm:"VARCHAR2(32 CHAR) NOTNULL UNIQUE PK 'UUID'"`
	Organ         sys.Organ   `xorm:"extends"`
	Product       dim.Product `xorm:"extends"`
	CustSize      string      `xorm:"VARCHAR2(44 CHAR)  'CUST_SIZE'"`
	CapitalCost   float64     `xorm:"NUMBER(32,6)  'CAPITAL_COST'"`
	CapitalProfit float64     `xorm:"NUMBER(32,6)  'CAPITAL_PROFIT'"`
	StartTime     time.Time   `xorm:"DATE  'START_TIME'"`
	CreateTime    time.Time   `xorm:"DATE  'CREATE_TIME'"`
	UpdateTime    time.Time   `xorm:"DATE  'UPDATE_TIME'"`
	CreateUser    string      `xorm:"VARCHAR2(44 CHAR)  'CREATE_USER'"`
	UpdateUser    string      `xorm:"VARCHAR2(44 CHAR)  'UPDATE_USER'"`
	Flag          string      `xorm:"VARCHAR2(2 CHAR)  'FLAG'"`
}

func (roc *Roc) scanRoc(rows *sql.Rows) (*Roc, error) {
	var one Roc
	values := []interface{}{
		&one.UUID,
		&one.CustSize,
		&one.CapitalCost,
		&one.CapitalProfit,
		&one.StartTime,
		&one.CreateTime,
		&one.UpdateTime,
		&one.CreateUser,
		&one.UpdateUser,
		&one.Flag,

		&one.Organ.OrganCode,
		&one.Organ.OrganName,

		&one.Product.ProductCode,
		&one.Product.ProductName,
	}
	err := util.OracleScan(rows, values)
	return &one, err
}

func (roc *Roc) SelectRocByParams(paramMap map[string]interface{}) ([]*Roc, error) {
	rows, err := modelsUtil.FindRows(rocTabls, rocCols, rocColsSort, paramMap)
	if nil != err {
		er := fmt.Errorf("查询资本回报率出错")
		zlog.Error(er.Error(), err)
		return nil, er
	}
	defer rows.Close()
	var rocs []*Roc
	for rows.Next() {
		one, err := roc.scanRoc(rows)
		if nil != err {
			er := fmt.Errorf("多参数查询资本回报率rows.Scan()错误")
			zlog.Error(er.Error(), err)
			return nil, er
		}
		rocs = append(rocs, one)
	}
	return rocs, nil
}

var rocTabls string = `
	     RPM_PAR_ROC T
	LEFT JOIN SYS_SEC_ORGAN ORGAN
	ON (T.ORGAN  = ORGAN.ORGAN_CODE)
	LEFT JOIN RPM_DIM_PRODUCT PRODUCT
	ON (T.PRODUCT = PRODUCT.PRODUCT_CODE)
`
var rocCols = map[string]string{
	"T.UUID":           "' '",
	"T.CUST_SIZE":      "' '",
	"T.CAPITAL_COST":   "0",
	"T.CAPITAL_PROFIT": "0",
	"T.START_TIME":     "sysdate",
	"T.CREATE_TIME":    "sysdate",
	"T.UPDATE_TIME":    "sysdate",
	"T.CREATE_USER":    "' '",
	"T.UPDATE_USER":    "' '",
	"T.FLAG":           "' '",

	"T.ORGAN":          "' '",
	"ORGAN.ORGAN_NAME": "' '",

	"T.PRODUCT":            "' '",
	"PRODUCT.PRODUCT_NAME": "' '",
}

var rocColsSort = []string{
	"T.UUID",
	"T.CUST_SIZE",
	"T.CAPITAL_COST",
	"T.CAPITAL_PROFIT",
	"T.START_TIME",
	"T.CREATE_TIME",
	"T.UPDATE_TIME",
	"T.CREATE_USER",
	"T.UPDATE_USER",
	"T.FLAG",

	"T.ORGAN",
	"ORGAN.ORGAN_NAME",

	"T.PRODUCT",
	"PRODUCT.PRODUCT_NAME",
}

// 分页操作
// by author Yeqc
// by time 2016-12-21 13:39:33
func (this Roc) List(param ...map[string]interface{}) (*util.PageData, error) {
	pageData, rows, err := modelsUtil.List(rocTabls, rocCols, rocColsSort, param...)
	defer rows.Close()
	if nil != err {
		er := fmt.Errorf("分页查询资本回报率信息出错")
		zlog.Error(er.Error(), err)
		return nil, er
	}

	var rst []*Roc
	for rows.Next() {
		one, err := this.scanRoc(rows)
		if nil != err {
			er := fmt.Errorf("分页查询资本回报率信息row.Scan()出错")
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
// by time 2016-12-21 13:39:33
func (this Roc) Find(param ...map[string]interface{}) ([]*Roc, error) {
	rows, err := modelsUtil.FindRows(rocTabls, rocCols, rocColsSort, param...)
	defer rows.Close()
	if nil != err {
		er := fmt.Errorf("多参数查询资本回报率信息出错")
		zlog.Error(er.Error(), err)
		return nil, er
	}
	var rst []*Roc
	for rows.Next() {
		one, err := this.scanRoc(rows)
		if nil != err {
			er := fmt.Errorf("多参数查询资本回报率信息row.Scan()出错")
			zlog.Error(er.Error(), err)
			return nil, er
		}
		rst = append(rst, one)
	}
	return rst, nil
}

// 新增资本回报率信息
// by author Yeqc
// by time 2016-12-21 13:39:33
func (this Roc) Add() error {
	paramMap := map[string]interface{}{
		"organ":          this.Organ.OrganCode,
		"product":        this.Product.ProductCode,
		"cust_size":      this.CustSize,
		"capital_cost":   this.CapitalCost,
		"capital_profit": this.CapitalProfit,
		"start_time":     this.StartTime,
		"create_time":    util.GetCurrentTime(),
		"update_time":    util.GetCurrentTime(),
		"create_user":    this.CreateUser,
		"update_user":    this.UpdateUser,
		"flag":           this.Flag,
	}
	err := util.OracleAdd("RPM_PAR_ROC", paramMap)
	if nil != err {
		er := fmt.Errorf("新增资本回报率信息出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

// 更新资本回报率信息
// by author Yeqc
// by time 2016-12-21 13:39:33
func (this Roc) Update() error {
	paramMap := map[string]interface{}{
		"organ":          this.Organ.OrganCode,
		"product":        this.Product.ProductCode,
		"cust_size":      this.CustSize,
		"capital_cost":   this.CapitalCost,
		"capital_profit": this.CapitalProfit,
		"start_time":     this.StartTime,
		"update_time":    util.GetCurrentTime(),
		"update_user":    this.UpdateUser,
		"flag":           this.Flag,
	}
	whereParam := map[string]interface{}{
		"uuid": this.UUID,
	}
	err := util.OracleUpdate("RPM_PAR_ROC", paramMap, whereParam)
	if nil != err {
		er := fmt.Errorf("更新资本回报率信息出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

// 删除资本回报率信息
// by author Yeqc
// by time 2016-12-21 13:39:33
func (this Roc) Delete() error {
	whereParam := map[string]interface{}{
		"uuid": this.UUID,
	}
	err := util.OracleDelete("RPM_PAR_ROC", whereParam)
	if nil != err {
		er := fmt.Errorf("删除资本回报率信息出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}
