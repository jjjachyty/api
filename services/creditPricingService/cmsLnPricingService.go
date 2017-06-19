package creditPricingService

import (
	"fmt"
	"strings"

	"pccqcpa.com.cn/app/rpm/api/models/biz/ln"
	"pccqcpa.com.cn/app/rpm/api/models/creditPricing"
	"pccqcpa.com.cn/app/rpm/api/services/loanService"
	"pccqcpa.com.cn/app/rpm/api/services/pricingService"
	"pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/components/zlog"
)

type CmsLnPricingService struct {
}

var custInfoService loanService.CustInfoService
var lnBusinessService loanService.LnBusinessService
var lnMortService loanService.LnMortService
var lnGuaranteService loanService.LnGuaranteService
var sceneDpService loanService.SceneDpService
var sceneItdService loanService.SceneItdService

// 定价服务类
var pricing pricingService.PricingService

func (c CmsLnPricingService) Handel(cnsLnPricingModel *creditPricing.CmsLnPricing) (*ln.LnPricing, error) {
	// 必输字段验证
	err := c.checkRequired(cnsLnPricingModel)
	if nil != err {
		return nil, err
	}

	// 保存定价单
	// 保存抵质押品
	// 保存保证人
	// 保存存款派生
	// 保存中间派生
	// 保存存量优惠
	err = c.save(cnsLnPricingModel)
	if nil != err {
		return nil, err
	}

	// 基础定价
	// 判断是否有派生，如果有派生定价
	// 反算
	fmt.Println("开始定价")
	businessCode := cnsLnPricingModel.LnBusiness.BusinessCode
	paramMap := map[string]interface{}{
		"IntRate":    cnsLnPricingModel.IntRate,
		"MarginType": cnsLnPricingModel.MarginType,
		"MarginInt":  cnsLnPricingModel.MarginInt,
	}
	businessTypes := util.LN_BUSINESS + "," + util.LN_INVERSE
	var pricingRstMsg *ln.LnPricing
	// 判断是否有派生信息，有则进行派生计算，没有就跳过
	if 0 != len(cnsLnPricingModel.SceneDps) || 0 != len(cnsLnPricingModel.SceneItds) || !c.checkNumberIsZero(cnsLnPricingModel.StockUsage) {
		paramMap["StockUsage"] = cnsLnPricingModel.StockUsage
		businessTypes += "," + util.LN_SCENE
		pricingRstMsg, err = pricing.LnBusinessPricing(businessCode, businessTypes, paramMap)
		if nil != err {
			return nil, err
		}
	} else {
		pricingRstMsg, err = pricing.LnBusinessPricing(businessCode, businessTypes, paramMap)
		if nil != err {
			return nil, err
		}
	}
	return pricingRstMsg, nil
}

// 客户信息直接新增或者修改
// 保存定价业务信息
// 保存抵质押品
// 保存保证人
// 保存存款派生
// 保存中间派生
// 保存存量优惠
func (c CmsLnPricingService) save(cnsLnPricingModel *creditPricing.CmsLnPricing) error {
	var errChLength int = 6
	var errCh = make(chan error, errChLength)
	go c.addOrUpdateCustInfo(&cnsLnPricingModel.CustInfo, errCh)
	go c.addLnBusiness(&cnsLnPricingModel.LnBusiness, errCh)
	go c.addLnMort(cnsLnPricingModel.LnBusiness.BusinessCode, cnsLnPricingModel.LnMorts, errCh)
	go c.addGuarantes(cnsLnPricingModel.LnBusiness.BusinessCode, cnsLnPricingModel.Guarantes, errCh)
	go c.addSceneDps(cnsLnPricingModel.LnBusiness.BusinessCode, cnsLnPricingModel.SceneDps, errCh)
	go c.addSceneItds(cnsLnPricingModel.LnBusiness.BusinessCode, cnsLnPricingModel.SceneItds, errCh)

	var errString string = ""
	for i := 0; i < errChLength; i++ {
		zlog.Infof("多协程保存定价业务信息等待返回【%v/%v】", nil, i+1, errChLength)
		er := <-errCh
		if nil != er {
			errString += er.Error() + "\n"
		}
		fmt.Println("----------", er)
	}
	fmt.Println("errString", errString)
	if "" != errString {
		return fmt.Errorf(errString)
	}
	return nil
}

// 保存中间派生
func (c CmsLnPricingService) addSceneItds(businessCode string, sceneItds []ln.SceneItd, er chan error) {
	// fmt.Println("新增派生中间", len(sceneItds))
	err := sceneItdService.DeleteByBusinessCode(businessCode)
	if nil != err {
		er <- err
		return
	}
	for _, sceneItd := range sceneItds {
		err := sceneItdService.Add(&sceneItd)
		if nil != err {
			er <- err
			return
		}
	}
	er <- nil
	return
}

// 保存存款派生
func (c CmsLnPricingService) addSceneDps(businessCode string, sceneDps []ln.SceneDp, er chan error) {
	// fmt.Println("新增派生存款", len(sceneDps))
	err := sceneDpService.DeleteByBusinessCode(businessCode)
	if nil != err {
		er <- err
		return
	}
	for _, sceneDp := range sceneDps {
		err := sceneDpService.Add(&sceneDp)
		if nil != err {
			er <- err
			return
		}
	}
	er <- nil
	return
}

// 保存保证人
func (c CmsLnPricingService) addGuarantes(businessCode string, guarantes []ln.LnGuarante, er chan error) {
	// fmt.Println("新增保证人", len(guarantes))
	err := lnGuaranteService.DeleteByBusinessCode(businessCode)
	if nil != err {
		er <- err
		return
	}
	for _, guarante := range guarantes {
		err := lnGuaranteService.Add(&guarante)
		if nil != err {
			er <- err
			return
		}
	}
	er <- nil
	return
}

// 保存抵质押品信息
func (c CmsLnPricingService) addLnMort(businessCode string, morts []ln.LnMort, er chan error) {
	// fmt.Println("新增抵质押品", len(morts))
	err := lnMortService.DeleteByBusinessCode(businessCode)
	if nil != err {
		er <- err
		return
	}
	for _, mort := range morts {
		err := lnMortService.Add(&mort)
		if nil != err {
			er <- err
			return
		}
	}
	er <- nil
	return
}

// 保存定价业务信息
func (c CmsLnPricingService) addLnBusiness(lnBusiness *ln.LnBusiness, er chan error) {
	err := lnBusinessService.Delete(lnBusiness)
	if nil != err {
		er <- err
		return
	}
	// fmt.Println("删除业务信息成功")
	err = lnBusinessService.Add(lnBusiness)
	if nil != err {
		er <- err
		return
	}
	er <- nil
	return
	// fmt.Println("新增业务信息成功")
}

// 新增或者更改客户信息
func (c CmsLnPricingService) addOrUpdateCustInfo(custInfo *ln.CustInfo, er chan error) {
	paramMap := map[string]interface{}{
		"cust_code": custInfo.CustCode,
	}
	custInfoOld, err := custInfoService.FindOne(paramMap)
	if nil != err {
		er <- err
		return
	}
	// fmt.Println("nil != custInfoOld", nil != custInfoOld)
	if nil != custInfoOld {
		custInfo.UUID = custInfoOld.UUID
		err := custInfoService.Update(custInfo)
		if nil != err {
			er <- err
			return
		}
	} else {
		err = custInfoService.Add(custInfo)
		if nil != err {
			er <- err
			return
		}
	}
	er <- nil
	return
}

// 检查必要字段
func (c CmsLnPricingService) checkRequired(cnsLnPricingModel *creditPricing.CmsLnPricing) error {
	var stringRequired = make(map[string]string)
	var numberRequired = make(map[string]interface{})

	stringRequired["LnBusiness.BusinessCode"] = cnsLnPricingModel.LnBusiness.BusinessCode
	stringRequired["CustInfo.CustCode"] = cnsLnPricingModel.CustInfo.CustCode

	var errString string = ""
	for k, v := range stringRequired {
		if ok := c.checkStringIsNull(v); ok {
			errString += fmt.Sprintf("[%v]是必输字段，不能为空\n", k)
		}
	}

	for k, v := range numberRequired {
		if ok := c.checkNumberIsZero(v); ok {
			errString += fmt.Sprintf("[%v]是必输字段，不能为0\n", k)
		}
	}

	if "" != errString {
		er := fmt.Errorf(errString)
		zlog.Error(er.Error(), er)
		return er
	}
	return nil
}

// 判断字符串是否为空
// 为空返回 true
// 不为空返回 false
func (c CmsLnPricingService) checkStringIsNull(str string) bool {
	if "" == strings.Replace(str, " ", "", -1) {
		return true
	}
	return false
}

// 判断数字是否为0
// 为0返回 true
// 不为0返回 false
func (c CmsLnPricingService) checkNumberIsZero(num interface{}) bool {
	if 0 == num {
		return true
	}
	return false
}
