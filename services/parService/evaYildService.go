package parService

import (
	"errors"

	"pccqcpa.com.cn/app/rpm/api/models/par"

	"pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/components/zlog"
)

type EvaYieldService struct{}

var evaYield par.EvaYield

// 分页操作
// by author Yeqc
// by time 2016-12-19 15:10:36
func (EvaYieldService) List(param ...map[string]interface{}) (*util.PageData, error) {
	return evaYield.List(param...)
}

// 多参数查询返回多条纪录
// by author Yeqc
// by time 2016-12-19 15:10:36
func (EvaYieldService) Find(param ...map[string]interface{}) ([]*par.EvaYield, error) {
	return evaYield.Find(param...)
}

// 多参数查询返回一条纪录
// by author Yeqc
// by time 2016-12-19 15:10:36
func (this EvaYieldService) FindOne(paramMap map[string]interface{}) (*par.EvaYield, error) {
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
	er := errors.New("查询【RPM_PAR_EVA_YIELD】EVA收益率表有多条返回纪录")
	zlog.Error(er.Error(), er)
	return nil, er
}

// 新增纪录
// by author Yeqc
// by time 2016-12-19 15:10:36
func (EvaYieldService) Add(model *par.EvaYield) error {
	return model.Add()
}

// 更新纪录
// by author Yeqc
// by time 2016-12-19 15:10:36
func (EvaYieldService) Update(model *par.EvaYield) error {
	return model.Update()
}

// 删除纪录
// by author Yeqc
// by time 2016-12-19 15:10:36
func (EvaYieldService) Delete(model *par.EvaYield) error {
	return model.Delete()
}
