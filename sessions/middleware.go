package sessions

import (
	//l4g "github.com/alecthomas/log4go"
	"djforgo/config"
	"github.com/gorilla/context"
	"net/http"
)

type SessionMiddleware struct {
}

func (this *SessionMiddleware) ProcessRequest(w http.ResponseWriter, r *http.Request) {

	session := G_SessionStore.GetSession(r)
	if session == nil {
		return
	}

	context.Set(r, config.SESSIONINFO, session.Values[config.SESSIONINFO])
	return
}

func (this *SessionMiddleware) ProcessResponse(w http.ResponseWriter, r *http.Request) {
	username := context.Get(r, config.SESSIONINFO)
	if username == nil {
		return
	}

	//l4g.Debug("%p-%v", r, username, "ProcessResponse", r.RequestURI)
	G_SessionStore.SetSession(w, r, config.SESSIONINFO, username)

	return
}
