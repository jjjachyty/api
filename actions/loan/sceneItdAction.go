package loan

import (
	"github.com/lunny/tango"

	"pccqcpa.com.cn/app/rpm/api/models/biz/ln"
	"pccqcpa.com.cn/app/rpm/api/services/loanService"
	"pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/app/rpm/api/util/currentMsg"
)

var sids loanService.SceneItdService

type SceneItdAction struct {
	tango.Ctx
	tango.Json
}

// url : /api/rpm/scenedp/(*param)
func (s *SceneItdAction) Get() util.RstMsg {
	paramMap, err := util.GetParmFromRouter(&s.Ctx)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	sceneItds, err := sids.Find(paramMap)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}

	return util.SuccessMsg("查询中间派生业务成功", sceneItds)
}

func (s *SceneItdAction) Post() util.RstMsg {
	var sdp = new(ln.SceneItd)
	err := currentMsg.DecodeJson(s.Ctx, sdp)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	err = sids.Add(sdp)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("新增中间派生信息成功", nil)
}

func (s *SceneItdAction) Put() util.RstMsg {
	var sdp = new(ln.SceneItd)
	err := currentMsg.DecodeJson(s.Ctx, sdp)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	err = sids.Update(sdp)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("修改中间派生信息成功", nil)
}

func (s *SceneItdAction) Delete() util.RstMsg {
	var sdp = new(ln.SceneItd)
	s.DecodeJson(&sdp)
	err := sids.Delete(sdp)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("删除中间派生信息成功", nil)
}
