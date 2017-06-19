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

// 结构体对应数据库表【RPM_PAR_DP_DIFFER】存款差异化定价参数表
// by author Jason
// by time 2016-11-30 09:31:42
/**
 * @apiDefine DpDiffer
 * @apiSuccess {string}     UUID            主键默认值sys_guid()
 * @apiSuccess {Organ}   	Branch    		机构
 * @apiSuccess {Product}   	Product      	产品
 * @apiSuccess {string}   	DimType   		维度编码类型
 * @apiSuccess {string}   	Dim    			维度编码
 * @apiSuccess {string}   	Percent      	上浮比例
 * @apiSuccess {string}   	Flag      		生效标志
 * @apiSuccess {string}   	FloatType      	浮动类型
 * @apiSuccess {time.Time}	CreateTime      创建时间
 * @apiSuccess {string}   	CreateUser      创建人
 * @apiSuccess {time.Time}	UpdateTime      更新时间
 * @apiSuccess {string}   	UpdateUser      更新人
 */
type DpDiffer struct {
	UUID       string      // 主键
	Branch     sys.Organ   // 机构
	Product    dim.Product // 产品
	DimType    string      // 维度编码类型
	Dim        string      // 维度编码
	Percent    float64     // 上浮比例
	Flag       string      // 生效标志
	CreateTime time.Time   // 创建时间
	UpdateTime time.Time   // 更新时间
	CreateUser string      // 创建人
	UpdateUser string      // 更新人
}

// 结构体scan
// by author Jason
// by time 2016-11-30 09:31:42
func (this DpDiffer) scan(rows *sql.Rows) (*DpDiffer, error) {
	var one = new(DpDiffer)
	values := []interface{}{
		&one.UUID,
		&one.Branch.OrganCode,
		&one.Product.ProductCode,
		&one.DimType,
		&one.Dim,
		&one.Percent,
		&one.Flag,
		&one.CreateTime,
		&one.UpdateTime,
		&one.CreateUser,
		&one.UpdateUser,
		&one.Branch.OrganName,
		&one.Product.ProductName,
	}
	err := util.OracleScan(rows, values)
	return one, err
}

// 分页操作
// by author Jason
// by time 2016-11-30 09:31:42
func (this DpDiffer) List(param ...map[string]interface{}) (*util.PageData, error) {
	pageData, rows, err := modelsUtil.List(dpDifferTabales, dpDifferCols, dpDifferColsSort, param...)
	defer rows.Close()
	if nil != err {
		er := fmt.Errorf("分页查询存款差异化定价参数表信息出错")
		zlog.Error(er.Error(), err)
		return nil, er
	}

	var rst []*DpDiffer
	for rows.Next() {
		one, err := this.scan(rows)
		if nil != err {
			er := fmt.Errorf("分页查询存款差异化定价参数表信息row.Scan()出错")
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
// by time 2016-11-30 09:31:42
func (this DpDiffer) Find(param ...map[string]interface{}) ([]*DpDiffer, error) {
	rows, err := modelsUtil.FindRows(dpDifferTabales, dpDifferCols, dpDifferColsSort, param...)
	defer rows.Close()
	if nil != err {
		er := fmt.Errorf("多参数查询存款差异化定价参数表信息出错")
		zlog.Error(er.Error(), err)
		return nil, er
	}
	var rst []*DpDiffer
	for rows.Next() {
		one, err := this.scan(rows)
		if nil != err {
			er := fmt.Errorf("多参数查询存款差异化定价参数表信息row.Scan()出错")
			zlog.Error(er.Error(), err)
			return nil, er
		}
		rst = append(rst, one)
	}
	return rst, nil
}

// 新增存款差异化定价参数表信息
// by author Jason
// by time 2016-11-30 09:31:42
func (this DpDiffer) Add() error {
	paramMap := map[string]interface{}{
		"branch":      this.Branch.OrganCode,
		"product":     this.Product.ProductCode,
		"dim_type":    this.DimType,
		"dim":         this.Dim,
		"percent":     this.Percent,
		"flag":        this.Flag,
		"create_time": util.GetCurrentTime(),
		"update_time": util.GetCurrentTime(),
		"create_user": this.CreateUser,
		"update_user": this.UpdateUser,
	}
	err := util.OracleAdd("RPM_PAR_DP_DIFFER", paramMap)
	if nil != err {
		er := fmt.Errorf("新增存款差异化定价参数表信息出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

// 更新存款差异化定价参数表信息
// by author Jason
// by time 2016-11-30 09:31:42
func (this DpDiffer) Update() error {
	paramMap := map[string]interface{}{
		"branch":      this.Branch.OrganCode,
		"product":     this.Product.ProductCode,
		"dim_type":    this.DimType,
		"dim":         this.Dim,
		"percent":     this.Percent,
		"flag":        this.Flag,
		"update_time": util.GetCurrentTime(),
		"update_user": this.UpdateUser,
	}
	whereParam := map[string]interface{}{
		"uuid": this.UUID,
	}
	err := util.OracleUpdate("RPM_PAR_DP_DIFFER", paramMap, whereParam)
	if nil != err {
		er := fmt.Errorf("更新存款差异化定价参数表信息出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

// 删除存款差异化定价参数表信息
// by author Jason
// by time 2016-11-30 09:31:42
func (this DpDiffer) Delete() error {
	whereParam := map[string]interface{}{
		"uuid": this.UUID,
	}
	err := util.OracleDelete("RPM_PAR_DP_DIFFER", whereParam)
	if nil != err {
		er := fmt.Errorf("删除存款差异化定价参数表信息出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

var dpDifferTabales string = `RPM_PAR_DP_DIFFER T 
						 LEFT JOIN SYS_SEC_ORGAN BRANCH
						   ON (T.BRANCH = BRANCH.ORGAN_CODE)
						 LEFT JOIN RPM_DIM_PRODUCT PRODUCT 
						   ON (T.PRODUCT = PRODUCT.PRODUCT_CODE AND PRODUCT.FLAG = 1)`

var dpDifferCols map[string]string = map[string]string{
	"T.UUID":        "' '",
	"T.BRANCH":      "' '",
	"T.PRODUCT":     "' '",
	"T.DIM_TYPE":    "' '",
	"T.DIM":         "' '",
	"T.PERCENT":     "0",
	"T.FLAG":        "' '",
	"T.CREATE_TIME": "SYSDATE",
	"T.UPDATE_TIME": "SYSDATE",
	"T.CREATE_USER": "' '",
	"T.UPDATE_USER": "' '",

	"BRANCH.ORGAN_NAME":    "' '",
	"PRODUCT.PRODUCT_NAME": "' '",
}

var dpDifferColsSort = []string{
	"T.UUID",
	"T.BRANCH",
	"T.PRODUCT",
	"T.DIM_TYPE",
	"T.DIM",
	"T.PERCENT",
	"T.FLAG",
	"T.CREATE_TIME",
	"T.UPDATE_TIME",
	"T.CREATE_USER",
	"T.UPDATE_USER",

	"BRANCH.ORGAN_NAME",
	"PRODUCT.PRODUCT_NAME",
}
