package tests

import (
	"github.com/sudachen/go-tables/tables"
	"github.com/sudachen/go-tables/util"
	"gotest.tools/assert"
	"testing"
)

func Test_New0(t *testing.T) {
	q := tables.New([]struct {
		Name string
		Age  int
		Rate float32
	}{})
	assert.DeepEqual(t, q.Names(), []string{"Name", "Age", "Rate"})
	assert.Assert(t, q.Len() == 0)
}

func Test_New1(t *testing.T) {
	q := tables.New([]struct {
		Name string
		Age  int
		Rate float32
	}{{"Ivanov", 32, 1.2}})
	assert.DeepEqual(t, q.Names(), []string{"Name", "Age", "Rate"})
	assert.Assert(t, q.Len() == 1)
	assert.DeepEqual(t, util.MapInterface(q.Row(0)),
		map[string]interface{}{
			"Name": "Ivanov",
			"Age":  32,
			"Rate": float32(1.2),
		})
}

func Test_New2(t *testing.T) {
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
}

func Test_New3(t *testing.T) {
	q := tables.New(map[string]interface{}{
		"Name": []string{"Ivanov", "Petrov"},
		"Age":  []int{32, 44},
		"Rate": []float32{1.2, 1.5}})
	assert.DeepEqual(t, q.Names(), []string{"Age", "Name", "Rate"})
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
}

func Test_New4(t *testing.T) {
	type R struct {
		Name string
		Age  int
		Rate float32
	}
	c := make(chan R)
	go func() {
		c <- R{"Ivanov", 32, 1.2}
		c <- R{"Petrov", 44, 1.5}
		close(c)
	}()
	q := tables.New(c)
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
}
