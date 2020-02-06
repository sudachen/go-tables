package tests

import (
	"github.com/sudachen/go-tables/tables"
	"github.com/sudachen/go-tables/util"
	"gotest.tools/assert"
	"testing"
)

func PrepareTable(t *testing.T) *tables.Table {
	q := tables.New([]struct {
		Name string
		Age  int
		Rate float32
	}{
		{"Ivanov", 32, 1.2},
		{"Petrov", 44, 1.5}})
	assert.DeepEqual(t, q.Names(), []string{"Name", "Age", "Rate"})
	assert.Assert(t, q.Len() == 2)
	assert.DeepEqual(t, util.MapInterface(q.Row(0)),
		map[string]interface{}{
			"Name": "Ivanov",
			"Age":  32,
			"Rate": float32(1.2),
		})
	assert.DeepEqual(t, util.MapInterface(q.Row(1)),
		map[string]interface{}{
			"Name": "Petrov",
			"Age":  44,
			"Rate": float32(1.5),
		})

	return q
}
