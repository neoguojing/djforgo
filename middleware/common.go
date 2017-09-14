package middleware

import (
	"djforgo/system"
	"github.com/gorilla/context"
	"net/http"
)

type CommonMiddleware struct {
}

func (this *CommonMiddleware) ProcessRequest(w http.ResponseWriter, r *http.Request) {

}

func (this *CommonMiddleware) ProcessResponse(w http.ResponseWriter, r *http.Request) {
	response := context.Get(r, system.RESPONSE)
	if response == nil {
		//redirect
		w.WriteHeader(http.StatusFound)
	} else {
		w.Write([]byte(response.(string)))
	}

	//clear the context
	context.Clear(r)
}
