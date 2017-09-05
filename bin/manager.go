package main

import (
	"djforgo/auth"
	"djforgo/config"
	"djforgo/dao"
	"flag"
	"fmt"
)

func main() {

	//命令行解析
	admin := flag.Bool("admin", false, "do admin init")
	appcfgfile := flag.String("f", "../config.json", "config file path")
	help := flag.Bool("help", false, "help")
	flag.Parse()

	config.LoadConfig(appcfgfile)
	fmt.Println(config.QasConfig)

	switch {
	case *help:
		flag.PrintDefaults()
	case *admin:
		adminInit()
	}

}

func adminInit() {
	err := dao.DB_Init()
	if err != nil {
		return
	}
	defer dao.DB_Destroy()

	if dao.DB_Instance.HasTable(&auth.User{}) {
		err = dao.DB_Instance.DropTable(&auth.User{}).Error
		if err != nil {
			fmt.Println("drop users", err)
			return
		}
	}

	err = dao.DB_Instance.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&auth.User{}).Error
	if err != nil {
		fmt.Println("create users", err)
		return
	}
}
