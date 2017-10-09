package utils

import (
	l4g "github.com/alecthomas/log4go"
	"github.com/bluele/gforms"
	"net/http"
	"reflect"
)

type Object interface {
}

type IForm interface {
	Object

	Init(r *http.Request)
	IsValid() bool
	Fields() []gforms.FieldInterface
	GetModel() interface{}
}

type IModel interface {
	Object
}

var G_ObjRegisterStore = initobjStore()

func initobjStore() *objStore {
	return &objStore{
		store: make(map[string]Object),
	}
}

type objStore struct {
	store map[string]Object
}

func (this *objStore) Set(value Object) error {
	if reflect.TypeOf(value).Kind() == reflect.Ptr {
		return l4g.Error("objStore::Set,value can not be ptr")
	}

	key := reflect.TypeOf(value).Name()

	if _, ok := this.store[key]; ok {
		return l4g.Error("objStore::Set,key=%s already exist", key)
	}
	this.store[key] = value
	return nil
}

func (this *objStore) Get(key string) Object {
	obj, ok := this.store[key]
	if !ok {
		l4g.Error("objStore::Get,key=%s was not exist", key)
		return nil
	}
	return obj
}

func (this *objStore) New(key string) Object {
	//copy object
	obj, ok := this.store[key]
	if !ok {
		l4g.Error("objStore::Get,key=%s was not exist in %v", key, this.store)
		return nil
	}
	return reflect.New(reflect.TypeOf(obj)).Interface()
}
