package dimService

import (
	"database/sql"
	"errors"

	"pccqcpa.com.cn/app/rpm/api/models/bi/dim"

	"pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/components/zlog"
)

type SubjectService struct{}

var dimSubject dim.DimSubject

// 分页操作
// by author Jason
// by time 2016-10-31 15:40:47
func (SubjectService) List(param ...map[string]interface{}) (*util.PageData, error) {
	return dimSubject.List(param...)
}

// 多参数查询返回多条纪录
// by author Jason
// by time 2016-10-31 15:40:47
func (SubjectService) Find(param ...map[string]interface{}) ([]*dim.DimSubject, error) {
	return dimSubject.Find(param...)
}

// 多参数查询返回一条纪录
// by author Jason
// by time 2016-10-31 15:40:47
func (this SubjectService) FindOne(paramMap map[string]interface{}) (*dim.DimSubject, error) {
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
	er := errors.New("查询【RPM_BI_DIM_SUBJECT】贷款业务分析科目维度有多条返回纪录")
	zlog.Error(er.Error(), er)
	return nil, er
}

// 新增纪录
// by author Jason
// by time 2016-10-31 15:40:47
func (SubjectService) Add(model *dim.DimSubject) error {
	return model.Add()
}

// 新增纪录带事务
// by author Jason
// by time 2016-10-31 15:34:54
func (SubjectService) BatchAdd(models []dim.DimSubject) (sql.Result, error) {
	return dimSubject.BatchAdd(models)
}

// 更新纪录
// by author Jason
// by time 2016-10-31 15:40:47
func (SubjectService) Update(model *dim.DimSubject) error {
	return model.Update()
}

// 删除纪录
// by author Jason
// by time 2016-10-31 15:40:47
func (SubjectService) Delete(model *dim.DimSubject) error {
	return model.Delete()
}
