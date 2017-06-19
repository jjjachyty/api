package actions

import (
	// "fmt"
	"github.com/lunny/tango"
	"github.com/tango-contrib/session"
	"pccqcpa.com.cn/app/rpm/api/util"
)

type AuthAction struct {
	tango.Json
	tango.Ctx
	session.Session
}

func (act AuthAction) Post() util.RspData {

	sid := act.Req().Header.Get("Sid")

	if "" == sid {
		return util.RspData{"F", "SID验证失败,未获取到sid", nil}
	}
	// sidUser := util.GetCacheByCacheName(util.RPM_SID_USER_CACHE, sid)
	sidUser := act.Session.Get(sid)

	if nil == sidUser {
		return util.RspData{"F", "SID验证失败", nil}
	}
	return util.RspData{"S", "SID验证通过", []byte("")}

}
