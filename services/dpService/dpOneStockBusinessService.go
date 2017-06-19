package dpService

import (
	"errors"

	"pccqcpa.com.cn/app/rpm/api/models/dp"

	"pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/components/zlog"
)

type DpOneStockBusinessService struct{}

var dpOneStockBusiness dp.DpOneStockBusiness

// 分页操作
// by author Jason
// by time 2017-01-04 10:46:11
func (DpOneStockBusinessService) List(param ...map[string]interface{}) (*util.PageData, error) {
	return dpOneStockBusiness.List(param...)
}

// 多参数查询返回多条纪录
// by author Jason
// by time 2017-01-04 10:46:11
func (DpOneStockBusinessService) Find(param ...map[string]interface{}) ([]*dp.DpOneStockBusiness, error) {
	return dpOneStockBusiness.Find(param...)
}

// 多参数查询返回一条纪录
// by author Jason
// by time 2017-01-04 10:46:11
func (this DpOneStockBusinessService) FindOne(paramMap map[string]interface{}) (*dp.DpOneStockBusiness, error) {
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
	er := errors.New("查询【RPM_BIZ_DP_ONE_STOCK_BUSINESS】一对一存款存量表有多条返回纪录")
	zlog.Error(er.Error(), er)
	return nil, er
}

// 新增纪录
// by author Jason
// by time 2017-01-04 10:46:11
func (DpOneStockBusinessService) Add(model *dp.DpOneStockBusiness) error {
	return model.Add()
}

// 更新纪录
// by author Jason
// by time 2017-01-04 10:46:11
func (DpOneStockBusinessService) Update(model *dp.DpOneStockBusiness) error {
	return model.Update()
}

// 删除纪录
// by author Jason
// by time 2017-01-04 10:46:11
func (DpOneStockBusinessService) Delete(model *dp.DpOneStockBusiness) error {
	return model.Delete()
}
