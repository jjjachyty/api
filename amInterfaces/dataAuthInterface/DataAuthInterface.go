package dataAuthInterface

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/components/zlog"
)

type DataAuthResp struct {
	Sid    string     `json:"-"`
	UserId string     `json:"-"`
	Msg    string     `json:"msg"`
	Code   string     `json:"code"`
	Data   []DataAuth `json:"data"`
}

type DataAuth struct {
	DomainId         string `json:"domain_id"`         // #域编码
	DomainName       string `json:"domain_name"`       // #域名称
	OrgUnitId        string `json:"org_unit_id"`       // #机构编码
	OrgUnitDesc      string `json:"org_unit_desc"`     // #机构名称
	UserId           string `json:"user_id"`           // #用户编码
	UserName         string `json:"user_name"`         // #用户名称
	GroupId          string `json:"group_id"`          // #权限组编码
	ReqUrl           string `json:"req_url"`           // #URL
	ConditionType    string `json:"condition_type"`    // #条件类型    （1：机构）
	ConditionContent string `json:"condition_content"` // #条件值列表  （值用逗号分隔
	ContentDesc      string `json:"content_desc"`      // #条件值描述  （机构描述
}

// 初始化AM系统的【数据权限】数据
// key: UserId
// val: map[reqUrl][]DataAuth
func (dataAuthResp *DataAuthResp) Init() error {
	amUrl := util.GetIniStringValue("am", "url")
	var url string = amUrl + "login=0&apitype=2" + "&userid=" + dataAuthResp.UserId + "&sid=" + dataAuthResp.Sid
	zlog.Infof("访问AM系统数据权限URL：", nil, url)
	resp, err := http.Get(url)
	defer resp.Body.Close()
	if nil != err {
		er := fmt.Errorf("访问AM系统获取数据权限接口出错")
		zlog.Error(er.Error(), err)
		return er
	}
	body, err := ioutil.ReadAll(resp.Body)
	if nil != err {
		er := fmt.Errorf("读取AM系统数据接口数据出错")
		zlog.Error(er.Error(), err)
		return er
	}
	err = json.Unmarshal(body, dataAuthResp)
	if nil != err {
		er := fmt.Errorf("反序列化AM系统数据接口出错")
		zlog.Error(er.Error(), err)
		return er
	}
	if "400" == dataAuthResp.Code {
		er := fmt.Errorf(dataAuthResp.Msg)
		zlog.Error(er.Error(), er)
		return er
	}
	return dataAuthResp.initCache()
}

func (d DataAuthResp) initCache() error {
	// util.FlushCacheByCacheName(util.AM_DATA_AUTH_CACHE)
	util.DeleteCacheByCacheName(util.AM_DATA_AUTH_CACHE, d.UserId)
	for _, dataAuth := range d.Data {
		if "" == strings.TrimSpace(dataAuth.ConditionContent) {
			continue
		}
		dataAuthMaps := util.GetCacheByCacheName(util.AM_DATA_AUTH_CACHE, dataAuth.UserId)
		if nil == dataAuthMaps {
			dataAuthMap := make(map[string][]DataAuth)
			var dataAuthArray = make([]DataAuth, 0)
			dataAuthArray = append(dataAuthArray, dataAuth)
			dataAuthMap[dataAuth.ReqUrl] = dataAuthArray
			err := util.PutCacheByCacheName(util.AM_DATA_AUTH_CACHE, dataAuth.UserId, dataAuthMap, 0)
			if nil != err {
				er := fmt.Errorf("数据权限写入缓存出错key:【%v】val:【%#v】", dataAuth.UserId, dataAuthMap)
				zlog.Error(er.Error(), err)
				return er
			}
		} else {
			dataAuths, ok := dataAuthMaps.(map[string][]DataAuth)[dataAuth.ReqUrl]
			if !ok {
				dataAuths = make([]DataAuth, 0)
			}
			dataAuths = append(dataAuths, dataAuth)
			dataAuthMaps.(map[string][]DataAuth)[dataAuth.ReqUrl] = dataAuths
			err := util.PutCacheByCacheName(util.AM_DATA_AUTH_CACHE, dataAuth.UserId, dataAuthMaps, 0)
			if nil != err {
				er := fmt.Errorf("数据权限写入缓存出错key:【%v】val:【%#v】", dataAuth.UserId, dataAuthMaps)
				zlog.Error(er.Error(), err)
				return er
			}
		}
		// fmt.Println("缓存内容【%#v】", dataAuthMaps)
	}
	zlog.Info("AM系统数据权限接口数据写入缓存成功", nil)
	return nil
}

// func init() {
// 	(&DataAuthResp{}).Init()
// }
