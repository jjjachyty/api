package loan

import (
	"fmt"

	"github.com/lunny/tango"

	"pccqcpa.com.cn/app/rpm/api/models/biz/ln"
	"pccqcpa.com.cn/app/rpm/api/services/loanService"
	"pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/app/rpm/api/util/currentMsg"
	"pccqcpa.com.cn/components/zlog"
)

var sdps loanService.SceneDpService

type SceneDpAction struct {
	tango.Ctx
	tango.Json
}

// url : /api/rpm/scenedp/(*param)
func (s *SceneDpAction) Get() util.RstMsg {

	paramMap, err := util.GetParmFromRouter(&s.Ctx)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}

	sceneDps, err := sdps.Find(paramMap)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}

	return util.SuccessMsg("查询存款派生业务成功", sceneDps)
}

func (s *SceneDpAction) Post() util.RstMsg {
	var sdp = new(ln.SceneDp)
	err := currentMsg.DecodeJson(s.Ctx, sdp)
	if nil != err {
		er := fmt.Errorf("json数据转换为存款派生结构体数据出错")
		zlog.Error(er.Error(), er)
		return util.ErrorMsg(er.Error(), err)
	}
	err = sdps.Add(sdp)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("新增存款派生信息成功", nil)
}

func (s *SceneDpAction) Put() util.RstMsg {
	var sdp = new(ln.SceneDp)
	err := currentMsg.DecodeJson(s.Ctx, sdp)
	if nil != err {
		er := fmt.Errorf("json数据转换为存款派生结构体数据出错")
		zlog.Error(er.Error(), er)
		return util.ErrorMsg(er.Error(), err)
	}
	err = sdps.Update(sdp)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("修改存款派生信息成功", nil)
}

func (s *SceneDpAction) Delete() util.RstMsg {
	var sdp = new(ln.SceneDp)
	err := s.DecodeJson(&sdp)
	if nil != err {
		er := fmt.Errorf("json数据转换为存款派生结构体数据出错")
		zlog.Error(er.Error(), er)
		return util.ErrorMsg(er.Error(), err)
	}
	err = sdps.Delete(sdp)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("删除存款派生信息成功", nil)
}
