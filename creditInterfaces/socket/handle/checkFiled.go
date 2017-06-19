package handle

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"strings"

	"pccqcpa.com.cn/app/rpm/api/creditInterfaces/socket/decoding"
	"pccqcpa.com.cn/components/zlog"
)

type checkAraay struct {
	StringReauired []string
	// numberRequired []string
}

type LnPricingFiledCheck struct {
	Header     checkAraay
	CustInfo   checkAraay
	LnBusiness checkAraay
	Mort       checkAraay
	Guarante   checkAraay
	SceneDp    checkAraay
	SceneItd   checkAraay
	LnPricing  checkAraay
}

func (l LnPricingFiledCheck) CheckPricingFiled(xmlStruct decoding.XmlStruct) error {

	// 检查Header
	err := checkAraay{}.checkField(lnPricingFiedCheck.Header.StringReauired, &xmlStruct.Header)
	if nil != err {
		er := fmt.Errorf("Header:%s", err.Error())
		return er
	}
	err = checkAraay{}.checkField(lnPricingFiedCheck.CustInfo.StringReauired, &xmlStruct.Body.CustInfo)
	if nil != err {
		er := fmt.Errorf("CustInfo:%s", err.Error())
		return er
	}
	err = checkAraay{}.checkField(lnPricingFiedCheck.LnBusiness.StringReauired, &xmlStruct.Body.LnBusiness)
	if nil != err {
		er := fmt.Errorf("Body:%s", err.Error())
		return er
	}
	err = checkAraay{}.checkField(lnPricingFiedCheck.LnPricing.StringReauired, &xmlStruct.Body)
	if nil != err {
		er := fmt.Errorf("LnPricing:%s", err.Error())
		return er
	}
	for i, mort := range xmlStruct.Body.MortList.LnMorts {
		err = checkAraay{}.checkField(lnPricingFiedCheck.Mort.StringReauired, &mort)
		if nil != err {
			er := fmt.Errorf("Mort[%v]:%s", i, err.Error())
			return er
		}
	}
	for i, guarante := range xmlStruct.Body.GuaranteList.LnGuarantes {
		err = checkAraay{}.checkField(lnPricingFiedCheck.Guarante.StringReauired, &guarante)
		if nil != err {
			er := fmt.Errorf("Guarante[%v]:%s", i, err.Error())
			return er
		}
	}
	for i, sceneDp := range xmlStruct.Body.SceneDpList.LnSceneDps {
		err = checkAraay{}.checkField(lnPricingFiedCheck.SceneDp.StringReauired, &sceneDp)
		if nil != err {
			er := fmt.Errorf("SceneDp[%v]:%s", i, err.Error())
			return er
		}
	}
	for i, sceneItd := range xmlStruct.Body.SceneItdList.LnSceneItds {
		err = checkAraay{}.checkField(lnPricingFiedCheck.SceneItd.StringReauired, &sceneItd)
		if nil != err {
			er := fmt.Errorf("SceneItd[%v]:%s", i, err.Error())
			return er
		}
	}

	return nil
}

func (c checkAraay) checkField(requiredStrs []string, checkedStruct interface{}) error {
	obj := reflect.ValueOf(checkedStruct)
	for _, requiredstr := range requiredStrs {
		objElem := obj.Elem()
		strs := strings.Split(requiredstr, ".")
		var refVal reflect.Value

		for i, str := range strs {
			// 反射判断该字段类型
			// 如果是结构体，继续
			// 如果是数字，判断是否为0
			// 如果是字符串，去空格后判断是否为空字符串
			if !objElem.FieldByName(str).IsValid() {
				er := fmt.Errorf("%s 结构体没有【%s】字段", obj.Type(), strings.SplitAfterN(requiredstr, ".", i+1)[0])
				zlog.Error(er.Error(), er)
				return er
			}
			refVal = objElem.FieldByName(str)
			filedVal := refVal.Interface()
			var flag bool
			switch refVal.Kind() {
			case reflect.String:
				flag = c.checkStringIsNull(filedVal.(string))
			case reflect.Int32, reflect.Int64, reflect.Float32, reflect.Float64, reflect.Int:
				flag = c.checkNumberIsZero(filedVal)
			case reflect.Struct:
				objElem = refVal
			default:
				continue
			}
			if flag {
				er := fmt.Errorf("字段【%s】为必要字段，不能为空或者0", requiredstr)
				zlog.Error(er.Error(), er)
				return er
			}
		}
	}
	return nil
}

// 判断字符串是否为空
// 为空返回 true
// 不为空返回 false
func (checkAraay) checkStringIsNull(str string) bool {
	if "" == strings.Replace(str, " ", "", -1) {
		return true
	}
	return false
}

// 判断数字是否为0
// 为0返回 true
// 不为0返回 false
// golang 整型和浮点类型是不能直接比较的
func (checkAraay) checkNumberIsZero(num interface{}) bool {
	switch reflect.TypeOf(num).Kind() {
	case reflect.Float32, reflect.Float64:
		if 0.0 == num {
			return true
		}
	default:
		if 0 == num {
			return true
		}
	}
	return false
}

var lnPricingFiedCheck LnPricingFiledCheck

// var lnPricingFiedCheck = LnPricingFiledCheck{
// 	Header: checkAraay{
// 		StringReauired: []string{
// 			"BusinessCode",
// 			"UserCode",
// 			"UserName",
// 		},
// 	},
// 	CustInfo: checkAraay{
// 		StringReauired: []string{
// 			"CustCode",
// 			"CustName",
// 			"CustType",
// 			"CustImplvl",
// 			"CustCredit",
// 			"Branch.OrganCode",
// 			"Industry.IndustryCode",
// 			"CustSize",
// 			"Status",
// 			"Owner",
// 		},
// 	},
// 	LnBusiness: checkAraay{
// 		StringReauired: []string{
// 			"Cust.CustCode",
// 			"Organ.OrganCode",
// 			"Product.ProductCode",
// 			"Currency",
// 			"Term",
// 			"TermMult",
// 			"RateType",
// 			"RpymType",
// 			"RepriceFreq",
// 			"RpymCapitalFreq",
// 			"Principal",
// 			"BaseRateType",
// 			"MainMortgageType",
// 		},
// 	},
// 	Mort: checkAraay{
// 		StringReauired: []string{
// 			"MortgageCode",
// 			"MortgageName",
// 			"Currency",
// 			"MortgageValue",
// 		},
// 	},
// 	Guarante: checkAraay{
// 		StringReauired: []string{
// 			// "Cust",
// 			"Guarante.CustCode",
// 			"GuaranteAmout",
// 			"GuaranteType",
// 		},
// 	},
// 	SceneDp: checkAraay{
// 		StringReauired: []string{
// 			"Product.ProductCode",
// 			"Currency",
// 			"Term",
// 			"Rate",
// 			"Value",
// 		},
// 	},
// 	SceneItd: checkAraay{
// 		StringReauired: []string{
// 			"Product.ProductCode",
// 			"Value",
// 		},
// 	},
// 	LnPricing: checkAraay{
// 		StringReauired: []string{
// 			"IntRate",
// 			"MarginType",
// 			"MarginInt",
// 		},
// 	},
// }

func init() {
	sysConf := os.Getenv("GOSYSCONFIG")
	file, err := os.Open(sysConf + "/socket/requiredField.xml")
	if nil != err {
		er := fmt.Errorf("打开socket必要字段配置文件出错", err)
		zlog.Error(er.Error(), err)
	}
	bytes, _ := ioutil.ReadAll(file)
	err = xml.Unmarshal(bytes, &lnPricingFiedCheck)
	fmt.Printf("%#v", lnPricingFiedCheck)
	if nil != err {
		zlog.Error(err.Error(), err)
	}
}
