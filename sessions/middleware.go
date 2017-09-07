package sessions

import (
	"context"
	l4g "github.com/alecthomas/log4go"
	"net/http"
)

const (
	SESSIONINFO = "username"
)

type SessionMiddleware struct {
}

func (this *SessionMiddleware) ProcessRequest(w http.ResponseWriter, r *http.Request) {

	session := G_SessionStore.GetSession(r)
	sessioninfo, ok := session.Values[SESSIONINFO]
	if !ok {
		l4g.Error("no session avalable")
		return
	}

	context.WithValue(r.Context(), SESSIONINFO, sessioninfo)
	return
}

func (this *SessionMiddleware) ProcessResponse(w http.ResponseWriter, r *http.Request) {
	username := r.Context().Value(SESSIONINFO)
	if username == nil {
		l4g.Error("no session info avaliable")
		return
	}

	G_SessionStore.SetSession(w, r, SESSIONINFO, username)
	return
}
