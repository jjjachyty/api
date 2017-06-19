package util

import (
	"strconv"

	"fmt"
	"github.com/go-ini/ini"
	"os"
	"pccqcpa.com.cn/components/zlog"
)

var cfg *ini.File

func init() {
	zlog.Debug("初始化system.ini配置文件", nil)
	confPath := os.Getenv("GOSYSCONFIG")
	fmt.Println("配置文件路径", confPath)
	file, err := ini.Load(confPath + "/system.ini")
	if nil == err {
		cfg = file
	} else {
		zlog.Errorf("初始化system.ini配置错误", err)
	}
}

// GetIniIntValue func
func GetIniStringValue(selectName string, key string) string {
	return getIniValue(selectName, key)
}

// GetIniIntValue func
func GetIniIntValue(selectName string, key string) int {
	intValue, err := strconv.Atoi(getIniValue(selectName, key))
	if nil == err {
		return intValue
	}
	zlog.Errorf("获取system.ini配置[%s]的[%s]init值出错", err, selectName, key)
	return 0

}
func getIniValue(selectName string, key string) string {
	if cfg.Section(selectName).Haskey(key) {
		return cfg.Section(selectName).Key(key).Value()
	}
	zlog.Errorf("查找system.ini配置[%s] 的[%s]值错误", nil, selectName, key)
	return ""
}
