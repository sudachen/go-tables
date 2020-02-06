package tables

import (
	"fmt"
	"github.com/sudachen/go-tables/util"
	"reflect"
)

type Column struct {
	column reflect.Value
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
var int8Type = reflect.TypeOf(int8(0))
var int16Type = reflect.TypeOf(int16(0))
var int32Type = reflect.TypeOf(int32(0))
var int64Type = reflect.TypeOf(int64(0))

var uintType = reflect.TypeOf(uint(0))
var uint8Type = reflect.TypeOf(uint8(0))
var uint16Type = reflect.TypeOf(uint16(0))
var uint32Type = reflect.TypeOf(uint32(0))
var uint64Type = reflect.TypeOf(uint64(0))

/*
t := table.New([]struct{Name string; Age int; Rate float}{{"Ivanov",32,1.2},{"Petrov",44,1.5}})
t.Col("Age").Int(0) -> 32
*/
func (c *Column) Int(row int) int {
	v := c.column.Index(row)
	return util.Convert(v, intType).(int)
}

/*
t := table.New([]struct{Name string; Age int; Rate float}{{"Ivanov",32,1.2},{"Petrov",44,1.5}})
t.Col("Age").Int8(0) -> 32
*/
func (c *Column) Int8(row int) int8 {
	v := c.column.Index(row)
	return util.Convert(v, int8Type).(int8)
}

/*
t := table.New([]struct{Name string; Age int; Rate float}{{"Ivanov",32,1.2},{"Petrov",44,1.5}})
t.Col("Age").Int16(0) -> 32
*/
func (c *Column) Int16(row int) int16 {
	v := c.column.Index(row)
	return util.Convert(v, int16Type).(int16)
}

/*
t := table.New([]struct{Name string; Age int; Rate float}{{"Ivanov",32,1.2},{"Petrov",44,1.5}})
t.Col("Age").Int64(0) -> 32
*/
func (c *Column) Int32(row int) int32 {
	v := c.column.Index(row)
	return util.Convert(v, int32Type).(int32)
}

/*
t := table.New([]struct{Name string; Age int; Rate float}{{"Ivanov",32,1.2},{"Petrov",44,1.5}})
t.Col("Age").Int64(0) -> 32
*/
func (c *Column) Int64(row int) int64 {
	v := c.column.Index(row)
	return util.Convert(v, int64Type).(int64)
}

/*
t := table.New([]struct{Name string; Age int; Rate float}{{"Ivanov",32,1.2},{"Petrov",44,1.5}})
t.Col("Age").Uint(0) -> 32
*/
func (c *Column) Uint(row int) uint {
	v := c.column.Index(row)
	return util.Convert(v, uintType).(uint)
}

/*
t := table.New([]struct{Name string; Age int; Rate float}{{"Ivanov",32,1.2},{"Petrov",44,1.5}})
t.Col("Age").Uint8(0) -> 32
*/
func (c *Column) Uint8(row int) uint8 {
	v := c.column.Index(row)
	return util.Convert(v, uint8Type).(uint8)
}

/*
t := table.New([]struct{Name string; Age int; Rate float}{{"Ivanov",32,1.2},{"Petrov",44,1.5}})
t.Col("Age").Uint16(0) -> 32
*/
func (c *Column) Uint16(row int) uint16 {
	v := c.column.Index(row)
	return util.Convert(v, uint16Type).(uint16)
}

/*
t := table.New([]struct{Name string; Age int; Rate float}{{"Ivanov",32,1.2},{"Petrov",44,1.5}})
t.Col("Age").Uint64(0) -> 32
*/
func (c *Column) Uint32(row int) uint32 {
	v := c.column.Index(row)
	return util.Convert(v, uint32Type).(uint32)
}

/*
t := table.New([]struct{Name string; Age int; Rate float}{{"Ivanov",32,1.2},{"Petrov",44,1.5}})
t.Col("Age").Int64(0) -> 32
*/
func (c *Column) Uint64(row int) uint64 {
	v := c.column.Index(row)
	return util.Convert(v, uint64Type).(uint64)
}

/*
t := table.New([]struct{Name string; Age int; Rate float}{{"Ivanov",32,1.2},{"Petrov",44,1.5}})
t.Col("Age").Ints(0) -> {32,44}
*/
func (c *Column) Ints() []int {
	return c.ExtractAs(intType).([]int)
}

/*
t := table.New([]struct{Name string; Age int; Rate float}{{"Ivanov",32,1.2},{"Petrov",44,1.5}})
t.Col("Age").Ints8(0) -> {32,44}
*/
func (c *Column) Ints8() []int8 {
	return c.ExtractAs(int8Type).([]int8)
}

/*
t := table.New([]struct{Name string; Age int; Rate float}{{"Ivanov",32,1.2},{"Petrov",44,1.5}})
t.Col("Age").Ints16(0) -> {32,44}
*/
func (c *Column) Ints16() []int16 {
	return c.ExtractAs(int16Type).([]int16)
}

/*
t := table.New([]struct{Name string; Age int; Rate float}{{"Ivanov",32,1.2},{"Petrov",44,1.5}})
t.Col("Age").Ints32(0) -> {32,44}
*/
func (c *Column) Ints32() []int32 {
	return c.ExtractAs(int32Type).([]int32)
}

/*
t := table.New([]struct{Name string; Age int; Rate float}{{"Ivanov",32,1.2},{"Petrov",44,1.5}})
t.Col("Age").Ints64(0) -> {32,44}
*/
func (c *Column) Ints64() []int64 {
	return c.ExtractAs(int64Type).([]int64)
}

/*
t := table.New([]struct{Name string; Age int; Rate float}{{"Ivanov",32,1.2},{"Petrov",44,1.5}})
t.Col("Age").Uints(0) -> {32,44}
*/
func (c *Column) Uints() []uint {
	return c.ExtractAs(uintType).([]uint)
}

/*
t := table.New([]struct{Name string; Age int; Rate float}{{"Ivanov",32,1.2},{"Petrov",44,1.5}})
t.Col("Age").Uints8(0) -> {32,44}
*/
func (c *Column) Uints8() []uint8 {
	return c.ExtractAs(uint8Type).([]uint8)
}

/*
t := table.New([]struct{Name string; Age int; Rate float}{{"Ivanov",32,1.2},{"Petrov",44,1.5}})
t.Col("Age").Uints16(0) -> {32,44}
*/
func (c *Column) Uints16() []uint16 {
	return c.ExtractAs(uint16Type).([]uint16)
}

/*
t := table.New([]struct{Name string; Age int; Rate float}{{"Ivanov",32,1.2},{"Petrov",44,1.5}})
t.Col("Age").Uints32(0) -> {32,44}
*/
func (c *Column) Uints32() []uint32 {
	return c.ExtractAs(uint32Type).([]uint32)
}

/*
t := table.New([]struct{Name string; Age int; Rate float}{{"Ivanov",32,1.2},{"Petrov",44,1.5}})
t.Col("Age").Uints64(0) -> {32,44}
*/
func (c *Column) Uints64() []uint64 {
	return c.ExtractAs(uint64Type).([]uint64)
}

var floatType = reflect.TypeOf(float32(0))
var float64Type = reflect.TypeOf(float64(0))

/*
t := table.New([]struct{Name string; Age int; Rate float}{{"Ivanov",32,1.2},{"Petrov",44,1.5}})
t.Col("Rate").Float(0) -> 1.2
*/
func (c *Column) Float(row int) float32 {
	v := c.column.Index(row)
	return util.Convert(v, floatType).(float32)
}

/*
t := table.New([]struct{Name string; Age int; Rate float}{{"Ivanov",32,1.2},{"Petrov",44,1.5}})
t.Col("Rate").Float64(0) -> 1.2
*/
func (c *Column) Float64(row int) float64 {
	v := c.column.Index(row)
	return util.Convert(v, float64Type).(float64)
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
t.Col("Rate").Floats64() -> {1.2,1.5}
*/
func (c *Column) Floats64() []float64 {
	return c.ExtractAs(float64Type).([]float64)
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
		r := reflect.MakeSlice(c.column.Type(), l, l)
		reflect.Copy(r, c.column)
		return r.Interface()
	} else {
		return util.Convert(c.column, tp)
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
	m := reflect.MakeMap(reflect.MapOf(c.column.Type().Elem(), v.Type()))
	r := reflect.MakeSlice(c.column.Type(), 0, 0)
	for i := 0; i < c.column.Len(); i++ {
		x := c.column.Index(i)
		q := m.MapIndex(x)
		if !q.IsValid() {
			r = reflect.Append(r, x)
			m.SetMapIndex(x, v)
		}
	}
	return &Column{r}
}
