package dpService

import (
	"errors"

	"github.com/lunny/tango"

	"pccqcpa.com.cn/app/rpm/api/models/dp"
	"pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/app/rpm/api/util/currentMsg"
	"pccqcpa.com.cn/components/zlog"
)

type DpOneBusinessService struct {
	Ctx tango.Ctx
}

var dpOneBusiness dp.DpOneBusiness

// 分页操作
// by author Jason
// by time 2016-12-06 16:59:09
func (DpOneBusinessService) List(param ...map[string]interface{}) (*util.PageData, error) {
	return dpOneBusiness.List(param...)
}

// 多参数查询返回多条纪录
// by author Jason
// by time 2016-12-06 16:59:09
func (DpOneBusinessService) Find(param ...map[string]interface{}) ([]*dp.DpOneBusiness, error) {
	return dpOneBusiness.Find(param...)
}

// 多参数查询返回一条纪录
// by author Jason
// by time 2016-12-06 16:59:09
func (this DpOneBusinessService) FindOne(paramMap map[string]interface{}) (*dp.DpOneBusiness, error) {
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
	er := errors.New("查询【RPM_BIZ_DP_ONE_BUSINESS】一对一存款业务表有多条返回纪录")
	zlog.Error(er.Error(), er)
	return nil, er
}

// 新增纪录
// by author Jason
// by time 2016-12-06 16:59:09
func (this DpOneBusinessService) Add(model *dp.DpOneBusiness) error {
	err := this.setOrganCode(model)
	if nil != err {
		return err
	}
	err = model.Add()
	if nil != err {
		return err
	}
	err = DpOnePricingService{}.Patch(map[string]interface{}{
		"business_code": model.BusinessCode,
		"status":        util.DP_ONE_PRICING_STATUS_UNFINISHED,
	})
	if nil != err {
		return err
	}
	return dp.DpOnePricing{BusinessCode: model.BusinessCode}.Patch(map[string]interface{}{"status": util.DP_ONE_PRICING_STATUS_UNFINISHED})
}

// 更新纪录
// by author Jason
// by time 2016-12-06 16:59:09
func (this DpOneBusinessService) Update(model *dp.DpOneBusiness) error {
	err := this.setOrganCode(model)
	if nil != err {
		return err
	}
	err = model.Update()
	if nil != err {
		return err
	}
	err = DpOnePricingService{}.Patch(map[string]interface{}{
		"business_code": model.BusinessCode,
		"status":        util.DP_ONE_PRICING_STATUS_UNFINISHED,
	})
	if nil != err {
		return err
	}
	return dp.DpOnePricing{BusinessCode: model.BusinessCode}.Patch(map[string]interface{}{"status": util.DP_ONE_PRICING_STATUS_UNFINISHED})
}

// 删除纪录
// by author Jason
// by time 2016-12-06 16:59:09
func (DpOneBusinessService) Delete(model *dp.DpOneBusiness) error {

	err := DpOnePricingService{}.Patch(map[string]interface{}{
		"business_code": model.BusinessCode,
		"status":        util.DP_ONE_PRICING_STATUS_UNFINISHED,
	})
	if nil != err {
		return err
	}
	return model.Delete()
}

func (this *DpOneBusinessService) setOrganCode(model *dp.DpOneBusiness) error {
	branchCode, err := currentMsg.GetCurrentUserBranchCode(this.Ctx)
	if nil != err {
		return err
	}
	model.Organ.OrganCode = branchCode
	return nil
}
