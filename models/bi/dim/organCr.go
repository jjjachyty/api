package dim

import (
	"database/sql"
	"fmt"

	"platform/dbobj"

	modelsUtil "pccqcpa.com.cn/app/rpm/api/models/util"
	"pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/components/zlog"
)

// 结构体对应数据库表【RPM_BI_DIM_ORGAN_CR】贷款业务分析方案信贷机构维度表
// by author Jason
// by time 2016-10-31 15:36:53
type DimOrganCr struct {
	UUID        string // 注释
	OrganCodeCr string // 信贷机构编码
	OrganNameCr string // 信贷机构名称
	OrganCodePc string // 定价机构编码
	OrganNamePc string // 定价机构名称
	AreaCodeOne string // 地区编码
	AreaNameOne string // 地区名称
	AreaCodeTwo string // 地区编码
	AreaNameTwo string // 地区名称
}

// 结构体scan
// by author Jason
// by time 2016-10-31 15:36:53
func (this DimOrganCr) scan(rows *sql.Rows) (*DimOrganCr, error) {
	var one = new(DimOrganCr)
	values := []interface{}{
		&one.UUID,
		&one.OrganCodeCr,
		&one.OrganNameCr,
		&one.OrganCodePc,
		&one.OrganNamePc,
		&one.AreaCodeOne,
		&one.AreaNameOne,
		&one.AreaCodeTwo,
		&one.AreaNameTwo,
	}
	err := util.OracleScan(rows, values)
	return one, err
}

// 分页操作
// by author Jason
// by time 2016-10-31 15:36:53
func (this DimOrganCr) List(param ...map[string]interface{}) (*util.PageData, error) {
	pageData, rows, err := modelsUtil.List(dimOrganCrTabales, dimOrganCrCols, dimOrganCrColsSort, param...)
	defer rows.Close()
	if nil != err {
		er := fmt.Errorf("分页查询贷款业务分析方案信贷机构维度表信息出错")
		zlog.Error(er.Error(), err)
		return nil, er
	}

	var rst []*DimOrganCr
	for rows.Next() {
		one, err := this.scan(rows)
		if nil != err {
			er := fmt.Errorf("分页查询贷款业务分析方案信贷机构维度表信息row.Scan()出错")
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
// by time 2016-10-31 15:36:53
func (this DimOrganCr) Find(param ...map[string]interface{}) ([]*DimOrganCr, error) {
	rows, err := modelsUtil.FindRows(dimOrganCrTabales, dimOrganCrCols, dimOrganCrColsSort, param...)
	defer rows.Close()
	if nil != err {
		er := fmt.Errorf("多参数查询贷款业务分析方案信贷机构维度表信息出错")
		zlog.Error(er.Error(), err)
		return nil, er
	}
	defer rows.Close()
	var rst []*DimOrganCr
	for rows.Next() {
		one, err := this.scan(rows)
		if nil != err {
			er := fmt.Errorf("多参数查询贷款业务分析方案信贷机构维度表信息row.Scan()出错")
			zlog.Error(er.Error(), err)
			return nil, er
		}
		rst = append(rst, one)
	}
	return rst, nil
}

// 新增贷款业务分析方案信贷机构维度表信息
// by author Jason
// by time 2016-10-31 15:36:53
func (this DimOrganCr) Add() error {
	paramMap := map[string]interface{}{
		"organ_code_cr": this.OrganCodeCr,
		"organ_name_cr": this.OrganNameCr,
		"organ_code_pc": this.OrganCodePc,
		"organ_name_pc": this.OrganNamePc,
		"area_code_one": this.AreaCodeOne,
		"area_name_one": this.AreaNameOne,
		"area_code_two": this.AreaCodeTwo,
		"area_name_two": this.AreaNameTwo,
	}
	err := util.OracleAdd("RPM_BI_DIM_ORGAN_CR", paramMap)
	if nil != err {
		er := fmt.Errorf("新增贷款业务分析方案信贷机构维度表信息出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

// 新增贷款业务分析基准利率降息维度表信息
// by author Jason
// by time 2016-10-31 15:34:54
func (this DimOrganCr) BatchAdd(models []DimOrganCr) (sql.Result, error) {
	sqlTx, err := dbobj.Default.Begin()
	var tableName = "RPM_BI_DIM_ORGAN_CR"
	var deleteSQL = "DELETE FROM " + tableName
	var msg = "[贷款业务分析方案信贷机构维度表信息-" + tableName + "]"
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
			"organ_code_cr": model.OrganCodeCr,
			"organ_name_cr": model.OrganNameCr,
			"organ_code_pc": model.OrganCodePc,
			"organ_name_pc": model.OrganNamePc,
			"area_code_one": model.AreaCodeOne,
			"area_name_one": model.AreaNameOne,
			"area_code_two": model.AreaCodeTwo,
			"area_name_two": model.AreaNameTwo,
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

// 更新贷款业务分析方案信贷机构维度表信息
// by author Jason
// by time 2016-10-31 15:36:53
func (this DimOrganCr) Update() error {
	paramMap := map[string]interface{}{
		"organ_code_cr": this.OrganCodeCr,
		"organ_name_cr": this.OrganNameCr,
		"organ_code_pc": this.OrganCodePc,
		"organ_name_pc": this.OrganNamePc,
		"area_code_one": this.AreaCodeOne,
		"area_name_one": this.AreaNameOne,
		"area_code_two": this.AreaCodeTwo,
		"area_name_two": this.AreaNameTwo,
	}
	whereParam := map[string]interface{}{
		"uuid": this.UUID,
	}
	err := util.OracleUpdate("RPM_BI_DIM_ORGAN_CR", paramMap, whereParam)
	if nil != err {
		er := fmt.Errorf("更新贷款业务分析方案信贷机构维度表信息出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

// 删除贷款业务分析方案信贷机构维度表信息
// by author Jason
// by time 2016-10-31 15:36:53
func (this DimOrganCr) Delete() error {
	whereParam := map[string]interface{}{
		"uuid": this.UUID,
	}
	err := util.OracleDelete("RPM_BI_DIM_ORGAN_CR", whereParam)
	if nil != err {
		er := fmt.Errorf("删除贷款业务分析方案信贷机构维度表信息出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

var dimOrganCrTabales string = `RPM_BI_DIM_ORGAN_CR`

var dimOrganCrCols map[string]string = map[string]string{
	"UUID":          "' '",
	"ORGAN_CODE_CR": "' '",
	"ORGAN_NAME_CR": "' '",
	"ORGAN_CODE_PC": "' '",
	"ORGAN_NAME_PC": "' '",
	"AREA_CODE_ONE": "' '",
	"AREA_NAME_ONE": "' '",
	"AREA_CODE_TWO": "' '",
	"AREA_NAME_TWO": "' '",
}

var dimOrganCrColsSort = []string{
	"UUID",
	"ORGAN_CODE_CR",
	"ORGAN_NAME_CR",
	"ORGAN_CODE_PC",
	"ORGAN_NAME_PC",
	"AREA_CODE_ONE",
	"AREA_NAME_ONE",
	"AREA_CODE_TWO",
	"AREA_NAME_TWO",
}
