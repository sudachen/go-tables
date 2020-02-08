package tests

import (
	"github.com/sudachen/go-tables/tables"
	"gotest.tools/assert"
	"gotest.tools/assert/cmp"
	"testing"
)

func Test_Collect(t *testing.T) {
	q := tables.New(trList)
	assertTrData(t, q)
	r := q.Collect(TR{}).([]TR)
	assert.DeepEqual(t, trList, r)
	r = q.Collect(&TR{}).([]TR)
	assert.DeepEqual(t, trList, r)
	assert.Assert(t, cmp.Panics(func() {
		r = q.Collect(false).([]TR)
	}))
}
