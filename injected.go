package gautumn

import "reflect"

func Injected(injectee interface{}, diContainer InjectionContainer) interface{} {
	funcType := reflect.TypeOf(injectee)
	args := []reflect.Value{}
	for i := 0; i < funcType.NumIn(); i++ {
		args = append(args, diContainer.GetDep(funcType.In(i).String()))
	}
	injectable := reflect.ValueOf(injectee)
	return injectable.Call(args)[0].Interface()
}

func DefaultInjected(injectee interface{}) interface{} {
	return Injected(injectee, DefaultContainer)
}
