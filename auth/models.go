package auth

import (
	"djforgo/auth/contenttype"
	"djforgo/dao"
	l4g "github.com/alecthomas/log4go"
	"github.com/jinzhu/gorm"
)

type PermissionManager struct {
	dao.Manager
}

type Permission struct {
	gorm.Model
	Name         string                  `gorm:"size:255"`
	Content      contenttype.ContentType `gorm:"ForeignKey:Contentrefer;unique"`
	Contentrefer uint
	CodeName     string `gorm:"size:100;unique"`

	PermissionManager `gorm:"-"`
}

func (this *Permission) Create() {

}

func (this *Permission) Add() {

}

func (this *Permission) Wrapper() (string, string) {
	this.Init()

	content := contenttype.ContentType{}
	err := this.DB.Model(this).Related(&content, "Contentrefer").Error
	if err != nil {
		l4g.Error("Permission::Wrapper-%v", err)
		return "", ""
	}

	this.Content = content

	return content.AppLabel, this.CodeName
}

type GroupManager struct {
	dao.Manager
}

type Group struct {
	gorm.Model
	Name        string       `gorm:"size:80;unique"`
	Permissions []Permission `gorm:"many2many:group_permissions;"`

	GroupManager `gorm:"-"`
}

type UserManager struct {
	dao.Manager
}

func (this *UserManager) GetQueryset(out interface{}) *gorm.DB {
	this.Init()
	db := this.DB.Select("id,name,email,is_active,is_admin,is_staff").
		Where("is_admin <> ?", true).Find(out)

	return db
}

func (this *UserManager) CreateUser(user *User) error {
	user.Is_Admin = false
	user.Is_staff = false

	return this.Save(user).Error
}

func (this *UserManager) CreateAdminUser(user *User) error {
	user.Is_Admin = true
	user.Is_staff = true
	_, err := user.GetAllPermissions()
	if err != nil {
		return err
	}

	return this.Save(user).Error
}

type User struct {
	BaseUser
	Is_staff    bool         `gorm:"default:False"`
	Groups      []Group      `gorm:"many2many:user_groups;"`
	Permissions []Permission `gorm:"many2many:user_permissions;"`

	UserManager `gorm:"-"`
}

func (this *User) SendEmail() error {
	return nil
}

func (this *User) copy(user *User) {
	*this = *user
}

func (this *User) GetAllPermissions() ([]Permission, error) {
	var perms []Permission
	if this.Is_Admin {
		perms = make([]Permission, 0)
		err := this.UserManager.Manager.GetQueryset(&perms).Error
		if err != nil {
			return nil, l4g.Error("User::GetAllPermissions", err)
		}
		this.Permissions = perms
	} else if this.Is_staff {
		if this.Email != "" {
			if this.ID == 0 {
				err := this.DB.Select("id,name,email,is_active,is_admin,is_staff").
					Where("email = ?", this.Email).First(this).Error
				if err != nil {
					return nil, l4g.Error("User::GetAllPermissions", err)
				}
			}
			err := this.DB.Model(this).Related(&perms, "Permissions").Error
			if err != nil {
				return nil, l4g.Error("User::GetAllPermissions", err)
			}
			this.Permissions = perms
		}

	} else {

	}
	return perms, nil
}

func (this *User) SetPermissions() error {
	var perms []Permission
	if this.Is_Admin {
		perms = make([]Permission, 0)
		err := this.UserManager.Manager.GetQueryset(&perms).Error
		if err != nil {
			return err
		}

		this.Permissions = perms
	}
	return nil
}

func (this *User) GetGroupPermissions() ([]Permission, error) {
	return nil, nil
}

func (this *User) UserHasPermission() bool {
	return false
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

func (this *AnonymousUser) GetAllPermissions() ([]Permission, error) {
	return nil, nil
}

func (this *AnonymousUser) GetGroupPermissions() ([]Permission, error) {
	return nil, nil
}

func (this *AnonymousUser) UserHasPermission() bool {
	return false
}
