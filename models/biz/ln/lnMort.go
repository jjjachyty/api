package ln

import (
	"database/sql"
	"fmt"
	"time"

	"pccqcpa.com.cn/app/rpm/api/models/dim"
	"pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/components/zlog"

	modelsUtil "pccqcpa.com.cn/app/rpm/api/models/util"
)

type LnMort struct {
	UUID          string
	BusinessCode  string
	MortgageCode  string
	MortgageName  string
	Mortgage      dim.Mortgage `json:-`
	Currency      string
	MortgageValue float64
	CreateTime    time.Time
	UpdateTime    time.Time
	CreateUser    string
	UpdateUser    string
}

func (l *LnMort) scan(rows *sql.Rows) (*LnMort, error) {
	var lnMort = new(LnMort)
	values := []interface{}{
		&lnMort.UUID,
		&lnMort.BusinessCode,
		&lnMort.MortgageCode,
		&lnMort.MortgageName,
		&lnMort.Currency,
		&lnMort.MortgageValue,
		&lnMort.CreateTime,
		&lnMort.UpdateTime,
		&lnMort.CreateUser,
		&lnMort.UpdateUser,
		&lnMort.Mortgage.MortgageName,
		&lnMort.Mortgage.MortgageBasel,
	}
	err := util.OracleScan(rows, values)
	return lnMort, err
}

func (l *LnMort) List(param ...map[string]interface{}) (*util.PageData, error) {
	pageData, rows, err := modelsUtil.List(lnMortTables, lnMortCols, lnMortColsSort, param...)
	defer rows.Close()
	if nil != err {
		er := fmt.Errorf("分页查询抵押品出错")
		zlog.Error(er.Error(), err)
		return nil, er
	}
	var lnMorts []*LnMort
	for rows.Next() {
		lnMort, err := l.scan(rows)
		if nil != err {
			er := fmt.Errorf("查询抵押品信息rows.Scan()出错")
			zlog.Error(er.Error(), err)
			return nil, er
		}
		lnMorts = append(lnMorts, lnMort)
	}
	pageData.Rows = lnMorts
	return pageData, nil
}

func (l *LnMort) Find(param ...map[string]interface{}) ([]*LnMort, error) {
	rows, err := modelsUtil.FindRows(lnMortTables, lnMortCols, lnMortColsSort, param...)

	defer rows.Close()
	if nil != err {
		er := fmt.Errorf("查询抵质押品信息出错")
		zlog.Error(er.Error(), err)
		return nil, er
	}

	var lnMorts []*LnMort
	for rows.Next() {
		lnMort, err := l.scan(rows)
		if nil != err {
			er := fmt.Errorf("查询抵押品信息rows.Scan()出错")
			zlog.Error(er.Error(), err)
			return nil, er
		}
		lnMorts = append(lnMorts, lnMort)
	}

	return lnMorts, nil
}

func (l *LnMort) Add() error {
	param := map[string]interface{}{
		"business_code":  l.BusinessCode,
		"mortgage_code":  l.MortgageCode,
		"mortgage_name":  l.MortgageName,
		"currency":       l.Currency,
		"mortgage_value": l.MortgageValue,
		"create_time":    util.GetCurrentTime(),
		"update_time":    util.GetCurrentTime(),
		"create_user":    l.CreateUser,
		"update_user":    l.UpdateUser,
	}
	err := util.OracleAdd(lnMortTableName, param)
	if nil != err {
		er := fmt.Errorf("新增抵质押品失败")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

func (l *LnMort) Update() error {
	param := map[string]interface{}{
		"mortgage_code":  l.MortgageCode,
		"mortgage_name":  l.MortgageName,
		"currency":       l.Currency,
		"mortgage_value": l.MortgageValue,
		"update_time":    util.GetCurrentTime(),
		"update_user":    l.UpdateUser,
	}
	whereParam := map[string]interface{}{
		"uuid": l.UUID,
	}
	err := util.OracleUpdate(lnMortTableName, param, whereParam)
	if nil != err {
		er := fmt.Errorf("更新抵质押品失败")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

func (l *LnMort) Delete() error {
	whereParam := map[string]interface{}{
		"uuid": l.UUID,
	}
	err := util.OracleDelete(lnMortTableName, whereParam)
	if nil != err {
		er := fmt.Errorf("删除抵质押品失败")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

func (l *LnMort) DeleteByBusinessCode(businessCode string) error {
	whereParam := map[string]interface{}{
		"business_code": businessCode,
	}
	err := util.OracleDelete(lnMortTableName, whereParam)
	if nil != err {
		er := fmt.Errorf("删除业务编号[%v]抵质押品失败", businessCode)
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

var lnMortTableName string = "rpm_biz_mortgage"

var lnMortTables string = `
	     RPM_BIZ_MORTGAGE T
	LEFT JOIN RPM_DIM_MORTGAGE MORTGAGE
	  ON (T.MORTGAGE_CODE = MORTGAGE.MORTGAGE_CODE)
`

var lnMortCols = map[string]string{
	"T.UUID":           "' '",
	"T.BUSINESS_CODE":  "' '",
	"T.MORTGAGE_CODE":  "' '",
	"T.MORTGAGE_NAME":  "' '",
	"T.CURRENCY":       "' '",
	"T.MORTGAGE_VALUE": "0",
	"T.CREATE_TIME":    "sysdate",
	"T.UPDATE_TIME":    "sysdate",
	"T.CREATE_USER":    "' '",
	"T.UPDATE_USER":    "' '",

	"MORTGAGE.MORTGAGE_NAME":  "' '",
	"MORTGAGE.MORTGAGE_BASEL": "' '",
}

var lnMortColsSort = []string{
	"T.UUID",
	"T.BUSINESS_CODE",
	"T.MORTGAGE_CODE",
	"T.MORTGAGE_NAME",
	"T.CURRENCY",
	"T.MORTGAGE_VALUE",
	"T.CREATE_TIME",
	"T.UPDATE_TIME",
	"T.CREATE_USER",
	"T.UPDATE_USER",

	"MORTGAGE.MORTGAGE_NAME",
	"MORTGAGE.MORTGAGE_BASEL",
}
