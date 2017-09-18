package invoker

import (
	"reflect"
)

func InvokeObjectMethod(object interface{}, methodName string, args ...interface{}) (result []reflect.Value, err error) {
	inputs := make([]reflect.Value, len(args))
	for i, _ := range args {
		inputs[i] = reflect.ValueOf(args[i])
	}
	result = reflect.ValueOf(object).MethodByName(methodName).Call(inputs)
	return
}
