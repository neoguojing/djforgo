package views

import (
	"djforgo/auth"
	l4g "github.com/alecthomas/log4go"
	"github.com/gorilla/schema"
	"html/template"
	"net/http"
)

var decoder = schema.NewDecoder()

func Login(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		html, err := template.ParseFiles("./admin/templates/login.html")
		if err != nil {
			l4g.Error("Login parse html faild", err)
			return
		}

		html.Execute(w, nil)
		return
	}

	err := req.ParseForm()
	if err != nil {
		l4g.Error(err)
		return
	}

	var user auth.UserLoginForm
	err = decoder.Decode(&user, req.PostForm)
	if err != nil {
		l4g.Error(err)
		return
	}

	if user.Valid() != nil {
		http.Redirect(w, req, "/login", http.StatusFound)
		return
	}

	if auth.Login_Check(&user) != nil {
		http.Redirect(w, req, "/login", http.StatusFound)
		return
	} else {

	}

	l4g.Debug(user)
}
