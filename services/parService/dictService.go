package parService

import (
	"errors"
	"fmt"
	"strings"

	"pccqcpa.com.cn/app/rpm/api/models/par"
	"pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/components/zlog"
)

type DictService struct{}

var dictModel par.Dict

func (d DictService) Init() {
	dictModel.Init()
}

func (d DictService) List(param ...map[string]interface{}) (*util.PageData, error) {
	d.handleParamMap(param...)
	return dictModel.List(param...)
}

// 判断是否可以直接从缓存中获取数据
func (d DictService) Find(param ...map[string]interface{}) ([]*par.Dict, error) {
	parentDict, canFind, isAsc := d.canFindByCache(param...)
	if canFind {
		value := util.GetCacheByCacheName(util.RPM_PARENT_DICT_CACHE, parentDict)
		if nil == value {
			zlog.Infof("从数据库中获取父级字典编号为【%s】的字典", nil, parentDict)
		} else if isAsc {
			zlog.Infof("从缓存中升序获取父级字典编号为【%s】的字典", nil, parentDict)
			dicts := value.([]*par.Dict)
			var rstDicts []*par.Dict
			for i := 0; i < len(dicts); i++ {
				if util.FLAG_TRUE == dicts[i].Flag {
					rstDicts = append(rstDicts, dicts[i])
				}
			}
			return rstDicts, nil
		} else {
			zlog.Infof("从缓存中降序获取父级字典编号为【%s】的字典", nil, parentDict)
			dicts := value.([]*par.Dict)
			var rstDicts []*par.Dict
			for i := len(dicts) - 1; i >= 0; i-- {
				if util.FLAG_TRUE == dicts[i].Flag {
					rstDicts = append(rstDicts, dicts[i])
				}
			}
			return rstDicts, nil
		}
	}
	var paramMap map[string]interface{}
	if 0 != len(param) {
		paramMap = param[0]
		paramMap["sort"] = "sort"
		paramMap["order"] = "ASC"
	}
	return dictModel.Find(paramMap)
}

func (d DictService) FindOne(paramMap map[string]interface{}) (*par.Dict, error) {
	dicts, err := d.Find(paramMap)
	if nil != err {
		return nil, err
	}
	switch len(dicts) {
	case 0:
		return nil, nil
	case 1:
		return dicts[0], nil
	}
	er := errors.New("查询字典有多条返回纪录")
	zlog.Error(er.Error(), er)
	return nil, er
}

// 查询缓存中是否有改编号的字典，如果有则报错
func (d DictService) Add(dict *par.Dict) error {
	dictCode := dict.DictCode
	sort := dict.Sort
	dictFlag := dict.Flag
	if "" == dict.ParentDict {
		er := fmt.Errorf("添加的字典的父节点为空，请刷新页面重新添加")
		zlog.Error(er.Error(), er)
		return er
	}
	value := util.GetCacheByCacheName(util.RPM_PARENT_DICT_CACHE, dict.ParentDict)
	if nil != value {
		dicts := value.([]*par.Dict)
		for _, dict := range dicts {
			if dictCode == dict.DictCode {
				err := fmt.Errorf("字典编号为【%s】的字典已存在，不可以新增", dictCode)
				zlog.Info(err.Error(), nil)
				return err
			} else if sort == dict.Sort && util.FLAG_TRUE == dictFlag && dictFlag == dict.Flag && util.DICT_TOP_CODE != dict.ParentDict {
				er := fmt.Errorf("排序重复，请设置不一样的排序")
				zlog.Error(er.Error(), er)
				return er
			}
		}
	}
	err := dict.Add()
	if nil != err {
		return err
	}
	// 添加至缓存中
	paramMap := map[string]interface{}{
		"dict_code":   dict.DictCode,
		"parent_dict": dict.ParentDict,
	}
	dictNew, err := d.FindOne(paramMap)
	// if nil != err || nil == dict {
	// 	go dict.Init()
	// }
	if nil == dictNew {
		zlog.Infof("查询字典为空", nil)
		return nil
	} else if nil != value {
		dicts := value.([]*par.Dict)
		var dictsCache = make([]*par.Dict, 0)
		if 0 == len(dicts) {
			dictsCache = append(dictsCache, dictNew)
			util.PutCacheByCacheName(util.RPM_PARENT_DICT_CACHE, dict.ParentDict, dictsCache, 0)
			return nil
		}
		for i, dict := range dicts {
			switch {
			case dict.Sort >= sort:
				dictsCache = append(dictsCache, dicts[0:i]...)
				dictsCache = append(dictsCache, dictNew)
				dictsCache = append(dictsCache, dicts[i:]...)
				util.PutCacheByCacheName(util.RPM_PARENT_DICT_CACHE, dict.ParentDict, dictsCache, 0)
			case dict.Sort < sort && i == len(dicts)-1:
				dictsCache = append(dicts, dictNew)
				util.PutCacheByCacheName(util.RPM_PARENT_DICT_CACHE, dict.ParentDict, dictsCache, 0)
			default:
				continue
			}
			break
		}
	} else {
		util.PutCacheByCacheName(util.RPM_PARENT_DICT_CACHE, dictNew.ParentDict, []*par.Dict{dictNew}, 0)
	}
	return nil
}

func (d DictService) Update(dict *par.Dict) error {
	dictCodeOld := dict.DictCode
	dictSortNew := dict.Sort
	UUIDOld := dict.UUID
	dictFlag := dict.Flag
	value := util.GetCacheByCacheName(util.RPM_PARENT_DICT_CACHE, dict.ParentDict)
	if nil != value {
		dicts := value.([]*par.Dict)
		for _, dict := range dicts {
			if dictCodeOld == dict.DictCode && dict.UUID != UUIDOld {
				err := fmt.Errorf("字典编号为【%s】的字典已存在，不可以更新", dictCodeOld)
				zlog.Info(err.Error(), nil)
				return err
			} else if dictSortNew == dict.Sort && dict.UUID != UUIDOld && util.FLAG_TRUE == dictFlag && dictFlag == dict.Flag && util.DICT_TOP_CODE != dict.ParentDict {
				er := fmt.Errorf("排序重复，请设置不一样的排序")
				zlog.Error(er.Error(), er)
				return er
			}

		}
	}
	err := dict.Update()
	if nil != err {
		return err
	}
	if value == nil {
		util.PutCacheByCacheName(util.RPM_PARENT_DICT_CACHE, dict.ParentDict, []*par.Dict{dict}, 0)
	} else {
		dicts := value.([]*par.Dict)
		var dictsCache []*par.Dict
		dictNew := dict
		var previousDict *par.Dict = nil
		for _, dict := range dicts {
			if dict.Sort >= dictNew.Sort && ((previousDict != nil && dictNew.Sort > previousDict.Sort) || previousDict == nil) && dictNew.UUID != dict.UUID {
				dictsCache = append(dictsCache, dictNew)
				dictsCache = append(dictsCache, dict)
			} else if dictNew.UUID == dict.UUID {
				continue
			} else {
				dictsCache = append(dictsCache, dict)
			}
			previousDict = dict
		}
		util.PutCacheByCacheName(util.RPM_PARENT_DICT_CACHE, dictNew.ParentDict, dictsCache, 0)

	}
	return nil

}

// 判断是否删除的父字典，判断改父字典是否还有子字典，如果有，则不可以删除
func (d DictService) Delete(dict *par.Dict) error {
	dictCodeOld := dict.DictCode
	if dict.DictCode != util.DICT_TOP_CODE {
		value := util.GetCacheByCacheName(util.RPM_PARENT_DICT_CACHE, dict.DictCode)
		if nil != value && 0 != len(value.([]*par.Dict)) {
			err := fmt.Errorf("字典编号为【%s】的字典存在子字典，不可以删除，请先删除子字典", dict.DictCode)
			zlog.Info(err.Error(), nil)
			zlog.Infof("存在的字典缓存[%#v]", nil, value)
			return err
		} else {
			dicts, err := dict.Find(map[string]interface{}{"parent_dict": dict.DictCode})
			if nil != err {
				return err
			}
			if 0 != len(dicts) {
				err := fmt.Errorf("字典编号为【%s】的字典存在子字典，不可以删除，请先删除子字典", dict.DictCode)
				zlog.Infof("子字典未进入缓存[%#v]", nil, dicts)
				return err
			}
		}
	}
	err := dict.Delete()
	if nil != err {
		return err
	}
	value := util.GetCacheByCacheName(util.RPM_PARENT_DICT_CACHE, dict.ParentDict)
	// if nil == value {
	// 	go dict.Init()
	// }
	dicts := value.([]*par.Dict)
	for i, dict := range dicts {
		if dictCodeOld == dict.DictCode {
			var dictsCache []*par.Dict
			if i >= len(dicts)-1 {
				dictsCache = dicts[:i]
			} else {
				dictsCache = append(dicts[:i], dicts[i+1:]...)
			}
			util.PutCacheByCacheName(util.RPM_PARENT_DICT_CACHE, dict.ParentDict, dictsCache, 0)
			break
		}
	}
	return nil
}

// 判断是否可以直接从缓存中获取
// 返回值 1:是否可以从缓存中获取 true 可以 false 不可以
// 返回值 2:是否生序 true 升序 false 降序
func (d DictService) canFindByCache(param ...map[string]interface{}) (string, bool, bool) {

	var parentDict string = ""
	var canFind, isAsc bool = false, true
	var paramMap map[string]interface{}
	switch len(param) {
	case 0:
		return parentDict, false, false
	default:
		paramMap = param[0]
	}
	for k, v := range paramMap {
		switch strings.ToUpper(k) {
		case "PARENT_DICT":
			parentDict = v.(string)
			canFind = true
		case "SORT":
			canFind = true
		case "ORDER":
			canFind = true
			switch strings.ToUpper(v.(string)) {
			case "ASC", "":
				isAsc = true
			case "DESC":
				isAsc = false
			default:
				return parentDict, false, false
			}
		default:
			return parentDict, false, false
		}
	}
	return parentDict, canFind, isAsc
}

// 处理模糊查询
func (d DictService) handleParamMap(param ...map[string]interface{}) {
	if 0 < len(param) {
		parmaMap := param[0]
		dictCode := parmaMap["dict_code"]
		key := []interface{}{
			"dict_code", "dict_name",
		}
		value := []interface{}{
			dictCode, dictCode,
		}
		parmaMap["searchLike"] = []map[string]interface{}{
			map[string]interface{}{
				"type":  "or",
				"key":   key,
				"value": value,
			},
		}
		delete(parmaMap, "dict_code")
		// param[0] = parmaMap
	}
}
