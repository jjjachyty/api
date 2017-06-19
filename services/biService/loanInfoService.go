package biService

import "pccqcpa.com.cn/app/rpm/api/models/bi/ln"
import "database/sql"

type LoanInfoService struct{}

var LoanInfoModel ln.LoanInfoModel

func (LoanInfoService) BatchInsert(loanInfos []ln.LoanInfo) (sql.Result, error) {
	r, e := LoanInfoModel.BatchInsert(loanInfos)
	return r, e
}

// 查询金额根据金额倒序排
func (LoanInfoService) FindAmount(asOfDate string) ([][]float64, error) {
	return LoanInfoModel.FindAmount(asOfDate)
}

func (LoanInfoService) ListDate() ([]string, error) {
	return LoanInfoModel.ListDate()
}
