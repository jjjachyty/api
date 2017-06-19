package dimService

import (
	"database/sql"
	"errors"

	"pccqcpa.com.cn/app/rpm/api/models/bi/dim"

	"pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/components/zlog"
)

type OrganMcService struct{}

var dimOrganMc dim.DimOrganMc

// 分页操作
// by author Jason
// by time 2016-10-31 15:37:34
func (OrganMcService) List(param ...map[string]interface{}) (*util.PageData, error) {
	return dimOrganMc.List(param...)
}

// 多参数查询返回多条纪录
// by author Jason
// by time 2016-10-31 15:37:34
func (OrganMcService) Find(param ...map[string]interface{}) ([]*dim.DimOrganMc, error) {
	return dimOrganMc.Find(param...)
}

// 多参数查询返回一条纪录
// by author Jason
// by time 2016-10-31 15:37:34
func (this OrganMcService) FindOne(paramMap map[string]interface{}) (*dim.DimOrganMc, error) {
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
	er := errors.New("查询【RPM_BI_DIM_ORGAN_MC】贷款业务分析方案管快机构维度表有多条返回纪录")
	zlog.Error(er.Error(), er)
	return nil, er
}

// 新增纪录
// by author Jason
// by time 2016-10-31 15:37:34
func (OrganMcService) Add(model *dim.DimOrganMc) error {
	return model.Add()
}

// 新增纪录带事务
// by author Jason
// by time 2016-10-31 15:34:54
func (OrganMcService) BatchAdd(models []dim.DimOrganMc) (sql.Result, error) {
	return dimOrganMc.BatchAdd(models)
}

// 更新纪录
// by author Jason
// by time 2016-10-31 15:37:34
func (OrganMcService) Update(model *dim.DimOrganMc) error {
	return model.Update()
}

// 删除纪录
// by author Jason
// by time 2016-10-31 15:37:34
func (OrganMcService) Delete(model *dim.DimOrganMc) error {
	return model.Delete()
}
