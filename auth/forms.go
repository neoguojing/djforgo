package auth

import (
	l4g "github.com/alecthomas/log4go"
)

type UserLoginForm struct {
	Name     string `schema:"-"`
	Email    string `schema:"email"`
	Password string `schema:"password"`
}

func (this *UserLoginForm) Valid() error {
	if this.Email == "" {
		return l4g.Error("Email was empty")
	}

	return nil
}

type UserCreationForm struct {
	Name      string `schema:"username"`
	Email     string `schema:"email"`
	Password1 string `schema:"password1"`
	Password2 string `schema:"password2"`
}

func (this *UserCreationForm) Valid() error {
	if !this.isNameValid() {
		return l4g.Error("username invalid")
	}
	if !this.isEmailValid() {
		return l4g.Error("email invalid")
	}
	if !this.isPasswordValid() {
		return l4g.Error("password invalid")
	}

	return nil
}

func (this *UserCreationForm) isNameValid() bool {
	if this.Name == "" {
		return false
	}
	return true
}

func (this *UserCreationForm) isEmailValid() bool {
	if this.Email == "" {
		return false
	}

	return true
}

func (this *UserCreationForm) isPasswordValid() bool {
	if this.Password1 == "" || this.Password2 == "" {
		return false
	}
	return this.Password1 == this.Password2
}
