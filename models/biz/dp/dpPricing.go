package dp

import (
	// "pccqcpa.com.cn/app/rpm/api/models/dim"
	// "fmt"
	"platform/dbobj"
	"time"

	"pccqcpa.com.cn/components/zlog"
)

type DpPricingModel struct {
}

//存款标准化定价单结果实体
type DpPricing struct {
	UUID            string    //主键
	ProductClassify string    //产品大类
	ProductName     string    //存款产品
	Term            string    //期限
	FtpRate         float64   //资金收益
	OcRate          float64   //运营费用率
	PremiumRate     float64   //保险费率
	NumOne          float64   //战略调整项
	BaseRate        float64   //基准利率
	TgtRate         float64   //目标收益率
	OpRate          float64   //操作风险率
	DpRate          float64   //标准化价格
	FloatingRatio   float64   //上浮比例
	Flag            string    //生效标志
	CreateTime      time.Time //创建时间
	UpdateTime      time.Time //更新时间
	CreateUser      string    //创建人
	Updateuser      string    //更新人
}

//查询存款标准化定价单结果表 20161214
func (DpPricingModel) List(whereSql string) ([]DpPricing, error) {
	var dpPricings []DpPricing
	var querySQL = `SELECT THIS_.UUID,T1_.PRODUCT_NAME,T2_.PRODUCT_NAME,THIS_.TERM,
	THIS_.FTP_RATE,THIS_.OC_RATE,THIS_.PREMIUM_RATE,THIS_.NUM_ONE,THIS_.BASE_RATE,
	THIS_.TGT_RATE,THIS_.OP_RATE,THIS_.DP_RATE,THIS_.FLOATING_RATIO,THIS_.FLAG,THIS_.CREATE_TIME,THIS_.CREATE_USER,
	THIS_.UPDATE_TIME,THIS_.UPDATE_USER
	FROM RPM_BIZ_DP_STANDARD THIS_
	LEFT JOIN RPM_DIM_PRODUCT T1_ ON T1_.PRODUCT_CODE = THIS_.PRODUCT_CLASSIFY
	LEFT JOIN RPM_DIM_PRODUCT T2_ ON T2_.PRODUCT_CODE = THIS_.PRODUCT
	` + whereSql + `
	ORDER BY THIS_.CREATE_TIME DESC,  THIS_.PRODUCT,THIS_.TERM ASC
	`
	zlog.Infof("存款标准化定价单-SQL%s", nil, querySQL)
	rows, err := dbobj.Default.Query(querySQL)
	if err == nil {
		var dpPricing DpPricing
		for rows.Next() {
			rows.Scan(
				&dpPricing.UUID,
				&dpPricing.ProductClassify,
				&dpPricing.ProductName,
				&dpPricing.Term,
				&dpPricing.FtpRate,
				&dpPricing.OcRate,
				&dpPricing.PremiumRate,
				&dpPricing.NumOne,
				&dpPricing.BaseRate,
				&dpPricing.TgtRate,
				&dpPricing.OpRate,
				&dpPricing.DpRate,
				&dpPricing.FloatingRatio,
				&dpPricing.Flag,
				&dpPricing.CreateTime,
				&dpPricing.CreateUser,
				&dpPricing.UpdateTime,
				&dpPricing.Updateuser,
			)
			dpPricings = append(dpPricings, dpPricing)
		}
		return dpPricings, nil
	}
	return dpPricings, err
}
