package auth

import (
	"djforgo/dao"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type IUser interface {
	GetUserName() string
	SetUserName(string)
	GetEmail() string
	SetEmail(string)
	IsAnonymous() bool
	IsAuthenticated() bool
	SetPassword(string) error
	CheckPassword(string) bool

	GetAllPermissions() ([]Permission, error)
	GetGroupPermissions() ([]Permission, error)
	UserHasPermission() bool
}

type BaseUserManager struct {
	dao.Manager
}

type BaseUser struct {
	gorm.Model
	Name      string `gorm:"type:varchar(50);unique"`
	Email     string `gorm:"type:varchar(50);not null;unique"`
	Password  string `gorm:"not null"`
	Is_active bool   `gorm:"default:True"`
	Is_Admin  bool   `gorm:"default:False"`

	BaseUserManager `gorm:"-"`
}

func (this *BaseUser) GetUserName() string {
	return this.Name
}

func (this *BaseUser) SetUserName(username string) {
	this.Name = username
}

func (this *BaseUser) GetEmail() string {
	return this.Email
}

func (this *BaseUser) SetEmail(email string) {
	this.Email = email
}

func (this *BaseUser) IsAnonymous() bool {
	return false
}

func (this *BaseUser) IsAuthenticated() bool {
	return true
}

func (this *BaseUser) SetPassword(rawpassword string) error {

	bytes, err := bcrypt.GenerateFromPassword([]byte(rawpassword), 14)
	if err != nil {
		return err
	}

	this.Password = string(bytes)
	return nil
}

func (this *BaseUser) CheckPassword(rawpassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(this.Password), []byte(rawpassword))
	return err == nil
}

func (this *BaseUser) GetAllPermissions() ([]Permission, error) {
	return nil, nil
}

func (this *BaseUser) GetGroupPermissions() ([]Permission, error) {
	return nil, nil
}

func (this *BaseUser) UserHasPermission() bool {
	return false
}
