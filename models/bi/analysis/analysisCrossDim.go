package analysis

import (
	"database/sql"
	"fmt"
	"time"

	"pccqcpa.com.cn/app/rpm/api/models/charts"
	modelsUtil "pccqcpa.com.cn/app/rpm/api/models/util"
	"pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/components/zlog"
)

// 结构体对应数据库表【RPM_BI_ANALYSIS_CROSS_DIM】贷款交叉分析-多维度
// by author Jason
// by time 2016-11-21 16:26:56
/**
 * @apiDefine AnalysisCrossDim
 * @apiSuccess {string}     UUID            主键默认值sys_guid()
 * @apiSuccess {string}   	DimOneCode      维度一码值
 * @apiSuccess {string}   	DimOneName      维度一名称
 * @apiSuccess {string}   	DimOneSubCode   维度一子码值
 * @apiSuccess {string}   	DimOneSubName   维度一子名称
 * @apiSuccess {string}   	DimTwoCode      维度二码值
 * @apiSuccess {string}   	DimTwoName      维度二名称
 * @apiSuccess {string}   	DimTwoSubCode   维度二子码值
 * @apiSuccess {string}   	DimTwoSubName   维度二子名称
 * @apiSuccess {float64}  	DimValue        交叉纬度值
 * @apiSuccess {time.Time}	CreateTime      创建时间
 * @apiSuccess {string}   	CreateUser      创建人
 * @apiSuccess {time.Time}	UpdateTime      更新时间
 * @apiSuccess {string}   	UpdateUser      更新人
 */
type AnalysisCrossDim struct {
	UUID          string    // 主键默认值sys_guid()
	DimOneCode    string    // 维度一码值
	DimOneName    string    // 维度一名称
	DimOneSubCode string    // 维度一子码值
	DimOneSubName string    // 维度一子名称
	DimTwoCode    string    // 维度二码值
	DimTwoName    string    // 维度二名称
	DimTwoSubCode string    // 维度二子码值
	DimTwoSubName string    // 维度二子名称
	DimValue      float64   // 交叉纬度值
	CreateTime    time.Time // 创建时间
	CreateUser    string    // 创建人
	UpdateTime    time.Time // 更新时间
	UpdateUser    string    // 更新人
}

// 结构体scan
// by author Jason
// by time 2016-11-21 16:26:56
func (this AnalysisCrossDim) scan(rows *sql.Rows) (*AnalysisCrossDim, error) {
	var one = new(AnalysisCrossDim)
	values := []interface{}{
		&one.UUID,
		&one.DimOneCode,
		&one.DimOneName,
		&one.DimOneSubCode,
		&one.DimOneSubName,
		&one.DimTwoCode,
		&one.DimTwoName,
		&one.DimTwoSubCode,
		&one.DimTwoSubName,
		&one.DimValue,
		&one.CreateTime,
		&one.CreateUser,
		&one.UpdateTime,
		&one.UpdateUser,
	}
	err := util.OracleScan(rows, values)
	return one, err
}

// 分页操作
// by author Jason
// by time 2016-11-21 16:26:56
func (this AnalysisCrossDim) List(param ...map[string]interface{}) (*util.PageData, error) {
	pageData, rows, err := modelsUtil.List(analysisCrossDimTabales, analysisCrossDimCols, analysisCrossDimColsSort, param...)
	defer rows.Close()
	if nil != err {
		er := fmt.Errorf("分页查询贷款交叉分析-多维度信息出错")
		zlog.Error(er.Error(), err)
		return nil, er
	}

	var rst []*AnalysisCrossDim
	for rows.Next() {
		one, err := this.scan(rows)
		if nil != err {
			er := fmt.Errorf("分页查询贷款交叉分析-多维度信息row.Scan()出错")
			zlog.Error(er.Error(), err)
			return nil, er
		}
		rst = append(rst, one)
	}
	pageData.Rows = rst
	return pageData, nil
}

// 多参数查询
// by author Jason
// by time 2016-11-21 16:26:56
func (this AnalysisCrossDim) Find(param ...map[string]interface{}) (*charts.AnalysisCrossDimChart, error) {
	rows, err := modelsUtil.FindRows(analysisCrossDimTabales, analysisCrossDimCols, analysisCrossDimColsSort, param...)
	defer rows.Close()
	if nil != err {
		er := fmt.Errorf("多参数查询贷款交叉分析-多维度信息出错")
		zlog.Error(er.Error(), err)
		return nil, er
	}

	var mapSeries = make(map[string]charts.Series) // 纵坐标数据，先放到map中，后循环到数组中
	var compareMap = make(map[string]int)          // 用于比较增加横坐标
	var categories []string                        // 横坐标
	var i = 0

	for rows.Next() {
		one, err := this.scan(rows)
		if nil != err {
			er := fmt.Errorf("多参数查询贷款交叉分析-多维度信息row.Scan()出错")
			zlog.Error(er.Error(), err)
			return nil, er
		}
		// 逻辑处理 	横坐标
		// fmt.Printf("横坐标【%s】【%s】比较值【%s】\n", one.DimOneSubCode, one.DimTwoCode, compareCode)
		var comparekey = one.DimOneSubCode + one.DimOneSubName
		index, ok := compareMap[comparekey]
		if !ok {
			// 横坐标处理
			categories = append(categories, one.DimOneSubName)
			compareMap[comparekey] = i
			index = i
			i++
		}

		// 纵坐标处理
		key := one.DimTwoSubCode + one.DimTwoSubName
		series, ok := mapSeries[key]
		if ok {
			// 如果存在，判断Data的长度与index+1的大小
			// 如果Data的长度大于等于index+1，则直接赋值
			// 如果Data的长度小于index+1，则make一个长度为inde+1 - len(Data)长度的slice，再append
			if len(series.Data) >= index+1 {
				series.Data[index] = one.DimValue
			} else {
				series.Data = append(series.Data, make([]float64, index-len(series.Data))...)
				series.Data = append(series.Data, one.DimValue)
			}
		} else {
			// 如果不存在，则创建一个slice，cap为index+1，设置下标index值
			var data = make([]float64, index+1)
			data[index] = one.DimValue
			series = charts.Series{
				Name: one.DimTwoSubName,
				Data: data,
			}
		}
		mapSeries[key] = series

		// fmt.Println("categories : ", categories)
		// fmt.Println("mapSeries : ", mapSeries)
	}
	var rst charts.AnalysisCrossDimChart
	rst.Categories = categories
	for _, val := range mapSeries {
		if len(val.Data) < i {
			val.Data = append(val.Data, make([]float64, i-len(val.Data))...)
		}
		rst.Series = append(rst.Series, val)
	}
	return &rst, nil
}

// 新增贷款交叉分析-多维度信息
// by author Jason
// by time 2016-11-21 16:26:56
func (this AnalysisCrossDim) Add() error {
	paramMap := map[string]interface{}{
		"dim_one_code":     this.DimOneCode,
		"dim_one_name":     this.DimOneName,
		"dim_one_sub_code": this.DimOneSubCode,
		"dim_one_sub_name": this.DimOneSubName,
		"dim_two_code":     this.DimTwoCode,
		"dim_two_name":     this.DimTwoName,
		"dim_two_sub_code": this.DimTwoSubCode,
		"dim_two_sub_name": this.DimTwoSubName,
		"dim_value":        this.DimValue,
		"create_time":      util.GetCurrentTime(),
		"create_user":      this.CreateUser,
		"update_time":      util.GetCurrentTime(),
		"update_user":      this.UpdateUser,
	}
	err := util.OracleAdd("RPM_BI_ANALYSIS_CROSS_DIM", paramMap)
	if nil != err {
		er := fmt.Errorf("新增贷款交叉分析-多维度信息出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

// 更新贷款交叉分析-多维度信息
// by author Jason
// by time 2016-11-21 16:26:56
func (this AnalysisCrossDim) Update() error {
	paramMap := map[string]interface{}{
		"dim_one_code":     this.DimOneCode,
		"dim_one_name":     this.DimOneName,
		"dim_one_sub_code": this.DimOneSubCode,
		"dim_one_sub_name": this.DimOneSubName,
		"dim_two_code":     this.DimTwoCode,
		"dim_two_name":     this.DimTwoName,
		"dim_two_sub_code": this.DimTwoSubCode,
		"dim_two_sub_name": this.DimTwoSubName,
		"dim_value":        this.DimValue,
		"update_time":      util.GetCurrentTime(),
		"update_user":      this.UpdateUser,
	}
	whereParam := map[string]interface{}{
		"uuid": this.UUID,
	}
	err := util.OracleUpdate("RPM_BI_ANALYSIS_CROSS_DIM", paramMap, whereParam)
	if nil != err {
		er := fmt.Errorf("更新贷款交叉分析-多维度信息出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

// 删除贷款交叉分析-多维度信息
// by author Jason
// by time 2016-11-21 16:26:56
func (this AnalysisCrossDim) Delete() error {
	whereParam := map[string]interface{}{
		"uuid": this.UUID,
	}
	err := util.OracleDelete("RPM_BI_ANALYSIS_CROSS_DIM", whereParam)
	if nil != err {
		er := fmt.Errorf("删除贷款交叉分析-多维度信息出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

var analysisCrossDimTabales string = `RPM_BI_ANALYSIS_CROSS_DIM T`

var analysisCrossDimCols map[string]string = map[string]string{
	"T.UUID":             "' '",
	"T.DIM_ONE_CODE":     "' '",
	"T.DIM_ONE_NAME":     "' '",
	"T.DIM_ONE_SUB_CODE": "' '",
	"T.DIM_ONE_SUB_NAME": "' '",
	"T.DIM_TWO_CODE":     "' '",
	"T.DIM_TWO_NAME":     "' '",
	"T.DIM_TWO_SUB_CODE": "' '",
	"T.DIM_TWO_SUB_NAME": "' '",
	"T.DIM_VALUE":        "0",
	"T.CREATE_TIME":      "sysdate",
	"T.CREATE_USER":      "' '",
	"T.UPDATE_TIME":      "sysdate",
	"T.UPDATE_USER":      "' '",
}

var analysisCrossDimColsSort = []string{
	"T.UUID",
	"T.DIM_ONE_CODE",
	"T.DIM_ONE_NAME",
	"T.DIM_ONE_SUB_CODE",
	"T.DIM_ONE_SUB_NAME",
	"T.DIM_TWO_CODE",
	"T.DIM_TWO_NAME",
	"T.DIM_TWO_SUB_CODE",
	"T.DIM_TWO_SUB_NAME",
	"T.DIM_VALUE",
	"T.CREATE_TIME",
	"T.CREATE_USER",
	"T.UPDATE_TIME",
	"T.UPDATE_USER",
}
