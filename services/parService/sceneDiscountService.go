package parService

import (
	"errors"
	"fmt"

	"pccqcpa.com.cn/app/rpm/api/models/par"
	"pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/components/zlog"
)

type SceneDiscountService struct {
}

var sceneDiscountModel par.SceneDiscountModel

var sceneDiscount par.SceneDiscount

func (SceneDiscountService) GetStockDiscount(custImplvl string, custSize string) (par.NullSceneDiscount, error) {
	var sceneDiscount par.NullSceneDiscount
	sceneDiscount.CustImplvl.String = custImplvl
	sceneDiscount.CustSize.String = custSize
	sceneDiscount.BizType.String = util.DISCOUNT_STOCK
	sceneDiscounts, err := sceneDiscountModel.QueryDiscountRate(sceneDiscount)
	if nil == err {
		len := len(sceneDiscounts)
		if 1 < len {
			return sceneDiscount, fmt.Errorf("存量优惠折让率查询重复,请检查存量优惠折让率参数表[RPM_PAR_SCENE_DISCOUNT]配置。查询条件[客户等级(CustImplvl):%s,客户规模(CustSize):%s,折让率类型(BizType):%s]", custImplvl, custSize, util.DISCOUNT_STOCK)
		} else if 1 == len {
			return sceneDiscounts[0], nil
		}
	}

	return sceneDiscount, fmt.Errorf("存量优惠折让率查询为空,请检查存量优惠折让率参数表[RPM_PAR_SCENE_DISCOUNT]配置。查询条件[客户等级(CustImplvl):%s,客户规模(CustSize):%s,折让率类型(BizType):%s]", custImplvl, custSize, util.DISCOUNT_STOCK)
}

func (SceneDiscountService) GetDerivedDiscount(custImplvl string, custSize string, product string) (par.NullSceneDiscount, error) {
	var sceneDiscount par.NullSceneDiscount
	sceneDiscount.CustImplvl.String = custImplvl
	sceneDiscount.CustSize.String = custSize
	sceneDiscount.BizType.String = util.DISCOUNT_DERIVED
	sceneDiscount.ProductCode.String = product
	sceneDiscounts, err := sceneDiscountModel.QueryDiscountRate(sceneDiscount)
	if nil == err {
		len := len(sceneDiscounts)
		if 1 < len {
			return sceneDiscount, fmt.Errorf("派生折让率查询重复,请检查派生折让率参数表[RPM_PAR_SCENE_DISCOUNT]配置。查询条件[客户等级(CustImplvl):%s,客户规模(CustSize):%s,折让率类型(BizType):%s,产品(ProductCode):%s]", custImplvl, custSize, util.DISCOUNT_DERIVED, product)

		} else if 1 == len {
			return sceneDiscounts[0], nil
		}
	}

	return sceneDiscount, fmt.Errorf("派生折让率查询为空,请检查派生折让率参数表[RPM_PAR_SCENE_DISCOUNT]配置。查询条件[客户等级(CustImplvl):%s,客户规模(CustSize):%s,折让率类型(BizType):%s,产品(ProductCode):%s]", custImplvl, custSize, util.DISCOUNT_DERIVED, product)
}

// 分页操作
// by author Jason
// by time 2017-02-17 10:25:04
func (SceneDiscountService) List(param ...map[string]interface{}) (*util.PageData, error) {
	return sceneDiscount.List(param...)
}

// 多参数查询返回多条纪录
// by author Jason
// by time 2017-02-17 10:25:04
func (SceneDiscountService) Find(param ...map[string]interface{}) ([]*par.SceneDiscount, error) {
	return sceneDiscount.Find(param...)
}

// 多参数查询返回一条纪录
// by author Jason
// by time 2017-02-17 10:25:04
func (this SceneDiscountService) FindOne(paramMap map[string]interface{}) (*par.SceneDiscount, error) {
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
	er := errors.New("查询【RPM_PAR_SCENE_DISCOUNT】派生优惠参数表有多条返回纪录")
	zlog.Error(er.Error(), er)
	return nil, er
}

// 新增纪录
// by author Jason
// by time 2017-02-17 10:25:04
func (SceneDiscountService) Add(model *par.SceneDiscount) error {
	return model.Add()
}

// 更新纪录
// by author Jason
// by time 2017-02-17 10:25:04
func (SceneDiscountService) Update(model *par.SceneDiscount) error {
	return model.Update()
}

// 删除纪录
// by author Jason
// by time 2017-02-17 10:25:04
func (SceneDiscountService) Delete(model *par.SceneDiscount) error {
	return model.Delete()
}
