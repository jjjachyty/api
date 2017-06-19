package amAuthData

import (
	"fmt"
	"strings"
	"sync"

	"github.com/lunny/tango"
	"pccqcpa.com.cn/app/rpm/api/amInterfaces/dataAuthInterface"
	"pccqcpa.com.cn/app/rpm/api/amInterfaces/dataUserInterface"
	"pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/components/zlog"
)

type AmAuthDataAction struct {
	tango.Ctx
	tango.Json
}

func (am *AmAuthDataAction) Post() util.RstMsg {
	var waitgroup sync.WaitGroup
	waitgroup.Add(2)
	userId := strings.TrimSpace(am.Form("userid"))
	sid := strings.TrimSpace(am.Req().Header.Get("Sid"))
	fmt.Println("userid,sid", userId, sid)
	if "" != userId && "" != sid {
		var err1, err2 error
		go func() {
			err1 = (&dataAuthInterface.DataAuthResp{Sid: sid, UserId: userId}).Init()
			waitgroup.Done()
		}()
		go func() {
			err2 = dataUserInterface.AmAuthUser{}.GetAmUser(sid)
			waitgroup.Done()
		}()
		waitgroup.Wait()
		if nil != err1 || nil != err2 {
			zlog.Errorf("请求AM系统数据出错[%v][%v]", nil, err1, err2)
			return util.ErrorMsg("请求AM系统数据出错", nil)
		}
	} else {
		er := fmt.Errorf("初始化数据权限未传userid[%s]或sid[%s]", userId, sid)
		zlog.Errorf(er.Error(), er)
		return util.ErrorMsg(er.Error(), er)
	}
	return util.SuccessMsg("【"+userId+"】用户数据权限对接成功", nil)
}
