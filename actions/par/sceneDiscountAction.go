package par

import (
	"fmt"

	"github.com/lunny/tango"

	"pccqcpa.com.cn/app/rpm/api/models/par"

	"pccqcpa.com.cn/app/rpm/api/services/parService"

	"pccqcpa.com.cn/app/rpm/api/util"

	"pccqcpa.com.cn/app/rpm/api/util/currentMsg"
	"pccqcpa.com.cn/components/zlog"
)

var sceneDiscountService parService.SceneDiscountService

type SceneDiscountAction struct {
	tango.Json
	tango.Ctx
}

// 分页查询
// by author Jason
// by time 2017-02-17 10:25:04
func (this *SceneDiscountAction) List() util.RstMsg {
	paramMap := util.GetPageMsg(&this.Ctx)
	pageData, err := sceneDiscountService.List(paramMap)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessListMsg("分页查询信息成功", pageData)
}

// 多参数查询
// by author Jason
// by time 2017-02-17 10:25:04
func (this *SceneDiscountAction) Get() util.RstMsg {
	paramMap, err := util.GetParmFromRouter(&this.Ctx)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	if util.IsPaginQuery(&this.Ctx) {
		util.GetPageMsg(&this.Ctx, paramMap)
		pageData, err := sceneDiscountService.List(paramMap)
		if nil != err {
			return util.ErrorMsg(err.Error(), err)
		}
		return util.SuccessListMsg("分页查询信息成功", pageData)
	}

	rst, err := sceneDiscountService.Find(paramMap)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("查询信息成功", rst)
}

// 新增信息
// by author Jason
// by time 2017-02-17 10:25:04
func (this *SceneDiscountAction) Post() util.RstMsg {
	var one = new(par.SceneDiscount)
	err := currentMsg.DecodeJson(this.Ctx, one)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	err = sceneDiscountService.Add(one)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("新增信息成功", one)
}

// 更新信息
// by author Jason
// by time 2017-02-17 10:25:04
func (this *SceneDiscountAction) Put() util.RstMsg {
	var one = new(par.SceneDiscount)
	err := currentMsg.DecodeJson(this.Ctx, one)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	err = sceneDiscountService.Update(one)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("更新信息成功", one)
}

// 删除信息
// by author Jason
// by time 2017-02-17 10:25:04
func (this *SceneDiscountAction) Delete() util.RstMsg {
	var one = new(par.SceneDiscount)
	err := this.DecodeJson(one)
	if nil != err {
		er := fmt.Errorf("json转换实体出错")
		zlog.Error(er.Error(), err)
		return util.ErrorMsg(er.Error(), err)
	}
	err = sceneDiscountService.Delete(one)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("删除信息成功", one)
}
