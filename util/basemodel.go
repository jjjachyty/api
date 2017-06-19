package util

import (
	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-oci8"
	"pccqcpa.com.cn/components/zlog"
)

var Engine *xorm.Engine

func init() {
	var err error
	Engine, err = xorm.NewEngine("oci8", "tjrpm10/tjrpm@172.168.171.250/APPDEV")

	Engine.ShowSQL(true)
	if err != nil {
		zlog.Error("orm 初始化出错", err)
	}
}

// //实体的属性获得列名
// func Struct2Column(bean interface{}, attr string) {
// 	refBean := reflect.TypeOf(bean)
// 	tag := refBean.FieldByName(attr).Tag.Get("xorm")
// }
