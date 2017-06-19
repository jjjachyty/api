package par

import (
	"database/sql"
	"fmt"
	"time"

	"pccqcpa.com.cn/app/rpm/api/models/dim"
	modelsUtil "pccqcpa.com.cn/app/rpm/api/models/util"
	"pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/components/zlog"
)

// 结构体对应数据库表【RPM_PAR_SCENE_DISCOUNT】派生优惠参数表
// by author Jason
// by time 2017-02-17 10:25:04
type SceneDiscount struct {
	UUID       string      // UUID
	BizType    string      // 业务类型
	CustImplvl string      // 客户级别
	CustSize   string      // 客户规模
	Product    dim.Product // 产品号
	Rate       float64     // 折让率
	CreateUser string      // 创建时间
	UpdateUser string      // 更新时间
	CreateTime time.Time   // 创建人
	UpdateTime time.Time   // 更新人
}

// 结构体scan
// by author Jason
// by time 2017-02-17 10:25:04
func (this SceneDiscount) scan(rows *sql.Rows) (*SceneDiscount, error) {
	var one = new(SceneDiscount)
	values := []interface{}{
		&one.UUID,
		&one.BizType,
		&one.CustImplvl,
		&one.CustSize,
		&one.Rate,
		&one.CreateUser,
		&one.UpdateUser,
		&one.CreateTime,
		&one.UpdateTime,

		&one.Product.ProductCode,
		&one.Product.ProductName,
	}
	err := util.OracleScan(rows, values)
	return one, err
}

// 分页操作
// by author Jason
// by time 2017-02-17 10:25:04
func (this SceneDiscount) List(param ...map[string]interface{}) (*util.PageData, error) {
	pageData, rows, err := modelsUtil.List(sceneDiscountTabales, sceneDiscountCols, sceneDiscountColsSort, param...)
	defer rows.Close()
	if nil != err {
		er := fmt.Errorf("分页查询派生优惠参数表信息出错")
		zlog.Error(er.Error(), err)
		return nil, er
	}

	var rst []*SceneDiscount
	for rows.Next() {
		one, err := this.scan(rows)
		if nil != err {
			er := fmt.Errorf("分页查询派生优惠参数表信息row.Scan()出错")
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
// by time 2017-02-17 10:25:04
func (this SceneDiscount) Find(param ...map[string]interface{}) ([]*SceneDiscount, error) {
	rows, err := modelsUtil.FindRows(sceneDiscountTabales, sceneDiscountCols, sceneDiscountColsSort, param...)
	defer rows.Close()
	if nil != err {
		er := fmt.Errorf("多参数查询派生优惠参数表信息出错")
		zlog.Error(er.Error(), err)
		return nil, er
	}
	var rst []*SceneDiscount
	for rows.Next() {
		one, err := this.scan(rows)
		if nil != err {
			er := fmt.Errorf("多参数查询派生优惠参数表信息row.Scan()出错")
			zlog.Error(er.Error(), err)
			return nil, er
		}
		rst = append(rst, one)
	}
	return rst, nil
}

// 新增派生优惠参数表信息
// by author Jason
// by time 2017-02-17 10:25:04
func (this SceneDiscount) Add() error {
	paramMap := map[string]interface{}{
		"biz_type":    this.BizType,
		"cust_implvl": this.CustImplvl,
		"cust_size":   this.CustSize,
		"product":     this.Product.ProductCode,
		"rate":        this.Rate,
		"create_user": this.CreateUser,
		"update_user": this.UpdateUser,
		"create_time": util.GetCurrentTime(),
		"update_time": util.GetCurrentTime(),
	}
	err := util.OracleAdd("RPM_PAR_SCENE_DISCOUNT", paramMap)
	if nil != err {
		er := fmt.Errorf("新增派生优惠参数表信息出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

// 更新派生优惠参数表信息
// by author Jason
// by time 2017-02-17 10:25:04
func (this SceneDiscount) Update() error {
	paramMap := map[string]interface{}{
		"biz_type":    this.BizType,
		"cust_implvl": this.CustImplvl,
		"cust_size":   this.CustSize,
		"product":     this.Product.ProductCode,
		"rate":        this.Rate,
		"create_user": this.CreateUser,
		"update_user": this.UpdateUser,
		"update_time": util.GetCurrentTime(),
	}
	whereParam := map[string]interface{}{
		"uuid": this.UUID,
	}
	err := util.OracleUpdate("RPM_PAR_SCENE_DISCOUNT", paramMap, whereParam)
	if nil != err {
		er := fmt.Errorf("更新派生优惠参数表信息出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

// 删除派生优惠参数表信息
// by author Jason
// by time 2017-02-17 10:25:04
func (this SceneDiscount) Delete() error {
	whereParam := map[string]interface{}{
		"uuid": this.UUID,
	}
	err := util.OracleDelete("RPM_PAR_SCENE_DISCOUNT", whereParam)
	if nil != err {
		er := fmt.Errorf("删除派生优惠参数表信息出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

var sceneDiscountTabales string = `
					RPM_PAR_SCENE_DISCOUNT T
			   LEFT JOIN RPM_DIM_PRODUCT PRODUCT
			     ON (T.PRODUCT = PRODUCT.PRODUCT_CODE)`

var sceneDiscountCols map[string]string = map[string]string{
	"T.UUID":        "' '",
	"T.BIZ_TYPE":    "' '",
	"T.CUST_IMPLVL": "' '",
	"T.CUST_SIZE":   "' '",
	"T.PRODUCT":     "' '",
	"T.RATE":        "0",
	"T.CREATE_USER": "' '",
	"T.UPDATE_USER": "' '",
	"T.CREATE_TIME": "sysdate",
	"T.UPDATE_TIME": "sysdate",

	"PRODUCT.PRODUCT_NAME": "' '",
}

var sceneDiscountColsSort = []string{
	"T.UUID",
	"T.BIZ_TYPE",
	"T.CUST_IMPLVL",
	"T.CUST_SIZE",
	"T.RATE",
	"T.CREATE_USER",
	"T.UPDATE_USER",
	"T.CREATE_TIME",
	"T.UPDATE_TIME",

	"T.PRODUCT",
	"PRODUCT.PRODUCT_NAME",
}
