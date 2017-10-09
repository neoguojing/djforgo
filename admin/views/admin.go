package views

import (
	"djforgo/admin"
	"djforgo/auth"
	"djforgo/dao"
	"djforgo/forms"
	"djforgo/system"
	"djforgo/templates"
	"djforgo/utils"
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

	parseEditForms(w, r)
	return
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
	case "group":
	default:
		l4g.Error("parseEditParams: invalid model")
		return
	}

	templates.RenderTemplate(r, "./admin/templates/edit.html", auth.Auth_Context(r, ctx))
}

func parseEditForms(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	model := vars["model"]
	id := vars["id"]

	if id == "" {
		l4g.Error("parseEditParam: invalid id param")
		return
	}

	switch model {
	case "user":
		err := r.ParseForm()
		if err != nil {
			l4g.Error("ParseForm,%v", err)
			return
		}

		var userFrom admin.UserEditForm

		err = decoder.Decode(&userFrom, r.PostForm)
		l4g.Debug("***", userFrom, r.PostForm)
		if err != nil {
			l4g.Error("Decode ,%v", err)
			return
		}

		if userFrom.Valid() != nil {
			return
		}

		userFrom.Save()
		templates.RedirectTo(w, "/edit/user/"+id)

	case "permition":
	case "group":
	default:
		l4g.Error("parseEditForms: invalid model")
		return
	}

}

func ModelEditHandler(w http.ResponseWriter, r *http.Request) {
	if !auth.IsAuthticated(r) {
		templates.RedirectTo(w, "/login")
		return
	}

	gformHandler(r)

}

func gformHandler(r *http.Request) {
	vars := mux.Vars(r)
	modelName := vars["modelname"]
	id := vars["id"]
	if id == "" {
		l4g.Error("ModelEditHandler: invalid id param")
		return
	}

	forms.Init()

	obj := utils.G_ObjRegisterStore.New(modelName)
	if obj == nil {
		panic(1)
		return
	}

	form := obj.(utils.IForm)
	form.Init(r)

	if r.Method != http.MethodPost {
		tContext := pongo2.Context{"model_id": id, "model_name": modelName, "fields": form.Fields()}
		templates.RenderTemplate(r, "./admin/templates/model_edit.html", auth.Auth_Context(r, tContext))
		return
	}

	if !form.IsValid() {
		l4g.Debug("form is invalid")
		return
	}

	model := form.GetModel()
	l4g.Debug(model)

	if model != nil {
		db := dao.NewManager().Update(model)
		if nil != db.Error {
			l4g.Error(db.Error)
			return
		}

	} else {
		panic("model was nil ")
	}
}

func DelHandler(w http.ResponseWriter, r *http.Request) {

	if !auth.IsAuthticated(r) {
		templates.RedirectTo(w, "/login")
		return
	}

	if r.Method != http.MethodPost {
		vars := mux.Vars(r)
		model := vars["model"]
		id := vars["id"]

		if id == "" {
			l4g.Error("DelHandler: invalid id param")
			return
		}

		//l4g.Debug("DelHandler", model, id)

		ctxUser := context.Get(r, system.USERINFO)
		if ctxUser == nil {
			l4g.Error("DelHandler: ctxUser was nil")
			return
		}
		currentUser := ctxUser.(*auth.User)
		ctx := pongo2.Context{}

		switch model {
		case "user":
			reqId, err := strconv.Atoi(id)
			if err != nil {
				l4g.Error("DelHandler:%v", err)
				return
			}

			var targetUser auth.IUser

			if currentUser.ID == uint(reqId) {
				templates.RedirectTo(w, "/index")
				return
			} else {
				targetUser, err = auth.GetUserByID(uint(reqId))
				if err != nil {
					return
				}
			}

			targetUser.GetAllGroups()
			targetUser.GetAllPermissions()

			err = targetUser.(*auth.User).Delete(targetUser.(*auth.User)).Error
			if err != nil {
				l4g.Error("DelHandler:%v", err)
			}
			templates.RedirectTo(w, "/index")
			return

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

		templates.RenderTemplate(r, "./admin/templates/index.html", auth.Auth_Context(r, ctx))
		return
	}

	return
}

func PasswordResetHandler(w http.ResponseWriter, r *http.Request) {
	if !auth.IsAuthticated(r) {
		templates.RedirectTo(w, "/login")
		return
	}

	if r.Method != http.MethodPost {
		templates.RenderTemplate(r, "./admin/templates/password_reset.html", auth.Auth_Context(r, nil))
		return
	}

	err := r.ParseForm()
	if err != nil {
		l4g.Error("PasswordResetHandler,%v", err)
		return
	}

	var passwordResetFrom admin.PasswordResetForm
	err = decoder.Decode(&passwordResetFrom, r.PostForm)
	if err != nil {
		l4g.Error("PasswordResetHandler ,%v", err)
		return
	}

	ctxUser := context.Get(r, system.USERINFO).(*auth.User)
	if passwordResetFrom.Valid(ctxUser) != nil {
		templates.RenderTemplate(r, "./admin/templates/password_reset.html", auth.Auth_Context(r, nil))
		return
	}

	err = passwordResetFrom.Save(ctxUser)
	if err != nil {
		templates.RenderTemplate(r, "./admin/templates/password_reset.html", auth.Auth_Context(r, nil))
	}

	templates.RedirectTo(w, "/index")
}
