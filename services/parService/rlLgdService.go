package parService

import (
	"errors"

	// "fmt"

	"pccqcpa.com.cn/app/rpm/api/models/par"

	"pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/components/zlog"
	// "platform/dbobj"
)

type RlLgdService struct{}

var rlLgd par.RlLgd

// 分页操作
// by author Yeqc
// by time 2016-12-30 10:06:36
func (RlLgdService) List(param ...map[string]interface{}) (*util.PageData, error) {
	if 0 < len(param) {
		paramMap := param[0]
		flag := util.CheckQueryParams(paramMap)
		if flag {
			return nil, errors.New("查询条件中不允许出现[^$.*+?{}()[]|'`~\\]特殊字符")
		}

	}
	return rlLgd.List(param...)
}

// 多参数查询返回多条纪录
// by author Yeqc
// by time 2016-12-30 10:06:36
func (RlLgdService) Find(param ...map[string]interface{}) ([]*par.RlLgd, error) {
	return rlLgd.Find(param...)
}

// 多参数查询返回一条纪录
// by author Yeqc
// by time 2016-12-30 10:06:36
func (this RlLgdService) FindOne(paramMap map[string]interface{}) (*par.RlLgd, error) {
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
// by time 2016-12-30 10:06:36
func (RlLgdService) Add(model *par.RlLgd) error {
	return model.Add()
}

// 更新纪录
// by author Yeqc
// by time 2016-12-30 10:06:36
func (RlLgdService) Update(model *par.RlLgd) error {
	return model.Update()
}

// 删除纪录
// by author Yeqc
// by time 2016-12-30 10:06:36
func (RlLgdService) Delete(model *par.RlLgd) error {
	return model.Delete()
}
