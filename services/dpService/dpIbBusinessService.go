package dpService

import (
	"errors"

	"pccqcpa.com.cn/app/rpm/api/models/dp"
	"pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/components/zlog"
)

type DpIbBusinessService struct{}

var dpIbBusiness dp.DpIbBusiness

// 分页操作
// by author Jason
// by time 2016-12-06 16:59:09
func (DpIbBusinessService) List(param ...map[string]interface{}) (*util.PageData, error) {
	return dpIbBusiness.List(param...)
}

// 多参数查询返回多条纪录
// by author Jason
// by time 2016-12-06 16:59:09
func (DpIbBusinessService) Find(param ...map[string]interface{}) ([]*dp.DpIbBusiness, error) {
	return dpIbBusiness.Find(param...)
}

// 多参数查询返回一条纪录
// by author Jason
// by time 2016-12-06 16:59:09
func (this DpIbBusinessService) FindOne(paramMap map[string]interface{}) (*dp.DpIbBusiness, error) {
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
	er := errors.New("查询【RPM_BIZ_DP_IB_BUSINESS】一对一存款中间收入表有多条返回纪录")
	zlog.Error(er.Error(), er)
	return nil, er
}

// 新增纪录
// by author Jason
// by time 2016-12-06 16:59:09
func (DpIbBusinessService) Add(model *dp.DpIbBusiness) error {
	err := model.Add()
	if nil != err {
		return err
	}
	err = DpOnePricingService{}.Patch(map[string]interface{}{
		"business_code": model.BusinessCode,
		"status":        util.DP_ONE_PRICING_STATUS_UNFINISHED,
	})
	if nil != err {
		return err
	}
	return dp.DpOnePricing{BusinessCode: model.BusinessCode}.Patch(map[string]interface{}{"status": util.DP_ONE_PRICING_STATUS_UNFINISHED})
}

// 更新纪录
// by author Jason
// by time 2016-12-06 16:59:09
func (DpIbBusinessService) Update(model *dp.DpIbBusiness) error {
	err := model.Update()
	if nil != err {
		return err
	}
	err = DpOnePricingService{}.Patch(map[string]interface{}{
		"business_code": model.BusinessCode,
		"status":        util.DP_ONE_PRICING_STATUS_UNFINISHED,
	})
	if nil != err {
		return err
	}
	return dp.DpOnePricing{BusinessCode: model.BusinessCode}.Patch(map[string]interface{}{"status": util.DP_ONE_PRICING_STATUS_UNFINISHED})
}

// 删除纪录
// by author Jason
// by time 2016-12-06 16:59:09
func (DpIbBusinessService) Delete(model *dp.DpIbBusiness) error {

	err := DpOnePricingService{}.Patch(map[string]interface{}{
		"business_code": model.BusinessCode,
		"status":        util.DP_ONE_PRICING_STATUS_UNFINISHED,
	})
	if nil != err {
		return err
	}
	return model.Delete()
}
