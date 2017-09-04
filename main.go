package main

import (
	"flag"
	"os"
	"os/signal"
	"runtime"
	"runtime/debug"
	"syscall"
	"time"

	"net/http"
	_ "net/http/pprof"

	//"context"
	l4g "github.com/alecthomas/log4go"
	"djforgo/config"
	"djforgo/server"
	"djforgo/utils"
)

func main() {
	defer time.Sleep(time.Millisecond * 50)
	//异常捕捉
	defer func() {
		if x := recover(); x != nil {
			l4g.Error("main recover:", x)
			l4g.Error(string(debug.Stack()))
			debug.PrintStack()
		}
	}()

	//命令行解析
	cmdVersion := flag.Bool("version", false, "show version information")
	logcfgfile := flag.String("lf", "log4go.xml", "log4go cfgfile")
	appcfgfile := flag.String("f", "config.json", "config file path")
	help := flag.Bool("help", false, "help")
	flag.Parse()

	//打印参数帮助列表
	if *help {
		flag.PrintDefaults()
		return
	}

	//打印版本
	if *cmdVersion {
		l4g.Info("%s.%s.%s.%s\n", MajorVersion, MinorVersion, ReviseVersion, BuildVersion)
		return
	}

	//初始化配置
	config.LoadConfig(appcfgfile)
	l4g.Info("%v", config.QasConfig)

	if len(*logcfgfile) > 0 {
		l4g.LoadConfiguration(*logcfgfile)
	}
	defer l4g.Close()

	//设置CUP数量
	l4g.Debug(runtime.NumCPU())
	l4g.Debug(runtime.GOMAXPROCS(runtime.NumCPU()))

	//启动堆栈状态监听线程
	if config.QasConfig != nil {
		addr := config.QasConfig.Pprof
		go func() {
			http.ListenAndServe(addr, nil)
		}()
	}

	//监控
	go utils.MetricStart()

	//>>>>>>>>>>>>>>>>工作区>>>>>>>>>>>>>>>>>>>>>>>>>>
	//ctx, cancel := context.WithCancel(context.Background())
	server.ServerInstance.OnInit()
	server.ServerInstance.OnWork()
	defer server.ServerInstance.OnClose()

	//<<<<<<<<<<<<<<<<<工作区<<<<<<<<<<<<<<<<<<<<<<<<<<<

	//等待退出
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGTERM)
	<-c
	//cancel()
	l4g.Info("Receive ctrl-c")
}
