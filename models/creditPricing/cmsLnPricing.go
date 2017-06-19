package creditPricing

import (
	"pccqcpa.com.cn/app/rpm/api/models/biz/ln"
)

type CmsLnPricing struct {
	CustInfo   ln.CustInfo     // 客户信息
	LnBusiness ln.LnBusiness   // 业务信息
	LnMorts    []ln.LnMort     // 押品信息
	Guarantes  []ln.LnGuarante // 保证人信息
	SceneDps   []ln.SceneDp    // 存款派生信息
	SceneItds  []ln.SceneItd   // 中间业务派生信息
	StockUsage float64         // 存量优惠
	IntRate    float64         // 执行利率
	MarginType string          // 浮动类型
	MarginInt  float64         // 浮动值
}
