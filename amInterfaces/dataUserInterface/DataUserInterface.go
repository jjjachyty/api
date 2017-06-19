package dataUserInterface

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"pccqcpa.com.cn/app/rpm/api/models/amData"
	"pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/components/zlog"
)

type AmAuthUser struct {
}

func (a AmAuthUser) GetAmUser(sid string) error {
	ssoSidUrl := util.GetIniStringValue("am", "url")
	sidUser := util.RpmSession{}.GetDefultSesion().Get(sid)

	if nil == sidUser { //本地没有，去请求服务器的
		ssomsge := util.RspData{}
		url := ssoSidUrl + "sid=" + sid + "&login=0&apitype=3"
		resp, err := http.Get(url)
		zlog.Infof("访问AM系统 验证sid是否生效URL【%s】", err, url)

		if err != nil {
			er := fmt.Errorf("验证Sid错误-请求AM系统错误")
			zlog.Error(er.Error(), err)
			return er
		}

		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		zlog.Debugf("验证Sid-返回信息:%s", nil, body)
		if err != nil {
			er := fmt.Errorf("验证Sid错误-解析错误")
			zlog.Error(er.Error(), err)
			return er
		}

		err = json.Unmarshal(body, &ssomsge)
		if nil != err {
			er := fmt.Errorf("验证SID-SSOMesg实体转换错误")
			zlog.Errorf(er.Error(), err)
			return er
		}
		// zlog.Debugf("\n\n验证Sid-返回信息ssomsge:%#v", nil, ssomsge)
		//验证通过
		if "200" == ssomsge.Status && "null" != string(ssomsge.Data) {
			var amUser = make([]amData.AmUser, 0)
			err := json.Unmarshal(ssomsge.Data, &amUser)
			if nil != err {
				er := fmt.Errorf("验证SID-AmUser实体转换错误")
				zlog.Error(er.Error(), err)
				return er
			}
			zlog.Infof("源数据【%v】\n验证SID-AmUser实体转换[%#v]", nil, string(ssomsge.Data), amUser)
			sidUser := new(amData.SidUser)
			sidUser.AmUserToSidUser(amUser[0])
			zlog.Infof("\n\n验证Sid-返回信息amUser:%#v", nil, amUser)
			util.RpmSession{}.GetDefultSesion().Set(sid, *sidUser)
			return nil
		}

		//验证失败
		return fmt.Errorf(ssomsge.Message)

	} else {
		//本地有
		//判断是否过期 最后操作时间+过期时间>现在时间
		return nil
	}
}
