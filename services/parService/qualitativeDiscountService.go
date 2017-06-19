package parService

import (
	"errors"

	// "fmt"

	"pccqcpa.com.cn/app/rpm/api/models/par"

	"pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/components/zlog"
	// "platform/dbobj"
)

type QualitativeDiscountService struct{}

var qualitativeDiscount par.QualitativeDiscount

// 分页操作
// by author Yeqc
// by time 2016-12-21 10:20:03
func (QualitativeDiscountService) List(param ...map[string]interface{}) (*util.PageData, error) {
	return qualitativeDiscount.List(param...)
}

// 多参数查询返回多条纪录
// by author Yeqc
// by time 2016-12-21 10:20:03
func (QualitativeDiscountService) Find(param ...map[string]interface{}) ([]*par.QualitativeDiscount, error) {
	return qualitativeDiscount.Find(param...)
}

// 多参数查询返回一条纪录
// by author Yeqc
// by time 2016-12-21 10:20:03
func (this QualitativeDiscountService) FindOne(paramMap map[string]interface{}) (*par.QualitativeDiscount, error) {
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
	er := errors.New("查询【RPM_PAR_QUALITATIVE_DISCOUNT】定性优惠点数有多条返回纪录")
	zlog.Error(er.Error(), er)
	return nil, er
}

// 新增纪录
// by author Yeqc
// by time 2016-12-21 10:20:03
func (QualitativeDiscountService) Add(model *par.QualitativeDiscount) error {
	return model.Add()
}

// 更新纪录
// by author Yeqc
// by time 2016-12-21 10:20:03
func (QualitativeDiscountService) Update(model *par.QualitativeDiscount) error {
	return model.Update()
}

// 删除纪录
// by author Yeqc
// by time 2016-12-21 10:20:03
func (QualitativeDiscountService) Delete(model *par.QualitativeDiscount) error {
	return model.Delete()
}
