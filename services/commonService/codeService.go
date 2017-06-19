package commonService

import (
	"strconv"
	"sync"
	"time"
	// "pccqcpa.com.cn/app/rpm/api/models/biz/ln"
)

type CodeService struct {
}

var lnBusinessCodeMutx = sync.Mutex{}
var lnBusinessNo int = 1001

// 通过时间戳加3位序号生成订单号
func (CodeService) GetLnBusinessCode() string {
	date := time.Now().Format("20060102150405")
	lnBusinessCodeMutx.Lock()
	defer lnBusinessCodeMutx.Unlock()
	str := strconv.Itoa(lnBusinessNo)
	lnBusinessNo++
	lnBusienssCode := "DK" + date + string([]rune(str)[len(str)-3:len(str)])

	// 保存定价单号
	// var lnbusiness ln.LnBusiness
	// lnbusiness.BusinessCode = lnBusienssCode
	// lnbusiness.Add()
	return lnBusienssCode
}

var custNoMutx = sync.Mutex{}
var custCodeByTimeMutex = sync.Mutex{}
var custNo int = 1001
var length int = 3

// 通过时间戳加3位序号数生成客户号
func (CodeService) GetCustCodeByTime() string {
	date := time.Now().Format("20060102150405")
	custCodeByTimeMutex.Lock()
	defer custCodeByTimeMutex.Unlock()
	str := strconv.Itoa(custNo)
	custNo++
	if len(str) > length+1 {
		length = 1 + length
	}
	custCode := "CUST" + date + string([]rune(str)[len(str)-length:len(str)])
	return custCode
}

var oneDpBusinessCodeMutex = sync.Mutex{}
var oneDpBusinessCodeNo = 1001

// 通过时间戳加3位序号数生成一对一存款业务单号
func (CodeService) GetOneDpBusinessByTime() string {
	date := time.Now().Format("20060102150405")
	oneDpBusinessCodeMutex.Lock()
	defer oneDpBusinessCodeMutex.Unlock()
	str := strconv.Itoa(oneDpBusinessCodeNo)
	oneDpBusinessCodeNo++
	if len(str) > length+1 {
		length = 1 + length
	}
	oneDpBusinessCodeCode := "DP" + date + string([]rune(str)[len(str)-length:len(str)])
	return oneDpBusinessCodeCode
}

// 生成新客户号
// func (CodeService) GetCustCodeBySeq() (string, error) {
// 	date := time.Now().Format("20060102")
// 	custNoMutx.Lock()
// 	defer custNoMutx.Unlock()

// 	sql := `
// 		select rpm_custno_seq.nextval from dual
// 	`
// 	var seq string
// 	rows, err := dbobj.Default.Query(sql)
// 	if nil != err {
// 		var msg string = "获取客户号出错，请检查数据库是否存在rpm_custno_seq序列"
// 		zlog.Error(msg, err)
// 		return "", err
// 	}
// 	defer rows.Close()
// 	for rows.Next() {
// 		rows.Scan(&seq)
// 	}
// 	for i := len(seq); i < 4; i++ {
// 		seq = "0" + seq
// 	}
// 	return date + seq, nil
// }
