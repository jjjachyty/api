package dim

import (
	"time"
)

type Mortgage struct {
	UUID           string    //主键
	MortgageCode   string    //押品编号
	MortgageName   string    //押品名称
	ParentMortgage string    //父级押品
	MortgageBasel  string    //抵押品－辛巴
	LeafFlag       string    //是否有叶子节点（1:有 0：无）
	Flag           string    //生效标志
	CreateTime     time.Time //创建时间
	CreateUser     string    //创建人
	UpdateTime     time.Time //更新时间
	UpdateUser     string    //更新人
}
