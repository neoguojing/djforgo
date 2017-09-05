package dao

import ()

type IModel interface {
	Create()
	Update()
	Delete()
	Query()
}
