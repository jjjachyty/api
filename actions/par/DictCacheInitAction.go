package par

import (
	"github.com/lunny/tango"
	"pccqcpa.com.cn/app/rpm/api/engine/pricing"
	"pccqcpa.com.cn/app/rpm/api/util"
)

type DictCacheInitAction struct {
	tango.Json
	tango.Ctx
}

func (d *DictCacheInitAction) Post() util.RstMsg {

	dictService.Init()

	pricing.PricingBus{}.Init()
	return util.SuccessMsg("初始化字典缓存成功", nil)
}
