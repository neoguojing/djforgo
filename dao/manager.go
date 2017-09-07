package dao

import (
	"github.com/jinzhu/gorm"
	"sync"
)

type IManager interface {
	GetQueryset(interface{}) *gorm.DB
	Save(interface{}) *gorm.DB
	Delete(interface{}) *gorm.DB
}

var managerPool = sync.Pool{
	// 当对象池中无对象时，如何创建新对象
	New: func() interface{} { return new(Manager) },
}

type Manager struct {
	db *gorm.DB
	sync.Once
}

func NewManager() *Manager {
	rtn := managerPool.Get().(*Manager)
	rtn.db = DB_Instance
	return rtn
}

func (this *Manager) init() {
	this.Do(func() {
		this.db = DB_Instance
	})
}

func (this *Manager) GetQueryset(out interface{}) *gorm.DB {
	this.init()
	return this.db.Find(out)
}

func (this *Manager) Save(in interface{}) *gorm.DB {
	this.init()

	var db *gorm.DB
	if this.db.NewRecord(in) {
		db = this.db.Create(in)
	} else {
		db = this.db.Save(in)
	}

	return db
}

func (this *Manager) Delete(in interface{}) *gorm.DB {
	this.init()

	var db *gorm.DB
	db = this.db.Delete(in)

	return db
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
