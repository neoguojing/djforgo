package dao

import ()

type IManager interface {
	GetQueryset()
}

type Manager struct {
}

type EmptyManager struct {
}
