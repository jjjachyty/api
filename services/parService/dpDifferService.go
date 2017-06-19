package parService

import (
	"errors"

	"pccqcpa.com.cn/app/rpm/api/models/par"

	"pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/components/zlog"
)

type DpDifferService struct{}

var dpDiffer par.DpDiffer

// 分页操作
// by author Jason
// by time 2016-11-30 09:31:42
func (DpDifferService) List(param ...map[string]interface{}) (*util.PageData, error) {
	if 0 < len(param) {
		paramMap := param[0]
		flag := util.CheckQueryParams(paramMap)
		if flag {
			return nil, errors.New("查询条件中不允许出现[^$.*+?{}()[]|'`~\\]特殊字符")
		}
	}
	return dpDiffer.List(param...)
}

// 多参数查询返回多条纪录
// by author Jason
// by time 2016-11-30 09:31:42
func (DpDifferService) Find(param ...map[string]interface{}) ([]*par.DpDiffer, error) {
	return dpDiffer.Find(param...)
}

// 多参数查询返回一条纪录
// by author Jason
// by time 2016-11-30 09:31:42
func (this DpDifferService) FindOne(paramMap map[string]interface{}) (*par.DpDiffer, error) {
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
	er := errors.New("查询【RPM_PAR_DP_DIFFER】存款差异化定价参数表有多条返回纪录")
	zlog.Error(er.Error(), er)
	return nil, er
}

// 新增纪录
// by author Jason
// by time 2016-11-30 09:31:42
func (DpDifferService) Add(model *par.DpDiffer) error {
	return model.Add()
}

// 更新纪录
// by author Jason
// by time 2016-11-30 09:31:42
func (DpDifferService) Update(model *par.DpDiffer) error {
	return model.Update()
}

// 删除纪录
// by author Jason
// by time 2016-11-30 09:31:42
func (DpDifferService) Delete(model *par.DpDiffer) error {
	return model.Delete()
}
