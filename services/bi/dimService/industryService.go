package dimService

import (
	"database/sql"
	"errors"

	"pccqcpa.com.cn/app/rpm/api/models/bi/dim"

	"pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/components/zlog"
)

type IndustryService struct{}

var dimIndustry dim.DimIndustry

// 分页操作
// by author Jason
// by time 2016-10-31 15:34:16
func (IndustryService) List(param ...map[string]interface{}) (*util.PageData, error) {
	return dimIndustry.List(param...)
}

// 多参数查询返回多条纪录
// by author Jason
// by time 2016-10-31 15:34:16
func (IndustryService) Find(param ...map[string]interface{}) ([]*dim.DimIndustry, error) {
	return dimIndustry.Find(param...)
}

// 多参数查询返回一条纪录
// by author Jason
// by time 2016-10-31 15:34:16
func (this IndustryService) FindOne(paramMap map[string]interface{}) (*dim.DimIndustry, error) {
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
	er := errors.New("查询【RPM_BI_DIM_INDUSTRY】业务方案分析行业维度表有多条返回纪录")
	zlog.Error(er.Error(), er)
	return nil, er
}

// 新增纪录
// by author Jason
// by time 2016-10-31 15:34:16
func (IndustryService) Add(model *dim.DimIndustry) error {
	return model.Add()
}

// 新增纪录带事务
// by author Jason
// by time 2016-10-31 15:34:54
func (IndustryService) BatchAdd(models []dim.DimIndustry) (sql.Result, error) {
	return dimIndustry.BatchAdd(models)
}

// 更新纪录
// by author Jason
// by time 2016-10-31 15:34:16
func (IndustryService) Update(model *dim.DimIndustry) error {
	return model.Update()
}

// 删除纪录
// by author Jason
// by time 2016-10-31 15:34:16
func (IndustryService) Delete(model *dim.DimIndustry) error {
	return model.Delete()
}
