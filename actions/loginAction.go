package actions

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"

	"pccqcpa.com.cn/components/zlog"

	"pccqcpa.com.cn/app/rpm/api/models/sys"

	"github.com/lunny/tango"
	"github.com/tango-contrib/renders"
	"github.com/tango-contrib/session"
	"pccqcpa.com.cn/app/rpm/api/middlewares/auth"
	"pccqcpa.com.cn/app/rpm/api/util"
)

type LoginAction struct {
	tango.Json
	tango.Ctx
	session.Session
	renders.Renderer
}

//登录
func (this *LoginAction) Get() interface{} {

	fmt.Println("sessionId", this.Session.Id())

	// util.CtxSet(this.Ctx)

	params, err := util.GetParmFromRouter(&this.Ctx)
	userName := params["username"]
	password := params["password"]

	//查询数据中有无此数据
	user, err := sys.SelectUserByParam(userName.(string), password.(string))
	if nil != err {
		return util.ErrorMsg("msg", err)
	}
	this.Session.Set(user.UserName, user.UserName)
	this.Session.Set("aa", "bb")
	this.Session.Set("cc", "ccc")
	fmt.Println("aaaaaa  ", this.Session.Get(user.UserName))
	fmt.Println(this.Session)

	if nil == err {
		v, _ := auth.GetAuthKey(user.UserName)
		fmt.Println("auth.AuthKey=============", auth.Get())
		fmt.Println("我是第一次getAuthKey", v)
		md5Str := stringToMd5(user.UserName)
		//fmt.Println(this.Session.Get("userKey"))
		//this.Session.Set(user.UserName, md5Str)
		ok := auth.SetAuthKey(user.UserName, md5Str)
		if ok {
			v, _ := auth.GetAuthKey(user.UserName)
			fmt.Println("我是第二次getAuthKey", v)

		}
		BB, _ := auth.GetAuthKey("BB")
		fmt.Println("我是测试BB的getAuthKey", BB)
		//如果数据库中没有改用户数据，则返回提示信息
		rst := util.SuccessMsg("登录成功", user)
		return rst
	}
	//如果数据库中没有改用户数据，则返回提示信息

	rst := util.ErrorMsg("用户名或密码错误，请核对后重新登陆", nil)
	return rst
}

//登出
func (this *LoginAction) Delete() interface{} {
	key := this.Ctx.Header().Get("userName")
	ok := auth.DeleteAuthKey(key)
	if ok {
		rst := util.SuccessMsg("登出成功", nil)
		return rst
	}
	rst := util.ErrorMsg("登出失败", nil)
	return rst
}

//参数1 ： 前台传送过来的密文
//参数2 ： 前台传的参数 + userKey
func isLogin(webCipherStr, str string) bool {
	cipherStr := stringToMd5(str)
	//判断前台的密文是否等于后台的密文
	//如果相等则有权限访问
	//如果不等则返回登陆页面
	if cipherStr == webCipherStr {
		return true
	}
	return false
}

func stringToMd5(str string) string {
	md5Ctx := md5.New()
	var bytes []byte
	copy(bytes[:], str)
	md5Ctx.Write(bytes)
	cipherStr := md5Ctx.Sum(nil)
	rst := hex.EncodeToString(cipherStr)
	zlog.Info("加密后的密文："+rst, nil)
	return rst
}
