package pricingService

import (
	"fmt"
	"reflect"

	"pccqcpa.com.cn/app/rpm/api/models/biz/ln"
)

type PricingBase struct {
	LnBusiness ln.LnBusiness
}

// LnBusinessToLnPricing func PricingBase 将业务基础信息转换到定价单中
func (pb *PricingBase) LnBusinessToLnPricing(pricing *ln.LnPricing) {
	pricing.BusinessCode = pb.LnBusiness.BusinessCode
	pricing.PlnCode = pb.LnBusiness.BusinessCode
	pricing.CustCode = pb.LnBusiness.Cust.CustCode
	pricing.CustName = pb.LnBusiness.Cust.CustName
	pricing.CustType = pb.LnBusiness.Cust.CustType
	pricing.CustImplvl = pb.LnBusiness.Cust.CustImplvl
	pricing.CustCredit = pb.LnBusiness.Cust.CustCredit
	pricing.BranchCode = pb.LnBusiness.Cust.Branch.OrganCode
	pricing.BranchName = pb.LnBusiness.Cust.Branch.OrganName
	pricing.IndustryCode = pb.LnBusiness.Cust.Industry.IndustryCode
	pricing.IndustryName = pb.LnBusiness.Cust.Industry.IndustryName

}

func (pb *PricingBase) ReflectToPricing(param map[string]interface{}, pricing interface{}) error {
	s := reflect.ValueOf(pricing).Elem()
	if !s.CanSet() {
		return fmt.Errorf("定价单实体不可赋值")
	}
	for k, v := range param {
		if s.FieldByName(k).IsValid() {
			switch s.FieldByName(k).Type().String() {
			case "float64":
				s.FieldByName(k).SetFloat(v.(float64))
			case "string":
				s.FieldByName(k).SetString(v.(string))
			}
		}
	}
	return nil
}
