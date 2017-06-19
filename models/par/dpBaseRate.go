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

// 结构体对应数据库表【RPM_PAR_DP_BASE_RATE】存款基准利率及挂牌-存款模块
// by author Yeqc
// by time 2017-01-05 14:25:41
type DpBaseRate struct {
	UUID       string      // 唯一建默认值sys_guid()
	Term       float64     // 期限 引用字典Term
	Name       string      // 名称
	Product    dim.Product // 产品 结构体
	Type       string      // 类型（基准 0 /挂牌 1）引用字典DpBaseRateType
	Rate       float64     // 利率 百分比
	Flag       string      // 生效标志 引用字典Flag
	CreateTime time.Time   // 创建时间
	UpdateTime time.Time   // 更新时间
	CreateUser string      // 创建人
	UpdateUser string      // 更新人
}

// 结构体scan
// by author Yeqc
// by time 2017-01-05 14:25:41
func (this DpBaseRate) scan(rows *sql.Rows) (*DpBaseRate, error) {
	var one = new(DpBaseRate)
	values := []interface{}{
		&one.UUID,
		&one.Term,
		&one.Name,
		// &one.Product,
		&one.Type,
		&one.Rate,
		&one.Flag,
		&one.CreateTime,
		&one.UpdateTime,
		&one.CreateUser,
		&one.UpdateUser,

		&one.Product.ProductCode,
		&one.Product.ProductName,
		&one.Product.ProductType,
		&one.Product.ProductTypeDesc,
		&one.Product.ProductLevel,
	}
	err := util.OracleScan(rows, values)
	return one, err
}

// 分页操作
// by author Yeqc
// by time 2017-01-05 14:25:41
func (this DpBaseRate) List(param ...map[string]interface{}) (*util.PageData, error) {
	pageData, rows, err := modelsUtil.List(dpBaseRateTabales, dpBaseRateCols, dpBaseRateColsSort, param...)
	defer rows.Close()
	if nil != err {
		er := fmt.Errorf("分页查询存款基准利率及挂牌-存款模块信息出错")
		zlog.Error(er.Error(), err)
		return nil, er
	}

	var rst []*DpBaseRate
	for rows.Next() {
		one, err := this.scan(rows)
		if nil != err {
			er := fmt.Errorf("分页查询存款基准利率及挂牌-存款模块信息row.Scan()出错")
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
// by time 2017-01-05 14:25:41
func (this DpBaseRate) Find(param ...map[string]interface{}) ([]*DpBaseRate, error) {
	rows, err := modelsUtil.FindRows(dpBaseRateTabales, dpBaseRateCols, dpBaseRateColsSort, param...)
	defer rows.Close()
	if nil != err {
		er := fmt.Errorf("多参数查询存款基准利率及挂牌-存款模块信息出错")
		zlog.Error(er.Error(), err)
		return nil, er
	}
	var rst []*DpBaseRate
	for rows.Next() {
		one, err := this.scan(rows)
		if nil != err {
			er := fmt.Errorf("多参数查询存款基准利率及挂牌-存款模块信息row.Scan()出错")
			zlog.Error(er.Error(), err)
			return nil, er
		}
		rst = append(rst, one)
	}
	return rst, nil
}

// 新增存款基准利率及挂牌-存款模块信息
// by author Yeqc
// by time 2017-01-05 14:25:41
func (this DpBaseRate) Add() error {
	paramMap := map[string]interface{}{
		"term":        this.Term,
		"name":        this.Name,
		"product":     this.Product.ProductCode,
		"type":        this.Type,
		"rate":        this.Rate,
		"flag":        this.Flag,
		"create_time": util.GetCurrentTime(),
		"update_time": util.GetCurrentTime(),
		"create_user": this.CreateUser,
		"update_user": this.UpdateUser,
	}
	err := util.OracleAdd("RPM_PAR_DP_BASE_RATE", paramMap)
	if nil != err {
		er := fmt.Errorf("新增存款基准利率及挂牌-存款模块信息出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

// 更新存款基准利率及挂牌-存款模块信息
// by author Yeqc
// by time 2017-01-05 14:25:41
func (this DpBaseRate) Update() error {
	paramMap := map[string]interface{}{
		"term":        this.Term,
		"name":        this.Name,
		"product":     this.Product.ProductCode,
		"type":        this.Type,
		"rate":        this.Rate,
		"flag":        this.Flag,
		"update_time": util.GetCurrentTime(),
		"update_user": this.UpdateUser,
	}
	whereParam := map[string]interface{}{
		"uuid": this.UUID,
	}
	err := util.OracleUpdate("RPM_PAR_DP_BASE_RATE", paramMap, whereParam)
	if nil != err {
		er := fmt.Errorf("更新存款基准利率及挂牌-存款模块信息出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

// 删除存款基准利率及挂牌-存款模块信息
// by author Yeqc
// by time 2017-01-05 14:25:41
func (this DpBaseRate) Delete() error {
	whereParam := map[string]interface{}{
		"uuid": this.UUID,
	}
	err := util.OracleDelete("RPM_PAR_DP_BASE_RATE", whereParam)
	if nil != err {
		er := fmt.Errorf("删除存款基准利率及挂牌-存款模块信息出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

var dpBaseRateTabales string = `
	     RPM_PAR_DP_BASE_RATE T 
	LEFT JOIN RPM_DIM_PRODUCT PRODUCT
	ON (PRODUCT.PRODUCT_CODE=T.PRODUCT)`

var dpBaseRateCols map[string]string = map[string]string{
	"T.UUID": "' '",
	"T.TERM": "0",
	"T.NAME": "' '",
	// "T.PRODUCT":     "' '",
	"T.TYPE":        "' '",
	"T.RATE":        "0",
	"T.FLAG":        "' '",
	"T.CREATE_TIME": "sysdate",
	"T.UPDATE_TIME": "sysdate",
	"T.CREATE_USER": "' '",
	"T.UPDATE_USER": "' '",

	"PRODUCT.PRODUCT_CODE":      "' '",
	"PRODUCT.PRODUCT_NAME":      "' '",
	"PRODUCT.PRODUCT_TYPE":      "' '",
	"PRODUCT.PRODUCT_TYPE_DESC": "' '",
	"PRODUCT.PRODUCT_LEVEL":     "' '",
}

var dpBaseRateColsSort = []string{
	"T.UUID",
	"T.TERM",
	"T.NAME",
	// "T.PRODUCT",
	"T.TYPE",
	"T.RATE",
	"T.FLAG",
	"T.CREATE_TIME",
	"T.UPDATE_TIME",
	"T.CREATE_USER",
	"T.UPDATE_USER",

	"PRODUCT.PRODUCT_CODE",
	"PRODUCT.PRODUCT_NAME",
	"PRODUCT.PRODUCT_TYPE",
	"PRODUCT.PRODUCT_TYPE_DESC",
	"PRODUCT.PRODUCT_LEVEL",
}
