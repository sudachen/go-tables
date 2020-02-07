package tables

import (
	"github.com/sudachen/go-tables/util"
	"reflect"
)

type Table struct {
	names   []string
	columns []reflect.Value
	cols    []*Column
	length  int
}

func (t *Table) Len() int {
	return t.length
}

func (t *Table) Names() []string {
	r := make([]string, len(t.names), len(t.names))
	copy(r, t.names)
	return r
}

func Empty() *Table {
	t := &Table{}
	return t
}

/*
tables.New(struct{Name string; Age int; Rate float32}{})
tables.New([]struct{Name string; Age int; Rate float32}{{"Ivanov",32,1.2},{"Petrov",44,1.5}})
tables.New(map[string]interface{}{"Name":[]string{"Ivanov","Petrov"},"Age":[]int{32,44},"Rate":[]float32{1.2,1.5}})

type R struct{Name string; Age int; Rate float32}
c := make(chan R)
go func(){
	c <- R{"Ivanov",32,1.2}
	c <- R{"Petrov",44,1.5}
	close(c)
}()
tables.New(c)
*/
func New(o interface{}) *Table {

	q := reflect.ValueOf(o)

	switch q.Kind() {
	case reflect.Ptr: // New(&struct{}{})
		q = q.Elem()
		fallthrough
	case reflect.Struct: // New(struct{}{})
		s := reflect.MakeSlice(reflect.SliceOf(q.Type()), 1, 1)
		s.Index(0).Set(q)
		q = s
		fallthrough

	case reflect.Slice: // New([]struct{}{{}})
		l := q.Len()
		tp := q.Type().Elem()
		if tp.Kind() != reflect.Struct {
			panic("slice of structures allowed only")
		}
		nl := tp.NumField()
		names := make([]string, 0, nl)
		columns := make([]reflect.Value, 0, nl)
		for i := 0; i < nl; i++ {
			fv := tp.Field(i)
			names = append(names, fv.Name)
			col := reflect.MakeSlice(reflect.SliceOf(fv.Type), l, l)
			columns = append(columns, col)
			for j := 0; j < l; j++ {
				col.Index(j).Set(q.Index(j).Field(i))
			}
		}

		return &Table{
			names:   names,
			columns: columns,
			length:  l,
		}

	case reflect.Chan: // New(chan struct{})
		tp := q.Type().Elem()
		nl := tp.NumField()
		names := make([]string, nl)
		columns := make([]reflect.Value, nl)
		scase := []reflect.SelectCase{{Dir: reflect.SelectRecv, Chan: q}}

		for i := 0; i < nl; i++ {
			fv := tp.Field(i)
			names[i] = fv.Name
			columns[i] = reflect.MakeSlice(reflect.SliceOf(fv.Type), 0, 1)
		}

		length := 0
		for {
			_, v, ok := reflect.Select(scase)
			if !ok {
				break
			}
			for i := 0; i < nl; i++ {
				columns[i] = reflect.Append(columns[i], v.Field(i))
			}
			length++
		}

		return &Table{
			names:   names,
			columns: columns,
			length:  length,
		}

	case reflect.Map: // New(map[string]interface{}{})
		m := o.(map[string]interface{})
		names := util.SortedDictKeys(m)
		columns := make([]reflect.Value, len(names), len(names))
		l := reflect.ValueOf(m[names[0]]).Len()

		for i, n := range names {
			vals := reflect.ValueOf(m[n])
			if vals.Len() != l {
				panic("bad count of elements in column " + n)
			}
			columns[i] = reflect.MakeSlice(vals.Type() /*[]type*/, l, l)
			reflect.Copy(columns[i], vals)
		}

		return &Table{
			names:   names,
			columns: columns,
			length:  l}
	}

	panic("bad argument type")
}

/*
t := tables.New([]struct{Name string; Age int; Rate float32}{{"Ivanov",42,1.2},{"Petrov",42,1.5}})
t.Col("Name").String(0) -> "Ivanov"
t.Col("Name").Len() -> 2
*/
func (t *Table) Col(column string) *Column {
	for i, n := range t.names {
		if n == column {
			if t.cols == nil {
				t.cols = make([]*Column, len(t.names), len(t.names))
			}
			if t.cols[i] == nil {
				c := &Column{column: t.columns[i]}
				t.cols[i] = c
			}
			return t.cols[i]
		}
	}
	panic("there is not column with name " + column)
}

/*
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
t := tables.New([]struct{Name string; Age int; Rate float32}{{"Ivanov",32,1.2},{"Petrov",44,1.5}})
t.Slice(0).Row(0) -> {"Ivanov",32,1.2}
t.Slice(1).Row(0) -> {"Petrov",44,1.5}
t.Slice(0,2).Len() -> 2
t.Slice(1,2).Len() -> 1
*/
func (t *Table) Slice(slice ...int) *Table {
	from, to := 0, t.length
	if len(slice) > 0 {
		from = slice[0]
		to = from + 1
	}
	if len(slice) > 1 {
		to = slice[1]
	}
	rv := make([]reflect.Value, len(t.columns), len(t.columns))
	for i, v := range t.columns {
		rv[i] = v.Slice(from, to)
	}
	return &Table{
		names:   t.Names(),
		columns: rv,
		length:  to - from,
	}
}

/*
t := tables.New([]struct{Name string; Age int; Rate float32}{{"Ivanov",32,1.2},{"Petrov",44,1.5}})
t.Only("Age","Rate").Row(0) -> {"Age": 32, "Rate": 1.2}
*/
func (t *Table) Only(column ...string) *Table {
	rn := make([]string, len(column), len(column))
	copy(rn, column)
	rv := make([]reflect.Value, len(column), len(column))
	for i, v := range t.columns {
		for j, n := range rn {
			if n == t.names[i] {
				rv[j] = v.Slice(0, t.length)
			}
		}
	}
	return &Table{
		names:   rn,
		columns: rv,
		length:  t.length,
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

insert from channel
type R struct{Name string; Age int; Rate float32}
c := make(chan R)
go func(){
	c <- R{"Ivanov",32,1.2}
	c <- R{"Petrov",44,1.5}
	close(c)
}()
t.Append(c)
*/
func (t *Table) Append(o interface{}) *Table {
	return t.Concat(New(o))
}

/*
t1 := tables.New(struct{Name string; Age int; Rate float32}{"Ivanov",32,1.2})
t2 := tables.New(struct{Name string; Age int; Rate float32}{"Petrov",44})
q := t1.Concat(t2)
q.Row(0) -> {"Ivanov",32,1.2}
q.Row(1) -> {"Petrov",44,0}
*/
func (t *Table) Concat(a *Table) *Table {
	names := t.Names()
	columns := make([]reflect.Value, len(names), len(names))
	copy(columns, t.columns)

	for i, n := range a.names {
		j := util.IndexOf(n, names)
		if j < 0 {
			col := reflect.MakeSlice(a.columns[i].Type() /*[]type*/, t.length, t.length+a.length)
			col = reflect.AppendSlice(col, a.columns[i])
			names = append(names, n)
			columns = append(columns, col)
		} else {
			columns[j] = reflect.AppendSlice(columns[j], a.columns[i])
		}
	}

	for i, col := range columns {
		if col.Len() < a.length+t.length {
			columns[i] = reflect.AppendSlice(
				col,
				reflect.MakeSlice(col.Type(), a.length, a.length))
		}
	}

	return &Table{
		names:   names,
		columns: columns,
		length:  t.length + a.length}
}

/*
t := tables.New([]struct{Name string; Age int; Rate float32}{{"Ivanov",32,1.2},{"Petrov",44,1.5}})
q := t.Transform(func(r struct{Name string}, i int) struct{Info string}{
			return struct{Info string}{fmt.Sprintf("rec %d for %s", i, r.Name)}
		})
q.Row(0) -> {Name: "Ivanov", "Age": 32, "Rate", 1.2, "Info": "rec 0 for Ivanov"}
*/
func (t *Table) Transform(f interface{}) *Table {
	t2 := t.Map(f)
	names := t.Names()
	columns := make([]reflect.Value, t.length, t.length)
	copy(columns, t.columns)
	for i, n := range t2.names {
		if j := util.IndexOf(n, names); j >= 0 {
			columns[j] = t2.columns[i]
		} else {
			names = append(names, n)
			columns = append(columns, t2.columns[i])
		}
	}
	return &Table{
		names:   names,
		columns: columns,
		length:  t.length,
	}
}

/*
t := tables.New([]struct{Name string; Age int; Rate float32}{{"Ivanov",32,1.2},{"Petrov",44,1.5}})
t.List(func(r struct{Rate float}, i int){
			fmt.Println(i, r.Rate)
		})
*/
func (t *Table) List(f interface{}) {
	q := reflect.ValueOf(f)
	tp := q.Type().In(0)
	for i := 0; i < t.length; i++ {
		iv := reflect.ValueOf(i)
		q.Call([]reflect.Value{t.MakeRow(i, tp), iv})
	}
}

/*
t := tables.New([]struct{Name string; Age int; Rate float32}{{"Ivanov",32,1.2},{"Petrov",44,1.5}})
q := t.Filter(func(r struct{Age int}) bool{
			return r.Age > 40
		})
q.Row(0) -> {Name: "Petrov", "Age": 44, "Rate", 1.5}
*/
func (t *Table) Filter(f interface{}) *Table {
	q := reflect.ValueOf(f)
	tp := q.Type().In(0)
	idxs := make([]int, 0, t.length)
	for i := 0; i < t.length; i++ {
		iv := reflect.ValueOf(i)
		rv := q.Call([]reflect.Value{t.MakeRow(i, tp), iv})
		if rv[0].Bool() {
			idxs = append(idxs, i)
		}
	}
	return &Table{}
}

/*
t := tables.New([]struct{Name string; Age int; Rate float32}{{"Ivanov",32,1.2},{"Petrov",44,1.5}})
t.Row(0) -> {Name: "Ivanov", "Age": 32, "Rate", 1.2}
q := t.Sort("Name",tables.DESC)
q.Row(0) -> {Name: "Petrov", "Age": 44, "Rate", 1.5}
*/
func (t *Table) Sort(opt interface{}) *Table {
	return nil
}

/*
t := tables.New([]struct{Name string; Age int; Rate float}{{"Ivanov",32,1.2},{"Ivanov",33,1.3},{"Petrov",44,1.5}})
t.Len() -> 3
q := t.Reduce(func(a struct{Age int}, r *struct{Age int}, i int){
			r.Age = func(a,b int)int{ if a.Age >= r.Age {return a.Age} return r.Age }(a,b)
			return
		}, "Name")
q.Len() -> 2
// "Name" is grouping field so it's retained, all other fields not presented in result will skipped
q.Row(0) -> {"Name":"Ivanov", "Age": 33}
q.Row(1) -> {"Name":"Petrov", "Age": 44}
*/
func (t *Table) Reduce(f interface{}, groupby ...string) *Table {
	return nil
}

/*
t := tables.New([]struct{Name string; Age int; Rate float32}{{"Ivanov",32,1.2},{"Petrov",44,1.5}})
q := t.Map(func(r struct{Name string}, i int) struct{Info string}){
			return struct{Info string}{fmt.Sprintf("rec %d for %s", i, r.Name)}
		})
q.Row(0) -> {"Info": "rec 0 for Ivanov"}
*/
func (t *Table) Map(f interface{}) *Table {
	l := 0
	names := make([]string, l, l)
	columns := make([]reflect.Value, l, l)

	return &Table{
		names:   names,
		columns: columns,
		length:  t.length,
	}
}

/*
t := tables.New([]struct{Name string; Age int; Rate float32}{{"Ivanov",32,1.2},{"Petrov",44,1.5}})
c := make(chan struct{Name string})
go t.Sink(c)
for x := range c {
	fmr.Println(x.Name)
}
*/
func (t *Table) Sink(c interface{}) {
}

/*
SinkMap transforms in parallel table rows and send results to a channel
in the order rows presented in the table.

type R struct{Info string}
c := make(chan R)
t.SinkMap(c, func(a struct{Name string; Age int})R{ return R{fmt.Sprint("%s: %d",a.Name,a.Age)} })
for x := range c {
	fmt.Println(x.Info)
}
*/
func (t *Table) SinkMap(c interface{}, f interface{}) {
}

/*
gets required fields as a struct
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
