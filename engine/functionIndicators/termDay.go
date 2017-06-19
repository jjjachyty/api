package functionIndicators

import (
	"fmt"
)

type TermDay struct {
}

// 查询所得税
func (day *TermDay) Calulate(paramMap map[string]interface{}) (float64, error) {
	var termDay int
	var err error
	if v, ok := paramMap["term"]; ok {
		termDay = v.(int)
	}
	if v, ok := paramMap["term_mult"]; ok {
		switch v.(string) {
		case "D":
		case "M":
			termDay = termDay * 30
		case "Y":
			termDay = termDay * 360
		default:
			err = fmt.Errorf("期限单位只指出年(Y)月(M)日(D)")
		}
	}
	return float64(termDay), err
}
