package pricing

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"
	"sync"
	"time"

	// "pccqcpa.com.cn/app/rpm/api/models/biz/ln"
	"pccqcpa.com.cn/app/rpm/api/models/par"
	"pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/components/zlog"
)

type RpmEngine struct {
	functionIndicators            map[string]*par.Indicator // key: method value: indicator
	functionIndicatorValue        map[string]interface{}    // key: indicatorCode value: 函数指标返回的值
	formulaIndicatorFormula       map[string]string         // key: indicatorCode value: 四则运算公式
	formulaIndicatorExchangeValue map[string]string         // key: indicatorCode value: 数字四则运算
	formulaIndicatorLength        int                       // 公式指标长度
	functionIndicatorsLength      int                       // 函数指标长度
	formulaIndicatorCalcRst       map[string]interface{}    // key: indicatorCode value: 四则运算结果
	formulaIndicatorMethod        map[string]*par.Indicator // key: 公式指标method字段 value: indicator
	formulaIndicatorMethodLength  int                       // 公式指标有method的长度
	formulaIndicators             map[string]*par.Indicator // 查询出来的公式指标 key为indicatorCode
}

var indicatorModel par.Indicator

var pricingBus = new(PricingBus)

var pricingFormulaBus = new(PricingFormulaBus)

var formulaMethodMutex sync.RWMutex // 公式指标调用函数赋值锁🔒

// 反射调用函数指标
var errString string
var errStringMap map[string]string

// golang的并发map是不可靠的，必须加锁
var mutex sync.RWMutex

// 开始调用引擎计算各公式指标
// 一、查询公式指标
// 二、查看依赖的函数指标
// 三、根据函数指标的method字段反射调用方法
// 四、将函数指标计算的值放入map中，返回结果
func (r RpmEngine) StartPricing(pricingEntity interface{}, businessTypes string) (map[string]interface{}, error) {
	// fmt.Println("开始调用计算引擎")
	r.functionIndicators = make(map[string]*par.Indicator)
	r.functionIndicatorValue = make(map[string]interface{})
	r.formulaIndicatorFormula = make(map[string]string)
	r.formulaIndicatorExchangeValue = make(map[string]string)
	r.formulaIndicatorCalcRst = make(map[string]interface{})
	r.formulaIndicatorMethod = make(map[string]*par.Indicator)
	r.formulaIndicators = make(map[string]*par.Indicator)
	errString = ""
	errStringMap = make(map[string]string)
	var businessType string = ""

	for _, v := range strings.Split(businessTypes, ",") {
		businessType += "'" + v + "',"
	}
	businessType = strings.TrimSuffix(businessType, ",")

	params := map[string]interface{}{
		"indicator_type": util.FIRST_INDICATOR,
		// "indicator_business_type": businessType,
		// "flag": util.FLAG_TRUE,
		"searchLike": []map[string]interface{}{
			map[string]interface{}{
				"type":  "in",
				"key":   "indicator_business_type",
				"value": businessType,
			},
		},
	}

	formulaIndicators, err := indicatorModel.Find(params)
	if nil != err {
		zlog.Error("查询应用情景为[%s]的公式指标错误", err)
		return nil, err
	}

	r.getFunctionIndicator(formulaIndicators)

	// 开协程计算有method的公式指标
	var fromulaMethodCh = make(chan int, r.formulaIndicatorMethodLength)
	pricingFormulaBusParams := make([]reflect.Value, 1)
	pricingFormulaBusParams[0] = reflect.ValueOf(pricingEntity)
	for k, v := range r.formulaIndicatorMethod {
		go r.reflectToFormulaMethod(k, v.IndicatorCode, pricingFormulaBusParams, fromulaMethodCh)
	}

	// 开始调用函数指标
	pricingBusParams := make([]reflect.Value, 2)
	pricingBusParams[0] = reflect.ValueOf(pricingEntity)
	pricingBusParams[1] = reflect.ValueOf(businessType)

	var ch = make(chan int, r.formulaIndicatorLength)
	for k, v := range r.functionIndicators {
		go r.reflectToFunction(k, v.IndicatorCode, pricingBusParams, ch)
	}

	r.wait(r.formulaIndicatorMethodLength, fromulaMethodCh)
	// fmt.Println("－－－－－－－－－－－开始等待协程")
	// time.Sleep(time.Second * 10)

	r.wait(r.functionIndicatorsLength, ch)

	// fmt.Println("----", pricingEntity)
	if "" != errString {
		er := fmt.Errorf(errString)
		zlog.Error(er.Error(), er)
		return nil, er
	}

	// 开始计算公式指标值
	// fmt.Println("－－－－－－－－－－－－函数指标返回值集合\n", r.functionIndicatorValue, errString)
	// time.Sleep(time.Second * 20)

	// 超时处理，判断是否进入死循环
	var timeout chan int = make(chan int, 1)
	go func() {
		time.Sleep(time.Second * 5)
		timeout <- 1
	}()
	r.formulaIndicatorExchange(r.formulaIndicatorFormula, r.functionIndicatorValue, timeout)

	ch = make(chan int, r.formulaIndicatorLength)
	// fmt.Println(r.formulaIndicatorExchangeValue)
	// time.Sleep(time.Second * 10)
	for k, v := range r.formulaIndicatorExchangeValue {
		go r.calculateFirstIndicator(k, v, ch)
	}
	r.wait(r.formulaIndicatorLength, ch)
	if "" != errString {
		er := fmt.Errorf(errString)
		zlog.Error(er.Error(), er)
		return r.formulaIndicatorCalcRst, er
	}
	return r.formulaIndicatorCalcRst, nil
}

// 公式指标计算
func (r *RpmEngine) calculateFirstIndicator(key, formulae string, ch chan int) {
	value, err := util.Calculate(formulae)
	// fmt.Printf("\n----------计算公式指标[%s]\n%s\n值：[%v,%v]\n", key, formulae, value, err)
	mutex.Lock()
	if nil != err {
		var msg string = "公式[" + r.formulaIndicators[key].IndicatorName + "]计算错误"
		zlog.Error(msg, err)
		r.checkErrString(msg)
		r.formulaIndicatorCalcRst[key] = -1
	} else {
		r.formulaIndicatorCalcRst[key] = value
	}
	mutex.Unlock()
	ch <- 0
}

// GetFunctionIndicator RpmEngine 获取函数指标map
func (r *RpmEngine) getFunctionIndicator(formulaIndicators []*par.Indicator) {
	for _, formulaIndicator := range formulaIndicators {
		// 赋值查询出来的公式指标
		r.formulaIndicators[formulaIndicator.IndicatorCode] = formulaIndicator

		if util.FIRST_INDICATOR == formulaIndicator.IndicatorType &&
			"" != strings.Replace(formulaIndicator.Method, " ", "", -1) {
			if _, ok := r.formulaIndicatorMethod[formulaIndicator.Method]; !ok {
				r.formulaIndicatorMethod[formulaIndicator.Method] = formulaIndicator
				r.formulaIndicatorMethodLength++
			}
			continue
		}

		if util.FIRST_INDICATOR == formulaIndicator.IndicatorType {
			if _, ok := r.formulaIndicatorFormula[formulaIndicator.IndicatorCode]; !ok {
				mutex.Lock()
				r.formulaIndicatorFormula[formulaIndicator.IndicatorCode] = strings.Replace(formulaIndicator.FormulaeCode, " ", "", -1)
				mutex.Unlock()
				r.formulaIndicatorLength++
				r.getFunctionIndicator(formulaIndicator.IndicatorBeforeRely)
			}
		}
		// fmt.Println("函数指标方法：", formulaIndicator.Method)
		if _, ok := r.functionIndicators[formulaIndicator.Method]; !ok && "" != strings.Replace(formulaIndicator.Method, " ", "", -1) {
			mutex.Lock()
			r.functionIndicators[formulaIndicator.Method] = formulaIndicator
			mutex.Unlock()
			r.functionIndicatorsLength++
		}
	}
}

// 递归转换公式
// 先将公式指标中依赖的函数指标换成数字，如果全部转换
func (r *RpmEngine) formulaIndicatorExchange(formulaMap map[string]string, functionIndicatorValue interface{}, timeout chan int) (err error) {
	// 超时处理
	select {
	case <-timeout:
		err = fmt.Errorf("公式指标转四则运算超时")
		return err
	default:

	}
	//循环第一层指标，将依赖指标换成数值，传给公式函数计算
	var length int
	var formulaIndicatorFormula = make(map[string]string)
	for key, formula := range formulaMap {
		exchangeEndFormulae := util.FormulaExchange(formula, functionIndicatorValue)
		var reg = regexp.MustCompile("[a-z_A-Z]+")
		if reg.MatchString(exchangeEndFormulae) {
			// 依赖了公式指标
			formulaMap[key] = exchangeEndFormulae
			// fmt.Println("循环获取指标值", exchangeEndFormulae)
			length++
		} else {
			delete(formulaMap, key)
			r.formulaIndicatorExchangeValue[key] = exchangeEndFormulae
			formulaIndicatorFormula[key] = "(" + exchangeEndFormulae + ")"
		}

		// fmt.Println("－－－－－－－－依赖指标值－－－－－－－", functionIndicatorValue)
		// fmt.Printf("\n\n\n\n\n－－－－－公式指标[%s]\n转换之前:%s\n转换之后:%s\n", key, formula, exchangeEndFormulae)
		zlog.Infof("\n\n\n\n\n－－－－－公式指标[%s]\n转换之前:%s\n转换之后:%s\n", nil, key, formula, exchangeEndFormulae)
		// time.Sleep(time.Second * 5)
	}
	if 0 < length {
		err := r.formulaIndicatorExchange(formulaMap, formulaIndicatorFormula, timeout)
		return err
	}
	return nil
}

func (r *RpmEngine) reflectToFunction(methodName, indicatorCode string, pricingBusParams []reflect.Value, ch chan int) {

	// fmt.Println("--＋＋＋＋-", reflect.ValueOf(pricingBus).MethodByName(methodName).IsValid(), methodName)
	if reflect.ValueOf(pricingBus).MethodByName(methodName).IsValid() {
		// fmt.Println("反射参数", pricingBusParams[0].Elem().Interface())
		// fmt.Println("反射方法", methodName)
		values := reflect.ValueOf(pricingBus).MethodByName(methodName).Call(pricingBusParams)
		if !values[1].IsNil() {
			// fmt.Println("=======~~~~~~~~~~~~~~~~~=====", values[1].Interface().(error).Error())
			r.checkErrString(values[1].Interface().(error).Error())

			ch <- 0
			return
		}
		mutex.Lock()
		// fmt.Printf("函数指标[%v]的值[%v]\n", indicatorCode, values[0].Interface())
		r.functionIndicatorValue[indicatorCode] = values[0].Interface()
		mutex.Unlock()
		ch <- 0
	} else {
		r.checkErrString("函数指标反射方法[" + methodName + "]出错")
		ch <- 0
	}

}

// 公式指标反射到方法
func (r *RpmEngine) reflectToFormulaMethod(methodName, indicatorCode string, pricingFormulaBusParams []reflect.Value, ch chan int) {
	// fmt.Println("\n\n\n\n\n公式指标反射是否可以反射\n\n\n\n\n", methodName, indicatorCode, pricingFormulaBusParams, reflect.ValueOf(pricingFormulaBus).MethodByName(methodName).IsValid())
	methodName = strings.Replace(methodName, " ", "", -1)
	if reflect.ValueOf(pricingFormulaBus).MethodByName(methodName).IsValid() {
		values := reflect.ValueOf(pricingFormulaBus).MethodByName(methodName).Call(pricingFormulaBusParams)
		// fmt.Println("\n\n\n\n\n...\n\n\n\n\n")
		if !values[1].IsNil() {
			r.checkErrString(values[1].Interface().(error).Error())
			ch <- 0
			return
		}
		mutex.Lock()
		// fmt.Printf("公式指标[%v]的值[%v]\n", indicatorCode, values[0].Interface())
		r.functionIndicatorValue[indicatorCode] = values[0].Interface()
		r.formulaIndicatorCalcRst[indicatorCode] = values[0].Interface()
		mutex.Unlock()
		ch <- 0
	} else {
		r.checkErrString("公式指标反射方法[" + methodName + "]出错")
		ch <- 0

	}

}

// 通过反射获取基础指标的值
func (r *RpmEngine) getValueByReflect(name string, business interface{}, ch chan int) {
	// 判断是否包含.
	var code string
	if strings.Contains(name, ".") {
		code = util.SubstrStartEnd(name, 0, strings.Index(name, "."))
	} else {
		code = name
	}
	// fmt.Println("code:", code)
	// 开始反射获取值
	s := reflect.ValueOf(business)
	if reflect.Struct != reflect.TypeOf(business).Kind() {
		s = s.Elem()
	}
	if s.CanInterface() { // 是否可以取值
		if s.FieldByName(code).IsValid() { // 判断code属性是否存在
			if s.FieldByName(code).CanInterface() { // 判断是否可以取值
				value := s.FieldByName(code).Interface()
				// 如果是结构体或者地址，则继续反射
				// fmt.Println("...", reflect.TypeOf(value).Kind(), reflect.Struct)
				if -1 != strings.Index(name, ".") {
					// if reflect.TypeOf(value).Kind() != reflect.Ptr {
					// 	err := fmt.Errorf("反射值是不可寻址，请检查基础指标code")
					// 	zlog.Error(err.Error(), err)
					// 	return err
					// }
					subName := util.SubstrStartEnd(name, strings.Index(name, ".")+1, len(name))
					// fmt.Println("subName", subName)
					r.getValueByReflect(subName, value, ch)
				}
				// fmt.Println("----反射取值---", code, fmt.Sprint(value))
				r.functionIndicatorValue[code] = value
				ch <- 0
				return
			} else {
				err := fmt.Errorf("获取指标[%v]值出错\n", code)
				zlog.Error("直接反射"+err.Error(), err)
				r.checkErrString(err.Error())
				ch <- -1
				return
			}
		} else {
			err := fmt.Errorf("指标[%s]对应属性不存在\n", code)
			zlog.Error("直接反射"+err.Error(), err)
			r.checkErrString(err.Error())
			ch <- -1
			return
		}
	} else {
		err := fmt.Errorf("获取基础指标反射取值出错")
		zlog.Error(err.Error(), err)
		ch <- -1
		return
	}
}

// 线程间通讯
func (r RpmEngine) wait(length int, ch chan int) {
	for i := 0; i < length; i++ {
		// fmt.Printf("－－－－－－等待线程返回结果进度[%v/%v]－－－－－－\n", i+1, length)
		// fmt.Printf("函数指标[%v]\n", r.functionIndicators)
		// fmt.Printf("函数指标返回值[%v]\n", r.functionIndicatorValue)
		<-ch
	}
}

func (r RpmEngine) checkErrString(str string) {
	if _, ok := errStringMap[str]; !ok {
		errString += str + "\n"
		errStringMap[str] = ""
	}
}
