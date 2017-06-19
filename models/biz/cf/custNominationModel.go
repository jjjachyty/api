package cf

import (
	"database/sql"
	"fmt"
	"platform/dbobj"

	modelsUtil "pccqcpa.com.cn/app/rpm/api/models/util"
	"pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/components/zlog"
)

//客户分类特殊名单model类
//@by Janly
type CustNomination struct {
	UUID      string
	CustCode  string
	CustName  string
	CustImpl  string
	CustType  string
	OrgCode   string
	Remark    string
	RowNumber int
}

type CustNominationModel struct {
}

func (c CustNominationModel) FindAllByPage(page util.Page, order util.Order, params map[string]interface{}) (util.Page, []CustNomination, error) {
	var custNominations []CustNomination
	var querySQL = `SELECT   THIS_.UUID,
                            THIS_.CUST_CODE,
                            THIS_.CUST_NAME,
                            THIS_.CUST_IMPL,
                            THIS_.CUST_TYPE,
                            THIS_.ORG_CODE,
                            THIS_.REMARK
   FROM RPM_CLASSIFY_NOMINATION THIS_`
	var whereSQL string
	if nil != params {
		var step = 0
		for k, v := range params {
			if 0 == step {
				whereSQL = " WHERE "
			} else if step > 0 {
				whereSQL += " OR "
			}
			whereSQL += k + " like '%" + v.(string) + "%' "
			step++
		}
	}
	querySQL += whereSQL

	var orderSQL string
	orderCol := field2Column[order.OrderAttr]

	if "" != orderCol {
		orderSQL = " ORDER BY " + orderCol + " " + string(order.OrderType)
	}
	querySQL += orderSQL
	page = util.GetPage(page.StartRowNumber, page.PageSize, querySQL)
	pageSQL := `SELECT PAGE_2.* FROM (
		SELECT PAGE_1.*,ROWNUM AS rowno FROM ( ` + querySQL + `) PAGE_1  WHERE ROWNUM <=:0) PAGE_2 
		WHERE PAGE_2.rowno >:1`
	zlog.Debugf("SQL:%s", nil, pageSQL)
	rows, err := dbobj.Default.Query(pageSQL, page.EndRowNumber, page.StartRowNumber)
	if nil == err {
		for rows.Next() {
			var custNomination CustNomination
			rows.Scan(&custNomination.UUID,
				&custNomination.CustCode,
				&custNomination.CustName,
				&custNomination.CustImpl,
				&custNomination.CustType,
				&custNomination.OrgCode,
				&custNomination.Remark,
				&custNomination.RowNumber,
			)
			custNominations = append(custNominations, custNomination)
		}
	}
	return page, custNominations, err
}

func (CustNominationModel) scan(rows *sql.Rows) (*CustNomination, error) {
	var custNomination CustNomination
	err := rows.Scan(&custNomination.UUID,
		&custNomination.CustCode,
		&custNomination.CustName,
		&custNomination.CustImpl,
		&custNomination.CustType,
		&custNomination.OrgCode,
		&custNomination.Remark,
	)
	return &custNomination, err
}

func (c CustNominationModel) Find(param ...map[string]interface{}) ([]*CustNomination, error) {
	rows, err := modelsUtil.FindRows("RPM_CLASSIFY_NOMINATION T", custCols, custColsSort, param...)
	if nil != err {
		er := fmt.Errorf("查询客户分类特殊名单信息出错")
		zlog.Error(er.Error(), err)
		return nil, er
	}
	defer rows.Close()
	var custNominations []*CustNomination
	for rows.Next() {
		custInfo, err := c.scan(rows)
		if nil != err {
			er := fmt.Errorf("查询客户分类特殊名单信息rows.Scan()出错")
			zlog.Error(er.Error(), err)
			return nil, er
		}
		custNominations = append(custNominations, custInfo)
	}
	return custNominations, nil
}

func (CustNominationModel) Save(custNomination CustNomination) error {
	var insertSQL = `
	  INSERT
INTO RPM_CLASSIFY_NOMINATION
  (
    CUST_CODE,
    CUST_NAME,
    CUST_IMPL,
    CUST_TYPE,
    ORG_CODE,
    REMARK
  )
  VALUES
  (
    :0,
    :1,
    :2,
    :3,
    :4,
    :5
  )`
	err := dbobj.Default.Exec(insertSQL, custNomination.CustCode, custNomination.CustName, custNomination.CustImpl, custNomination.CustType, custNomination.OrgCode, custNomination.Remark)
	if err != nil {
		zlog.Error("新增客户分类[白名单]失败", err)
	}
	return err
}

// 新增客户分类特殊名单信息带事务
// by author Jason
// by time 2016-10-31 15:34:54
func (this CustNominationModel) BatchAdd(models []CustNomination) (sql.Result, error) {
	sqlTx, err := dbobj.Default.Begin()
	var tableName = "RPM_CLASSIFY_NOMINATION"
	var deleteSQL = "DELETE FROM " + tableName
	var msg = "[客户分类特殊名单信息-" + tableName + "]"
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
			"CUST_CODE": model.CustCode,
			"CUST_NAME": model.CustName,
			"CUST_IMPL": model.CustImpl,
			"CUST_TYPE": model.CustType,
			"ORG_CODE":  model.OrgCode,
			"REMARK":    model.Remark,
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

func (CustNominationModel) Delete(custNomination CustNomination) error {
	var deleteSQL = `DELETE
FROM RPM_CLASSIFY_NOMINATION
WHERE CUST_CODE    = :0`
	err := dbobj.Default.Exec(deleteSQL, custNomination.CustCode)
	if err != nil {
		zlog.Error("删除客户分类[白名单]失败", err)
	}
	return err
}

func (CustNominationModel) Update(custNomination CustNomination) error {
	var updateSQL = `UPDATE RPM_CLASSIFY_NOMINATION
SET CUST_IMPL =:0,
    REMARK =:1
WHERE CUST_CODE = :2`
	err := dbobj.Default.Exec(updateSQL, custNomination.CustImpl, custNomination.Remark, custNomination.CustCode)
	if err != nil {
		zlog.Error("更新客户分类[白名单]失败", err)
	}
	return err
}

var field2Column = map[string]string{
	"CustImpl": "THIS_.CUST_IMPL",
	"CustCode": "THIS_.CUST_CODE",
	"CustName": "THIS_.CUST_NAME",
	"CustType": "THIS_.CUST_TYPE",
}

var custCols = map[string]string{
	"uuid":      "' '",
	"cust_code": "' '",
	"cust_name": "' '",
	"cust_impl": "' '",
	"cust_type": "' '",
	"org_code":  "' '",
	"remark":    "' '",
}
var custColsSort = []string{
	"uuid",
	"cust_code",
	"cust_name",
	"cust_impl",
	"cust_type",
	"org_code",
	"remark",
}
