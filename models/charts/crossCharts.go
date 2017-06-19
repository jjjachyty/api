package charts

type Series struct {
	Name string    `json:"name"`
	Data []float64 `json:"data"`
}

type AnalysisCrossDimChart struct {
	Categories []string // 横坐标
	Series     []Series // 数据
}
