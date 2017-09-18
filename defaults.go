package gautumn

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
