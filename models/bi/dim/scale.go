package dim

import (
	"database/sql"
	"fmt"

	"platform/dbobj"

	modelsUtil "pccqcpa.com.cn/app/rpm/api/models/util"
	"pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/components/zlog"
)

// 结构体对应数据库表【RPM_BI_DIM_SCALE】贷款业务分析方案规模维度
// by author Jason
// by time 2016-10-31 15:39:12
type DimScale struct {
	UUID string // 主键
	Code string // 规模编码
	Name string // 规模名称
	Type string // 规模类型(国标/行标)
	Sort int
}

// 结构体scan
// by author Jason
// by time 2016-10-31 15:39:12
func (this DimScale) scan(rows *sql.Rows) (*DimScale, error) {
	var one = new(DimScale)
	values := []interface{}{
		&one.UUID,
		&one.Code,
		&one.Name,
		&one.Type,
		&one.Sort,
	}
	err := util.OracleScan(rows, values)
	return one, err
}

// 分页操作
// by author Jason
// by time 2016-10-31 15:39:12
func (this DimScale) List(param ...map[string]interface{}) (*util.PageData, error) {
	pageData, rows, err := modelsUtil.List(dimScaleTabales, dimScaleCols, dimScaleColsSort, param...)
	defer rows.Close()
	if nil != err {
		er := fmt.Errorf("分页查询贷款业务分析方案规模维度信息出错")
		zlog.Error(er.Error(), err)
		return nil, er
	}

	var rst []*DimScale
	for rows.Next() {
		one, err := this.scan(rows)
		if nil != err {
			er := fmt.Errorf("分页查询贷款业务分析方案规模维度信息row.Scan()出错")
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
// by time 2016-10-31 15:39:12
func (this DimScale) Find(param ...map[string]interface{}) ([]*DimScale, error) {
	rows, err := modelsUtil.FindRows(dimScaleTabales, dimScaleCols, dimScaleColsSort, param...)
	defer rows.Close()
	if nil != err {
		er := fmt.Errorf("多参数查询贷款业务分析方案规模维度信息出错")
		zlog.Error(er.Error(), err)
		return nil, er
	}
	defer rows.Close()
	var rst []*DimScale
	for rows.Next() {
		one, err := this.scan(rows)
		if nil != err {
			er := fmt.Errorf("多参数查询贷款业务分析方案规模维度信息row.Scan()出错")
			zlog.Error(er.Error(), err)
			return nil, er
		}
		rst = append(rst, one)
	}
	return rst, nil
}

// 新增贷款业务分析方案规模维度信息
// by author Jason
// by time 2016-10-31 15:39:12
func (this DimScale) Add() error {
	paramMap := map[string]interface{}{
		"code": this.Code,
		"name": this.Name,
		"type": this.Type,
		"sort": this.Sort,
	}
	err := util.OracleAdd("RPM_BI_DIM_SCALE", paramMap)
	if nil != err {
		er := fmt.Errorf("新增贷款业务分析方案规模维度信息出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

// 新增贷款业务分析基准利率降息维度表信息
// by author Jason
// by time 2016-10-31 15:34:54
func (this DimScale) BatchAdd(models []DimScale) (sql.Result, error) {
	sqlTx, err := dbobj.Default.Begin()
	var tableName = "RPM_BI_DIM_SCALE"
	var deleteSQL = "DELETE FROM " + tableName
	var msg = "[贷款业务分析方案规模维度信息-" + tableName + "]"
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
			"type": model.Type,
			"sort": model.Sort,
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

// 更新贷款业务分析方案规模维度信息
// by author Jason
// by time 2016-10-31 15:39:12
func (this DimScale) Update() error {
	paramMap := map[string]interface{}{
		"code": this.Code,
		"name": this.Name,
		"type": this.Type,
		"sort": this.Sort,
	}
	whereParam := map[string]interface{}{
		"uuid": this.UUID,
	}
	err := util.OracleUpdate("RPM_BI_DIM_SCALE", paramMap, whereParam)
	if nil != err {
		er := fmt.Errorf("更新贷款业务分析方案规模维度信息出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

// 删除贷款业务分析方案规模维度信息
// by author Jason
// by time 2016-10-31 15:39:12
func (this DimScale) Delete() error {
	whereParam := map[string]interface{}{
		"uuid": this.UUID,
	}
	err := util.OracleDelete("RPM_BI_DIM_SCALE", whereParam)
	if nil != err {
		er := fmt.Errorf("删除贷款业务分析方案规模维度信息出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

var dimScaleTabales string = `RPM_BI_DIM_SCALE`

var dimScaleCols map[string]string = map[string]string{
	"UUID": "' '",
	"CODE": "' '",
	"NAME": "' '",
	"TYPE": "' '",
	"SORT": "0",
}

var dimScaleColsSort = []string{
	"UUID",
	"CODE",
	"NAME",
	"TYPE",
	"SORT",
}
