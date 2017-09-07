package urls

import (
	"djforgo/admin/views"
	"net/http"
)

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
		HandlerFunc: views.Login,
	},
	Route{
		Name:        "register",
		Method1:     "GET",
		Method2:     "POST",
		Pattern:     "/register",
		HandlerFunc: views.Register,
	},
}
