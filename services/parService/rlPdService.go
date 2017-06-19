package parService

import (
	"errors"

	// "fmt"

	"pccqcpa.com.cn/app/rpm/api/models/par"

	"pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/components/zlog"
	// "platform/dbobj"
)

type RlPdService struct{}

var rlPd par.RlPd

// 分页操作
// by author Yeqc
// by time 2016-12-30 11:03:34
func (RlPdService) List(param ...map[string]interface{}) (*util.PageData, error) {
	if 0 < len(param) {
		paramMap := param[0]
		flag := util.CheckQueryParams(paramMap)
		if flag {
			return nil, errors.New("查询条件中不允许出现[^$.*+?{}()[]|'`~\\]特殊字符")
		}

	}
	return rlPd.List(param...)
}

// 多参数查询返回多条纪录
// by author Yeqc
// by time 2016-12-30 11:03:34
func (RlPdService) Find(param ...map[string]interface{}) ([]*par.RlPd, error) {
	return rlPd.Find(param...)
}

// 多参数查询返回一条纪录
// by author Yeqc
// by time 2016-12-30 11:03:34
func (this RlPdService) FindOne(paramMap map[string]interface{}) (*par.RlPd, error) {
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
	er := errors.New("查询【RPM_PAR_RL_PD】零售违约概率有多条返回纪录")
	zlog.Error(er.Error(), er)
	return nil, er
}

// 新增纪录
// by author Yeqc
// by time 2016-12-30 11:03:34
func (RlPdService) Add(model *par.RlPd) error {
	model.Pd = model.BreachAcount / model.TotalAcount
	return model.Add()
}

// 更新纪录
// by author Yeqc
// by time 2016-12-30 11:03:34
func (RlPdService) Update(model *par.RlPd) error {
	model.Pd = model.BreachAcount / model.TotalAcount
	return model.Update()
}

// 删除纪录
// by author Yeqc
// by time 2016-12-30 11:03:34
func (RlPdService) Delete(model *par.RlPd) error {
	return model.Delete()
}
