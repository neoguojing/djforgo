package dao

import (
	"djforgo/system"
	"github.com/RangelReale/osin"
	"github.com/felipeweb/osin-mysql"
)

var (
	G_OAuthStore  *mysql.Storage
	G_OauthServer *osin.Server
)

func InitSrores() {
	if system.SysConfig.Services.OAuth == 1 {
		G_OAuthStore = mysql.New(DB_Instance.DB(), "oauth_")
		serverConfig := osin.NewServerConfig()
		serverConfig.AllowGetAccessRequest = true
		G_OauthServer = osin.NewServer(serverConfig, G_OAuthStore)
	}
}
