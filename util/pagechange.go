package util

import (
	"github.com/lunny/tango"
)

func ChangePageMsgHeadToMap(param map[string]interface{}, ctx *tango.Ctx) {
	start := ctx.Req().Header.Get("start-row-number")
	length := ctx.Req().Header.Get("page-size")
	// 如果没有设置开始则默认从第一条开始
	if "" == start {
		start = "0"
	}
	// 如果没有设置每页的条数，则默认为每页15条记录
	if "" == length {
		length = "10"
	}
	param["start"] = start
	param["length"] = length
}
