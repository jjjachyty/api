package decoding

import (
	"encoding/xml"
	"fmt"

	"pccqcpa.com.cn/app/rpm/api/models/biz/ln"
	"pccqcpa.com.cn/components/zlog"
)

type Decoding struct{}

type XmlStruct struct {
	Header XmlHeader `xml:"header"`
	Body   XmlBody   `xml:"body"`
}

type XmlHeader struct {
	BusinessCode string
	UserCode     string
	UserName     string
}

type XmlBody struct {
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
	Remark       string      // 备注
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

func (Decoding) Start(data []byte) (*XmlStruct, error) {
	zlog.Info("开始解码", nil)
	var xmlStruct = new(XmlStruct)
	err := xml.Unmarshal(data, &xmlStruct)
	if nil != err {
		er := fmt.Errorf("socket接口解码XML出错")
		zlog.Error(er.Error(), err)
		return nil, er
	}
	zlog.Info("socket服务端解码XML成功", nil)
	// zlog.Infof("解码后实体：%#v", nil, xmlStruct)
	return xmlStruct, nil
}
