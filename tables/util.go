package tables

import (
	"fmt"
	"reflect"
	"sort"
)

func SortedDictKeys(m interface{}) []string {
	v := reflect.ValueOf(m)
	if v.Kind() != reflect.Map || v.Type().Key() != reflect.TypeOf("") {
		panic("parameter is not a map")
	}
	keys := KeysOf(m).([]string)
	sort.Strings(keys)
	return keys
}

func KeysOf(m interface{}) interface{} {
	v := reflect.ValueOf(m)
	if v.Kind() != reflect.Map {
		panic("parameter is not a map")
	}
	k := v.MapKeys()
	keys := reflect.MakeSlice(reflect.SliceOf(v.Type().Key()), len(k), len(k))
	for i, s := range k {
		keys.Index(i).Set(s)
	}
	return keys.Interface()
}

func IndexOf(a string, b []string) int {
	for i, v := range b {
		if v == a {
			return i
		}
	}
	return -1
}

func Convert(v reflect.Value, tp reflect.Type) interface{} {
	if v.Kind() == reflect.Slice {
		if v.Type().Elem() == tp { return v.Interface() }
		if v.Type().Elem().ConvertibleTo(tp) {
			r := reflect.MakeSlice(reflect.TypeOf(tp),v.Len(),v.Len())
			for i:=0; i<v.Len(); i++ {
				r.Index(i).Set(v.Index(i).Convert(tp))
			}
			return r.Interface()
		} else if tp == reflect.TypeOf("") {
			rs := make([]string,v.Len(),v.Len())
			for i := range rs {
				rs[i] = fmt.Sprint(v.Index(i).Interface())
			}
			return rs
		}
	} else {
		if v.Type() == tp { return v.Interface() }
		if v.Type().ConvertibleTo(tp) {
			return v.Convert(tp).Interface()
		} else if tp == reflect.TypeOf("") {
			return fmt.Sprint(v.Interface())
		}
	}
	panic("can't convert")
}
