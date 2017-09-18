package admin

import (
	//l4g "github.com/alecthomas/log4go"
	"djforgo/auth"
)

type UserEditForm struct {
	auth.User
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
