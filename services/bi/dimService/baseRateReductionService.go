package dimService

import (
	"database/sql"
	"errors"

	"pccqcpa.com.cn/app/rpm/api/models/bi/dim"

	"pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/components/zlog"
)

type BaseRateReductionService struct{}

var dimBaseRateReduction dim.DimBaseRateReduction

// 分页操作
// by author Jason
// by time 2016-10-31 15:34:54
func (BaseRateReductionService) List(param ...map[string]interface{}) (*util.PageData, error) {
	return dimBaseRateReduction.List(param...)
}

// 多参数查询返回多条纪录
// by author Jason
// by time 2016-10-31 15:34:54
func (BaseRateReductionService) Find(param ...map[string]interface{}) ([]*dim.DimBaseRateReduction, error) {
	return dimBaseRateReduction.Find(param...)
}

// 多参数查询返回一条纪录
// by author Jason
// by time 2016-10-31 15:34:54
func (this BaseRateReductionService) FindOne(paramMap map[string]interface{}) (*dim.DimBaseRateReduction, error) {
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
	er := errors.New("查询【RPM_BI_DIM_BASE_RATE_REDUCTION】贷款业务分析基准利率降息维度表有多条返回纪录")
	zlog.Error(er.Error(), er)
	return nil, er
}

// 新增纪录
// by author Jason
// by time 2016-10-31 15:34:54
func (BaseRateReductionService) Add(model *dim.DimBaseRateReduction) error {
	return model.Add()
}

// 新增纪录带事务
// by author Jason
// by time 2016-10-31 15:34:54
func (BaseRateReductionService) BatchAdd(models []dim.DimBaseRateReduction) (sql.Result, error) {
	return dimBaseRateReduction.BatchAdd(models)
}

// 更新纪录
// by author Jason
// by time 2016-10-31 15:34:54
func (BaseRateReductionService) Update(model *dim.DimBaseRateReduction) error {
	return model.Update()
}

// 删除纪录
// by author Jason
// by time 2016-10-31 15:34:54
func (BaseRateReductionService) Delete(model *dim.DimBaseRateReduction) error {
	return model.Delete()
}
