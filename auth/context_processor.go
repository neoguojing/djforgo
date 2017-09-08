package auth

import (
	"djforgo/config"
	l4g "github.com/alecthomas/log4go"
	"github.com/flosch/pongo2"
	"github.com/gorilla/context"
	"net/http"
)

func Auth_Context(r *http.Request, tcontext pongo2.Context) pongo2.Context {
	userObj := context.Get(r, config.USERINFO)
	if userObj == nil {
		l4g.Error("Auth_Context userObj was nil")
		return nil
	}

	userContext := pongo2.Context{config.USERINFO: userObj}
	if tcontext == nil {
		return userContext
	}
	return tcontext.Update(userContext)
}
