package dimService

import (
	"database/sql"
	"errors"

	"pccqcpa.com.cn/app/rpm/api/models/bi/dim"

	"pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/components/zlog"
)

type LineService struct{}

var dimLine dim.DimLine

// 分页操作
// by author Jason
// by time 2016-10-31 15:33:31
func (LineService) List(param ...map[string]interface{}) (*util.PageData, error) {
	return dimLine.List(param...)
}

// 多参数查询返回多条纪录
// by author Jason
// by time 2016-10-31 15:33:31
func (LineService) Find(param ...map[string]interface{}) ([]*dim.DimLine, error) {
	return dimLine.Find(param...)
}

// 多参数查询返回一条纪录
// by author Jason
// by time 2016-10-31 15:33:31
func (this LineService) FindOne(paramMap map[string]interface{}) (*dim.DimLine, error) {
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
	er := errors.New("查询【RPM_BI_DIM_LINE】贷款业务方案分析业务条线表有多条返回纪录")
	zlog.Error(er.Error(), er)
	return nil, er
}

// 新增纪录
// by author Jason
// by time 2016-10-31 15:33:31
func (LineService) Add(model *dim.DimLine) error {
	return model.Add()
}

// 新增纪录带事务
// by author Jason
// by time 2016-10-31 15:34:54
func (LineService) BatchAdd(models []dim.DimLine) (sql.Result, error) {
	return dimLine.BatchAdd(models)
}

// 更新纪录
// by author Jason
// by time 2016-10-31 15:33:31
func (LineService) Update(model *dim.DimLine) error {
	return model.Update()
}

// 删除纪录
// by author Jason
// by time 2016-10-31 15:33:31
func (LineService) Delete(model *dim.DimLine) error {
	return model.Delete()
}
