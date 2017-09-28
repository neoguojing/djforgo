package utils

import (
	l4g "github.com/alecthomas/log4go"
	"reflect"
)

var G_ObjRegisterStore = initobjStore()

func initobjStore() *objStore {
	return &objStore{
		store: make(map[string]interface{}),
	}
}

type objStore struct {
	store map[string]interface{}
}

func (this *objStore) Set(value interface{}) error {
	if reflect.TypeOf(value).Kind() == reflect.Ptr {
		return l4g.Error("objStore::Set,value can not be ptr")
	}

	key := reflect.TypeOf(value).Name()

	if _, ok := this.store[key]; ok {
		return l4g.Error("objStore::Set,key already exist")
	}
	this.store[key] = value
	return nil
}

func (this *objStore) Get(key string) interface{} {
	return this.store[key]
}

func (this *objStore) New(key string) interface{} {
	//copy object
	rtn := this.store[key]
	return rtn
}
