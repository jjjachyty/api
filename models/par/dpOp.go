package par

import (
	"database/sql"
	"fmt"
	"pccqcpa.com.cn/app/rpm/api/models/dim"
	"pccqcpa.com.cn/app/rpm/api/models/sys"
	modelsUtil "pccqcpa.com.cn/app/rpm/api/models/util"
	"pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/components/zlog"
	"time"
)

// 结构体对应数据库表【RPM_PAR_DP_OP】存款操作风险率
// by author Yeqc
// by time 2016-12-28 10:37:13
/**
 * @apiDefine DpOp
 * @apiSuccess {string}     	UUID            主键默认值sys_guid()
 * @apiSuccess {Organ}   		Branch    		机构
 * @apiSuccess {Product}   		Product      	产品
 * @apiSuccess {float64}   		OpRate    		操作风险率
 * @apiSuccess {string}   		Flag      		生效标志
 * @apiSuccess {time.Time}   	StartTime      	生效日期
 * @apiSuccess {time.Time}		CreateTime      创建时间
 * @apiSuccess {string}   		CreateUser      创建人
 * @apiSuccess {time.Time}		UpdateTime      更新时间
 * @apiSuccess {string}   		UpdateUser      更新人
 */
type DpOp struct {
	UUID       string      // 主键
	Branch     sys.Organ   // 机构
	Product    dim.Product // 产品
	OpRate     float64     // 操作风险率
	Flag       string      // 生效标志
	StartTime  time.Time   // 生效日期
	CreateTime time.Time   // 创建时间
	CreateUser string      // 创建人
	UpdateTime time.Time   // 更新时间
	UpdateUser string      // 更新人
}

// 结构体scan
// by author Yeqc
// by time 2016-12-28 10:37:13
func (this DpOp) scan(rows *sql.Rows) (*DpOp, error) {
	var one = new(DpOp)
	values := []interface{}{
		&one.UUID,
		// &one.Branch,
		// &one.ProductCd,
		// &one.Currency,
		&one.OpRate,
		&one.Flag,
		&one.StartTime,
		&one.CreateTime,
		&one.CreateUser,
		&one.UpdateTime,
		&one.UpdateUser,

		&one.Product.ProductCode,
		&one.Product.ProductName,
		&one.Product.ProductType,
		&one.Product.ProductTypeDesc,
		&one.Product.ProductLevel,

		&one.Branch.OrganCode,
		&one.Branch.OrganName,
	}
	err := util.OracleScan(rows, values)
	return one, err
}

// 分页操作
// by author Yeqc
// by time 2016-12-28 10:37:13
func (this DpOp) List(param ...map[string]interface{}) (*util.PageData, error) {
	pageData, rows, err := modelsUtil.List(dpOpTabales, dpOpCols, dpOpColsSort, param...)
	defer rows.Close()
	if nil != err {
		er := fmt.Errorf("分页查询存款操作风险率信息出错")
		zlog.Error(er.Error(), err)
		return nil, er
	}

	var rst []*DpOp
	for rows.Next() {
		one, err := this.scan(rows)
		if nil != err {
			er := fmt.Errorf("分页查询存款操作风险率信息row.Scan()出错")
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
// by time 2016-12-28 10:37:13
func (this DpOp) Find(param ...map[string]interface{}) ([]*DpOp, error) {
	rows, err := modelsUtil.FindRows(dpOpTabales, dpOpCols, dpOpColsSort, param...)
	defer rows.Close()
	if nil != err {
		er := fmt.Errorf("多参数查询存款操作风险率信息出错")
		zlog.Error(er.Error(), err)
		return nil, er
	}
	var rst []*DpOp
	for rows.Next() {
		one, err := this.scan(rows)
		if nil != err {
			er := fmt.Errorf("多参数查询存款操作风险率信息row.Scan()出错")
			zlog.Error(er.Error(), err)
			return nil, er
		}
		rst = append(rst, one)
	}
	return rst, nil
}

// 新增存款操作风险率信息
// by author Yeqc
// by time 2016-12-28 10:37:13
func (this DpOp) Add() error {
	paramMap := map[string]interface{}{
		"organ":   this.Branch.OrganCode,
		"product": this.Product.ProductCode,
		// "currency":    this.Currency,
		"op_rate":     this.OpRate,
		"flag":        this.Flag,
		"start_time":  this.StartTime,
		"create_time": util.GetCurrentTime(),
		"create_user": this.CreateUser,
		"update_time": util.GetCurrentTime(),
		"update_user": this.UpdateUser,
	}
	err := util.OracleAdd("RPM_PAR_DP_OP", paramMap)
	if nil != err {
		er := fmt.Errorf("新增存款操作风险率信息出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

// 更新存款操作风险率信息
// by author Yeqc
// by time 2016-12-28 10:37:13
func (this DpOp) Update() error {
	paramMap := map[string]interface{}{
		"organ":   this.Branch.OrganCode,
		"product": this.Product.ProductCode,
		// "currency":    this.Currency,
		"op_rate":     this.OpRate,
		"flag":        this.Flag,
		"start_time":  this.StartTime,
		"update_time": util.GetCurrentTime(),
		"update_user": this.UpdateUser,
	}
	whereParam := map[string]interface{}{
		"uuid": this.UUID,
	}
	err := util.OracleUpdate("RPM_PAR_DP_OP", paramMap, whereParam)
	if nil != err {
		er := fmt.Errorf("更新存款操作风险率信息出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

// 删除存款操作风险率信息
// by author Yeqc
// by time 2016-12-28 10:37:13
func (this DpOp) Delete() error {
	whereParam := map[string]interface{}{
		"uuid": this.UUID,
	}
	err := util.OracleDelete("RPM_PAR_DP_OP", whereParam)
	if nil != err {
		er := fmt.Errorf("删除存款操作风险率信息出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

var dpOpTabales string = `
	     RPM_PAR_DP_OP T
	LEFT JOIN RPM_DIM_PRODUCT PRODUCT 
	ON (PRODUCT.PRODUCT_CODE=T.PRODUCT)
	LEFT JOIN SYS_SEC_ORGAN BRANCH
	ON (BRANCH.ORGAN_CODE = T.ORGAN)`

var dpOpCols map[string]string = map[string]string{
	"T.UUID": "' '",
	// "T.BRANCH": "' '",
	// "T.PRODUCT_CD":  "' '",
	// "T.CURRENCY":    "' '",
	"T.OP_RATE":     "0",
	"T.FLAG":        "' '",
	"T.START_TIME":  "sysdate",
	"T.CREATE_TIME": "sysdate",
	"T.CREATE_USER": "' '",
	"T.UPDATE_TIME": "sysdate",
	"T.UPDATE_USER": "' '",

	"PRODUCT.PRODUCT_CODE":      "' '",
	"PRODUCT.PRODUCT_NAME":      "' '",
	"PRODUCT.PRODUCT_TYPE":      "' '",
	"PRODUCT.PRODUCT_TYPE_DESC": "' '",
	"PRODUCT.PRODUCT_LEVEL":     "' '",

	"BRANCH.ORGAN_CODE": "' '",
	"BRANCH.ORGAN_NAME": "' '",
}

var dpOpColsSort = []string{
	"T.UUID",
	// "T.BRANCH",
	// "T.PRODUCT_CD",
	// "T.CURRENCY",
	"T.OP_RATE",
	"T.FLAG",
	"T.START_TIME",
	"T.CREATE_TIME",
	"T.CREATE_USER",
	"T.UPDATE_TIME",
	"T.UPDATE_USER",

	"PRODUCT.PRODUCT_CODE",
	"PRODUCT.PRODUCT_NAME",
	"PRODUCT.PRODUCT_TYPE",
	"PRODUCT.PRODUCT_TYPE_DESC",
	"PRODUCT.PRODUCT_LEVEL",

	"BRANCH.ORGAN_CODE",
	"BRANCH.ORGAN_NAME",
}
