package ln

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/lunny/tango"

	"pccqcpa.com.cn/app/rpm/api/models/biz/ln"
	"pccqcpa.com.cn/app/rpm/api/services/loanService"
	"pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/app/rpm/api/util/currentMsg"
	"pccqcpa.com.cn/components/zlog"
)

var custInfoService loanService.CustInfoService

type CustInfoAction struct {
	tango.Json
	tango.Ctx
}

func (c *CustInfoAction) List() util.RstMsg {
	paramMap := util.GetPageMsg(&c.Ctx)
	flag := util.CheckQueryParams(paramMap)

	if flag {
		return util.ErrorMsg("查询条件中不允许出现[^$.*+?{}()[]|'`~\\]特殊字符", errors.New("查询条件中不允许出现[^$.*+?{}()[]|'`~\\]特殊字符"))
	}

	user, err := currentMsg.GetCurrentUser(c.Ctx)

	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	paramMap["owner"] = user
	pageData, err := custInfoService.List(paramMap)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessListMsg("分页查询客户信息成功", pageData)
}

func (c *CustInfoAction) Get() util.RstMsg {
	paramMap, err := util.GetParmFromRouter(&c.Ctx)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	if util.IsPaginQuery(&c.Ctx) {
		util.GetPageMsg(&c.Ctx, paramMap)
		flag := util.CheckQueryParams(paramMap)
		if flag {
			return util.ErrorMsg("查询条件中不允许出现[^$.*+?{}()[]|'`~\\]特殊字符", errors.New("查询条件中不允许出现[^$.*+?{}()[]|'`~\\]特殊字符"))
		}
		user, err := currentMsg.GetCurrentUser(c.Ctx)
		if nil != err {
			return util.ErrorMsg(err.Error(), err)
		}
		paramMap["owner"] = user

		pageData, err := custInfoService.List(paramMap)
		if nil != err {
			return util.ErrorMsg(err.Error(), err)
		}
		return util.SuccessListMsg("分页查询客户信息成功", pageData)
	}
	cust, err := custInfoService.FindOne(paramMap)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("查询单个客户信息成功", cust)
}

//新增客户信息
func (c *CustInfoAction) Post() util.RstMsg {
	var custInfo = new(ln.CustInfo)

	err := currentMsg.DecodeJson(c.Ctx, custInfo)

	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}

	user, err := currentMsg.GetCurrentUser(c.Ctx)

	if nil == err {
		custInfo.Owner = user.UserId
	}
	err = custInfoService.Add(custInfo)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("新增客户信息成功", custInfo)
}

// 删除模拟客户信息
func (this *CustInfoAction) Delete() util.RstMsg {
	var custInfo = new(ln.CustInfo)
	err := this.DecodeJson(&custInfo)
	if nil != err {
		er := fmt.Errorf("json转换实体出错")
		zlog.Error(er.Error(), err)
		return util.ErrorMsg(er.Error(), err)
	}
	data, err := custInfoService.Delete(custInfo)
	if nil != err && nil == data {
		return util.ErrorMsg(err.Error(), err)
	} else if nil != data {
		strs := strings.Split(data.(string), " ")
		rstCode, strconvErr := strconv.Atoi(strs[0])
		if nil != strconvErr {
			er := fmt.Errorf("字符串【%v】转换数字出错", strs[0])
			zlog.Error(er.Error(), strconvErr)
			return util.ErrorMsg(er.Error(), strconvErr)
		}
		return util.ReturnErrorRedirectMsg(rstCode, err.Error(), strs[1], err)
	} else {
		return util.SuccessMsg("删除客户信息成功", custInfo)
	}
}

func (this *CustInfoAction) Put() util.RstMsg {
	var custInfo = new(ln.CustInfo)
	err := currentMsg.DecodeJson(this.Ctx, custInfo)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	err = custInfoService.Update(custInfo)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("更新客户信息成功", custInfo)
}
