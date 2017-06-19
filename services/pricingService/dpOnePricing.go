package pricingService

import (
	"fmt"
	"pccqcpa.com.cn/app/rpm/api/models/dp"
	"pccqcpa.com.cn/app/rpm/api/services/dpService"
	"pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/components/zlog"
)

type dpOnePricing struct {
	dpOneBusiness     []*dp.DpOneBusiness     // 存款业务信息
	dpOneIbBusiness   []*dp.DpIbBusiness      // 中间业务信息
	dpOneStockPricing []*dp.DpOneStockPricing // 存款存量威胁业务信息
}

// 创建一对一存款定价结构体
func NewDpOnePricing() *dpOnePricing {
	return &dpOnePricing{}
}

func (d *dpOnePricing) SetDpOneBusiess(dpOneBusiness []*dp.DpOneBusiness) {
	d.dpOneBusiness = dpOneBusiness
}

func (d *dpOnePricing) SetDpIbBusiess(dpOneIbBusiness []*dp.DpIbBusiness) {
	d.dpOneIbBusiness = dpOneIbBusiness
}

func (d *dpOnePricing) SetDpOneStockPricing(dpOneStockPricing []*dp.DpOneStockPricing) {
	d.dpOneStockPricing = dpOneStockPricing
}

func (d *dpOnePricing) Pricing(dpOnePricingOld *dp.DpOnePricing) (*dp.DpOnePricing, error) {
	paramMap := map[string]interface{}{
		"business_code": dpOnePricingOld.BusinessCode,
	}
	dpOnePricing, err := dpService.DpOnePricingService{}.FindOne(paramMap)
	if nil != err {
		return nil, err
	} else if nil == dpOnePricing {
		er := fmt.Errorf("查询存款一对一定价单为空")
		zlog.Error(er.Error(), er)
		return nil, er
	}

	if 0 == len(d.dpOneBusiness) {
		paramMap = map[string]interface{}{
			"business_code": dpOnePricing.BusinessCode,
		}
		dpOneBusiness, err := dpService.DpOneBusinessService{}.Find(paramMap)
		if nil != err {
			return nil, err
		}
		d.dpOneBusiness = dpOneBusiness
	}
	if 0 == len(d.dpOneStockPricing) {
		paramMap := map[string]interface{}{
			"business_code": dpOnePricing.BusinessCode,
			// "cust": dpOnePricing.Cust.CustCode,
		}
		dpOneStockPricing, err := dpService.DpOneStockPricingService{}.Find(paramMap)
		if nil != err {
			return nil, err
		}
		d.dpOneStockPricing = dpOneStockPricing
	}
	if 0 == len(d.dpOneIbBusiness) {
		paramMap := map[string]interface{}{
			"business_code": dpOnePricing.BusinessCode,
		}
		dpIbBusiness, err := dpService.DpIbBusinessService{}.Find(paramMap)
		if nil != err {
			return nil, err
		}
		d.dpOneIbBusiness = dpIbBusiness
	}

	// 变量 当前EVA，派生EVA，流失机会成本，未来EVA,办理当前存款损失,存量存款业务流失
	var CurrentEVA, IbEVA, LostCost, NextEVA, currentDpLost, stockDpLost float64 = 0, 0, 0, 0, 0, 0
	// 计算当前业务EVA
	var sumDp, sumIb float64 // 存款总额，中间业务总额
	for _, business := range d.dpOneBusiness {
		if business.Organ.OrganCode != dpOnePricingOld.Organ.OrganCode {
			business.Organ.OrganCode = dpOnePricingOld.Organ.OrganCode
			go business.Update()
		}
		pricingRst, err := rpmEngine.StartPricing(business, util.ONE_DP_CURRENT_EVA)
		if nil != err {
			return nil, err
		}
		for _, val := range pricingRst {
			CurrentEVA += val.(float64)
		}
		sumDp += business.Amount
	}
	// 计算派生中间业务收入EVA
	for _, ibBusiness := range d.dpOneIbBusiness {
		pricingRst, err := rpmEngine.StartPricing(ibBusiness, util.ONE_DP_IB_EVA)
		if nil != err {
			return nil, err
		}
		for _, val := range pricingRst {
			IbEVA += val.(float64)
		}
		sumIb += ibBusiness.Amount
	}

	// 计算存款业务流失机会成本
	for _, lostCost := range d.dpOneStockPricing {
		pricingRst, err := rpmEngine.StartPricing(lostCost, util.ONE_DP_STOCK)
		if nil != err {
			return nil, err
		}
		for key, val := range pricingRst {
			switch key {
			case "DpOneStockLoss":
				currentDpLost += val.(float64)
			case "DpOneStockLost":
				stockDpLost += val.(float64)
			default:
				er := fmt.Errorf("计算存款业务流失时不支持该指标【%v】", key)
				zlog.Error(er.Error(), er)
				return nil, er
			}
		}
		// fmt.Printf("\n\n\n\n\n\n\n存量流失：%+v\n", pricingRst)
	}

	// 计算客户未来贡献
	NextEVA = CurrentEVA + IbEVA

	dpOnePricing.CurrentEva = CurrentEVA // 当前存款EVA
	dpOnePricing.IbEva = IbEVA           // 派生中间业务EVA
	dpOnePricing.NextEva = NextEVA       // 客户未来贡献合计
	dpOnePricing.LostCost = LostCost
	dpOnePricing.CurrentDpLost = NextEVA - currentDpLost // 办理当前存款损失
	dpOnePricing.StockDpLost = stockDpLost               // 存量存款业务流失
	dpOnePricing.Status = util.DP_ONE_PRICING_STATUS_FINISHED
	dpOnePricing.Organ.OrganCode = dpOnePricingOld.Organ.OrganCode
	dpOnePricing.SumDp = sumDp
	dpOnePricing.SumIb = sumIb
	dpOnePricing.SumBreak = sumDp + sumIb

	err = dpService.DpOnePricingService{}.Update(dpOnePricing)
	if nil != err {
		return nil, err
	}

	return dpOnePricing, nil
}
