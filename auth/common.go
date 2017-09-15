package auth

import (
	"djforgo/dao"
	"djforgo/system"
	l4g "github.com/alecthomas/log4go"
	"github.com/gorilla/context"
	"net/http"
)

func GetUserByUsername(username *string) (IUser, error) {
	if username == nil {
		return nil, l4g.Error("GetUserByUsername username is invalid", username)
	}

	var user User
	err := dao.DB_Instance.Where("name = ?", username).First(&user).Error
	if err != nil {
		return nil, l4g.Error("GetUserByUsername", err)
	}

	return &user, err
}

func GetAnonymousUser() IUser {
	return &AnonymousUser{}
}

func GetUserByEmail(email *string) (IUser, error) {
	if email == nil {
		return &AnonymousUser{}, nil
	}

	var user User
	err := dao.DB_Instance.Where("email = ?", *email).First(&user).Error
	if err != nil {
		return nil, l4g.Error(err)
	}

	return &user, err
}

func Login_Check(loginform *UserLoginForm) (IUser, error) {
	var user IUser
	user, err := GetUserByEmail(&loginform.Email)
	if err != nil {
		return nil, err
	}

	if !user.CheckPassword(loginform.Password) {
		return nil, l4g.Error("Password is invalid")
	}

	return user, nil
}

func IsAuthticated(r *http.Request) bool {
	user := context.Get(r, system.USERINFO)
	if user == nil {
		return false
	}

	return user.(IUser).IsAuthenticated()
}

func GetUsers(r *http.Request) []User {
	user := context.Get(r, system.USERINFO)
	if user == nil {
		return nil
	}

	var users []User = make([]User, 0)

	if user.(IUser).IsAuthenticated() {
		userObj := user.(*User)
		err := userObj.GetQueryset(&users).Error
		if err != nil {
			l4g.Error("GetUsers", err)
			return nil
		} else {
			l4g.Debug(userObj)
			return users
		}
	}
	return nil
}
