package login

import (
	"github.com/lunny/tango"

	"pccqcpa.com.cn/app/rpm/api/amInterfaces"
	"pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/app/rpm/api/util/currentMsg"
)

type LoginOutAmAction struct {
	tango.Ctx
	tango.Json
}

func (l *LoginOutAmAction) Post() util.RstMsg {
	sidUser, err := currentMsg.GetCurrentUser(l.Ctx)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	err = amInterfaces.AmInterface{Sid: sidUser.Sid, UserId: sidUser.UserId}.LoginOutByAm()
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("注销成功", nil)
}
