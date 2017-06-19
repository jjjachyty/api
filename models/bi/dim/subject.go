package dim

import (
	"database/sql"
	"fmt"

	modelsUtil "pccqcpa.com.cn/app/rpm/api/models/util"
	"pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/components/zlog"
	"platform/dbobj"
)

// 结构体对应数据库表【RPM_BI_DIM_SUBJECT】贷款业务分析科目维度
// by author Jason
// by time 2016-10-31 15:40:47
type DimSubject struct {
	UUID         string // 主键
	SubjectCode1 string // 一级科目代码
	SubjectName1 string // 一级科目名称
	SubjectCode2 string // 二级科目代码
	SubjectName2 string // 二级科目名称
	SubjectCode3 string // 三级科目代码
	SubjectName3 string // 三级科目名称
	SubjectCode  string // 科目代码
	SubjectName  string // 科目名称
}

// 结构体scan
// by author Jason
// by time 2016-10-31 15:40:47
func (this DimSubject) scan(rows *sql.Rows) (*DimSubject, error) {
	var one = new(DimSubject)
	values := []interface{}{
		&one.UUID,
		&one.SubjectCode1,
		&one.SubjectName1,
		&one.SubjectCode2,
		&one.SubjectName2,
		&one.SubjectCode3,
		&one.SubjectName3,
		&one.SubjectCode,
		&one.SubjectName,
	}
	err := util.OracleScan(rows, values)
	return one, err
}

// 分页操作
// by author Jason
// by time 2016-10-31 15:40:47
func (this DimSubject) List(param ...map[string]interface{}) (*util.PageData, error) {
	pageData, rows, err := modelsUtil.List(dimSubjectTabales, dimSubjectCols, dimSubjectColsSort, param...)
	defer rows.Close()
	if nil != err {
		er := fmt.Errorf("分页查询贷款业务分析科目维度信息出错")
		zlog.Error(er.Error(), err)
		return nil, er
	}

	var rst []*DimSubject
	for rows.Next() {
		one, err := this.scan(rows)
		if nil != err {
			er := fmt.Errorf("分页查询贷款业务分析科目维度信息row.Scan()出错")
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
// by time 2016-10-31 15:40:47
func (this DimSubject) Find(param ...map[string]interface{}) ([]*DimSubject, error) {
	rows, err := modelsUtil.FindRows(dimSubjectTabales, dimSubjectCols, dimSubjectColsSort, param...)
	defer rows.Close()
	if nil != err {
		er := fmt.Errorf("多参数查询贷款业务分析科目维度信息出错")
		zlog.Error(er.Error(), err)
		return nil, er
	}
	defer rows.Close()
	var rst []*DimSubject
	for rows.Next() {
		one, err := this.scan(rows)
		if nil != err {
			er := fmt.Errorf("多参数查询贷款业务分析科目维度信息row.Scan()出错")
			zlog.Error(er.Error(), err)
			return nil, er
		}
		rst = append(rst, one)
	}
	return rst, nil
}

// 新增贷款业务分析科目维度信息
// by author Jason
// by time 2016-10-31 15:40:47
func (this DimSubject) Add() error {
	paramMap := map[string]interface{}{
		"subject_code_1": this.SubjectCode1,
		"subject_name_1": this.SubjectName1,
		"subject_code_2": this.SubjectCode2,
		"subject_name_2": this.SubjectName2,
		"subject_code_3": this.SubjectCode3,
		"subject_name_3": this.SubjectName3,
		"subject_code":   this.SubjectCode,
		"subject_name":   this.SubjectName,
	}
	err := util.OracleAdd("RPM_BI_DIM_SUBJECT", paramMap)
	if nil != err {
		er := fmt.Errorf("新增贷款业务分析科目维度信息出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

// 新增贷款业务分析基准利率降息维度表信息
// by author Jason
// by time 2016-10-31 15:34:54
func (this DimSubject) BatchAdd(models []DimSubject) (sql.Result, error) {

	sqlTx, err := dbobj.Default.Begin()
	var tableName = "RPM_BI_DIM_SUBJECT"
	var deleteSQL = "DELETE FROM " + tableName
	var msg = "[贷款业务分析科目维度信息-" + tableName + "]"
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
			"subject_code_1": model.SubjectCode1,
			"subject_name_1": model.SubjectName1,
			"subject_code_2": model.SubjectCode2,
			"subject_name_2": model.SubjectName2,
			"subject_code_3": model.SubjectCode3,
			"subject_name_3": model.SubjectName3,
			"subject_code":   model.SubjectCode,
			"subject_name":   model.SubjectName,
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

// 更新贷款业务分析科目维度信息
// by author Jason
// by time 2016-10-31 15:40:47
func (this DimSubject) Update() error {
	paramMap := map[string]interface{}{
		"subject_code_1": this.SubjectCode1,
		"subject_name_1": this.SubjectName1,
		"subject_code_2": this.SubjectCode2,
		"subject_name_2": this.SubjectName2,
		"subject_code_3": this.SubjectCode3,
		"subject_name_3": this.SubjectName3,
		"subject_code":   this.SubjectCode,
		"subject_name":   this.SubjectName,
	}
	whereParam := map[string]interface{}{
		"uuid": this.UUID,
	}
	err := util.OracleUpdate("RPM_BI_DIM_SUBJECT", paramMap, whereParam)
	if nil != err {
		er := fmt.Errorf("更新贷款业务分析科目维度信息出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

// 删除贷款业务分析科目维度信息
// by author Jason
// by time 2016-10-31 15:40:47
func (this DimSubject) Delete() error {
	whereParam := map[string]interface{}{
		"uuid": this.UUID,
	}
	err := util.OracleDelete("RPM_BI_DIM_SUBJECT", whereParam)
	if nil != err {
		er := fmt.Errorf("删除贷款业务分析科目维度信息出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

var dimSubjectTabales string = `RPM_BI_DIM_SUBJECT`

var dimSubjectCols map[string]string = map[string]string{
	"UUID":           "' '",
	"SUBJECT_CODE_1": "' '",
	"SUBJECT_NAME_1": "' '",
	"SUBJECT_CODE_2": "' '",
	"SUBJECT_NAME_2": "' '",
	"SUBJECT_CODE_3": "' '",
	"SUBJECT_NAME_3": "' '",
	"SUBJECT_CODE":   "' '",
	"SUBJECT_NAME":   "' '",
}

var dimSubjectColsSort = []string{
	"UUID",
	"SUBJECT_CODE_1",
	"SUBJECT_NAME_1",
	"SUBJECT_CODE_2",
	"SUBJECT_NAME_2",
	"SUBJECT_CODE_3",
	"SUBJECT_NAME_3",
	"SUBJECT_CODE",
	"SUBJECT_NAME",
}
