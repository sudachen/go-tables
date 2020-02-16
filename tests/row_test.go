package tests

import (
	"gotest.tools/assert"
	"gotest.tools/assert/cmp"
	"testing"
)

func Test_Row1(t *testing.T) {
	q := TrTable()
	r := TR{}
	for i, v := range trList {
		q.Fetch(i, &r)
		assert.DeepEqual(t, r, v)
	}
}

func Test_Row2(t *testing.T) {
	q := TrTable()
	r := struct{ A int }{}
	assert.Assert(t, cmp.Panics(func() {
		q.Fetch(0, &r)
	}))
	x := map[int]interface{}{}
	assert.Assert(t, cmp.Panics(func() {
		q.Fetch(0, &x)
	}))
}
