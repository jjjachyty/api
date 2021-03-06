package util

import (
	"bytes"
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"pccqcpa.com.cn/components/zlog"
)

//FormulaCalculate func 公式计算器
//formula 公式表达式 valus 表达式中指标项的各个值
//	FormulaCalculate("(ftp+oc+ca+ec+el)/(1-tax)", map[string]string{"ftp": "0.257851", "oc": "0.987456", "ca": "0.364476", "ec": "12.134", "el": "1.21466", "tax": "0.2500"})
// func FormulaCalculate(formula string, valus map[string]string) {
// 	zlog.Debugf("gfc 将转换公式%s", nil, formula)
// 	formulaSlice := getExprSlice(formula)
// 	fmt.Println(formulaSlice)
// 	for i := 0; i < len(formulaSlice); i++ {
// 		fmt.Println(formulaSlice[i])
// 		val := valus[formulaSlice[i]]
// 		if "" != val {
// 			formulaSlice[i] = val
// 		}

// 	}

// 	zlog.Debugf("gfc 公式转换结果%s", nil, formula)
// 	fmt.Println(strings.Join(formulaSlice, ""))
// 	Calculate(strings.Join(formulaSlice, ""))
// }

// 将公式中的字母转换成为数字
// formula 公式表达式 (ftp+oc+ca+ec+el)/(1-tax)
// relyFormula 依赖公式 map[string]string{"ftp": "0.257851", "oc": "0.987456", "ca": "0.364476", "ec": "12.134", "el": "1.21466", "tax": "0.2500"}
func FormulaExchange(formula string, relyFormula interface{}) string {
	var one, rst string = "", ""
	reg := regexp.MustCompile(`[\+\-\*\/\(\)]`)
	for _, v := range formula {
		ok := reg.MatchString(string(v))
		if ok {
			rst += getExchageValue(one, relyFormula) + string(v)
			one = ""
			continue
		}
		one += string(v)
	}
	if "" != one {
		rst += getExchageValue(one, relyFormula)
	}
	return rst
}

// 获取转换的值
func getExchageValue(key string, relyFormula interface{}) string {
	switch reflect.TypeOf(relyFormula).String() {
	case "map[string]string":
		if v, ok := relyFormula.(map[string]string)[key]; ok {
			return fmt.Sprint(v)
		}
	case "map[string]float64":
		if v, ok := relyFormula.(map[string]float64)[key]; ok {
			return fmt.Sprintf("%f", v)
		}
	case "map[string]interface {}":
		if v, ok := relyFormula.(map[string]interface{})[key]; ok {
			var typeOfV = reflect.TypeOf(v)
			switch typeOfV.Kind() {
			case reflect.Float64:
				return fmt.Sprintf("%f", v)
			}
			return fmt.Sprint(v)
		}
	default:
		zlog.Errorf("relyFormula参数目前不支持[%v]", nil, reflect.TypeOf(relyFormula).String())
	}
	return key
}

type stackNode struct {
	Data interface{}
	next *stackNode
}

type linkStack struct {
	top   *stackNode
	Count int
}

func (link *linkStack) init() {
	link.top = nil
	link.Count = 0
}

func (link *linkStack) push(data interface{}) {
	node := new(stackNode)
	node.Data = data
	node.next = link.top
	link.top = node
	link.Count++
}

func (link *linkStack) pop() interface{} {
	if link.top == nil {
		return nil
	}
	returnData := link.top.Data
	link.top = link.top.next
	link.Count--
	return returnData
}

//lookTop func Look up the top element in the stack, but not pop.
func (link *linkStack) lookTop() interface{} {
	if link.top == nil {
		return nil
	}
	return link.top.Data
}

//Calculate func 计算方法 -1*(-9--8/-4)/(1-9)*8--8
func Calculate(expr string) (float64, error) {
	zlog.Debugf("gfc 将解析表达式%s", nil, expr)
	var exprTmp string

L:
	exprSlice := strings.Split(expr, "")
	index := strings.LastIndex(expr, "(")
	lenExpr := len(exprSlice)
	if expr == exprTmp {
		zlog.Errorf("gfc 计算死循环,请检查公式%s是否存在中文字符", nil, expr)
		return 0, errors.New("gfc 计算死循环,请检查公式是否存在中文字符")
	}
	exprTmp = expr
	if index != -1 {
		var strBuff bytes.Buffer
		for i := index + 1; i < lenExpr; i++ {
			if exprSlice[i] != ")" {
				strBuff.WriteString(exprSlice[i])
				continue
			}
			break
		}
		r := strBuff.String()

		tmpResult, err := count(r)
		if err != nil {
			return 0, err
		}

		strBuff.Reset()
		strBuff.WriteString("(")
		strBuff.WriteString(r)
		strBuff.WriteString(")")

		expr = strings.Replace(expr, strBuff.String(), strconv.FormatFloat(tmpResult, 'f', 14, 64), -1)

		goto L
	}
	tmpResult, err := count(expr)
	zlog.Debugf("gfc 计算结果:%f", nil, tmpResult)
	return tmpResult, err
}

func count(data string) (float64, error) {
	if 1 == len(data) {
		return strconv.ParseFloat(data, 64)
	}
	arr := generateRPN(data)
	zlog.Debugf("gfc 解析结果:%s", nil, arr)
	return calculateRPN(arr)
}

func calculateRPN(datas []string) (float64, error) {
	var stack linkStack
	var flag bool = false
	var symbolChar string
	stack.init()
	for i := 0; i < len(datas); i++ {
		if isNumberString(datas[i]) {
			if f, err := strconv.ParseFloat(datas[i], 64); err != nil {
				zlog.Errorf("gfc 计算解析错误,无法将%s,转化成一个float64类型,请检查公式", err, datas[i])
				return 0, errors.New(fmt.Sprintf("gfc 计算解析错误,无法将%s,转化成一个float64类型,请检查公式", datas[i]))
			} else {
				if flag { //处理同时两个符号
					switch symbolChar {
					case "-":
						stack.push(0 - f)
					case "+":
						stack.push(f)
					}
					flag = false
					continue
				}
				stack.push(f)
			}
		} else {
			p1 := stack.pop()
			p2 := stack.pop()
			if p2 == nil && !isNumberString(datas[i]) { //如果p2为空 同时两个符号
				stack.push(p1)
				symbolChar = datas[i]
				flag = true
				continue
			}
			f1 := p1.(float64)

			f2 := p2.(float64)

			p3, err := normalCalculate(f2, f1, datas[i])

			if err != nil {
				return 0, err
			}
			stack.push(p3)
		}
	}
	res := stack.pop().(float64)
	//zlog.Debugf("gfc 计算结果:%f", nil, res)
	return res, nil
}

func normalCalculate(a, b float64, operation string) (float64, error) {
	switch operation {
	case "*":
		return a * b, nil
	case "-":
		return a - b, nil
	case "+":
		return a + b, nil
	case "/":
		if 0 == b {
			zlog.Error("gfc 计算遇到除数为0，默认返回0", nil)
			return 0, errors.New("gfc 计算遇到除数为0，默认返回0")
		} else {
			return a / b, nil
		}
	default:
		zlog.Errorf("gfc 不支持的运算符%s", nil, operation)
		return 0, errors.New(fmt.Sprintf("gfc 不支持的运算符%s", operation))
	}
}

func getExprSlice(exp string) []string {
	var symbolSlic []string

	var preBytes bytes.Buffer

	expBytes := []byte(exp)
	lenExp := len(expBytes)
	for i := 0; i < lenExp; i++ {
		if !isNumber(expBytes[i]) && "" != preBytes.String() { //符号
			symbolSlic = append(symbolSlic, preBytes.String())
			symbolSlic = append(symbolSlic, string(expBytes[i]))
			preBytes.Reset()
			continue
		}
		preBytes.WriteByte(expBytes[i])
	}
	if preBytes.Len() > 0 {
		symbolSlic = append(symbolSlic, preBytes.String())
	}
	return symbolSlic
}

func generateRPN(exp string) []string {

	var stack linkStack
	stack.init()

	var spiltedStr = getExprSlice(exp)
	var datas []string

	for i := 0; i < len(spiltedStr); i++ { // 遍历每一个字符
		tmp := spiltedStr[i] //当前字符

		if !isNumberString(tmp) { //是否是数字
			// 四种情况入栈
			// 1 左括号直接入栈
			// 2 栈内为空直接入栈
			// 3 栈顶为左括号，直接入栈
			// 4 当前元素不为右括号时，在比较栈顶元素与当前元素，如果当前元素大，直接入栈。
			if tmp == "(" ||
				stack.lookTop() == nil || stack.lookTop().(string) == "(" ||
				(compareOperator(tmp, stack.lookTop().(string)) == 1 && tmp != ")") {
				stack.push(tmp)
			} else { // ) priority
				if tmp == ")" { //当前元素为右括号时，提取操作符，直到碰见左括号
					for {
						popi := stack.pop()
						if popi != nil {
							if pop := popi.(string); pop == "(" {
								break
							} else {
								datas = append(datas, pop)
							}
						}
						break
					}
				} else { //当前元素为操作符时，不断地与栈顶元素比较直到遇到比自己小的（或者栈空了），然后入栈。
					for {
						pop := stack.lookTop()
						if pop != nil && compareOperator(tmp, pop.(string)) != 1 {
							datas = append(datas, stack.pop().(string))
						} else {
							stack.push(tmp)
							break
						}
					}
				}
			}

		} else {
			datas = append(datas, tmp)
		}
	}

	//将栈内剩余的操作符全部弹出。
	for {
		if pop := stack.pop(); pop != nil {
			datas = append(datas, pop.(string))
		} else {
			break
		}
	}
	return datas
}

// if return 1, o1 > o2.
// if return 0, o1 = 02
// if return -1, o1 < o2
func compareOperator(o1, o2 string) int {
	// + - * /
	var o1Priority int
	if o1 == "+" || o1 == "-" {
		o1Priority = 1
	} else {
		o1Priority = 2
	}
	var o2Priority int
	if o2 == "+" || o2 == "-" {
		o2Priority = 1
	} else {
		o2Priority = 2
	}
	if o1Priority > o2Priority {
		return 1
	} else if o1Priority == o2Priority {
		return 0
	} else {
		return -1
	}
}

func isNumberString(o1 string) bool {
	if o1 == "+" || o1 == "-" || o1 == "*" || o1 == "/" || o1 == "(" || o1 == ")" {
		return false
	} else {
		return true
	}
}

func convertToStrings(s string) []string {
	var strs []string
	bys := []byte(s)
	var tmp string
	for i := 0; i < len(bys); i++ {
		if !isNumber(bys[i]) {
			if tmp != "" {
				strs = append(strs, tmp)
				tmp = ""
			}
			strs = append(strs, string(bys[i]))
		} else {
			tmp = tmp + string(bys[i])
		}
	}
	strs = append(strs, tmp)
	return strs
}

func isNumber(o1 byte) bool {
	if o1 == '+' || o1 == '-' || o1 == '*' || o1 == '/' || o1 == '(' || o1 == ')' {
		return false
	} else {
		return true
	}
}
