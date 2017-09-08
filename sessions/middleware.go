package sessions

import (
	"djforgo/config"
	//l4g "github.com/alecthomas/log4go"
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

	//l4g.Debug("ProcessRequest %p", r, session.Values[config.SESSIONINFO], r.RequestURI)
	return
}

func (this *SessionMiddleware) ProcessResponse(w http.ResponseWriter, r *http.Request) {
	username := context.Get(r, config.SESSIONINFO)
	if username == nil {
		return
	}

	G_SessionStore.SetSession(w, r, config.SESSIONINFO, username)
	//l4g.Debug("ProcessResponse %p", r, username, r.RequestURI)

	return
}
