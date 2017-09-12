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
	sessionStatu := context.Get(r, config.SESSIONSTATUS).(config.SessionStatus)
	if sessionStatu == config.Session_Exist {
		username := context.Get(r, config.SESSIONINFO).(string)
		user, err := GetUserByUsername(&username)
		if err == nil {
			context.Set(r, config.USERINFO, user)
		}
	} else {
		user := GetAnonymousUser()
		context.Set(r, config.USERINFO, user)
	}

}

func (this *AuthenticationMiddleware) ProcessResponse(w http.ResponseWriter, r *http.Request) {

}