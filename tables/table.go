package tables

import (
	"reflect"
)

type Table struct {
	names   []string
	columns []reflect.Value
	length  int
}

func (t *Table) Len() int {
	return t.length
}

func (t *Table) Names() []string {
	r := make([]string,len(t.names),len(t.names))
	copy(r,t.names)
	return r
}

func Empty() *Table {
	t := &Table{}
	return t
}

/*
tables.New(struct{Name string; Age int; Rate float32}{})
tables.New([]struct{Name string; Age int; Rate float32}{{"Ivanov",32,1.2},{"Petrov",44,1.5}})
tables.New(map[string]interface{}{"Name":[]string{"Ivanov","Petrov"},"Age":[]int{32,44},"Rate":[]float{1.2,1.5}})
*/
func New(o interface{}) *Table {

	q := reflect.ValueOf(o)

	switch q.Kind() {
	case reflect.Ptr: // t.Append(&struct{}{})
		q = q.Elem()
		fallthrough
	case reflect.Struct: // t.Append(struct{}{})
		s := reflect.MakeSlice(reflect.SliceOf(q.Type()),1,1)
		s.Index(0).Set(q)
		q = s
		fallthrough

	case reflect.Slice: // t.Append([]struct{}{{}})
		// by rows

		l := q.Len()

	case reflect.Map: // t.Append(map[string]interface{}{})
		// by columns

		keys := q.MapKeys()
		l := q.MapIndex(keys[0]).Len()
		for _,k := range keys[1:] {
			if q.MapIndex(k).Len() != l {
				panic("bad count of column elements")
			}
		}

		names := make([]string, len(keys), len(keys))
		columns := make([]reflect.Value, len(names), len(names))

		for i, k :=range keys {
			names[i] = k.String()
			vals := q.MapIndex(k)
			columns[i] = reflect.MakeSlice(vals.Type() /*[]type*/, l, l)
			reflect.Copy(columns[i],vals)
		}

		return &Table{names,columns,l}
	}

	panic("bad argument type")
}

/*
t := tables.New([]struct{Name string; Age int; Rate float}{{"Ivanov",32,1.2},{"Petrov",44,1.5}})
t.Slice(0).Row(0) -> {"Ivanov",32,1.2}
t.Slice(1).Row(0) -> {"Petrov",44,1.5}
t.Slice(0,2).Len() -> 2
t.Slice(1,2).Len() -> 1
*/
func (t *Table) Slice(slice... int) *Table {
	from, to := 0, t.length
	if len(slice) > 0 {
		from = slice[0]
		to = from + 1
	}
	if len(slice) > 1 {
		to = slice[1]
	}
	rv := make([]reflect.Value, len(t.columns), len(t.columns))
	for i,v := range t.columns {
		rv[i] = v.Slice(from,to)
	}
	return &Table{
		names: t.Names(),
		columns: rv,
		length: to-from,
	}
}

/*
t := tables.New([]struct{Name string; Age int; Rate float}{{"Ivanov",32,1.2},{"Petrov",44,1.5}})
t.Only("Age","Rate").Row(0) -> {"Age": 32, "Rate": 1.2}
*/
func (t *Table) Only(column... string) *Table {
	rn := make([]string,len(column),len(column))
	copy(rn,column)
	rv := make([]reflect.Value, len(column), len(column))
	for i,v := range t.columns {
		for j,n := range rn {
			if n == t.names[i] {
				rv[j] = v.Slice(0,t.length)
			}
		}
	}
	return &Table{
		names: rn,
		columns: rv,
		length: t.length,
	}
}

/*
t := tables.Empty()
t = t.Append(struct{Name string; Age int; Rate float32}{"Ivanov",32,1.2})
t = t.Append([]struct{Name string; Age int; Rate: float32}{{"Ivanov",32,1.2},{"Petrov",44,1.5}})
t = t.Append(map[string]interface{}{"Name":[]string{"Ivanov","Petrov"},"Age":[]int{32,44},"Rate":[]float32{1.2,1.5}})

insert empty column
t = t.Append([]struct{Info string}{})
t = t.Append(map[string]interface{}{"Info":[]string{})
*/
func (t *Table) Append(o interface{}) *Table{
	return t.Concat(New(o))
}

func (t *Table) Concat(a *Table) *Table {

	names := t.Names()
	columns := make([]reflect.Value,t.length,t.length)
	copy(columns,t.columns)

	for i,n := range a.names {
		j := IndexOf(n,names)
		if j < 0 {
			col := reflect.MakeSlice(a.columns[i].Type() /*[]type*/, t.length, t.length+a.length)
			col.Set(reflect.AppendSlice(col,a.columns[i]))
			t.names = append(names,n)
			columns = append(columns,col)
		} else {
			columns[j].Set(reflect.AppendSlice(columns[j],a.columns[i]))
		}
	}

	for i, col := range columns {
		if col.Len() < a.length + t.length {
			columns[i].Set(reflect.AppendSlice(col,reflect.MakeSlice(col.Type(),a.length,a.length)))
		}
	}

	return &Table{names,columns,t.length+a.length}
}

/*
t := tables.New([]struct{Name string; Age int; Rate float}{{"Ivanov",32,1.2},{"Petrov",44,1.5}})
t.List(func(r struct{Name string; Age int; Rate float}){
			fmt.Println(r.Name,r.Age,r.Rate)
		},0)
t.List(func(r struct{Age int; Rate float}, i int){
			fmt.Println(i, r.Rate)
		},0,t.Len())
t.List(func(r struct{Name string}, i int){
			fmt.Println(i, r.Name)
		}) // all
*/
func (t *Table) List(r interface{}, i... int) {

}

/*
t := tables.New([]struct{Name string; Age int; Rate float}{{"Ivanov",32,1.2},{"Petrov",44,1.5}})
t.Row(0) -> {Name: "Ivanov", "Age": 32, "Rate", 1.2}
q := t.Sort("Name",tables.DESC)
q.Row(0) -> {Name: "Petrov", "Age": 44, "Rate", 1.5}
*/
func Sort(opt interface{}) *Table {

}

/*
t := table.New([]struct{Name string; Age int; Rate float}{{"Ivanov",32,1.2},{"Petrov",44,1.5}})
q := t.Map(func(r struct{Name string}, i int) struct{Info string}){
			return struct{Info string}{fmt.Sprintf("rec %d for %s", i, r.Name)}
		})
q.Row(0) -> {"Info": "rec 0 for Ivanov"}
*/
func (t *Table) Map(r interface{}) *Table{

}

/*
t := table.New([]struct{Name string; Age int; Rate float}{{"Ivanov",32,1.2},{"Petrov",44,1.5}})
q := t.Update(func(r struct{Name string}, i int) struct{Info string}{
			return struct{Info string}{fmt.Sprintf("rec %d for %s", i, r.Name)}
		})
q.Row(0) -> {Name: "Ivanov", "Age": 32, "Rate", 1.2, "Info": "rec 0 for Ivanov"}
*/
func (t *Table) Update(r interface{}) *Table{

}

func (t *Table) Filter(r interface{}) *Table{

}

/*
t := tables.New([]struct{Name string; Age int; Rate float}{{"Ivanov",32,1.2},{"Ivanov",33,1.3},{"Petrov",44,1.5}})
t.Len() -> 3
q := t.Reduce(func(a, b struct{Age int})(r struct{Age int}){
			r.Age = func(a,b int)int{ if a.Age >= b.Age {return a.Age} return b.Age }(a,b)
			return
		}, "Name")
q.Len() -> 2
// "Name" is grouping field so it's retained, all other fields not presented in result will skipped
q.Row(0) -> {"Name":"Ivanov", "Age": 33}
q.Row(1) -> {"Name":"Petrov", "Age": 44}
*/
func (t *Table) Reduce(f interface{}, groupby... string) *Table{

}

/*
t := table.New([]struct{Name string; Age int; Rate float}{{"Ivanov",32,1.2},{"Petrov",44,1.5}})
t.Row(0)["Name"].String() -> "Ivanov"
*/
func (t *Table) Row(row int) map[string]reflect.Value {
	r := map[string]reflect.Value{}
	for i,n := range t.names {
		// prevent to change value in slice via returned reflect.Value
		r[n] = reflect.ValueOf(t.columns[i].Index(row).Interface())
	}
	return r
}

