package classfyService

import (
	"fmt"

	"pccqcpa.com.cn/app/rpm/api/models/biz/cf"
	"pccqcpa.com.cn/components/zlog"
)

type ClassifyResultService struct{}

var classifyResultModel cf.ClassifyResult

func (c ClassifyResultService) FindOne(paramMap map[string]interface{}) (*cf.ClassifyResult, error) {
	rst, err := classifyResultModel.Find(paramMap)
	if nil != err {
		return nil, err
	}
	switch len(rst) {
	case 0:
		return nil, nil
	case 1:
		return rst[0], nil
	}
	er := fmt.Errorf("查询客户分类结果表有多条纪录")
	zlog.Error(er.Error(), er)
	return nil, er
}

func (c ClassifyResultService) Add(model cf.ClassifyResult) error {
	return model.Add()
}

func (c ClassifyResultService) Update(model cf.ClassifyResult) error {
	return model.Update()
}
