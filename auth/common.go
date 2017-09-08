package auth

import (
	"djforgo/dao"
	l4g "github.com/alecthomas/log4go"
)

func GetUserByUsername(username string) (IUser, error) {
	if username == "" {
		return &AnonymousUser{}, nil
	}

	var user User
	err := dao.DB_Instance.Where("name = ?", username).First(&user).Error
	if err != nil {
		return nil, l4g.Error(err)
	}

	return &user, err
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
