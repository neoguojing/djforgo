package server

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

var routes = Routes{
	Route{
		Name:        "login",
		Method1:     "GET",
		Method2:     "POST",
		Pattern:     "/login",
		HandlerFunc: views.Login,
	},
	Route{
		Name:        "HttpServerV1",
		Method1:     "GET",
		Method2:     "POST",
		Pattern:     "/v1/http",
		HandlerFunc: nil,
	},
	Route{
		Name:        "HttpServerV2",
		Method1:     "GET",
		Method2:     "POST",
		Pattern:     "/v2/http",
		HandlerFunc: nil,
	},
	Route{
		Name:        "HttpServerJson",
		Method1:     "GET",
		Method2:     "POST",
		Pattern:     "/httpforjson",
		HandlerFunc: nil,
	},
}
