package main

import (
	"github.com/lunny/tango"
	// "github.com/tango-contrib/renders"
	"os"
	"runtime"

	"pccqcpa.com.cn/app/rpm/api/routers"
	"pccqcpa.com.cn/app/rpm/api/util"

	// 引用pprof package

	// "log"
	// "net/http"
	// _ "net/http/pprof"
	_ "pccqcpa.com.cn/app/rpm/api/amInterfaces/dataAuthInterface"
	_ "pccqcpa.com.cn/app/rpm/api/creditInterfaces/socket"
	// _ "pccqcpa.com.cn/app/rpm/api/models/par"
	_ "pccqcpa.com.cn/app/rpm/api/util"

	"fmt"

	"pccqcpa.com.cn/app/rpm/api/middlewares"
	"pccqcpa.com.cn/app/rpm/api/middlewares/dataAuth"
	"pccqcpa.com.cn/components/zlog"
)

func main() {

	// f, _ := os.Create("profile_file")
	// pprof.StartCPUProfile(f)     // 开始cpu profile，结果写到文件f中
	// defer pprof.StopCPUProfile() // 结束profile

	// 注意，有时候 defer f.Close()， defer pprof.StopCPUProfile() 会执行不到，这时候我们就会看到 prof 文件是空的， 我们需要在自己代码退出的地方，增加上下面两行，确保写文件内容了。

	// flag.Parse()
	// go func() {
	// 	log.Println(http.ListenAndServe("localhost:7777", nil))
	// }()
	tg := tango.Classic()
	// tg.Use(renders.New(renders.Options{
	// 	Reload:    false,
	// 	Directory: "./templates",
	// }))
	//使用登陆验证中间件

	tg.Use(middlewares.CorsSeting())
	tg.Use(middlewares.VerificationSid())
	tg.Use(dataAuth.HandleAuthData())
	// tg.Use(auth.LoginAuths(sessions, "/api/sign/"))
	// tg.Get("/login", tango.File("./public/RPM/pages/login.html"))
	// tg.Get("/dianCan", tango.File("./public/菜单.jpg"))
	// tg.Get("/:name", tango.Dir("./../public"))
	routers.Init(tg)

	zlog.Info("欢迎使用重庆天健金融有限公司外部定价产品", nil)
	appPort := util.GetIniIntValue("system", "appPort")
	fmt.Println(appPort)
	if 0 == appPort {
		zlog.Error("读取conf/system.ini中[system]的port端口号为空", nil)
		os.Exit(1)
	}

	runtime.GOMAXPROCS(runtime.NumCPU() * 2) // 设置多核数
	tg.Use(tango.Static())
	tg.Run(appPort)

	//tg.RunTLS("server.crt", "server.key", appPort)

}
