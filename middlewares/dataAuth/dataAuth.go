package dataAuth

import (
	"errors"
	"fmt"
	"strings"

	"github.com/lunny/tango"
	"pccqcpa.com.cn/app/rpm/api/amInterfaces/dataAuthInterface"
	"pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/app/rpm/api/util/currentMsg"
	"pccqcpa.com.cn/components/zlog"
)

func HandleAuthData() tango.HandlerFunc {
	return func(ctx *tango.Context) {
		fmt.Println(ctx.Req().URL.Path, "/" == ctx.Req().URL.Path || "/index.html" == ctx.Req().URL.Path)
		if "/" == ctx.Req().URL.Path || strings.Contains(ctx.Req().URL.Path, ".") ||
			("POST" == ctx.Req().Method && "/api/rpm/work/flow/ln/pricing" == ctx.Req().URL.Path) {
			ctx.Next()
			return
		}

		sid := ctx.Req().Header.Get("Sid")

		if "" == sid {
			er := errors.New("请求Header未传Sid")
			zlog.Error(er.Error(), er)
			ctx.Abort(403, er.Error())
		}
		sidUser, err := currentMsg.GetCurrentUser(tango.Ctx{ctx})
		if nil != err {
			ctx.Abort(403, err.Error())
		}
		var userId string = sidUser.UserId

		path := ctx.Req().URL.Path
		paramsLength := len(*ctx.Params())
		var routerUrl string = path
		for _, param := range *ctx.Params() {
			// fmt.Printf("参数Name【%v】参数值【%v】\n", param.Name, param.Value)
			routerUrl = strings.Replace(path, param.Value, "("+param.Name+")", 1)
		}

		amAuthStr := getAuthData(userId, routerUrl)
		// fmt.Printf("\n\n\n\n\n\nuserId[%v],routerUrl[%v],amAuthStr[%v]\n", userId, routerUrl, amAuthStr)
		if 0 == paramsLength {
			if util.IsPaginQuery(&tango.Ctx{ctx}) {
				ctx.Params().Set(util.IS_PAGIN_QUERY, "true")
			}
		}
		if nil != amAuthStr {
			for k, v := range amAuthStr {
				v = strings.Replace(v, "''", "'", -1)
				v = strings.Replace(v, ",,", ",", -1)
				v = strings.TrimRight(v, ",")

				ctx.Params().Set(k, v)
			}
		}

		// fmt.Printf("%#v", ctx.Params())
		ctx.Next()
	}
}

func getAuthData(userId, reqUrl string) map[string]string {
	data := util.GetCacheByCacheName(util.AM_DATA_AUTH_CACHE, userId)
	// fmt.Printf("data[%#v]\n\n\n\n", data)
	if nil == data {
		return nil
	}
	authDatas := data.(map[string][]dataAuthInterface.DataAuth)[reqUrl]
	// fmt.Printf("authDatas[%#v]\n\n\n\n", authDatas)
	if 0 == len(authDatas) {
		return nil
	}

	var amAuthMap = make(map[string]string)

	for _, dataAuth := range authDatas {
		amAuthStr := amAuthMap[dataAuth.ConditionType]

		if strings.Contains(dataAuth.ConditionContent, ",") {
			str := strings.Replace(dataAuth.ConditionContent, ",", "','", -1)
			amAuthStr += "'" + str + "',"
		} else {
			amAuthStr += "'" + dataAuth.ConditionContent + "',"
		}
		amAuthMap[dataAuth.ConditionType] = amAuthStr
	}

	// organs = strings.Replace(organs, "''", "'", -1)
	// organs = strings.Replace(organs, ",,", ",", -1)
	// organs = strings.TrimRight(organs, ",")

	// fmt.Println("\n\n\norgans", organs)
	return amAuthMap
}
