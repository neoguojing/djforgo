package contenttype

import (
	"djforgo/dao"
	l4g "github.com/alecthomas/log4go"
	"github.com/jinzhu/gorm"
)

type ContentTypeManager struct {
	dao.Manager
}

func (this *ContentTypeManager) GetQueryset() ([]ContentType, error) {
	var outPut = make([]ContentType, 0)
	err := this.Manager.GetQueryset(&outPut).Error
	if err != nil {
		return nil, err
	}

	return outPut, nil
}

type ContentType struct {
	gorm.Model
	Name      string `gorm:"size:100"`
	AppLabel  string `gorm:"size:100"`
	ModelName string `gorm:"size:100"`

	ContentTypeManager `gorm:"-"`
}

func (this *ContentType) GetId() (uint, error) {

	if this.AppLabel == "" || this.ModelName == "" {
		return 0, l4g.Error("AppLabel or ModelName was empty")
	}

	this.Init()
	var content ContentType
	err := this.DB.Where("app_label = ? AND model_name = ?", this.AppLabel, this.ModelName).First(&content).Error
	if err != nil {
		return 0, err
	}

	return content.ID, nil
}
