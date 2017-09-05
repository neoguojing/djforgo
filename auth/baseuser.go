package auth

import (
	"github.com/jinzhu/gorm"
	"neoproj/djforgo/dao"
)

type IUserManager interface {
	NormalizeEmail()
}

type IUser interface {
	GetUserName()
	Save()
	IsAnonymous()
	IsAuthenticated()
	SetPassword()
	CheckPassword()
	NormalizeUserName()
}

type BaseUserManager struct {
	dao.Manager
}

type BaseUser struct {
	gorm.Model
}
