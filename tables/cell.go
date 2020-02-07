package tables

import (
	"fmt"
	"github.com/sudachen/go-tables/util"
	"reflect"
)

type Cell struct {
	reflect.Value
}

func (c Cell) String() string {
	if c.Kind() == reflect.String {
		return c.Interface().(string)
	}
	return fmt.Sprint(c.Interface())
}

func (c Cell) Int() int {
	return util.Convert(c.Value, intType).(int)
}

func (c Cell) Int8() int8 {
	return util.Convert(c.Value, int8Type).(int8)
}

func (c Cell) Int16() int16 {
	return util.Convert(c.Value, int16Type).(int16)
}

func (c Cell) Int32() int32 {
	return util.Convert(c.Value, int32Type).(int32)
}

func (c Cell) Int64() int64 {
	return util.Convert(c.Value, int64Type).(int64)
}

func (c Cell) Uint() uint {
	return util.Convert(c.Value, uintType).(uint)
}

func (c Cell) Uint8() uint8 {
	return util.Convert(c.Value, uint8Type).(uint8)
}

func (c Cell) Uint16() uint16 {
	return util.Convert(c.Value, uint16Type).(uint16)
}

func (c Cell) Uint32() uint32 {
	return util.Convert(c.Value, uint32Type).(uint32)
}

func (c Cell) Uint64() uint64 {
	return util.Convert(c.Value, uint64Type).(uint64)
}

func (c Cell) Float() float32 {
	return util.Convert(c.Value, floatType).(float32)
}

func (c Cell) Float64() float64 {
	return util.Convert(c.Value, float64Type).(float64)
}
