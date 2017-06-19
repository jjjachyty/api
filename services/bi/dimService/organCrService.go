package dimService

import (
	"database/sql"
	"errors"

	"pccqcpa.com.cn/app/rpm/api/models/bi/dim"

	"pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/components/zlog"
)

type OrganCrService struct{}

var dimOrganCr dim.DimOrganCr

// 分页操作
// by author Jason
// by time 2016-10-31 15:36:53
func (OrganCrService) List(param ...map[string]interface{}) (*util.PageData, error) {
	return dimOrganCr.List(param...)
}

// 多参数查询返回多条纪录
// by author Jason
// by time 2016-10-31 15:36:53
func (OrganCrService) Find(param ...map[string]interface{}) ([]*dim.DimOrganCr, error) {
	return dimOrganCr.Find(param...)
}

// 多参数查询返回一条纪录
// by author Jason
// by time 2016-10-31 15:36:53
func (this OrganCrService) FindOne(paramMap map[string]interface{}) (*dim.DimOrganCr, error) {
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
	er := errors.New("查询【RPM_BI_DIM_ORGAN_CR】贷款业务分析方案信贷机构维度表有多条返回纪录")
	zlog.Error(er.Error(), er)
	return nil, er
}

// 新增纪录
// by author Jason
// by time 2016-10-31 15:36:53
func (OrganCrService) Add(model *dim.DimOrganCr) error {
	return model.Add()
}

// 新增纪录带事务
// by author Jason
// by time 2016-10-31 15:34:54
func (OrganCrService) BatchAdd(models []dim.DimOrganCr) (sql.Result, error) {
	return dimOrganCr.BatchAdd(models)
}

// 更新纪录
// by author Jason
// by time 2016-10-31 15:36:53
func (OrganCrService) Update(model *dim.DimOrganCr) error {
	return model.Update()
}

// 删除纪录
// by author Jason
// by time 2016-10-31 15:36:53
func (OrganCrService) Delete(model *dim.DimOrganCr) error {
	return model.Delete()
}
