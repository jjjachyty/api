package main

import (
	//"encoding/json"
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"pccqcpa.com.cn/app/rpm/api/models/biz/ln"
	// "strconv"
	"time"
)

type XmlStruct struct {
	Header XmlHeader `xml:"header"`
	Body   LnPricing `xml:"body"`
}

type XmlHeader struct {
	UserCode string
	UserName string
}

type LnPricing struct {
	CustInfo     ln.CustInfo
	LnBusiness   ln.LnBusiness
	MortList     LnMorts     `xml:"mortArry"`
	GuaranteList LnGuarantes `xml:"guaranteArry"`
	SceneDpList  LnSceneDps  `xml:"sceneDpArry"`
	SceneItdList LnSceneItds `xml:"sceneItdArry"`
	StockUsage   float64     // 存量优惠
	IntRate      float64     // 执行利率
	MarginType   string      // 浮动类型
	MarginInt    float64     // 浮动值
	Remark       string      //备注
}

type LnMorts struct {
	LnMorts []ln.LnMort `xml:"mort"`
}

type LnGuarantes struct {
	LnGuarantes []ln.LnGuarante `xml:"guarante"`
}

type LnSceneDps struct {
	LnSceneDps []ln.SceneDp `xml:"sceneDp"`
}

type LnSceneItds struct {
	LnSceneItds []ln.SceneItd `xml:"sceneItd"`
}

func sender(conn net.Conn) {
	resultSend := []byte(xml.Header)
	file, _ := os.Open("./socket_test.xml")
	bytes, _ := ioutil.ReadAll(file)
	file.Close()
	resultSend = append(resultSend, bytes...)
	resultSend = append(resultSend, "\n"...)
	ioutil.WriteFile("./socket_test1.xml", resultSend, 0666)

	conn.Write([]byte(resultSend))
	fmt.Println("send over")
}

func sender1(conn net.Conn) {
	var a = new(LnPricing)
	a.CustInfo.CustCode = "1111111111111"
	a.CustInfo.CustName = "信贷接口测试1"
	a.CustInfo.Organization = "0"
	a.CustInfo.CustType = "2"
	a.CustInfo.CustImplvl = "1"
	a.CustInfo.CustCredit = "A"
	a.CustInfo.Branch.OrganCode = "panchina001"
	a.CustInfo.Industry.IndustryCode = "A"
	a.CustInfo.CustSize = "2"
	a.CustInfo.StockContribute = 6000000.000
	a.CustInfo.Status = "01"
	a.CustInfo.Flag = "1"
	a.CustInfo.Owner = "rpm"

	a.IntRate = 0.065
	a.MarginType = "1"
	a.MarginInt = 20
	a.Remark = "Janlysocket接口测试"

	a.LnBusiness.BusinessCode = "jjjjjjjjjjjjjjjjjjjjj"
	a.LnBusiness.Cust.CustCode = a.CustInfo.CustCode
	a.LnBusiness.Organ.OrganCode = "panchina001"
	a.LnBusiness.Product.ProductCode = "19011001"
	a.LnBusiness.Currency = "CNY"
	a.LnBusiness.Term = 2
	a.LnBusiness.TermMult = "Y"
	a.LnBusiness.RateType = "1"
	a.LnBusiness.RpymType = "1"
	a.LnBusiness.RepriceFreq = 0
	a.LnBusiness.RpymInterestFreq = 0
	a.LnBusiness.RpymCapitalFreq = 0
	a.LnBusiness.Principal = 1400000
	a.LnBusiness.BaseRateType = "1"
	a.LnBusiness.MainMortgageType = "1"
	a.LnBusiness.Flag = "1"
	a.LnBusiness.CreateUser = "RPM系统管理员"
	//抵质押
	var mort1 ln.LnMort
	var mort2 ln.LnMort

	mort1.MortgageCode = "004007001003"
	mort1.MortgageName = "其他押品-无形资产-可转让知识产权中的财产权-著作权"
	mort1.MortgageValue = 100000
	mort1.Currency = "CNY"
	//mort1.BusinessCode=a.LnBusiness.BusinessCode

	mort2.MortgageCode = "003004002001"
	mort2.MortgageName = "居住用房产类在建工程"
	mort2.MortgageValue = 200000
	mort2.Currency = "CNY"
	//mort2.BusinessCode=a.LnBusiness.BusinessCode

	a.MortList.LnMorts = append(a.MortList.LnMorts, mort1)
	a.MortList.LnMorts = append(a.MortList.LnMorts, mort2)

	//担保人
	var guarte1 ln.LnGuarante
	var guarte2 ln.LnGuarante

	guarte1.Guarante.CustCode = "91500000709426199C"
	guarte1.Guarante.CustName = "重庆长安民生物流股份有限公司"
	guarte1.GuaranteType = "1"

	guarte2.Guarante.CustCode = "100000000040049"
	guarte2.Guarante.CustName = "重庆长安汽车股份有限公司"
	guarte2.GuaranteType = "2"

	a.GuaranteList.LnGuarantes = append(a.GuaranteList.LnGuarantes, guarte1)
	a.GuaranteList.LnGuarantes = append(a.GuaranteList.LnGuarantes, guarte2)

	//派生存款

	var sceneDp1 ln.SceneDp
	var sceneDp2 ln.SceneDp

	sceneDp1.Product.ProductCode = "20010205"
	sceneDp1.Product.ProductName = "其它单位定期存款"
	sceneDp1.Currency = "CNY"
	sceneDp1.Term = 360
	sceneDp1.Rate = 0.0195
	sceneDp1.Value = 200000

	sceneDp2.Product.ProductCode = "20010101"
	sceneDp2.Product.ProductName = "单位一般活期存款"
	sceneDp2.Currency = "CNY"
	sceneDp2.Term = 0
	sceneDp2.Rate = 0.00385
	sceneDp2.Value = 200000

	a.SceneDpList.LnSceneDps = append(a.SceneDpList.LnSceneDps, sceneDp1)
	a.SceneDpList.LnSceneDps = append(a.SceneDpList.LnSceneDps, sceneDp2)

	//派生中间
	var sceneItd1 ln.SceneItd
	var sceneItd2 ln.SceneItd

	sceneItd1.Product.ProductCode = "300108"
	sceneItd1.Product.ProductName = "银证转账"
	sceneItd1.Value = 100000

	sceneItd2.Product.ProductCode = "300503"
	sceneItd2.Product.ProductName = "理财咨询服务"
	sceneItd2.Value = 100000

	a.SceneItdList.LnSceneItds = append(a.SceneItdList.LnSceneItds, sceneItd1)
	a.SceneItdList.LnSceneItds = append(a.SceneItdList.LnSceneItds, sceneItd2)

	// jsonMsg, _ := json.MarshalIndent(a, "", "")
	// fmt.Println(string(jsonMsg))

	var xmlStruct XmlStruct
	xmlStruct.Body = *a

	result, _ := xml.MarshalIndent(xmlStruct, "", "")
	resultSend := []byte(xml.Header)
	resultSend = append(resultSend, result...)
	// 处理前八位长度
	// length := len(resultSend)
	// var lengthTemp string = strconv.Itoa(length)
	// for i := len(lengthTemp); i < 8; i++ {
	// 	lengthTemp = "0" + lengthTemp
	// }
	// fmt.Println(lengthTemp)
	// resultSend = append([]byte(lengthTemp), resultSend...)

	ioutil.WriteFile("./socket.xml", resultSend, 0666)
	//fmt.Println("发送报文：-----\n",string(resultSend))

	// words := "00000012hello world!"
	resultSend = append(resultSend, "\n"...)
	conn.Write([]byte(resultSend))
	fmt.Println("send over")

}

func main() {
	timeBegain := time.Now()

	server := "172.168.171.10:9092"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", server)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}

	fmt.Println("connect success")
	sender(conn)

	// 获取返回信息
	var buf bytes.Buffer
	for {
		var rstMsg = make([]byte, 1024)
		n, _ := conn.Read(rstMsg)
		buf.Write(rstMsg[:n])
		if string(rstMsg[n-1]) == "\n" {
			break
		}
	}
	fmt.Printf("返回信息：\n%v", buf.String())
	timeEnd := time.Now()
	fmt.Printf("调用接口定价花销时间【%v】\n", timeEnd.Sub(timeBegain))
}
