package dpService

import (
	"errors"

	"github.com/lunny/tango"

	"pccqcpa.com.cn/app/rpm/api/models/dp"
	"pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/app/rpm/api/util/currentMsg"
	"pccqcpa.com.cn/components/zlog"
)

type DpOneStockPricingService struct {
	Ctx tango.Ctx
}

var dpOneStockPricing dp.DpOneStockPricing

// 分页操作
// by author Jason
// by time 2017-01-04 10:46:11
func (DpOneStockPricingService) List(param ...map[string]interface{}) (*util.PageData, error) {
	return dpOneStockPricing.List(param...)
}

// 多参数查询返回多条纪录
// by author Jason
// by time 2017-01-04 10:46:11
func (DpOneStockPricingService) Find(param ...map[string]interface{}) ([]*dp.DpOneStockPricing, error) {
	return dpOneStockPricing.Find(param...)
}

// 多参数查询返回一条纪录
// by author Jason
// by time 2017-01-04 10:46:11
func (this DpOneStockPricingService) FindOne(paramMap map[string]interface{}) (*dp.DpOneStockPricing, error) {
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
	er := errors.New("查询【RPM_BIZ_DP_ONE_STOCK_PRICING】一对一存款存量威胁表有多条返回纪录")
	zlog.Error(er.Error(), er)
	return nil, er
}

// 新增纪录
// by author Jason
// by time 2017-01-04 10:46:11
func (dpone DpOneStockPricingService) Add(model *dp.DpOneStockPricing) error {
	paramMap := map[string]interface{}{
		"business_code": model.DpOneStockBusiness,
	}
	dpOneStock, err := DpOneStockBusinessService{}.FindOne(paramMap)
	if nil != err {
		return err
	}
	if dpOneStock == nil {
		er := errors.New("查询存量存款为空")
		zlog.Error(er.Error(), er)
		return er
	}

	// 获取当前办理业务机构信息
	organ, err := currentMsg.GetCurrentUserBranchCode(dpone.Ctx)
	if nil != err {
		return err
	} else {
		model.Organ.OrganCode = organ
	}
	// 将存量存款信息赋值到pricing中
	model.Cust = dpOneStock.Cust
	// model.Organ = dpOneStock.Organ
	model.Product = dpOneStock.Product
	model.Ccy = dpOneStock.Ccy
	model.RateOld = dpOneStock.Rate
	model.AmountOld = dpOneStock.Amount
	model.StartOfDate = dpOneStock.StartOfDate
	model.EndOfDate = dpOneStock.EndOfDate
	model.CurrentAmount = dpOneStock.CurrentAmount
	model.RestTerm = dpOneStock.RestTerm

	param := map[string]interface{}{
		"business_code":         model.BusinessCode,
		"dp_one_stock_business": model.DpOneStockBusiness,
	}
	rst, err := dpone.Find(param)
	if nil != err {
		return err
	}
	if 0 != len(rst) {
		er := errors.New("此笔存款已经交流过，请修改存量业务交流数据")
		zlog.Error(er.Error(), er)
		return er
	}

	err = model.Add()
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
// by time 2017-01-04 10:46:11
func (DpOneStockPricingService) Update(model *dp.DpOneStockPricing) error {
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
// by time 2017-01-04 10:46:11
func (DpOneStockPricingService) Delete(model *dp.DpOneStockPricing) error {
	err := DpOnePricingService{}.Patch(map[string]interface{}{
		"business_code": model.BusinessCode,
		"status":        util.DP_ONE_PRICING_STATUS_UNFINISHED,
	})
	if nil != err {
		return err
	}
	return model.Delete()
}
