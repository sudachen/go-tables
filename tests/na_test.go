package tests

import (
	"github.com/sudachen/go-tables/tables"
	"gotest.tools/assert"
	"gotest.tools/assert/cmp"
	"testing"
)

func Test_NA1(t *testing.T) {
	q := tables.Empty()
	q2 := q.Append([]struct{ Name string }{{"Hello"}})
	q3 := q2.Append([]struct{ Age int }{})
	assert.Assert(t, q3.Col("Name").Len() == 1)
	assert.Assert(t, q3.Col("Age").Len() == 1)
	assert.Assert(t, !q3.Col("Name").Na(0))
	assert.Assert(t, q3.Col("Age").Na(0))
}

func Test_NA2(t *testing.T) {
	q := tables.Empty()
	q2 := q.Append([]struct {
		Name string
		Rate float32
	}{{"Hello", 1.2}})
	q3 := q2.Append([]struct {
		Age  int
		Rate float32
	}{{0, 0}})

	q4 := q3.Append([]struct {
		Name string
		Age  int
		Rate float32
	}{{"Hello", 0, 0}})

	q5 := q4.FillNa(struct {
		Name string
		Age  int
	}{"Empty", -1})
	q6 := q4.FillNa(map[string]interface{}{"Name": "Empty", "Age": -1})
	q7 := q4.FillNa(map[string]interface{}{"Rate": 0})
	q8 := q4.FillNa(map[string]interface{}{"Name": 0, "Age": -1.0})

	assert.Assert(t, q4.Col("Name").Len() == 3)
	assert.Assert(t, q4.Col("Age").Len() == 3)

	assert.Assert(t, !q4.Col("Name").Na(0))
	assert.Assert(t, q4.Col("Age").Na(0))
	assert.Assert(t, q4.Col("Name").Na(1))
	assert.Assert(t, !q4.Col("Age").Na(1))
	assert.Assert(t, !q4.Col("Name").Na(2))
	assert.Assert(t, !q4.Col("Age").Na(2))

	assert.Assert(t, q4.DropNa().Len() == 1)
	assert.Assert(t, q4.DropNa("Name").Len() == 2)
	assert.Assert(t, q4.DropNa("Age").Len() == 2)
	assert.Assert(t, q2.DropNa().Len() == 1)

	assert.Assert(t, cmp.Panics(func() {
		q2.DropNa("pigs")
	}))

	assert.Assert(t, cmp.Panics(func() {
		q4.FillNa("pigs")
	}))

	assert.Assert(t, cmp.Panics(func() {
		q4.FillNa(struct{ Name1 string }{})
	}))

	assert.Assert(t, q5.DropNa().Len() == 3)
	assert.Assert(t, !q5.Col("Name").Na(0))
	assert.Assert(t, !q5.Col("Age").Na(0))
	assert.Assert(t, !q5.Col("Name").Na(1))
	assert.Assert(t, !q5.Col("Age").Na(1))
	assert.Assert(t, !q5.Col("Name").Na(2))
	assert.Assert(t, !q5.Col("Age").Na(2))
	assert.Assert(t, q5.Col("Age").Int(0) == -1)
	assert.Assert(t, q5.Col("Name").String(1) == "Empty")

	assert.Assert(t, q6.DropNa().Len() == 3)
	assert.Assert(t, !q6.Col("Name").Na(0))
	assert.Assert(t, !q6.Col("Age").Na(0))
	assert.Assert(t, !q6.Col("Name").Na(1))
	assert.Assert(t, !q6.Col("Age").Na(1))
	assert.Assert(t, !q6.Col("Name").Na(2))
	assert.Assert(t, !q6.Col("Age").Na(2))
	assert.Assert(t, q6.Col("Age").Int(0) == -1)
	assert.Assert(t, q6.Col("Name").String(1) == "Empty")

	assert.Assert(t, q7.DropNa().Len() == 1)
	assert.Assert(t, q8.DropNa().Len() == 3)
	assert.Assert(t, q8.Col("Age").Int(0) == -1)
	assert.Assert(t, q8.Col("Name").String(1) == "0")
}
