package creditInterface

import (
	"github.com/lunny/tango"

	"pccqcpa.com.cn/app/rpm/api/models/creditPricing"
	"pccqcpa.com.cn/app/rpm/api/services/creditPricingService"
	"pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/components/zlog"
)

type CmsLnPricingAction struct {
	tango.Ctx
	tango.Json
}

func (c *CmsLnPricingAction) Post() util.RstMsg {
	cmsLnPricingModel := new(creditPricing.CmsLnPricing)
	err := c.DecodeJson(cmsLnPricingModel)
	if nil != err {
		zlog.Error("json数据转换为信息定价实体出错", err)
		return util.ErrorMsg("json数据转换为信息定价实体出错", err)
	}
	lnPricing, err := creditPricingService.CmsLnPricingService{}.Handel(cmsLnPricingModel)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("定价计算成功并保存", lnPricing)
}
