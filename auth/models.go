package auth

import (
	"djforgo/dao"
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

type User struct {
	BaseUser
	Is_staff bool `schema:"-"`
}

func (this *User) SendEmail() error {
	return nil
}

type AnonymousUser struct {
}

func (this *AnonymousUser) GetUserName() string {
	return ""
}

func (this *AnonymousUser) GetEmail() string {
	return ""
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

func (this *AnonymousUser) SetPassword() error {
	return nil
}

func (this *AnonymousUser) CheckPassword() error {
	return nil
}
