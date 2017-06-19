package rl

import (
	"platform/dbobj"

	"fmt"

	"pccqcpa.com.cn/components/zlog"
)

type RLPriceMatrixModel struct {
}

//RLPriceMatrix type 零售价格矩阵表结构体
type RLPriceMatrix struct {
	UUID                 string
	ProductClassify      string
	ProductClassifyName  string
	ProductCode          string
	ProductName          string
	OrgCode              string
	OrgName              string
	Currency             string
	CurrencyDescribe     string
	MortgageType         string
	MortgageTypeDescribe string
	SCORE                string
	OthDimOne            string
	OthDimTwo            string
	OthDimThree          string
	OthDimFour           string
	OthDimFive           string
	Amount               float64
	AmountDescribe       string
	MeanTerm             int
	MeanTermDescribe     string
	ActualTerm           int
	BottomRate           float64
	BottomFloat          float64
	TgtRate              float64
	TgtFloat             float64
	SceneRate            float64
	SceneFloat           float64
}

//查询零售价格矩阵表
func (RLPriceMatrixModel) List() ([]RLPriceMatrix, error) {
	var rlPriceMatrixs []RLPriceMatrix
	var querySQL = `SELECT THIS_.UUID,THIS_.PRODUCT_CLASSIFY,T2_.PRODUCT_NAME,THIS_.PRODUCT_CODE,T1_.PRODUCT_NAME,
coalesce(THIS_.ORG_CODE,' '),
coalesce(T6_.ORGAN_NAME,' '),
coalesce(THIS_.CURRENCY,' '),
coalesce(T7_.DICT_NAME,' '),
coalesce(THIS_.MORTGAGE_TYPE,' '),
coalesce(T5_.DICT_NAME,' '),
coalesce(THIS_.SCORE,' '),
coalesce(THIS_.OTH_DIM_ONE,' '),
coalesce(THIS_.OTH_DIM_TWO,' '),
coalesce(THIS_.OTH_DIM_THREE,' '),
coalesce(THIS_.OTH_DIM_FOUR,' '),
coalesce(THIS_.OTH_DIM_FIVE,' '),
coalesce(THIS_.AMOUNT,' '),
coalesce(T4_.DICT_NAME,' '),
THIS_.MEAN_TERM,
T3_.DICT_NAME,
THIS_.ACTUAL_TERM,
THIS_.BOTTOM_RATE,
THIS_.BOTTOM_FLOAT,
THIS_.TGT_RATE, 
THIS_.TGT_FLOAT,
coalesce(THIS_.SCENE_RATE,THIS_.TGT_RATE),
coalesce(THIS_.SCENE_FLOAT,THIS_.TGT_FLOAT)
FROM RPM_RL_PRICE_MATRIX THIS_
LEFT JOIN RPM_DIM_PRODUCT T1_ ON T1_.PRODUCT_CODE = THIS_.PRODUCT_CODE
LEFT JOIN RPM_DIM_PRODUCT T2_ ON T2_.PRODUCT_CODE = THIS_.PRODUCT_CLASSIFY
LEFT JOIN RPM_PAR_DICT T3_ ON((T3_.PARENT_DICT ='TERM' AND T3_.DICT_TYPE='RL' AND T3_.DICT_CODE = THIS_.MEAN_TERM))
LEFT JOIN RPM_PAR_DICT T4_ ON((T4_.PARENT_DICT ='AMOUNT' AND T4_.DICT_TYPE='RL' AND T4_.DICT_CODE = THIS_.AMOUNT))
LEFT JOIN RPM_PAR_DICT T5_ ON((T5_.PARENT_DICT ='GuaranteeWay' AND T5_.DICT_TYPE='BIZ' AND T4_.DICT_CODE = THIS_.MORTGAGE_TYPE))
LEFT JOIN SYS_SEC_ORGAN T6_ ON T6_.ORGAN_CODE = THIS_.ORG_CODE
LEFT JOIN RPM_PAR_DICT T7_ ON((T7_.PARENT_DICT ='CurrencyType' AND T7_.DICT_TYPE='BIZ' AND T7_.DICT_CODE = THIS_.CURRENCY))
ORDER BY THIS_.PRODUCT_CODE,THIS_.AMOUNT, THIS_.MEAN_TERM`
	zlog.Infof("零售贷款价格矩阵-SQL:%s", nil, querySQL)
	rows, err := dbobj.Default.Query(querySQL)
	if err == nil {
		var rlPriceMatrix RLPriceMatrix
		for rows.Next() {
			rows.Scan(
				&rlPriceMatrix.UUID,
				&rlPriceMatrix.ProductClassify,
				&rlPriceMatrix.ProductClassifyName,
				&rlPriceMatrix.ProductCode,
				&rlPriceMatrix.ProductName,
				&rlPriceMatrix.OrgCode,
				&rlPriceMatrix.OrgName,
				&rlPriceMatrix.Currency,
				&rlPriceMatrix.CurrencyDescribe,
				&rlPriceMatrix.MortgageType,
				&rlPriceMatrix.MortgageTypeDescribe,
				&rlPriceMatrix.SCORE,
				&rlPriceMatrix.OthDimOne,
				&rlPriceMatrix.OthDimTwo,
				&rlPriceMatrix.OthDimThree,
				&rlPriceMatrix.OthDimFour,
				&rlPriceMatrix.OthDimFive,
				&rlPriceMatrix.Amount,
				&rlPriceMatrix.AmountDescribe,
				&rlPriceMatrix.MeanTerm,
				&rlPriceMatrix.MeanTermDescribe,
				&rlPriceMatrix.ActualTerm,
				&rlPriceMatrix.BottomRate,
				&rlPriceMatrix.BottomFloat,
				&rlPriceMatrix.TgtRate,
				&rlPriceMatrix.TgtFloat,
				&rlPriceMatrix.SceneRate,
				&rlPriceMatrix.SceneFloat,
			)
			rlPriceMatrixs = append(rlPriceMatrixs, rlPriceMatrix)
		}
		fmt.Println("矩阵---------", rlPriceMatrixs)
		return rlPriceMatrixs, nil
	}
	return rlPriceMatrixs, err
}

func (RLPriceMatrixModel) Find(rlPriceMatrix RLPriceMatrix) ([]RLPriceMatrix, error) {
	var rlPriceMatrixs []RLPriceMatrix
	var querySQL = `SELECT THIS_.BOTTOM_RATE,THIS_.TGT_RATE,THIS_.SCENE_RATE FROM RPM_RL_PRICE_MATRIX THIS_ WHERE THIS_.PRODUCT_CODE=:1 AND  
THIS_.AMOUNT = (SELECT MAX(AMOUNT) FROM RPM_RL_PRICE_MATRIX THIS_ WHERE THIS_.PRODUCT_CODE=:2 AND THIS_.AMOUNT <=:3)
AND
THIS_.MEAN_TERM=(SELECT MAX(MEAN_TERM) FROM RPM_RL_PRICE_MATRIX THIS_ WHERE THIS_.PRODUCT_CODE=:4 AND THIS_.MEAN_TERM <=:5)`
	zlog.Infof("零售贷款定价-SQL:%s", nil, querySQL)
	rows, err := dbobj.Default.Query(querySQL, rlPriceMatrix.ProductCode, rlPriceMatrix.ProductCode, rlPriceMatrix.Amount, rlPriceMatrix.ProductCode, rlPriceMatrix.MeanTerm)
	if nil == err {
		var rlPriceMatrix RLPriceMatrix
		for rows.Next() {
			rows.Scan(&rlPriceMatrix.BottomRate, &rlPriceMatrix.TgtRate, &rlPriceMatrix.SceneRate)
			rlPriceMatrixs = append(rlPriceMatrixs, rlPriceMatrix)
		}
		return rlPriceMatrixs, nil
	}
	return rlPriceMatrixs, err
}

//叶全才
func (RLPriceMatrixModel) GetRate(productCode string, amountAct float64, meanTermAct int) ([]RLPriceMatrix, error) {
	var rlPriceMatrixs []RLPriceMatrix
	var querySQL = `SELECT THIS_.TGT_RATE FROM RPM_RL_PRICE_MATRIX THIS_ WHERE THIS_.PRODUCT_CODE=:1 AND THIS_.AMOUNT =(
SELECT MAX(T2_.AMOUNT) FROM RPM_RL_PRICE_MATRIX T2_ WHERE T2_.amount <= :2)
 AND THIS_.MEAN_TERM =(
 SELECT MAX(T3_.MEAN_TERM) FROM RPM_RL_PRICE_MATRIX T3_ WHERE T3_.MEAN_TERM <= :3)`
	zlog.Infof("零售贷款利率获取-SQL:%s", nil, querySQL)
	rows, err := dbobj.Default.Query(querySQL, productCode, amountAct, meanTermAct)
	if err == nil {
		var rlPriceMatrix RLPriceMatrix
		for rows.Next() {
			rows.Scan(&rlPriceMatrix.TgtRate)
			rlPriceMatrixs = append(rlPriceMatrixs, rlPriceMatrix)
		}
		return rlPriceMatrixs, nil
	}
	return rlPriceMatrixs, err
}
