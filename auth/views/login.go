package views

import (
	"djforgo/auth"
	"djforgo/config"
	"djforgo/templates"
	l4g "github.com/alecthomas/log4go"
	"github.com/gorilla/context"
	"github.com/gorilla/schema"
	"net/http"
)

var decoder = schema.NewDecoder()

func Login(w http.ResponseWriter, req *http.Request) {
	sessionStatu := context.Get(req, config.SESSIONSTATUS).(config.SessionStatus)
	if sessionStatu == config.Session_Exist {
		session_user := context.Get(req, config.USERINFO).(auth.IUser)
		if session_user.IsAuthenticated() {
			templates.RedirectTo(w, "/index")
			return
		}
	}

	if req.Method != http.MethodPost {

		templates.RenderTemplate(req, "./auth/templates/login.html", nil)

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
		templates.RedirectTo(w, "/login")
		return
	}

	var user auth.IUser
	user, err = auth.Login_Check(&userFrom)
	if err != nil {
		templates.RedirectTo(w, "/login")
		return
	}

	context.Set(req, config.SESSIONINFO, user.GetUserName())

	templates.RedirectTo(w, "/index")

}

func Logout(w http.ResponseWriter, req *http.Request) {
	sessionStatu := context.Get(req, config.SESSIONSTATUS).(config.SessionStatus)
	if sessionStatu != config.Session_Exist {
		templates.RedirectTo(w, "/login")
		return
	}

	if req.Method != http.MethodPost {
		context.Set(req, config.SESSIONSTATUS, config.Session_Delete)
		templates.RedirectTo(w, "/login")
		return
	}

	return
}
