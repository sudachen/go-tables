package tests

import (
	"gotest.tools/assert"
	"gotest.tools/assert/cmp"
	"testing"
)

func Test_ColumnString(t *testing.T) {
	q := PrepareTable(t)

	assert.DeepEqual(t, q.Col("Name").String(0), "Ivanov")
	assert.DeepEqual(t, q.Col("Name").String(1), "Petrov")
	assert.DeepEqual(t, q.Col("Age").String(0), "32")
	assert.DeepEqual(t, q.Col("Age").String(1), "44")
	assert.DeepEqual(t, q.Col("Rate").String(0), "1.2")
	assert.DeepEqual(t, q.Col("Rate").String(1), "1.5")

	assert.Assert(t, cmp.Panics(func(){q.Col("name")}))
}

func Test_ColumnStrings(t *testing.T) {
	q := PrepareTable(t)

	assert.DeepEqual(t, q.Col("Name").Strings(),[]string{"Ivanov","Petrov"})
	assert.DeepEqual(t, q.Col("Age").Strings(),[]string{"32","44"})
	assert.DeepEqual(t, q.Col("Rate").Strings(),[]string{"1.2","1.5"})
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

	assert.Assert(t, cmp.Panics(func(){q.Col("age")}))
}

func Test_ColumnInts(t *testing.T) {
	q := PrepareTable(t)

	assert.DeepEqual(t, q.Col("Age").Ints(),[]int{32,44})
	assert.DeepEqual(t, q.Col("Age").Ints8(),[]int8{32,44})
	assert.DeepEqual(t, q.Col("Age").Ints16(),[]int16{32,44})
	assert.DeepEqual(t, q.Col("Age").Ints32(),[]int32{32,44})
	assert.DeepEqual(t, q.Col("Age").Ints64(),[]int64{32,44})
	assert.DeepEqual(t, q.Col("Age").Uints(),[]uint{32,44})
	assert.DeepEqual(t, q.Col("Age").Uints8(),[]uint8{32,44})
	assert.DeepEqual(t, q.Col("Age").Uints16(),[]uint16{32,44})
	assert.DeepEqual(t, q.Col("Age").Uints32(),[]uint32{32,44})
	assert.DeepEqual(t, q.Col("Age").Uints64(),[]uint64{32,44})
}

func Test_ColumnFloat(t *testing.T) {
	q := PrepareTable(t)

	assert.DeepEqual(t, q.Col("Rate").Float(0), float32(1.2))
	assert.DeepEqual(t, q.Col("Rate").Float(1), float32(1.5))
	assert.DeepEqual(t, q.Col("Rate").Float64(0), float64(float32(1.2)))
	assert.DeepEqual(t, q.Col("Rate").Float64(1), float64(float32(1.5)))

	assert.DeepEqual(t, q.Col("Age").Float(0), float32(32))
	assert.DeepEqual(t, q.Col("Age").Float64(0), float64(32))

	assert.Assert(t, cmp.Panics(func(){q.Col("rate")}))
}

func Test_ColumnFloats(t *testing.T) {
	q := PrepareTable(t)

	assert.DeepEqual(t, q.Col("Age").Floats(),[]float32{32,44})
	assert.DeepEqual(t, q.Col("Rate").Floats(),[]float32{1.2,1.5})
	assert.DeepEqual(t, q.Col("Age").Floats64(),[]float64{32,44})
	assert.DeepEqual(t, q.Col("Rate").Floats64(),[]float64{float64(float32(1.2)),float64(float32(1.5))})
}

func Test_ColumnLen(t *testing.T) {
	q := PrepareTable(t)

	assert.Assert(t,q.Len() == 2)

	q2 := q.Append(struct{Name string; Age int}{"Sidorov",55})

	assert.Assert(t, q.Len() == 2)
	assert.Assert(t, q2.Len() == 3)
}

func Test_ColumnUnique(t *testing.T) {
	q := PrepareTable(t)
	assert.DeepEqual(t, q.Col("Name").Unique().Strings(), []string{"Ivanov","Petrov"})

	q2 := q.Append([]struct{Name string; Age int}{
						{"Sidorov",55},{"Ivanov",55}})

	assert.Assert(t, q2.Len() == 4)
	assert.DeepEqual(t, q2.Col("Name").Unique().Strings(), []string{"Ivanov","Petrov","Sidorov"})
	assert.DeepEqual(t, q2.Col("Age").Unique().Ints(), []int{32,44,55})
	assert.DeepEqual(t, q2.Col("Rate").Unique().Floats(), []float32{1.2,1.5,0})
}
