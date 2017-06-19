package par

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

// 结构体对应数据库表【RPM_PAR_DP_TARGET_RATE】目标收益率-存款标准化
// by author Jason
// by time 2017-02-15 10:19:49
type DpTargetRate struct {
	UUID       string      // 唯一字段-默认sys_guid()
	Organ      sys.Organ   // 机构
	Product    dim.Product // 产品
	Term       float64     // 期限
	TargetRate float64     // 目标收益率
	Flag       string      // 生效标志
	CreateUser string      // 创建人
	UpdateUser string      // 更新人
	CreateTime time.Time   // 创建时间
	UpdateTime time.Time   // 更新时间
}

// 结构体scan
// by author Jason
// by time 2017-02-15 10:19:49
func (this DpTargetRate) scan(rows *sql.Rows) (*DpTargetRate, error) {
	var one = new(DpTargetRate)
	values := []interface{}{
		&one.UUID,
		&one.Term,
		&one.TargetRate,
		&one.Flag,
		&one.CreateUser,
		&one.UpdateUser,
		&one.CreateTime,
		&one.UpdateTime,

		&one.Product.ProductCode,
		&one.Product.ProductName,

		&one.Organ.OrganCode,
		&one.Organ.OrganName,
	}
	err := util.OracleScan(rows, values)
	return one, err
}

// 分页操作
// by author Jason
// by time 2017-02-15 10:19:49
func (this DpTargetRate) List(param ...map[string]interface{}) (*util.PageData, error) {
	pageData, rows, err := modelsUtil.List(dpTargetRateTabales, dpTargetRateCols, dpTargetRateColsSort, param...)
	defer rows.Close()
	if nil != err {
		er := fmt.Errorf("分页查询目标收益率-存款标准化信息出错")
		zlog.Error(er.Error(), err)
		return nil, er
	}

	var rst []*DpTargetRate
	for rows.Next() {
		one, err := this.scan(rows)
		if nil != err {
			er := fmt.Errorf("分页查询目标收益率-存款标准化信息row.Scan()出错")
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
// by time 2017-02-15 10:19:49
func (this DpTargetRate) Find(param ...map[string]interface{}) ([]*DpTargetRate, error) {
	rows, err := modelsUtil.FindRows(dpTargetRateTabales, dpTargetRateCols, dpTargetRateColsSort, param...)
	defer rows.Close()
	if nil != err {
		er := fmt.Errorf("多参数查询目标收益率-存款标准化信息出错")
		zlog.Error(er.Error(), err)
		return nil, er
	}
	var rst []*DpTargetRate
	for rows.Next() {
		one, err := this.scan(rows)
		if nil != err {
			er := fmt.Errorf("多参数查询目标收益率-存款标准化信息row.Scan()出错")
			zlog.Error(er.Error(), err)
			return nil, er
		}
		rst = append(rst, one)
	}
	return rst, nil
}

// 新增目标收益率-存款标准化信息
// by author Jason
// by time 2017-02-15 10:19:49
func (this DpTargetRate) Add() error {
	paramMap := map[string]interface{}{
		"organ":       this.Organ.OrganCode,
		"product":     this.Product.ProductCode,
		"term":        this.Term,
		"target_rate": this.TargetRate,
		"flag":        this.Flag,
		"create_user": this.CreateUser,
		"update_user": this.UpdateUser,
		"create_time": util.GetCurrentTime(),
		"update_time": util.GetCurrentTime(),
	}
	err := util.OracleAdd("RPM_PAR_DP_TARGET_RATE", paramMap)
	if nil != err {
		er := fmt.Errorf("新增目标收益率-存款标准化信息出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

// 更新目标收益率-存款标准化信息
// by author Jason
// by time 2017-02-15 10:19:49
func (this DpTargetRate) Update() error {
	paramMap := map[string]interface{}{
		"organ":       this.Organ.OrganCode,
		"product":     this.Product.ProductCode,
		"term":        this.Term,
		"target_rate": this.TargetRate,
		"flag":        this.Flag,
		"create_user": this.CreateUser,
		"update_user": this.UpdateUser,
		"update_time": util.GetCurrentTime(),
	}
	whereParam := map[string]interface{}{
		"uuid": this.UUID,
	}
	err := util.OracleUpdate("RPM_PAR_DP_TARGET_RATE", paramMap, whereParam)
	if nil != err {
		er := fmt.Errorf("更新目标收益率-存款标准化信息出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

// 删除目标收益率-存款标准化信息
// by author Jason
// by time 2017-02-15 10:19:49
func (this DpTargetRate) Delete() error {
	whereParam := map[string]interface{}{
		"uuid": this.UUID,
	}
	err := util.OracleDelete("RPM_PAR_DP_TARGET_RATE", whereParam)
	if nil != err {
		er := fmt.Errorf("删除目标收益率-存款标准化信息出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

var dpTargetRateTabales string = `
						RPM_PAR_DP_TARGET_RATE T
				   LEFT JOIN RPM_DIM_PRODUCT PRODUCT
				     ON (T.PRODUCT = PRODUCT.PRODUCT_CODE)
				   LEFT JOIN SYS_SEC_ORGAN ORGAN
				     ON (T.ORGAN = ORGAN.ORGAN_CODE)`

var dpTargetRateCols map[string]string = map[string]string{
	"T.UUID":        "' '",
	"T.TERM":        "0",
	"T.TARGET_RATE": "0",
	"T.FLAG":        "' '",
	"T.CREATE_USER": "' '",
	"T.UPDATE_USER": "' '",
	"T.CREATE_TIME": "sysdate",
	"T.UPDATE_TIME": "sysdate",

	"T.PRODUCT":            "' '",
	"PRODUCT.PRODUCT_NAME": "' '",
	"T.ORGAN":              "' '",
	"ORGAN.ORGAN_NAME":     "' '",
}

var dpTargetRateColsSort = []string{
	"T.UUID",
	"T.TERM",
	"T.TARGET_RATE",
	"T.FLAG",
	"T.CREATE_USER",
	"T.UPDATE_USER",
	"T.CREATE_TIME",
	"T.UPDATE_TIME",

	"T.PRODUCT",
	"PRODUCT.PRODUCT_NAME",
	"T.ORGAN",
	"ORGAN.ORGAN_NAME",
}
