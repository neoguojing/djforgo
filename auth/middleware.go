package auth

import (
	"context"
	"net/http"
)

const (
	SESSIONINFO = "username"
	USERINFO    = "user"
)

type AuthenticationMiddleware struct {
}

func (this *AuthenticationMiddleware) ProcessRequest(w http.ResponseWriter, r *http.Request) {
	username := r.Context().Value(SESSIONINFO).(string)
	user, err := GetUserByUsername(&username)
	if err == nil {
		context.WithValue(r.Context(), USERINFO, user)
	}
}

func (this *AuthenticationMiddleware) ProcessResponse(w http.ResponseWriter, r *http.Request) {

}
