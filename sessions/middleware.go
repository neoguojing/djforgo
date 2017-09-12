package sessions

import (
	"djforgo/config"
	l4g "github.com/alecthomas/log4go"
	"github.com/gorilla/context"
	"net/http"
)

const (
	MaxAge_Delete = -1
)

type SessionMiddleware struct {
}

func (this *SessionMiddleware) ProcessRequest(w http.ResponseWriter, r *http.Request) {

	session := G_SessionStore.GetSession(r)
	if session == nil {
		context.Set(r, config.SESSIONSTATUS, config.Session_Invalid)
	} else {
		username := session.Values[config.SESSIONINFO]
		if username == nil {
			context.Set(r, config.SESSIONSTATUS, config.Session_New)
		} else {
			context.Set(r, config.SESSIONSTATUS, config.Session_Exist)
			context.Set(r, config.SESSIONINFO, session.Values[config.SESSIONINFO])
		}
	}

	//l4g.Debug("ProcessRequest %p", r, session.Values[config.SESSIONINFO], r.RequestURI)
	return
}

func (this *SessionMiddleware) ProcessResponse(w http.ResponseWriter, r *http.Request) {
	setSession := func() {
		username := context.Get(r, config.SESSIONINFO)
		if username == nil {
			l4g.Error("SessionMiddleware:ProcessResponse invalid SESSIONINFO", username)
			return
		}
		G_SessionStore.SetSession(w, r, config.QasConfig.Session.MaxAge, config.SESSIONINFO, username)
	}
	ssn_status := context.Get(r, config.SESSIONSTATUS).(config.SessionStatus)
	switch ssn_status {
	case config.Session_New:
		setSession()
	case config.Session_Invalid:
		setSession()
	case config.Session_Exist:
		return
	case config.Session_Delete:
		G_SessionStore.SetSession(w, r, MaxAge_Delete, config.SESSIONINFO, "")
	default:
		l4g.Error("SessionMiddleware:ProcessResponse invalid ssn_status", ssn_status)
	}

	//l4g.Debug("ProcessResponse %p", r, username, r.RequestURI)

	return
}
