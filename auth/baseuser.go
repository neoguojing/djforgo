package auth

import(
	"neoproj/djforgo/dao"
	"github.com/jinzhu/gorm"
)

type IUserManager interface{
	NormalizeEmail()
}

type IUser interface{
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

