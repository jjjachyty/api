package dimService

import (
	"database/sql"
	"errors"

	"pccqcpa.com.cn/app/rpm/api/models/bi/dim"

	"pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/components/zlog"
)

type TermRangeService struct{}

var dimTermRange dim.DimTermRange

// 分页操作
// by author Jason
// by time 2016-10-31 16:46:57
func (TermRangeService) List(param ...map[string]interface{}) (*util.PageData, error) {
	return dimTermRange.List(param...)
}

// 多参数查询返回多条纪录
// by author Jason
// by time 2016-10-31 16:46:57
func (TermRangeService) Find(param ...map[string]interface{}) ([]*dim.DimTermRange, error) {
	return dimTermRange.Find(param...)
}

// 多参数查询返回一条纪录
// by author Jason
// by time 2016-10-31 16:46:57
func (this TermRangeService) FindOne(paramMap map[string]interface{}) (*dim.DimTermRange, error) {
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
	er := errors.New("查询【RPM_BI_DIM_Term_RANGE】贷款业务方案分析期限区间维度有多条返回纪录")
	zlog.Error(er.Error(), er)
	return nil, er
}

// 新增纪录
// by author Jason
// by time 2016-10-31 16:46:57
func (TermRangeService) Add(model *dim.DimTermRange) error {
	return model.Add()
}

// 新增纪录带事务
// by author Jason
// by time 2016-10-31 15:34:54
func (TermRangeService) BatchAdd(models []dim.DimTermRange) (sql.Result, error) {
	return dimTermRange.BatchAdd(models)
}

// 更新纪录
// by author Jason
// by time 2016-10-31 16:46:57
func (TermRangeService) Update(model *dim.DimTermRange) error {
	return model.Update()
}

// 删除纪录
// by author Jason
// by time 2016-10-31 16:46:57
func (TermRangeService) Delete(model *dim.DimTermRange) error {
	return model.Delete()
}
