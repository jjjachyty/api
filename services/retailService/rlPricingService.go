package retailService

import (
	"errors"
	"github.com/lunny/tango"
	"pccqcpa.com.cn/app/rpm/api/models/biz/rl"
	"pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/app/rpm/api/util/currentMsg"
	// "strings"
	// "fmt"
)

type RLPricingService struct {
	Ctx tango.Ctx
}

var rlPricingModel rl.RlPricing

//分页查询
func (rlp *RLPricingService) List(param ...map[string]interface{}) (*util.PageData, error) {
	if 0 < len(param) {
		paramMap := param[0]
		flag := util.CheckQueryParams(paramMap)
		if flag {
			return nil, errors.New("查询条件中不允许出现[^$.*+?{}()[]|'`~\\]特殊字符")
		}

		//处理排序条件
		if _, ok := paramMap["sort"]; ok {
			if paramMap["sort"] == "Cust.CustCode" {
				paramMap["sort"] = "t2.cust_code"
			}
			if paramMap["sort"] == "Cust.CustName" {
				paramMap["sort"] = "t2.cust_name"
			}
			if paramMap["sort"] == "Product.ProductName" {
				paramMap["sort"] = "t4.product_name"
			}
		}

		//客户经理分页
		if _, ok := paramMap["type"]; ok {
			userName, err := currentMsg.GetCurrentUserName(rlp.Ctx)
			if nil != err {
				return nil, err
			}
			param[0]["cust_manager"] = userName //添加当前用户(客户经理)
			var flag bool = false               //标志是否是条件查询
			//条件查询并分页：custCode custName status 只要存在一个参数即判定为分页下的条件查询
			if custCode, ok := paramMap["cust._cust_code"]; ok {
				paramMap["t2.cust_code"] = custCode
				delete(paramMap, "cust._cust_code")
				flag = true
			}
			if custName, ok := paramMap["cust._cust_name"]; ok {
				paramMap["t2.cust_name"] = custName
				delete(paramMap, "cust._cust_name")
				flag = true
			}
			if status, ok := paramMap["status"]; ok {
				paramMap["t1.status"] = status
				delete(paramMap, "status")
				flag = true
			}
			param[0]["flag"] = flag //添加标志状态

			return rlPricingModel.ListCusMan(param...)
		}

	}
	//客户分页
	return rlPricingModel.List(param...)
}

//叶全才  添加贷款申请信息
func (r *RLPricingService) Add(rlPricing *rl.RlPricing) error {
	err := rlPricing.Add()
	if nil != err {
		return err
	}
	return nil
}

//更新零售贷款状态信息
func (r *RLPricingService) Update(rlPricing *rl.RlPricing) error {
	return rlPricing.Update()
}
