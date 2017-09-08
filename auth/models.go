package auth

import (
	"djforgo/dao"
	//l4g "github.com/alecthomas/log4go"
	"github.com/jinzhu/gorm"
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

type UserManager struct {
	dao.Manager
}

func (this *UserManager) GetQueryset(out interface{}) *gorm.DB {
	this.Init()
	return this.DB.Where("is_admin <> ?",true).Find(out)
}


func (this *UserManager) CreateUser(user *User) error {
	user.Is_Admin = false
	user.Is_staff = false
	user.SetPassword("")

	return this.Save(user).Error
}

func (this *UserManager) CreateAdminUser(user *User) error {
	user.Is_Admin = true
	user.Is_staff = true
	user.SetPassword("")

	return this.Save(user).Error
}

type User struct {
	BaseUser
	Is_staff    bool `gorm:"default:False"`
	UserManager `gorm:"-"`
}

func (this *User) SendEmail() error {
	return nil
}

type AnonymousUser struct {
}

func (this *AnonymousUser) GetUserName() string {
	return ""
}

func (this *AnonymousUser) SetUserName(username string) {

}

func (this *AnonymousUser) GetEmail() string {
	return ""
}

func (this *AnonymousUser) SetEmail(email string) {
}

func (this *AnonymousUser) Save() error {
	return nil
}

func (this *AnonymousUser) IsAnonymous() bool {
	return true
}

func (this *AnonymousUser) IsAuthenticated() bool {
	return false
}

func (this *AnonymousUser) SetPassword(rawpassword string) error {
	return nil
}

func (this *AnonymousUser) CheckPassword(rawPassword string) bool {
	return true
}
