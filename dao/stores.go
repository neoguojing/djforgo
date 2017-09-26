package dao

import (
	"djforgo/system"
	"github.com/felipeweb/osin-mysql"
)

var (
	G_OAuthStore *mysql.Storage
)

func InitSrores() {
	if system.SysConfig.Services.OAuth == 1 {
		G_OAuthStore = mysql.New(DB_Instance.DB(), "oauth_")
	}
}
