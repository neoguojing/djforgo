package main

import (
	"djforgo/auth"
	"djforgo/dao"
	"djforgo/system"
	"djforgo/system/register"
	"flag"
	"fmt"
	"github.com/felipeweb/osin-mysql"
)

func main() {
	//命令行解析
	init := flag.Bool("init", false, "do admin init")
	appcfgfile := flag.String("f", "../config.json", "config file path")
	help := flag.Bool("help", false, "help")
	flag.Parse()

	system.LoadConfig(appcfgfile)
	fmt.Println(system.SysConfig)

	switch {
	case *help:
		flag.PrintDefaults()
	case *init:
		register.ModelSetInstance.CreateTables()
		register.ModelSetInstance.CreateContentType()
		register.ModelSetInstance.CreatePermissions()
		createAdmin()
		createOAuthStore()
	}

}

func createAdmin() {
	err := dao.DB_Init()
	if err != nil {
		return
	}
	defer dao.DB_Destroy()

	var user auth.User

	if !dao.DB_Instance.HasTable(user) {
		fmt.Println("users table does not exist")
		return
	}

	user.SetEmail(system.SysConfig.Admin.Email)
	user.SetUserName("admin")
	user.SetPassword("admin")
	if err = user.CreateAdminUser(&user); err != nil {
		fmt.Println("CreateAdminUser", err)
	}

}

func createOAuthStore() {
	if system.SysConfig.Services.OAuth != 1 {
		return
	}

	err := dao.DB_Init()
	if err != nil {
		return
	}
	defer dao.DB_Destroy()

	dao.G_OAuthStore = mysql.New(dao.DB_Instance.DB(), "oauth_")
	dao.G_OAuthStore.CreateSchemas()
}
