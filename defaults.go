package gautumn

import "reflect"

var DefaultContainer Container = Container{}

func RegisterByName(name string, value interface{}) {
	DefaultContainer.RegisterByName(name, value)
}
func RegisterByType(value interface{}) {
	DefaultContainer.RegisterByType(value)
}
func RegisterByInterface(inter interface{}, value interface{}) {
	DefaultContainer.RegisterByInterface(inter, value)
}
func RegisterByConstructor(constructor interface{}) {
	DefaultContainer.RegisterByConstructor(constructor)
}

func GetDep(name string) reflect.Value {
	return DefaultContainer.GetDep(name)
}

func Clean() {
	DefaultContainer = Container{}
}
