package loanService

import (
	"fmt"

	"pccqcpa.com.cn/app/rpm/api/models/biz/ln"
	"pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/components/zlog"
)

var lnMortModel ln.LnMort

//LnMortService type 抵质押品服务类
type LnMortService struct {
}

// List func LnMortService 获取定价单列表服务方法
func (mort LnMortService) List(param ...map[string]interface{}) (*util.PageData, error) {
	return lnMortModel.List(param...)
}

// Find func LnMortService 获取定价单服务方法
func (mort LnMortService) Find(param ...map[string]interface{}) ([]*ln.LnMort, error) {
	return lnMortModel.Find(param...)
}

func (mort LnMortService) Add(lnMort *ln.LnMort) error {
	err := lnMort.Add()
	if nil != err {
		return err
	}
	paramMap := map[string]interface{}{
		"business_code": lnMort.BusinessCode,
		"mortgage_code": lnMort.MortgageCode,
	}
	lnMorts, _ := lnMort.Find(paramMap)
	if 1 < len(lnMorts) {
		er := fmt.Errorf("新增抵押品失败，业务号为[%v]的抵押品编码为[%v]的信息有多条", lnMort.BusinessCode, lnMort.MortgageCode)
		zlog.Error(er.Error(), nil)
		return er
	} else if 1 == len(lnMorts) {
		lnMort = lnMorts[0]
	}

	return nil
}

func (mort LnMortService) Update(lnMort *ln.LnMort) error {
	return lnMort.Update()
}

func (mort LnMortService) Delete(lnMort *ln.LnMort) error {
	return lnMort.Delete()
}

func (mort LnMortService) DeleteByBusinessCode(businessCode string) error {
	return lnMortModel.DeleteByBusinessCode(businessCode)
}
