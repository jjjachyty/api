package encoding

import (
	// "encoding/xml"
	"fmt"
	"net"
	// "time"

	"pccqcpa.com.cn/app/rpm/api/creditInterfaces/socket/decoding"
	"pccqcpa.com.cn/app/rpm/api/models/biz/ln"
	// "pccqcpa.com.cn/components/zlog"
)

type Encoding struct{}

type ResXml struct {
	Header ResXmlHeader `xml:"header"`
	Body   ResXmlBody   `xml:"body"`
}

type ResXmlHeader struct {
	decoding.XmlHeader
	SumTime string
}

type ResXmlBody struct {
	StatusCode string
	Msg        string
	BottomRate float64 `json:"omitempty"`
	SceneRate  float64 `json:"omitempty"`
	TgtRate    float64 `json:"omitempty"`
}

func (Encoding) Start(lnPricing *ln.LnPricing, conn net.Conn) ResXml {
	fmt.Println("开始编码")
	var resXml ResXml
	resXml.Body.BottomRate = lnPricing.BottomRate
	resXml.Body.SceneRate = lnPricing.SceneRate
	resXml.Body.TgtRate = lnPricing.TgtRate

	resXml.Body.StatusCode = "200"
	resXml.Body.Msg = "计算成功"
	return resXml
}
