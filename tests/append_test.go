package tests

import (
	"github.com/sudachen/go-tables/tables"
	"gotest.tools/assert"
	"gotest.tools/assert/cmp"
	"testing"
)

func Test_Append0(t *testing.T) {
	q := tables.Empty()
	assert.Assert(t, cmp.Panics(func() { q.Append([]int{0}) }))
	assert.Assert(t, cmp.Panics(func() { q.Append(0) }))
	assert.Assert(t, cmp.Panics(func() {
		q.Append(map[string]interface{}{
			"Name": []string{"a", "b"},
			"Age":  []int{0},
		})
	}))
	q2 := q.Append([]struct{ Name string }{})
	assert.Assert(t, q.Len() == q2.Len())
	assert.Assert(t, cmp.Panics(func() { q2.Append(struct{ Name int }{0}) }))
	assert.Assert(t, q.Append([]struct{ Age int }{{0}}).Len() == q.Len()+1)
	assert.Assert(t, q.Append([]struct{ Tall int }{{0}}).Len() == q.Len()+1)
}
