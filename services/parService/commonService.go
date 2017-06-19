package parService

import (
	"errors"

	"pccqcpa.com.cn/app/rpm/api/models/par"

	"pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/components/zlog"
)

type CommonService struct{}

var common par.Common

// 分页操作
// by author Yeqc
// by time 2016-12-20 10:58:09
func (CommonService) List(param ...map[string]interface{}) (*util.PageData, error) {
	return common.List(param...)
}

// 多参数查询返回多条纪录
// by author Yeqc
// by time 2016-12-20 10:58:09
func (CommonService) Find(param ...map[string]interface{}) ([]*par.Common, error) {
	return common.Find(param...)
}

// 多参数查询返回一条纪录
// by author Yeqc
// by time 2016-12-20 10:58:09
func (this CommonService) FindOne(paramMap map[string]interface{}) (*par.Common, error) {
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
	er := errors.New("查询【RPM_PAR_COMMON】通用参数设置有多条返回纪录")
	zlog.Error(er.Error(), er)
	return nil, er
}

// 新增纪录
// by author Yeqc
// by time 2016-12-20 10:58:09
func (CommonService) Add(model *par.Common) error {
	return model.Add()
}

// 更新纪录
// by author Yeqc
// by time 2016-12-20 10:58:09
func (CommonService) Update(model *par.Common) error {
	return model.Update()
}

// 删除纪录
// by author Yeqc
// by time 2016-12-20 10:58:09
func (CommonService) Delete(model *par.Common) error {
	return model.Delete()
}
