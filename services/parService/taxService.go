package parService

import (
	"errors"

	// "fmt"

	"pccqcpa.com.cn/app/rpm/api/models/par"

	"pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/components/zlog"
	// "platform/dbobj"
)

type TaxService struct{}

var taxModel par.Tax
var tax par.Tax

func (this *TaxService) SelectEcByParmas(paramMap map[string]interface{}) ([]*par.Tax, error) {
	return taxModel.SelectTaxByParams(paramMap)
}

// 分页操作
// by author Yeqc
// by time 2016-12-21 11:10:11
func (TaxService) List(param ...map[string]interface{}) (*util.PageData, error) {
	return tax.List(param...)
}

// 多参数查询返回多条纪录
// by author Yeqc
// by time 2016-12-21 11:10:11
func (TaxService) Find(param ...map[string]interface{}) ([]*par.Tax, error) {
	return tax.Find(param...)
}

// 多参数查询返回一条纪录
// by author Yeqc
// by time 2016-12-21 11:10:11
func (this TaxService) FindOne(paramMap map[string]interface{}) (*par.Tax, error) {
	models, err := this.Find(paramMap)
	if nil != err {
		return nil, err
	}
	switch len(models) {
	case 0:
		return nil, nil
	case 1:
		return models[0], nil
	}
	er := errors.New("查询【RPM_PAR_TAX】税率有多条返回纪录")
	zlog.Error(er.Error(), er)
	return nil, er
}

// 新增纪录
// by author Yeqc
// by time 2016-12-21 11:10:11
func (TaxService) Add(model *par.Tax) error {
	return model.Add()
}

// 更新纪录
// by author Yeqc
// by time 2016-12-21 11:10:11
func (TaxService) Update(model *par.Tax) error {
	return model.Update()
}

// 删除纪录
// by author Yeqc
// by time 2016-12-21 11:10:11
func (TaxService) Delete(model *par.Tax) error {
	return model.Delete()
}
