package rl

import (
	"github.com/lunny/tango"
	"pccqcpa.com.cn/app/rpm/api/models/biz/rl"
	"pccqcpa.com.cn/app/rpm/api/services/retailService"
	"pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/components/zlog"
)

var rlPriceMatrixService retailService.RLPriceMatrixService

type RLPriceMatrixAction struct {
	tango.Json
	tango.Ctx
}

//api/rpm/rlpricematrix
func (RLPriceMatrixAction) List() util.RstMsg {
	var pageData util.PageData
	priceMatrixs, err := rlPriceMatrixService.GetAllRLPriceMatrix()
	if nil == err {
		pageData.Rows = priceMatrixs
		return util.ReturnSuccess("零售价格矩阵查询成功", pageData)
	}
	return util.ErrorMsg("零售价格矩阵查询失败", err)
}

//api/rpm/rlpricematrix
func (rlpm RLPriceMatrixAction) Post() util.RstMsg {

	var rlPriceMatrix rl.RLPriceMatrix
	err := rlpm.DecodeJson(&rlPriceMatrix)
	if nil != err {
		zlog.Error("零售贷款定价实体json转换实体出错", err)
		return util.ErrorMsg("零售贷款定价实体json转换实体出错", err)
	}
	rlPriceMatrix, err = rlPriceMatrixService.GetRLPriceMatrix(rlPriceMatrix)

	if nil == err {
		return util.SuccessMsg("零售贷款定价成功", rlPriceMatrix)
	}
	return util.ErrorMsg("零售贷款定价失败", err)

}
