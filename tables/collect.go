package tables

import "reflect"

func (t *Table) Collect(s interface{}) interface{} {
	tp := reflect.TypeOf(s)
	if tp.Kind() == reflect.Ptr {
		tp = tp.Elem()
	}
	if tp.Kind() != reflect.Struct {
		panic("only struct{...} is allowed as an argument")
	}
	r := reflect.MakeSlice(reflect.SliceOf(tp),t.length,t.length)
	for i:=0; i<t.length; i++ {
		t.FillRow(i, tp, r.Index(i).Addr())
	}
	return r.Interface()
}
