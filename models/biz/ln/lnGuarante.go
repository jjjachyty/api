package ln

import (
	"database/sql"
	"fmt"
	"time"

	"pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/components/zlog"

	modelsUtil "pccqcpa.com.cn/app/rpm/api/models/util"
)

type LnGuarante struct {
	UUID          string    //主键
	BusinessCode  string    //业务编号
	Cust          string    //客户
	Guarante      CustInfo  //保证人
	GuaranteAmout float64   //保证金额
	GuaranteType  string    //保证人类型
	CreateTime    time.Time //创建时间
	UpdateTime    time.Time // 更新时间
	UpdateUser    string    // 更新人
	CreateUser    string    //创建人
}

func (l *LnGuarante) scan(rows *sql.Rows) (*LnGuarante, error) {
	var lnGuarante = new(LnGuarante)
	values := []interface{}{
		&lnGuarante.UUID,
		&lnGuarante.BusinessCode,
		&lnGuarante.Cust,
		&lnGuarante.GuaranteAmout,
		&lnGuarante.GuaranteType,
		&lnGuarante.CreateTime,
		&lnGuarante.UpdateTime,
		&lnGuarante.CreateUser,
		&lnGuarante.UpdateUser,
		&lnGuarante.Guarante.CustCode,
		&lnGuarante.Guarante.CustName,
	}
	err := util.OracleScan(rows, values)
	return lnGuarante, err

}
func (l *LnGuarante) List(param ...map[string]interface{}) (*util.PageData, error) {
	pageDate, rows, err := modelsUtil.List(lnGuaranteTables, lnGuaranteCols, lnGuaranteColsSort, param...)
	defer rows.Close()
	if nil != err {
		er := fmt.Errorf("分页查询保证人出错")
		zlog.Error(er.Error(), err)
		return nil, er
	}
	var lnGuarantes []*LnGuarante
	for rows.Next() {
		lnGuarante, err := l.scan(rows)
		if nil != err {
			er := fmt.Errorf("分页查询保证人rows.Scan()出错")
			zlog.Error(er.Error(), err)
			return nil, er
		}
		lnGuarantes = append(lnGuarantes, lnGuarante)
	}
	pageDate.Rows = lnGuarantes
	return pageDate, nil
}

func (l *LnGuarante) Find(param ...map[string]interface{}) ([]*LnGuarante, error) {
	rows, err := modelsUtil.FindRows(lnGuaranteTables, lnGuaranteCols, lnGuaranteColsSort, param...)
	defer rows.Close()
	if nil != err {
		er := fmt.Errorf("多参数查询保证人出错")
		zlog.Error(er.Error(), err)
		return nil, er
	}
	var lnGuarantes []*LnGuarante

	for rows.Next() {
		lnGuarante, err := l.scan(rows)
		if nil != err {
			er := fmt.Errorf("多参数查询保证人rows.Scan()出错")
			zlog.Error(er.Error(), err)
			return nil, er
		}
		lnGuarantes = append(lnGuarantes, lnGuarante)
	}
	return lnGuarantes, nil
}

func (l *LnGuarante) Add() error {
	param := map[string]interface{}{
		"business_code":  l.BusinessCode,
		"cust":           l.Cust,
		"guarante_amout": l.GuaranteAmout,
		"guarante_type":  l.GuaranteType,
		"create_time":    util.GetCurrentTime(),
		"update_time":    util.GetCurrentTime(),
		"create_user":    l.CreateUser,
		"update_user":    l.UpdateUser,
		"guarante":       l.Guarante.CustCode,
	}
	err := util.OracleAdd(lnGuaranteTableName, param)
	if nil != err {
		er := fmt.Errorf("新增保证人出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

func (l *LnGuarante) Update() error {
	param := map[string]interface{}{
		// "uuid":           l.UUID,
		// "business_code":  l.BusinessCode,
		"cust":           l.Cust,
		"guarante_amout": l.GuaranteAmout,
		"guarante_type":  l.GuaranteType,
		"create_time":    util.GetCurrentTime(),
		"update_time":    util.GetCurrentTime(),
		"update_user":    l.UpdateUser,
		"guarante":       l.Guarante.CustCode,
	}
	whereParam := map[string]interface{}{
		"uuid": l.UUID,
	}
	err := util.OracleUpdate(lnGuaranteTableName, param, whereParam)
	if nil != err {
		er := fmt.Errorf("更新保证人出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}
func (l *LnGuarante) Delete() error {
	whereParam := map[string]interface{}{
		"business_code": l.BusinessCode,
		"guarante":      l.Guarante.CustCode,
	}
	err := util.OracleDelete(lnGuaranteTableName, whereParam)
	if nil != err {
		er := fmt.Errorf("删除保证人出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

func (l *LnGuarante) DeleteByBusinessCode(businessCode string) error {
	whereParam := map[string]interface{}{
		"business_code": businessCode,
	}
	err := util.OracleDelete(lnGuaranteTableName, whereParam)
	if nil != err {
		er := fmt.Errorf("通过业务单号删除保证人出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

var lnGuaranteTableName = "rpm_biz_guarante"

const lnGuaranteTables string = `
	     RPM_BIZ_GUARANTE T
	LEFT JOIN RPM_BIZ_CUST_INFO GUARANTE
	  ON (T.GUARANTE = GUARANTE.CUST_CODE)
	`

var lnGuaranteCols = map[string]string{
	"T.UUID":             "' '",
	"T.BUSINESS_CODE":    "' '",
	"T.CUST":             "' '",
	"T.GUARANTE_AMOUT":   "0",
	"T.GUARANTE_TYPE":    "' '",
	"T.CREATE_TIME":      "sysdate",
	"T.UPDATE_TIME":      "sysdate",
	"T.CREATE_USER":      "' '",
	"T.UPDATE_USER":      "' '",
	"T.GUARANTE":         "' '",
	"GUARANTE.CUST_NAME": "' '",
}
var lnGuaranteColsSort = []string{
	"T.UUID",
	"T.BUSINESS_CODE",
	"T.CUST",
	"T.GUARANTE_AMOUT",
	"T.GUARANTE_TYPE",
	"T.CREATE_TIME",
	"T.UPDATE_TIME",
	"T.CREATE_USER",
	"T.UPDATE_USER",
	"T.GUARANTE",
	"GUARANTE.CUST_NAME",
}
