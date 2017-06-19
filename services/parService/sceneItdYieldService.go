package parService

import (
	"errors"

	// "fmt"

	"pccqcpa.com.cn/app/rpm/api/models/par"

	"pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/components/zlog"
	// "platform/dbobj"
)

var sceneItdYieldModel par.SceneItdYield

type SceneItdYieldService struct{}

// 分页操作
// by author Yeqc
// by time 2016-12-21 10:30:46
func (SceneItdYieldService) List(param ...map[string]interface{}) (*util.PageData, error) {
	return sceneItdYieldModel.List(param...)
}

func (SceneItdYieldService) Find(param ...map[string]interface{}) ([]*par.SceneItdYield, error) {
	return sceneItdYieldModel.Find(param...)
}

func (s SceneItdYieldService) FindOne(param ...map[string]interface{}) (*par.SceneItdYield, error) {
	sceneItdYields, err := s.Find(param...)
	if nil != err {
		return nil, err
	} else if nil == sceneItdYields || 0 == len(sceneItdYields) {
		return nil, nil
	} else if 1 == len(sceneItdYields) {
		return sceneItdYields[0], nil
	} else {
		er := errors.New("查询派生中间收益率有多条")
		zlog.Error(er.Error(), er)
		return sceneItdYields[0], er
	}
}

// 新增纪录
// by author Yeqc
// by time 2016-12-21 10:30:46
func (SceneItdYieldService) Add(model *par.SceneItdYield) error {
	return model.Add()
}

// 更新纪录
// by author Yeqc
// by time 2016-12-21 10:30:46
func (SceneItdYieldService) Update(model *par.SceneItdYield) error {
	return model.Update()
}

// 删除纪录
// by author Yeqc
// by time 2016-12-21 10:30:46
func (SceneItdYieldService) Delete(model *par.SceneItdYield) error {
	return model.Delete()
}
