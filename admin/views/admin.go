package views

import (
	//l4g "github.com/alecthomas/log4go"
	"github.com/flosch/pongo2"
	"github.com/gorilla/context"
	"net/http"
)

const (
	SESSIONINFO = "username"
	USERINFO    = "user"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	//l4g.Debug(context.Get(r, SESSIONINFO))
	context.Set(r, SESSIONINFO, "neo")
	if r.Method != http.MethodPost {
		tmplate := pongo2.Must(pongo2.FromFile("./admin/templates/index.html"))

		tmplate.ExecuteWriter(nil, w)
		return
	}
}
