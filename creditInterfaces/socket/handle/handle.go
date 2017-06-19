package handle

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"net"
	"strings"
	// "sync"
	"time"

	"pccqcpa.com.cn/app/rpm/api/creditInterfaces/socket/decoding"
	"pccqcpa.com.cn/app/rpm/api/creditInterfaces/socket/encoding"
	"pccqcpa.com.cn/app/rpm/api/models/biz/ln"
	"pccqcpa.com.cn/app/rpm/api/services/loanService"
	"pccqcpa.com.cn/app/rpm/api/services/pricingService"
	"pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/components/zlog"
)

type Handle struct{}

var custInfoService loanService.CustInfoService
var lnBusinessService loanService.LnBusinessService
var lnMortService loanService.LnMortService
var lnGuaranteService loanService.LnGuaranteService
var sceneDpService loanService.SceneDpService
var sceneItdService loanService.SceneItdService

// 定价服务类
var pricing pricingService.PricingService

// 连接处理
// 去除前面八位长度字节
func (h Handle) HandleConnection(conn net.Conn) {
	// 处理报文
	timeBegain := time.Now()
	length := 1024
	var socketXmlBuffer bytes.Buffer
	for {
		bytes := make([]byte, length)
		n, err := conn.Read(bytes)
		if nil != err {
			zlog.Error("接口读取报文出错", err)
			return
		}
		n, err = socketXmlBuffer.Write(bytes[:n])
		if nil != err {
			zlog.Error("报文写入Buffer出错", err)
			return
		}
		if strings.Contains(socketXmlBuffer.String(), "</XmlStruct>") {
			break
		}
	}

	// 开始解码
	zlog.Infof("读取xml信息：", nil, socketXmlBuffer.String())
	lnbusiness, err := (decoding.Decoding{}).Start(socketXmlBuffer.Bytes())
	if nil != err {
		er := fmt.Errorf("接收xml字节流转换为实体出错，请检查类型")
		zlog.Error(er.Error(), err)
		h.returnErrMsg(er, decoding.XmlHeader{}, timeBegain, conn)
		return
	}
	// fmt.Println(lnbusiness)

	// 必输字段验证
	err = h.checkRequired(*lnbusiness)
	if nil != err {
		h.returnErrMsg(err, lnbusiness.Header, timeBegain, conn)
		return
	}

	// 保存定价单
	// 保存抵质押品
	// 保存保证人
	// 保存存款派生
	// 保存中间派生
	// 保存存量优惠
	err = h.saveLnBusiness(lnbusiness)
	if nil != err {
		h.returnErrMsg(err, lnbusiness.Header, timeBegain, conn)
		return
	}

	// 基础定价
	// 判断是否有派生，如果有派生定价
	// 反算
	fmt.Println("开始定价")
	businessCode := lnbusiness.Body.LnBusiness.BusinessCode
	paramMap := map[string]interface{}{
		"IntRate":    lnbusiness.Body.IntRate,
		"Remark":     lnbusiness.Body.Remark,
		"MarginType": lnbusiness.Body.MarginType,
		"MarginInt":  lnbusiness.Body.MarginInt,
	}
	businessTypes := util.LN_BUSINESS + "," + util.LN_INVERSE
	var pricingRstMsg *ln.LnPricing
	// 判断是否有派生信息，有则进行派生计算，没有就跳过
	if 0 != len(lnbusiness.Body.SceneDpList.LnSceneDps) ||
		0 != len(lnbusiness.Body.SceneItdList.LnSceneItds) ||
		!(0 == lnbusiness.Body.StockUsage) {
		paramMap["StockUsage"] = lnbusiness.Body.StockUsage
		businessTypes += "," + util.LN_SCENE
		pricingRstMsg, err = pricing.LnBusinessPricing(businessCode, businessTypes, paramMap)
		if nil != err {
			fmt.Println("计算存量优惠出错", err)
			h.returnErrMsg(err, lnbusiness.Header, timeBegain, conn)
			return
		}
	} else {
		pricingRstMsg, err = pricing.LnBusinessPricing(businessCode, businessTypes, paramMap)
		if nil != err {
			fmt.Println("计算定价出错", err)
			h.returnErrMsg(err, lnbusiness.Header, timeBegain, conn)
			return
		}
	}

	// 开始编码返回值
	resXml := (encoding.Encoding{}).Start(pricingRstMsg, conn)
	resXml.Header.XmlHeader = lnbusiness.Header
	resXml.Header.SumTime = fmt.Sprint(time.Since(timeBegain))
	resXmlBytes, _ := xml.Marshal(resXml)
	resBytes := append([]byte(xml.Header), resXmlBytes...)
	resBytes = append(resBytes, "\n"...)
	conn.Write(resBytes)
	zlog.Infof("返回信息：%s", nil, string(resBytes))
}

func (h Handle) returnErrMsg(err error, header decoding.XmlHeader, timeBegain time.Time, conn net.Conn) {
	var resXml encoding.ResXml
	resXml.Header.XmlHeader = header
	resXml.Header.SumTime = fmt.Sprint(time.Since(timeBegain))
	resXml.Body.StatusCode = "401"
	resXml.Body.Msg = err.Error()

	bytes, _ := xml.Marshal(resXml)

	rstMsgBytes := append([]byte(xml.Header), bytes...)
	rstMsgBytes = append(rstMsgBytes, "\n"...)

	conn.Write(rstMsgBytes)

}

// 检查必要字段
func (h Handle) checkRequired(lnbusiness decoding.XmlStruct) error {
	return LnPricingFiledCheck{}.CheckPricingFiled(lnbusiness)
}

// 客户信息直接新增或者修改
// 保存定价业务信息
// 保存抵质押品
// 保存保证人
// 保存存款派生
// 保存中间派生
// 保存存量优惠
// func (h Handle) saveLnBusiness(lnbusiness *decoding.XmlStruct) error {
// 	var waitGroup sync.WaitGroup

// 	var errCh chan error
// 	go h.addOrUpdateCustInfo(lnbusiness.Header, &lnbusiness.Body.CustInfo, errCh, waitGroup)
// 	go h.addLnBusiness(lnbusiness.Header, &lnbusiness.Body.LnBusiness, errCh, waitGroup)
// 	go h.addLnMort(lnbusiness.Header, lnbusiness.Body.MortList.LnMorts, errCh, waitGroup)
// 	go h.addGuarantes(lnbusiness.Header, lnbusiness.Body.GuaranteList.LnGuarantes, errCh, waitGroup)
// 	go h.addSceneDps(lnbusiness.Header, lnbusiness.Body.SceneDpList.LnSceneDps, errCh, waitGroup)
// 	go h.addSceneItds(lnbusiness.Header, lnbusiness.Body.SceneItdList.LnSceneItds, errCh, waitGroup)

// 	go func() {
// 		er := <-errCh
// 		if nil != er {
// 			return
// 		}
// 	}()
// 	waitGroup.Wait()
// }

func (h Handle) saveLnBusiness(lnbusiness *decoding.XmlStruct) error {
	var errChLength int = 6
	var errCh = make(chan error, errChLength)
	go h.addOrUpdateCustInfo(lnbusiness.Header, &lnbusiness.Body.CustInfo, errCh)
	go h.addLnBusiness(lnbusiness.Header, &lnbusiness.Body.LnBusiness, errCh)
	go func() {
		if util.MAIN_MORTGAGE_TYPE_CREDIT == lnbusiness.Body.LnBusiness.MainMortgageType {
			errCh <- nil
		} else {
			h.addLnMort(lnbusiness.Header, lnbusiness.Body.MortList.LnMorts, errCh)
		}
	}()
	go h.addGuarantes(lnbusiness.Header, lnbusiness.Body.GuaranteList.LnGuarantes, errCh)
	go h.addSceneDps(lnbusiness.Header, lnbusiness.Body.SceneDpList.LnSceneDps, errCh)
	go h.addSceneItds(lnbusiness.Header, lnbusiness.Body.SceneItdList.LnSceneItds, errCh)

	var errString string = ""
	for i := 0; i < errChLength; i++ {
		zlog.Infof("多协程保存定价业务信息等待返回【%v/%v】", nil, i+1, errChLength)
		er := <-errCh
		if nil != er {
			errString += er.Error() + "\n"
		}
		fmt.Println("----------", er)
	}
	fmt.Println("errString", errString)
	if "" != errString {
		return fmt.Errorf(errString)
	}
	return nil
}

// 保存中间派生
func (h Handle) addSceneItds(header decoding.XmlHeader, sceneItds []ln.SceneItd, er chan error) {
	// fmt.Println("新增派生中间", len(sceneItds))
	err := sceneItdService.DeleteByBusinessCode(header.BusinessCode)
	if nil != err {
		er <- err
		return
	}
	for _, sceneItd := range sceneItds {
		sceneItd.BusinessCode = header.BusinessCode
		sceneItd.CreateUser = header.UserName
		sceneItd.UpdateUser = header.UserName
		err := sceneItdService.Add(&sceneItd)
		if nil != err {
			er <- err
			return
		}
	}
	er <- nil
	return
}

// 保存存款派生
func (h Handle) addSceneDps(header decoding.XmlHeader, sceneDps []ln.SceneDp, er chan error) {
	// fmt.Println("新增派生存款", len(sceneDps))
	err := sceneDpService.DeleteByBusinessCode(header.BusinessCode)
	if nil != err {
		er <- err
		return
	}
	for _, sceneDp := range sceneDps {
		sceneDp.BusinessCode = header.BusinessCode
		sceneDp.CreateUser = header.UserName
		sceneDp.UpdateUser = header.UserName
		err := sceneDpService.Add(&sceneDp)
		if nil != err {
			er <- err
			return
		}
	}
	er <- nil
	return
}

// 保存保证人
func (h Handle) addGuarantes(header decoding.XmlHeader, guarantes []ln.LnGuarante, er chan error) {
	//fmt.Println("-------先删除保证人---------", businessCode)
	err := lnGuaranteService.DeleteByBusinessCode(header.BusinessCode)
	if nil != err {
		er <- err
		return
	}
	//fmt.Println("-------循环添加保证人---------", len(guarantes))
	for _, guarante := range guarantes {
		guarante.BusinessCode = header.BusinessCode
		guarante.CreateUser = header.UserName
		guarante.UpdateUser = header.UserName
		err := lnGuaranteService.Add(&guarante)
		if nil != err {
			er <- err
			return
		}
	}
	er <- nil
	return
}

// 保存抵质押品信息
func (h Handle) addLnMort(header decoding.XmlHeader, morts []ln.LnMort, er chan error) {
	err := lnMortService.DeleteByBusinessCode(header.BusinessCode)
	if nil != err {
		er <- err
		return
	}
	for _, mort := range morts {
		mort.BusinessCode = header.BusinessCode
		mort.CreateUser = header.UserName
		mort.UpdateUser = header.UserName
		err := lnMortService.Add(&mort)
		if nil != err {
			er <- err
			return
		}
	}
	er <- nil
	return
}

// 保存定价业务信息
func (h Handle) addLnBusiness(header decoding.XmlHeader, lnBusiness *ln.LnBusiness, er chan error) {
	err := lnBusinessService.Delete(lnBusiness)
	if nil != err {
		er <- err
		return
	}
	// fmt.Println("删除业务信息成功")
	lnBusiness.BusinessCode = header.BusinessCode
	lnBusiness.CreateUser = header.UserName
	lnBusiness.UpdateUser = header.UserName
	err = lnBusinessService.Add(lnBusiness)
	if nil != err {
		er <- err
		return
	}
	er <- nil
	return
	// fmt.Println("新增业务信息成功")
}

// 新增或者更改客户信息
func (h Handle) addOrUpdateCustInfo(header decoding.XmlHeader, custInfo *ln.CustInfo, er chan error) {
	paramMap := map[string]interface{}{
		"cust_code": custInfo.CustCode,
	}
	custInfoOld, err := custInfoService.FindOne(paramMap)
	if nil != err {
		er <- err
		return
	}
	// fmt.Println("nil != custInfoOld", nil != custInfoOld)
	custInfo.CreateUser = header.UserName
	custInfo.UpdateUser = header.UserName
	if nil != custInfoOld {
		custInfo.UUID = custInfoOld.UUID
		err := custInfoService.Update(custInfo)
		if nil != err {
			er <- err
			return
		}
	} else {
		err = custInfoService.Add(custInfo)
		if nil != err {
			er <- err
			return
		}
	}
	er <- nil
	return
}
