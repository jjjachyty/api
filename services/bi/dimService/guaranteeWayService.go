package dimService

import (
	"database/sql"
	"errors"

	"pccqcpa.com.cn/app/rpm/api/models/bi/dim"

	"pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/components/zlog"
)

type GuaranteeWayService struct{}

var dimGuaranteeWay dim.DimGuaranteeWay

// 分页操作
// by author Jason
// by time 2016-10-31 15:34:35
func (GuaranteeWayService) List(param ...map[string]interface{}) (*util.PageData, error) {
	return dimGuaranteeWay.List(param...)
}

// 多参数查询返回多条纪录
// by author Jason
// by time 2016-10-31 15:34:35
func (GuaranteeWayService) Find(param ...map[string]interface{}) ([]*dim.DimGuaranteeWay, error) {
	return dimGuaranteeWay.Find(param...)
}

// 多参数查询返回一条纪录
// by author Jason
// by time 2016-10-31 15:34:35
func (this GuaranteeWayService) FindOne(paramMap map[string]interface{}) (*dim.DimGuaranteeWay, error) {
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
	er := errors.New("查询【RPM_BI_DIM_GUARANTEE_WAY】业务分析担保方式维度表有多条返回纪录")
	zlog.Error(er.Error(), er)
	return nil, er
}

// 新增纪录
// by author Jason
// by time 2016-10-31 15:34:35
func (GuaranteeWayService) Add(model *dim.DimGuaranteeWay) error {
	return model.Add()
}

// 新增纪录带事务
// by author Jason
// by time 2016-10-31 15:34:54
func (GuaranteeWayService) BatchAdd(models []dim.DimGuaranteeWay) (sql.Result, error) {
	return dimGuaranteeWay.BatchAdd(models)
}

// 更新纪录
// by author Jason
// by time 2016-10-31 15:34:35
func (GuaranteeWayService) Update(model *dim.DimGuaranteeWay) error {
	return model.Update()
}

// 删除纪录
// by author Jason
// by time 2016-10-31 15:34:35
func (GuaranteeWayService) Delete(model *dim.DimGuaranteeWay) error {
	return model.Delete()
}
