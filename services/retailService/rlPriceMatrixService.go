package retailService

import (
	"errors"

	"pccqcpa.com.cn/app/rpm/api/models/biz/rl"
)

type RLPriceMatrixService struct {
}

var rlPriceMatrixModel rl.RLPriceMatrixModel

func (RLPriceMatrixService) GetAllRLPriceMatrix() ([]rl.RLPriceMatrix, error) {
	return rlPriceMatrixModel.List()
}

func (RLPriceMatrixService) GetRLPriceMatrix(rlPriceMatrix rl.RLPriceMatrix) (rl.RLPriceMatrix, error) {
	rlPriceMatrixs, err := rlPriceMatrixModel.Find(rlPriceMatrix)
	if nil == err {
		length := len(rlPriceMatrixs)
		switch length {
		case 1:
			return rlPriceMatrixs[0], nil
		case 0:
			return rlPriceMatrix, errors.New("零售定价查询数据为空，请检查表[RPM_RETAIL_PRICE_MATRIX]")
		default:
			return rlPriceMatrix, errors.New("零售定价查询数据重复，请检查表[RPM_RETAIL_PRICE_MATRIX]")
		}

	}
	return rlPriceMatrix, err
}

//叶全才
func (RLPriceMatrixService) GetRate(productCode string, amount float64, deadline int, unit string) (rl.RLPriceMatrix, error) {
	var rlPriceMatrix rl.RLPriceMatrix
	var meanTerm int
	if unit == "Y" {
		meanTerm = deadline * 360
	} else if unit == "M" {
		meanTerm = deadline * 30
	} else if unit == "D" {
		meanTerm = deadline
	} else {
		//非法数据
	}
	rlPriceMatrixs, err := rlPriceMatrixModel.GetRate(productCode, amount, meanTerm)
	if nil == err {
		length := len(rlPriceMatrixs)
		switch length {
		case 1:
			return rlPriceMatrixs[0], nil
		case 0:
			return rlPriceMatrix, errors.New("利率浮动查询数据为空，请检查表[RPM_RETAIL_PRICE_MATRIX]")
		default:
			return rlPriceMatrix, errors.New("利率定价查询数据重复，请检查表[RPM_RETAIL_PRICE_MATRIX]")
		}
	}
	return rlPriceMatrix, err
}
