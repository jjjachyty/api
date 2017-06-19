package par

import (
	"database/sql"
	"time"

	"fmt"

	"pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/components/zlog"

	modelsUtil "pccqcpa.com.cn/app/rpm/api/models/util"
)

/**
 * @apiDefine LgdBasel
 * @apiSuccess {string}     UUID            主键默认值sys_guid()
 * @apiSuccess {string}   	PledgeType    	抵押品类型
 * @apiSuccess {float64}   	LowLgd      	最低违约损失率
 * @apiSuccess {float64}   	HighPledgeRadio 最低抵质押水平
 * @apiSuccess {float64}   	LowPledgeRadio  最高抵质押水平
 * @apiSuccess {string}   	StartTime      	生效日期
 * @apiSuccess {string}   	Flag      		生效标志
 * @apiSuccess {time.Time}	CreateTime      创建时间
 * @apiSuccess {string}   	CreateUser      创建人
 * @apiSuccess {time.Time}	UpdateTime      更新时间
 * @apiSuccess {string}   	UpdateUser      更新人
 */
type LgdBasel struct {
	UUID            string
	PledgeType      string    //抵押品类型
	LowLgd          float64   //最低违约损失率
	HighPledgeRadio float64   //最高抵质押水平
	LowPledgeRadio  float64   //最低抵质押水平
	StartTime       time.Time //生效日期
	CreateTime      time.Time //创建日期
	CreateUser      string    //创建人
	UpdateTime      time.Time //更新日期
	UpdateUser      string    //更新人
	Flag            string    //生效标志
}

var lgdBaselTables string = `RPM_PAR_LGD_BASEL T`

var lgdBaselCols = map[string]string{
	"T.UUID":              "' '",
	"T.PLEDGE_TYPE":       "' '",
	"T.LOW_LGD":           "0",
	"T.HIGH_PLEDGE_RADIO": "0",
	"T.LOW_PLEDGE_RADIO":  "0",
	"T.START_TIME":        "sysdate",
	"T.CREATE_TIME":       "sysdate",
	"T.CREATE_USER":       "' '",
	"T.UPDATE_TIME":       "sysdate",
	"T.UPDATE_USER":       "' '",
	"T.FLAG":              "' '",
}

var lgdBaselSortCols = []string{
	"T.UUID",
	"T.PLEDGE_TYPE",
	"T.LOW_LGD",
	"T.HIGH_PLEDGE_RADIO",
	"T.LOW_PLEDGE_RADIO",
	"T.START_TIME",
	"T.CREATE_TIME",
	"T.CREATE_USER",
	"T.UPDATE_TIME",
	"T.UPDATE_USER",
	"T.FLAG",
}

func (this *LgdBasel) scanLgdBasel(rows *sql.Rows) (*LgdBasel, error) {
	var lgdBasel = new(LgdBasel)
	var values = []interface{}{
		&lgdBasel.UUID,
		&lgdBasel.PledgeType,
		&lgdBasel.LowLgd,
		&lgdBasel.HighPledgeRadio,
		&lgdBasel.LowPledgeRadio,
		&lgdBasel.StartTime,
		&lgdBasel.CreateTime,
		&lgdBasel.CreateUser,
		&lgdBasel.UpdateTime,
		&lgdBasel.UpdateUser,
		&lgdBasel.Flag,
	}
	err := util.OracleScan(rows, values)
	if nil != err {
		return nil, err
	}
	return lgdBasel, nil
}

func (this *LgdBasel) SelectLgdBaselByParams(paramMap map[string]interface{}) ([]*LgdBasel, error) {
	rows, err := modelsUtil.FindRows(lgdBaselTables, lgdBaselCols, lgdBaselSortCols, paramMap)
	if nil != err {
		zlog.Error("多参数查询违约损失率错误", err)
		return nil, err
	}
	defer rows.Close()
	var lgdBasels []*LgdBasel
	for rows.Next() {
		lgdBasel, err := this.scanLgdBasel(rows)
		if nil != err {
			zlog.Error("多参数查询违约损失率rows.Scan()错误", err)
			return nil, err
		}
		lgdBasels = append(lgdBasels, lgdBasel)
	}
	return lgdBasels, nil
}

// 分页操作
// by author Yeqc
// by time 2016-12-19 10:16:29
func (this LgdBasel) List(param ...map[string]interface{}) (*util.PageData, error) {
	pageData, rows, err := modelsUtil.List(lgdBaselTables, lgdBaselCols, lgdBaselSortCols, param...)
	defer rows.Close()
	if nil != err {
		er := fmt.Errorf("分页查询违约损失率信息出错")
		zlog.Error(er.Error(), err)
		return nil, er
	}

	var rst []*LgdBasel
	for rows.Next() {
		one, err := this.scanLgdBasel(rows)
		if nil != err {
			er := fmt.Errorf("分页查询违约损失率信息row.Scan()出错")
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
// by time 2016-12-19 10:16:29
func (this LgdBasel) Find(param ...map[string]interface{}) ([]*LgdBasel, error) {
	rows, err := modelsUtil.FindRows(lgdBaselTables, lgdBaselCols, lgdBaselSortCols, param...)
	defer rows.Close()
	if nil != err {
		er := fmt.Errorf("多参数查询违约损失率信息出错")
		zlog.Error(er.Error(), err)
		return nil, er
	}
	var rst []*LgdBasel
	for rows.Next() {
		one, err := this.scanLgdBasel(rows)
		if nil != err {
			er := fmt.Errorf("多参数查询违约损失率信息row.Scan()出错")
			zlog.Error(er.Error(), err)
			return nil, er
		}
		rst = append(rst, one)
	}
	return rst, nil
}

// 新增违约损失率信息
// by author Yeqc
// by time 2016-12-19 10:16:29
func (this LgdBasel) Add() error {
	paramMap := map[string]interface{}{
		"pledge_type":       this.PledgeType,
		"low_lgd":           this.LowLgd,
		"high_pledge_radio": this.HighPledgeRadio,
		"low_pledge_radio":  this.LowPledgeRadio,
		"start_time":        this.StartTime,
		"create_time":       util.GetCurrentTime(),
		"create_user":       this.CreateUser,
		"update_time":       util.GetCurrentTime(),
		"update_user":       this.CreateUser,
		"flag":              this.Flag,
	}
	err := util.OracleAdd("RPM_PAR_LGD_BASEL", paramMap)
	if nil != err {
		er := fmt.Errorf("新增违约损失率信息出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

// 更新违约损失率信息
// by author Yeqc
// by time 2016-12-19 10:16:29
func (this LgdBasel) Update() error {
	paramMap := map[string]interface{}{
		"pledge_type":       this.PledgeType,
		"low_lgd":           this.LowLgd,
		"high_pledge_radio": this.HighPledgeRadio,
		"low_pledge_radio":  this.LowPledgeRadio,
		"start_time":        this.StartTime,
		"update_time":       util.GetCurrentTime(),
		"update_user":       this.UpdateUser,
		"flag":              this.Flag,
	}
	whereParam := map[string]interface{}{
		"uuid": this.UUID,
	}
	err := util.OracleUpdate("RPM_PAR_LGD_BASEL", paramMap, whereParam)
	if nil != err {
		er := fmt.Errorf("更新违约损失率信息出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

// 删除违约损失率信息
// by author Yeqc
// by time 2016-12-19 10:16:29
func (this LgdBasel) Delete() error {
	whereParam := map[string]interface{}{
		"uuid": this.UUID,
	}
	err := util.OracleDelete("RPM_PAR_LGD_BASEL", whereParam)
	if nil != err {
		er := fmt.Errorf("删除违约损失率信息出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

func (this LgdBasel) Init() error {
	zlog.Info("初始化押品映射缓存", nil)
	var paramMap = map[string]interface{}{"flag": util.FLAG_TRUE}

	// 加入生效期限
	err := util.GetStartTimeParam(paramMap)
	if nil != err {
		zlog.Error(err.Error(), err)
		return err
	}

	lgdBasels, err := this.SelectLgdBaselByParams(paramMap)
	if nil != err {
		er := fmt.Errorf("初始化[查询]违约违约损失率至缓存[%s]中错误", util.RPM_LGD_CACHE)
		zlog.Errorf(er.Error(), err)
		return er
	}
	for _, lgd := range lgdBasels {
		err := util.PutCacheByCacheName(util.RPM_LGD_CACHE, lgd.PledgeType, lgd, 0)
		if nil != err {
			er := fmt.Errorf("初始化违约损失率至缓存[%s]中错误", util.RPM_LGD_CACHE)
			zlog.Error(er.Error(), err)
			return er
		}
	}
	return nil
}

// 查询违约损失率表放入缓存中
// key   : 抵质押类型
// value : 违约损失率实体
func init() {
	err := LgdBasel{}.Init()
	if nil != err {
		zlog.Infof(err.Error(), err)
	}
}
