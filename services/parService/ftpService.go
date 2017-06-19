package parService

import (
	"errors"
	"pccqcpa.com.cn/app/rpm/api/models/par"
	"pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/components/zlog"
)

type FtpRateService struct{}

var ftpModel par.Ftp

// 分页操作
// by author Yeqc
// by time 2016-12-19 10:23:37
func (FtpRateService) List(param ...map[string]interface{}) (*util.PageData, error) {
	if 0 < len(param) {
		paramMap := param[0]
		flag := util.CheckQueryParams(paramMap)
		if flag {
			return nil, errors.New("查询条件中不允许出现[^$.*+?{}()[]|'`~\\]特殊字符")
		}
		if organName, ok := paramMap["organ_name"]; ok {
			key := []interface{}{
				"organ.organ_name", "product.product_name",
			}
			value := []interface{}{
				organName, organName,
			}
			searchLike := []map[string]interface{}{
				map[string]interface{}{
					"type":  "or",
					"key":   key,
					"value": value,
				},
			}
			paramMap["searchLike"] = searchLike
			delete(paramMap, "organ_name")
		}

	}
	return ftpModel.List(param...)
}

func (f *FtpRateService) Find(param ...map[string]interface{}) ([]*par.Ftp, error) {
	return ftpModel.Find(param...)
}

// 多参数查询返回一条纪录
// by author Yeqc
// by time 2016-12-19 10:23:37
func (this FtpRateService) FindOne(paramMap map[string]interface{}) (*par.Ftp, error) {
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
	er := errors.New("查询【RPM_PAR_FTP】资金成本率有多条返回纪录")
	zlog.Error(er.Error(), er)
	return nil, er
}

// 新增纪录
// by author Yeqc
// by time 2016-12-19 10:23:37
func (FtpRateService) Add(model *par.Ftp) error {
	return model.Add()
}

// 更新纪录
// by author Yeqc
// by time 2016-12-19 10:23:37
func (FtpRateService) Update(model *par.Ftp) error {
	return model.Update()
}

// 删除纪录
// by author Yeqc
// by time 2016-12-19 10:23:37
func (FtpRateService) Delete(model *par.Ftp) error {
	return model.Delete()
}
