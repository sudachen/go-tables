package tables

import (
	"github.com/sudachen/go-fp/lazy"
	"reflect"
	"sync"
)

/*
Lazy creates new lazy transformation stream from the table and empty struct or some transformation function
*/
func (t *Table) Lazy(x interface{}) *lazy.Stream {
	v := reflect.ValueOf(x)
	vt := v.Type()
	if v.Kind() == reflect.Struct || (v.Kind() == reflect.Ptr && v.Elem().Kind() == reflect.Struct) {
		if vt.Kind() == reflect.Ptr {
			vt = vt.Elem()
		}
		getf := func(index int, ctx interface{}) reflect.Value {
			ctx.(*lazy.Counter).IncAwait(index)
			if index < t.Len() {
				return t.MakeRow(index, vt)
			}
			return reflect.ValueOf(false)
		}
		return &lazy.Stream{Get: getf, Ctx: &lazy.Counter{0}, Tp: vt}
	} else if v.Kind() == reflect.Func &&
		vt.NumIn() == 1 && vt.NumOut() == 1 &&
		vt.In(0).Kind() == reflect.Struct &&
		(vt.Out(0).Kind() == reflect.Struct || vt.Out(0).Kind() == reflect.Bool) {
		ti := vt.In(0)
		to := vt.Out(0)
		isFilter := to.Kind() == reflect.Bool
		getf := func(index int, ctx interface{}) reflect.Value {
			ctx.(*lazy.Counter).IncAwait(index)
			if index < t.Len() {
				q := []reflect.Value{t.MakeRow(index, ti)}
				r := v.Call(q)
				if isFilter {
					if r[0].Bool() {
						return q[0]
					}
					return reflect.ValueOf(true)
				}
				return r[0]
			}
			return reflect.ValueOf(false)
		}
		if isFilter {
			return &lazy.Stream{Get: getf, Ctx: &lazy.Counter{0}, Tp: ti}
		}
		return &lazy.Stream{Get: getf, Ctx: &lazy.Counter{0}, Tp: to}
	} else {
		panic("only struct{...}, func(struct{...})struct{...} or func(struct{...})bool are allowed as an argument")
	}
}

/*
FillUp fills new table from the transformation source
*/
func FillUp(z *lazy.Stream) *Table {
	c := reflect.MakeChan(z.Tp, 0)
	go func() {
		index := 0
		for {
			v := z.Next(index)
			index++
			if v.Kind() == reflect.Bool {
				if !v.Bool() {
					break
				}
			} else {
				c.Send(v)
			}
		}
		c.Close()
	}()
	return New(c)
}

/*
ConqFillUp fills new table from the transformation source concurrently
*/
func ConqFillUp(z *lazy.Stream, concurrency int) *Table {
	index := &lazy.Counter{0}
	c := reflect.MakeChan(z.Tp, concurrency)
	gw := sync.WaitGroup{}
	gw.Add(concurrency)
	for i := 0; i < concurrency; i++ {
		go func() {
			defer gw.Done()
			for {
				v := z.Next(index.GetInc())
				if v.Kind() == reflect.Bool {
					if !v.Bool() {
						break
					}
				} else {
					c.Send(v)
				}
			}
		}()
	}
	go func() {
		gw.Wait()
		c.Close()
	}()
	return New(c)
}