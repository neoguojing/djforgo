package admin

import (
	"djforgo/auth"
	l4g "github.com/alecthomas/log4go"
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
	if !this.IsAdmin() {
		_, err := this.GetAllPermissions()
		if err != nil {
			return err
		}

		delPermitionIds := make([]uint, 0)
		for _, id := range this.PermissionID {
			for _, v := range this.Permissions {
				if id == v.ID {
					v.ID = 0
				}
			}
		}

		for _, v := range this.Permissions {
			if 0 != v.ID {
				delPermitionIds = append(delPermitionIds, v.ID)
			}
		}

		this.DelPermitions(delPermitionIds)
		this.AddPermitions(this.PermissionID)

		_, err = this.GetAllGroups()
		if err != nil {
			return err
		}

		delGroupIds := make([]uint, 0)
		for _, id := range this.GroupID {
			for _, v := range this.Groups {
				if id == v.ID {
					v.ID = 0
				}
			}
		}

		for _, v := range this.Groups {
			if 0 != v.ID {
				delGroupIds = append(delGroupIds, v.ID)
			}
		}

		this.DelGroups(delGroupIds)
		this.AddGroups(this.GroupID)

	}

	err := this.Update(&this.User).Error
	if err != nil {
		l4g.Error("UserEditForm:Save %v", err)
	}
	return err
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
