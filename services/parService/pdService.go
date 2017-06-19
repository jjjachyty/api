package parService

import (
	"errors"

	// "fmt"

	"pccqcpa.com.cn/app/rpm/api/models/par"

	"pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/components/zlog"
	// "platform/dbobj"
)

type PdService struct{}

var pd par.Pd

// 分页操作
// by author Yeqc
// by time 2016-12-21 09:52:26
func (PdService) List(param ...map[string]interface{}) (*util.PageData, error) {
	return pd.List(param...)
}

// 多参数查询返回多条纪录
// by author Yeqc
// by time 2016-12-21 09:52:26
func (PdService) Find(param ...map[string]interface{}) ([]*par.Pd, error) {
	return pd.Find(param...)
}

// 多参数查询返回一条纪录
// by author Yeqc
// by time 2016-12-21 09:52:26
func (this PdService) FindOne(paramMap map[string]interface{}) (*par.Pd, error) {
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
	er := errors.New("查询有多条返回纪录")
	zlog.Error(er.Error(), er)
	return nil, er
}

// 新增纪录
// by author Yeqc
// by time 2016-12-21 09:52:26
func (PdService) Add(model *par.Pd) error {
	return model.Add()
}

// 更新纪录
// by author Yeqc
// by time 2016-12-21 09:52:26
func (PdService) Update(model *par.Pd) error {
	return model.Update()
}

// 删除纪录
// by author Yeqc
// by time 2016-12-21 09:52:26
func (PdService) Delete(model *par.Pd) error {
	return model.Delete()
}
