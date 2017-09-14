package register

import (
	"djforgo/auth"
	"djforgo/auth/contenttype"
	"djforgo/config"
	"djforgo/dao"
	"fmt"
	"github.com/jinzhu/gorm"
	"reflect"
	"strings"
)

var permitionType = []string{"add", "edit", "del", "view"}

var ModelSetInstance *ModelSet

func newModelSet() *ModelSet {
	return &ModelSet{
		models: make(map[string]interface{}),
	}
}

type ModelSet struct {
	models map[string]interface{}
}

func (this *ModelSet) Register(pmodel interface{}) bool {
	type_model := reflect.TypeOf(pmodel)

	key := type_model.PkgPath() + "." + type_model.Name()
	fmt.Println(key)
	if _, ok := this.models[key]; ok {
		fmt.Println("%v already exist", key)
		return false
	}

	this.models[key] = pmodel

	return true
}

func (this *ModelSet) CreateTables() bool {

	err := dao.DB_Init()
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer dao.DB_Destroy()

	for key, v := range this.models {
		if dao.DB_Instance.HasTable(v) {
			err = dao.DB_Instance.DropTable(v).Error
			if err != nil {
				fmt.Println("drop table", key, err)
				return false
			}
		}

		err = dao.DB_Instance.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(v).Error
		if err != nil {
			fmt.Println("create ", key, err)
			return false
		}

		this.handleIndex(dao.DB_Instance, v)
	}

	return true
}

func (this *ModelSet) handleIndex(db *gorm.DB, modelObj interface{}) {
	switch reflect.TypeOf(modelObj).Name() {
	case "ContentType":
		db.Model(modelObj).AddUniqueIndex("app_label", "model_name")
	}
}

func (this *ModelSet) CreateContentType() bool {

	err := dao.DB_Init()
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer dao.DB_Destroy()

	for _, v := range this.models {
		if !dao.DB_Instance.HasTable(v) {
			fmt.Println("content_type table does not exit ")
			return false
		}

		content_type := contenttype.ContentType{
			Name:      reflect.TypeOf(v).Name(),
			AppLabel:  reflect.TypeOf(v).PkgPath(),
			ModelName: reflect.TypeOf(v).Name(),
		}
		err = content_type.Save(&content_type).Error
		if err != nil {
			fmt.Println("CreateContentType", err)
			return false
		}
	}

	return true
}

func (this *ModelSet) CreatePermissions() bool {

	err := dao.DB_Init()
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer dao.DB_Destroy()

	for _, v := range this.models {
		if !dao.DB_Instance.HasTable(v) {
			fmt.Println("Permissions table does not exit ")
			return false
		}

		content := &contenttype.ContentType{
			Name:      reflect.TypeOf(v).Name(),
			AppLabel:  reflect.TypeOf(v).PkgPath(),
			ModelName: reflect.TypeOf(v).Name(),
		}

		contentId, err := content.GetId()
		if err != nil {
			fmt.Println("CreatePermissions", err)
			return false
		}

		for _, perm := range permitionType {
			permission := auth.Permission{
				Name:         fmt.Sprintf("Can %s %s", perm, reflect.TypeOf(v).Name()),
				CodeName:     fmt.Sprintf("%s_%s", perm, strings.ToLower(reflect.TypeOf(v).Name())),
				Contentrefer: contentId,
			}
			err = permission.Save(&permission).Error
			if err != nil {
				fmt.Println("CreatePermissions", err)
				return false
			}
		}

	}

	return true
}

func init() {
	ModelSetInstance = newModelSet()
	ModelSetInstance.Register(auth.User{})
	ModelSetInstance.Register(auth.Group{})
	ModelSetInstance.Register(auth.Permission{})
	ModelSetInstance.Register(contenttype.ContentType{})

	for _, v := range config.AppModels {
		ModelSetInstance.Register(v)
	}
}
