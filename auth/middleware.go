package auth

import (
	//l4g "github.com/alecthomas/log4go"
	"github.com/gorilla/context"
	"net/http"
)

const (
	SESSIONINFO = "username"
	USERINFO    = "user"
)

type AuthenticationMiddleware struct {
}

func (this *AuthenticationMiddleware) ProcessRequest(w http.ResponseWriter, r *http.Request) {

	username := context.Get(r, SESSIONINFO)
	if username == nil {
		username = ""
	}

	user, err := GetUserByUsername(username.(string))
	if err == nil {
		context.Set(r, USERINFO, user)
	}

	//l4g.Debug("%v", username.(string), "AuthenticationMiddleware", r.RequestURI)
}

func (this *AuthenticationMiddleware) ProcessResponse(w http.ResponseWriter, r *http.Request) {

}
