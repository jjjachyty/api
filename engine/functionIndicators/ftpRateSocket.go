package functionIndicators

import (
	"fmt"
)

type FtpRateSocket struct {
}

// FTPRATE
func (this *FtpRateSocket) Calulate(paramMap map[string]interface{}) float64 {
	fmt.Println("开始计算资金成本率街口试算")
	return 2.2
}
