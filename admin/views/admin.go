package views

import (
	"djforgo/admin"
	"djforgo/auth"
	"djforgo/system"
	"djforgo/templates"
	l4g "github.com/alecthomas/log4go"
	"github.com/flosch/pongo2"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"net/http"
	"strconv"
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

	if r.Method != http.MethodPost {
		parseEditParam(w, r)
		return
	}
}

func parseEditParam(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	model := vars["model"]
	id := vars["id"]

	if id == "" {
		l4g.Error("parseEditParam: invalid id param")
		return
	}

	ctxUser := context.Get(r, system.USERINFO)
	if ctxUser == nil {
		l4g.Error("parseEditParam: ctxUser was nil")
		return
	}
	currentUser := ctxUser.(*auth.User)
	ctx := pongo2.Context{}

	switch model {
	case "user":
		reqId, err := strconv.Atoi(id)
		if err != nil {
			l4g.Error("parseEditParam:%v", err)
			return
		}

		var targetUser auth.IUser

		if currentUser.ID == uint(reqId) {
			targetUser = currentUser
		} else {
			targetUser, err = auth.GetUserByID(uint(reqId))
			if err != nil {
				return
			}
		}

		permissions, _ := targetUser.GetAllPermissions()
		groups, _ := targetUser.GetAllGroups()
		ctx.Update(pongo2.Context{"targetuser": targetUser})
		ctx.Update(pongo2.Context{"groups": groups})
		ctx.Update(pongo2.Context{"permissions": permissions})

		if currentUser.IsAdmin() && currentUser.ID != uint(reqId) {
			unUsedPermitions, err := targetUser.GetUnUsedPermitions()
			if err == nil && unUsedPermitions != nil {
				ctx.Update(pongo2.Context{"permissions_without": unUsedPermitions})
			}

			unUsedGroups, err := targetUser.GetUnUsedGroups()
			if err == nil && unUsedGroups != nil {
				ctx.Update(pongo2.Context{"groups_without": unUsedGroups})
			}
		}

	case "perm":
		currentUser.GetAllPermissions()
		ctx.Update(pongo2.Context{"targetuser": currentUser})
	case "group":
		currentUser.GetAllGroups()
		ctx.Update(pongo2.Context{"targetuser": currentUser})
	default:
		l4g.Error("parseEditParams: invalid model")
		return
	}

	templates.RenderTemplate(r, "./admin/templates/edit.html", auth.Auth_Context(r, ctx))
}

func parseEditForms(model string, w http.ResponseWriter, r *http.Request) {
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
		l4g.Error("parseEditForms: invalid model")
		return
	}

}
