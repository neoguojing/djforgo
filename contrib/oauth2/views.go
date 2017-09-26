package oauth2

import (
	"djforgo/dao"
	"github.com/RangelReale/osin"
	"net/http"
)

var oauthServer = osin.NewServer(osin.NewServerConfig(), dao.G_OAuthStore)

func AuthorizeHandler(w http.ResponseWriter, r *http.Request) {
	resp := oauthServer.NewResponse()
	defer resp.Close()

	if ar := oauthServer.HandleAuthorizeRequest(resp, r); ar != nil {
		//handle login

		ar.Authorized = true
		oauthServer.FinishAuthorizeRequest(resp, r, ar)
	}

	osin.OutputJSON(resp, w, r)
}

func TokenHandler(w http.ResponseWriter, r *http.Request) {
	resp := oauthServer.NewResponse()
	defer resp.Close()

	if ar := oauthServer.HandleAccessRequest(resp, r); ar != nil {
		ar.Authorized = true
		oauthServer.FinishAccessRequest(resp, r, ar)
	}

	osin.OutputJSON(resp, w, r)
}
