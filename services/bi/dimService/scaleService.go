package dimService

import (
	"database/sql"
	"errors"

	"pccqcpa.com.cn/app/rpm/api/models/bi/dim"

	"pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/components/zlog"
)

type ScaleService struct{}

var dimScale dim.DimScale

// 分页操作
// by author Jason
// by time 2016-10-31 15:39:12
func (ScaleService) List(param ...map[string]interface{}) (*util.PageData, error) {
	return dimScale.List(param...)
}

// 多参数查询返回多条纪录
// by author Jason
// by time 2016-10-31 15:39:12
func (ScaleService) Find(param ...map[string]interface{}) ([]*dim.DimScale, error) {
	return dimScale.Find(param...)
}

// 多参数查询返回一条纪录
// by author Jason
// by time 2016-10-31 15:39:12
func (this ScaleService) FindOne(paramMap map[string]interface{}) (*dim.DimScale, error) {
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
	er := errors.New("查询【RPM_BI_DIM_SCALE】贷款业务分析方案规模维度有多条返回纪录")
	zlog.Error(er.Error(), er)
	return nil, er
}

// 新增纪录
// by author Jason
// by time 2016-10-31 15:39:12
func (ScaleService) Add(model *dim.DimScale) error {
	return model.Add()
}

// 新增纪录带事务
// by author Jason
// by time 2016-10-31 15:34:54
func (ScaleService) BatchAdd(models []dim.DimScale) (sql.Result, error) {
	return dimScale.BatchAdd(models)
}

// 更新纪录
// by author Jason
// by time 2016-10-31 15:39:12
func (ScaleService) Update(model *dim.DimScale) error {
	return model.Update()
}

// 删除纪录
// by author Jason
// by time 2016-10-31 15:39:12
func (ScaleService) Delete(model *dim.DimScale) error {
	return model.Delete()
}
