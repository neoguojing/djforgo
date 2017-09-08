package views

import (
	"djforgo/auth"
	l4g "github.com/alecthomas/log4go"
	"github.com/flosch/pongo2"
	"github.com/gorilla/context"
	"github.com/gorilla/schema"
	"net/http"
)

const (
	USERINFO    = "user"
	SESSIONINFO = "username"
)

var decoder = schema.NewDecoder()

func Login(w http.ResponseWriter, req *http.Request) {
	session_user := context.Get(req, USERINFO)
	if session_user != nil {
		if session_user.(auth.IUser).IsAuthenticated() {
			http.Redirect(w, req, "/index", http.StatusFound)
			return
		}
	}

	if req.Method != http.MethodPost {
		tmplate := pongo2.Must(pongo2.FromFile("./auth/templates/login.html"))

		tmplate.ExecuteWriter(nil, w)
		return
	}

	err := req.ParseForm()
	if err != nil {
		l4g.Error(err)
		return
	}

	var userFrom auth.UserLoginForm
	err = decoder.Decode(&userFrom, req.PostForm)
	if err != nil {
		l4g.Error(err)
		return
	}

	if userFrom.Valid() != nil {
		http.Redirect(w, req, "/login", http.StatusFound)
		return
	}

	var user auth.IUser
	user, err = auth.Login_Check(&userFrom)
	if err != nil {
		http.Redirect(w, req, "/login", http.StatusFound)
		return
	}

	context.Set(req, SESSIONINFO, user.GetUserName())
	//l4g.Debug("%p,%v", req, userFrom)

	//http.Redirect(w, req, "/index", http.StatusFound)

}
