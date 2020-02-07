package lazy

import (
	"reflect"
)

type Source struct {
	Ctx  interface{}
	Tp   reflect.Type
	Next func(ctx interface{}) reflect.Value
}

func New(c interface{}) *Source {
	v := reflect.ValueOf(c)
	if v.Kind() == reflect.Chan && v.Elem().Kind() == reflect.Struct {
		scase := []reflect.SelectCase{{Dir: reflect.SelectRecv, Chan: v}}
		f := func(_ interface{}) reflect.Value {
			_, r, ok := reflect.Select(scase)
			if !ok {
				return reflect.ValueOf(false)
			}
			return r
		}
		return &Source{nil, v.Elem().Type(), f }
	} else {
		panic("only chan struct{...} is allowed as an argument")
	}
}

