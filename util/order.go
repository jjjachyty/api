package util

import (
	"net/http"

	"strings"

	"pccqcpa.com.cn/components/zlog"
)

//DESC 降序排序
const DESC OrderType = " DESC "

// ASC 升序排序
const ASC OrderType = " ASC "

//OrderType type 排序类型
type OrderType string

//Order type 排序实体
type Order struct {
	OrderAttr string
	OrderType OrderType
}

//GetOrder func 从header获得排序参数
func GetOrder(reqHeader http.Header) Order {
	var ord Order
	orderAttr := reqHeader.Get("order-attr")
	orderType := reqHeader.Get("order-type")

	col := strings.ToUpper(UperChange(orderAttr))
	zlog.Debugf("根据%s %s 排序", nil, col, orderType)
	if orderAttr == "" {
		zlog.Warning("Order 无排序字段,自动默认为[CREATE_TIME]排序", nil)
		ord.OrderAttr = "CREATE_TIME"
	} else {

		if col != "" {
			ord.OrderAttr = (col)
		}
		// else {
		// 	zlog.Warningf("Order 实体无法找到字段%s对应列,自动默认为[CREATE_TIME]排序", nil, orderAttr)
		// 	ord.OrderAttr = "CREATE_TIME"
		// }

	}
	if orderType == "" {
		zlog.Warning("Order 无排序规则,自动默认为[DESC]排序", nil)
		ord.OrderType = DESC
	} else {
		switch strings.ToUpper(orderType) {
		case "DESC":
			ord.OrderType = DESC
		case "ASC":
			ord.OrderType = ASC
		default:
			zlog.Warning("Order 不支持的排序规则,自动默认为[DESC]排序", nil)
			ord.OrderType = DESC
		}
	}
	return ord
}
