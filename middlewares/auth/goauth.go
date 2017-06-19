package auth

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"net/http"
	"time"

	"strings"

	"pccqcpa.com.cn/app/rpm/api/middlewares"

	"github.com/lunny/tango"
	"github.com/tango-contrib/session"
)

type GOAuth struct {
	BaseUrl string
	session.Session
}

// var authKey = session.NewMemoryStore(time.Minute * 1)

// var authKey = session.NewCookieTracker(name, maxAge, secure, rootPath)

// func SetAuthKey(key, value string) bool {
// 	authKey.Run()
// 	err := authKey.Set(session.Id(key), key, value)
// 	if nil != err {
// 		return false
// 	}
// 	authKey.SetIdMaxAge(session.Id(key), time.Second*20)
// 	return true
// }

// func GetAuthKey(key string) (string, bool) {

// 	authKey.Exist(session.Id(key))
// 	v := authKey.Get(session.Id(key), key)
// 	if nil == v {

// 		return "", false
// 	}
// 	return v.(string), true

// }

// func DeleteAuthKey(key string) bool {
// 	ok := authKey.Del(session.Id(key), key)
// 	return ok
// }

var authKey = middlewares.NewMemoryCooke(time.Second * 10)

func SetAuthKey(key, value string) bool {
	err := authKey.Set(key, value)
	if nil == err {
		return true
	}
	return false
}

func GetAuthKey(key string) (string, bool) {

	return authKey.Get(key)

}

func DeleteAuthKey(key string) bool {
	ok := authKey.Del(key)
	return ok
}

func Get() interface{} {
	return authKey
}

func LoginAuths(sessions *session.Sessions, baseUrl string) tango.HandlerFunc {

	return func(ctx *tango.Context) {

		// fmt.Println("------------", ctx.Req().URL)

		// aa := sessions.Session(ctx.Req(), ctx.ResponseWriter)
		// fmt.Println("中间件的session：", aa)
		// usernaem := ctx.Req().Header.Get("userName")
		// aaa := aa.Get("aa")
		// fmt.Println("keys, has:----------", aaa)
		// fmt.Println("sessionId kkkk ", aa.Id())
		// fmt.Println("afadfa", aa.Get("aa"))
		// fmt.Println("begainLoginAuth....")
		// fmt.Println(usernaem)
		url := ctx.Req().URL
		//判断访问url是否为登陆的url，如果是则跳过验证，
		//如果不是，则验证是否登陆
		//如果已登陆过，则继续，否则跳转至重定向到页面
		if 0 == strings.Index(url.String(), baseUrl) {

			ctx.Next()
			return
		}

		//不是登陆url
		token := ""

		ctx.Next()
		return

		//根据不同的访问方法调用,使用不同的方式获取参数
		method := ctx.Req().Method
		fmt.Println(ctx.Req().Header.Get("userName"))
		fmt.Println(ctx.Req().Header.Get("token"))
		switch method {
		case "POST", "PUT", "PATCH":
			fmt.Println("POST", "PUT")

		case "GET", "DELETE", "HEAD", "TRACE", "OPTIONS":
			fmt.Println("GET", "DELETE")
		}

		fmt.Println("...")
		userName := ctx.Form("userKey")
		userKey, ok := GetAuthKey(userName)
		fmt.Println("userKey:   ", userKey)

		if !ok {

			ctx.ServeJson(map[string]string{"Code": "400", "Msg": "尚未登录，请先登录"})
			return
		}
		fmt.Println(".........", userKey)
		if "" == userKey {
			var a []byte
			copy(a[:], "尚未登录，请先登录")
			ctx.ResponseWriter.Write(a)
			return
		}

		//加密验证

		form := ctx.Forms().Form
		for k, v := range form {
			fmt.Println(k, v)
		}
		params := ctx.Header().Get("params")

		m := md5.New()
		var str string = params + userKey
		var bytes []byte
		copy(bytes[:], str)
		m.Write(bytes)
		encoding := m.Sum(nil)
		encodingstring := hex.EncodeToString(encoding)
		if token == encodingstring {
			ctx.Next()
			return
		}

		return
	}
}

// defaultUnauthorizedHandler provides a default HTTP 401 Unauthorized response.
func defaultUnauthorizedHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("defaultUnauthorizedHandler")
	http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
}
