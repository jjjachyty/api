package ln

import (
	"errors"

	"github.com/lunny/tango"

	"pccqcpa.com.cn/app/rpm/api/services/loanService"
	"pccqcpa.com.cn/app/rpm/api/util"
)

type CustInfoGuranteAction struct {
	tango.Json
	tango.Ctx
}

// 保证人要求能查看全行的
func (c *CustInfoGuranteAction) Get() util.RstMsg {
	var custInfoService loanService.CustInfoService
	paramMap, err := util.GetParmFromRouter(&c.Ctx)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	if util.IsPaginQuery(&c.Ctx) {
		util.GetPageMsg(&c.Ctx, paramMap)
		flag := util.CheckQueryParams(paramMap)
		if flag {
			return util.ErrorMsg("查询条件中不允许出现[^$.*+?{}()[]|'`~\\]特殊字符", errors.New("查询条件中不允许出现[^$.*+?{}()[]|'`~\\]特殊字符"))
		}

		pageData, err := custInfoService.List(paramMap)
		if nil != err {
			return util.ErrorMsg(err.Error(), err)
		}
		return util.SuccessListMsg("分页查询客户信息成功", pageData)
	}
	cust, err := custInfoService.FindOne(paramMap)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("查询单个客户信息成功", cust)
}
