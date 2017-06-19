package par

import (
	"database/sql"
	"fmt"
	"time"

	"pccqcpa.com.cn/app/rpm/api/models/dim"
	"pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/components/zlog"

	modelsUtil "pccqcpa.com.cn/app/rpm/api/models/util"
)

// 结构体对应数据库表【RPM_PAR_SCENE_ITD_YIELD】派生中间收入收益率
/**
 * @apiDefine SceneItdYield
 * @apiSuccess {string}     UUID            主键默认值sys_guid()
 * @apiSuccess {float64}   	ItdYield    	中收收益率
 * @apiSuccess {string}   	Flag      		生效标志
 * @apiSuccess {time.Time}	CreateTime      创建时间
 * @apiSuccess {string}   	CreateUser      创建人
 * @apiSuccess {time.Time}	UpdateTime      更新时间
 * @apiSuccess {string}   	UpdateUser      更新人
 */
type SceneItdYield struct {
	UUID       string      // 主键
	Product    dim.Product // 产品
	ItdYield   float64     // 中收收益率
	Flag       string      // 生效标志
	CreateTime time.Time   // 创建时间
	UpdateTime time.Time   // 更新时间
	CreateUser string      // 创建人
	UpdateUser string      // 更新人
}

//结构体scan
func (s *SceneItdYield) scan(rows *sql.Rows) (*SceneItdYield, error) {
	var sceneItdYield = new(SceneItdYield)
	values := []interface{}{
		&sceneItdYield.UUID,
		&sceneItdYield.ItdYield,
		&sceneItdYield.Flag,
		&sceneItdYield.CreateTime,
		&sceneItdYield.UpdateTime,
		&sceneItdYield.CreateUser,
		&sceneItdYield.UpdateUser,

		&sceneItdYield.Product.ProductCode,
		&sceneItdYield.Product.ProductName,
	}
	err := util.OracleScan(rows, values)
	return sceneItdYield, err
}

// 分页操作
// by author Yeqc
// by time 2016-12-21 10:30:46
func (this SceneItdYield) List(param ...map[string]interface{}) (*util.PageData, error) {
	pageData, rows, err := modelsUtil.List(sceneItdYieldTables, sceneItdYieldCols, sceneItdYieldColsSort, param...)
	defer rows.Close()
	if nil != err {
		er := fmt.Errorf("分页查询派生中间收入收益率信息出错")
		zlog.Error(er.Error(), err)
		return nil, er
	}

	var rst []*SceneItdYield
	for rows.Next() {
		one, err := this.scan(rows)
		if nil != err {
			er := fmt.Errorf("分页查询派生中间收入收益率信息row.Scan()出错")
			zlog.Error(er.Error(), err)
			return nil, er
		}
		rst = append(rst, one)
	}
	pageData.Rows = rst
	return pageData, nil
}

// 多参数查询
func (s *SceneItdYield) Find(param ...map[string]interface{}) ([]*SceneItdYield, error) {
	rows, err := modelsUtil.FindRows(sceneItdYieldTables, sceneItdYieldCols, sceneItdYieldColsSort, param...)
	defer rows.Close()
	if nil != err {
		er := fmt.Errorf("查询中收收益率出错")
		zlog.Error(er.Error(), err)
		return nil, er
	}
	var sceneItdYields []*SceneItdYield
	for rows.Next() {
		sceneItdYield, err := s.scan(rows)
		if nil != err {
			er := fmt.Errorf("查询中收收益率rows.Scan()出错")
			zlog.Error(er.Error(), err)
			return nil, er
		}
		sceneItdYields = append(sceneItdYields, sceneItdYield)
	}
	return sceneItdYields, nil
}

// 新增派生中间收入收益率信息
// by author Yeqc
// by time 2016-12-21 10:30:46
func (this SceneItdYield) Add() error {
	// fmt.Println("---------", this.ItdYield)
	paramMap := map[string]interface{}{
		"product":     this.Product.ProductCode,
		"itd_yield":   this.ItdYield,
		"flag":        this.Flag,
		"create_time": util.GetCurrentTime(),
		"update_time": util.GetCurrentTime(),
		"create_user": this.CreateUser,
		"update_user": this.UpdateUser,
	}
	err := util.OracleAdd("RPM_PAR_SCENE_ITD_YIELD", paramMap)
	if nil != err {
		er := fmt.Errorf("新增派生中间收入收益率信息出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

// 更新派生中间收入收益率信息
// by author Yeqc
// by time 2016-12-21 10:30:46
func (this SceneItdYield) Update() error {
	paramMap := map[string]interface{}{
		"product":     this.Product.ProductCode,
		"itd_yield":   this.ItdYield,
		"flag":        this.Flag,
		"update_time": util.GetCurrentTime(),
		"update_user": this.UpdateUser,
	}
	whereParam := map[string]interface{}{
		"uuid": this.UUID,
	}
	err := util.OracleUpdate("RPM_PAR_SCENE_ITD_YIELD", paramMap, whereParam)
	if nil != err {
		er := fmt.Errorf("更新派生中间收入收益率信息出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

// 删除派生中间收入收益率信息
// by author Yeqc
// by time 2016-12-21 10:30:46
func (this SceneItdYield) Delete() error {
	whereParam := map[string]interface{}{
		"uuid": this.UUID,
	}
	err := util.OracleDelete("RPM_PAR_SCENE_ITD_YIELD", whereParam)
	if nil != err {
		er := fmt.Errorf("删除派生中间收入收益率信息出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

var sceneItdYieldTables string = `
			 RPM_PAR_SCENE_ITD_YIELD T
		LEFT JOIN RPM_DIM_PRODUCT PRODUCT
		  ON (T.PRODUCT = PRODUCT.PRODUCT_CODE)`

var sceneItdYieldCols = map[string]string{
	"T.UUID":        "' '",
	"T.ITD_YIELD":   "0",
	"T.FLAG":        "' '",
	"T.CREATE_TIME": "sysdate",
	"T.UPDATE_TIME": "sysdate",
	"T.CREATE_USER": "' '",
	"T.UPDATE_USER": "' '",

	"PRODUCT.PRODUCT_CODE": "' '",
	"PRODUCT.PRODUCT_NAME": "' '",
}

var sceneItdYieldColsSort = []string{
	"T.UUID",
	"T.ITD_YIELD",
	"T.FLAG",
	"T.CREATE_TIME",
	"T.UPDATE_TIME",
	"T.CREATE_USER",
	"T.UPDATE_USER",

	"PRODUCT.PRODUCT_CODE",
	"PRODUCT.PRODUCT_NAME",
}
