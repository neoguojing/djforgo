package sessions

import (
	"djforgo/config"
	l4g "github.com/alecthomas/log4go"
	"github.com/gorilla/sessions"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

const (
	DEFALT_PATH = "/admin/"
)

var G_SessionStore *SessionStore

type SessionStore struct {
	store sessions.Store
}

func newSessionStore() *SessionStore {
	serstr := config.QasConfig.Session.Salt + time.Now().String()
	l4g.Debug(serstr)
	secret, _ := bcrypt.GenerateFromPassword([]byte(serstr), 14)
	l4g.Debug(string(secret))
	return &SessionStore{
		store: sessions.NewCookieStore(secret),
	}
}

func (this *SessionStore) GetSession(r *http.Request) *sessions.Session {
	session, _ := this.store.Get(r, config.QasConfig.Session.Name)
	return session
}

func (this *SessionStore) SetSession(w http.ResponseWriter, r *http.Request, key, value interface{}) {
	session, _ := this.store.Get(r, config.QasConfig.Session.Name)
	session.Values[key] = value
	session.Options.MaxAge = config.QasConfig.Session.MaxAge
	session.Options.Path = DEFALT_PATH
	session.Save(r, w)
}

func InitSessionStore() {
	G_SessionStore = newSessionStore()
}
