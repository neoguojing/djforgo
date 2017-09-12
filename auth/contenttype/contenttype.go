package contenttype

import (
	"djforgo/dao"
	"github.com/jinzhu/gorm"
)

type ContentTypeManager struct {
	dao.Manager
}

type ContentType struct {
	gorm.Model
	Name      string `gorm:"size:100"`
	AppLabel  string `gorm:"size:100;unique"`
	ModelName string `gorm:"size:100"`

	ContentTypeManager `gorm:"-"`
}
