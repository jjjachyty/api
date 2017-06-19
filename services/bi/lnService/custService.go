package lnService

import (
	"errors"
	"fmt"

	"pccqcpa.com.cn/app/rpm/api/models/bi/ln"

	"pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/components/zlog"

	"platform/dbobj"
)

type CustService struct{}

var lnCust ln.LnCust

// 分页操作
// by author Jason
// by time 2016-10-31 15:44:32
func (CustService) List(param ...map[string]interface{}) (*util.PageData, error) {
	return lnCust.List(param...)
}

// 多参数查询返回多条纪录
// by author Jason
// by time 2016-10-31 15:44:32
func (CustService) Find(param ...map[string]interface{}) ([]*ln.LnCust, error) {
	return lnCust.Find(param...)
}

// 多参数查询返回一条纪录
// by author Jason
// by time 2016-10-31 15:44:32
func (this CustService) FindOne(paramMap map[string]interface{}) (*ln.LnCust, error) {
	models, err := this.Find(paramMap)
	if nil != err {
		return nil, err
	}
	switch len(models) {
	case 0:
		return nil, nil
	case 1:
		return models[0], nil
	}
	er := errors.New("查询【RPM_BI_LN_CUST】贷款业务分析客户维度表有多条返回纪录")
	zlog.Error(er.Error(), er)
	return nil, er
}

// 新增纪录
// by author Jason
// by time 2016-10-31 15:44:32
func (CustService) Add(model *ln.LnCust) error {
	return model.Add()
}

// 新增纪录带事务
// by author Jason
// by time 2016-10-31 15:34:54
func (CustService) BatchAdd(models []ln.LnCust) error {
	tx, err := dbobj.Default.Begin()
	if nil == err {
		for k, v := range models {
			result, err := v.BatchAdd(tx)
			if err != nil {
				tx.Rollback()
				return err
			}
			fmt.Print("第", k, "条")
			fmt.Print(result.RowsAffected())
		}
		tx.Commit()
	} else {
		tx.Rollback()
	}
	return err
}

// 更新纪录
// by author Jason
// by time 2016-10-31 15:44:32
func (CustService) Update(model *ln.LnCust) error {
	return model.Update()
}

// 删除纪录
// by author Jason
// by time 2016-10-31 15:44:32
func (CustService) Delete(model *ln.LnCust) error {
	return model.Delete()
}
