package par

import (
	"database/sql"
	"time"

	"fmt"
	"pccqcpa.com.cn/app/rpm/api/models/dim"
	"pccqcpa.com.cn/app/rpm/api/models/sys"
	modelsUtil "pccqcpa.com.cn/app/rpm/api/models/util"
	"pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/components/zlog"
)

// 结构体对应数据库表【RPM_PAR_OC】运营成本率
/**
 * @apiDefine Oc
 * @apiSuccess {string}     Uuid            主键默认值sys_guid()
 * @apiSuccess {Organ}   	Organ    		机构
 * @apiSuccess {Product}   	Product      	产品
 * @apiSuccess {string}   	CustSize   		客户规模
 * @apiSuccess {float64}   	AutoOc    		费用率
 * @apiSuccess {float64}   	ManualOc      	手工值
 * @apiSuccess {string}   	Flag      		生效标志
 * @apiSuccess {string}   	ParamType      	参数类型
 * @apiSuccess {time.Time}  StartTime      	生效日期
 * @apiSuccess {time.Time}	CreateTime      创建时间
 * @apiSuccess {string}   	CreateUser      创建人
 * @apiSuccess {time.Time}	UpdateTime      更新时间
 * @apiSuccess {string}   	UpdateUser      更新人
 */
type Oc struct {
	UUID       string      //主键
	Organ      sys.Organ   //机构
	Product    dim.Product //产品
	CustSize   string      //客户规模
	AutoOc     float64     //费用率
	ManualOc   float64     //手工值
	Flag       string      //生效标志
	ParamType  string      //参数类型
	StartTime  time.Time   //生效日期
	CreateTime time.Time   //创建时间
	CreateUser string      //创建人
	UpdateTime time.Time   //更新时间
	UpdateUser string      //更新人
}

var ocTables = `
	     RPM_PAR_OC T
	LEFT JOIN SYS_SEC_ORGAN ORGAN
	ON (T.ORGAN = ORGAN.ORGAN_CODE)
	LEFT JOIN RPM_DIM_PRODUCT PRODUCT
	ON (T.PRODUCT = PRODUCT.PRODUCT_CODE)
`

var ocCols = map[string]string{
	"T.UUID":        "' '",
	"T.CUST_SIZE":   "' '",
	"T.AUTO_OC":     "0",
	"T.MANUAL_OC":   "0",
	"T.FLAG":        "' '",
	"T.PARAM_TYPE":  "' '",
	"T.START_TIME":  "sysdate",
	"T.CREATE_TIME": "sysdate",
	"T.CREATE_USER": "' '",
	"T.UPDATE_TIME": "sysdate",
	"T.UPDATE_USER": "' '",

	"T.ORGAN":            "' '",
	"ORGAN.ORGAN_NAME":   "' '",
	"ORGAN.PARENT_ORGAN": "' '",

	"T.PRODUCT":            "' '",
	"PRODUCT.PRODUCT_NAME": "' '",
}

var ocColsSort = []string{
	"T.UUID",
	"T.CUST_SIZE",
	"T.AUTO_OC",
	"T.MANUAL_OC",
	"T.FLAG",
	"T.PARAM_TYPE",
	"T.START_TIME",
	"T.CREATE_TIME",
	"T.CREATE_USER",
	"T.UPDATE_TIME",
	"T.UPDATE_USER",

	"T.ORGAN",
	"ORGAN.ORGAN_NAME",
	"ORGAN.PARENT_ORGAN",

	"T.PRODUCT",
	"PRODUCT.PRODUCT_NAME",
}

func (o *Oc) scanOc(rows *sql.Rows) (*Oc, error) {
	var one Oc
	//var parentO0rgan sys.Organ
	values := []interface{}{
		&one.UUID,
		&one.CustSize,
		&one.AutoOc,
		&one.ManualOc,
		&one.Flag,
		&one.ParamType,
		&one.StartTime,
		&one.CreateTime,
		&one.CreateUser,
		&one.UpdateTime,
		&one.UpdateUser,

		&one.Organ.OrganCode,
		&one.Organ.OrganName,
		&one.Organ.ParentOrgan,

		&one.Product.ProductCode,
		&one.Product.ProductName,
	}
	//one.Organ.ParentOrgan = &parentO0rgan
	err := util.OracleScan(rows, values)
	if nil != err {
		return nil, err
	}
	return &one, nil
}

// 分页操作
// by author Yeqc
// by time 2016-12-19 10:21:15
func (this Oc) List(param ...map[string]interface{}) (*util.PageData, error) {
	pageData, rows, err := modelsUtil.List(ocTables, ocCols, ocColsSort, param...)
	defer rows.Close()
	if nil != err {
		er := fmt.Errorf("分页查询运营成本率信息出错")
		zlog.Error(er.Error(), err)
		return nil, er
	}

	var rst []*Oc
	for rows.Next() {
		one, err := this.scanOc(rows)
		if nil != err {
			er := fmt.Errorf("分页查询运营成本率信息row.Scan()出错")
			zlog.Error(er.Error(), err)
			return nil, er
		}
		rst = append(rst, one)
	}
	pageData.Rows = rst
	return pageData, nil
}

func (o Oc) Find(param ...map[string]interface{}) ([]*Oc, error) {
	rows, err := modelsUtil.FindRows(ocTables, ocCols, ocColsSort, param...)
	if nil != err {
		var msg string = "查询运营费用率失败"
		zlog.Error(msg, err)
		return nil, err
	}
	defer rows.Close()
	var rsts []*Oc
	for rows.Next() {
		one, err := o.scanOc(rows)
		if nil != err {
			var msg string = "查询运营费用率rows.Scan()错误"
			zlog.Error(msg, err)
			return nil, err
		}
		rsts = append(rsts, one)
	}
	return rsts, nil
}

// 新增运营成本率信息
// by author Yeqc
// by time 2016-12-19 10:21:15
func (this Oc) Add() error {
	paramMap := map[string]interface{}{
		"organ":       this.Organ.OrganCode,
		"product":     this.Product.ProductCode,
		"cust_size":   this.CustSize,
		"auto_oc":     this.AutoOc,
		"manual_oc":   this.ManualOc,
		"flag":        this.Flag,
		"param_type":  this.ParamType,
		"start_time":  this.StartTime,
		"create_time": util.GetCurrentTime(),
		"create_user": this.CreateUser,
		"update_time": util.GetCurrentTime(),
		"update_user": this.UpdateUser,
	}
	err := util.OracleAdd("RPM_PAR_OC", paramMap)
	if nil != err {
		er := fmt.Errorf("新增运营成本率信息出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

// 更新运营成本率信息
// by author Yeqc
// by time 2016-12-19 10:21:15
func (this Oc) Update() error {
	paramMap := map[string]interface{}{
		"organ":       this.Organ.OrganCode,
		"product":     this.Product.ProductCode,
		"cust_size":   this.CustSize,
		"auto_oc":     this.AutoOc,
		"manual_oc":   this.ManualOc,
		"flag":        this.Flag,
		"param_type":  this.ParamType,
		"start_time":  this.StartTime,
		"update_time": util.GetCurrentTime(),
		"update_user": this.UpdateUser,
	}
	whereParam := map[string]interface{}{
		"uuid": this.UUID,
	}
	err := util.OracleUpdate("RPM_PAR_OC", paramMap, whereParam)
	if nil != err {
		er := fmt.Errorf("更新运营成本率信息出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

// 删除运营成本率信息
// by author Yeqc
// by time 2016-12-19 10:21:15
func (this Oc) Delete() error {
	whereParam := map[string]interface{}{
		"uuid": this.UUID,
	}
	err := util.OracleDelete("RPM_PAR_OC", whereParam)
	if nil != err {
		er := fmt.Errorf("删除运营成本率信息出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}
