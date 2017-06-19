package dpService

import (
	"errors"

	"pccqcpa.com.cn/app/rpm/api/models/dp"
	"pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/components/zlog"
)

type DpReserveService struct{}

var dpReserve dp.DpReserve

// 分页操作
// by author Jason
// by time 2016-12-06 17:00:38
func (DpReserveService) List(param ...map[string]interface{}) (*util.PageData, error) {
	return dpReserve.List(param...)
}

// 多参数查询返回多条纪录
// by author Jason
// by time 2016-12-06 17:00:38
func (DpReserveService) Find(param ...map[string]interface{}) ([]*dp.DpReserve, error) {
	return dpReserve.Find(param...)
}

// 多参数查询返回一条纪录
// by author Jason
// by time 2016-12-06 17:00:38
func (this DpReserveService) FindOne(paramMap map[string]interface{}) (*dp.DpReserve, error) {
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
	er := errors.New("查询【RPM_PAR_DP_RESERVE】存款准备金有多条返回纪录")
	zlog.Error(er.Error(), er)
	return nil, er
}

// 新增纪录
// by author Jason
// by time 2016-12-06 17:00:38
func (DpReserveService) Add(model *dp.DpReserve) error {
	return model.Add()
}

// 更新纪录
// by author Jason
// by time 2016-12-06 17:00:38
func (DpReserveService) Update(model *dp.DpReserve) error {
	return model.Update()
}

// 删除纪录
// by author Jason
// by time 2016-12-06 17:00:38
func (DpReserveService) Delete(model *dp.DpReserve) error {
	return model.Delete()
}
