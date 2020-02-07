package tables

import (
	"github.com/sudachen/go-tables/util"
	"reflect"
)

/*
Row returns table row as a map of reflect.Value
	t := tables.New([]struct{Name string; Age int; Rate float32}{{"Ivanov",32,1.2},{"Petrov",44,1.5}})
	t.Row(0)["Name"].String() -> "Ivanov"
*/
func (t *Table) Row(row int) map[string]reflect.Value {
	r := map[string]reflect.Value{}
	for i, n := range t.names {
		// prevent to change value in slice via returned reflect.Value
		r[n] = reflect.ValueOf(t.columns[i].Index(row).Interface())
	}
	return r
}

/*
MakeRow gets required fields as a struct
*/
func (t *Table) MakeRow(i int, tp reflect.Type) reflect.Value {
	v := reflect.New(tp)
	fl := tp.NumField()
	for i := 0; i < fl; i++ {
		n := tp.Field(i).Name
		j := util.IndexOf(n, t.names)
		if j < 0 {
			panic("table does not have field " + n)
		}
		v.Field(i).Set(t.columns[j].Index(i))
	}
	return v
}

/*
Fetch fills struct with table' row data

	t := tables.New([]struct{Name string; Age int; Rate float32}{{"Ivanov",32,1.2},{"Petrov",44,1.5}})
	r := struct{Name string; Age int}{}
	t.Fetch(0,&r)
	r.Name -> "Ivanov"
	r.Age -> 32
*/
func (t *Table) Fetch(i int, r interface{}) {
	q := reflect.ValueOf(r)
	if q.Kind() != reflect.Ptr && q.Elem().Kind() != reflect.Struct {
		panic("only pointer to a struct is allowed as argumet")
	}
	q.Elem().Set(t.MakeRow(i, q.Elem().Type()))
}