package common

import (
	"fmt"
	"github.com/lunny/tango"
	"strings"

	"pccqcpa.com.cn/app/rpm/api/services/commonService"
	"pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/components/zlog"
)

type CodeAction struct {
	tango.Json
	tango.Ctx
}

var codeService commonService.CodeService

func (c *CodeAction) Get() util.RstMsg {
	action := c.Param(":action")
	switch strings.ToLower(action) {
	case "lnbusiness":
		code := codeService.GetLnBusinessCode()
		return util.SuccessMsg("获取对公业务编码成功", code)
	case "cust":
		code := codeService.GetCustCodeByTime()
		return util.SuccessMsg("获取对公业务编码成功", code)
	case "onedp":
		code := codeService.GetOneDpBusinessByTime()
		return util.SuccessMsg("获取一对一存款业务编码成功", code)
	default:
		er := fmt.Errorf("不支持[%v]操作", action)
		zlog.Error(er.Error(), er)
		return util.ErrorMsg(er.Error(), er)
	}
	return util.ErrorMsg("未传路有参数", nil)
}
