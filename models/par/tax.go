package par

import (
	"database/sql"
	"fmt"
	"time"

	"pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/components/zlog"

	modelsUtil "pccqcpa.com.cn/app/rpm/api/models/util"
)

// 结构体对应数据库表【RPM_PAR_TAX】税率
/**
 * @apiDefine Tax
 * @apiSuccess {string}     UUID       	主键默认值sys_guid()
 * @apiSuccess {string}   	TaxType    	税率类型
 * @apiSuccess {float64}   	TaxRate    	税率
 * @apiSuccess {time.Time}  StartTime  	生效日期
 * @apiSuccess {string}   	Flag       	生效标志
 * @apiSuccess {time.Time}	CreateTime  创建时间
 * @apiSuccess {string}   	CreateUser  创建人
 * @apiSuccess {time.Time}	UpdateTime  更新时间
 * @apiSuccess {string}   	UpdateUser  更新人
 */
type Tax struct {
	UUID       string    //主键
	TaxType    string    //税率类型
	TaxRate    float64   //税率
	StartTime  time.Time //生效日期
	CreateTime time.Time //创建日期
	UpdateTime time.Time //更新日期
	CreateUser string    //创建人
	UpdateUser string    //更新人
	Flag       string    //生效标志
}

var taxTables = ` 
	RPM_PAR_TAX T
`

var taxCols = map[string]string{
	"T.UUID":        "' '",
	"T.TAX_TYPE":    "' '",
	"T.TAX_RATE":    "0",
	"T.START_TIME":  "sysdate",
	"T.CREATE_TIME": "sysdate",
	"T.UPDATE_TIME": "sysdate",
	"T.CREATE_USER": "' '",
	"T.UPDATE_USER": "' '",
	"T.FLAG":        "' '",
}

var taxColsSort = []string{
	"T.UUID",
	"T.TAX_TYPE",
	"T.TAX_RATE",
	"T.START_TIME",
	"T.CREATE_TIME",
	"T.UPDATE_TIME",
	"T.CREATE_USER",
	"T.UPDATE_USER",
	"T.FLAG",
}

// 结构体scan
func (this *Tax) sanTax(rows *sql.Rows) (*Tax, error) {
	var tax = new(Tax)
	var values = []interface{}{
		&tax.UUID,
		&tax.TaxType,
		&tax.TaxRate,
		&tax.StartTime,
		&tax.CreateTime,
		&tax.UpdateTime,
		&tax.CreateUser,
		&tax.UpdateUser,
		&tax.Flag,
	}
	err := util.OracleScan(rows, values)
	if nil != err {
		return nil, err
	}
	return tax, nil
}

func (this *Tax) SelectTaxByParams(paramMap map[string]interface{}) ([]*Tax, error) {
	rows, err := modelsUtil.FindRows(taxTables, taxCols, taxColsSort, paramMap)
	if nil != err {
		var er = fmt.Errorf("多参数查询税率出错")
		zlog.Error(er.Error(), err)
		return nil, er
	}
	defer rows.Close()
	var taxs []*Tax
	for rows.Next() {
		tax, err := this.sanTax(rows)
		if nil != err {
			var er = fmt.Errorf("多参数查询税率rows.Scan()错误")
			zlog.Error(er.Error(), err)
			return nil, er
		}
		taxs = append(taxs, tax)
	}
	return taxs, nil
}

// 分页操作
// by author Yeqc
// by time 2016-12-21 11:10:11
func (this Tax) List(param ...map[string]interface{}) (*util.PageData, error) {
	pageData, rows, err := modelsUtil.List(taxTables, taxCols, taxColsSort, param...)
	defer rows.Close()
	if nil != err {
		er := fmt.Errorf("分页查询税率信息出错")
		zlog.Error(er.Error(), err)
		return nil, er
	}

	var rst []*Tax
	for rows.Next() {
		one, err := this.sanTax(rows)
		if nil != err {
			er := fmt.Errorf("分页查询税率信息row.Scan()出错")
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
// by time 2016-12-21 11:10:11
func (this Tax) Find(param ...map[string]interface{}) ([]*Tax, error) {
	rows, err := modelsUtil.FindRows(taxTables, taxCols, taxColsSort, param...)
	defer rows.Close()
	if nil != err {
		er := fmt.Errorf("多参数查询税率信息出错")
		zlog.Error(er.Error(), err)
		return nil, er
	}
	var rst []*Tax
	for rows.Next() {
		one, err := this.sanTax(rows)
		if nil != err {
			er := fmt.Errorf("多参数查询税率信息row.Scan()出错")
			zlog.Error(er.Error(), err)
			return nil, er
		}
		rst = append(rst, one)
	}
	return rst, nil
}

// 新增税率信息
// by author Yeqc
// by time 2016-12-21 11:10:11
func (this Tax) Add() error {
	paramMap := map[string]interface{}{
		"tax_type":    this.TaxType,
		"tax_rate":    this.TaxRate,
		"start_time":  this.StartTime,
		"create_time": util.GetCurrentTime(),
		"update_time": util.GetCurrentTime(),
		"create_user": this.CreateUser,
		"update_user": this.UpdateUser,
		"flag":        this.Flag,
	}
	err := util.OracleAdd("RPM_PAR_TAX", paramMap)
	if nil != err {
		er := fmt.Errorf("新增税率信息出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

// 更新税率信息
// by author Yeqc
// by time 2016-12-21 11:10:11
func (this Tax) Update() error {
	paramMap := map[string]interface{}{
		"tax_type":    this.TaxType,
		"tax_rate":    this.TaxRate,
		"start_time":  this.StartTime,
		"update_time": util.GetCurrentTime(),
		"create_user": this.CreateUser,
		"update_user": this.UpdateUser,
		"flag":        this.Flag,
	}
	whereParam := map[string]interface{}{
		"uuid": this.UUID,
	}
	err := util.OracleUpdate("RPM_PAR_TAX", paramMap, whereParam)
	if nil != err {
		er := fmt.Errorf("更新税率信息出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

// 删除税率信息
// by author Yeqc
// by time 2016-12-21 11:10:11
func (this Tax) Delete() error {
	whereParam := map[string]interface{}{
		"uuid": this.UUID,
	}
	err := util.OracleDelete("RPM_PAR_TAX", whereParam)
	if nil != err {
		er := fmt.Errorf("删除税率信息出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}
