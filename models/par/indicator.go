package par

import (
	"database/sql"
	"fmt"
	"regexp"
	"strings"
	"time"

	"pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/components/zlog"

	modlesUtil "pccqcpa.com.cn/app/rpm/api/models/util"
)

/**
 * 指标CRUD，对应数据库表RPM_PAR_BASE_INDICATORS
 */
/**
 * @apiDefine Indicator
 * @apiSuccess {string}     UUID            		主键默认值sys_guid()
 * @apiSuccess {string}   	IndicatorCode   		指标码值
 * @apiSuccess {string}   	IndicatorName      		指标名称
 * @apiSuccess {Product}   	IndicatorType   		指标类型
 * @apiSuccess {string}   	IndicatorBeforeRely    	指标依赖
 * @apiSuccess {string}   	Method      			函数指标调用的方法
 * @apiSuccess {string}   	FormulaeCode      		指标公式
 * @apiSuccess {string}   	FormulaeCn      		指标公式中文
 * @apiSuccess {string}   	IndicatorBusinessType   业务类型
 * @apiSuccess {string}   	Status      			状态（0：内置 1:非内置）
 * @apiSuccess {string}   	Flag      				生效标志 0：停用 1:启用
 * @apiSuccess {time.Time}	CreateTime      		创建时间
 * @apiSuccess {string}   	CreateUser      		创建人
 * @apiSuccess {time.Time}	UpdateTime      		更新时间
 * @apiSuccess {string}   	UpdateUser      		更新人
 */
type Indicator struct {
	UUID                  string       // 主键
	IndicatorCode         string       // 指标码值
	IndicatorName         string       // 指标名称
	IndicatorType         string       // 指标类型
	IndicatorBeforeRely   []*Indicator // 指标依赖
	Method                string       // 函数指标调用的方法
	FormulaeCode          string       // 指标公式
	FormulaeCn            string       // 指标公式中文
	IndicatorBusinessType string       // 业务类型
	Status                string       // 状态（0：内置 1:非内置）
	Flag                  string       // 生效标志 0：停用 1:启用
	CreateTime            time.Time    // 指标创建时间
	UpdateTime            time.Time    // 指标更新时间
	CreateUser            string       // 指标创建人
	UpdateUser            string       // 指标更新人
}

// 指标查询赋值
func (ind *Indicator) scan(rows *sql.Rows) (*Indicator, error) {
	var indicator = new(Indicator)
	var indicatorBeforeRely string
	values := []interface{}{
		&indicator.UUID,
		&indicator.IndicatorCode,
		&indicator.IndicatorName,
		&indicator.IndicatorType,
		&indicatorBeforeRely,
		&indicator.FormulaeCode,
		&indicator.FormulaeCn,
		&indicator.IndicatorBusinessType,
		&indicator.Flag,
		&indicator.Status,
		&indicator.Method,
		&indicator.CreateTime,
		&indicator.UpdateTime,
		&indicator.CreateUser,
		&indicator.UpdateUser,
	}
	// 获取前置指标
	err := util.OracleScan(rows, values)
	indicatorBeforeRelyNew := strings.Replace(indicatorBeforeRely, " ", "", -1)
	indicator.setBeforeIndicators(strings.TrimSpace(indicatorBeforeRelyNew))
	return indicator, err
}

// 分页查询
func (ind *Indicator) List(param ...map[string]interface{}) (*util.PageData, error) {
	pageData, rows, err := modlesUtil.List(indicatorTables, indicatorCols, indicatorColSort, param...)
	if nil != err {
		er := fmt.Errorf("分页查询指标出错")
		zlog.Error(er.Error(), err)
		return nil, er
	}
	defer rows.Close()
	var indicators []*Indicator
	for rows.Next() {
		indicator, err := ind.scan(rows)
		if nil != err {
			er := fmt.Errorf("多参数查询指标rows.Scan()出错")
			zlog.Error(er.Error(), err)
			return nil, er
		}
		indicators = append(indicators, indicator)
	}
	pageData.Rows = indicators
	return pageData, nil
}

// 多参数查询
func (ind *Indicator) Find(param ...map[string]interface{}) ([]*Indicator, error) {
	rows, err := modlesUtil.FindRows(indicatorTables, indicatorCols, indicatorColSort, param...)
	if nil != err {
		er := fmt.Errorf("多参数查询指标出错")
		zlog.Error(er.Error(), err)
		return nil, er
	}
	defer rows.Close()
	var indicators []*Indicator
	for rows.Next() {
		indicator, err := ind.scan(rows)
		if nil != err {
			er := fmt.Errorf("多参数查询指标rows.Scan()出错")
			zlog.Error(er.Error(), err)
			return nil, er
		}
		indicators = append(indicators, indicator)
	}
	return indicators, nil
}

func (ind *Indicator) Add() error {
	param := map[string]interface{}{
		"indicator_code":          ind.IndicatorCode,
		"indicator_name":          ind.IndicatorName,
		"indicator_type":          ind.IndicatorType,
		"indicator_before_rely":   ind.FormulaToRely(),
		"formulae_code":           ind.FormulaeCode,
		"formulae_cn":             ind.FormulaeCn,
		"indicator_business_type": ind.IndicatorBusinessType,
		"flag":        ind.Flag,
		"status":      ind.Status,
		"method":      ind.Method,
		"create_time": util.GetCurrentTime(),
		"update_time": util.GetCurrentTime(),
		"create_user": ind.CreateUser,
		"update_user": ind.UpdateUser,
	}
	err := util.OracleAdd(indicatorTables, param)
	if nil != err {
		er := fmt.Errorf("新增指标出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

func (ind *Indicator) Update() error {
	param := map[string]interface{}{
		"indicator_name":          ind.IndicatorName,
		"indicator_type":          ind.IndicatorType,
		"indicator_before_rely":   ind.FormulaToRely(),
		"formulae_code":           ind.FormulaeCode,
		"formulae_cn":             ind.FormulaeCn,
		"status":                  ind.Status,
		"method":                  ind.Method,
		"indicator_business_type": ind.IndicatorBusinessType,
		"flag":        ind.Flag,
		"update_time": util.GetCurrentTime(),
		"update_user": ind.UpdateUser,
	}
	whereParam := map[string]interface{}{
		"uuid": ind.UUID,
	}
	err := util.OracleUpdate(indicatorTables, param, whereParam)
	if nil != err {
		er := fmt.Errorf("更新指标失败")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

func (ind *Indicator) Delete() error {
	whereParam := map[string]interface{}{
		"uuid": ind.UUID,
	}
	err := util.OracleDelete(indicatorTables, whereParam)
	if nil != err {
		er := fmt.Errorf("删除指标失败")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

func (ind Indicator) getIndicatoBeforeString() string {
	var indicatorBeforeString string
	for _, v := range ind.IndicatorBeforeRely {
		indicatorBeforeString += v.IndicatorCode + ","
	}
	if 0 < len(indicatorBeforeString) {
		indicatorBeforeString = "," + indicatorBeforeString
	}
	return indicatorBeforeString
}

func (ind Indicator) FormulaToRely() string {
	if util.FIRST_INDICATOR == ind.IndicatorType {
		reg := regexp.MustCompile(`[,\+\-\\*\/\(\)]+`)
		return reg.ReplaceAllString(","+ind.FormulaeCode+",", ",")
	} else if util.SECOND_INDICATOR == ind.IndicatorType {
		return ind.getIndicatoBeforeString()
	}
	return ""
}

// 通过数据库中查询的前置指标字符串 在缓存中获取前置指标
// 先缓存查找，没有再从数据查找
func (ind *Indicator) setBeforeIndicators(indicatorBeforeRely string) {
	// fmt.Println("+++++++++", indicatorBeforeRely)
	var beforeIndicators []*Indicator
	for _, v := range strings.Split(indicatorBeforeRely, ",") {
		indicator := util.GetCacheByCacheName(util.RPM_INDICATOR_CACHE, v)
		if nil != indicator {
			// 从缓存中查找
			zlog.Infof("从缓存中查找指标[%v(%v)]", nil, indicator.(*Indicator).IndicatorName, indicator.(*Indicator).IndicatorCode)
			beforeIndicators = append(beforeIndicators, indicator.(*Indicator))
		} else {
			// 从数据库中查找
			param := map[string]interface{}{
				"indicator_code": v,
			}
			indicators, err := ind.Find(param)
			if nil == err && 0 < len(indicators) {
				beforeIndicators = append(beforeIndicators, indicators[0])
				zlog.Infof("从数据库中查找指标[%v(%v)]", nil, indicators[0].IndicatorName, indicators[0].IndicatorCode)
				util.PutCacheByCacheName(util.RPM_INDICATOR_CACHE, indicators[0].IndicatorCode, indicators[0], 0)
			}
		}

	}
	ind.IndicatorBeforeRely = beforeIndicators
}

var indicatorTables string = "RPM_PAR_BASE_INDICATORS T"

var indicatorCols = map[string]string{
	"T.UUID":                    "''",
	"T.INDICATOR_CODE":          "' '",
	"T.INDICATOR_NAME":          "' '",
	"T.INDICATOR_TYPE":          "' '",
	"T.INDICATOR_BEFORE_RELY":   "' '",
	"T.FORMULAE_CODE":           "' '",
	"T.FORMULAE_CN":             "' '",
	"T.INDICATOR_BUSINESS_TYPE": "' '",
	"T.FLAG":                    "' '",
	"T.STATUS":                  "' '",
	"T.METHOD":                  "' '",
	"T.CREATE_TIME":             "SYSDATE",
	"T.UPDATE_TIME":             "SYSDATE",
	"T.CREATE_USER":             "' '",
	"T.UPDATE_USER":             "' '",
}

var indicatorColSort = []string{
	"T.UUID",
	"T.INDICATOR_CODE",
	"T.INDICATOR_NAME",
	"T.INDICATOR_TYPE",
	"T.INDICATOR_BEFORE_RELY",
	"T.FORMULAE_CODE",
	"T.FORMULAE_CN",
	"T.INDICATOR_BUSINESS_TYPE",
	"T.FLAG",
	"T.STATUS",
	"T.METHOD",
	"T.CREATE_TIME",
	"T.UPDATE_TIME",
	"T.CREATE_USER",
	"T.UPDATE_USER",
}

func init() {
	zlog.Infof("初始化指标缓存", nil)
	util.FlushCacheByCacheName(util.RPM_INDICATOR_CACHE)
	var ind Indicator
	paramMap := map[string]interface{}{
		"flag": util.FLAG_TRUE,
	}
	indicators, _ := ind.Find(paramMap)
	for _, indicator := range indicators {
		util.PutCacheByCacheName(util.RPM_INDICATOR_CACHE, indicator.IndicatorCode, indicator, 0)
	}
}
