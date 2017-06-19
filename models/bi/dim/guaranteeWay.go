package dim

import (
	"database/sql"
	"fmt"

	modelsUtil "pccqcpa.com.cn/app/rpm/api/models/util"
	"pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/components/zlog"
	"platform/dbobj"
)

// 结构体对应数据库表【RPM_BI_DIM_GUARANTEE_WAY】业务分析担保方式维度表
// by author Jason
// by time 2016-10-31 15:34:35
type DimGuaranteeWay struct {
	UUID string // 主键
	Code string // 担保方式代码
	Name string // 担保方式名称
}

// 结构体scan
// by author Jason
// by time 2016-10-31 15:34:35
func (this DimGuaranteeWay) scan(rows *sql.Rows) (*DimGuaranteeWay, error) {
	var one = new(DimGuaranteeWay)
	values := []interface{}{
		&one.UUID,
		&one.Code,
		&one.Name,
	}
	err := util.OracleScan(rows, values)
	return one, err
}

// 分页操作
// by author Jason
// by time 2016-10-31 15:34:35
func (this DimGuaranteeWay) List(param ...map[string]interface{}) (*util.PageData, error) {
	pageData, rows, err := modelsUtil.List(dimGuaranteeWayTabales, dimGuaranteeWayCols, dimGuaranteeWayColsSort, param...)
	defer rows.Close()
	if nil != err {
		er := fmt.Errorf("分页查询业务分析担保方式维度表信息出错")
		zlog.Error(er.Error(), err)
		return nil, er
	}

	var rst []*DimGuaranteeWay
	for rows.Next() {
		one, err := this.scan(rows)
		if nil != err {
			er := fmt.Errorf("分页查询业务分析担保方式维度表信息row.Scan()出错")
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
// by time 2016-10-31 15:34:35
func (this DimGuaranteeWay) Find(param ...map[string]interface{}) ([]*DimGuaranteeWay, error) {
	rows, err := modelsUtil.FindRows(dimGuaranteeWayTabales, dimGuaranteeWayCols, dimGuaranteeWayColsSort, param...)
	defer rows.Close()
	if nil != err {
		er := fmt.Errorf("多参数查询业务分析担保方式维度表信息出错")
		zlog.Error(er.Error(), err)
		return nil, er
	}
	defer rows.Close()
	var rst []*DimGuaranteeWay
	for rows.Next() {
		one, err := this.scan(rows)
		if nil != err {
			er := fmt.Errorf("多参数查询业务分析担保方式维度表信息row.Scan()出错")
			zlog.Error(er.Error(), err)
			return nil, er
		}
		rst = append(rst, one)
	}
	return rst, nil
}

// 新增业务分析担保方式维度表信息
// by author Jason
// by time 2016-10-31 15:34:35
func (this DimGuaranteeWay) Add() error {
	paramMap := map[string]interface{}{
		"code": this.Code,
		"name": this.Name,
	}
	err := util.OracleAdd("RPM_BI_DIM_GUARANTEE_WAY", paramMap)
	if nil != err {
		er := fmt.Errorf("新增业务分析担保方式维度表信息出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

// 新增贷款业务分析基准利率降息维度表信息
// by author Jason
// by time 2016-10-31 15:34:54
func (this DimGuaranteeWay) BatchAdd(models []DimGuaranteeWay) (sql.Result, error) {
	sqlTx, err := dbobj.Default.Begin()
	var tableName = "RPM_BI_DIM_GUARANTEE_WAY"
	var deleteSQL = "DELETE FROM " + tableName
	var msg = "[业务方案分析行业维度表信息-" + tableName + "]"
	if nil != err {
		zlog.Error(msg+"导入,事物开启失败", nil)
		return nil, err
	}
	//先清除表数据
	stmt, err := sqlTx.Prepare(deleteSQL)
	if nil != err {
		zlog.Error(msg+"导入,获取Stmt失败", nil)
		return nil, err
	}
	result, err := stmt.Exec()
	zlog.Infof(msg+"导入,DELETE-SQL:%s", nil, deleteSQL)
	if nil != err {
		zlog.Error(msg+"导入,DELETE出错", err)
		sqlTx.Rollback()
		return result, err
	}
	//关闭连接
	stmt.Close()

	for i, model := range models {
		zlog.Debugf(msg+"导入第%d条", nil, i+1)

		paramMap := map[string]interface{}{
			"code": model.Code,
			"name": model.Name,
		}
		result, err = util.OracleBatchStmtAdd(tableName, paramMap, sqlTx)
		if nil != err {
			zlog.Errorf(msg+"导入第%d行出错", err, i+1)
			err = stmt.Close()
			if nil != err {
				zlog.Error(msg+"链接关闭出错", nil)
				sqlTx.Rollback()
				return nil, err
			}
			sqlTx.Rollback()
			return result, fmt.Errorf(msg+"导入第%d行出错,%s", i+1, err)
		}
	}

	//关闭连接
	err = stmt.Close()
	if nil != err {
		zlog.Error(msg+"链接关闭出错", nil)

		sqlTx.Rollback()
		return nil, err
	}
	zlog.Info(msg+"导入成功", nil)
	sqlTx.Commit()

	return nil, nil
}

// 更新业务分析担保方式维度表信息
// by author Jason
// by time 2016-10-31 15:34:35
func (this DimGuaranteeWay) Update() error {
	paramMap := map[string]interface{}{
		"code": this.Code,
		"name": this.Name,
	}
	whereParam := map[string]interface{}{
		"uuid": this.UUID,
	}
	err := util.OracleUpdate("RPM_BI_DIM_GUARANTEE_WAY", paramMap, whereParam)
	if nil != err {
		er := fmt.Errorf("更新业务分析担保方式维度表信息出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

// 删除业务分析担保方式维度表信息
// by author Jason
// by time 2016-10-31 15:34:35
func (this DimGuaranteeWay) Delete() error {
	whereParam := map[string]interface{}{
		"uuid": this.UUID,
	}
	err := util.OracleDelete("RPM_BI_DIM_GUARANTEE_WAY", whereParam)
	if nil != err {
		er := fmt.Errorf("删除业务分析担保方式维度表信息出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

var dimGuaranteeWayTabales string = `RPM_BI_DIM_GUARANTEE_WAY`

var dimGuaranteeWayCols map[string]string = map[string]string{
	"UUID": "' '",
	"CODE": "' '",
	"NAME": "' '",
}

var dimGuaranteeWayColsSort = []string{
	"UUID",
	"CODE",
	"NAME",
}
