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
	sessions.InitSessionStore()
	http.Handle("/", urls.G_Router)
	l4g.Info("http://%s:%s/\n", system.QasConfig.Downnet.HttpIP, system.QasConfig.Downnet.Port)

	err := dao.DB_Init()

	return err
}

func (this *Server) OnWork() {
	err := http.ListenAndServe(fmt.Sprintf("%s:%s", system.QasConfig.Downnet.HttpIP,
		system.QasConfig.Downnet.Port), nil)
	if err != nil {
		l4g.Error(err)
	}
}

func (this *Server) OnClose() {
	dao.DB_Destroy()
}
