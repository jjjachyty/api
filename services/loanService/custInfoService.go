package loanService

import (
	"errors"
	"fmt"
	"strings"

	"pccqcpa.com.cn/app/rpm/api/models/amData"
	"pccqcpa.com.cn/app/rpm/api/models/biz/ln"
	"pccqcpa.com.cn/app/rpm/api/models/dp"
	"pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/components/zlog"
)

var custInfoModel ln.CustInfo
var utilBase util.UtilBase

// var lnBusinessModel ln.LnBusiness

type CustInfoService struct {
}

// List func CustInfoService 分页查询客户信息
func (c CustInfoService) List(param ...map[string]interface{}) (*util.PageData, error) {
	var paramMap = make(map[string]interface{})
	if 0 < len(param) {
		// util.CopyMap(param[0], &paramMap)
		paramMap = utilBase.DeepCopy(param[0]).(map[string]interface{})
	}
	// 判断是否有非法字符
	flag := util.CheckQueryParams(paramMap)
	if flag {
		return nil, errors.New("查询条件中不允许出现[^$.*+?{}()[]|'`~\\]特殊字符")
	}
	var branchSearchLike []map[string]interface{}
	if v, ok := paramMap[util.SEARCH_LIKE]; ok {
		branchSearchLike = v.([]map[string]interface{})
	}
	// 模拟查询
	if custCode, ok := paramMap["cust_code"]; ok {
		key := []interface{}{
			"t.cust_code", "cust_name",
		}
		value := []interface{}{
			custCode, custCode,
		}
		var searchLike []map[string]interface{}
		if val, ok := paramMap[util.SEARCH_LIKE]; ok {
			searchLike = val.([]map[string]interface{})
		}
		searchLike = append(searchLike, map[string]interface{}{
			"type":  "or",
			"key":   key,
			"value": value,
		})

		paramMap[util.SEARCH_LIKE] = searchLike
		delete(paramMap, "cust_code")
		delete(paramMap, "cust_name")
	}

	if _, ok := paramMap[util.RPM_CUST_OWNER]; !ok {
		return custInfoModel.List(paramMap)
	}

	var currentUser = paramMap[util.RPM_CUST_OWNER].(amData.SidUser)
	var owner = currentUser.UserId
	delete(paramMap, util.RPM_CUST_OWNER)
	// 判断是否为客户经理
	if util.RPM_CM == currentUser.RoleIds {
		// 客户经理查询规则
		var custListCmParamMap = make(map[string]interface{})
		// err := util.CopyMap(paramMap, &custListCmParamMap)
		custListCmParamMap = utilBase.DeepCopy(paramMap).(map[string]interface{})
		// if nil != err {
		// 	return nil, err
		// }

		custListCmParamMap[util.RPM_CUST_OWNER] = owner
		return custInfoModel.List(custListCmParamMap)
	} else {
		// if status, ok := paramMap["status"]; ok {
		// 	switch status.(string) {
		// 	case util.REAL_CUST:
		// 		return c.findPageDataWithRealCust(paramMap)
		// 	case util.SIMULATION_CUST:
		// 		return c.findPageDataWithSimCust(paramMap, owner)
		// 	default:
		// 		er := fmt.Errorf("客户状态只有模拟与存量")
		// 		zlog.Error(er.Error(), er)
		// 		return nil, er
		// 	}
		// } else {
		// 	pageDataReal, err := c.findPageDataWithRealCust(paramMap)
		// 	if nil != err {
		// 		return nil, err
		// 	}
		// 	pageDataSim, err := c.findPageDataWithSimCust(paramMap, owner)
		// 	if nil != err {
		// 		return nil, err
		// 	}
		// 	return util.GetUtilBase().JoinPageData(c.callBack, pageDataReal, pageDataSim), nil
		// }

		// -- not cm
		// SELECT * FROM RPM_BIZ_CUST_INFO T
		//  WHERE (T.OWNER = 'rpmttt' and t.status = '01')
		//     or (T.BRANCH IN ('panchina001') and T.STATUS = '02');

		var custListMap = make(map[string]interface{})
		custListMap = utilBase.DeepCopy(paramMap).(map[string]interface{})
		c.handleCmParam(custListMap)
		return custInfoModel.UnionList(
			custListMap,
			map[string]interface{}{
				"T.OWNER":  owner,
				"T.STATUS": util.SIMULATION_CUST,
			},
			map[string]interface{}{
				"T.STATUS":       util.REAL_CUST,
				util.SEARCH_LIKE: branchSearchLike,
			})
	}
}

func (c CustInfoService) findPageDataWithRealCust(paramMap map[string]interface{}) (*util.PageData, error) {
	var custListMapReal = make(map[string]interface{})
	custListMapReal = utilBase.DeepCopy(paramMap).(map[string]interface{})
	custListMapReal["status"] = util.REAL_CUST
	return custInfoModel.List(custListMapReal)
}

func (c CustInfoService) findPageDataWithSimCust(paramMap map[string]interface{}, owner string) (*util.PageData, error) {
	var custListMapSim = make(map[string]interface{})
	custListMapSim = utilBase.DeepCopy(paramMap).(map[string]interface{})
	custListMapSim["status"] = util.SIMULATION_CUST
	custListMapSim[util.RPM_CUST_OWNER] = owner
	c.handleCmParam(custListMapSim)
	return custInfoModel.List(custListMapSim)
}

// Find func CustInfoService 多参数查询客户信息
func (CustInfoService) Find(param ...map[string]interface{}) ([]*ln.CustInfo, error) {
	if 0 < len(param) {
		if custCode, ok := param[0]["cust_code"]; ok {
			param[0]["t.cust_code"] = custCode
			delete(param[0], "cust_code")
		}
	}
	return custInfoModel.Find(param...)
}

func (c CustInfoService) FindOne(paramMap map[string]interface{}) (*ln.CustInfo, error) {
	custInfos, err := c.Find(paramMap)
	if nil != err {
		return nil, err
	}
	switch len(custInfos) {
	case 0:
		zlog.Info("查询客户信息为空", nil)
		return nil, nil
	case 1:
		return custInfos[0], nil
	}
	er := fmt.Errorf("查询客户信息有多条")
	zlog.Error(er.Error(), er)
	return nil, er
}

// Add func CustInfoService 新增客户信息
func (CustInfoService) Add(cust *ln.CustInfo) error {
	// 判断客户信息是否存在，然后再新增
	paramMap := map[string]interface{}{
		"cust_code": cust.CustCode,
	}
	custInfos, err := cust.Find(paramMap)
	if nil != err {
		er := fmt.Errorf("判断客户信息是否存在时报错")
		zlog.Error(er.Error(), err)
		return er
	}
	if 0 < len(custInfos) {
		er := fmt.Errorf("客户号为[%v]的客户信息已存在", cust.CustCode)
		zlog.Error(er.Error(), err)
		return er
	}
	err = cust.Add()
	if nil == err {
		custInfos, _ := cust.Find(paramMap)
		cust = custInfos[0]
	}
	return err
}

// Update func CustInfoService 更新客户信息
func (CustInfoService) Update(cust *ln.CustInfo) error {
	return cust.Update()
}

// Delete func CustInfoService 删除客户信息
// 判断是否有业务信息引用
// 第一个返回值，判断是否需要重定向
func (CustInfoService) Delete(cust *ln.CustInfo) (interface{}, error) {
	paramMap := map[string]interface{}{
		"cust": cust.CustCode,
	}
	lnbusiness, err := lnBusinessModel.Find(paramMap)
	if nil != err {
		return nil, err
	}
	if 0 < len(lnbusiness) {
		er := fmt.Errorf("有贷款业务单号[%v]引用该客户，不能删除", lnbusiness[0].BusinessCode)
		zlog.Error(er.Error(), er)
		return util.NUM_300 + cust.CustCode, er
	}

	dpOnePricings, err := dp.DpOnePricing{}.Find(paramMap)
	if nil != err {
		return nil, err
	}
	if 0 < len(dpOnePricings) {
		er := fmt.Errorf("有存款业务单号[%v]引用该客户，不能删除", dpOnePricings[0].BusinessCode)
		zlog.Error(er.Error(), er)
		return util.NUM_301 + cust.CustCode, er
	}
	return nil, cust.Delete()
}

// BUG(Jason)： #1： 未实现删除逻辑判断业务信息是否有引用客户信息

// 处理客户经理与非客户经理
// -- cm
// SELECT * FROM RPM_BIZ_CUST_INFO T
//  WHERE T.OWNER = 'rpmttt'
//    AND T.BRANCH IN ('panchina001');

// -- not cm
// SELECT * FROM RPM_BIZ_CUST_INFO T
//  WHERE (T.OWNER = 'rpmttt' and t.status = '01')
//     or (T.BRANCH IN ('panchina001') and T.STATUS = '02');
func (CustInfoService) handleCmParam(paramMap map[string]interface{}) {
	if v, ok := paramMap[util.SEARCH_LIKE]; ok {
		searchLike := v.([]map[string]interface{})
		for index, param := range searchLike {
			if "in" == strings.ToLower(param["type"].(string)) && ("t.branch" == strings.ToLower(param["key"].(string)) || "branch" == strings.ToLower(param["key"].(string))) {
				searchLike = append(searchLike[:index], searchLike[index+1:]...)
				paramMap[util.SEARCH_LIKE] = searchLike
				break
			}
		}
	}
}

func (c CustInfoService) callBack(slice1, slice2 interface{}) interface{} {
	return append(slice1.([]*ln.CustInfo), slice2.([]*ln.CustInfo)...)
}
