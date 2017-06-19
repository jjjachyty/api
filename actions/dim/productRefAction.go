package dim
import (
	"fmt"

	"github.com/lunny/tango"

	"pccqcpa.com.cn/app/rpm/api/models/dim"

	"pccqcpa.com.cn/app/rpm/api/services/dimService"


	"pccqcpa.com.cn/app/rpm/api/util"

	"pccqcpa.com.cn/app/rpm/api/util/currentMsg"
	"pccqcpa.com.cn/components/zlog"
)

var productRefService dimService.ProductRefService

type ProductRefAction struct {
	tango.Json
	tango.Ctx
}

// 分页查询
// by author Jason
// by time 2017-03-30 13:51:39
func(this *ProductRefAction) List() util.RstMsg {
	paramMap := util.GetPageMsg(&this.Ctx)
	pageData, err := productRefService.List(paramMap)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessListMsg("分页查询信息成功", pageData)
}

// 多参数查询
// by author Jason
// by time 2017-03-30 13:51:39
func(this *ProductRefAction) Get() util.RstMsg {
	paramMap, err := util.GetParmFromRouter(&this.Ctx)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	if util.IsPaginQuery(&this.Ctx) {
		util.GetPageMsg(&this.Ctx, paramMap)
		pageData,err := productRefService.List(paramMap)
		if nil != err {
			return util.ErrorMsg(err.Error(), err)
		}
		return util.SuccessListMsg("分页查询信息成功", pageData)
	}

	rst, err := productRefService.Find(paramMap)
	if nil != err {
			return util.ErrorMsg(err.Error(), err)
		}
	return util.SuccessMsg("查询信息成功", rst)
}

// 新增信息
// by author Jason
// by time 2017-03-30 13:51:39
func(this *ProductRefAction) Post() util.RstMsg {
	var one = new(dim.ProductRef)
	err := currentMsg.DecodeJson(this.Ctx, one)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	err = productRefService.Add(one)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("新增信息成功",one)
}

// 更新信息
// by author Jason
// by time 2017-03-30 13:51:39
func(this *ProductRefAction) Put() util.RstMsg {
	var one = new(dim.ProductRef)
	err := currentMsg.DecodeJson(this.Ctx, one)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	err = productRefService.Update(one)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("更新信息成功",one)
}

// 删除信息
// by author Jason
// by time 2017-03-30 13:51:39
func(this *ProductRefAction) Delete() util.RstMsg {
	var one = new(dim.ProductRef)
	err := this.DecodeJson(one)
	if nil != err {
		er := fmt.Errorf("json转换实体出错")
		zlog.Error(er.Error(), err)
		return util.ErrorMsg(er.Error(), err)
	}
	err = productRefService.Delete(one)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("删除信息成功",one)
}

