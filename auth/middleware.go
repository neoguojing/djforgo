package auth

import (
	//l4g "github.com/alecthomas/log4go"
	"djforgo/system"
	"github.com/gorilla/context"
	"net/http"
)

type AuthenticationMiddleware struct {
}

func (this *AuthenticationMiddleware) ProcessRequest(w http.ResponseWriter, r *http.Request) {
	sessionStatu := context.Get(r, system.SESSIONSTATUS).(system.SessionStatus)
	if sessionStatu == system.Session_Exist {
		username := context.Get(r, system.SESSIONINFO).(string)
		user, err := GetUserByUsername(&username)
		if err == nil {
			context.Set(r, system.USERINFO, user)
		}
	} else {
		user := GetAnonymousUser()
		context.Set(r, system.USERINFO, user)
	}

}

func (this *AuthenticationMiddleware) ProcessResponse(w http.ResponseWriter, r *http.Request) {

}
