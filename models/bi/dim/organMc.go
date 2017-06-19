package dim

import (
	"database/sql"
	"fmt"

	modelsUtil "pccqcpa.com.cn/app/rpm/api/models/util"
	"pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/components/zlog"
	"platform/dbobj"
)

// 结构体对应数据库表【RPM_BI_DIM_ORGAN_MC】贷款业务分析方案管快机构维度表
// by author Jason
// by time 2016-10-31 15:37:34
type DimOrganMc struct {
	UUID        string // 注释
	OrganCodeMc string // 管快机构编码
	OrganNameMc string // 管快机构名称
	OrganCodePc string // 定价机构编码
	OrganNamePc string // 定价机构名称
	AreaCode    string // 地区编码
	AreaName    string // 地区名称
}

// 结构体scan
// by author Jason
// by time 2016-10-31 15:37:34
func (this DimOrganMc) scan(rows *sql.Rows) (*DimOrganMc, error) {
	var one = new(DimOrganMc)
	values := []interface{}{
		&one.UUID,
		&one.OrganCodeMc,
		&one.OrganNameMc,
		&one.OrganCodePc,
		&one.OrganNamePc,
		&one.AreaCode,
		&one.AreaName,
	}
	err := util.OracleScan(rows, values)
	return one, err
}

// 分页操作
// by author Jason
// by time 2016-10-31 15:37:34
func (this DimOrganMc) List(param ...map[string]interface{}) (*util.PageData, error) {
	pageData, rows, err := modelsUtil.List(dimOrganMcTabales, dimOrganMcCols, dimOrganMcColsSort, param...)
	defer rows.Close()
	if nil != err {
		er := fmt.Errorf("分页查询贷款业务分析方案管快机构维度表信息出错")
		zlog.Error(er.Error(), err)
		return nil, er
	}

	var rst []*DimOrganMc
	for rows.Next() {
		one, err := this.scan(rows)
		if nil != err {
			er := fmt.Errorf("分页查询贷款业务分析方案管快机构维度表信息row.Scan()出错")
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
// by time 2016-10-31 15:37:34
func (this DimOrganMc) Find(param ...map[string]interface{}) ([]*DimOrganMc, error) {
	rows, err := modelsUtil.FindRows(dimOrganMcTabales, dimOrganMcCols, dimOrganMcColsSort, param...)
	defer rows.Close()
	if nil != err {
		er := fmt.Errorf("多参数查询贷款业务分析方案管快机构维度表信息出错")
		zlog.Error(er.Error(), err)
		return nil, er
	}
	defer rows.Close()
	var rst []*DimOrganMc
	for rows.Next() {
		one, err := this.scan(rows)
		if nil != err {
			er := fmt.Errorf("多参数查询贷款业务分析方案管快机构维度表信息row.Scan()出错")
			zlog.Error(er.Error(), err)
			return nil, er
		}
		rst = append(rst, one)
	}
	return rst, nil
}

// 新增贷款业务分析方案管快机构维度表信息
// by author Jason
// by time 2016-10-31 15:37:34
func (this DimOrganMc) Add() error {
	paramMap := map[string]interface{}{
		"organ_code_mc": this.OrganCodeMc,
		"organ_name_mc": this.OrganNameMc,
		"organ_code_pc": this.OrganCodePc,
		"organ_name_pc": this.OrganNamePc,
		"area_code":     this.AreaCode,
		"area_name":     this.AreaName,
	}
	err := util.OracleAdd("RPM_BI_DIM_ORGAN_MC", paramMap)
	if nil != err {
		er := fmt.Errorf("新增贷款业务分析方案管快机构维度表信息出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

// 新增贷款业务分析基准利率降息维度表信息
// by author Jason
// by time 2016-10-31 15:34:54
func (this DimOrganMc) BatchAdd(models []DimOrganMc) (sql.Result, error) {

	sqlTx, err := dbobj.Default.Begin()
	var tableName = "RPM_BI_DIM_ORGAN_MC"
	var deleteSQL = "DELETE FROM " + tableName
	var msg = "[贷款业务分析方案管快机构维度表信息-" + tableName + "]"
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
			"organ_code_mc": model.OrganCodeMc,
			"organ_name_mc": model.OrganNameMc,
			"organ_code_pc": model.OrganCodePc,
			"organ_name_pc": model.OrganNamePc,
			"area_code":     model.AreaCode,
			"area_name":     model.AreaName,
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

// 更新贷款业务分析方案管快机构维度表信息
// by author Jason
// by time 2016-10-31 15:37:34
func (this DimOrganMc) Update() error {
	paramMap := map[string]interface{}{
		"organ_code_mc": this.OrganCodeMc,
		"organ_name_mc": this.OrganNameMc,
		"organ_code_pc": this.OrganCodePc,
		"organ_name_pc": this.OrganNamePc,
		"area_code":     this.AreaCode,
		"area_name":     this.AreaName,
	}
	whereParam := map[string]interface{}{
		"uuid": this.UUID,
	}
	err := util.OracleUpdate("RPM_BI_DIM_ORGAN_MC", paramMap, whereParam)
	if nil != err {
		er := fmt.Errorf("更新贷款业务分析方案管快机构维度表信息出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

// 删除贷款业务分析方案管快机构维度表信息
// by author Jason
// by time 2016-10-31 15:37:34
func (this DimOrganMc) Delete() error {
	whereParam := map[string]interface{}{
		"uuid": this.UUID,
	}
	err := util.OracleDelete("RPM_BI_DIM_ORGAN_MC", whereParam)
	if nil != err {
		er := fmt.Errorf("删除贷款业务分析方案管快机构维度表信息出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

var dimOrganMcTabales string = `RPM_BI_DIM_ORGAN_MC`

var dimOrganMcCols map[string]string = map[string]string{
	"UUID":          "' '",
	"ORGAN_CODE_MC": "' '",
	"ORGAN_NAME_MC": "' '",
	"ORGAN_CODE_PC": "' '",
	"ORGAN_NAME_PC": "' '",
	"AREA_CODE":     "' '",
	"AREA_NAME":     "' '",
}

var dimOrganMcColsSort = []string{
	"UUID",
	"ORGAN_CODE_MC",
	"ORGAN_NAME_MC",
	"ORGAN_CODE_PC",
	"ORGAN_NAME_PC",
	"AREA_CODE",
	"AREA_NAME",
}
