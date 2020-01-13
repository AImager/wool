package wool

import (
	"errors"
	"reflect"
)

// not export , so others can't implement the package
type dic struct {
	container map[reflect.Type]*__obj
}

type DIC interface {
	Bind(it interface{}, v interface{}) error
	BindSingleton(it interface{}, v interface{}) error
	Make(it interface{}, args ...interface{}) error
}

var di *dic

type __obj struct {
	instance  interface{}
	__type    reflect.Type
	__value   reflect.Value
	singleton bool
}

// DI singleton
func DI() DIC {
	if di == nil {
		di = &dic{
			container: make(map[reflect.Type]*__obj),
		}
	}
	return di
}

func (di *dic) Bind(it interface{}, v interface{}) error {
	return di.bind(it, v, false)
}

func (di *dic) BindSingleton(it interface{}, v interface{}) error {
	return di.bind(it, v, true)
}

// Bind it's type must be no-empty reflect.Interface
// c's type can be reflect.Func or reflect.Struct
func (di *dic) bind(it interface{}, v interface{}, singleton bool) error {
	itType := reflect.TypeOf(it)
	vType := reflect.TypeOf(v)
	if itType.Kind() != reflect.Ptr {
		return errors.New("it is not ptr")
	}
	// just check
	switch vType.Kind() {
	case reflect.Func:
		if vType.In(0) != reflect.TypeOf(di) {
			return errors.New("first params must be DIC")
		}
		if vType.Out(0) != itType {
			return errors.New("first return params must same as it's type")
		}
	default:
	}
	obj := &__obj{
		__type:  vType,
		__value: reflect.ValueOf(v),
	}
	di.set(itType, obj)
	return nil
}

// Make it must be reflect.Ptr
// it's point obj must be reflect.Struct or reflect.Interface
func (di *dic) Make(it interface{}, args ...interface{}) error {
	itType := reflect.TypeOf(it)
	obj, ok := di.load(itType)
	if itType.Kind() != reflect.Ptr || !ok {
		return errors.New("can't find interface")
	}
	switch obj.__type.Kind() {
	case reflect.Func:
		var argValues = make([]reflect.Value, 0)
		for _, arg := range args {
			argValues = append(argValues, reflect.ValueOf(arg))
		}
		// todo error handle
		buildValue := obj.__value.Call(argValues)
		reflect.ValueOf(it).Elem().Set(buildValue[0])
	default:
		reflect.ValueOf(it).Elem().Set(obj.__value)
		return nil
	}
	return nil
}

// todo lock
func (di *dic) set(it reflect.Type, c *__obj) {
	di.container[it] = c
}

func (di *dic) load(it reflect.Type) (obj *__obj, exist bool) {
	obj, exist = di.container[it]
	return
}
