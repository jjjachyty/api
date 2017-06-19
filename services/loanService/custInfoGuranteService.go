package loanService

import (
	"pccqcpa.com.cn/app/rpm/api/util"
)

type CustInfoGuranteService struct {
}

func (c *CustInfoGuranteService) List(param ...map[string]interface{}) (*util.PageData, error) {
	return custInfoModel.List(param...)
}
