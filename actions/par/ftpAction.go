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

type FtpAction struct {
	tango.Ctx
	tango.Json
}

var ftpRateService parService.FtpRateService

// 分页查询
// by author Yeqc
// by time 2016-12-19 10:23:37
/**
 * @api {get} /ftp 分页查询资金成本率
 * @apiName ListFtp
 * @apiGroup Ftp
 *
 * @apiVersion 1.0.0
 *
 * @apiUse Ftp
 *
 */
func (this *FtpAction) List() util.RstMsg {
	paramMap := util.GetPageMsg(&this.Ctx)
	pageData, err := ftpRateService.List(paramMap)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessListMsg("分页查询信息成功", pageData)
}

// 多参数查询
// by author Yeqc
// by time 2016-12-19 10:21:15
/**
 * @api {get} /ftp/(*param) 多参数查询资金成本率
 * @apiName GetFtp
 * @apiGroup Ftp
 *
 * @apiVersion 1.0.0
 *
 * @apiUse Ftp
 *
 */
func (f *FtpAction) Get() util.RstMsg {
	paramMap, err := util.GetParmFromRouter(&f.Ctx)
	if nil != err {
		zlog.Error(err.Error(), err)
		return util.ErrorMsg(err.Error(), err)
	}
	if util.IsPaginQuery(&f.Ctx) {
		util.GetPageMsg(&f.Ctx, paramMap)
		pageData, err := ftpRateService.List(paramMap)
		if nil != err {
			return util.ErrorMsg(err.Error(), err)
		}
		return util.SuccessListMsg("分页查询信息成功", pageData)
	}
	//多参数查询
	ftps, err := ftpRateService.Find(paramMap)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("多参数查询指标成功", ftps)
}

// 新增信息
// by author Yeqc
// by time 2016-12-19 10:23:37
/**
 * @api {post} /ftp 新增资金成本率
 * @apiName PostFtp
 * @apiGroup Ftp
 *
 * @apiVersion 1.0.0
 *
 * @apiUse Ftp
 *
 */
func (this *FtpAction) Post() util.RstMsg {
	var one = new(par.Ftp)
	err := currentMsg.DecodeJson(this.Ctx, one)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	err = ftpRateService.Add(one)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("新增信息成功", one)
}

// 更新信息
// by author Yeqc
// by time 2016-12-19 10:23:37
/**
 * @api {put} /ftp 更新资金成本率
 * @apiName PutFtp
 * @apiGroup Ftp
 *
 * @apiVersion 1.0.0
 *
 * @apiUse Ftp
 *
 */
func (this *FtpAction) Put() util.RstMsg {
	var one = new(par.Ftp)
	err := currentMsg.DecodeJson(this.Ctx, one)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	err = ftpRateService.Update(one)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("更新信息成功", one)
}

// 删除信息
// by author Yeqc
// by time 2016-12-19 10:23:37
/**
 * @api {delete} /ftp 删除资金成本率
 * @apiName DeleteFtp
 * @apiGroup Ftp
 *
 * @apiVersion 1.0.0
 *
 * @apiUse Ftp
 *
 */
func (this *FtpAction) Delete() util.RstMsg {
	var one = new(par.Ftp)
	err := this.DecodeJson(one)
	if nil != err {
		er := fmt.Errorf("json转换实体出错")
		zlog.Error(er.Error(), err)
		return util.ErrorMsg(er.Error(), err)
	}
	err = ftpRateService.Delete(one)
	if nil != err {
		return util.ErrorMsg(err.Error(), err)
	}
	return util.SuccessMsg("删除信息成功", one)
}
