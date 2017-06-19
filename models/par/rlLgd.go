package par

import (
	"database/sql"
	"fmt"

	"pccqcpa.com.cn/app/rpm/api/models/dim"
	modelsUtil "pccqcpa.com.cn/app/rpm/api/models/util"
	"pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/components/zlog"
)

// 结构体对应数据库表【RPM_PAR_RL_LGD】零售违约损失率-零售贷款模块
// by author Yeqc
// by time 2016-12-30 10:06:36
/**
 * @apiDefine RlLgd
 * @apiSuccess {string}     UUID            主键默认值sys_guid()
 * @apiSuccess {Product}   	ProductCode    	产品
 * @apiSuccess {string}   	Lgd      		违约损失率
 */
type RlLgd struct {
	UUID        string      // 默认值sys_guid()
	ProductCode dim.Product // 产品
	Lgd         float64     // 违约损失率
}

// 结构体scan
// by author Yeqc
// by time 2016-12-30 10:06:36
func (this RlLgd) scan(rows *sql.Rows) (*RlLgd, error) {
	var one = new(RlLgd)
	values := []interface{}{
		&one.UUID,
		// &one.ProductCode,
		&one.Lgd,

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
// by time 2016-12-30 10:06:36
func (this RlLgd) List(param ...map[string]interface{}) (*util.PageData, error) {
	pageData, rows, err := modelsUtil.List(rlLgdTabales, rlLgdCols, rlLgdColsSort, param...)
	defer rows.Close()
	if nil != err {
		er := fmt.Errorf("分页查询零售违约损失率-零售贷款模块信息出错")
		zlog.Error(er.Error(), err)
		return nil, er
	}

	var rst []*RlLgd
	for rows.Next() {
		one, err := this.scan(rows)
		if nil != err {
			er := fmt.Errorf("分页查询零售违约损失率-零售贷款模块信息row.Scan()出错")
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
// by time 2016-12-30 10:06:36
func (this RlLgd) Find(param ...map[string]interface{}) ([]*RlLgd, error) {
	rows, err := modelsUtil.FindRows(rlLgdTabales, rlLgdCols, rlLgdColsSort, param...)
	defer rows.Close()
	if nil != err {
		er := fmt.Errorf("多参数查询零售违约损失率-零售贷款模块信息出错")
		zlog.Error(er.Error(), err)
		return nil, er
	}
	var rst []*RlLgd
	for rows.Next() {
		one, err := this.scan(rows)
		if nil != err {
			er := fmt.Errorf("多参数查询零售违约损失率-零售贷款模块信息row.Scan()出错")
			zlog.Error(er.Error(), err)
			return nil, er
		}
		rst = append(rst, one)
	}
	return rst, nil
}

// 新增零售违约损失率-零售贷款模块信息
// by author Yeqc
// by time 2016-12-30 10:06:36
func (this RlLgd) Add() error {
	paramMap := map[string]interface{}{
		"product_code": this.ProductCode.ProductCode,
		"lgd":          this.Lgd,
	}
	err := util.OracleAdd("RPM_PAR_RL_LGD", paramMap)
	if nil != err {
		er := fmt.Errorf("新增零售违约损失率-零售贷款模块信息出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

// 更新零售违约损失率-零售贷款模块信息
// by author Yeqc
// by time 2016-12-30 10:06:36
func (this RlLgd) Update() error {
	paramMap := map[string]interface{}{
		"product_code": this.ProductCode.ProductCode,
		"lgd":          this.Lgd,
	}
	whereParam := map[string]interface{}{
		"uuid": this.UUID,
	}
	err := util.OracleUpdate("RPM_PAR_RL_LGD", paramMap, whereParam)
	if nil != err {
		er := fmt.Errorf("更新零售违约损失率-零售贷款模块信息出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

// 删除零售违约损失率-零售贷款模块信息
// by author Yeqc
// by time 2016-12-30 10:06:36
func (this RlLgd) Delete() error {
	whereParam := map[string]interface{}{
		"uuid": this.UUID,
	}
	err := util.OracleDelete("RPM_PAR_RL_LGD", whereParam)
	if nil != err {
		er := fmt.Errorf("删除零售违约损失率-零售贷款模块信息出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

var rlLgdTabales string = `
	RPM_PAR_RL_LGD T
	LEFT JOIN RPM_DIM_PRODUCT PRODUCTCODE
	ON (T.PRODUCT_CODE=PRODUCTCODE.PRODUCT_CODE)
	`

var rlLgdCols map[string]string = map[string]string{
	"T.UUID": "' '",
	// "T.PRODUCT_CODE": "' '",
	"T.LGD": "0",

	"PRODUCTCODE.PRODUCT_CODE":      "' '",
	"PRODUCTCODE.PRODUCT_NAME":      "' '",
	"PRODUCTCODE.PRODUCT_TYPE":      "' '",
	"PRODUCTCODE.PRODUCT_TYPE_DESC": "' '",
	"PRODUCTCODE.PRODUCT_LEVEL":     "' '",
}

var rlLgdColsSort = []string{
	"T.UUID",
	// "T.PRODUCT_CODE",
	"T.LGD",

	"PRODUCTCODE.PRODUCT_CODE",
	"PRODUCTCODE.PRODUCT_NAME",
	"PRODUCTCODE.PRODUCT_TYPE",
	"PRODUCTCODE.PRODUCT_TYPE_DESC",
	"PRODUCTCODE.PRODUCT_LEVEL",
}
