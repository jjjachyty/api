package par

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"time"

	"pccqcpa.com.cn/components/zlog"

	modelsUtil "pccqcpa.com.cn/app/rpm/api/models/util"
	"pccqcpa.com.cn/app/rpm/api/util"
)

/**
 * @apiDefine Dict
 * @apiSuccess {string}     UUID            主键默认值sys_guid()
 * @apiSuccess {string}   	DictCode    	字典码值
 * @apiSuccess {string}   	DictName      	字典名称
 * @apiSuccess {string}   	ParentDict   	父级字典
 * @apiSuccess {string}   	DictType    	字典类型
 * @apiSuccess {string}   	Flag      		生效标志
 * @apiSuccess {int}   		Sort      		排序
 * @apiSuccess {time.Time}	CreateTime      创建时间
 * @apiSuccess {string}   	CreateUser      创建人
 * @apiSuccess {time.Time}	UpdateTime      更新时间
 * @apiSuccess {string}   	UpdateUser      更新人
 */
type Dict struct {
	UUID       string    //主键
	DictCode   string    //字典码值 唯一
	DictName   string    //字典名称
	ParentDict string    //父级字典
	DictType   string    //字典类型
	Flag       string    //生效标志
	Sort       int       //排序
	CreateTime time.Time //创建时间
	CreateUser string    //创建人
	UpdateTime time.Time //更新时间
	UpdateUser string    //更新人
}

func (d *Dict) scan(rows *sql.Rows) (*Dict, error) {
	var dict = new(Dict)
	values := []interface{}{
		&dict.UUID,
		&dict.DictCode,
		&dict.DictName,
		&dict.ParentDict,
		&dict.DictType,
		&dict.Flag,
		&dict.Sort,
		&dict.CreateTime,
		&dict.CreateUser,
		&dict.UpdateTime,
		&dict.UpdateUser,
	}
	err := util.OracleScan(rows, values)
	return dict, err
}

func (d Dict) List(param ...map[string]interface{}) (*util.PageData, error) {
	pageData, rows, err := modelsUtil.List(dictTables, dictCols, dictColsSort, param...)
	defer rows.Close()
	if nil != err {
		er := errors.New("分页查询数据字典出错")
		zlog.Error(er.Error(), err)
		return nil, er
	}
	var dicts []*Dict
	for rows.Next() {
		dict, err := d.scan(rows)
		if nil != err {
			er := fmt.Errorf("分页查询数据字典rows.Scan()出错")
			zlog.Error(er.Error(), err)
			return nil, er
		}
		dicts = append(dicts, dict)
	}
	pageData.Rows = dicts
	return pageData, nil
}

func (d *Dict) Find(param ...map[string]interface{}) ([]*Dict, error) {
	rows, err := modelsUtil.FindRows(dictTables, dictCols, dictColsSort, param...)
	defer rows.Close()
	if nil != err {
		er := fmt.Errorf("查询数据字典出错")
		zlog.Error(er.Error(), err)
		return nil, er
	}
	var dicts []*Dict
	for rows.Next() {
		dict, err := d.scan(rows)
		if nil != err {
			er := fmt.Errorf("查询数据字典rows.Scan()出错")
			zlog.Error(er.Error(), err)
			return nil, er
		}
		dicts = append(dicts, dict)
	}
	return dicts, nil
}

func (d Dict) Add() error {
	paramMap := map[string]interface{}{
		"dict_code":   d.DictCode,
		"dict_name":   d.DictName,
		"parent_dict": d.ParentDict,
		"dict_type":   d.DictType,
		"flag":        d.Flag,
		"sort":        d.Sort,
		"create_time": util.GetCurrentTime(),
		"create_user": d.CreateUser,
		"update_time": util.GetCurrentTime(),
		"update_user": d.UpdateUser,
	}
	err := util.OracleAdd(dictTables, paramMap)
	if nil != err {
		er := errors.New("新增数据字典出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

func (d Dict) Update() error {
	paramMap := map[string]interface{}{
		"dict_code":   d.DictCode,
		"dict_name":   d.DictName,
		"parent_dict": d.ParentDict,
		"dict_type":   d.DictType,
		"flag":        d.Flag,
		"sort":        d.Sort,
		"create_time": util.GetCurrentTime(),
		"update_user": d.UpdateUser,
		"update_time": util.GetCurrentTime(),
	}
	whereParamMap := map[string]interface{}{
		"uuid": d.UUID,
	}
	err := util.OracleUpdate(dictTables, paramMap, whereParamMap)
	if nil != err {
		er := errors.New("更新数据字典出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

func (d Dict) Delete() error {
	whereParamMap := map[string]interface{}{
		"uuid": d.UUID,
	}
	err := util.OracleDelete(dictTables, whereParamMap)
	if nil != err {
		er := errors.New("删除数据字典出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

var dictTables string = " RPM_PAR_DICT T"

var dictCols map[string]string = map[string]string{
	"T.UUID":        "''",
	"T.DICT_CODE":   "' '",
	"T.DICT_NAME":   "' '",
	"T.PARENT_DICT": "' '",
	"T.DICT_TYPE":   "' '",
	"T.FLAG":        "' '",
	"T.SORT":        "0",
	"T.CREATE_TIME": "sysdate",
	"T.CREATE_USER": "' '",
	"T.UPDATE_TIME": "sysdate",
	"T.UPDATE_USER": "' '",
}

var dictColsSort []string = []string{
	"T.UUID",
	"T.DICT_CODE",
	"T.DICT_NAME",
	"T.PARENT_DICT",
	"T.DICT_TYPE",
	"T.FLAG",
	"T.SORT",
	"T.CREATE_TIME",
	"T.CREATE_USER",
	"T.UPDATE_TIME",
	"T.UPDATE_USER",
}

func (d Dict) Init() {
	util.FlushCacheByCacheName(util.RPM_PARENT_DICT_CACHE)
	paramMap := map[string]interface{}{
		"parent_dict": util.DICT_TOP_CODE,
		"sort":        "sort",
		"order":       "ASC",
	}
	dicts, err := d.Find(paramMap)
	if nil != err {
		zlog.Errorf("初始化字典时查询父级节点为【%s】的字典出错", err, util.DICT_TOP_CODE)
		os.Exit(1)
	}
	util.PutCacheByCacheName(util.RPM_PARENT_DICT_CACHE, util.DICT_TOP_CODE, dicts, 0)
	for _, dict := range dicts {
		paramMap := map[string]interface{}{
			"parent_dict": dict.DictCode,
			"sort":        "sort",
			"order":       "ASC",
		}
		subDicts, err := d.Find(paramMap)
		if nil != err {
			zlog.Errorf("初始化字典时查询父级节点为【%s】的字典出错", err, dict.DictCode)
			os.Exit(1)
		}
		zlog.Infof("缓存字典父级字典编号【%s】成功", nil, dict.DictCode)
		util.PutCacheByCacheName(util.RPM_PARENT_DICT_CACHE, dict.DictCode, subDicts, 0)
	}
}

func init() {
	Dict{}.Init()
}
