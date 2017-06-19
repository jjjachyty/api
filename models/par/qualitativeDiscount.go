package par

import (
	"database/sql"
	"fmt"
	"time"

	modelsUtil "pccqcpa.com.cn/app/rpm/api/models/util"
	"pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/components/zlog"
)

// 结构体对应数据库表【RPM_PAR_QUALITATIVE_DISCOUNT】定性优惠点数
// by author Yeqc
// by time 2016-12-21 10:20:03
/**
 * @apiDefine QualitativeDiscount
 * @apiSuccess {string}     UUID            	主键默认值sys_guid()
 * @apiSuccess {float64}   	UseProduct    		使用产品数
 * @apiSuccess {float64}   	CooperationPeriod   合作年限数
 * @apiSuccess {float64}   	Discount   			优惠点数
 * @apiSuccess {string}   	Flag   				生效标志
 * @apiSuccess {time.Time}	CreateTime      	创建时间
 * @apiSuccess {string}   	CreateUser      	创建人
 * @apiSuccess {time.Time}	UpdateTime      	更新时间
 * @apiSuccess {string}   	UpdateUser      	更新人
 */
type QualitativeDiscount struct {
	UUID           string    // 主键
	StockSceneType string    // 存量优惠类型
	StockSceneVal  float64   // 存量优惠值
	Discount       float64   // 优惠点数
	Flag           string    // 生效标志
	CreateTime     time.Time // 创建时间
	UpdateTime     time.Time // 更新时间
	CreateUser     string    // 创建人
	UpdateUser     string    // 更新人
}

// 结构体scan
// by author Yeqc
// by time 2016-12-21 10:20:03
func (this QualitativeDiscount) scan(rows *sql.Rows) (*QualitativeDiscount, error) {
	var one = new(QualitativeDiscount)
	values := []interface{}{
		&one.UUID,
		&one.StockSceneType,
		&one.StockSceneVal,
		&one.Discount,
		&one.Flag,
		&one.CreateTime,
		&one.UpdateTime,
		&one.CreateUser,
		&one.UpdateUser,
	}
	err := util.OracleScan(rows, values)
	return one, err
}

// 分页操作
// by author Yeqc
// by time 2016-12-21 10:20:03
func (this QualitativeDiscount) List(param ...map[string]interface{}) (*util.PageData, error) {
	pageData, rows, err := modelsUtil.List(qualitativeDiscountTabales, qualitativeDiscountCols, qualitativeDiscountColsSort, param...)
	defer rows.Close()
	if nil != err {
		er := fmt.Errorf("分页查询定性优惠点数信息出错")
		zlog.Error(er.Error(), err)
		return nil, er
	}

	var rst []*QualitativeDiscount
	for rows.Next() {
		one, err := this.scan(rows)
		if nil != err {
			er := fmt.Errorf("分页查询定性优惠点数信息row.Scan()出错")
			zlog.Error(er.Error(), err)
			return nil, er
		}
		rst = append(rst, one)
	}
	pageData.Rows = rst
	return pageData, nil
}

// 多参数查询
// by author Yeqc
// by time 2016-12-21 10:20:03
func (this QualitativeDiscount) Find(param ...map[string]interface{}) ([]*QualitativeDiscount, error) {
	rows, err := modelsUtil.FindRows(qualitativeDiscountTabales, qualitativeDiscountCols, qualitativeDiscountColsSort, param...)
	defer rows.Close()
	if nil != err {
		er := fmt.Errorf("多参数查询定性优惠点数信息出错")
		zlog.Error(er.Error(), err)
		return nil, er
	}
	var rst []*QualitativeDiscount
	for rows.Next() {
		one, err := this.scan(rows)
		if nil != err {
			er := fmt.Errorf("多参数查询定性优惠点数信息row.Scan()出错")
			zlog.Error(er.Error(), err)
			return nil, er
		}
		rst = append(rst, one)
	}
	return rst, nil
}

// 新增定性优惠点数信息
// by author Yeqc
// by time 2016-12-21 10:20:03
func (this QualitativeDiscount) Add() error {
	paramMap := map[string]interface{}{
		"stock_scene_type": this.StockSceneType,
		"stock_scene_val":  this.StockSceneVal,
		"discount":         this.Discount,
		"flag":             this.Flag,
		"create_time":      util.GetCurrentTime(),
		"update_time":      util.GetCurrentTime(),
		"create_user":      this.CreateUser,
		"update_user":      this.UpdateUser,
	}
	err := util.OracleAdd("RPM_PAR_QUALITATIVE_DISCOUNT", paramMap)
	if nil != err {
		er := fmt.Errorf("新增定性优惠点数信息出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

// 更新定性优惠点数信息
// by author Yeqc
// by time 2016-12-21 10:20:03
func (this QualitativeDiscount) Update() error {
	paramMap := map[string]interface{}{
		"stock_scene_type": this.StockSceneType,
		"stock_scene_val":  this.StockSceneVal,
		"discount":         this.Discount,
		"flag":             this.Flag,
		"update_time":      util.GetCurrentTime(),
		"update_user":      this.UpdateUser,
	}
	whereParam := map[string]interface{}{
		"uuid": this.UUID,
	}
	err := util.OracleUpdate("RPM_PAR_QUALITATIVE_DISCOUNT", paramMap, whereParam)
	if nil != err {
		er := fmt.Errorf("更新定性优惠点数信息出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

// 删除定性优惠点数信息
// by author Yeqc
// by time 2016-12-21 10:20:03
func (this QualitativeDiscount) Delete() error {
	whereParam := map[string]interface{}{
		"uuid": this.UUID,
	}
	err := util.OracleDelete("RPM_PAR_QUALITATIVE_DISCOUNT", whereParam)
	if nil != err {
		er := fmt.Errorf("删除定性优惠点数信息出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

var qualitativeDiscountTabales string = `RPM_PAR_QUALITATIVE_DISCOUNT T`

var qualitativeDiscountCols map[string]string = map[string]string{
	"T.UUID":             "' '",
	"T.STOCK_SCENE_TYPE": "' '",
	"T.STOCK_SCENE_VAL":  "0",
	"T.DISCOUNT":         "0",
	"T.FLAG":             "' '",
	"T.CREATE_TIME":      "sysdate",
	"T.UPDATE_TIME":      "sysdate",
	"T.CREATE_USER":      "' '",
	"T.UPDATE_USER":      "' '",
}

var qualitativeDiscountColsSort = []string{
	"T.UUID",
	"T.STOCK_SCENE_TYPE",
	"T.STOCK_SCENE_VAL",
	"T.DISCOUNT",
	"T.FLAG",
	"T.CREATE_TIME",
	"T.UPDATE_TIME",
	"T.CREATE_USER",
	"T.UPDATE_USER",
}
