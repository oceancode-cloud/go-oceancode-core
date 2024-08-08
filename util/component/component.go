package _component

import "reflect"

var beans = make(map[reflect.Type]interface{})

func Bind(iFace interface{}, impl interface{}) {
	beans[reflect.TypeOf(iFace).Elem()] = impl
}

func Resolve[T interface{}](beanType interface{}) T {
	return beans[reflect.TypeOf(beanType).Elem()].(T)
}
