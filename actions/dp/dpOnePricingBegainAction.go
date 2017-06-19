package dp

import (
	"fmt"
	"github.com/lunny/tango"

	"pccqcpa.com.cn/app/rpm/api/models/biz/ln"
	"pccqcpa.com.cn/app/rpm/api/models/dp"
	"pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/app/rpm/api/util/currentMsg"
	"pccqcpa.com.cn/components/zlog"
)

type DpOneBegainPricingAction struct {
	tango.Json
	tango.Ctx
}

// 新增信息
// by author Jason
// by time 2016-12-12 10:44:29
func (this *DpOneBegainPricingAction) Post() util.RstMsg {
	// var custCode = this.Form("custCode")
	// if "" == strings.TrimSpace(custCode) {
	// 	er := fmt.Errorf("未传客户编号【custCode】")
	// 	zlog.Error(er.Error(), er)
	// 	return util.ErrorMsg(er.Error(), er)
	// }

	var one = new(ln.CustInfo)
	err := this.DecodeJson(one)
	if nil != err {
		er := fmt.Errorf("json转换实体出错")
		zlog.Error(er.Error(), err)
		return util.ErrorMsg(er.Error(), err)
	}

	var dpOnePricing = new(dp.DpOnePricing)
	err = currentMsg.DecodeJsonUser(this.Ctx, dpOnePricing) // 添加当前用户
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	organCode, err := currentMsg.GetCurrentUserBranchCode(this.Ctx)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	dpOnePricing.Organ.OrganCode = organCode

	err = dpOnePricingService.BegainPricing(dpOnePricing, *one)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("新增存款一对一信息成功", dpOnePricing)
}
