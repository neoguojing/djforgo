package forms

import ()

type IForm interface {
	Valid() error
}

type Form struct {
}

func (this *Form) Valid() error {
	return nil
}

func (this *Form) isEmailValid() bool {
	return true
}
