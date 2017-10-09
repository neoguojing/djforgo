package dao

import (
	"github.com/jinzhu/gorm"
	"sync"
)

type IManager interface {
	GetQueryset(interface{}) *gorm.DB
	Save(interface{}) *gorm.DB
	Update(interface{}) *gorm.DB
	Delete(interface{}) *gorm.DB
}

var managerPool = sync.Pool{
	// 当对象池中无对象时，如何创建新对象
	New: func() interface{} { return new(Manager) },
}

type Manager struct {
	DB *gorm.DB
	sync.Once
}

func NewManager() *Manager {
	rtn := managerPool.Get().(*Manager)
	rtn.DB = DB_Instance
	return rtn
}

func (this *Manager) Init() {
	this.Do(func() {
		this.DB = DB_Instance
	})
}

func (this *Manager) GetQueryset(out interface{}) *gorm.DB {
	this.Init()
	return this.DB.Find(out)
}

func (this *Manager) Save(in interface{}) *gorm.DB {
	this.Init()

	var db *gorm.DB
	if this.DB.NewRecord(in) {
		db = this.DB.Create(in)
	} else {
		db = this.DB.Save(in)
	}

	return db
}

func (this *Manager) Update(in interface{}) *gorm.DB {
	this.Init()

	db := this.DB.Set("gorm:save_associations", false).Model(&in).Updates(in)

	return db
}

func (this *Manager) Delete(in interface{}) *gorm.DB {
	this.Init()

	var db *gorm.DB
	db = this.DB.Delete(in)

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
