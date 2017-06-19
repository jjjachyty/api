package pricingService

import (
	"fmt"
	// "time"

	"pccqcpa.com.cn/app/rpm/api/engine/pricing"
	"pccqcpa.com.cn/app/rpm/api/models/biz/ln"
	"pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/components/zlog"
	"strings"
)

var rpmEngine pricing.RpmEngine
var base PricingBase
var lnPricing ln.LnPricing

type PricingService struct {
	LnBusiness ln.LnBusiness
}

// 一、根据参数code查询对公定价结果单
// 二、将实体与type参数传入定价引擎计算出定价结果
// 三、保存定价单、返回结果
func (ps PricingService) LnBusinessPricing(businessCode, businessType string, param ...map[string]interface{}) (*ln.LnPricing, error) {

	paramMap := map[string]interface{}{
		"business_code": businessCode,
	}
	lbp, err := lnPricing.Find(paramMap)
	// fmt.Println("--我是查询的定价单实体－－－-", lbp[0])
	if nil != err {
		return nil, err
	}
	if 0 == len(lbp) {
		er := fmt.Errorf("未查询到业务编号[%v]的定价单或者业务信息", businessCode)
		zlog.Error(er.Error(), er)
		return nil, er
	}
	// 反算时，将执行利率、浮动比例、浮动值赋值到pricing结构体
	if 0 != len(param) {
		err = base.ReflectToPricing(param[0], lbp[0]) //执行利率
		if nil != err {
			return nil, err
		}
	}
	// fmt.Println("查询到的实体", lbp[0].LnBusiness.Term, lbp[0].LnBusiness.TermMult)

	// time.Sleep(time.Second * 10)
	pricingRst, err := rpmEngine.StartPricing(lbp[0], businessType)
	// fmt.Printf("\n\n\n定价结果[%v]\n\n\n", pricingRst)
	if nil != err {
		zlog.Error("对公贷款定价引擎计算失败", err)
		return nil, err
	}
	base.LnBusiness = lbp[0].LnBusiness
	// 反射定价结果到定价单中
	err = base.ReflectToPricing(pricingRst, lbp[0])
	if nil != err {
		return nil, err
	}

	// 将业务基础信息赋值给定价单
	base.LnBusinessToLnPricing(lbp[0])
	if util.LN_INVERSE == businessType {
		lbp[0].Status = util.PRICING_STATUS_FINISHED_NO_SAVE
	} else if strings.Contains(businessType, ",LN_INVERSE") { //信贷socket单独处理  by Janly 2017-05-02
		lbp[0].Status = util.PRICING_STATUS_FINISHED_SAVE
	} else {
		lbp[0].Status = util.PRICING_STATUS_UNFINISHED_INVERSE
	}

	// 保存定价单信息
	if "" == lbp[0].UUID {
		err = lbp[0].Add()
		if nil != err {
			return nil, err
		}
	} else {
		err = lbp[0].Update()
		if nil != err {
			return nil, err
		}
	}

	// 更新时间
	ln.LnBusiness{}.Patch(map[string]interface{}{}, lbp[0].BusinessCode)

	return lbp[0], nil

}
