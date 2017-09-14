package sessions

import (
	"djforgo/system"
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
		context.Set(r, system.SESSIONSTATUS, system.Session_Invalid)
	} else {
		username := session.Values[system.SESSIONINFO]
		if username == nil {
			context.Set(r, system.SESSIONSTATUS, system.Session_New)
		} else {
			context.Set(r, system.SESSIONSTATUS, system.Session_Exist)
			context.Set(r, system.SESSIONINFO, session.Values[system.SESSIONINFO])
		}
	}

	//l4g.Debug("ProcessRequest %p", r, session.Values[config.SESSIONINFO], r.RequestURI)
	return
}

func (this *SessionMiddleware) ProcessResponse(w http.ResponseWriter, r *http.Request) {
	setSession := func() {
		username := context.Get(r, system.SESSIONINFO)
		if username == nil {
			l4g.Error("SessionMiddleware:ProcessResponse invalid SESSIONINFO", username)
			return
		}
		G_SessionStore.SetSession(w, r, system.QasConfig.Session.MaxAge, system.SESSIONINFO, username)
	}
	ssn_status := context.Get(r, system.SESSIONSTATUS).(system.SessionStatus)
	switch ssn_status {
	case system.Session_New:
		setSession()
	case system.Session_Invalid:
		setSession()
	case system.Session_Exist:
		return
	case system.Session_Delete:
		G_SessionStore.SetSession(w, r, MaxAge_Delete, system.SESSIONINFO, "")
	default:
		l4g.Error("SessionMiddleware:ProcessResponse invalid ssn_status", ssn_status)
	}

	//l4g.Debug("ProcessResponse %p", r, username, r.RequestURI)

	return
}
