package util

import (
	"reflect"
)

func Option(t interface{}, o []interface{}) reflect.Value {
	tv := reflect.ValueOf(t)
	for tv.Kind() == reflect.Interface || tv.Kind() == reflect.Ptr {
		tv = tv.Elem()
	}
	for _, x := range o {
		v := reflect.ValueOf(x)
		for v.Kind() == reflect.Interface || v.Kind() == reflect.Ptr {
			v = v.Elem()
		}
		if v.Type() == tv.Type() {
			return v
		}
	}
	return tv
}

func StrOption(t interface{}, o []interface{}) string {
	return Option(t, o).String()
}

func IntOption(t interface{}, o []interface{}) int {
	return int(Option(t, o).Int())
}

func FloatOption(t interface{}, o []interface{}) float64 {
	return Option(t, o).Float()
}
