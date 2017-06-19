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

type SceneItd struct {
	UUID         string
	BusinessCode string
	Business     LnBusiness `json:"-"`
	Product      dim.Product
	Value        float64
	Eva          float64
	Diacount     float64
	CreateTime   time.Time
	UpdateTime   time.Time
	CreateUser   string
	UpdateUser   string
}

var sceneItdTableName = "rpm_biz_scene_itd"

// Add func SceneItd 新增中间派生
func (s *SceneItd) scan(rows *sql.Rows) (*SceneItd, error) {
	var sceneItd = new(SceneItd)
	values := []interface{}{
		&sceneItd.UUID,
		&sceneItd.BusinessCode,
		&sceneItd.Value,
		&sceneItd.Eva,
		&sceneItd.Diacount,
		&sceneItd.CreateTime,
		&sceneItd.UpdateTime,
		&sceneItd.CreateUser,
		&sceneItd.UpdateUser,

		&sceneItd.Product.ProductCode,
		&sceneItd.Product.ProductName,
	}
	err := util.OracleScan(rows, values)
	return sceneItd, err
}

func (s *SceneItd) Find(param ...map[string]interface{}) ([]*SceneItd, error) {
	rows, err := modelsUtil.FindRows(sceneItdTables, sceneItdCols, sceneItdColsSort, param...)
	defer rows.Close()
	if nil != err {
		er := fmt.Errorf("查询派生中间业务出错")
		zlog.Error(er.Error(), err)
		return nil, er
	}
	var sceneItds []*SceneItd
	for rows.Next() {
		sceneItd, err := s.scan(rows)
		if nil != err {
			er := fmt.Errorf("查询派生中间业务rows.Scan()出错")
			zlog.Error(er.Error(), err)
			return nil, er
		}
		sceneItds = append(sceneItds, sceneItd)
	}
	return sceneItds, nil
}

// Add func SceneItd 新增中间派生
func (s *SceneItd) Add() error {
	var param = map[string]interface{}{
		"business_code": s.BusinessCode,
		"product":       s.Product.ProductCode,
		"value":         s.Value,
		"eva":           s.Eva,
		"diacount":      s.Diacount,
		"create_time":   util.GetCurrentTime(),
		"update_time":   util.GetCurrentTime(),
		"create_user":   s.CreateUser,
		"update_user":   s.UpdateUser,
	}
	err := util.OracleAdd(sceneItdTableName, param)
	if nil != err {
		er := fmt.Errorf("新增中间业务记录出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

// Update func SceneItd 更新中间派生
func (s *SceneItd) Update() error {
	var param = map[string]interface{}{
		"business_code": s.BusinessCode,
		"product":       s.Product.ProductCode,
		"value":         s.Value,
		"eva":           s.Eva,
		"diacount":      s.Diacount,
		"update_time":   util.GetCurrentTime(),
		"update_user":   s.UpdateUser,
	}
	var whereParam = map[string]interface{}{
		"uuid": s.UUID,
	}
	err := util.OracleUpdate(sceneItdTableName, param, whereParam)
	if nil != err {
		er := fmt.Errorf("更新中间业务记录出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

// Delete func SceneItd 删除中间派生
func (s *SceneItd) Delete() error {
	var whereParam = map[string]interface{}{
		"uuid": s.UUID,
	}
	err := util.OracleDelete(sceneItdTableName, whereParam)
	if nil != err {
		er := fmt.Errorf("删除中间业务记录出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

// Delete func SceneItd 删除中间派生
func (s *SceneItd) DeleteByBusinessCode(businessCode string) error {
	var whereParam = map[string]interface{}{
		"business_code": businessCode,
	}
	err := util.OracleDelete(sceneItdTableName, whereParam)
	if nil != err {
		er := fmt.Errorf("业务编号删除中间业务记录出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

var sceneItdTables string = `
	     RPM_BIZ_SCENE_ITD T
	LEFT JOIN RPM_DIM_PRODUCT PRODUCT
	  ON (T.PRODUCT = PRODUCT.PRODUCT_CODE)
`

var sceneItdCols = map[string]string{
	"T.uuid":               "' '",
	"T.business_code":      "' '",
	"T.value":              "0",
	"T.eva":                "0",
	"T.diacount":           "0",
	"T.create_time":        "sysdate",
	"T.update_time":        "sysdate",
	"T.create_user":        "' '",
	"T.update_user":        "' '",
	"T.product":            "' '",
	"PRODUCT.product_name": "' '",
}

var sceneItdColsSort = []string{
	"T.uuid",
	"T.business_code",
	"T.value",
	"T.eva",
	"T.diacount",
	"T.create_time",
	"T.update_time",
	"T.create_user",
	"T.update_user",
	"T.product",
	"PRODUCT.product_name",
}
