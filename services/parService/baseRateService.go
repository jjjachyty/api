package parService

import (
	"errors"

	"pccqcpa.com.cn/app/rpm/api/models/par"

	"pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/components/zlog"
)

type BaseRateService struct{}

var baseRate par.BaseRate

// 分页操作
// by author Yeqc
// by time 2016-12-19 10:49:29
func (BaseRateService) List(param ...map[string]interface{}) (*util.PageData, error) {
	if 0 < len(param) {
		paramMap := param[0]
		flag := util.CheckQueryParams(paramMap)
		if flag {
			return nil, errors.New("查询条件中不允许出现[^$.*+?{}()[]|'`~\\]特殊字符")
		}
		if baseRateName, ok := paramMap["base_rate_name"]; ok {
			key := []interface{}{
				"t.base_rate_name",
			}
			value := []interface{}{
				baseRateName, baseRateName,
			}
			searchLike := []map[string]interface{}{
				map[string]interface{}{
					"type":  "or",
					"key":   key,
					"value": value,
				},
			}
			paramMap["searchLike"] = searchLike
			delete(paramMap, "base_rate_name")
		}

	}
	return baseRate.List(param...)
}

// 多参数查询返回多条纪录
// by author Yeqc
// by time 2016-12-19 10:49:29
func (BaseRateService) Find(param ...map[string]interface{}) ([]*par.BaseRate, error) {
	return baseRate.Find(param...)
}

// 多参数查询返回一条纪录
// by author Yeqc
// by time 2016-12-19 10:49:29
func (this BaseRateService) FindOne(paramMap map[string]interface{}) (*par.BaseRate, error) {
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
	er := errors.New("查询【RPM_PAR_BASE_RATE】基准利率有多条返回纪录")
	zlog.Error(er.Error(), er)
	return nil, er
}

// 新增纪录
// by author Yeqc
// by time 2016-12-19 10:49:29
func (BaseRateService) Add(model *par.BaseRate) error {
	return model.Add()
}

// 更新纪录
// by author Yeqc
// by time 2016-12-19 10:49:29
func (BaseRateService) Update(model *par.BaseRate) error {
	return model.Update()
}

// 删除纪录
// by author Yeqc
// by time 2016-12-19 10:49:29
func (BaseRateService) Delete(model *par.BaseRate) error {
	return model.Delete()
}
