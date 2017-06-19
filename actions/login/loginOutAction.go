package login

import (
	"github.com/lunny/tango"

	"pccqcpa.com.cn/app/rpm/api/amInterfaces"
	"pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/app/rpm/api/util/currentMsg"
)

type LoginOutAction struct {
	tango.Ctx
	tango.Json
}

func (l *LoginOutAction) Post() util.RstMsg {
	// var sid = l.Req().Header.Get("Sid")
	sidUser, err := currentMsg.GetCurrentUser(l.Ctx)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	err = amInterfaces.AmInterface{Sid: sidUser.Sid, UserId: sidUser.UserId}.LoginOut()
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("注销成功", nil)
}
