package classfyService

import (
	"database/sql"
	"errors"
	"fmt"
	"pccqcpa.com.cn/app/rpm/api/models/biz/cf"
	"pccqcpa.com.cn/app/rpm/api/services/loanService"
	"pccqcpa.com.cn/app/rpm/api/util"

	"pccqcpa.com.cn/components/zlog"
)

//客户分类 特殊客户名单 服务类
//@auth by Janly

type CustNominationService struct {
}

var custNominationModel cf.CustNominationModel

var custInfoService loanService.CustInfoService

var classifyResultService ClassifyResultService

func (CustNominationService) GetAll(startRowNumber int64, pageSize int64, orderAttr string, orderType util.OrderType, params map[string]interface{}) (util.PageData, error) {
	var page util.Page
	var order util.Order
	var pageData util.PageData
	page.StartRowNumber = startRowNumber
	page.PageSize = pageSize
	order.OrderAttr = orderAttr
	order.OrderType = orderType
	page, custNominations, err := custNominationModel.FindAllByPage(page, order, params)
	if nil == err {
		pageData.Rows = custNominations
		pageData.Page = page
	}
	return pageData, err
}

// 查询单一纪录
// by author Jason
// by time 2016-10-31 15:34:54
func (c CustNominationService) FindOne(paramMap map[string]interface{}) (*cf.CustNomination, error) {
	custNominations, err := custNominationModel.Find(paramMap)
	if nil != err {
		return nil, err
	}
	if 0 == len(custNominations) {
		return nil, nil
	} else if 1 != len(custNominations) {
		er := errors.New("查询客户白名单有多条纪录")
		zlog.Error(er.Error(), er)
		return nil, er
	}
	return custNominations[0], nil
}

func (c CustNominationService) Add(custNomination cf.CustNomination) error {
	paramMap := map[string]interface{}{
		"cust_code": custNomination.CustCode,
		"sort":      "cust_code",
	}
	one, err := c.FindOne(paramMap)
	if nil != err {
		return err
	}
	if nil != one {
		er := fmt.Errorf("已存在【%v】客户号的客户名单\n", custNomination.CustCode)
		zlog.Error(er.Error(), er)
		return er
	}
	err = custNominationModel.Save(custNomination)
	if nil != err {
		return err
	}
	return c.cascadeAddOrUpdate(custNomination)
}

func (c CustNominationService) Remove(custNomination cf.CustNomination) error {
	err := c.casecadeUpdate(custNomination)
	if nil != err {
		return err
	}
	return custNominationModel.Delete(custNomination)
}

func (c CustNominationService) Update(custNomination cf.CustNomination) error {
	err := custNominationModel.Update(custNomination)
	if nil != err {
		return err
	}
	return c.cascadeAddOrUpdate(custNomination)
}

// 新增纪录带事务
// by author Jason
// by time 2016-10-31 15:34:54
func (CustNominationService) BatchAdd(models []cf.CustNomination) (sql.Result, error) {
	return custNominationModel.BatchAdd(models)
}

// 删除白名单时更新客户分类结果表RPM_CLASSIFY_RESULT
// by author Jason
// by time 2016-10-31 15:34:54
func (c CustNominationService) casecadeUpdate(model cf.CustNomination) error {
	paramMap := map[string]interface{}{
		"cust_id": model.CustCode,
		"sort":    "cust_id",
	}
	one, err := classifyResultService.FindOne(paramMap)
	if nil != err {
		return err
	} else if one != nil {
		one.FinalCustClassification = ""
		err = classifyResultService.Update(*one)
	}
	return nil
}

// 级连更新或者新增客户分类目标表RPM_CLASSIFY_RESULT，与客户表RPM_BIZ_CUST_INFO
// by author Jason
// by time 2016-11-07 15:34:54
func (CustNominationService) cascadeAddOrUpdate(model cf.CustNomination) error {
	paramMap := map[string]interface{}{
		"cust_code": model.CustCode,
		"sort":      "cust_code",
	}
	custInfo, err := custInfoService.FindOne(paramMap)
	if nil != err {
		return err
	}
	if nil != custInfo {
		//修改客户等级
		custInfo.CustImplvl = model.CustImpl
		err := custInfoService.Update(custInfo)
		if nil != err {
			return err
		}
	}
	paramMap = map[string]interface{}{
		"cust_id": model.CustCode,
		"sort":    "cust_id",
	}
	one, err := classifyResultService.FindOne(paramMap)
	if nil != err {
		return err
	} else if one != nil {
		one.FinalCustClassification = model.CustImpl
		err = classifyResultService.Update(*one)
	} else {
		one = new(cf.ClassifyResult)
		one.FinalCustClassification = model.CustImpl
		one.CustId = model.CustCode
		err = classifyResultService.Add(*one)
	}
	return err
}
