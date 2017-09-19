package views

import (
	l4g "github.com/alecthomas/log4go"
	//"djforgo/config"
	//"github.com/gorilla/context"
	"djforgo/admin"
	"djforgo/auth"
	"djforgo/templates"
	"github.com/flosch/pongo2"
	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"net/http"
)

var decoder = schema.NewDecoder()

func IndexHandler(w http.ResponseWriter, r *http.Request) {

	if !auth.IsAuthticated(r) {
		templates.RedirectTo(w, "/login")
		return
	}

	if r.Method != http.MethodPost {
		ctx := pongo2.Context{"users": auth.GetUsers(r)}
		ctx.Update(pongo2.Context{"permitions": auth.GetAllPermitionsOfUser(r)})
		ctx.Update(pongo2.Context{"groups": auth.GetAllGroupsOfUser(r)})
		templates.RenderTemplate(r, "./admin/templates/index.html", auth.Auth_Context(r, ctx))
		return
	}
}

func EditHandler(w http.ResponseWriter, r *http.Request) {

	if !auth.IsAuthticated(r) {
		templates.RedirectTo(w, "/login")
		return
	}

	vars := mux.Vars(r)
	model := vars["model"]
	id := vars["id"]
	
	l4g.Debug(model,id)
	_ = model
	_ = id

	if r.Method != http.MethodPost {
		ctx := pongo2.Context{"permitions": auth.GetAllPermitionsOfUser(r)}
		ctx.Update(pongo2.Context{"groups": auth.GetAllGroupsOfUser(r)})
		templates.RenderTemplate(r, "./admin/templates/edit.html", auth.Auth_Context(r, ctx))
		return
	}
}

func parseEditParams(model string, w http.ResponseWriter, r *http.Request) {
	switch model {
	case "user":
		err := r.ParseForm()
		if err != nil {
			l4g.Error(err)
			return
		}
		var userFrom admin.UserEditForm
		err = decoder.Decode(&userFrom, r.PostForm)
		if err != nil {
			l4g.Error(err)
			return
		}

		if userFrom.Valid() != nil {
			templates.RedirectTo(w, r.RequestURI)
			return
		}

	case "permition":
	case "group":
	default:
		l4g.Error("parseEditParams: invalid model")
		return
	}

}
