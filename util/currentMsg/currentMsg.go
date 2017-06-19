package currentMsg

import (
	"encoding/json"
	"fmt"

	"github.com/lunny/tango"

	"pccqcpa.com.cn/app/rpm/api/models/amData"
	"pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/components/zlog"
)

func GetCurrentUser(ctx tango.Ctx) (amData.SidUser, error) {
	sid := ctx.Req().Header.Get("Sid")
	if "" == sid {
		er := fmt.Errorf("验证Sid错误-SID为空")
		zlog.Error(er.Error(), er)
		return amData.SidUser{}, er
	}
	sidUser := util.RpmSession{}.GetDefultSesion().Get(sid)
	if nil == sidUser {
		return amData.SidUser{}, fmt.Errorf("Sid已过期")
	}
	return sidUser.(amData.SidUser), nil
}

func GetCurrentUserName(ctx tango.Ctx) (string, error) {
	sid := ctx.Req().Header.Get("Sid")
	if "" == sid {
		er := fmt.Errorf("验证Sid错误-SID为空")
		zlog.Error(er.Error(), er)
		return "", er
	}

	//获取本地sid
	sidUser := util.RpmSession{}.GetDefultSesion().Get(sid)

	if nil == sidUser {
		er := fmt.Errorf("获取当前登录用户出错")
		zlog.Error(er.Error(), er)
		return "", er
	}
	return sidUser.(amData.SidUser).UserName, nil
}

// decodeJson时加入创建人 更新人
func DecodeJson(ctx tango.Ctx, obj interface{}) error {
	bytes, err := ctx.Body()
	if nil != err {
		return err
	} else if 0 != len(bytes) {
		err := ctx.DecodeJson(obj)
		if nil != err {
			er := fmt.Errorf("json转换实体出错")
			zlog.Infof("实体：%#v\n转换json：%v\n", err, obj, string(bytes))
			zlog.Error(er.Error(), er)
			return er
		}
	}
	return DecodeJsonUser(ctx, obj)
}

func DecodeJsonUser(ctx tango.Ctx, obj interface{}) error {
	userName, err := GetCurrentUserName(ctx)
	if nil != err {
		return err
	}
	var jsonStr string = ""
	if "POST" == ctx.Req().Method {
		jsonStr = "{\"CreateUser\":\"" + userName + "\",\"UpdateUser\":\"" + userName + "\"}"
		err = json.Unmarshal([]byte(jsonStr), obj)
	} else if "PUT" == ctx.Req().Method {
		jsonStr = "{\"UpdateUser\":\"" + userName + "\"}"
		err = json.Unmarshal([]byte(jsonStr), obj)
	}
	if nil != err {
		er := fmt.Errorf("json转换实体出错(加入当前用户)")
		zlog.Error(er.Error(), err)
		zlog.Infof("json转换实体出错(加入当前用户)\n实体：%#v\n转换json：%v", err, obj, jsonStr)
		return er
	}
	return nil
}

func GetCurrentUserBranchCode(ctx tango.Ctx) (string, error) {
	sid := ctx.Req().Header.Get("Sid")
	if "" == sid {
		er := fmt.Errorf("验证Sid错误-SID为空")
		zlog.Error(er.Error(), er)
		return "", er
	}

	//获取本地sid
	// sidUser := util.GetCacheByCacheName(util.RPM_SID_USER_CACHE, sid)
	sidUser := util.RpmSession{}.GetDefultSesion().Get(sid)
	if nil == sidUser {
		er := fmt.Errorf("获取当前登录用户出错")
		zlog.Error(er.Error(), er)
		return "", er
	}
	return sidUser.(amData.SidUser).BranchCode, nil
}
