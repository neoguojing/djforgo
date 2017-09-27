package oauth2

import (
	"djforgo/auth"
	"djforgo/dao"
	"djforgo/templates"
	"github.com/RangelReale/osin"
	l4g "github.com/alecthomas/log4go"
	"github.com/flosch/pongo2"
	"github.com/gorilla/schema"
	"net/http"
)

var decoder = schema.NewDecoder()

func handleLogin(ar *osin.AuthorizeRequest, w http.ResponseWriter, r *http.Request) bool {
	relogin := func() {

		tContext := pongo2.Context{"query": r.URL.RawQuery}
		templates.RenderAndResponse(w, r, "./contrib/oauth2/templates/login.html", tContext)
	}

	if r.Method != http.MethodPost {
		relogin()
		return false
	}

	err := r.ParseForm()
	if err != nil {
		l4g.Error(err)
		relogin()
		return false
	}

	var userFrom auth.UserLoginForm
	err = decoder.Decode(&userFrom, r.PostForm)
	if err != nil {
		l4g.Error(err)
		relogin()
		return false
	}

	if userFrom.Valid() != nil {
		relogin()
		return false
	}

	_, err = auth.Login_Check(&userFrom)
	if err != nil {
		relogin()
		return false
	}

	return true
}

func AuthorizeHandler(w http.ResponseWriter, r *http.Request) {
	resp := dao.G_OauthServer.NewResponse()
	defer resp.Close()

	if ar := dao.G_OauthServer.HandleAuthorizeRequest(resp, r); ar != nil {
		//handle login
		if !handleLogin(ar, w, r) {
			return
		}

		ar.Authorized = true
		dao.G_OauthServer.FinishAuthorizeRequest(resp, r, ar)
	}

	osin.OutputJSON(resp, w, r)
}

//Token request must be post
func TokenHandler(w http.ResponseWriter, r *http.Request) {
	resp := dao.G_OauthServer.NewResponse()
	defer resp.Close()
	r.ParseForm()
	l4g.Debug(r.Form)
	if ar := dao.G_OauthServer.HandleAccessRequest(resp, r); ar != nil {
		ar.Authorized = true
		dao.G_OauthServer.FinishAccessRequest(resp, r, ar)
	}

	osin.OutputJSON(resp, w, r)
}
