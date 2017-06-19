package ln

import (
	"fmt"

	"github.com/lunny/tango"

	"pccqcpa.com.cn/app/rpm/api/models/biz/ln"
	"pccqcpa.com.cn/app/rpm/api/services/loanService"
	"pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/app/rpm/api/util/currentMsg"
	"pccqcpa.com.cn/components/zlog"
)

type LnMortAction struct {
	tango.Json
	tango.Ctx
}

var lnMortService loanService.LnMortService

func (l *LnMortAction) List() util.RstMsg {
	paramMap := util.GetPageMsg(&l.Ctx)
	pageData, err := lnMortService.List(paramMap)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessListMsg("分页查询抵质押品信息成功", pageData)
}

func (l *LnMortAction) Get() util.RstMsg {
	paramMap, err := util.GetParmFromRouter(&l.Ctx)
	if nil != err {
		zlog.Error(err.Error(), err)
		return util.ErrorMsg(err.Error(), err)
	}
	lnMorts, err := lnMortService.Find(paramMap)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("查询押品信息成功", lnMorts)
}

func (l *LnMortAction) Post() util.RstMsg {
	var lnMort = new(ln.LnMort)
	err := currentMsg.DecodeJson(l.Ctx, lnMort)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	err = lnMortService.Add(lnMort)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("新增押品信息成功", lnMort)
}

func (l *LnMortAction) Delete() util.RstMsg {
	var lnMort = new(ln.LnMort)
	err := l.DecodeJson(&lnMort)
	if nil != err {
		er := fmt.Errorf("删除抵质押品信息出错")
		return util.ErrorMsg(er.Error(), er)
	}
	err = lnMortService.Delete(lnMort)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("删除抵质押品信息成功", nil)
}

func (l *LnMortAction) Put() util.RstMsg {
	var lnMort = new(ln.LnMort)
	err := currentMsg.DecodeJson(l.Ctx, lnMort)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	err = lnMortService.Update(lnMort)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("更新押品信息成功", lnMort)
}
