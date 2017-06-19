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

var lnGuaranteService loanService.LnGuaranteService

type LnGuaranteAction struct {
	tango.Json
	tango.Ctx
}

func (l *LnGuaranteAction) Get() util.RstMsg {
	paramMap, err := util.GetParmFromRouter(&l.Ctx)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	lnGuarantes, err := lnGuaranteService.Find(paramMap)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("多参数查询保证人信息成功", lnGuarantes)
}

func (l *LnGuaranteAction) Post() util.RstMsg {
	var lnGuarante = new(ln.LnGuarante)
	err := currentMsg.DecodeJson(l.Ctx, lnGuarante)
	if nil != err {
		er := fmt.Errorf("json数据转换为保证人结构体数据出错")
		zlog.Error(er.Error(), err)
		return util.ErrorMsg(er.Error(), err)
	}
	err = lnGuaranteService.Add(lnGuarante)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("新增保证人信息成功", lnGuarante)
}

func (l *LnGuaranteAction) Put() util.RstMsg {
	var lnGuarante = new(ln.LnGuarante)
	err := currentMsg.DecodeJson(l.Ctx, lnGuarante)
	if nil != err {
		er := fmt.Errorf("json数据转换为保证人结构体数据出错")
		zlog.Error(er.Error(), err)
		return util.ErrorMsg(er.Error(), err)
	}
	err = lnGuaranteService.Update(lnGuarante)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("更新保证人信息成功", lnGuarante)
}

func (l *LnGuaranteAction) Delete() util.RstMsg {
	var lnGuarante = new(ln.LnGuarante)
	err := l.DecodeJson(lnGuarante)
	if nil != err {
		er := fmt.Errorf("json数据转换为保证人结构体数据出错")
		zlog.Error(er.Error(), err)
		return util.ErrorMsg(er.Error(), err)
	}
	err = lnGuaranteService.Delete(lnGuarante)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("更新保证人信息成功", nil)
}
