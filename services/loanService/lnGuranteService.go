package loanService

import (
	"fmt"

	"pccqcpa.com.cn/app/rpm/api/models/biz/ln"
	"pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/components/zlog"
)

var lnGuaranteModel ln.LnGuarante

//LnGuaranteService type 保证人服务类
type LnGuaranteService struct {
}

// List func LnGuaranteService 获取保证人列表服务方法
func (g LnGuaranteService) List(param ...map[string]interface{}) (*util.PageData, error) {
	return lnGuaranteModel.List(param...)
}

// Find func LnGuaranteService 获取保证人服务方法
func (g LnGuaranteService) Find(param ...map[string]interface{}) ([]*ln.LnGuarante, error) {
	return lnGuaranteModel.Find(param...)
}

// 不能重复添加同一个保证人
// 不能添加自己为保证人
func (g LnGuaranteService) Add(lnGuarante *ln.LnGuarante) error {
	paramMap := map[string]interface{}{
		"business_code": lnGuarante.BusinessCode,
		"cust":          lnGuarante.Guarante.CustCode,
	}
	lnbusiness, err := lnBusinessModel.Find(paramMap)
	if nil != err {
		er := fmt.Errorf("添加保证人时验证保证人是否为本人时出错")
		zlog.Error(er.Error(), err)
		return er
	}
	if 0 != len(lnbusiness) {
		er := fmt.Errorf("不能添加贷款客户本身为保证人")
		zlog.Error(er.Error(), er)
		return er
	}
	paramMap = map[string]interface{}{
		"business_code": lnGuarante.BusinessCode,
		"guarante":      lnGuarante.Guarante.CustCode,
	}

	lnGuarantes, _ := lnGuarante.Find(paramMap)
	if 0 < len(lnGuarantes) {
		er := fmt.Errorf("业务号为的[%v]的保证人编码为[%v]的信息已存在，不可再添加改保证人", lnGuarante.BusinessCode, lnGuarante.Guarante.CustCode)
		zlog.Error(er.Error(), er)
		return er
	}

	err = lnGuarante.Add()
	if nil != err {
		return err
	}

	lnGuarantes, _ = lnGuarante.Find(paramMap)
    if 1 == len(lnGuarantes) {
			lnGuarante = lnGuarantes[0]
	}

	return nil
}

func (g LnGuaranteService) Update(lnGuarante *ln.LnGuarante) error {
	paramMap := map[string]interface{}{
		"business_code": lnGuarante.BusinessCode,
		"cust":          lnGuarante.Guarante.CustCode,
	}
	lnbusiness, err := lnBusinessModel.Find(paramMap)
	if nil != err {
		er := fmt.Errorf("添加保证人时验证保证人是否为本人时出错")
		zlog.Error(er.Error(), err)
		return er
	}
	if 0 != len(lnbusiness) {
		er := fmt.Errorf("不能添加贷款客户本身为保证人")
		zlog.Error(er.Error(), er)
		return er
	}

	return lnGuarante.Update()
}

func (g LnGuaranteService) Delete(lnGuarante *ln.LnGuarante) error {
	return lnGuarante.Delete()
}

func (g LnGuaranteService) DeleteByBusinessCode(businessCode string) error {
	return lnGuaranteModel.DeleteByBusinessCode(businessCode)
}
