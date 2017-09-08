package views

import (
	//l4g "github.com/alecthomas/log4go"
	"djforgo/config"
	"github.com/flosch/pongo2"
	"github.com/gorilla/context"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	//l4g.Debug(context.Get(r, SESSIONINFO))
	context.Set(r, config.SESSIONINFO, "neo")
	if r.Method != http.MethodPost {
		tmplate := pongo2.Must(pongo2.FromFile("./admin/templates/index.html"))

		tmplate.ExecuteWriter(nil, w)
		return
	}
}
