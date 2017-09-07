package forms

import ()

type IForm interface {
	Valid() error
	isEmailValid() bool
}

type Form struct {
}

func (this *Form) Valid() error {
	return nil
}

func (this *Form) isEmailValid() bool {
	return true
}
