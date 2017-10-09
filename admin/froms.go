package admin

import (
	"djforgo/auth"
	"djforgo/utils"
	l4g "github.com/alecthomas/log4go"
	"github.com/bluele/gforms"
	"net/http"
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

type PasswordResetForm struct {
	OldPassword string `schema:"oldpassword"`
	Password1   string `schema:"password1"`
	Password2   string `schema:"password2"`
}

func (this *PasswordResetForm) Save(user *auth.User) error {
	user.SetPassword(this.Password1)
	if err := user.Update(user).Error; err != nil {
		l4g.Error("PasswordResetForm::Save %v", err)
		return err
	}
	return nil
}

func (this *PasswordResetForm) Valid(user *auth.User) error {
	if this.OldPassword == "" || this.Password1 == "" || this.Password2 == "" {
		return l4g.Error("Passwords was empty")
	}

	if !user.CheckPassword(this.OldPassword) {
		return l4g.Error("Old password was invalid")
	}

	if this.Password1 != this.Password2 {
		return l4g.Error("New passwords was not equal")
	}

	return nil
}

type PermitionForm struct {
	gforms.ModelFormInstance
}

func (this *PermitionForm) Init(r *http.Request) {
	this.ModelFormInstance = *gforms.DefineModelForm(auth.Permission{}, gforms.NewFields(
		gforms.NewTextField(
			"name",
			gforms.Validators{
				gforms.Required(),
			},
		),
		gforms.NewTextField(
			"codename",
			gforms.Validators{
				gforms.Required(),
			},
		),
	))(r)
}

type GroupForm struct {
	ID uint `schema:"id"`
	auth.Group
}

func init() {
	utils.G_ObjRegisterStore.Set(PermitionForm{})
}
