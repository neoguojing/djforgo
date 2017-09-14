package sessions

import (
	"djforgo/system"
	l4g "github.com/alecthomas/log4go"
	"github.com/gorilla/sessions"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

var G_SessionStore *SessionStore

type SessionStore struct {
	store sessions.Store
}

func newSessionStore() *SessionStore {
	serstr := system.QasConfig.Session.Salt + time.Now().String()
	secret, _ := bcrypt.GenerateFromPassword([]byte(serstr), 14)
	return &SessionStore{
		store: sessions.NewCookieStore(secret),
	}
}

func (this *SessionStore) GetSession(r *http.Request) *sessions.Session {
	session, err := this.store.New(r, system.QasConfig.Session.Name)
	if err != nil {
		l4g.Error(err)
		return nil
	}
	return session
}

func (this *SessionStore) SetSession(w http.ResponseWriter, r *http.Request, maxage int, key, value interface{}) {
	session, _ := this.store.Get(r, system.QasConfig.Session.Name)
	session.Values[key] = value
	session.Options.MaxAge = maxage
	session.Options.Path = system.QasConfig.Session.Path
	session.Save(r, w)
}

func InitSessionStore() {
	G_SessionStore = newSessionStore()
}
