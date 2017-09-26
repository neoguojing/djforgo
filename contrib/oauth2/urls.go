package oauth2

import (
	"djforgo/system"
	"djforgo/urls"
)

var authUrl = urls.Routes{
	urls.Route{
		Name:        "authorize",
		Method1:     "GET",
		Method2:     "",
		Pattern:     "/authorize",
		HandlerFunc: AuthorizeHandler,
	},
	urls.Route{
		Name:        "token",
		Method1:     "GET",
		Method2:     "",
		Pattern:     "/token",
		HandlerFunc: TokenHandler,
	},
}

func init() {
	if system.SysConfig.Services.OAuth == 1 {
		urls.RegisterRouters(authUrl)
	}
}
