package util

import (
	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-oci8"
)

var Engine *xorm.Engine

func init() {

}

// //实体的属性获得列名
// func Struct2Column(bean interface{}, attr string) {
// 	refBean := reflect.TypeOf(bean)
// 	tag := refBean.FieldByName(attr).Tag.Get("xorm")
// }
