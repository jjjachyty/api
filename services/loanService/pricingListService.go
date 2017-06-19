package loanService

import (
	"errors"
	"sync"

	"pccqcpa.com.cn/app/rpm/api/models/biz/ln"

	"pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/components/zlog"
)

//PricingListService type 定价单实体
type PricingListService struct {
}

var lpm ln.LnPricingModel

// GetLnPringList func PricingListService 获取定价单列表服务方法
// func (pls PricingListService) GetLnPringList() ([]lnpricing.LnPricing, error) {
// 	zlog.AppOperateLog("", "PricingListService.GetLnPringList", zlog.SELECT, nil, nil, "对公贷款定价单服务")
// 	return lnpricing.LnPricingModel{pls.Ctx}.QueryPricingList()
// }

// GetLnPringListByPage func PricingListService 获取定价单列表服务方法
func (pls PricingListService) GetLnPringListByPageWithCustCodeName(startRowNumber int64, pageSize int64, orderAttr string, orderType util.OrderType, parms map[string]interface{}) (util.PageData, error) {
	var pageData util.PageData
	zlog.AppOperateLog("", "PricingListService.GetLnPringList", zlog.SELECT, nil, nil, "对公贷款定价单服务")
	result, page, err := lpm.QueryPricingListByPageAndOrderWithCustCodeName(startRowNumber, pageSize, orderAttr, orderType, parms)
	pageData.Page = page
	pageData.Rows = result
	if nil != err {
		zlog.Error("获取定价单列表服务出错", err)
	}
	return pageData, err
}

// GetLnPringWithPID func PricingListService 获取定价单服务方法
func (pls PricingListService) GetLnPringWithPID(pid string) (util.PageData, error) {
	var pageData util.PageData
	var parms = make(map[string]interface{})
	var lnbs = new(ln.LnBusiness)
	zlog.AppOperateLog("", "PricingListService.GetLnPringWithPID", zlog.SELECT, nil, nil, "对公贷款定价单"+pid+"服务")
	lnp, err := ln.LnPricingModel{}.QueryPricingListWithPID(pid)
	if nil != err {
		return pageData, err
	}
	// if len(lnp) == 0 {
	// 	return pageData, nil
	// 	// return pageData, errors.New("未查询到定价单号为:" + pid + "的信息")
	// }

	//查询基础定价单
	parms["business_code"] = pid
	lnbStrct, err1 := lnbs.Find(parms)
	if nil != err1 {
		return pageData, err1
	}
	if 0 != len(lnbStrct) {
		if len(lnp) == 0 {
			lnp = append(lnp, ln.LnPricing{LnBusiness: *lnbStrct[0]})
		} else if 1 == len(lnp) {
			lnp[0].LnBusiness = *lnbStrct[0]
		}
	}

	pageData.Rows = lnp[0]
	return pageData, err

}

// GetLnPringWithPID func PricingListService 获取定价单服务方法
func (pls PricingListService) Update(lnPricing *ln.LnPricing) error {
	zlog.AppOperateLog("", "PricingListService.Update", zlog.SELECT, nil, nil, "对公贷款定价单"+lnPricing.PlnCode+"更新服务")
	var waitGroup sync.WaitGroup
	waitGroup.Add(2)
	var err1, err2 error
	go func() {
		err1 = lnPricing.Update()
		waitGroup.Done()
	}()
	go func() {
		err2 = ln.LnBusiness{}.Patch(map[string]interface{}{}, lnPricing.BusinessCode)
		waitGroup.Done()
	}()
	waitGroup.Wait()
	if nil != err1 {
		return err1
	}
	return err2
}

func (p PricingListService) Patch(prams map[string]interface{}) error {
	if val, ok := prams["business_code"]; !ok || "" == val {
		er := errors.New("更新对公贷款定价单结果信息，没有传送businessKey主键")
		zlog.Error(er.Error(), er)
		return er
	}
	PKMap := map[string]interface{}{
		"business_code": prams["business_code"],
	}
	delete(prams, "business_code")
	return lpm.Patch(prams, PKMap)
}

func (pls PricingListService) Delete(lnPricing *ln.LnPricing) error {
	if util.PRICING_STATUS_UNFINISHED_LN != lnPricing.Status &&
		util.PRICING_STATUS_UNFINISHED_INVERSE != lnPricing.Status &&
		util.PRICING_STATUS_FINISHED_NO_SAVE != lnPricing.Status &&
		util.PRICING_STATUS_FINISHED_SAVE != lnPricing.Status {
		er := errors.New("不能删审批中的定价单")
		zlog.Infof(er.Error(), er)
		return er
	}
	lnPricing.LnBusiness.BusinessCode = lnPricing.BusinessCode
	err := LnBusinessService{}.Delete(&lnPricing.LnBusiness)
	if nil != err {
		return err
	}
	err = lnPricing.Delete()
	return err
}
