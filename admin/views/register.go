package views

import (
	"djforgo/auth"
	l4g "github.com/alecthomas/log4go"
	"html/template"
	"net/http"
)

func Register(w http.ResponseWriter, req *http.Request) {

	if req.Method != http.MethodPost {
		html, err := template.ParseFiles("./admin/templates/register.html")
		if err != nil {
			l4g.Error("Register parse html faild", err)
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

	var registerForm auth.UserCreationForm
	err = decoder.Decode(&registerForm, req.PostForm)
	if err != nil {
		l4g.Error(err)
		return
	}

	if registerForm.Valid() != nil {
		http.Redirect(w, req, "/register", http.StatusFound)
		return
	}

	err = createUser(&registerForm)
	if err != nil {
		l4g.Error("Register", err)
		http.Redirect(w, req, "/register", http.StatusFound)
		return
	}

	http.Redirect(w, req, "/login", http.StatusFound)
	l4g.Debug(registerForm)
}

func createUser(rguser *auth.UserCreationForm) error {
	var user auth.User
	user.SetUserName(rguser.Name)
	user.SetEmail(rguser.Email)
	user.SetPassword(rguser.Password1)

	return user.CreateUser(&user)
}
