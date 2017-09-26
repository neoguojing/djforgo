package server

import (
	"djforgo/dao"
	"djforgo/sessions"
	"djforgo/system"
	"djforgo/urls"
	"fmt"
	l4g "github.com/alecthomas/log4go"
	"net/http"
)

type Server struct {
}

func NewServer() *Server {
	return &Server{}
}

var ServerInstance = NewServer()

func (this *Server) OnInit() error {
	//database relate init
	err := dao.DB_Init()
	dao.InitSrores()
	
	//session init
	sessions.InitSessionStore()
	http.Handle("/", urls.G_Router)
	l4g.Info("http://%s:%s/\n", system.SysConfig.Downnet.HttpIP, system.SysConfig.Downnet.Port)

	

	return err
}

func (this *Server) OnWork() {
	err := http.ListenAndServe(fmt.Sprintf("%s:%s", system.SysConfig.Downnet.HttpIP,
		system.SysConfig.Downnet.Port), nil)
	if err != nil {
		l4g.Error(err)
	}
}

func (this *Server) OnClose() {
	dao.DB_Destroy()
}
