package cf

import (
	"database/sql"
	"fmt"
	"platform/dbobj"
	"strings"
	"time"

	modelsUtil "pccqcpa.com.cn/app/rpm/api/models/util"
	"pccqcpa.com.cn/components/zlog"
)

type ClassifyResult struct {
	CustId                       string    // 客户号
	FinalCustClassification      string    // 最终客户分类 可修改
	AutoCustClassification       string    // 自动客户分类 可修改
	LastCustClassification       string    // 上次客户分类
	ClassificationDate           time.Time // 客户分类时间
	ClassificationAdjustmentDate time.Time // 客户分类调整时间
	AdjustBy                     string    // 调整人员
	UpdateUser                   string    `json:'-'` // 更新人员
}

func (ClassifyResult) scan(rows *sql.Rows) (*ClassifyResult, error) {
	var one = new(ClassifyResult)
	err := rows.Scan(
		&one.CustId,
		&one.FinalCustClassification,
		&one.AutoCustClassification,
		&one.LastCustClassification,
		&one.ClassificationDate,
		&one.ClassificationAdjustmentDate,
		&one.AdjustBy,
	)
	return one, err
}

func (c ClassifyResult) Find(param ...map[string]interface{}) ([]*ClassifyResult, error) {
	rows, err := modelsUtil.FindRows(classifyResultTables, classifyResultCols, classifyResultColsSort, param...)
	if nil != err {
		er := fmt.Errorf("查询客户分类结果表信息出错")
		zlog.Error(er.Error(), err)
		return nil, er
	}
	defer rows.Close()
	var rst []*ClassifyResult
	for rows.Next() {
		one, err := c.scan(rows)
		if nil != err {
			er := fmt.Errorf("查询客户分类结果表信息rows.Scan()出错")
			zlog.Error(er.Error(), err)
			return nil, er
		}
		rst = append(rst, one)
	}
	return rst, nil
}

func (c ClassifyResult) Add() error {
	err := dbobj.Default.Exec(`insert into RPM_CLASSIFY_RESULT(
			cust_id,
			final_cust_classification,
			auto_cust_classification,
			last_cust_classification,
			classification_date,
			classification_adjustment_date,
			adjust_by
		) values(:1, :2, :3, :4, :5, :6, :7)`,
		strings.TrimSpace(c.CustId),
		strings.TrimSpace(c.FinalCustClassification),
		strings.TrimSpace(c.AutoCustClassification),
		strings.TrimSpace(c.LastCustClassification),
		time.Now(),
		time.Now(),
		strings.TrimSpace(c.UpdateUser),
	)
	if nil != err {
		er := fmt.Errorf("新增客户分类结果表信息出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

func (c ClassifyResult) Update() error {
	err := dbobj.Default.Exec(`update RPM_CLASSIFY_RESULT set 
			final_cust_classification = :1,
			auto_cust_classification = :2,
			last_cust_classification = :3,
			classification_date = :4,
			classification_adjustment_date = :5,
			adjust_by = :6
			where cust_id = :7`,
		strings.TrimSpace(c.FinalCustClassification),
		strings.TrimSpace(c.AutoCustClassification),
		strings.TrimSpace(c.LastCustClassification),
		time.Now(),
		time.Now(),
		strings.TrimSpace(c.UpdateUser),
		strings.TrimSpace(c.CustId),
	)

	if nil != err {
		er := fmt.Errorf("更新客户分类结果表信息出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

var classifyResultTableName = "RPM_CLASSIFY_RESULT"

var classifyResultTables = "RPM_CLASSIFY_RESULT T"

var classifyResultCols = map[string]string{
	"cust_id":                        "' '",
	"final_cust_classification":      "' '",
	"auto_cust_classification":       "' '",
	"last_cust_classification":       "' '",
	"classification_date":            "sysdate",
	"classification_adjustment_date": "sysdate",
	"adjust_by":                      "' '",
}

var classifyResultColsSort = []string{
	"cust_id",
	"final_cust_classification",
	"auto_cust_classification",
	"last_cust_classification",
	"classification_date",
	"classification_adjustment_date",
	"adjust_by",
}
