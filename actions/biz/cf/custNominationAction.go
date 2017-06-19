package cf

//客户分类 白名单 控制器
//@auth by Janly
import (
	"errors"

	"github.com/lunny/tango"
	"pccqcpa.com.cn/app/rpm/api/models/biz/cf"
	"pccqcpa.com.cn/app/rpm/api/services/classfyService"
	"pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/components/zlog"
)

type CustNominationAction struct {
	tango.Json
	tango.Ctx
}

var custNominationService classfyService.CustNominationService

// excel 导入
// by author Jason
// by time 2016-10-31 15:34:54
func (this CustNominationAction) ExcelImport() util.RstMsg {
	var models []cf.CustNomination
	var excelName string = "xlsTemp/custNomination.xlsx"

	err := this.SaveToFile("custNomination", excelName)

	if nil != err {
		return util.ErrorMsg("导入失败", err)
	}

	err = util.GetValueFormExcel(&models, excelName)

	if nil != err {
		return util.ErrorMsg("导入失败", err)
	}

	rst, err := custNominationService.BatchAdd(models)

	if nil != err {
		return util.ErrorMsg("导入失败", err)
	}
	return util.SuccessMsg("导入成功", rst)
}

func (cna CustNominationAction) List() util.RstMsg {

	zlog.AppOperateLog("", "CustNominationAction.List", zlog.SELECT, nil, nil, "查询客户分类白名单")

	startRowNumber, pageSize, orderAttr, orderType, err := util.GetPageAndOrder(cna.Req().Header)

	params, err := util.GetParmFromRouter(&cna.Ctx)

	flag := util.CheckQueryParams(params)

	if flag {
		return util.ErrorMsg("查询条件中不允许出现[^$.*+?{}()[]|'`~\\]特殊字符", errors.New("查询条件中不允许出现[^$.*+?{}()[]|'`~\\]特殊字符"))
	}
	if err == nil {
		pageData, err := custNominationService.GetAll(startRowNumber, pageSize, orderAttr, orderType, params)
		if err == nil {
			return util.ReturnSuccess("获取客户分类白名单成功", pageData)
		}
	}

	return util.ErrorMsg("获取客户分类白名单失败", err)
}

func (cna CustNominationAction) Post() util.RstMsg {

	zlog.AppOperateLog("", "CustNominationAction.Post", zlog.ADD, nil, nil, "新增客户分类[白名单]")
	custNomination, err := cna.getCustNomination()
	if nil == err {

		err = custNominationService.Add(custNomination)
		if nil == err {
			return util.SuccessMsg("新增客户分类[白名单]成功", custNomination)
		}
	}
	// return util.ErrorMsg("新增客户分类[白名单]失败", err)
	return util.ErrorMsg(err.Error(), err)
}

func (cna CustNominationAction) Delete() util.RstMsg {
	custNomination, err := cna.getCustNomination()
	if nil == err {

		err := custNominationService.Remove(custNomination)
		if nil == err {
			return util.SuccessMsg("删除客户["+custNomination.CustName+"]成功", custNomination)
		}
	}
	return util.ErrorMsg("删除客户["+custNomination.CustName+"]失败", err)
}

//PATCH func 局部更新
func (cna CustNominationAction) Patch() util.RstMsg {
	custNomination, err := cna.getCustNomination()
	if nil == err {

		err := custNominationService.Update(custNomination)
		if nil == err {
			return util.SuccessMsg("更新客户["+custNomination.CustName+"]成功", custNomination)
		}
	}
	return util.ErrorMsg("更新客户["+custNomination.CustName+"]失败", err)
}

func (cna CustNominationAction) getCustNomination() (cf.CustNomination, error) {
	var custNomination cf.CustNomination
	errStr := "客户分类[白名单]json转换实体出错"
	err := cna.DecodeJson(&custNomination)
	if nil != err {
		zlog.Error(errStr, err)
	}
	return custNomination, err
}
