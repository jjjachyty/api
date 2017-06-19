package util

import (
	"time"
)

type BaseEntity struct {
	Uuid       string    //主键
	StartDate  time.Time //创建时间
	UpdateDate time.Time //更新时间
	CreateUser string    //创建人
	Flag       string    //生效标志
	ResChr1    string    //预留字符字段1
	Reschr2    string    //预留字符字段2
	ResChr3    string    //预留字符字段3
	Reschr4    string    //预留字符字段4
	ResChr5    string    //预留字符字段5
	Reschr6    string    //预留字符字段6
	ResChr7    string    //预留字符字段7
	Reschr8    string    //预留字符字段8
	ResChr9    string    //预留字符字段9
	Reschr10   string    //预留字符字段10
	ResNum1    float64   //预留数字字段1
	ResNum2    float64   //预留数字字段2
	ResNum3    float64   //预留数字字段3
	ResNum4    float64   //预留数字字段4
	ResNum5    float64   //预留数字字段5
	ResNum6    float64   //预留数字字段6
	ResNum7    float64   //预留数字字段7
	ResNum8    float64   //预留数字字段8
	ResNum9    float64   //预留数字字段9
	ResNum10   float64   //预留数字字段10
}

type BaseActionEntity struct {
	ErrMsg  string
	ErrCode string
}
