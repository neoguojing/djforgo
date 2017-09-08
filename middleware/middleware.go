package middleware

import (
	"djforgo/auth"
	"djforgo/sessions"
	//l4g "github.com/alecthomas/log4go"
	"net/http"
)

type IMiddleware interface {
	ProcessRequest(http.ResponseWriter, *http.Request)
	ProcessResponse(http.ResponseWriter, *http.Request)
}

var Middlewares []IMiddleware

func MiddlewareHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for _, m := range Middlewares {
			m.ProcessRequest(w, r)
		}

		next.ServeHTTP(w, r)

		for i := len(Middlewares) - 1; i >= 0; i-- {
			Middlewares[i].ProcessResponse(w, r)
		}
	})
}

func MiddlewareHandlerFunc(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for _, m := range Middlewares {
			m.ProcessRequest(w, r)
		}

		next(w, r)

		for i := len(Middlewares) - 1; i >= 0; i-- {
			Middlewares[i].ProcessResponse(w, r)
		}
	})
}

func init() {
	Middlewares = make([]IMiddleware, 0)
	Middlewares = append(Middlewares, &CommonMiddleware{})
	Middlewares = append(Middlewares, &sessions.SessionMiddleware{})
	Middlewares = append(Middlewares, &auth.AuthenticationMiddleware{})
}
