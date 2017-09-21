package auth

import (
	"djforgo/auth/contenttype"
	"djforgo/dao"
	l4g "github.com/alecthomas/log4go"
	set "github.com/deckarep/golang-set"
	"github.com/jinzhu/gorm"
)

type PermissionManager struct {
	dao.Manager
}

func (this *PermissionManager) GetAllPermitions() ([]Permission, error) {
	permitions := make([]Permission, 0)
	err := this.GetQueryset(&permitions).Error
	if err != nil {
		l4g.Error("PermissionManager::GetAllPermitions %v", err)
		return nil, err
	}

	return permitions, nil
}

type Permission struct {
	gorm.Model
	Name         string                  `gorm:"size:255" schema:"name"`
	Content      contenttype.ContentType `gorm:"ForeignKey:Contentrefer;unique" schema:"-"`
	Contentrefer uint                    `schema:"contentid"`
	CodeName     string                  `gorm:"size:100;unique" schema:"codename"`

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
	Name        string       `gorm:"size:80;unique" schema:"name"`
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

func (this *UserManager) Update(user *User) *gorm.DB {
	this.Init()
	db := this.DB.Set("gorm:save_associations", false).Model(user).Updates(*user)

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
	Is_staff    bool         `gorm:"default:False" schema:"istaff"`
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
		this.Init()

		if this.Email != "" {
			if this.ID == 0 {
				err := this.DB.Select("id,name,email,is_active,is_admin,is_staff").
					Where("email = ?", this.Email).First(this).Error
				if err != nil {
					return nil, l4g.Error("User::GetAllPermissions", err)
				}
			}
			perms = make([]Permission, 0)
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

func (this *User) GetAllGroups() ([]Group, error) {
	var groups []Group
	if this.Is_Admin {
		groups = make([]Group, 0)
		err := this.UserManager.Manager.GetQueryset(&groups).Error
		if err != nil {
			return nil, l4g.Error("User::GetAllGroups", err)
		}
		this.Groups = groups
	} else if this.Is_staff {
		this.Init()
		if this.Email != "" {
			if this.ID == 0 {
				err := this.DB.Select("id,name,email,is_active,is_admin,is_staff").
					Where("email = ?", this.Email).First(this).Error
				if err != nil {
					return nil, l4g.Error("User::GetAllGroups", err)
				}
			}
			groups = make([]Group, 0)
			err := this.DB.Model(this).Related(&groups, "Groups").Error
			if err != nil {
				return nil, l4g.Error("User::GetAllGroups", err)
			}
			this.Groups = groups
		}

	} else {

	}
	return groups, nil
}

func (this *User) GetGroupPermissions() ([]Permission, error) {
	return nil, nil
}

func (this *User) UserHasPermission() bool {
	return false
}

func (this *User) GetUnUsedPermitions() ([]interface{}, error) {
	if this.IsAdmin() {
		return nil, nil
	}

	permissions, err := this.GetAllPermissions()
	if err != nil {
		return nil, err
	}

	inUseSet := set.NewSet()
	for _, v := range permissions {
		inUseSet.Add(v)
	}

	allPermitons, err1 := GetAllPermitions()
	if err1 != nil {
		return nil, err1
	}

	fullSet := set.NewSet()
	for _, v := range allPermitons {
		fullSet.Add(v)
	}

	noUseSet := fullSet.Difference(inUseSet)
	return noUseSet.ToSlice(), nil
}

func (this *User) AddPermitions(ids []uint) bool {
	this.Init()

	tx := this.DB.Begin()

	for _, v := range ids {
		perm := Permission{}
		perm.ID = v
		if err := tx.Model(this).Association("Permissions").Append(perm).Error; err != nil {
			tx.Rollback()
			l4g.Error("AddPermitions:%v", err)
			return false
		}
	}

	tx.Commit()
	return true
}

//ids is the permitions need to delete
func (this *User) DelPermitions(ids []uint) bool {
	this.Init()

	tx := this.DB.Begin()

	for _, v := range ids {
		perm := Permission{}
		perm.ID = v
		if err := tx.Model(this).Association("Permissions").Delete(perm).Error; err != nil {
			tx.Rollback()
			l4g.Error("DelPermitions:%v", err)
			return false
		}
	}

	tx.Commit()

	this.Permissions = this.Permissions[:0]
	return true
}

func (this *User) AddGroups(ids []uint) bool {
	this.Init()

	tx := this.DB.Begin()

	for _, v := range ids {
		group := Group{}
		group.ID = v
		if err := tx.Model(this).Association("Groups").Append(group).Error; err != nil {
			tx.Rollback()
			l4g.Error("AddGroups:%v", err)
			return false
		}
	}

	tx.Commit()
	return true
}

//ids is the groups` id need to delete
func (this *User) DelGroups(ids []uint) bool {
	this.Init()

	tx := this.DB.Begin()

	for _, v := range ids {
		group := Group{}
		group.ID = v
		if err := tx.Model(this).Association("Groups").Delete(group).Error; err != nil {
			tx.Rollback()
			l4g.Error("DelGroups:%v", err)
			return false
		}
	}

	tx.Commit()

	this.Groups = this.Groups[:0]
	return true
}

func (this *User) GetUnUsedGroups() ([]interface{}, error) {
	if this.IsAdmin() {
		return nil, nil
	}

	groups, err := this.GetAllGroups()
	if err != nil {
		return nil, err
	}

	all, err1 := GetAllGroups()
	if err1 != nil {
		return nil, err1
	}

	rtn := make([]interface{}, 0)
	for _, v1 := range all {
		hasSameId := false
		for _, v2 := range groups {
			if v1.ID == v2.ID {
				hasSameId = true
				goto Jump
			}
		}
	Jump:
		if !hasSameId {
			rtn = append(rtn, v1)
		}
	}

	return rtn, nil
}

type AnonymousUser struct {
}

func (this *AnonymousUser) GetUserID() uint {
	return 0
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

func (this *AnonymousUser) GetAllGroups() ([]Group, error) {
	return nil, nil
}

func (this *AnonymousUser) IsAdmin() bool {
	return false
}

func (this *AnonymousUser) GetUnUsedPermitions() ([]interface{}, error) {
	return nil, nil
}

func (this *AnonymousUser) GetUnUsedGroups() ([]interface{}, error) {
	return nil, nil
}
