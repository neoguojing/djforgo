package utils

import (
	"reflect"
)

func Struct2Map(objs ...interface{}) map[string]interface{} {
	var data = make(map[string]interface{})
	
	for _,obj := range  objs {
		t := reflect.TypeOf(obj)
		v := reflect.ValueOf(obj)
	
		for i := 0; i < t.NumField(); i++ {
			data[t.Field(i).Name] = v.Field(i).Interface()
		}
	}
	
	return data
}
