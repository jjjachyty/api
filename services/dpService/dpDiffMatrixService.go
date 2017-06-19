package dpService

import (
	"errors"

	"pccqcpa.com.cn/app/rpm/api/models/dp"
	"pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/components/zlog"
)

type DpDiffMatrixService struct{}

var dpDiffMatrix dp.DpDiffMatrix

// 分页操作
// by author Jason
// by time 2017-02-07 16:48:18
func (DpDiffMatrixService) List(param ...map[string]interface{}) (*util.PageData, error) {
	return dpDiffMatrix.List(param...)
}

// 多参数查询返回多条纪录
// by author Jason
// by time 2017-02-07 16:48:18
func (DpDiffMatrixService) Find(param ...map[string]interface{}) ([]*dp.DpDiffMatrix, error) {
	var paramMap map[string]interface{}
	if 0 != len(param) {
		paramMap = param[0]
	}
	paramMap["sort"] = "t.branch, t.product, t.term, t.cust_grade, t.amount_grade, t.channel"
	return dpDiffMatrix.Find(paramMap)
}

// 多参数查询返回一条纪录
// by author Jason
// by time 2017-02-07 16:48:18
func (this DpDiffMatrixService) FindOne(paramMap map[string]interface{}) (*dp.DpDiffMatrix, error) {
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
	er := errors.New("查询【RPM_BIZ_DP_DIFF_MATRIX】存款差异化矩阵表有多条返回纪录")
	zlog.Error(er.Error(), er)
	return nil, er
}

// 新增纪录
// by author Jason
// by time 2017-02-07 16:48:18
func (DpDiffMatrixService) Add(model *dp.DpDiffMatrix) error {
	return model.Add()
}

// 更新纪录
// by author Jason
// by time 2017-02-07 16:48:18
func (DpDiffMatrixService) Update(model *dp.DpDiffMatrix) error {
	return model.Update()
}

// 删除纪录
// by author Jason
// by time 2017-02-07 16:48:18
func (DpDiffMatrixService) Delete(model *dp.DpDiffMatrix) error {
	return model.Delete()
}
