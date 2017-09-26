package urls

import (
	admin "djforgo/admin/views"
	"djforgo/auth/views"
	"djforgo/middleware"
	"github.com/gorilla/mux"
	"net/http"
)

var (
	G_Router *mux.Router
)

func RegisterRouters(routes Routes) {
	for _, route := range routes {
		G_Router.Methods(route.Method1, route.Method2).Path(route.Pattern).
			Name(route.Name).HandlerFunc(route.HandlerFunc)
	}
}

type Route struct {
	Name        string
	Method1     string
	Method2     string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var Urls = Routes{
	Route{
		Name:        "login",
		Method1:     "GET",
		Method2:     "POST",
		Pattern:     "/login",
		HandlerFunc: middleware.MiddlewareHandlerFunc(views.Login),
	},
	Route{
		Name:        "logout",
		Method1:     "GET",
		Method2:     "POST",
		Pattern:     "/logout",
		HandlerFunc: middleware.MiddlewareHandlerFunc(views.Logout),
	},
	Route{
		Name:        "register",
		Method1:     "GET",
		Method2:     "POST",
		Pattern:     "/register",
		HandlerFunc: views.Register,
	},
	Route{
		Name:        "index",
		Method1:     "GET",
		Method2:     "POST",
		Pattern:     "/index",
		HandlerFunc: middleware.MiddlewareHandlerFunc(admin.IndexHandler),
	},
	Route{
		Name:        "edit",
		Method1:     "GET",
		Method2:     "POST",
		Pattern:     "/edit/{model:[a-z]+}/{id:[0-9]+}",
		HandlerFunc: middleware.MiddlewareHandlerFunc(admin.EditHandler),
	},
	Route{
		Name:        "del",
		Method1:     "GET",
		Method2:     "POST",
		Pattern:     "/del/{model:[a-z]+}/{id:[0-9]+}",
		HandlerFunc: middleware.MiddlewareHandlerFunc(admin.DelHandler),
	},
	Route{
		Name:        "password_reset",
		Method1:     "GET",
		Method2:     "POST",
		Pattern:     "/reset",
		HandlerFunc: middleware.MiddlewareHandlerFunc(admin.PasswordResetHandler),
	},
}

func newRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range Urls {
		router.Methods(route.Method1, route.Method2).Path(route.Pattern).Name(route.Name).HandlerFunc(route.HandlerFunc)
	}

	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	return router
}

func init() {
	G_Router = newRouter()
}
