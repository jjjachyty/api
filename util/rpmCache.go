package util

import (
	"fmt"

	"github.com/tango-contrib/cache"
)

// 指标缓存
var rpmCache = cache.New()

func PutCache(key string, val interface{}, timeout int64) error {
	return rpmCache.Put(key, val, timeout)
}

func DeleteCache(key string) error {
	return rpmCache.Delete(key)
}

func GetCache(key string) interface{} {
	ok := rpmCache.IsExist(key)
	if ok {
		return rpmCache.Get(key)
	}
	return nil
}

func FlushCache() {
	rpmCache.Flush()
}

// 引擎缓存
// put first
var rpmEngineCache = cache.New()

func PutRpmEngineCache(key string, val interface{}, timeout int64) error {
	return rpmEngineCache.Put(key, val, timeout)
}

func DeleteRpmEngineCache(key string) error {
	return rpmEngineCache.Delete(key)
}

func GetRpmEngineCache(key string) interface{} {
	ok := rpmEngineCache.IsExist(key)
	if ok {
		return rpmEngineCache.Get(key)
	}
	return nil
}

func FlushRpmEngineCache() {
	rpmEngineCache.Flush()
}

// 违约损失率缓存
// var rpmLgdCache = cache.New(cache.Options{Adapter: "rpmLgdCache"})
var rpmLgdCache *cache.Caches

// 指标缓存
// key:uuid value:indicator
// var rpmIndicatorCache = cache.New(cache.Options{Adapter: "rpmIndicatorCache"})
var rpmIndicatorCache *cache.Caches

// 字典缓存
// key : parentDict value: dict
var rpmParentDictCache *cache.Caches

// AM系统数据权限缓存
// key : userId
// val : map[reqUrl][]DataAuth
var amDataAuthCache *cache.Caches

// 存放挤出人信息缓存
// key : 前台传的Sid
// val : 挤出人信息
var rpmSqueezeCache *cache.Caches

// 存放登陆用户信息
// key : sid
// val : SidUser
var rpmSiduserCache *cache.Caches

func PutCacheByCacheName(cacheName, key string, val interface{}, timeout int64) error {
	switch cacheName {
	case RPM_LGD_CACHE:
		return rpmLgdCache.Put(key, val, timeout)
	case RPM_INDICATOR_CACHE:
		return rpmIndicatorCache.Put(key, val, timeout)
	case RPM_PARENT_DICT_CACHE:
		return rpmParentDictCache.Put(key, val, timeout)
	case AM_DATA_AUTH_CACHE:
		return amDataAuthCache.Put(key, val, timeout)
	case RPM_SQUEEZE_CACHE:
		return rpmSqueezeCache.Put(key, val, timeout)
	case RPM_SID_USER_CACHE:
		return rpmSiduserCache.Put(key, val, timeout)
	default:
		return fmt.Errorf("没有名为[%s]的缓存", cacheName)
	}
}

func GetCacheByCacheName(cacheName, key string) interface{} {
	switch cacheName {
	case RPM_LGD_CACHE:
		return rpmLgdCache.Get(key)
	case RPM_INDICATOR_CACHE:
		return rpmIndicatorCache.Get(key)
	case RPM_PARENT_DICT_CACHE:
		return rpmParentDictCache.Get(key)
	case AM_DATA_AUTH_CACHE:
		return amDataAuthCache.Get(key)
	case RPM_SQUEEZE_CACHE:
		return rpmSqueezeCache.Get(key)
	case RPM_SID_USER_CACHE:
		return rpmSiduserCache.Get(key)
	default:
		return fmt.Errorf("没有名为[%s]的缓存", cacheName)
	}
}

type RpmCache struct {
}

func FlushCacheByCacheName(cacheName string) {
	switch cacheName {
	case RPM_LGD_CACHE:
		rpmLgdCache.Flush()
	case RPM_INDICATOR_CACHE:
		rpmIndicatorCache.Flush()
	case RPM_PARENT_DICT_CACHE:
		rpmParentDictCache.Flush()
	case AM_DATA_AUTH_CACHE:
		amDataAuthCache.Flush()
	case RPM_SQUEEZE_CACHE:
		rpmSqueezeCache.Flush()
	case RPM_SID_USER_CACHE:
		rpmSiduserCache.Flush()
	default:
		fmt.Errorf("没有名为[%s]的缓存", cacheName)
	}
}
func DeleteCacheByCacheName(cacheName, key string) {
	switch cacheName {
	case RPM_LGD_CACHE:
		rpmLgdCache.Delete(key)
	case RPM_INDICATOR_CACHE:
		rpmIndicatorCache.Delete(key)
	case RPM_PARENT_DICT_CACHE:
		rpmParentDictCache.Delete(key)
	case AM_DATA_AUTH_CACHE:
		amDataAuthCache.Delete(key)
	case RPM_SQUEEZE_CACHE:
		rpmSqueezeCache.Delete(key)
	case RPM_SID_USER_CACHE:
		rpmSiduserCache.Delete(key)
	default:
		fmt.Errorf("没有名为[%s]的缓存", cacheName)
	}
}

func init() {
	cache.Register("rpmIndicatorCache", cache.NewMemoryCacher())
	cache.Register("rpmLgdCache", cache.NewMemoryCacher())
	cache.Register("rpmParentDictCache", cache.NewMemoryCacher())
	cache.Register("amDataAuthCache", cache.NewMemoryCacher())
	rpmIndicatorCache = cache.New(cache.Options{Adapter: "rpmIndicatorCache"})
	rpmLgdCache = cache.New(cache.Options{Adapter: "rpmLgdCache"})
	rpmParentDictCache = cache.New(cache.Options{Adapter: "rpmParentDictCache"})
	amDataAuthCache = cache.New(cache.Options{Adapter: "amDataAuthCache"})

	cache.Register("rpmSqueezeCache", cache.NewMemoryCacher())
	rpmSqueezeCache = cache.New(cache.Options{Adapter: "rpmSqueezeCache"})

	cache.Register("rpmSiduserCache", cache.NewMemoryCacher())
	rpmSiduserCache = cache.New(cache.Options{Adapter: "rpmSiduserCache"})
}
