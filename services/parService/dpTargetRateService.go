package parService

import (
	"errors"

	"pccqcpa.com.cn/app/rpm/api/models/par"
	"pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/components/zlog"
)

type DpTargetRateService struct{}

var dpTargetRate par.DpTargetRate

// 分页操作
// by author Jason
// by time 2017-02-15 10:19:49
func (DpTargetRateService) List(param ...map[string]interface{}) (*util.PageData, error) {
	return dpTargetRate.List(param...)
}

// 多参数查询返回多条纪录
// by author Jason
// by time 2017-02-15 10:19:49
func (DpTargetRateService) Find(param ...map[string]interface{}) ([]*par.DpTargetRate, error) {
	return dpTargetRate.Find(param...)
}

// 多参数查询返回一条纪录
// by author Jason
// by time 2017-02-15 10:19:49
func (this DpTargetRateService) FindOne(paramMap map[string]interface{}) (*par.DpTargetRate, error) {
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
	er := errors.New("查询【RPM_PAR_DP_TARGET_RATE】目标收益率有多条返回纪录")
	zlog.Error(er.Error(), er)
	return nil, er
}

// 新增纪录
// by author Jason
// by time 2017-02-15 10:19:49
func (DpTargetRateService) Add(model *par.DpTargetRate) error {
	return model.Add()
}

// 更新纪录
// by author Jason
// by time 2017-02-15 10:19:49
func (DpTargetRateService) Update(model *par.DpTargetRate) error {
	return model.Update()
}

// 删除纪录
// by author Jason
// by time 2017-02-15 10:19:49
func (DpTargetRateService) Delete(model *par.DpTargetRate) error {
	return model.Delete()
}
