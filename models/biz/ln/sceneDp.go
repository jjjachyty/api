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

type SceneDp struct {
	UUID         string
	BusinessCode string      `xorm:"business"`
	Business     LnBusiness  `json:"-"`
	Product      dim.Product `xorm:<- `
	Currency     string
	Term         int
	Rate         float64
	Value        float64
	Eva          float64
	Discount     float64
	CreateTime   time.Time
	UpdateTime   time.Time
	CreateUser   string
	UpdateUser   string
}

// scan func SceneDp 赋值
func (sdp *SceneDp) scan(rows *sql.Rows) (*SceneDp, error) {
	var sceneDp = new(SceneDp)
	values := []interface{}{
		&sceneDp.UUID,
		&sceneDp.BusinessCode,
		&sceneDp.Currency,
		&sceneDp.Term,
		&sceneDp.Rate,
		&sceneDp.Value,
		&sceneDp.Eva,
		&sceneDp.Discount,
		&sceneDp.CreateTime,
		&sceneDp.UpdateTime,
		&sceneDp.CreateUser,
		&sceneDp.UpdateUser,
		&sceneDp.Product.ProductCode,
		&sceneDp.Product.ProductName,
	}
	err := util.OracleScan(rows, values)
	return sceneDp, err
}

// Find func SceneDp 多参数查询存款派生信息
func (sdp *SceneDp) Find(param ...map[string]interface{}) ([]*SceneDp, error) {
	rows, err := modelsUtil.FindRows(sceneDpTables, sceneDpTCols, sceneDpTColsSort, param...)
	if nil != err {
		er := fmt.Errorf("查询存款派生信息失败")
		zlog.Error(er.Error(), err)
		return nil, er
	}
	defer rows.Close()
	var sceneDps []*SceneDp
	for rows.Next() {
		sceneDp, err := sdp.scan(rows)
		if nil != err {
			er := fmt.Errorf("查询存款派生信息rows.Scan()出错")
			zlog.Error(er.Error(), err)
			return nil, er
		}
		sceneDps = append(sceneDps, sceneDp)
	}
	return sceneDps, nil
}

func (sdp *SceneDp) Add() error {
	param := map[string]interface{}{
		"business_code": sdp.BusinessCode,
		"product":       sdp.Product.ProductCode,
		"currency":      sdp.Currency,
		"term":          sdp.Term,
		"rate":          sdp.Rate,
		"value":         sdp.Value,
		"eva":           sdp.Eva,
		"discount":      sdp.Discount,
		"create_time":   util.GetCurrentTime(),
		"update_time":   util.GetCurrentTime(),
		"create_user":   sdp.CreateUser,
		"update_user":   sdp.UpdateUser,
	}

	err := util.OracleAdd(sceneDpTable, param)
	if nil != err {
		er := fmt.Errorf("新增存款派生记录失败")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

func (sdp *SceneDp) Update() error {
	param := map[string]interface{}{
		"business_code": sdp.BusinessCode,
		"product":       sdp.Product.ProductCode,
		"currency":      sdp.Currency,
		"term":          sdp.Term,
		"rate":          sdp.Rate,
		"value":         sdp.Value,
		"eva":           sdp.Eva,
		"discount":      sdp.Discount,
		"create_time":   util.GetCurrentTime(),
		"update_time":   util.GetCurrentTime(),
		"update_user":   sdp.UpdateUser,
	}
	whereParam := map[string]interface{}{
		"uuid": sdp.UUID,
	}

	err := util.OracleUpdate(sceneDpTable, param, whereParam)
	if nil != err {
		er := fmt.Errorf("更新存款派生记录失败")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

func (sdp *SceneDp) Delete() error {
	whereParam := map[string]interface{}{
		"uuid": sdp.UUID,
	}

	err := util.OracleDelete(sceneDpTable, whereParam)
	if nil != err {
		er := fmt.Errorf("删除存款派生记录失败")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

func (sdp *SceneDp) DeleteByBusienssCode(businessCode string) error {
	whereParam := map[string]interface{}{
		"business_code": businessCode,
	}

	err := util.OracleDelete(sceneDpTable, whereParam)
	if nil != err {
		er := fmt.Errorf("业务编号删除存款派生记录失败")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

var sceneDpTable string = "rpm_biz_scene_dp"

var sceneDpTables string = `
			     RPM_BIZ_SCENE_DP T
			LEFT JOIN RPM_DIM_PRODUCT PRODUCT
			  ON (T.PRODUCT = PRODUCT.PRODUCT_CODE)
`

var sceneDpTCols = map[string]string{
	"T.UUID":               "' '",
	"T.BUSINESS_CODE":      "' '",
	"T.CURRENCY":           "' '",
	"T.TERM":               "0",
	"T.RATE":               "0",
	"T.VALUE":              "0",
	"T.EVA":                "0",
	"T.DISCOUNT":           "0",
	"T.CREATE_TIME":        "sysdate",
	"T.UPDATE_TIME":        "sysdate",
	"T.CREATE_USER":        "' '",
	"T.UPDATE_USER":        "' '",
	"T.PRODUCT":            "' '",
	"PRODUCT.PRODUCT_NAME": "' '",
}

var sceneDpTColsSort = []string{
	"T.UUID",
	"T.BUSINESS_CODE",
	"T.CURRENCY",
	"T.TERM",
	"T.RATE",
	"T.VALUE",
	"T.EVA",
	"T.DISCOUNT",
	"T.CREATE_TIME",
	"T.UPDATE_TIME",
	"T.CREATE_USER",
	"T.UPDATE_USER",
	"T.PRODUCT",
	"PRODUCT.PRODUCT_NAME",
}
