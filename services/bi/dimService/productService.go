package dimService

import (
	"database/sql"
	"errors"

	"pccqcpa.com.cn/app/rpm/api/models/bi/dim"

	"pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/components/zlog"
)

type ProductService struct{}

var dimProduct dim.DimProduct

// 分页操作
// by author Jason
// by time 2016-10-31 15:38:16
func (ProductService) List(param ...map[string]interface{}) (*util.PageData, error) {
	return dimProduct.List(param...)
}

// 多参数查询返回多条纪录
// by author Jason
// by time 2016-10-31 15:38:16
func (ProductService) Find(param ...map[string]interface{}) ([]*dim.DimProduct, error) {
	return dimProduct.Find(param...)
}

// 多参数查询返回一条纪录
// by author Jason
// by time 2016-10-31 15:38:16
func (this ProductService) FindOne(paramMap map[string]interface{}) (*dim.DimProduct, error) {
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
	er := errors.New("查询【RPM_BI_DIM_PRODUCT】贷款方案分析产品维度表有多条返回纪录")
	zlog.Error(er.Error(), er)
	return nil, er
}

// 新增纪录
// by author Jason
// by time 2016-10-31 15:38:16
func (ProductService) Add(model *dim.DimProduct) error {
	return model.Add()
}

// 新增纪录带事务
// by author Jason
// by time 2016-10-31 15:34:54
func (ProductService) BatchAdd(models []dim.DimProduct) (sql.Result, error) {
	return dimProduct.BatchAdd(models)
}

// 更新纪录
// by author Jason
// by time 2016-10-31 15:38:16
func (ProductService) Update(model *dim.DimProduct) error {
	return model.Update()
}

// 删除纪录
// by author Jason
// by time 2016-10-31 15:38:16
func (ProductService) Delete(model *dim.DimProduct) error {
	return model.Delete()
}
