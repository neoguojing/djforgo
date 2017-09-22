package main

import (
	"flag"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"net/http"
	_ "net/http/pprof"

	//"context"
	"djforgo/server"
	"djforgo/system"
	"djforgo/utils"
	l4g "github.com/alecthomas/log4go"
)

func main() {

	//命令行解析
	logcfgfile := flag.String("lf", "log4go.xml", "log4go cfgfile")
	appcfgfile := flag.String("f", "config.json", "config file path")
	help := flag.Bool("help", false, "help")
	flag.Parse()

	//打印参数帮助列表
	if *help {
		flag.PrintDefaults()
		return
	}

	system.LoadConfig(appcfgfile)
	l4g.Info("%v", system.SysConfig)

	if len(*logcfgfile) > 0 {
		l4g.LoadConfiguration(*logcfgfile)
	}
	defer l4g.Close()

	l4g.Debug(runtime.NumCPU())
	l4g.Debug(runtime.GOMAXPROCS(runtime.NumCPU()))

	if system.SysConfig != nil {
		addr := system.SysConfig.Pprof
		go func() {
			http.ListenAndServe(addr, nil)
		}()
	}

	go utils.PrometheusMonitorStart()

	//ctx, cancel := context.WithCancel(context.Background())
	err := server.ServerInstance.OnInit()
	if err != nil {
		return
	}

	server.ServerInstance.OnWork()
	defer server.ServerInstance.OnClose()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGTERM)
	<-c
	//cancel()
	l4g.Info("Receive ctrl-c")
}
