package middlewares

import (
	"github.com/lunny/tango"
	"pccqcpa.com.cn/app/rpm/api/util"
)

func CorsSeting() tango.HandlerFunc {
	return func(res *tango.Context) {

		origin := util.GetIniStringValue("cors", "origin")

		res.Header().Set("Access-Control-Allow-Origin", origin)
		// ctx.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		res.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With,Start-Row-Number,Order-Attr,Order-Type,Page-Size,Sid")
		res.Header().Set("Access-Control-Allow-Credentials", "true")
		res.ResponseWriter.Header().Set("Access-Control-Allow-Credentials", "true")
		res.ResponseWriter.Header().Set("Access-Control-Allow-Origin", origin)
		res.ResponseWriter.Header().Set("Access-Control-Allow-Headers", "Content-Type, text/plain, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With,Start-Row-Number,Order-Attr,Order-Type,Page-Size,Sid")
		res.ResponseWriter.Header().Set("Access-Control-Allow-Credentials", "true")

		//复杂请求
		// fmt.Println(res.Req().Method)
		if res.Req().Method != "OPTIONS" {
			res.Next()
			return
		}

		// fmt.Println(res.Header().Get("Access-Control-Request-headers"))
		res.Header().Add("Vary", "Origin")
		res.Header().Add("Vary", "Access-Control-Request-Method")
		res.Header().Add("Vary", "Access-Control-Request-Headers")
		res.Header().Set("Access-Control-Allow-Origin", origin)
		res.Header().Set("Access-Control-Allow-Methods", "PUT,DELETE,POST,GET,PATCH")
		res.Header().Set("Access-Control-Allow-Credentials", "true")
		// res.Header().Set("Access-Control-Allow-Headers", "x-requested-with,cache-control,Start-Row-Number,Page-Size,Order-Attr,Order-Type,Sid")
		res.Header().Set("Access-Control-Allow-Headers", "Content-Type, text/plain, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With,Start-Row-Number,Order-Attr,Order-Type,Page-Size,Sid")
		res.ServeJson("")

		return
	}
}
