package par

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/components/zlog"

	modelsUtil "pccqcpa.com.cn/app/rpm/api/models/util"
)

// 派生调节系数
// TABLE：RPM_PAR_SCENE_DISCOUNT_ADJ
/**
 * @apiDefine SceneDiscountAdj
 * @apiSuccess {string}     UUID            主键默认值sys_guid()
 * @apiSuccess {string}   	BizType    		业务类型(1存量,2派生)
 * @apiSuccess {float64}   	GapProportion   实现缺口比例（实际eva/预期eva）
 * @apiSuccess {float64}   	AdjValue   		调节系数
 * @apiSuccess {string}   	Flag    		生效标志
 * @apiSuccess {time.Time}	CreateTime      创建时间
 * @apiSuccess {string}   	CreateUser      创建人
 * @apiSuccess {time.Time}	UpdateTime      更新时间
 * @apiSuccess {string}   	UpdateUser      更新人
 */
type SceneDiscountAdj struct {
	UUID          string    // 主键
	BizType       string    // 业务类型(1存量,2派生)
	GapProportion float64   // 实现缺口比例（实际eva/预期eva）
	AdjValue      float64   // 调节系数
	Flag          string    // 生效标志
	CreateTime    time.Time // 创建时间
	UpdateTime    time.Time // 更新时间
	CreateUser    string    // 创建人
	UpdateUser    string    // 更新人
}

func (d SceneDiscountAdj) scan(rows *sql.Rows) (*SceneDiscountAdj, error) {
	var sceneDiscountAdj = new(SceneDiscountAdj)
	values := []interface{}{
		&sceneDiscountAdj.UUID,
		&sceneDiscountAdj.BizType,
		&sceneDiscountAdj.GapProportion,
		&sceneDiscountAdj.AdjValue,
		&sceneDiscountAdj.Flag,
		&sceneDiscountAdj.CreateTime,
		&sceneDiscountAdj.UpdateTime,
		&sceneDiscountAdj.CreateUser,
		&sceneDiscountAdj.UpdateUser,
	}
	err := util.OracleScan(rows, values)
	return sceneDiscountAdj, err
}

func (s SceneDiscountAdj) List(param ...map[string]interface{}) (*util.PageData, error) {
	pageData, rows, err := modelsUtil.List(sceneDiscountAjdTables, sceneDiscountAjdCols, sceneDiscountAjdColsSort, param...)
	defer rows.Close()
	if nil != err {
		er := errors.New("分页查询派生调节系数出错")
		zlog.Error(er.Error(), err)
		return nil, er
	}
	var sceneDiscountAdjs []*SceneDiscountAdj
	for rows.Next() {
		sceneDiscountAdj, err := s.scan(rows)
		if nil != err {
			er := fmt.Errorf("查询派生调节系数row.Scan()出错")
			zlog.Error(er.Error(), err)
			return nil, er
		}
		sceneDiscountAdjs = append(sceneDiscountAdjs, sceneDiscountAdj)
	}
	pageData.Rows = sceneDiscountAdjs
	return pageData, nil
}

// 多参数查询
func (d SceneDiscountAdj) Find(param ...map[string]interface{}) ([]*SceneDiscountAdj, error) {
	rows, err := modelsUtil.FindRows(sceneDiscountAjdTables, sceneDiscountAjdCols, sceneDiscountAjdColsSort, param...)
	defer rows.Close()
	if nil != err {
		er := fmt.Errorf("查询派生调节系数出错")
		zlog.Error(er.Error(), err)
		return nil, er
	}
	var sceneDiscountAdjs []*SceneDiscountAdj
	for rows.Next() {
		sceneDiscountAdj, err := d.scan(rows)
		if nil != err {
			er := fmt.Errorf("查询派生调节系数row.Scan()出错")
			zlog.Error(er.Error(), err)
			return nil, er
		}
		sceneDiscountAdjs = append(sceneDiscountAdjs, sceneDiscountAdj)
	}
	return sceneDiscountAdjs, nil
}

// 新增
func (s SceneDiscountAdj) Add() error {
	paramMap := map[string]interface{}{
		"BIZ_TYPE":       s.BizType,
		"GAP_PROPORTION": s.GapProportion,
		"ADJ_VALUE":      s.AdjValue,
		"FLAG":           s.Flag,
		"CREATE_TIME":    util.GetCurrentTime(),
		"UPDATE_TIME":    util.GetCurrentTime(),
		"CREATE_USER":    s.CreateUser,
		"UPDATE_USER":    s.UpdateUser,
	}
	err := util.OracleAdd(sceneDiscountAjdTables, paramMap)
	if nil != err {
		er := errors.New("新增派生调节系数出错")
		return er
	}
	return nil
}

// 更新
func (s SceneDiscountAdj) Update() error {
	paramMap := map[string]interface{}{
		"BIZ_TYPE":       s.BizType,
		"GAP_PROPORTION": s.GapProportion,
		"ADJ_VALUE":      s.AdjValue,
		"FLAG":           s.Flag,
		"UPDATE_TIME":    util.GetCurrentTime(),
		"UPDATE_USER":    s.UpdateUser,
	}
	whereParamMap := map[string]interface{}{
		"UUID": s.UUID,
	}
	err := util.OracleUpdate(sceneDiscountAjdTables, paramMap, whereParamMap)
	if nil != err {
		er := errors.New("更新派生调节系数出错")
		return er
	}
	return nil
}

// 删除
func (s SceneDiscountAdj) Delete() error {
	whereParamMap := map[string]interface{}{
		"UUID": s.UUID,
	}
	err := util.OracleDelete(sceneDiscountAjdTables, whereParamMap)
	if nil != err {
		er := errors.New("删除派生调节系数出错")
		return er
	}
	return nil
}

var sceneDiscountAjdTables string = "RPM_PAR_SCENE_DISCOUNT_ADJ T"

var sceneDiscountAjdCols = map[string]string{
	"T.UUID":           "' '",
	"T.BIZ_TYPE":       "' '",
	"T.GAP_PROPORTION": "0",
	"T.ADJ_VALUE":      "0",
	"T.FLAG":           "' '",
	"T.CREATE_TIME":    "sysdate",
	"T.UPDATE_TIME":    "sysdate",
	"T.CREATE_USER":    "' '",
	"T.UPDATE_USER":    "' '",
}

var sceneDiscountAjdColsSort = []string{
	"T.UUID",
	"T.BIZ_TYPE",
	"T.GAP_PROPORTION",
	"T.ADJ_VALUE",
	"T.FLAG",
	"T.CREATE_TIME",
	"T.UPDATE_TIME",
	"T.CREATE_USER",
	"T.UPDATE_USER",
}
