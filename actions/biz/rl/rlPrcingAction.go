package rl

import (
	"fmt"
	"strconv"

	"github.com/lunny/tango"

	"pccqcpa.com.cn/app/rpm/api/models/biz/rl"
	"pccqcpa.com.cn/app/rpm/api/services/retailService"
	"pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/app/rpm/api/util/currentMsg"
	"pccqcpa.com.cn/components/zlog"
)

var rlPricingService retailService.RLPricingService

type RlPrcingAction struct {
	tango.Json
	tango.Ctx
}

//api/rpm/rlpricing
func (rlp *RlPrcingAction) List() util.RstMsg {
	paramMap := util.GetPageMsg(&rlp.Ctx)
	rlPricingService.Ctx = rlp.Ctx
	pageData, err := rlPricingService.List(paramMap)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessListMsg("查询所有指标分页成功", pageData)
}

//带参数查询，判断header里面的start-row-number是否为空，如果为空，则不分页 放在service层判断
func (rlp *RlPrcingAction) Get() util.RstMsg {

	paramMap, err := util.GetParmFromRouter(&rlp.Ctx)
	if nil != err {
		zlog.Error(err.Error(), err)
		return util.ErrorMsg(err.Error(), err)
	}
	//判断是否为分页查询
	if util.IsPaginQuery(&rlp.Ctx) {
		zlog.Debug("多参数分页查询", nil)
		util.GetPageMsg(&rlp.Ctx, paramMap)
		rlPricingService.Ctx = rlp.Ctx
		pageData, err := rlPricingService.List(paramMap)
		if nil != err {
			return util.ErrorMsg(err.Error(), err)
		}
		return util.SuccessListMsg("多参数查询指标分页成功", pageData)
	} else {
		zlog.Debug("非分页查询", nil)
	}

	amount, err := strconv.ParseFloat(paramMap["amount"].(string), 64)
	deadline, err := strconv.Atoi(paramMap["term"].(string))
	var rlPriceMatrix rl.RLPriceMatrix
	rlPriceMatrix, err = rlPriceMatrixService.GetRate(paramMap["product_code"].(string), amount, deadline, paramMap["term_mult"].(string))
	if nil == err {
		return util.SuccessMsg("目标利率查询成功", rlPriceMatrix)
	}
	return util.ErrorMsg("目标利率查询失败", err)
}

//保存零售货款定价单
func (rlp *RlPrcingAction) Post() util.RstMsg {
	var rlPricing = new(rl.RlPricing)

	var rlMatrix rl.RLPriceMatrix
	err := currentMsg.DecodeJson(rlp.Ctx, rlPricing)
	if nil != err {
		zlog.Error("零售贷款信息实体json转换实体出错", err)
		return util.ErrorMsg("零售贷款信息实体json转换实体出错", err)
	}

	//提交前计算利率
	rlMatrix, err = rlPriceMatrixService.GetRate(rlPricing.Product.ProductCode, rlPricing.Amount, rlPricing.Term, rlPricing.TermMult)
	if nil != err {
		zlog.Error("计算零售贷款利率出错", err)
		return util.ErrorMsg("保存零售贷款定价单出错", err)
	}
	rlPricing.TgtRate = rlMatrix.TgtRate      //设置目标利率
	rlPricing.Status = util.RL_PRICING_STATUS //设置默认状态为0
	err2 := rlPricingService.Add(rlPricing)
	if nil == err2 {
		return util.SuccessMsg("零售货款信息添加成功", rlPricing)
	}
	return util.ErrorMsg("零售货款信息添加失败", err)

}

//修改零售定价状态
func (r *RlPrcingAction) Put() util.RstMsg {
	var rlPricing = new(rl.RlPricing)
	err := r.DecodeJson(&rlPricing)
	if nil != err {
		er := fmt.Errorf("json数据转换为零售定价信息出错")
		zlog.Errorf(er.Error(), err)
		return util.ErrorMsg(er.Error(), err)
	}
	err = rlPricingService.Update(rlPricing)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("更新零售货款定价单状态信息成功", rlPricing)
}
