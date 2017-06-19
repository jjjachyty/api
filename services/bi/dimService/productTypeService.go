package dimService

import (
	"database/sql"
	"errors"

	"pccqcpa.com.cn/app/rpm/api/models/bi/dim"

	"pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/components/zlog"
)

type ProductTypeService struct{}

var dimProductType dim.DimProductType

// 分页操作
// by author Jason
// by time 2016-10-31 15:38:44
func (ProductTypeService) List(param ...map[string]interface{}) (*util.PageData, error) {
	return dimProductType.List(param...)
}

// 多参数查询返回多条纪录
// by author Jason
// by time 2016-10-31 15:38:44
func (ProductTypeService) Find(param ...map[string]interface{}) ([]*dim.DimProductType, error) {
	return dimProductType.Find(param...)
}

// 多参数查询返回一条纪录
// by author Jason
// by time 2016-10-31 15:38:44
func (this ProductTypeService) FindOne(paramMap map[string]interface{}) (*dim.DimProductType, error) {
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
	er := errors.New("查询【RPM_BI_DIM_PRODUCT_TYPE】业务分析方案产品类型维度表有多条返回纪录")
	zlog.Error(er.Error(), er)
	return nil, er
}

// 新增纪录
// by author Jason
// by time 2016-10-31 15:38:44
func (ProductTypeService) Add(model *dim.DimProductType) error {
	return model.Add()
}

// 新增纪录带事务
// by author Jason
// by time 2016-10-31 15:34:54
func (ProductTypeService) BatchAdd(models []dim.DimProductType) (sql.Result, error) {
	return dimProductType.BatchAdd(models)
}

// 更新纪录
// by author Jason
// by time 2016-10-31 15:38:44
func (ProductTypeService) Update(model *dim.DimProductType) error {
	return model.Update()
}

// 删除纪录
// by author Jason
// by time 2016-10-31 15:38:44
func (ProductTypeService) Delete(model *dim.DimProductType) error {
	return model.Delete()
}
