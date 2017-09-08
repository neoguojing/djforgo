package auth

import (
	//l4g "github.com/alecthomas/log4go"
	"djforgo/config"
	"github.com/gorilla/context"
	"net/http"
)

type AuthenticationMiddleware struct {
}

func (this *AuthenticationMiddleware) ProcessRequest(w http.ResponseWriter, r *http.Request) {

	username := context.Get(r, config.SESSIONINFO)
	if username == nil {
		username = ""
	}

	user, err := GetUserByUsername(username.(string))
	if err == nil {
		context.Set(r, config.USERINFO, user)
	}

}

func (this *AuthenticationMiddleware) ProcessResponse(w http.ResponseWriter, r *http.Request) {

}
