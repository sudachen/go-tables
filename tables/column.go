package tables

import (
	"fmt"
	"reflect"
)

type Column struct {
	column reflect.Value
}

/*
t := table.New([]struct{Name string; Age int; Rate float}{{"Ivanov",42,1.2},{"Petrov",42,1.5}})
t.Col("Name").String(0) -> "Ivanov"
t.Col("Name").Extract().([]string)[0] -> "Ivanov"
t.Col("Name").Len() -> 2
t.Col("Age").Unique().Extract().([]int) -> {42}
*/
func (t *Table) Col(column string) *Column {
	for i,n :=range t.names {
		if n == column {
			return &Column{column: t.columns[i]}
		}
	}
	panic("there is not column with name "+column)
}

/*
t := table.New([]struct{Name string; Age int; Rate float}{{"Ivanov",32,1.2},{"Petrov",44,1.5}})
t.Col("Name").String(0) -> "Ivanov"
*/
func (c *Column) String(row int) string {
	v := c.column.Index(row)
	if v.Kind() == reflect.String {
		return v.Interface().(string)
	}
	return fmt.Sprint(v.Interface())
}

var stringType = reflect.TypeOf("")

/*
t := table.New([]struct{Name string; Age int; Rate float}{{"Ivanov",32,1.2},{"Petrov",44,1.5}})
t.Col("Name").Strings() -> {"Ivanov","Petrow"}
*/
func (c *Column) Strings() []string {
	return c.ExtractAs(stringType).([]string)
}

var intType = reflect.TypeOf(int(0))

/*
t := table.New([]struct{Name string; Age int; Rate float}{{"Ivanov",32,1.2},{"Petrov",44,1.5}})
t.Col("Age").Int(0) -> 32
*/
func (c *Column) Int(row int) int {
	v := c.column.Index(row)
	return Convert(v,intType).(int)
}

/*
t := table.New([]struct{Name string; Age int; Rate float}{{"Ivanov",32,1.2},{"Petrov",44,1.5}})
t.Col("Age").Ints(0) -> {32,44}
*/
func (c *Column) Ints() []int {
	return c.ExtractAs(intType).([]int)
}

var floatType = reflect.TypeOf(float32(0))

/*
t := table.New([]struct{Name string; Age int; Rate float}{{"Ivanov",32,1.2},{"Petrov",44,1.5}})
t.Col("Rate").Float(0) -> 1.2
*/
func (c *Column) Float(row int) float32 {
	v := c.column.Index(row)
	return Convert(v,floatType).(float32)
}

/*
t := table.New([]struct{Name string; Age int; Rate float}{{"Ivanov",32,1.2},{"Petrov",44,1.5}})
t.Col("Rate").Floats() -> {1.2,1.5}
*/
func (c *Column) Floats() []float32 {
	return c.ExtractAs(floatType).([]float32)
}

/*
t := table.New([]struct{Name string; Age int; Rate float}{{"Ivanov",32,1.2},{"Petrov",44,1.5}})
t.Col("Rate").Interface(0).(float32) -> 1.2
*/
func (c *Column) Interface(row int) interface{} {
	return c.column.Index(row).Interface()
}

/*
t := table.New([]struct{Name string; Age int; Rate float}{{"Ivanov",32,1.2},{"Petrov",44,1.5}})
t.Col("Age").ExtractAs(reflect.TypeOf("")).([]string)[0] -> "32"
t.Col("Rate").ExtractAs(reflect.TypeOf(int(0))).([]int)[0] -> 1
*/
func (c *Column) ExtractAs(tp reflect.Type) interface{} {
	if c.column.Type().Elem() == tp {
		l := c.column.Len()
		r := reflect.MakeSlice(c.column.Type(),l,l)
		reflect.Copy(r,c.column)
		return r.Interface()
	} else {
		return Convert(c.column,tp)
	}
}

/*
t := table.New([]struct{Name string; Age int; Rate float}{{"Ivanov",32,1.2},{"Petrov",44,1.5}})
t.Col("Name").Inspect().([]string)[0] -> "Ivanov"
t.Col("Age").Inspect().([]int)[0] -> 32
t.Col("Rate").Inspect().([]float32)[0] -> 1.2
*/
func (c *Column) Inspect() interface{} {
	return c.column.Interface()
}

/*
t := tables.New([]struct{Name string}{{"Ivanov"}})
c1 := t.Col("Name")
t.Append([]struct{Name string}{{"Petrov"}})
c2 := t.Col("Name")
c1.Len() -> 1
c2.Len() -> 2
*/
func (c *Column) Len() int {
	return c.column.Len()
}

/*
t := tables.New([]struct{Name string}{{"Ivanov"}})
u1 := t.Col("Name").Unique()
t = t.Append([]struct{Name string}{{"Petrov"},{"Petrov"}})
u2 := t.Col("Name").Unique()
u1.Unique().Inspect() -> {}
u2.Unique().Len() -> 2
*/
func (c *Column) Unique() *Column {
	v := reflect.ValueOf(true)
	m := reflect.MakeMap(reflect.MapOf(c.column.Type().Elem(),v.Type()))
	r := reflect.MakeSlice(c.column.Type(),0,0)
	for i:=0; i < c.column.Len(); i++ {
		x := c.column.Index(i)
		q := m.MapIndex(x)
		if !q.Bool() {
			r = reflect.Append(r,x)
			q.Set(v)
		}
	}
	return &Column{r}
}
