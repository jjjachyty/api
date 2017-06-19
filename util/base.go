package util

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode"

	"github.com/lunny/tango"

	"pccqcpa.com.cn/components/zlog"
)

type RstMsg struct {
	RstCode  int
	RstMsg   string
	Err      error
	ErrMsg   string
	Data     interface{}
	PageData PageData
}
type RspData struct {
	Status  string          `json:"code,"`
	Message string          `json:"msg,"`
	Data    json.RawMessage `json:"data,"`
}

//PageData type分页查询结果集
type PageData struct {
	Page Page
	Rows interface{}
}

/**
* 获取路由参数并转换为map返回
* Jason_he
 */
func GetParmFromRouter(this *tango.Ctx) (map[string]interface{}, error) {

	rst := make(map[string]interface{})

	for _, param := range *this.Params() {
		switch param.Name {
		case "*param":
			paramArray := strings.Split(param.Value, "/")
			if len(paramArray)%2 != 0 {
				err := fmt.Errorf("路由参数不规范，请保证为key/value形式")
				return rst, err
			}
			for i := 0; i < len(paramArray); i = i + 2 {
				key := UperChange(paramArray[i])
				rst[key] = paramArray[i+1]
			}

		case ":organ":
			searchLike, ok := rst[SEARCH_LIKE]
			if !ok {
				searchLike = make([]map[string]interface{}, 0)
			}
			searchLike = append(searchLike.([]map[string]interface{}), map[string]interface{}{
				"key":   "organ",
				"type":  "in",
				"value": param.Value,
			})
			rst[SEARCH_LIKE] = searchLike

		case ":branch":
			searchLike, ok := rst[SEARCH_LIKE]
			if !ok {
				searchLike = make([]map[string]interface{}, 0)
			}
			searchLike = append(searchLike.([]map[string]interface{}), map[string]interface{}{
				"key":   "branch",
				"type":  "in",
				"value": param.Value,
			})
			rst[SEARCH_LIKE] = searchLike

		default:
			zlog.Infof("路由参数[%v]未处理", nil, param.Name)
		}
	}
	fmt.Printf("参数【%#v】", rst)
	return rst, nil
}

//post、put参数转换为结构体
func ParamsToStruct1(values url.Values, entity interface{}) error {
	ok := reflect.ValueOf(entity).CanSet()
	if !ok {
		err := fmt.Errorf("请传实体地址")
		zlog.Error("请传实体地址", err)
		return err
	}
	s := reflect.ValueOf(entity).Elem()
	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fieldName := typeOfT.Field(i).Name //属性名称
		fieldType := f.Type().String()     //属性类型
		value, ok := values[fieldName]     //前台是否传了该参数
		if !f.CanSet() {
			zlog.Info("不能操作该属性:"+fieldName, nil)
			continue
		}
		if ok {
			switch fieldType {
			case "string":
				f.SetString(value[0])
				continue
			case "int", "int32", "int64":
				//参数1 ：字符串
				//参数2 ：转换到进制 8 10 16
				//参数3 ：bitSize 32 64
				v, err := strconv.ParseInt(value[0], 10, 32)
				if nil != err {
					f.SetInt(v)
				}
				continue
			case "float32", "float64":
				v, err := strconv.ParseFloat(value[0], 64)
				if nil != err {
					f.SetFloat(v)
				}
				continue
			case "time.Time":
				date, err := time.Parse("2006-01-02 15:04:05", value[0])
				if nil != err {
					date = time.Now()
				}
				d := reflect.ValueOf(date)
				f.Set(d)
				continue
			case "interface {}":
				v := reflect.ValueOf(value[0])
				f.Set(v)
				continue
			case "struct":
				fmt.Println("这是一个结构体")
			default:
				fmt.Println("暂不支持该类型反射：", fieldType)
			}
		}
	}
	fmt.Println(entity)
	return nil
}

//post、put参数转换为结构体
func ParamsToStruct(values url.Values, entity interface{}) error {
	fmt.Println("前台参数", values)
	fmt.Println("Branch实体：", values["Branch"])
	s := reflect.ValueOf(entity).Elem()
	for k, v := range values {
		if s.CanSet() {
			fieldName := strings.Title(k)
			//过滤参数
			if "$$hashkey" == strings.ToLower(fieldName) || "indicatorbeforerelyarr" == strings.ToLower(fieldName) {
				continue
			}
			//打印前台传的参数名称
			fmt.Println("fieldName: ", fieldName)

			//判断是否可操作
			if !s.FieldByName(fieldName).IsValid() {
				return fmt.Errorf("前台传的参数(%s)与对应的结构体属性不对应，请检查！", fieldName)
			}
			fmt.Println("结构体属性类型", s.FieldByName(fieldName).Type().String())
			switch s.FieldByName(fieldName).Type().String() {
			case "string":
				s.FieldByName(fieldName).SetString(v[0])
			case "int32", "int":
				//参数1 ：字符串
				//参数2 ：转换到进制 8 10 16
				//参数3 ：bitSize 32 64
				value, err := strconv.ParseInt(v[0], 10, 32)
				if nil == err {
					s.FieldByName(fieldName).SetInt(value)
				}
			case "float64", "float32":
				value, err := strconv.ParseFloat(v[0], 64)
				if nil == err {
					s.FieldByName(fieldName).SetFloat(value)
				}
			case "time.Time":
				date, err := time.Parse("2006-01-02 15:04:05", v[0])
				if nil != err {
					date = time.Now()
				}
				value := reflect.ValueOf(date)
				s.FieldByName(fieldName).Set(value)
			case "interface {}":
				value := reflect.ValueOf(interface{}(v[0]))
				s.FieldByName(fieldName).Set(value)
			default:
				return fmt.Errorf("前台传的参数(%s)与对应的结构体属性不对应，请检查！", fieldName)

			}
		} else {
			zlog.Info("反射不能被赋值", nil)
			return fmt.Errorf("反射不能被赋值！")
		}
	}
	return nil
}

// BUG(Jason) : #1: 暂时不知道怎么处理time.Time类型      已解决
// BUG(Jason) : #2: 暂时不知道怎么处理interface{}类型    已解决

//get请求参数转化为map
func ParamsToMap(this *tango.Ctx) map[string]interface{} {
	rst := make(map[string]interface{})
	urlParams := this.Forms().Form
	for k, v := range urlParams {
		//处理模糊查询字段
		if "search" == k {

			var dat []map[string]interface{}
			err := json.Unmarshal([]byte(v[0]), &dat)
			if nil != err {
				zlog.Error("json转换map出错", err)
			}

			rst["searchLike"] = dat
			continue
		}

		key := UperChange(k)
		rst[key] = v[0]
	}
	return rst
}
func SuccessListMsg(format string, pageData *PageData, a ...interface{}) RstMsg {
	var rst RstMsg
	rst.RstCode = 200
	rst.RstMsg = fmt.Sprintf(format, a...)
	rst.PageData = *pageData
	return rst
}
func SuccessMsg(msg string, data interface{}) RstMsg {
	var rst RstMsg
	rst.RstCode = 200
	rst.RstMsg = msg
	rst.Err = nil
	rst.Data = nil
	var pageData PageData
	pageData.Rows = data
	rst.PageData = pageData
	return rst
}
func SuccessDataMsg(msg string, data interface{}) RstMsg {
	var rst RstMsg
	rst.RstCode = 200
	rst.RstMsg = msg
	rst.Err = nil
	rst.Data = data
	return rst
}

func RetrunSuccessPage(msg string, param map[string]interface{}) RstMsg {
	var rst RstMsg
	rst.RstCode = 200
	rst.RstMsg = msg
	rst.Err = nil
	rst.Data = nil
	var pageData PageData
	pageData.Page.TotalRows = param["Total"].(int64)
	pageData.Rows = param["Items"]
	rst.PageData = pageData
	return rst
}

func ReturnSuccess(msg string, pageData PageData) RstMsg {
	var rst RstMsg
	rst.RstCode = 200
	rst.RstMsg = msg
	rst.Err = nil
	rst.PageData = pageData
	return rst
}

func ErrorMsg(msg string, err error) RstMsg {
	var rst RstMsg
	rst.RstCode = 400
	rst.RstMsg = msg
	rst.Err = err
	if nil != err {
		rst.ErrMsg = err.Error()
	}
	rst.Data = nil
	return rst
}

func ReturnErrorRedirectMsg(rstCode int, msg string, data interface{}, err error) RstMsg {
	var rst RstMsg
	rst.RstCode = rstCode
	rst.RstMsg = msg
	rst.Err = err
	rst.ErrMsg = err.Error()
	rst.PageData.Rows = data
	return rst
}

//将字符串中的大写字母转换为小写字母，并加下划线
//第一个字母如果是大写，则直接转换为小写，不加下划线
func UperChange(str string) string {
	var key string
	for i, v := range str {
		if 0 < i && unicode.IsUpper(v) {
			key += "_"
		}
		key += strings.ToLower(string(v))
	}
	return key
}

// 获取当前时间字符串
func GetCurrentTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// 获取当前登录用户
// func GetCurrentUser() string {
// 	return "Jason"
// }

// 获取当前登陆用户机构信息
// func GetCurrentOrgan() sys.Organ {
// 	return sys.Organ{
// 		OrganCode: "panchina001",
// 	}
// }

// 判断当前登录用户是否为客户经理
func CurrentUserIsCm() bool {

	return false
}

// 判断多参数查询是否为分页查询
func IsPaginQuery(ctx *tango.Ctx) bool {
	// startRowNum := ctx.Req().Header.Get("start-row-number")
	if "" != ctx.Params().Get(":"+IS_PAGIN_QUERY) {
		return true
	}
	if "" == ctx.Req().Header.Get("Start-Row-Number") || "-1" == ctx.Req().Header.Get("Page-Size") {
		return false
	}
	return true
}

// 将header中的分页信息放入param中
func GetPageMsg(ctx *tango.Ctx, param ...map[string]interface{}) map[string]interface{} {
	var paramMap map[string]interface{}
	var err error = nil
	if 0 < len(param) {
		paramMap = param[0]
	} else {
		paramMap, err = GetParmFromRouter(ctx)
		if nil != err {
			return paramMap
		}
	}
	start := ctx.Req().Header.Get("start-row-number") // 开始记录数
	length := ctx.Req().Header.Get("Page-Size")       // 每页记录数
	sort := ctx.Req().Header.Get("order-attr")        // 排序字段
	order := ctx.Req().Header.Get("order-type")       // 排序desc asc
	if "" == start {
		start = "0"
	}
	if "" == length {
		length = "15"
	}

	paramMap["sort"] = sort
	paramMap["order"] = order
	paramMap["start"] = start
	paramMap["length"] = length

	return paramMap
}

func CheckQueryParams(pars map[string]interface{}) bool {
	var flag = false
	reg := regexp.MustCompile("[\\^\\$\\*\\+\\?\\{\\}\\(\\)\\[\\]\\|\\'\\`\\~@＃&、～｀]")
	if len(pars) != 0 {
		for _, v := range pars {
			switch v.(type) {
			case string:
				flag = reg.Match([]byte(v.(string)))
				if flag {
					break
				}
			}
		}
	}
	return flag
}

func GetStartTimeParam(paramMap map[string]interface{}) error {
	timeNow, err := time.Parse("2006-01-02", time.Now().Format("2006-01-02"))
	if nil != err {
		er := fmt.Errorf("格式化时间出错")
		zlog.Error(er.Error(), err)
		return er
	}
	var searchLike = []map[string]interface{}{
		map[string]interface{}{
			"type":  "le",
			"key":   "to_date(to_char(t.start_time,'yyyy-mm-dd'),'yyyy-mm-dd')",
			"value": timeNow,
		},
	}
	paramMap["searchLike"] = searchLike
	paramMap["sort"] = "startTime"
	paramMap["order"] = "DESC"
	return nil
}

// func CopyMap(map1, map2 interface{}) error {
// 	bytes, err := json.Marshal(map1)
// 	if nil != err {
// 		er := fmt.Errorf("深度拷贝MAP序列化实体出错")
// 		zlog.Error(er.Error(), err)
// 		return er
// 	}
// 	err = json.Unmarshal(bytes, map2)
// 	if nil != err {
// 		er := fmt.Errorf("深度拷贝MAP反序列化实体出错")
// 		zlog.Error(er.Error(), err)
// 		return er
// 	}
// 	return nil
// }

type UtilBase struct {
}

func GetUtilBase() *UtilBase {
	return new(UtilBase)
}

func (UtilBase) JoinPageData(callBack func(interface{}, interface{}) interface{}, pageDatas ...*PageData) *PageData {

	for index, pageData := range pageDatas {
		if 0 == index {
			continue
		} else {
			pageDatas[0].Page.TotalRows += pageData.Page.TotalRows
			switch iType := pageDatas[0].Rows.(type) {

			default:
				if reflect.Slice == reflect.TypeOf(iType).Kind() {
					// pageDatas[0].Rows = append(pageDatas[0].Rows.(iType), pageData.Rows.(iType)...)
					pageDatas[0].Rows = callBack(pageDatas[0].Rows, pageData.Rows)
				}
			}
		}
	}
	if 0 == len(pageDatas) {
		return nil
	} else {
		if 0 < pageDatas[0].Page.PageSize {
			pageDatas[0].Page.TotalPage = pageDatas[0].Page.TotalRows / pageDatas[0].Page.PageSize
			if 0 != pageDatas[0].Page.TotalRows%pageDatas[0].Page.PageSize {
				pageDatas[0].Page.TotalPage += 1
			}
		}
		return pageDatas[0]
	}
}

func (UtilBase) HandleParamToLike(paramMap map[string]interface{}, keys ...string) {
	for _, key := range keys {
		if value, ok := paramMap[key]; ok {
			var searchLike []map[string]interface{}
			if val, ok := paramMap[SEARCH_LIKE]; ok {
				searchLike = val.([]map[string]interface{})
			}
			searchLike = append(searchLike, map[string]interface{}{
				"key":   key,
				"type":  "like",
				"value": value,
			})
			delete(paramMap, key)
			paramMap[SEARCH_LIKE] = searchLike
		}
	}
}

// 深度拷贝Map
func (u UtilBase) DeepCopy(value interface{}) interface{} {
	if valueMap, ok := value.(map[string]interface{}); ok {
		newMap := make(map[string]interface{})
		for k, v := range valueMap {
			newMap[k] = u.DeepCopy(v)
		}

		return newMap
	} else if valueSlice, ok := value.([]interface{}); ok {
		newSlice := make([]interface{}, len(valueSlice))
		for k, v := range valueSlice {
			newSlice[k] = u.DeepCopy(v)
		}

		return newSlice
	}

	return value
}
func GetApiPath() string {

	confPath := os.Getenv("GOSYSCONFIG")
	by := []byte(confPath)
	return string(by[:len(by)-5])
}

// 参数1: 主map
// 参数2: 目标map
// 参数3: 需要转换的key数组
// 若rMap中已存在key，则替换值
func SplitMap(mMap, rMap map[string]interface{}, keys []string) {
	if nil == mMap || nil == rMap {
		return
	}
	for _, key := range keys {
		if val, ok := mMap[key]; ok {
			rMap[key] = val
			delete(mMap, key)
		}
	}
}

// func CtxSet(ctx tango.Ctx) {
// 	ctx.Header().Set("Access-Control-Allow-Origin", "127.0.0.1")
// 	ctx.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
// 	ctx.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, Total-Page, Total-Rows, Start-Row-Number")
// 	ctx.Header().Set("Access-Control-Allow-Credentials", "true")
// 	ctx.ResponseWriter.Header().Set("Access-Control-Allow-Credentials", "true")
// 	ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", "127.0.0.1")
// 	ctx.ResponseWriter.Header().Set("Access-Control-Allow-Headers", "Content-Type, text/plain, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With,Total-Page,Total-Rows,Start-Row-Number")
// 	ctx.ResponseWriter.Header().Set("Access-Control-Allow-Credentials", "true")
// }
