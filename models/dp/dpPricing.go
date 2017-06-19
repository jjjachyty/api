package dp

import (
	"database/sql"
	"fmt"
	"time"

	modelsUtil "pccqcpa.com.cn/app/rpm/api/models/util"
	"pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/components/zlog"
)

// 结构体对应数据库表【RPM_BIZ_DP_PRICING】存款定价表
// by author Jason
// by time 2016-12-01 10:24:52
type DpPricing struct {
	UUID         string    // 唯一键，默认值sys_guid()
	BusinessCode string    // 业务编号
	Procudt      string    // 产品
	Organ        string    // 机构
	Term         float64   // 期限（单位：天）
	Amount       float64   // 存款金额
	Channel      string    // 存款渠道
	DpRate       float64   // 存款利率
	Status       string    // 状态信息
	CreateTime   time.Time // 创建时间
	CreateUser   string    // 创建人
	UpdateTime   time.Time // 更新时间
	UpdateUser   string    // 更新人
}

// 结构体scan
// by author Jason
// by time 2016-12-01 10:24:52
func (this DpPricing) scan(rows *sql.Rows) (*DpPricing, error) {
	var one = new(DpPricing)
	values := []interface{}{
		&one.UUID,
		&one.BusinessCode,
		&one.Procudt,
		&one.Organ,
		&one.Term,
		&one.Amount,
		&one.Channel,
		&one.DpRate,
		&one.Status,
		&one.CreateTime,
		&one.CreateUser,
		&one.UpdateTime,
		&one.UpdateUser,
	}
	err := util.OracleScan(rows, values)
	return one, err
}

// 分页操作
// by author Jason
// by time 2016-12-01 10:24:52
func (this DpPricing) List(param ...map[string]interface{}) (*util.PageData, error) {
	pageData, rows, err := modelsUtil.List(dpPricingTabales, dpPricingCols, dpPricingColsSort, param...)
	defer rows.Close()
	if nil != err {
		er := fmt.Errorf("分页查询存款定价表信息出错")
		zlog.Error(er.Error(), err)
		return nil, er
	}

	var rst []*DpPricing
	for rows.Next() {
		one, err := this.scan(rows)
		if nil != err {
			er := fmt.Errorf("分页查询存款定价表信息row.Scan()出错")
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
// by time 2016-12-01 10:24:52
func (this DpPricing) Find(param ...map[string]interface{}) ([]*DpPricing, error) {
	rows, err := modelsUtil.FindRows(dpPricingTabales, dpPricingCols, dpPricingColsSort, param...)
	defer rows.Close()
	if nil != err {
		er := fmt.Errorf("多参数查询存款定价表信息出错")
		zlog.Error(er.Error(), err)
		return nil, er
	}
	var rst []*DpPricing
	for rows.Next() {
		one, err := this.scan(rows)
		if nil != err {
			er := fmt.Errorf("多参数查询存款定价表信息row.Scan()出错")
			zlog.Error(er.Error(), err)
			return nil, er
		}
		rst = append(rst, one)
	}
	return rst, nil
}

// 新增存款定价表信息
// by author Jason
// by time 2016-12-01 10:24:52
func (this DpPricing) Add() error {
	paramMap := map[string]interface{}{
		"business_code": this.BusinessCode,
		"procudt":       this.Procudt,
		"organ":         this.Organ,
		"term":          this.Term,
		"amount":        this.Amount,
		"channel":       this.Channel,
		"dp_rate":       this.DpRate,
		"status":        this.Status,
		"create_time":   util.GetCurrentTime(),
		"create_user":   this.CreateUser,
		"update_time":   util.GetCurrentTime(),
		"update_user":   this.UpdateUser,
	}
	err := util.OracleAdd("RPM_BIZ_DP_PRICING", paramMap)
	if nil != err {
		er := fmt.Errorf("新增存款定价表信息出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

// 更新存款定价表信息
// by author Jason
// by time 2016-12-01 10:24:52
func (this DpPricing) Update() error {
	paramMap := map[string]interface{}{
		"business_code": this.BusinessCode,
		"procudt":       this.Procudt,
		"organ":         this.Organ,
		"term":          this.Term,
		"amount":        this.Amount,
		"channel":       this.Channel,
		"dp_rate":       this.DpRate,
		"status":        this.Status,
		"update_time":   util.GetCurrentTime(),
		"update_user":   this.UpdateUser,
	}
	whereParam := map[string]interface{}{
		"uuid": this.UUID,
	}
	err := util.OracleUpdate("RPM_BIZ_DP_PRICING", paramMap, whereParam)
	if nil != err {
		er := fmt.Errorf("更新存款定价表信息出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

// 删除存款定价表信息
// by author Jason
// by time 2016-12-01 10:24:52
func (this DpPricing) Delete() error {
	whereParam := map[string]interface{}{
		"business_code": this.BusinessCode,
	}
	err := util.OracleDelete("RPM_BIZ_DP_PRICING", whereParam)
	if nil != err {
		er := fmt.Errorf("删除存款定价表信息出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

var dpPricingTabales string = `RPM_BIZ_DP_PRICING T`

var dpPricingCols map[string]string = map[string]string{
	"T.UUID":          "' '",
	"T.BUSINESS_CODE": "' '",
	"T.PROCUDT":       "' '",
	"T.ORGAN":         "' '",
	"T.TERM":          "0",
	"T.AMOUNT":        "0",
	"T.CHANNEL":       "' '",
	"T.DP_RATE":       "0",
	"T.STATUS":        "' '",
	"T.CREATE_TIME":   "sysdate",
	"T.CREATE_USER":   "' '",
	"T.UPDATE_TIME":   "sysdate",
	"T.UPDATE_USER":   "' '",
}

var dpPricingColsSort = []string{
	"T.UUID",
	"T.BUSINESS_CODE",
	"T.PROCUDT",
	"T.ORGAN",
	"T.TERM",
	"T.AMOUNT",
	"T.CHANNEL",
	"T.DP_RATE",
	"T.STATUS",
	"T.CREATE_TIME",
	"T.CREATE_USER",
	"T.UPDATE_TIME",
	"T.UPDATE_USER",
}
