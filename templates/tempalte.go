package templates

import (
	l4g "github.com/alecthomas/log4go"
	"github.com/flosch/pongo2"
	"github.com/gorilla/context"
	"net/http"
)

const (
	RESPONCECONTENT = "RESPONCE"
)

func RenderTemplate(r *http.Request, path string, tcontext pongo2.Context) error {
	tempateStr, err := pongo2.RenderTemplateFile(path, tcontext)
	if err != nil {
		return err
	}

	context.Set(r, RESPONCECONTENT, tempateStr)

	return nil
}

func RenderAndResponse(w http.ResponseWriter, r *http.Request, path string, context pongo2.Context) error {
	tmplate := pongo2.Must(pongo2.FromFile(path))

	if err := tmplate.ExecuteWriter(context, w); err != nil {
		l4g.Error(err)
		return err
	}

	return nil
}

func RedirectTo(w http.ResponseWriter, url string) {
	w.Header().Set("Location", url)
	w.WriteHeader(http.StatusFound)
}
