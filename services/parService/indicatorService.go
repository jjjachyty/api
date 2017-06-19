package parService

import (
	"fmt"
	"strings"

	"pccqcpa.com.cn/app/rpm/api/models/par"
	"pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/components/zlog"
)

type IndicatorService struct {
}

var indicatorModel par.Indicator

// 分页查询
func (ind *IndicatorService) List(param ...map[string]interface{}) (*util.PageData, error) {
	var paramMap = ind.handleParam(param...)
	return indicatorModel.List(paramMap)
}

// 多参数查询
func (ind *IndicatorService) Find(param ...map[string]interface{}) ([]*par.Indicator, error) {
	var paramMap = ind.handleParam(param...)
	return indicatorModel.Find(paramMap)
}

func (ind *IndicatorService) Add(indicator *par.Indicator) (*par.Indicator, error) {
	param := map[string]interface{}{
		"indicator_code": indicator.IndicatorCode,
	}
	indicators, err := indicator.Find(param)
	if nil != err {
		return nil, err
	}
	if 0 < len(indicators) {
		er := fmt.Errorf("指标码值为[%v]的指标已存在，不可以新增", indicator.IndicatorCode)
		zlog.Error(er.Error(), err)
		return nil, er
	}
	err = indicator.Add()
	if nil != err {
		return nil, err
	}
	indi, _ := indicator.Find(param)
	util.PutCacheByCacheName(util.RPM_INDICATOR_CACHE, indi[0].IndicatorCode, indi[0], 0)

	return indi[0], nil
}

func (ind *IndicatorService) Update(indicator *par.Indicator) (*par.Indicator, error) {
	err := indicator.Update()
	if nil != err {
		return nil, err
	}
	param := map[string]interface{}{
		"indicator_code": indicator.IndicatorCode,
	}
	indicators, _ := indicator.Find(param)
	// 更新指标，默认不更新指标码值
	util.PutCacheByCacheName(util.RPM_INDICATOR_CACHE, indicators[0].IndicatorCode, indicators[0], 0)

	return indicators[0], nil
}

// 判断是否有指标引用
// 没有其他指标引用方可删除
func (ind *IndicatorService) Delete(indicator *par.Indicator) error {
	searchLike := []map[string]interface{}{
		map[string]interface{}{
			"type":  "like",
			"key":   "indicator_before_rely",
			"value": "," + indicator.IndicatorCode + ",",
		},
	}
	param := map[string]interface{}{
		"searchLike": searchLike,
	}
	indicators, err := indicator.Find(param)
	if nil != err {
		return err
	} else if 0 < len(indicators) {
		indicatorCodes := ind.getIndicatorCodes(indicators)
		er := fmt.Errorf("有指标码值为[%v]引用本指标，不可删除", indicatorCodes)
		zlog.Error(er.Error(), err)
		return er
	}
	err = indicator.Delete()
	if nil == err {
		util.DeleteCacheByCacheName(util.RPM_INDICATOR_CACHE, indicator.IndicatorCode)
	}
	return err
}

func (ind *IndicatorService) getIndicatorCodes(indicators []*par.Indicator) string {
	var indicatorCodes string
	for _, indicator := range indicators {
		indicatorCodes += indicator.IndicatorCode + ","
	}
	indicatorCodes = strings.TrimRight(indicatorCodes, ",")
	return indicatorCodes
}

func (ind *IndicatorService) handleParam(param ...map[string]interface{}) map[string]interface{} {
	var searchLike = make([]map[string]interface{}, 0)
	var paramMap = make(map[string]interface{})
	if 0 < len(param) {
		paramMap = param[0]
		if indicatorType, ok := paramMap["indicator_type"]; ok {
			var keys []interface{}
			var values []interface{}
			for _, v := range strings.Split(indicatorType.(string), "_") {
				keys = append(keys, "indicator_type")
				values = append(values, v)
			}
			searchLike = append(searchLike, map[string]interface{}{
				"key":   keys,
				"value": values,
				"type":  "or",
			})
			delete(paramMap, "indicator_type")
		}
		if indicatorName, ok := paramMap["indicator_name"]; ok {
			searchLike = append(searchLike, map[string]interface{}{
				"key":   []interface{}{"indicator_name", "indicator_code"},
				"value": []interface{}{indicatorName, indicatorName},
				"type":  "or",
			})
			delete(paramMap, "indicator_code")
			delete(paramMap, "indicator_name")
		}
	}
	paramMap["searchLike"] = searchLike
	// fmt.Println("---------", paramMap)
	return paramMap
}
