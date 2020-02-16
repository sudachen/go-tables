package tests

import (
	"fmt"
	"github.com/sudachen/go-tables/tables"
	"github.com/sudachen/go-tables/util"
	"gotest.tools/assert"
	"gotest.tools/assert/cmp"
	"math/rand"
	"reflect"
	"testing"
	"time"
)

func Test_ColumnString(t *testing.T) {
	q := PrepareTable(t)

	assert.DeepEqual(t, q.Col("Name").String(0), "Ivanov")
	assert.DeepEqual(t, q.Col("Name").String(1), "Petrov")
	assert.DeepEqual(t, q.Col("Age").String(0), "32")
	assert.DeepEqual(t, q.Col("Age").String(1), "44")
	assert.DeepEqual(t, q.Col("Rate").String(0), "1.2")
	assert.DeepEqual(t, q.Col("Rate").String(1), "1.5")

	assert.Assert(t, cmp.Panics(func() { q.Col("name") }))
}

func Test_ColumnStrings(t *testing.T) {
	q := PrepareTable(t)

	assert.DeepEqual(t, q.Col("Name").Strings(), []string{"Ivanov", "Petrov"})
	assert.DeepEqual(t, q.Col("Age").Strings(), []string{"32", "44"})
	assert.DeepEqual(t, q.Col("Rate").Strings(), []string{"1.2", "1.5"})
}

func Test_ColumnInt(t *testing.T) {
	q := PrepareTable(t)

	assert.DeepEqual(t, q.Col("Age").Int(0), 32)
	assert.DeepEqual(t, q.Col("Age").Int(1), 44)
	assert.DeepEqual(t, q.Col("Rate").Int(0), 1)
	assert.DeepEqual(t, q.Col("Rate").Int(1), 1)

	assert.DeepEqual(t, q.Col("Age").Int8(0), int8(32))
	assert.DeepEqual(t, q.Col("Rate").Int8(0), int8(1))

	assert.DeepEqual(t, q.Col("Age").Int16(0), int16(32))
	assert.DeepEqual(t, q.Col("Rate").Int16(0), int16(1))

	assert.DeepEqual(t, q.Col("Age").Int32(0), int32(32))
	assert.DeepEqual(t, q.Col("Rate").Int32(0), int32(1))

	assert.DeepEqual(t, q.Col("Age").Int64(0), int64(32))
	assert.DeepEqual(t, q.Col("Rate").Int64(0), int64(1))

	assert.DeepEqual(t, q.Col("Age").Uint(0), uint(32))
	assert.DeepEqual(t, q.Col("Rate").Uint(0), uint(1))

	assert.DeepEqual(t, q.Col("Age").Uint8(0), uint8(32))
	assert.DeepEqual(t, q.Col("Rate").Uint8(0), uint8(1))

	assert.DeepEqual(t, q.Col("Age").Uint16(0), uint16(32))
	assert.DeepEqual(t, q.Col("Rate").Uint16(0), uint16(1))

	assert.DeepEqual(t, q.Col("Age").Uint32(0), uint32(32))
	assert.DeepEqual(t, q.Col("Rate").Uint32(0), uint32(1))

	assert.DeepEqual(t, q.Col("Age").Uint64(0), uint64(32))
	assert.DeepEqual(t, q.Col("Rate").Uint64(0), uint64(1))

	assert.Assert(t, cmp.Panics(func() { q.Col("age") }))
}

func Test_ColumnInts(t *testing.T) {
	q := PrepareTable(t)

	assert.DeepEqual(t, q.Col("Age").Ints(), []int{32, 44})
	assert.DeepEqual(t, q.Col("Age").Ints8(), []int8{32, 44})
	assert.DeepEqual(t, q.Col("Age").Ints16(), []int16{32, 44})
	assert.DeepEqual(t, q.Col("Age").Ints32(), []int32{32, 44})
	assert.DeepEqual(t, q.Col("Age").Ints64(), []int64{32, 44})
	assert.DeepEqual(t, q.Col("Age").Uints(), []uint{32, 44})
	assert.DeepEqual(t, q.Col("Age").Uints8(), []uint8{32, 44})
	assert.DeepEqual(t, q.Col("Age").Uints16(), []uint16{32, 44})
	assert.DeepEqual(t, q.Col("Age").Uints32(), []uint32{32, 44})
	assert.DeepEqual(t, q.Col("Age").Uints64(), []uint64{32, 44})
}

func Test_ColumnInt2(t *testing.T) {
	q := PrepareTable(t)

	c := q.Col("Age")
	assert.Assert(t, c.Index(0).Int() == 32)
	assert.Assert(t, c.Index(0).Int8() == 32)
	assert.Assert(t, c.Index(0).Int16() == 32)
	assert.Assert(t, c.Index(0).Int32() == 32)
	assert.Assert(t, c.Index(0).Int64() == 32)
	assert.Assert(t, c.Index(0).Uint() == 32)
	assert.Assert(t, c.Index(0).Uint8() == 32)
	assert.Assert(t, c.Index(0).Uint16() == 32)
	assert.Assert(t, c.Index(0).Uint32() == 32)
	assert.Assert(t, c.Index(0).Uint64() == 32)
}

func Test_ColumnFloat(t *testing.T) {
	q := PrepareTable(t)

	assert.DeepEqual(t, q.Col("Rate").Float(0), float32(1.2))
	assert.DeepEqual(t, q.Col("Rate").Float(1), float32(1.5))
	assert.DeepEqual(t, q.Col("Rate").Float64(0), float64(float32(1.2)))
	assert.DeepEqual(t, q.Col("Rate").Float64(1), float64(float32(1.5)))

	assert.DeepEqual(t, q.Col("Age").Float(0), float32(32))
	assert.DeepEqual(t, q.Col("Age").Float64(0), float64(32))

	assert.Assert(t, cmp.Panics(func() { q.Col("rate") }))
}

func Test_ColumnFloats(t *testing.T) {
	q := PrepareTable(t)

	assert.DeepEqual(t, q.Col("Age").Floats(), []float32{32, 44})
	assert.DeepEqual(t, q.Col("Rate").Floats(), []float32{1.2, 1.5})
	assert.DeepEqual(t, q.Col("Age").Floats64(), []float64{32, 44})
	assert.DeepEqual(t, q.Col("Rate").Floats64(), []float64{float64(float32(1.2)), float64(float32(1.5))})
}

func Test_ColumnLen(t *testing.T) {
	q := PrepareTable(t)

	assert.Assert(t, q.Len() == 2)

	q2 := q.Append([]struct {
		Name string
		Age  int
	}{{"Sidorov", 55}})

	assert.Assert(t, q.Len() == 2)
	assert.Assert(t, q2.Len() == 3)
}

func Test_ColumnUnique(t *testing.T) {
	q := PrepareTable(t)
	assert.DeepEqual(t, q.Col("Name").Unique().Strings(), []string{"Ivanov", "Petrov"})

	q2 := q.Append([]struct {
		Name string
		Age  int
	}{{"Sidorov", 55}, {"Ivanov", 55}})

	assert.Assert(t, q2.Len() == 4)
	assert.DeepEqual(t, q2.Col("Name").Unique().Strings(), []string{"Ivanov", "Petrov", "Sidorov"})
	assert.DeepEqual(t, q2.Col("Age").Unique().Ints(), []int{32, 44, 55})
	assert.DeepEqual(t, q2.Col("Rate").Unique().Floats(), []float32{1.2, 1.5, 0})

	q3 := q.Append([]struct {
		Name string
		Tall int
	}{{"Sidorov", 55}, {"Ivanov", 55}})

	assert.DeepEqual(t, q3.Col("Tall").Unique().Strings(), []string{"0", "55"})
}

func Test_Col0(t *testing.T) {
	r := map[int]interface{}{}
	assert.Assert(t, cmp.Panics(func() {
		tables.Col(r)
	}))
}

func Test_Col1(t *testing.T) {
	r := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	c := tables.Col(r)
	assert.Assert(t, c.Len() == len(r))
	assert.Assert(t, c.Type() == reflect.TypeOf(r[0]))
	for i, v := range r {
		assert.Assert(t, c.Int(i) == v)
		assert.Assert(t, c.Interface(i).(int) == v)
		assert.Assert(t, c.Inspect().([]int)[i] == v)
	}
}

func Test_Col2(t *testing.T) {
	r := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(r), func(i, j int) { r[i], r[j] = r[j], r[i] })
	c := tables.Col(r)
	assert.Assert(t, c.Max().Int() == 9)
	assert.Assert(t, c.Min().Int() == 0)
	assert.Assert(t, r[c.MaxIndex()] == 9)
	assert.Assert(t, c.Index(c.MaxIndex()).Int() == 9)
	assert.Assert(t, c.Index(c.MinIndex()).Int() == 0)
}

type ColR3 int

func (a ColR3) Less(b ColR3) bool {
	return b < a
}

func Test_Col3(t *testing.T) {
	r := []ColR3{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(r), func(i, j int) { r[i], r[j] = r[j], r[i] })
	c := tables.Col(r)
	assert.Assert(t, c.Max().Int() == 0)
	assert.Assert(t, c.Min().Int() == 9)
	assert.Assert(t, c.Index(c.MaxIndex()).Int() == 0)
	assert.Assert(t, c.Index(c.MinIndex()).Int() == 9)
}

type ColR4 struct {
	a int
	b uint
	c float64
	e [2]byte
	d string
}

func MkColR4(i int) *ColR4 {
	return &ColR4{
		0,
		uint(1),
		float64(2) * 0.1,
		[2]byte{0, byte(i)},
		fmt.Sprintf("col4:%d", i),
	}
}

func Test_Col4(t *testing.T) {
	r := []*ColR4{MkColR4(0), MkColR4(1), MkColR4(1), MkColR4(2), MkColR4(3), MkColR4(4), MkColR4(5)}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(r), func(i, j int) { r[i], r[j] = r[j], r[i] })
	c := tables.Col(r)
	assert.Assert(t, c.Max().Interface().(*ColR4).d == "col4:5")
	assert.Assert(t, c.Min().Interface().(*ColR4).d == "col4:0")
}

func Test_Col5(t *testing.T) {
	r := []*ColR4{MkColR4(0), MkColR4(1), MkColR4(1), nil, MkColR4(4), MkColR4(5)}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(r), func(i, j int) { r[i], r[j] = r[j], r[i] })
	c := tables.Col(r)
	assert.Assert(t, c.Max().Interface().(*ColR4).d == "col4:5")
	assert.Assert(t, c.Min().Interface().(*ColR4) == nil)
}

func Test_Col6(t *testing.T) {
	r := []*ColR4{MkColR4(0), MkColR4(0), MkColR4(0), MkColR4(1)}
	c := tables.Col(r)
	assert.Assert(t, c.Max().Interface().(*ColR4).d == "col4:1")
	assert.Assert(t, c.Min().Interface().(*ColR4).d == "col4:0")
}

func Test_Less1(t *testing.T) {
	a := map[int]interface{}{0: 0}
	assert.Assert(t, cmp.Panics(func() {
		util.Less(reflect.ValueOf(a), reflect.ValueOf(a))
	}))
	assert.Assert(t, cmp.Panics(func() {
		util.Less(reflect.ValueOf(1), reflect.ValueOf(""))
	}))
	assert.Assert(t, util.Less(reflect.ValueOf([2]int{0, 1}), reflect.ValueOf([2]int{0, 2})))
	assert.Assert(t, cmp.Panics(func() {
		util.Less(reflect.ValueOf([2]int{0, 1}), reflect.ValueOf([1]int{0}))
	}))
}
