package par

import (
	"database/sql"
	"fmt"
	"time"

	"pccqcpa.com.cn/app/rpm/api/models/dim"
	"pccqcpa.com.cn/app/rpm/api/models/sys"
	"pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/components/zlog"

	modelsUtil "pccqcpa.com.cn/app/rpm/api/models/util"
)

// 经济资本占用率 RPM_PAR_EC
/**
 * @apiDefine Ec
 * @apiSuccess {string}    	 	Uuid            主键默认值sys_guid()
 * @apiSuccess {Organ}   	 	Organ    		机构
 * @apiSuccess {Product}   		Product      	产品
 * @apiSuccess {float64}   		EcRate   		经济资本占用率
 * @apiSuccess {string}   		Flag    		生效标志
 * @apiSuccess {time.Time }   	StartTime      	生效日期
 * @apiSuccess {time.Time}		CreateTime      创建时间
 * @apiSuccess {string}   		CreateUser      创建人
 * @apiSuccess {time.Time}		UpdateTime      更新时间
 * @apiSuccess {string}   		UpdateUser      更新人
 */
type Ec struct {
	UUID       string      //主键
	Organ      sys.Organ   //机构
	Product    dim.Product //产品
	EcRate     float64     //经济资本占用率
	Flag       string      //生效标志
	StartTime  time.Time   //生效日期
	CreateTime time.Time   //创建日期
	UpdateTime time.Time   //更新日期
	CreateUser string      //创建人
	UpdateUser string      //更新人
}

var ecTables string = `
	RPM_PAR_EC T
	LEFT JOIN SYS_SEC_ORGAN ORGAN 
	ON (T.ORGAN = ORGAN.ORGAN_CODE)
	LEFT JOIN RPM_DIM_PRODUCT PRODUCT
	ON (T.PRODUCT = PRODUCT.PRODUCT_CODE)
`

var ecCols = map[string]string{
	"T.UUID":        "' '",
	"T.EC_RATE":     "0",
	"T.FLAG":        "' '",
	"T.START_TIME":  "sysdate",
	"T.CREATE_TIME": "sysdate",
	"T.UPDATE_TIME": "sysdate",
	"T.CREATE_USER": "' '",
	"T.UPDATE_USER": "' '",

	"T.ORGAN":          "' '",
	"ORGAN.ORGAN_NAME": "' '",

	"T.PRODUCT":            "' '",
	"PRODUCT.PRODUCT_NAME": "' '",
}

var ecColsSort = []string{
	"T.UUID",
	"T.EC_RATE",
	"T.FLAG",
	"T.START_TIME",
	"T.CREATE_TIME",
	"T.UPDATE_TIME",
	"T.CREATE_USER",
	"T.UPDATE_USER",

	"T.ORGAN",
	"ORGAN.ORGAN_NAME",

	"T.PRODUCT",
	"PRODUCT.PRODUCT_NAME",
}

func (this *Ec) scanEc(rows *sql.Rows) (*Ec, error) {
	var ec = new(Ec)
	var values = []interface{}{
		&ec.UUID,
		&ec.EcRate,
		&ec.Flag,
		&ec.StartTime,
		&ec.CreateTime,
		&ec.UpdateTime,
		&ec.CreateUser,
		&ec.UpdateUser,
		&ec.Organ.OrganCode,
		&ec.Organ.OrganName,
		&ec.Product.ProductCode,
		&ec.Product.ProductName,
	}
	err := util.OracleScan(rows, values)
	if nil != err {
		return nil, err
	}
	return ec, nil
}

// 分页操作
// by author Yeqc
// by time 2016-12-19 10:21:15
func (this Ec) List(param ...map[string]interface{}) (*util.PageData, error) {
	pageData, rows, err := modelsUtil.List(ecTables, ecCols, ecColsSort, param...)
	defer rows.Close()
	if nil != err {
		er := fmt.Errorf("分页查询经济资本占用率信息出错")
		zlog.Error(er.Error(), err)
		return nil, er
	}

	var rst []*Ec
	for rows.Next() {
		one, err := this.scanEc(rows)
		if nil != err {
			er := fmt.Errorf("分页查询经济资本占用率信息row.Scan()出错")
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
// by time 2016-12-19 10:21:15
func (this *Ec) Find(param ...map[string]interface{}) ([]*Ec, error) {
	rows, err := modelsUtil.FindRows(ecTables, ecCols, ecColsSort, param...)
	defer rows.Close()
	if nil != err {
		er := fmt.Errorf("多参数查询经济资本占用率信息出错")
		zlog.Error(er.Error(), err)
		return nil, er
	}
	var rst []*Ec
	for rows.Next() {
		one, err := this.scanEc(rows)
		if nil != err {
			er := fmt.Errorf("多参数查询经济资本占用率信息row.Scan()出错")
			zlog.Error(er.Error(), err)
			return nil, er
		}
		rst = append(rst, one)
	}
	return rst, nil
}

// 新增经济资本占用率信息
// by author Yeqc
// by time 2016-12-19 10:21:15
func (this Ec) Add() error {
	paramMap := map[string]interface{}{
		"organ":       this.Organ.OrganCode,
		"product":     this.Product.ProductCode,
		"ec_rate":     this.EcRate,
		"flag":        this.Flag,
		"start_time":  this.StartTime,
		"create_time": util.GetCurrentTime(),
		"update_time": util.GetCurrentTime(),
		"create_user": this.CreateUser,
		"update_user": this.UpdateUser,
	}
	err := util.OracleAdd("RPM_PAR_EC", paramMap)
	if nil != err {
		er := fmt.Errorf("新增经济资本占用率信息出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

// 更新经济资本占用率信息
// by author Yeqc
// by time 2016-12-19 10:21:15
func (this Ec) Update() error {
	paramMap := map[string]interface{}{
		"organ":       this.Organ.OrganCode,
		"product":     this.Product.ProductCode,
		"ec_rate":     this.EcRate,
		"flag":        this.Flag,
		"start_time":  this.StartTime,
		"update_time": util.GetCurrentTime(),
		"update_user": this.UpdateUser,
	}
	whereParam := map[string]interface{}{
		"uuid": this.UUID,
	}
	err := util.OracleUpdate("RPM_PAR_EC", paramMap, whereParam)
	if nil != err {
		er := fmt.Errorf("更新经济资本占用率信息出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

// 删除经济资本占用率信息
// by author Yeqc
// by time 2016-12-19 10:21:15
func (this Ec) Delete() error {
	whereParam := map[string]interface{}{
		"uuid": this.UUID,
	}
	err := util.OracleDelete("RPM_PAR_EC", whereParam)
	if nil != err {
		er := fmt.Errorf("删除经济资本占用率信息出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

func (this *Ec) SelectEcByParmas(paramMap map[string]interface{}) ([]*Ec, error) {
	rows, err := modelsUtil.FindRows(ecTables, ecCols, ecColsSort, paramMap)
	if nil != err {
		er := fmt.Errorf("多参数查询经济资本占用率出错")
		zlog.Error(er.Error(), err)
		return nil, er
	}
	defer rows.Close()
	var ecs []*Ec
	for rows.Next() {
		ec, err := this.scanEc(rows)
		if nil != err {
			var er = fmt.Errorf("查询经济资本占用率rows.Scan()错误")
			zlog.Error(er.Error(), err)
			return nil, er
		}
		ecs = append(ecs, ec)
	}
	return ecs, nil

}
