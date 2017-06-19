package dimService

import (
	"pccqcpa.com.cn/app/rpm/api/models/dim"
)

type ProductService struct {
}

var productModel dim.Product

func (this *ProductService) SelectPorductByParams(paramMap map[string]interface{}) ([]*dim.Product, error) {
	return productModel.SelectPorductByParams(paramMap)
}
