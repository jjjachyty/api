package lnService

import (
	"errors"
	"fmt"

	"pccqcpa.com.cn/app/rpm/api/models/bi/ln"

	"pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/components/zlog"

	"platform/dbobj"
)

type AmountRangeService struct{}

var lnAmountRange ln.LnAmountRange

// 分页操作
// by author Jason
// by time 2016-10-31 15:42:39
func (AmountRangeService) List(param ...map[string]interface{}) (*util.PageData, error) {
	return lnAmountRange.List(param...)
}

// 多参数查询返回多条纪录
// by author Jason
// by time 2016-10-31 15:42:39
func (AmountRangeService) Find(param ...map[string]interface{}) ([]*ln.LnAmountRange, error) {
	return lnAmountRange.Find(param...)
}

// 多参数查询返回一条纪录
// by author Jason
// by time 2016-10-31 15:42:39
func (this AmountRangeService) FindOne(paramMap map[string]interface{}) (*ln.LnAmountRange, error) {
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
	er := errors.New("查询【RPM_BI_LN_AMOUNT_RANGE】金额区间维度表有多条返回纪录")
	zlog.Error(er.Error(), er)
	return nil, er
}

// 新增纪录
// by author Jason
// by time 2016-10-31 15:42:39
func (AmountRangeService) Add(model *ln.LnAmountRange) error {
	return model.Add()
}

// 新增纪录带事务
// by author Jason
// by time 2016-10-31 15:34:54
func (AmountRangeService) BatchAdd(models []ln.LnAmountRange) error {
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
// by time 2016-10-31 15:42:39
func (AmountRangeService) Update(model *ln.LnAmountRange) error {
	return model.Update()
}

// 删除纪录
// by author Jason
// by time 2016-10-31 15:42:39
func (AmountRangeService) Delete(model *ln.LnAmountRange) error {
	return model.Delete()
}
