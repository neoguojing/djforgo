package admin

import (
	//l4g "github.com/alecthomas/log4go"
	"djforgo/auth"
)

type UserEditForm struct {
	auth.User
	PermissionID []uint `schema:"f_permitions"`
	GroupID      []uint `schema:"f_groups"`
}

func (this *UserEditForm) Init() error {
	return nil
}

func (this *UserEditForm) Save() error {
	for _, v := range this.PermissionID {
		this.AddPermition(v)
	}

	for _, v := range this.GroupID {
		this.AddGroup(v)
	}

	return this.User.Save(&this.User).Error
}

func (this *UserEditForm) Valid() error {
	return nil
}

type PermitionEditForm struct {
	ID uint `schema:"id"`
	auth.Permission
}

func (this *PermitionEditForm) Valid() error {
	return nil
}

type GroupEditForm struct {
	ID uint `schema:"id"`
	auth.Group
}

func (this *GroupEditForm) Valid() error {
	return nil
}
