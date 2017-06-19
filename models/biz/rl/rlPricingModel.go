package rl

import (
	"database/sql"
	"fmt"
	"pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/components/zlog"
	"platform/dbobj"
	"time"

	"pccqcpa.com.cn/app/rpm/api/models/biz/ln"
	"pccqcpa.com.cn/app/rpm/api/models/dim"
	modlesUtil "pccqcpa.com.cn/app/rpm/api/models/util"

	"strconv"
	"strings"
)

type RLPricingModel struct {
}

//RlPricingModel type 零售贷款定价单实体
type RlPricing struct {
	UUID        string      //主键
	Cust        ln.CustInfo //客户ln.CustInfo
	CustManager string      //客户经理
	Product     dim.Product //产品dim.Product
	Amount      float64     //金额
	Term        int         //期限
	TermMult    string      //期限单位
	BottomRate  float64     //保本利率
	SceneRate   float64     //优惠利率
	TgtRate     float64     //目标利率
	Eva         float64     //EVA
	Raroc       float64     //RAROC
	Status      string      //定价状态
	Remark      string      //备注
	CreateTime  time.Time   //创建时间
	UpdateTime  time.Time   //更新时间
	CreateUser  string      //创建人
	UpdateUser  string      //更新人
}

func (r *RlPricing) scan(rows *sql.Rows) (*RlPricing, error) {
	var rlPricing = new(RlPricing)
	values := []interface{}{
		&rlPricing.UUID,
		// &rlPricing.Cust.CustCode,
		&rlPricing.CustManager,
		// &rlPricing.Product,
		&rlPricing.Amount,
		&rlPricing.Term,
		&rlPricing.TermMult,
		&rlPricing.BottomRate,
		&rlPricing.SceneRate,
		&rlPricing.TgtRate,
		&rlPricing.Eva,
		&rlPricing.Raroc,
		&rlPricing.Status,
		&rlPricing.Remark,
		&rlPricing.CreateTime,
		&rlPricing.UpdateTime,
		&rlPricing.CreateUser,
		&rlPricing.UpdateUser,

		&rlPricing.Cust.CustCode,
		&rlPricing.Cust.CustName,
		&rlPricing.Cust.Organization,
		&rlPricing.Cust.CustType,
		&rlPricing.Cust.CustImplvl,
		&rlPricing.Cust.CustCredit,
		&rlPricing.Cust.Branch.OrganCode,
		&rlPricing.Cust.Industry.IndustryCode,
		&rlPricing.Cust.CustSize,
		&rlPricing.Cust.CustCapital,
		&rlPricing.Cust.GapProportion,
		&rlPricing.Cust.StockContribute,
		&rlPricing.Cust.StockFreeze,
		&rlPricing.Cust.StockUsage,
		&rlPricing.Cust.UseProduct,
		&rlPricing.Cust.CooperationPeriod,

		&rlPricing.Product.ProductCode,
		&rlPricing.Product.ProductName,
		&rlPricing.Product.ProductType,
		&rlPricing.Product.ProductTypeDesc,
		&rlPricing.Product.ProductLevel,
	}
	err := util.OracleScan(rows, values)
	return rlPricing, err
}

//分页查询"我的贷款"
func (rlp *RlPricing) List(param ...map[string]interface{}) (*util.PageData, error) {

	pageData, rows, err := modlesUtil.List(rlPricingTables, rlPricingCols, rlPricingSort, param...)
	if nil != err {
		er := fmt.Errorf("分页查询指标出错")
		zlog.Error(er.Error(), err)
		return nil, er
	}
	defer rows.Close()
	var rlpricings []*RlPricing
	for rows.Next() {
		rlPricing, err := rlp.scan(rows)
		if nil != err {
			er := fmt.Errorf("分页查询指标rows.Scan()出错")
			zlog.Error(er.Error(), err)
			return nil, er
		}
		rlpricings = append(rlpricings, rlPricing)
	}
	pageData.Rows = rlpricings
	return pageData, nil
}

//分页查询客户经理下的零售定价单
func (rlp *RlPricing) ListCusMan(param ...map[string]interface{}) (*util.PageData, error) {
	var paramMap map[string]interface{}
	if 0 == len(param) {
		paramMap = make(map[string]interface{})
	} else {
		paramMap = param[0]
	}
	var addSql string = ""
	if flag, ok := paramMap["flag"]; ok && flag.(bool) {
		//处理条件查询参数
		if status, ok := paramMap["t1.status"]; ok {
			addSql += ` and (upper(t1.status)=` + status.(string) + `)`
		}
		if custCode, ok := paramMap["t2.cust_code"]; ok {
			custName, _ := paramMap["t2.cust_name"]
			addSql += ` and (
				upper(t2.cust_code) like upper('%%` + custCode.(string) + `%%')
				or upper(t2.cust_name) like upper('%%` + custName.(string) + `%%')
				)`
		}
	}

	zlog.Info("开始分页查询", nil)
	var tableAlias string = ""
	orderBy := getOrderBy(tableAlias, paramMap)
	var querySQL = `select uuid,
  coalesce(cust_manager,' '),
  coalesce(amount,0),
  coalesce(term,0),
  coalesce(term_mult,' '),
  coalesce(bottom_rate,0),
  coalesce(scene_rate,0),
  coalesce(tgt_rate,0),
  coalesce(eva,0),
  coalesce(raroc,0),
  coalesce(status,'0'),
  coalesce(remark,' '),
  coalesce(create_time,sysdate),
  coalesce(update_time,sysdate),
  coalesce(create_user,' '),
  coalesce(update_user,' '),
  coalesce(cust_code,' '),
  coalesce(cust_name,' '),
  coalesce(organization,' '),
  coalesce(cust_type,' '),
  coalesce(cust_implvl,' '),
  coalesce(cust_credit,' '),
  coalesce(branch,' '),
  coalesce(industry,' '),
  coalesce(cust_size,' '),
  coalesce(cust_capital,0),
  coalesce(gap_proportion,0),
  coalesce(stock_contribute,0),
  coalesce(stock_freeze,0),
  coalesce(stock_usage,0),
  coalesce(use_product,0),
  coalesce(cooperation_period,0),
  coalesce(product_code,' '),
  coalesce(product_name,' '),
  coalesce(product_type,' '),
  coalesce(product_type_desc,' '),
  coalesce(product_level,' ')
from
  (select t.*,
    rownum as num
  from
    (select t1.uuid,
      t1.cust_manager,
      t1.amount,
      t1.term,
      t1.term_mult,
      t1.bottom_rate,
      t1.scene_rate,
      t1.tgt_rate,
      t1.eva,
      t1.raroc,
      t1.status,
      t1.remark,
      t1.create_time,
      t1.update_time,
      t1.create_user,
      t1.update_user,
      t2.cust_code,
      t2.cust_name,
      t2.organization,
      t2.cust_type,
      t2.cust_implvl,
      t2.cust_credit,
      t2.branch,
      t2.industry,
      t2.cust_size,
      t2.cust_capital,
      t2.gap_proportion,
      t2.stock_contribute,
      t2.stock_freeze,
      t2.stock_usage,
      t2.use_product,
      t2.cooperation_period,
      t4.product_code,
      t4.product_name,
      t4.product_type,
      t4.product_type_desc,
      t4.product_level
    from rpm_biz_rl_pricing t1
    left join rpm_biz_cust_info t2
    on (t1.cust = t2.cust_code)
    left join rpm_dim_product t4
    on (t1.product = t4.product_code)
    where 1        =1
    and 
    ( (t1.cust_manager) = :1
    	or upper(t1.cust_manager) is null )
    ` + addSql + `

	` + orderBy + `
    ) t
  where rownum <= :2
  )
where num > :3`
	//初始化返回结果
	var pageData = new(util.PageData)
	//获取分页数据
	start, length := getPageCount(paramMap)

	delete(paramMap, "start")
	delete(paramMap, "length")

	custManager, _ := paramMap["cust_manager"]
	//查询sql
	rows, err := dbobj.Default.Query(querySQL, custManager, strconv.Itoa(start+length), strconv.Itoa(start))
	if nil != err {
		er := fmt.Errorf("分页查询指标出错")
		zlog.Error(er.Error(), err)
		return nil, er
	}
	defer rows.Close()
	zlog.Infof("\n分页查询sql: %s\n参数\n%s", nil, querySQL)
	var rlpricings []*RlPricing
	for rows.Next() {
		rlPricing, err := rlp.scan(rows)
		if nil != err {
			er := fmt.Errorf("分页查询指标rows.Scan()出错")
			zlog.Error(er.Error(), err)
			return nil, er
		}
		rlpricings = append(rlpricings, rlPricing)
	}
	//查询总条数
	var total int64
	countSql := `select count(1) from rpm_biz_rl_pricing t1
	left join rpm_biz_cust_info t2
    on (t1.cust = t2.cust_code)
    where 1=1
	and ( 
	(t1.cust_manager) = :1
	or (t1.cust_manager) is null )
	` + addSql + ``

	countRows, err := dbobj.Default.Query(countSql, custManager)
	defer countRows.Close()
	if nil != err {
		zlog.Errorf("分页查询总数出错\nsql:", err, countSql)
	}
	for countRows.Next() {
		err = countRows.Scan(&total)
		countRows.Next()
		if nil != err {
			zlog.Errorf("查询分页总数row.Scan()报错", err)
			return nil, err
		}
	}
	//排序方式
	pageData.Rows = rlpricings
	pageData.Page.TotalRows = total
	return pageData, nil
}

func (r *RlPricing) Add() error {
	param := map[string]interface{}{
		"cust":         r.Cust.CustCode,
		"cust_manager": r.CustManager,
		"product":      r.Product.ProductCode,
		"amount":       r.Amount,
		"term":         r.Term,
		"term_mult":    r.TermMult,
		"bottom_rate":  r.BottomRate,
		"scene_rate":   r.SceneRate,
		"tgt_rate":     r.TgtRate,
		"eva":          r.Eva,
		"raroc":        r.Raroc,
		"status":       r.Status,
		"remark":       r.Remark,
		"create_time":  util.GetCurrentTime(),
		"update_time":  util.GetCurrentTime(),
		"create_user":  r.CreateUser,
		"update_user":  r.UpdateUser,
	}
	err := util.OracleAdd(rlPricingTableName, param)
	if nil != err {
		er := fmt.Errorf("新增零售货款定价信息出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

//修改零售货款定价状态
func (r *RlPricing) Update() error {
	param := map[string]interface{}{
		"status": r.Status,
		"remark": r.Remark,
	}
	whereParam := map[string]interface{}{
		"uuid": r.UUID,
	}
	err := util.OracleUpdate(rlPricingTableName, param, whereParam)
	if nil != err {
		er := fmt.Errorf("更新零售定价状态信息出错")
		zlog.Errorf(er.Error(), err)
		return er
	}
	return nil
}

var rlPricingTableName = "RPM_BIZ_RL_PRICING"

var rlPricingTables string = `
	RPM_BIZ_RL_PRICING T
		LEFT JOIN RPM_BIZ_CUST_INFO CUST
			ON (T.CUST = CUST.CUST_CODE)
		LEFT JOIN RPM_DIM_PRODUCT PRODUCT
			ON (T.PRODUCT = PRODUCT.PRODUCT_CODE)
`

var rlPricingCols = map[string]string{
	"T.UUID": "' '",
	// "t.cust":         "' '",
	"T.CUST_MANAGER": "' '",
	// "T4.PRODUCT_NAME": "' '",
	"T.AMOUNT":                  "0",
	"T.TERM":                    "0",
	"T.TERM_MULT":               "' '",
	"T.BOTTOM_RATE":             "0",
	"T.SCENE_RATE":              "0",
	"T.TGT_RATE":                "0",
	"T.EVA":                     "0",
	"T.RAROC":                   "0",
	"T.STATUS":                  "'0'",
	"T.REMARK":                  "' '",
	"T.CREATE_TIME":             "sysdate",
	"T.UPDATE_TIME":             "sysdate",
	"T.CREATE_USER":             "' '",
	"T.UPDATE_USER":             "' '",
	"CUST.CUST_CODE":            "' '",
	"CUST.CUST_NAME":            "' '",
	"CUST.ORGANIZATION":         "' '",
	"CUST.CUST_TYPE":            "' '",
	"CUST.CUST_IMPLVL":          "' '",
	"CUST.CUST_CREDIT":          "' '",
	"CUST.BRANCH":               "' '",
	"CUST.INDUSTRY":             "' '",
	"CUST.CUST_SIZE":            "' '",
	"CUST.CUST_CAPITAL":         "0",
	"CUST.GAP_PROPORTION":       "0",
	"CUST.STOCK_CONTRIBUTE":     "0",
	"CUST.STOCK_FREEZE":         "0",
	"CUST.STOCK_USAGE":          "0",
	"CUST.USE_PRODUCT":          "0",
	"CUST.COOPERATION_PERIOD":   "0",
	"PRODUCT.PRODUCT_CODE":      "' '",
	"PRODUCT.PRODUCT_NAME":      "' '",
	"PRODUCT.PRODUCT_TYPE":      "' '",
	"PRODUCT.PRODUCT_TYPE_DESC": "' '",
	"PRODUCT.PRODUCT_LEVEL":     "' '",
}
var rlPricingSort = []string{
	"T.UUID",
	// "T.CUST",
	"T.CUST_MANAGER",
	// "PRODUCT.PRODUCT_NAME",
	"T.AMOUNT",
	"T.TERM",
	"T.TERM_MULT",
	"T.BOTTOM_RATE",
	"T.SCENE_RATE",
	"T.TGT_RATE",
	"T.EVA",
	"T.RAROC",
	"T.STATUS",
	"T.REMARK",
	"T.CREATE_TIME",
	"T.UPDATE_TIME",
	"T.CREATE_USER",
	"T.UPDATE_USER",
	"CUST.CUST_CODE",
	"CUST.CUST_NAME",
	"CUST.ORGANIZATION",
	"CUST.CUST_TYPE",
	"CUST.CUST_IMPLVL",
	"CUST.CUST_CREDIT",
	"CUST.BRANCH",
	"CUST.INDUSTRY",
	"CUST.CUST_SIZE",
	"CUST.CUST_CAPITAL",
	"CUST.GAP_PROPORTION",
	"CUST.STOCK_CONTRIBUTE",
	"CUST.STOCK_FREEZE",
	"CUST.STOCK_USAGE",
	"CUST.USE_PRODUCT",
	"CUST.COOPERATION_PERIOD",
	"PRODUCT.PRODUCT_CODE",
	"PRODUCT.PRODUCT_NAME",
	"PRODUCT.PRODUCT_TYPE",
	"PRODUCT.PRODUCT_TYPE_DESC",
	"PRODUCT.PRODUCT_LEVEL",
}

//处理分页数据
func getPageCount(paramMap map[string]interface{}) (int, int) {
	var start, length int
	_, ok := paramMap["start"]
	if ok {
		start, _ = strconv.Atoi(paramMap["start"].(string))
	} else {
		start = 0
	}

	_, ok = paramMap["length"]
	if ok {
		length, _ = strconv.Atoi(paramMap["length"].(string))
	} else {
		length = 20
	}
	return start, length
}

//获取排序方式
func getOrderBy(tableAlias string, paramMap map[string]interface{}) string {
	orderBy := ""
	sort, ok := paramMap["sort"]
	// fmt.Println("－－－－－－－－", ok, sort)
	if ok && "" != sort {
		// 判断是否为关联表排序
		var orderStrings = strings.Split(sort.(string), ".")
		var orderStringsLength = len(orderStrings)
		if 1 < orderStringsLength {
			orderBy += " order by " + orderStrings[orderStringsLength-2] + "." + util.UperChange(orderStrings[orderStringsLength-1]) + " "
		} else {
			orderBy += " order by " + util.UperChange(sort.(string)) + " "
		}
	} else {
		//默认排序方式
		orderBy += " order by " + tableAlias + "create_time "
		paramMap["order"] = "DESC"
	}
	order, ok := paramMap["order"]
	if ok {
		orderBy += order.(string)
	}
	return orderBy
}
