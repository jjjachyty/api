package middlewares

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/lunny/tango"
	"pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/components/zlog"
)

// var SidSession = make(map[string]SidUser)

var mutes sync.Mutex

func VerificationSid() tango.HandlerFunc {
	return func(res *tango.Context) {

		if "/" == res.Req().URL.Path ||
			strings.Contains(res.Req().URL.Path, ".js") ||
			strings.Contains(res.Req().URL.Path, ".html") ||
			("POST" == res.Req().Method && "/api/rpm/work/flow/ln/pricing" == res.Req().URL.Path) {
			res.Next()
			return
		}
		mutes.Lock()
		if "/api/rpm/am/authdata" == res.Req().URL.Path && "POST" == res.Req().Method {

			res.Next()
			mutes.Unlock()
			return
		}
		mutes.Unlock()

		sid := res.Req().Header.Get("Sid")
		zlog.Infof("前台跳转访问HOST【%v】，参数Sid【%v】，res.Req().URL.Path【%v】", nil, res.Req().Host, sid, res.Req().URL.Path)
		if !strings.HasPrefix(res.Req().URL.Path, "/api/rpm") {
			res.Next()
		} else if "" == sid {
			zlog.Error("验证Sid错误-SID为空", nil)
			res.Abort(403, "{\"RstCode\": 403, \"RstMsg\": \"拒绝访问\"}")
		} else {
			err := checkSidValid(sid)
			// zlog.Infof("\n\n\n\n\nSidSession:\n%#v", nil, SidSession)
			if nil == err {
				res.Next()
			} else {
				// 检查是否是被挤出去的
				mutes.Lock()
				defer mutes.Unlock()
				msg := util.GetCacheByCacheName(util.RPM_SQUEEZE_CACHE, sid)
				if nil != msg {
					go func() {
						time.AfterFunc(time.Second*5, func() {
							util.DeleteCacheByCacheName(util.RPM_SQUEEZE_CACHE, sid)
						})
					}()
					res.Abort(403, "{\"RstCode\": 403, \"RstMsg\":\"本用户已在其他终端登录！\"}")
				} else {
					// res.Abort(403, err.Error())
					res.Abort(403, "{\"RstCode\": 403, \"RstMsg\": \""+err.Error()+"\"}")
				}
			}
		}
	}
}

// func checkSidValid(sid string) error {
// 	//获取本地sid
// 	sidUser := SidSession[sid]
// 	ssoSidUrl := util.GetIniStringValue("am", "url")
// 	if "" == sidUser.UserName { //本地没有，去请求服务器的
// 		ssomsge := util.RspData{}
// 		url := ssoSidUrl + "sid=" + sid + "&login=0&apitype=3"
// 		resp, err := http.Get(url)
// 		zlog.Infof("访问AM系统 验证sid是否生效URL【%s】", err, url)

// 		if err != nil {
// 			er := fmt.Errorf("验证Sid错误-请求AM系统错误")
// 			zlog.Error(er.Error(), err)
// 			return er
// 		}

// 		defer resp.Body.Close()
// 		body, err := ioutil.ReadAll(resp.Body)
// 		zlog.Debugf("验证Sid-返回信息:%s", nil, body)
// 		if err != nil {
// 			er := fmt.Errorf("验证Sid错误-解析错误")
// 			zlog.Error(er.Error(), err)
// 			return er
// 		}

// 		err = json.Unmarshal(body, &ssomsge)
// 		if nil != err {
// 			er := fmt.Errorf("验证SID-SSOMesg实体转换错误")
// 			zlog.Errorf(er.Error(), err)
// 			return er
// 		}
// 		zlog.Debugf("\n\n验证Sid-返回信息ssomsge:%#v", nil, ssomsge)
// 		//验证通过
// 		if "200" == ssomsge.Status {
// 			var amUser = make([]AmUser, 0)
// 			err := json.Unmarshal(ssomsge.Data, &amUser)
// 			if nil != err {
// 				er := fmt.Errorf("验证SID-AmUser实体转换错误")
// 				zlog.Error(er.Error(), err)
// 				return er
// 			}
// 			zlog.Infof("源数据【%v】\n验证SID-AmUser实体转换[%#v]", nil, string(ssomsge.Data), amUser)
// 			// amUser := ssomsge.Data.([]interface{})[0].(map[string]interface{})
// 			(&sidUser).amUserToSidUser(amUser[0])
// 			zlog.Infof("\n\n验证Sid-返回信息amUser:%#v", nil, amUser)
// 			SidSession[sid] = sidUser
// 			return nil
// 		}

// 		//验证失败
// 		return fmt.Errorf(ssomsge.Message)

// 	}

// 	//本地有
// 	//判断是否过期 最后操作时间+过期时间>现在时间

// 	if sidUser.LastOpTime.Add(sidUser.ExpTime * time.Minute).After(time.Now()) {
// 		sidUser.LastOpTime = time.Now()
// 		return nil
// 	}

// 	//session过期
// 	delete(SidSession, sid)
// 	zlog.Debugf("验证Sid-Sid:%s已过期", nil, sid)
// 	return fmt.Errorf("会话已过期")
// }

func checkSidValid(sid string) error {
	//获取本地sid
	val := util.GetCacheByCacheName(util.RPM_SQUEEZE_CACHE, sid)
	if nil != val {
		return fmt.Errorf("用户在其他终端登录,该终端已下线,请确保您的密码是否已泄露。")
	}
	// sidUser := util.GetCacheByCacheName(util.RPM_SID_USER_CACHE, sid)
	sidUser := util.RpmSession{}.GetDefultSesion().Get(sid)
	if nil != sidUser {
		return nil
	} else {
		er := fmt.Errorf("您已长时间未操作本系统,系统默认开启账户保护,请重新登录")
		zlog.Error(er.Error(), er)
		return er
	}
}
