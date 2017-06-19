package par

import (
	"database/sql"
	"platform/dbobj"
	"time"

	"pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/components/zlog"
)

type NullSceneDiscount struct {
	UUID        sql.NullString
	BizType     sql.NullString
	CustImplvl  sql.NullString
	CustSize    sql.NullString
	ProductCode sql.NullString
	Rate        sql.NullFloat64
	CreateUser  sql.NullString
	UpdateUser  sql.NullString
	CreateTime  time.Time
	UpdateTime  time.Time
}
type SceneDiscountModel struct {
}

func (SceneDiscountModel) QueryDiscountRate(discount NullSceneDiscount) ([]NullSceneDiscount, error) {
	var discounts []NullSceneDiscount
	var rows *sql.Rows
	var err error
	var querySQL = "SELECT * FROM RPM_PAR_SCENE_DISCOUNT DIS_ WHERE "
	if util.DISCOUNT_STOCK == discount.BizType.String {
		querySQL = querySQL + "BIZ_TYPE = '1' AND CUST_IMPLVL ＝:0 AND CUST_SIZE=:1"
	} else if util.DISCOUNT_DERIVED == discount.BizType.String {
		querySQL = querySQL + "BIZ_TYPE = '2' AND CUST_IMPLVL ＝:0 AND CUST_SIZE=:1 AND PRODUCT=:2"
	}
	//fmt.Printf("SQL:%s---%v", querySQL, discount)
	zlog.Infof("SQL-:%s BIZ_TYPE=%v CUST_IMPLVL=%v CUST_SIZE=%v PRODUCT=%v", nil, querySQL, discount.BizType, discount.CustImplvl, discount.CustSize, discount.ProductCode)

	if util.DISCOUNT_STOCK == discount.BizType.String {
		rows, err = dbobj.Default.Query(querySQL, discount.CustImplvl.String, discount.CustSize.String)
	} else if util.DISCOUNT_DERIVED == discount.BizType.String {
		rows, err = dbobj.Default.Query(querySQL, discount.CustImplvl.String, discount.CustSize.String, discount.ProductCode.String)
	}
	defer rows.Close()
	if err == nil {
		discounts = NullSceneDiscount{}.getResult(rows)
	}
	return discounts, err
}

func (NullSceneDiscount) getResult(rows *sql.Rows) []NullSceneDiscount {
	var discounts []NullSceneDiscount
	defer rows.Close()
	for rows.Next() {
		var discount NullSceneDiscount
		rows.Scan(&discount.UUID, &discount.BizType, &discount.CustImplvl, &discount.CustSize, &discount.ProductCode, &discount.Rate, &discount.CreateUser, &discount.UpdateUser, &discount.CreateTime, &discount.UpdateTime)

		discounts = append(discounts, discount)
	}
	return discounts
}
