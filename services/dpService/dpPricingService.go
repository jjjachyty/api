package dpService

import (
	"errors"

	"pccqcpa.com.cn/app/rpm/api/models/dp"
	"pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/components/zlog"
)

type DpPricingService struct{}

var dpPricing dp.DpPricing

// 分页操作
// by author Jason
// by time 2016-12-01 10:24:52
func (DpPricingService) List(param ...map[string]interface{}) (*util.PageData, error) {
	return dpPricing.List(param...)
}

// 多参数查询返回多条纪录
// by author Jason
// by time 2016-12-01 10:24:52
func (DpPricingService) Find(param ...map[string]interface{}) ([]*dp.DpPricing, error) {
	return dpPricing.Find(param...)
}

// 多参数查询返回一条纪录
// by author Jason
// by time 2016-12-01 10:24:52
func (this DpPricingService) FindOne(paramMap map[string]interface{}) (*dp.DpPricing, error) {
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
	er := errors.New("查询有多条返回纪录")
	zlog.Error(er.Error(), er)
	return nil, er
}

// 新增纪录
// by author Jason
// by time 2016-12-01 10:24:52
func (dp DpPricingService) Add(model *dp.DpPricing) error {
	return model.Add()
}

// 更新纪录
// by author Jason
// by time 2016-12-01 10:24:52
func (DpPricingService) Update(model *dp.DpPricing) error {
	return model.Update()
}

// 删除纪录
// by author Jason
// by time 2016-12-01 10:24:52
func (DpPricingService) Delete(model *dp.DpPricing) error {
	return model.Delete()
}
