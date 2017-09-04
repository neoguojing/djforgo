package server

import (
	"fmt"
	l4g "github.com/alecthomas/log4go"
	"github.com/gorilla/mux"
	"neoproj/djforgo/config"
	"net/http"
)

type Server struct {
}

func NewServer() *Server {
	return &Server{}
}

var ServerInstance = NewServer()

func newRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc

		router.Methods(route.Method1, route.Method2).Path(route.Pattern).Name(route.Name).Handler(handler)
	}

	return router
}

func (this *Server) OnInit() error {

	http.Handle("/", newRouter())
	l4g.Info("http://%s:%s/\n", config.QasConfig.Downnet.HttpIP, config.QasConfig.Downnet.Port)

	return nil
}

func (this *Server) OnWork() {
	err := http.ListenAndServe(fmt.Sprintf("%s:%s", config.QasConfig.Downnet.HttpIP,
		config.QasConfig.Downnet.Port), nil)
	if err != nil {
		l4g.Error(err)
	}
}

func (this *Server) OnClose() {

}
