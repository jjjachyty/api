package socket

import (
	"fmt"
	"net"
	"os"

	"pccqcpa.com.cn/app/rpm/api/creditInterfaces/socket/handle"
	u "pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/components/zlog"
)

func socketServer() {

	//conf.Parse()
	port := u.GetIniIntValue("socket", "listenAddr")
	httpsPort := u.GetIniIntValue("system", "appPort")
	// 判断端口号是否已经占用
	//tcpaddr := strings.SplitN(port, ":", 2)
	// 检查端口是否与其他启用端口重复
	if httpsPort == port {
		zlog.Error("socket端口与https端口重复", nil)
		os.Exit(1)
	}

	//ip := tcpaddr[0]
	// localPort, err := strconv.Atoi(tcpaddr[1])
	// if nil != err {
	// 	zlog.Error("端口字符串转换为数字出错", err)
	// 	os.Exit(1)
	// }

	addr := &net.TCPAddr{net.ParseIP(""), port, ""}
	tcpConn, err := net.DialTCP("tcp", nil, addr)
	if nil == err {
		zlog.Errorf("本地端口[%i]已经被占用，请使用其他端口作为socket的监听端口", nil, port)
		tcpConn.Close()
		os.Exit(1)
	}

	fmt.Printf("开始启动socket端口[%s]接口监听\n", port)
	netListen, err := net.Listen("tcp", addr.String())
	if err != nil {
		zlog.Errorf("启动socket端口号[%s]监听接口出错,退出程序\n", err, port)
		os.Exit(1)
	}
	fmt.Printf("socket接口监听端口[%s]启动成功\n", port)
	for {
		fmt.Println("开始监听端口")
		conn, err := netListen.Accept()
		if nil != err {
			fmt.Println(err)
			continue
		}
		go func() {
			zlog.Info(conn.RemoteAddr().String()+" tcp connect success", nil)
			(handle.Handle{}).HandleConnection(conn)
		}()
	}
}

// 建立socekt，监听端口9091
func init() {
	go socketServer()
}
