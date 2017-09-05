package auth

import (
	"github.com/jinzhu/gorm"
	"neoproj/djforgo/dao"
)

type PermissionManager struct {
	dao.Manager
}

type Permission struct {
	gorm.Model
	Name string `gorm:"size:255"`
}

type GroupManager struct {
	dao.Manager
}

type Group struct {
	gorm.Model
	Name string `gorm:"size:255;unique"`
}
