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
	q2 := q.Append([]struct{ Name string }{{"Hello"}})
	q3 := q2.Append([]struct{ Age int }{{0}})
	q4 := q3.Append([]struct {
		Name string
		Age  int
	}{{"Hello", 0}})

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
}
