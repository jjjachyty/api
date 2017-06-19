package dim

import (
	"database/sql"
	"fmt"
	"time"

	"pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/components/zlog"

	modelsUtil "pccqcpa.com.cn/app/rpm/api/models/util"
)

type Product struct {
	UUID            string    `xorm:"varchar(44) notnull unique 'UUID'"`
	ProductCode     string    `xorm:"varchar(44) notnull unique 'PRODUCT_CODE'"`
	ProductName     string    `xorm:"varchar(44) notnull  		'PRODUCT_NAME'"`
	ProductType     string    `xorm:"varchar(44) notnull  		'PRODUCT_TYPE'"`
	ProductTypeDesc string    `xorm:"varchar(44)   			    'PRODUCT_TYPE_DESC'"`
	ProductLevel    string    `xorm:"varchar(44) notnull  		'PRODUCT_LEVEL'"`
	ParentProduct   *Product  `xorm:"extends"`
	CreateTime      time.Time `xorm:"DATE created               'CREATE_TIME'"`
	CreateUser      string    `xorm:"VARCHAR2(44)               'CREATE_USER'"`
	UpdateTime      time.Time `xorm:"DATE updated               'UPDATE_TIME'"`
	UpdateUser      string    `xorm:"VARCHAR2(44)               'UPDATE_USER'"`
	LeafFlag        string    `xorm:"varchar(44)                'PRODUCT_LEVEL'"`
	Flag            string    `xorm:"varchar(44)                'FLAG'"`
}
type NullProduct struct {
	UUID            sql.NullString
	ProductCode     sql.NullString
	ProductName     sql.NullString
	ProductType     sql.NullString
	ProductTypeDesc sql.NullString
	ProductLevel    sql.NullString
	ParentProduct   *Product
	CreateTime      time.Time
	CreateUser      sql.NullString
	UpdateTime      time.Time
	UpdateUser      sql.NullString
	LeafFlag        sql.NullString
	Flag            sql.NullString
}

var productTables string = `
    RPM_DIM_PRODUCT T`

var productCols = map[string]string{
	"T.UUID":              "' '",
	"T.PRODUCT_CODE":      "' '",
	"T.PRODUCT_NAME":      "' '",
	"T.PRODUCT_TYPE":      "' '",
	"T.PRODUCT_TYPE_DESC": "' '",
	"T.PRODUCT_LEVEL":     "' '",
	"T.PARENT_PRODUCT":    "' '",
	"T.CREATE_TIME":       "sysdate",
	"T.CREATE_USER":       "' '",
	"T.UPDATE_TIME":       "sysdate",
	"T.UPDATE_USER":       "' '",
	"T.LEAF_FLAG":         "' '",
	"T.FLAG":              "' '",
}

var productColsSort = []string{
	"T.UUID",
	"T.PRODUCT_CODE",
	"T.PRODUCT_NAME",
	"T.PRODUCT_TYPE",
	"T.PRODUCT_TYPE_DESC",
	"T.PRODUCT_LEVEL",
	"T.PARENT_PRODUCT",
	"T.CREATE_TIME",
	"T.CREATE_USER",
	"T.UPDATE_TIME",
	"T.UPDATE_USER",
	"T.LEAF_FLAG",
	"T.FLAG",
}

func (this *Product) scanProduct(rows *sql.Rows) (*Product, error) {
	var product = new(Product)
	var parentProduct = new(Product)
	var values = []interface{}{
		&product.UUID,
		&product.ProductCode,
		&product.ProductName,
		&product.ProductType,
		&product.ProductTypeDesc,
		&product.ProductLevel,
		&parentProduct.ProductCode,
		&product.CreateTime,
		&product.CreateUser,
		&product.UpdateTime,
		&product.UpdateUser,
		&product.LeafFlag,
		&product.Flag,
	}
	product.ParentProduct = parentProduct
	err := util.OracleScan(rows, values)
	if nil != err {
		return nil, err
	}
	return product, nil
}

func (this *Product) SelectPorductByParams(paramMap map[string]interface{}) ([]*Product, error) {
	rows, err := modelsUtil.FindRows(productTables, productCols, productColsSort, paramMap)
	if nil != err {
		var er error = fmt.Errorf("多参数查询产品错误")
		zlog.Error(er.Error(), err)
		return nil, er
	}
	defer rows.Close()
	var products []*Product
	for rows.Next() {
		product, err := this.scanProduct(rows)
		if nil != err {
			var er = fmt.Errorf("多参数查询产品rows.Scan()错误")
			zlog.Error(er.Error(), err)
			return nil, er
		}
		products = append(products, product)
	}
	return products, nil
}
