package routers

import (
	"github.com/lunny/tango"
	"pccqcpa.com.cn/app/rpm/api/actions"
	"pccqcpa.com.cn/app/rpm/api/actions/amAuthData"
	"pccqcpa.com.cn/app/rpm/api/actions/bi"
	"pccqcpa.com.cn/app/rpm/api/actions/bi/analysisCrossDim"
	"pccqcpa.com.cn/app/rpm/api/actions/bi/analysisSigDim"
	"pccqcpa.com.cn/app/rpm/api/actions/bi/dim"
	"pccqcpa.com.cn/app/rpm/api/actions/biz/cf"
	"pccqcpa.com.cn/app/rpm/api/actions/biz/dp"
	"pccqcpa.com.cn/app/rpm/api/actions/biz/ln"
	"pccqcpa.com.cn/app/rpm/api/actions/biz/rl"
	"pccqcpa.com.cn/app/rpm/api/actions/common"
	"pccqcpa.com.cn/app/rpm/api/actions/creditInterface"
	jasondp "pccqcpa.com.cn/app/rpm/api/actions/dp"
	"pccqcpa.com.cn/app/rpm/api/actions/engine"
	"pccqcpa.com.cn/app/rpm/api/actions/loan"
	"pccqcpa.com.cn/app/rpm/api/actions/login"
	"pccqcpa.com.cn/app/rpm/api/actions/par"
	"pccqcpa.com.cn/app/rpm/api/actions/workFlow"
)

func Init(t *tango.Tango) {

	t.Group("/api/rpm", func(tg *tango.Group) {
		// 何培兵
		tg.Any("/sign/(*param)", new(actions.LoginAction))
		//张力
		tg.Any("/auth", new(actions.AuthAction))
		// 指标维护
		// 默认查询全部带分页查询

		tg.Group("", func(tan *tango.Group) {
			tg.Route([]string{"GET:List", "POST", "PUT", "DELETE"}, "/indicator", new(par.IndicatorAction))
			// 带参数查询，判断header里面的start-row-number是否为空，如果为空，则不分页
			tg.Get("/indicator/(*param)", new(par.IndicatorAction))
		})

		// 客户信息维护
		// 默认查询带分页查询
		tg.Group("", func(tan *tango.Group) {
			tan.Route([]string{"GET:List", "POST", "PUT", "DELETE"}, "/cust", new(ln.CustInfoAction))
			// 带参数查询，判断header里面的start-row-number是否为空，如果为空，则不分页
			tan.Get("/cust/(*param)", new(ln.CustInfoAction))

			// 贷款定价单添加保证人专用路由
			tan.Get("/cust/guarante", new(ln.CustInfoGuranteAction))
			tan.Get("/cust/guarante/(*param)", new(ln.CustInfoGuranteAction))
		})

		//字典管理
		tg.Group("", func(tan *tango.Group) {
			tg.Route([]string{"GET:List", "POST", "PUT", "DELETE"}, "/dict", new(par.DictAction))
			tan.Get("/dict/(*param)", new(par.DictAction))
			tg.Post("/dict/cache/init", new(par.DictCacheInitAction))
		})

		tg.Group("", func(tan *tango.Group) {
			tan.Route([]string{"GET:List", "POST", "PUT", "DELETE"}, "/lnmort", new(ln.LnMortAction))
			tan.Get("/lnmort/(*param)", new(ln.LnMortAction))
		})

		tg.Group("", func(tan *tango.Group) {
			tan.Route([]string{"GET:List", "POST", "PUT", "DELETE"}, "/lnguarante", new(ln.LnGuaranteAction))
			tan.Get("/lnguarante/(*param)", new(ln.LnGuaranteAction))
		})

		tg.Group("", func(tan *tango.Group) {
			tan.Route([]string{"GET:List", "POST", "PUT", "DELETE"}, "/lnbusiness", new(ln.LnBusinessAction))
			tan.Get("/lnbusiness/(*param)", new(ln.LnBusinessAction))
		})

		// 调节系数
		tg.Group("", func(tan *tango.Group) {
			tg.Route([]string{"GET:List", "POST", "PUT", "DELETE"}, "/scenediscountadj", new(par.SceneDiscountAdjAction))
			tg.Get("/scenediscountadj/(*param)", new(par.SceneDiscountAdjAction))
		})

		// 基准利率
		tg.Group("", func(tan *tango.Group) {
			tg.Route([]string{"GET:List", "POST", "PUT", "DELETE"}, "/baserate", new(par.BaseRateAction))
			tg.Get("/baserate/(*param)", new(par.BaseRateAction))
		})

		// 经济资本占用率
		tg.Group("", func(tan *tango.Group) {
			tg.Route([]string{"GET:List", "POST", "PUT", "DELETE"}, "/ec", new(par.EcAction))
			tg.Get("/ec/(*param)", new(par.EcAction))
		})

		// EVA收益率
		tg.Group("", func(tan *tango.Group) {
			tg.Route([]string{"GET:List", "POST", "PUT", "DELETE"}, "/evaYield", new(par.EvaYieldAction))
			tg.Get("/evaYield/(*param)", new(par.EvaYieldAction))
		})

		// 通用参数
		tg.Group("", func(tan *tango.Group) {
			tg.Route([]string{"GET:List", "POST", "PUT", "DELETE"}, "/common", new(par.CommonAction))
			tg.Get("/common/(*param)", new(par.CommonAction))
		})

		// 资金成本率
		tg.Group("", func(tan *tango.Group) {
			tg.Route([]string{"GET:List", "POST", "PUT", "DELETE"}, "/ftp", new(par.FtpAction))
			tg.Get("/ftp/(*param)", new(par.FtpAction))
		})

		//违约损失率
		tg.Group("", func(tan *tango.Group) {
			tg.Route([]string{"GET:List", "POST", "PUT", "DELETE"}, "/lgdBasel", new(par.LgdBaselAction))
			tg.Get("/lgdBasel/(*param)", new(par.LgdBaselAction))
		})

		//运营成本率
		tg.Group("", func(tan *tango.Group) {
			tg.Route([]string{"GET:List", "POST", "PUT", "DELETE"}, "/oc", new(par.OcAction))
			tg.Get("/oc/(*param)", new(par.OcAction))
		})

		//违约概率
		tg.Group("", func(tan *tango.Group) {
			tg.Route([]string{"GET:List", "POST", "PUT", "DELETE"}, "/pd", new(par.PdAction))
			tg.Get("/pd/(*param)", new(par.PdAction))
		})

		//派生中间收入收益率
		tg.Group("", func(tan *tango.Group) {
			tg.Route([]string{"GET:List", "POST", "PUT", "DELETE"}, "/scene", new(par.SceneltdYieldAction))
			tg.Get("/scene/(*param)", new(par.SceneltdYieldAction))
		})

		//税率
		tg.Group("", func(tan *tango.Group) {
			tg.Route([]string{"GET:List", "POST", "PUT", "DELETE"}, "/tax", new(par.TaxAction))
			tg.Get("/tax/(*param)", new(par.TaxAction))
		})

		//存款操作风险率
		tg.Group("", func(tan *tango.Group) {
			tg.Route([]string{"GET:List", "POST", "PUT", "DELETE"}, "/dpOp", new(par.DpOpAction))
			tg.Get("/dpOp/(*param)", new(par.DpOpAction))
		})

		//零售资本成本率
		tg.Group("", func(tan *tango.Group) {
			tg.Route([]string{"GET:List", "POST", "PUT", "DELETE"}, "/rlEc", new(par.RlEcAction))
			tg.Get("/rlEc/(*param)", new(par.RlEcAction))
		})

		//零售违约损失率
		tg.Group("", func(tan *tango.Group) {
			tg.Route([]string{"GET:List", "POST", "PUT", "DELETE"}, "/rlLgd", new(par.RlLgdAction))
			tg.Get("/rlLgd/(*param)", new(par.RlLgdAction))
		})

		//零售违约概率
		tg.Group("", func(tan *tango.Group) {
			tg.Route([]string{"GET:List", "POST", "PUT", "DELETE"}, "/rlPd", new(par.RlPdAction))
			tg.Get("/rlPd/(*param)", new(par.RlPdAction))
		})

		//零售资本回报率
		tg.Group("", func(tan *tango.Group) {
			tg.Route([]string{"GET:List", "POST", "PUT", "DELETE"}, "/rlRoc", new(par.RlRocAction))
			tg.Get("/rlRoc/(*param)", new(par.RlRocAction))
		})

		//资本回报率
		tg.Group("", func(tan *tango.Group) {
			tg.Route([]string{"GET:List", "POST", "PUT", "DELETE"}, "/roc", new(par.RocAction))
			tg.Get("/roc/(*param)", new(par.RocAction))
		})

		// 定性优惠项
		tg.Group("", func(tan *tango.Group) {
			tg.Route([]string{"GET:List", "POST", "PUT", "DELETE"}, "/qualitativeDiscount", new(par.QualitativeDiscountAction))
			tg.Get("/qualitativeDiscount/(*param)", new(par.QualitativeDiscountAction))
		})

		//中收收益率
		tg.Group("", func(tan *tango.Group) {
			tg.Route([]string{"GET:List", "POST", "PUT", "DELETE"}, "/sceneltdYield", new(par.SceneDiscountAdjAction))
			tg.Get("/sceneltdYield/(*param)", new(par.SceneDiscountAdjAction))
		})

		//存款基准利率及挂牌－存款模块
		tg.Group("", func(tan *tango.Group) {
			tg.Route([]string{"GET:List", "POST", "PUT", "DELETE"}, "/dpBaseRate", new(par.DpBaseRateAction))
			tg.Get("/dpBaseRate/(*param)", new(par.DpBaseRateAction))
		})

		tg.Group("", func(tan *tango.Group) {
			tg.Route([]string{"POST", "PUT", "DELETE"}, "/scenedp", new(loan.SceneDpAction))
			tg.Get("/scenedp/(*param)", new(loan.SceneDpAction))
			tg.Route([]string{"POST", "PUT", "DELETE"}, "/sceneitd", new(loan.SceneItdAction))
			tg.Get("/sceneitd/(*param)", new(loan.SceneItdAction))
		})

		// Jason 定价水平分析报表
		tg.Group("", func(tan *tango.Group) {
			tg.Route([]string{"GET:List", "POST"}, "/analysisSigDim", new(analysisSigDim.AnalysisSigDimAction))
			tg.Route([]string{"GET:ListDate"}, "/analysisSigDim/date", new(analysisSigDim.AnalysisSigDimAction))
			tg.Get("/analysisSigDim/(*param)", new(analysisSigDim.AnalysisSigDimAction))
		})
		// Jason 定价薯片分析报表-交叉分析
		tg.Group("", func(tan *tango.Group) {
			tg.Get("/analysisCrossDim/(*param)", new(analysisCrossDim.AnalysisCrossDimAction))
		})

		// 公共业务表信息维护
		tg.Get("/tree/(:structName)", new(common.TreeAction))
		tg.Get("/code/(:action)", new(common.CodeAction))

		// 计算类
		tg.Group("/pricing", func(tan *tango.Group) {
			tan.Post("/lnbase", new(engine.LnBusinessPricingAction))
			tan.Post("/lnscene", new(engine.LnScenePricingAction))
			tan.Post("/lninverse", new(engine.LnBusinessInverseAction))
			// 信贷接口定价
			tan.Post("/cmslnpricing", new(creditInterface.CmsLnPricingAction))
		})

		// 	Janly贷款定价单
		tg.Group("", func(tgln *tango.Group) {
			//tgln.Get("/pricelists", new(loan.PricingListAction))
			tgln.Route([]string{"GET:List", "DELETE"}, "/pricelists", new(loan.PricingListAction))
			tgln.Any("/pricelist/businesscode/(:businesscode)", new(loan.PricingListAction))
			tgln.Route([]string{"GET:List"}, "/pricelists/(*param)", new(loan.PricingListAction))

		})
		// 	Janly零售价格矩阵
		tg.Group("/rl", func(tgln *tango.Group) {
			//tgln.Get("/pricelists", new(loan.PricingListAction))
			tgln.Route([]string{"GET:List", "POST"}, "/pricematrixs", new(rl.RLPriceMatrixAction))
			//tgln.Any("/rl/businesscode/(:businesscode)", new(rl.RLPriceMatrixAction))

		})
		// 	Janly客户分类
		tg.Group("/classfy", func(tgln *tango.Group) {
			//tgln.Get("/pricelists", new(loan.PricingListAction))
			tgln.Route([]string{"GET:List", "POST", "PATCH", "DELETE"}, "/nomination", new(cf.CustNominationAction))
			tgln.Route([]string{"GET:List"}, "/nomination/(*param)", new(cf.CustNominationAction))
			tgln.Route([]string{"POST:ExcelImport"}, "/nomination/excel", new(cf.CustNominationAction))

		})

		// Jason存款服务类
		tg.Group("/dp", func(tgln *tango.Group) {
			// 存款参数服务类
			tgln.Route([]string{"GET:List", "POST", "PUT", "DELETE"}, "/differ", new(par.DpDifferAction))
			tgln.Get("/differ/(*param)", new(par.DpDifferAction))

			//存款定价
			tgln.Route([]string{"GET:List", "POST", "PUT", "DELETE"}, "/pricing", new(jasondp.DpPricingAction))

			// 存款一对一定价
			tgln.Route([]string{"GET:List", "POST", "PUT", "DELETE"}, "/one/business", new(jasondp.DpOneBusinessAction))
			tgln.Get("/one/business/(*param)", new(jasondp.DpOneBusinessAction))

			tgln.Route([]string{"GET:List", "POST", "PUT", "DELETE"}, "/one/ib", new(jasondp.DpIbBusinessAction))
			tgln.Get("/one/ib/(*param)", new(jasondp.DpIbBusinessAction))

			tgln.Route([]string{"GET:List", "POST", "PUT", "DELETE"}, "/onepricing", new(jasondp.DpOnePricingAction))
			tgln.Get("/onepricing/(*param)", new(jasondp.DpOnePricingAction))

			tgln.Post("/one/pricing", new(engine.DpOnePricing))

			tgln.Get("/one/stock/business", new(jasondp.DpOneStockBusinessAction))

			tgln.Route([]string{"POST", "PUT", "DELETE"}, "/one/stock/pricing", new(jasondp.DpOneStockPricingAction))
			tgln.Get("/one/stock/pricing/(*param)", new(jasondp.DpOneStockPricingAction))

			tgln.Get("/one/stock/business/(*param)", new(jasondp.DpOneStockBusinessAction))

			tgln.Post("/one/begainpricing", new(jasondp.DpOneBegainPricingAction))

			// 存款差异化价格矩阵
			tgln.Get("/diff/matrix", new(jasondp.DpDiffMatrixAction))
			tgln.Get("/diff/matrix/(*param)", new(jasondp.DpDiffMatrixAction))

			tgln.Route([]string{"GET:List", "POST", "PUT", "DELETE"}, "/target/rate", new(par.DpTargetRateAction))
			tgln.Get("/target/rate/(*param)", new(par.DpTargetRateAction))
		})

		tg.Group("/par", func(tgln *tango.Group) {
			//  派生优惠参数表
			tgln.Route([]string{"GET:List", "POST", "PUT", "DELETE"}, "/scenediscount", new(par.SceneDiscountAction))
			tgln.Get("/scenediscount/(*param)", new(par.SceneDiscountAction))
		})

		// 	Janly客户分类
		tg.Group("/bi", func(tgln *tango.Group) {
			//tgln.Get("/pricelists", new(loan.PricingListAction))
			tgln.Route([]string{"POST:LoanInfo", "GET"}, "/loaninfo", new(bi.ExcelImportAction))
			tgln.Get("/loaninfo/(*param)", new(bi.ExcelImportAction))
			tgln.Route([]string{"GET:ListDate"}, "/loaninfo/listdate", new(bi.ExcelImportAction))
			// tgln.Route([]string{"POST:AmountRange"}, "/amountrange", new(bi.ExcelImportAction))
			tgln.Route([]string{"POST:ExcelImport"}, "/baseratereduction", new(dim.BaseRateReductionAction))
			tgln.Route([]string{"POST:ExcelImport"}, "/guaranteeway", new(dim.GuaranteeWayAction))
			tgln.Route([]string{"POST:ExcelImport"}, "/industry", new(dim.IndustryAction))
			tgln.Route([]string{"POST:ExcelImport"}, "/termrange", new(dim.TermRangeAction))
			tgln.Route([]string{"POST:ExcelImport"}, "/amountrange", new(dim.AmountRangeAction))

			tgln.Route([]string{"POST:ExcelImport"}, "/line", new(dim.LineAction))
			tgln.Route([]string{"POST:ExcelImport"}, "/organcr", new(dim.OrganCrAction))
			tgln.Route([]string{"POST:ExcelImport"}, "/organmc", new(dim.OrganMcAction))
			tgln.Route([]string{"POST:ExcelImport"}, "/product", new(dim.ProductAction))
			tgln.Route([]string{"POST:ExcelImport"}, "/producttype", new(dim.ProductTypeAction))
			tgln.Route([]string{"POST:ExcelImport"}, "/scale", new(dim.ScaleAction))
			tgln.Route([]string{"POST:ExcelImport"}, "/subject", new(dim.SubjectAction))

		})

		tg.Group("/login", func(tgln *tango.Group) {
			tgln.Post("/out", new(login.LoginOutAction))
		})
		tg.Post("/loginout/am", new(login.LoginOutAmAction))

		//叶全才 零售定价
		tg.Group("", func(tgln *tango.Group) {
			tgln.Route([]string{"GET:List", "POST", "PUT"}, "/rl/prcing", new(rl.RlPrcingAction))
			tgln.Get("/rl/prcing/(*param)", new(rl.RlPrcingAction))
		})

		//叶全才　存款标准化定价
		tg.Group("", func(tgln *tango.Group) {
			tgln.Route([]string{"GET:List"}, "/dp/prcing", new(dp.DpPricingAction))
			tgln.Get("/dp/prcing/(*param)", new(dp.DpPricingAction))
		})

		// Jason AM系统数据权限
		tg.Group("/am", func(tgln *tango.Group) {
			tgln.Post("/authdata", new(amAuthData.AmAuthDataAction))
		})

		// Jason WorkFlow系统访问API
		tg.Group("/work/flow", func(tgln *tango.Group) {
			tgln.Post("/ln/pricing", new(workFlow.WorkFlowLnPricingAction))
		})

	})

}
