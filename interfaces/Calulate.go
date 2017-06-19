package interfaces

type Calulate interface {
	Calulate(paramMap map[string]interface{}) (float64, error)
}
