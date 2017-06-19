package engine

import (
	"errors"
	"log"
	"os"
	"runtime/pprof"

	"github.com/lunny/tango"

	"pccqcpa.com.cn/app/rpm/api/models/biz/ln"
	"pccqcpa.com.cn/app/rpm/api/util"
)

type LnScenePricingAction struct {
	tango.Ctx
	tango.Json
}

//URL POST:/api/rpm/pricing/scene/
func (l *LnScenePricingAction) Post() util.RstMsg {

	// flag.Parse()

	// //这里实现了远程获取pprof数据的接口
	// go func() {
	// 	log.Println(http.ListenAndServe("localhost:6060", nil))
	// }()
	// defer f.Close()
	// defer pprof.StopCPUProfile()
	var lnPrincing *ln.LnPricing
	var err error
	f, err := os.OpenFile("scene.prof", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}

	pprof.StartCPUProfile(f)

	businessCode := l.Form("BusinessCode")
	StockUsage := l.FormFloat64("StockUsage")

	if "" == businessCode {
		return util.ErrorMsg("[BusinessCode.业务单号]为空,不能计算对公贷款情景优惠", errors.New("[BusinessCode.业务单号]为空,不能计算对公贷款情景优惠"))
	} else if "" == l.Form("StockUsage") {
		return util.ErrorMsg("[StockUsage.存量优惠使用]为空,不能计算对公贷款情景优惠", errors.New("[StockUsage.存量优惠使用]为空,不能计算对公贷款情景优惠"))
	} else {
		lnPrincing, err = service.LnBusinessPricing(businessCode, util.LN_SCENE, map[string]interface{}{"StockUsage": StockUsage})
		if nil != err {
			return util.ErrorMsg(err.Error(), err)
		}
	}

	return util.SuccessMsg("情景优惠计算成功", lnPrincing)
}
