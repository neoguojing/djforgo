package dao

import (
	"github.com/jinzhu/gorm"
	"sync"
)

type IManager interface {
	Init()
	GetQueryset(interface{}) (*gorm.DB, error)
	Release()
}

var managerPool = sync.Pool{
	// 当对象池中无对象时，如何创建新对象
	New: func() interface{} { return new(Manager) },
}

type Manager struct {
	db *gorm.DB
}

func NewManager() *Manager {
	rtn := managerPool.Get().(*Manager)
	rtn.db = DB_Instance
	return rtn
}

func (this *Manager) Init() {
	this.db = DB_Instance
}

func (this *Manager) GetQueryset(out interface{}) *gorm.DB {
	return this.db.Find(out)
}

func (this *Manager) Release() {
	managerPool.Put(this)
}

var emptymanagerPool = sync.Pool{
	// 当对象池中无对象时，如何创建新对象
	New: func() interface{} { return new(EmptyManager) },
}

type EmptyManager struct {
}

func (this *EmptyManager) GetQueryset() error {
	return nil
}
