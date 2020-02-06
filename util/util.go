package util

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

func Contains(cont interface{}, val interface{}) bool {
	cv := reflect.ValueOf(cont)
	if cv.Kind() == reflect.Slice || cv.Kind() == reflect.Array {
		for i := 0; i < cv.Len(); i++ {
			if cv.Index(i).Interface() == val {
				return true
			}
		}
	}
	return false
}

func ValsOf(m interface{}) interface{} {
	v := reflect.ValueOf(m)
	if v.Kind() != reflect.Map {
		panic("parameter is not a map")
	}
	k := v.MapKeys()
	vals := reflect.MakeSlice(reflect.SliceOf(v.Type().Elem()), len(k), len(k))
	for i, s := range k {
		vals.Index(i).Set(v.MapIndex(s))
	}
	return vals.Interface()
}

func Convert(v reflect.Value, tp reflect.Type) interface{} {
	vt := v.Type()
	if v.Kind() == reflect.Slice {
		if vt.Elem() == tp {
			return v.Interface()
		}
		if tp.Kind() == reflect.String {
			rs := make([]string, v.Len(), v.Len())
			for i := range rs {
				rs[i] = fmt.Sprint(v.Index(i).Interface())
			}
			return rs
		} else if vt.Elem().ConvertibleTo(tp) {
			r := reflect.MakeSlice(reflect.SliceOf(tp), v.Len(), v.Len())
			for i := 0; i < v.Len(); i++ {
				x := v.Index(i).Convert(tp)
				r.Index(i).Set(x)
			}
			return r.Interface()
		}
	} else {
		if vt == tp {
			return v.Interface()
		}
		if tp.Kind() == reflect.String {
			return fmt.Sprint(v.Interface())
		} else if v.Type().ConvertibleTo(tp) {
			return v.Convert(tp).Interface()
		}
	}
	panic("can't convert")
}

func MapInterface(m map[string]reflect.Value) map[string]interface{} {
	r := map[string]interface{}{}
	for k, v := range m {
		r[k] = v.Interface()
	}
	return r
}
