package views

import (
	//l4g "github.com/alecthomas/log4go"
	//"djforgo/config"
	//"github.com/gorilla/context"
	"djforgo/auth"
	"djforgo/templates"
	"github.com/flosch/pongo2"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	//l4g.Debug(context.Get(r, SESSIONINFO))
	//context.Set(r, config.SESSIONINFO, "neo")

	if !auth.IsAuthticated(r) {
		templates.RedirectTo(w, "/login")
		return
	}

	if r.Method != http.MethodPost {
		ctx := pongo2.Context{"users": auth.GetUsers(r)}
		templates.RenderTemplate(r, "./admin/templates/index.html", auth.Auth_Context(r, ctx))
		return
	}
}
