package middleware

import (
	"djforgo/config"
	"github.com/gorilla/context"
	"net/http"
)

type CommonMiddleware struct {
}

func (this *CommonMiddleware) ProcessRequest(w http.ResponseWriter, r *http.Request) {

}

func (this *CommonMiddleware) ProcessResponse(w http.ResponseWriter, r *http.Request) {
	response := context.Get(r, config.RESPONSE)
	if response == nil {
		w.Write(nil)
	} else {
		w.Write([]byte(response.(string)))
	}

	//clear the context
	context.Clear(r)
}
