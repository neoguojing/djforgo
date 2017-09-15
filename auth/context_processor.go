package auth

import (
	"djforgo/system"
	l4g "github.com/alecthomas/log4go"
	"github.com/flosch/pongo2"
	"github.com/gorilla/context"
	"net/http"
)

func PermissionWrapper(user IUser) pongo2.Context {
	perms, err := user.GetAllPermissions()
	if err != nil {
		l4g.Error(err)
	}

	permMap := make(map[string][]string)
	for _, perm := range perms {
		app, op := perm.Wrapper()
		if app == "" || op == "" {
			continue
		}

		if _, ok := permMap[app]; !ok {
			ops := make([]string, 0)
			permMap[app] = ops
		}

		permMap[app] = append(permMap[app], op)
	}
	l4g.Debug(permMap)
	return pongo2.Context{system.PERMITIONINFO: permMap}
}

func Auth_Context(r *http.Request, tcontext pongo2.Context) pongo2.Context {
	userObj := context.Get(r, system.USERINFO)
	if userObj == nil {
		l4g.Error("Auth_Context userObj was nil")
		return nil
	}

	userContext := pongo2.Context{system.USERINFO: userObj}
	userContext.Update(PermissionWrapper(userObj.(IUser)))
	if tcontext == nil {
		return userContext
	}
	return tcontext.Update(userContext)
}
