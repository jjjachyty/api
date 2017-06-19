package parService

import (
	"errors"

	"pccqcpa.com.cn/app/rpm/api/models/par"

	"pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/components/zlog"
)

type LgdBaselService struct{}

var lgdBasel par.LgdBasel

// 分页操作
// by author Yeqc
// by time 2016-12-19 10:16:29
func (LgdBaselService) List(param ...map[string]interface{}) (*util.PageData, error) {
	return lgdBasel.List(param...)
}

// 多参数查询返回多条纪录
// by author Yeqc
// by time 2016-12-19 10:16:29
func (LgdBaselService) Find(param ...map[string]interface{}) ([]*par.LgdBasel, error) {
	return lgdBasel.Find(param...)
}

// 多参数查询返回一条纪录
// by author Yeqc
// by time 2016-12-19 10:16:29
func (this LgdBaselService) FindOne(paramMap map[string]interface{}) (*par.LgdBasel, error) {
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
	er := errors.New("查询【RPM_PAR_LGD_BASEL】违约损失率信息有多条返回纪录")
	zlog.Error(er.Error(), er)
	return nil, er
}

// 新增纪录
// by author Yeqc
// by time 2016-12-19 10:16:29
func (this LgdBaselService) Add(model *par.LgdBasel) error {
	err := model.Add()
	if nil != err {
		return err
	}
	return model.Init()
}

// 更新纪录
// by author Yeqc
// by time 2016-12-19 10:16:29
func (LgdBaselService) Update(model *par.LgdBasel) error {
	err := model.Update()
	if nil != err {
		return err
	}
	return model.Init()
}

// 删除纪录
// by author Yeqc
// by time 2016-12-19 10:16:29
func (LgdBaselService) Delete(model *par.LgdBasel) error {
	err := model.Delete()
	if nil != err {
		return err
	}
	return model.Init()
}
