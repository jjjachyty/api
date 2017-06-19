package util

// 客户状态
const (
	SIMULATION_CUST string = "01" //模拟客户
	REAL_CUST       string = "02" //存量客户
)

const (
	IS_PAGIN_QUERY   string = "IsPaginQuery"
	SEARCH_LIKE      string = "searchLike"
	PRICING_BUS_DICT string = "PricingBusDict"
	DP_BASE_RATE_GP  string = "1"
	RPM_CM           string = "rpmcm"
	RPM_CUST_OWNER   string = "owner"
)

// 表名
const (
	ORGAN    string = "sys_sec_organ"    //机构表
	INDUSTRY string = "rpm_dim_industry" //行业表
	PRODUCT  string = "rpm_dim_product"  //产品表
)

// 指标业务类型
const (
	FIRST_INDICATOR  string = "F" //公式管理
	SECOND_INDICATOR string = "I" //公式指标
	THIRD_INCATOR    string = "B" //基础指标
)

// 业务类型
const (
	LN_BUSINESS        string = "LN"                 // 对公贷款
	LN_SCENE           string = "LN_SCENE"           // 对公情景
	LN_INVERSE         string = "LN_INVERSE"         // 对公反算
	DP_BUSINESS        string = "DP"                 // 存款
	ONE_DP             string = "ONE_DP"             // 一对一存款
	ONE_DP_CURRENT_EVA string = "ONE_DP_CURRENT_EVA" // 存款一对一当前EVA
	ONE_DP_IB_EVA      string = "ONE_DP_IB_EVA"      // 存款一对一派生EVA
	ONE_DP_STOCK       string = "ONE_DP_STOCK"       // 存款一对一存量
)

// 参数类型param——type
const (
	PARAM_DP string = "DP" //存款
	PARAM_LN string = "LN" //贷款
)

// 树形结构总节点
const (
	TREE_TOP_CODE = "#"
	DICT_TOP_CODE = "*"
)

// 计算引擎函数指标报错标志
const (
	CALCULATE_ERROR_KEY = "error"
	ERROR               = -1
	SUCCESS             = 0
)

//生效标志
const (
	FLAG_TRUE  = "1"
	FLAG_FALSE = "0"
)

// 定价状态
const (
	// PRICING_STATUS_UNFINISHED         = "0"  // 未完成状态
	PRICING_STATUS_UNFINISHED_LN      = "-1" // 未完成计算贷款状态
	PRICING_STATUS_UNFINISHED_INVERSE = "0"  // 未完成利率反算状态
	PRICING_STATUS_FINISHED_NO_SAVE   = "1"  // 计算完成未保存状态
	PRICING_STATUS_FINISHED_SAVE      = "2"  // 计算完成并保存状态

	DP_ONE_PRICING_STATUS_UNFINISHED = "0" // 存款一对一定价未完成状态
	DP_ONE_PRICING_STATUS_FINISHED   = "1" // 存款一对一定价完成状态
)

// 缓存名称
const (
	RPM_LGD_CACHE         string = "rpmLgdCache"
	RPM_INDICATOR_CACHE   string = "rpmIndicatorCache"
	RPM_PARENT_DICT_CACHE string = "rpmParentDictCache"
	AM_DATA_AUTH_CACHE    string = "amDataAuthCache"
	RPM_SQUEEZE_CACHE     string = "rpmSqueezeCache"
	RPM_SID_USER_CACHE    string = "rpmSiduserCache"
)

// 税率类型
const (
	INCOME_TAX string = "Income"
	ADD_TAX    string = "Add"
)

// 折让率类型
const (
	DISCOUNT_STOCK   = "1" // 存量
	DISCOUNT_DERIVED = "2" // 派生
)

// 主担保方式
const (
	MAIN_MORTGAGE_TYPE_CREDIT = "4"
)

// LGD 信用或者未覆盖的的最低违约损失率
const CREDIT_LOW_LGD float64 = 0.45

// 零售贷款定价审核状态
const RL_PRICING_STATUS string = "0" //提交审批

// rpm_par_common表key
const (
	PREMIUM_RATE string = "PREMIUM_RATE" // 存款保险费率
)

const ONE_YEAR_TERM int = 360 // 一年期限值

const (
	USE_PRODUCT        string = "product"     // 使用产品数
	COOPERATION_PERIOD string = "cooperation" // 合作年限数
)

const (
	NUM_300 string = "300 "
	NUM_301 string = "301 "
)
