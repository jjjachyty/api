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

// 结构体对应数据库表【RPM_PAR_RL_EC】零售资本成本率-零售贷款模块
// by author Yeqc
// by time 2016-12-30 09:32:23
/**
 * @apiDefine RlEc
 * @apiSuccess {string}     	UUID            主键默认值sys_guid()
 * @apiSuccess {Product}   		ProductCode    	产品
 * @apiSuccess {float64}   		Capitail      	资本充足率
 * @apiSuccess {float64}   		RiskWeight   	产品风险权重
 * @apiSuccess {string}   		Flag    		生效标志
 * @apiSuccess {time.Time}   	StartTime      	生效时间
 * @apiSuccess {time.Time}		CreateTime      创建时间
 * @apiSuccess {string}   		CreateUser      创建人
 * @apiSuccess {time.Time}		UpdateTime      更新时间
 * @apiSuccess {string}   		UpdateUser      更新人
 */
type RlEc struct {
	UUID        string      // 默认值sys_guid()
	ProductCode dim.Product // 产品
	Capitail    float64     // 资本充足率
	RiskWeight  float64     // 产品风险权重
	Flag        string      // 生效标志
	CreateTime  time.Time   // 创建时间
	CreateUser  string      // 创建人
	UpdateTime  time.Time   // 更新时间
	UpdateUser  string      // 更新人
}

// 结构体scan
// by author Yeqc
// by time 2016-12-30 09:32:23
func (this RlEc) scan(rows *sql.Rows) (*RlEc, error) {
	var one = new(RlEc)
	values := []interface{}{
		&one.UUID,
		// &one.ProductCode,
		&one.Capitail,
		&one.RiskWeight,
		&one.Flag,
		&one.CreateTime,
		&one.CreateUser,
		&one.UpdateTime,
		&one.UpdateUser,

		&one.ProductCode.ProductCode,
		&one.ProductCode.ProductName,
		&one.ProductCode.ProductType,
		&one.ProductCode.ProductTypeDesc,
		&one.ProductCode.ProductLevel,
	}
	err := util.OracleScan(rows, values)
	return one, err
}

// 分页操作
// by author Yeqc
// by time 2016-12-30 09:32:23
func (this RlEc) List(param ...map[string]interface{}) (*util.PageData, error) {
	pageData, rows, err := modelsUtil.List(rlEcTabales, rlEcCols, rlEcColsSort, param...)
	defer rows.Close()
	if nil != err {
		er := fmt.Errorf("分页查询零售资本成本率-零售贷款模块信息出错")
		zlog.Error(er.Error(), err)
		return nil, er
	}

	var rst []*RlEc
	for rows.Next() {
		one, err := this.scan(rows)
		if nil != err {
			er := fmt.Errorf("分页查询零售资本成本率-零售贷款模块信息row.Scan()出错")
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
// by time 2016-12-30 09:32:23
func (this RlEc) Find(param ...map[string]interface{}) ([]*RlEc, error) {
	rows, err := modelsUtil.FindRows(rlEcTabales, rlEcCols, rlEcColsSort, param...)
	defer rows.Close()
	if nil != err {
		er := fmt.Errorf("多参数查询零售资本成本率-零售贷款模块信息出错")
		zlog.Error(er.Error(), err)
		return nil, er
	}
	var rst []*RlEc
	for rows.Next() {
		one, err := this.scan(rows)
		if nil != err {
			er := fmt.Errorf("多参数查询零售资本成本率-零售贷款模块信息row.Scan()出错")
			zlog.Error(er.Error(), err)
			return nil, er
		}
		rst = append(rst, one)
	}
	return rst, nil
}

// 新增零售资本成本率-零售贷款模块信息
// by author Yeqc
// by time 2016-12-30 09:32:23
func (this RlEc) Add() error {
	paramMap := map[string]interface{}{
		"product_code": this.ProductCode.ProductCode,
		"capitail":     this.Capitail,
		"risk_weight":  this.RiskWeight,
		"flag":         this.Flag,
		"create_time":  util.GetCurrentTime(),
		"create_user":  this.CreateUser,
		"update_time":  util.GetCurrentTime(),
		"update_user":  this.UpdateUser,
	}
	err := util.OracleAdd("RPM_PAR_RL_EC", paramMap)
	if nil != err {
		er := fmt.Errorf("新增零售资本成本率-零售贷款模块信息出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

// 更新零售资本成本率-零售贷款模块信息
// by author Yeqc
// by time 2016-12-30 09:32:23
func (this RlEc) Update() error {
	paramMap := map[string]interface{}{
		"product_code": this.ProductCode.ProductCode,
		"capitail":     this.Capitail,
		"risk_weight":  this.RiskWeight,
		"flag":         this.Flag,
		"update_time":  util.GetCurrentTime(),
		"update_user":  this.UpdateUser,
	}
	whereParam := map[string]interface{}{
		"uuid": this.UUID,
	}
	err := util.OracleUpdate("RPM_PAR_RL_EC", paramMap, whereParam)
	if nil != err {
		er := fmt.Errorf("更新零售资本成本率-零售贷款模块信息出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

// 删除零售资本成本率-零售贷款模块信息
// by author Yeqc
// by time 2016-12-30 09:32:23
func (this RlEc) Delete() error {
	whereParam := map[string]interface{}{
		"uuid": this.UUID,
	}
	err := util.OracleDelete("RPM_PAR_RL_EC", whereParam)
	if nil != err {
		er := fmt.Errorf("删除零售资本成本率-零售贷款模块信息出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

var rlEcTabales string = `
	RPM_PAR_RL_EC T
	LEFT JOIN RPM_DIM_PRODUCT PRODUCTCODE
	ON (PRODUCTCODE.PRODUCT_CODE=T.PRODUCT_CODE)
	`

var rlEcCols map[string]string = map[string]string{
	"T.UUID": "' '",
	// "T.PRODUCT_CODE": "' '",
	"T.CAPITAIL":    "0",
	"T.RISK_WEIGHT": "0",
	"T.FLAG":        "' '",
	"T.CREATE_TIME": "sysdate",
	"T.CREATE_USER": "' '",
	"T.UPDATE_TIME": "sysdate",
	"T.UPDATE_USER": "' '",

	"productcode.product_code":      "' '",
	"productcode.product_name":      "' '",
	"productcode.product_type":      "' '",
	"productcode.product_type_desc": "' '",
	"productcode.product_level":     "' '",
}

var rlEcColsSort = []string{
	"T.UUID",
	// "T.PRODUCT_CODE",
	"T.CAPITAIL",
	"T.RISK_WEIGHT",
	"T.FLAG",
	"T.CREATE_TIME",
	"T.CREATE_USER",
	"T.UPDATE_TIME",
	"T.UPDATE_USER",

	"productcode.product_code",
	"productcode.product_name",
	"productcode.product_type",
	"productcode.product_type_desc",
	"productcode.product_level",
}
