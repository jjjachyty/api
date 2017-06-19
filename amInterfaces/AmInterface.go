package amInterfaces

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/components/zlog"
)

type AmInterface struct {
	UserId  string
	Apitype string
	Sid     string
	Login   string
}

// 注销方法
func (am AmInterface) LoginOut() error {
	amUrl := util.GetIniStringValue("am", "url")
	url := amUrl + "login=1&userid=" + am.UserId
	resp, err := http.Get(url)
	defer resp.Body.Close()
	if nil != err {
		er := fmt.Errorf("注销用户失败")
		zlog.Error(er.Error(), err)
		return er
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	if nil != err {
		er := fmt.Errorf("获取注销用户接口数据出错")
		zlog.Error(er.Error(), err)
		return er
	}
	var rspData = new(util.RspData)
	err = json.Unmarshal(bytes, rspData)
	if nil != err {
		er := fmt.Errorf("注销用户接口：接口数据反序列化出错")
		zlog.Error(er.Error(), err)
		return er
	}
	zlog.Infof("请求注销接口 URL【%s】rspData【%#v】", nil, url, resp)
	if "101" != rspData.Status {
		er := fmt.Errorf(rspData.Message)
		zlog.Error(er.Error(), nil)
		return er
	}
	// util.DeleteCacheByCacheName(util.RPM_SID_USER_CACHE, am.Sid)
	util.RpmSession{}.GetDefultSesion().Del(am.Sid)

	return nil
}

// am系统调用退出方法
func (am AmInterface) LoginOutByAm() error {
	// util.DeleteCacheByCacheName(util.RPM_SID_USER_CACHE, am.Sid)
	util.RpmSession{}.GetDefultSesion().Del(am.Sid)
	return util.PutCacheByCacheName(util.RPM_SQUEEZE_CACHE, am.Sid, am.Sid, 0)
}
