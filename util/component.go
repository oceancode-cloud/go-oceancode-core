package _util

import "reflect"

var beans = make(map[reflect.Type]interface{})

func Bind(iFace interface{}, impl interface{}) {
	beans[reflect.TypeOf(iFace).Elem()] = impl
}

func Resolve(beanType interface{}) interface{} {
	return beans[reflect.TypeOf(beanType).Elem()]
}
