package dimService

import (
	"errors"

	"pccqcpa.com.cn/app/rpm/api/models/dim"
	"pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/components/zlog"
)

type ProductRefService struct{}

var productRef dim.ProductRef

// 分页操作
// by author Jason
// by time 2017-03-30 13:51:39
func (ProductRefService) List(param ...map[string]interface{}) (*util.PageData, error) {
	return productRef.List(param...)
}

// 多参数查询返回多条纪录
// by author Jason
// by time 2017-03-30 13:51:39
func (ProductRefService) Find(param ...map[string]interface{}) ([]*dim.ProductRef, error) {
	return productRef.Find(param...)
}

// 多参数查询返回一条纪录
// by author Jason
// by time 2017-03-30 13:51:39
func (this ProductRefService) FindOne(paramMap map[string]interface{}) (*dim.ProductRef, error) {
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
	er := errors.New("查询【RPM_DIM_PRODUCT_REF】产品映射关系表有多条返回纪录")
	zlog.Error(er.Error(), er)
	return nil, er
}

// 新增纪录
// by author Jason
// by time 2017-03-30 13:51:39
func (ProductRefService) Add(model *dim.ProductRef) error {
	return model.Add()
}

// 更新纪录
// by author Jason
// by time 2017-03-30 13:51:39
func (ProductRefService) Update(model *dim.ProductRef) error {
	return model.Update()
}

// 删除纪录
// by author Jason
// by time 2017-03-30 13:51:39
func (ProductRefService) Delete(model *dim.ProductRef) error {
	return model.Delete()
}
