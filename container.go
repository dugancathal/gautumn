package gautumn

import (
	"fmt"
	"reflect"
)

type InjectionContainer interface {
	GetDep(name string) reflect.Value
	RegisterByName(name string, value interface{})
	RegisterByType(value interface{})
	RegisterByInterface(inter interface{}, value interface{})
}

type Container map[string]reflect.Value

func (c Container) GetDep(name string) reflect.Value {
	if val, ok := c[name]; ok {
		if val.Type().Kind() == reflect.Func {
			return reflect.ValueOf(Injected(val.Interface(), c))
		}
		return val
	} else {
		panic(fmt.Errorf("Could not find dependency for type: %s", name))
	}
}

func (c Container) RegisterByName(name string, value interface{}) {
	if old, exists := c[name]; exists {
		panic(fmt.Errorf("Attempted to register two injectables with name %s \n\t1: %#v\n\t2: %#v", name, old, value))
	}
	c[name] = reflect.ValueOf(value)
}

func (c Container) RegisterByType(value interface{}) {
	c.RegisterByName(reflect.TypeOf(value).String(), value)
}

func (c Container) RegisterByInterface(inter interface{}, value interface{}) {
	c.RegisterByName(reflect.TypeOf(inter).Elem().String(), value)
}

func (c Container) RegisterByConstructor(constructor interface{}) {
	construct := reflect.TypeOf(constructor)
	if construct.NumOut() != 1 {
		panic("Registering by injectable requires a function that returns exactly 1 item: the item to construct")
	}

	outType := construct.Out(0)
	c.RegisterByName(outType.String(), constructor)
}
