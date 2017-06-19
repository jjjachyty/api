package parService

import (
	"errors"

	"pccqcpa.com.cn/app/rpm/api/models/par"

	"pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/components/zlog"
	// "platform/dbobj"
)

type RocService struct{}

var roc par.Roc

// 分页操作
// by author Yeqc
// by time 2016-12-21 13:39:33
func (RocService) List(param ...map[string]interface{}) (*util.PageData, error) {
	if 0 < len(param) {
		paramMap := param[0]
		flag := util.CheckQueryParams(paramMap)
		if flag {
			return nil, errors.New("查询条件中不允许出现[^$.*+?{}()[]|'`~\\]特殊字符")
		}
		if organName, ok := paramMap["organ_name"]; ok {
			key := []interface{}{
				"organ.organ_name", "product.product_name",
			}
			value := []interface{}{
				organName, organName,
			}
			searchLike := []map[string]interface{}{
				map[string]interface{}{
					"type":  "or",
					"key":   key,
					"value": value,
				},
			}
			paramMap["searchLike"] = searchLike
			delete(paramMap, "organ_name")
		}
	}
	return roc.List(param...)
}

// 多参数查询返回多条纪录
// by author Yeqc
// by time 2016-12-21 13:39:33
func (RocService) Find(param ...map[string]interface{}) ([]*par.Roc, error) {
	return roc.Find(param...)
}

// 多参数查询返回一条纪录
// by author Yeqc
// by time 2016-12-21 13:39:33
func (this RocService) FindOne(paramMap map[string]interface{}) (*par.Roc, error) {
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
	er := errors.New("查询【RPM_PAR_ROC】资本回报率有多条返回纪录")
	zlog.Error(er.Error(), er)
	return nil, er
}

// 新增纪录
// by author Yeqc
// by time 2016-12-21 13:39:33
func (RocService) Add(model *par.Roc) error {
	return model.Add()
}

// 更新纪录
// by author Yeqc
// by time 2016-12-21 13:39:33
func (RocService) Update(model *par.Roc) error {
	return model.Update()
}

// 删除纪录
// by author Yeqc
// by time 2016-12-21 13:39:33
func (RocService) Delete(model *par.Roc) error {
	return model.Delete()
}
