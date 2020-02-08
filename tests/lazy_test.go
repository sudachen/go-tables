package tests

import (
	"github.com/sudachen/go-fp/lazy"
	"github.com/sudachen/go-tables/tables"
	"github.com/sudachen/go-tables/util"
	"gotest.tools/assert"
	"gotest.tools/assert/cmp"
	"testing"
)

func Test_Lazy1(t *testing.T) {
	q := tables.FillUp(lazy.New(trList))
	assertTrData(t, q)

	q = tables.ConqFillUp(lazy.New(trList), 6)
	assertTrData(t, q)
}

func Test_Lazy2(t *testing.T) {
	q := tables.New(trList)
	r := lazy.New(trList).Filter(func(r TR) bool { return r.Age > 30 }).Collect().([]TR)
	q2 := tables.ConqFillUp(q.Lazy(func(r TR) bool { return r.Age > 30 }), 6)
	for i, v := range r {
		assert.DeepEqual(t, util.MapInterface(q2.Row(i)),
			map[string]interface{}{
				"Name": v.Name,
				"Age":  v.Age,
				"Rate": v.Rate,
			})
		assert.Assert(t, v.Age > 30)
	}
}

func Test_Lazy3(t *testing.T) {
	q := tables.New(trList)
	q2 := tables.ConqFillUp(q.Lazy(func(r TR) TR { return r }), 6)
	assertTrData(t, q2)
}

func Test_Lazy4(t *testing.T) {
	q := tables.New(trList)
	q2 := tables.ConqFillUp(q.Lazy(TR{}), 6)
	assertTrData(t, q2)
	q2 = tables.ConqFillUp(q.Lazy(&TR{}), 6)
	assertTrData(t, q2)
}

func Test_Lazy5(t *testing.T) {
	q := tables.New(trList)
	assert.Assert(t, cmp.Panics(func() {
		q.Lazy(func(int) int { return 0 })
	}))
}
