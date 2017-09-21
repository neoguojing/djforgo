package views

import (
	"djforgo/auth"
	"djforgo/system"
	"djforgo/templates"
	l4g "github.com/alecthomas/log4go"
	"github.com/gorilla/context"
	"github.com/gorilla/schema"
	"net/http"
)

var decoder = schema.NewDecoder()

func Login(w http.ResponseWriter, req *http.Request) {
	sessionStatu := context.Get(req, system.SESSIONSTATUS).(system.SessionStatus)
	if sessionStatu == system.Session_Exist {
		session_user := context.Get(req, system.USERINFO)
		if session_user != nil {
			if session_user.(auth.IUser).IsAuthenticated() {
				templates.RedirectTo(w, "/index")
				return
			}
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

	context.Set(req, system.SESSIONINFO, user.GetUserID())

	templates.RedirectTo(w, "/index")

}

func Logout(w http.ResponseWriter, req *http.Request) {
	sessionStatu := context.Get(req, system.SESSIONSTATUS).(system.SessionStatus)
	if sessionStatu != system.Session_Exist {
		templates.RedirectTo(w, "/login")
		return
	}

	if req.Method != http.MethodPost {
		context.Set(req, system.SESSIONSTATUS, system.Session_Delete)
		templates.RedirectTo(w, "/login")
		return
	}

	return
}
