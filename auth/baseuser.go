package auth

import (
	"djforgo/dao"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type IUser interface {
	GetUserName() string
	GetEmail() string
	Save() error
	IsAnonymous() bool
	IsAuthenticated() bool
	SetPassword() error
	CheckPassword(string) bool
}

type BaseUserManager struct {
	dao.Manager
}

type BaseUser struct {
	gorm.Model
	Name      string `schema:"name" gorm:"type:varchar(50);unique"`
	Email     string `schema:"email" gorm:"type:varchar(50);not null;unique"`
	Password  string `schema:"password" gorm:"not null"`
	Is_active bool   `schema:"-"`
}

func (this *BaseUser) GetUserName() string {
	return this.Name
}

func (this *BaseUser) GetEmail() string {
	return this.Email
}

func (this *BaseUser) Save() error {
	if err := this.SetPassword(); err != nil {
		return err
	}

	return nil
}

func (this *BaseUser) IsAnonymous() bool {
	return false
}

func (this *BaseUser) IsAuthenticated() bool {
	return true
}

func (this *BaseUser) SetPassword() error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(this.Password), 14)
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
