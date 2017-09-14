package main

import (
	"djforgo/auth"
	"djforgo/dao"
	"djforgo/system"
	"djforgo/system/register"
	"flag"
	"fmt"
)

func main() {
	//命令行解析
	init := flag.Bool("init", false, "do admin init")
	appcfgfile := flag.String("f", "../config.json", "config file path")
	help := flag.Bool("help", false, "help")
	flag.Parse()

	system.LoadConfig(appcfgfile)
	fmt.Println(system.QasConfig)

	switch {
	case *help:
		flag.PrintDefaults()
	case *init:
		register.ModelSetInstance.CreateTables()
		register.ModelSetInstance.CreateContentType()
		register.ModelSetInstance.CreatePermissions()
		createAdmin()

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

	user.SetEmail(system.QasConfig.Admin.Email)
	user.SetUserName("admin")
	user.SetPassword("admin")
	if err = user.CreateAdminUser(&user); err != nil {
		fmt.Println(err)
	}

}
