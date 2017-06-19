package parService

import (
	"errors"

	// "fmt"

	"pccqcpa.com.cn/app/rpm/api/models/par"

	"pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/components/zlog"
	// "platform/dbobj"
)

type DpBaseRateService struct{}

var dpBaseRate par.DpBaseRate

// 分页操作
// by author Yeqc
// by time 2017-01-05 14:25:41
func (DpBaseRateService) List(param ...map[string]interface{}) (*util.PageData, error) {
	if 0 < len(param) {
		paramMap := param[0]
		flag := util.CheckQueryParams(paramMap)
		if flag {
			return nil, errors.New("查询条件中不允许出现[^$.*+?{}()[]|'`~\\]特殊字符")
		}
		if name, ok := paramMap["name"]; ok {
			key := []interface{}{
				"name", "product.product_name",
			}
			value := []interface{}{
				name, name,
			}
			searchLike := []map[string]interface{}{
				map[string]interface{}{
					"type":  "or",
					"key":   key,
					"value": value,
				},
			}
			paramMap["searchLike"] = searchLike
			delete(paramMap, "name")
		}
	}
	return dpBaseRate.List(param...)
}

// 多参数查询返回多条纪录
// by author Yeqc
// by time 2017-01-05 14:25:41
func (DpBaseRateService) Find(param ...map[string]interface{}) ([]*par.DpBaseRate, error) {
	return dpBaseRate.Find(param...)
}

// 多参数查询返回一条纪录
// by author Yeqc
// by time 2017-01-05 14:25:41
func (this DpBaseRateService) FindOne(paramMap map[string]interface{}) (*par.DpBaseRate, error) {
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
	er := errors.New("查询【RPM_PAR_DP_BASE_RATE】存款基准利率及挂牌有多条返回纪录")
	zlog.Error(er.Error(), er)
	return nil, er
}

// 新增纪录
// by author Yeqc
// by time 2017-01-05 14:25:41
func (DpBaseRateService) Add(model *par.DpBaseRate) error {
	return model.Add()
}

// 更新纪录
// by author Yeqc
// by time 2017-01-05 14:25:41
func (DpBaseRateService) Update(model *par.DpBaseRate) error {
	return model.Update()
}

// 删除纪录
// by author Yeqc
// by time 2017-01-05 14:25:41
func (DpBaseRateService) Delete(model *par.DpBaseRate) error {
	return model.Delete()
}
