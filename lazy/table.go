package lazy

import (
	"github.com/sudachen/go-tables/tables"
	"reflect"
)

func Lazy(t *tables.Table, x interface{}) *Source {
	v := reflect.ValueOf(x)
	vt := v.Type()
	if v.Kind() == reflect.Struct || (v.Kind() == reflect.Ptr && v.Elem().Kind() == reflect.Struct) {
		if vt.Kind() == reflect.Ptr {
			vt = vt.Elem()
		}
		f := func(ctx interface{}) reflect.Value {
			no := ctx.(*int)
			i := *no
			if i < t.Len() {
				*no++
				return t.MakeRow(i, vt)
			}
			return reflect.ValueOf(false)
		}
		return &Source{new(int), vt, f}
	} else if v.Kind() == reflect.Func &&
		vt.NumIn() == 1 && vt.NumOut() == 1 &&
		vt.In(0).Kind() == reflect.Struct &&
		(vt.Out(0).Kind() == reflect.Struct || vt.Out(0).Kind() == reflect.Bool) {
		ti := vt.In(0)
		to := vt.Out(0)
		isFilter := to.Kind() == reflect.Bool
		f := func(ctx interface{}) reflect.Value {
			no := ctx.(*int)
			i := *no
			if i < t.Len() {
				*no++
				q := []reflect.Value{t.MakeRow(i, ti)}
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
			return &Source{new(int), ti, f}
		}
		return &Source{new(int), to, f}
	} else {
		panic("only struct{...}, func(struct{...})struct{...} or func(struct{...})bool are allowed as an argument")
	}
}

func (z *Source) FillUp() *tables.Table {
	c := reflect.MakeChan(z.Tp, 0)
	go func() {
		for {
			v := z.Next(z.Ctx)
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
	return tables.New(c)
}
