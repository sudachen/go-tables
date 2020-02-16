package tests

import (
	"github.com/sudachen/go-tables/internal"
	"github.com/sudachen/go-tables/util"
	"gotest.tools/assert"
	"gotest.tools/assert/cmp"
	"reflect"
	"strings"
	"testing"
)

type Option1 bool
type Option2 string
type Option3 int

func option1(o ...interface{}) bool {
	return util.Option(Option1(false), o).Bool()
}

func option2(o ...interface{}) string {
	return util.Option(Option2(""), o).String()
}

func option3(o ...interface{}) int {
	return int(util.Option(Option3(0), o).Int())
}

func Test_Option1(t *testing.T) {

	assert.Assert(t, option1(Option1(true)) == true)
	assert.Assert(t, option1(Option1(true), Option1(false)) == true)
	assert.Assert(t, option1(Option1(false), Option1(true)) == false)
	assert.Assert(t, option1(Option2(0)) == false)
	assert.Assert(t, option1() == false)

}

func Test_Option2(t *testing.T) {
	assert.Assert(t, option2(Option2("hello")) == "hello")
	assert.Assert(t, option2(Option1(false)) == "")
}

func Test_Option3(t *testing.T) {
	assert.Assert(t, option3(Option3(42)) == 42)
	assert.Assert(t, option3(Option1(false)) == 0)
}

func Test_BitsAppend(t *testing.T) {
	b := util.Bits{}
	b = b.Append(util.Bits{}, 0)
	assert.Assert(t, b.Len() == 0)
	q := b.Append(util.FillBits(1), 33)
	assert.Assert(t, q.Len() == 34)
	assert.Assert(t, cmp.Panics(func() {
		q.Append(util.Bits{}, 33)
	}))
}

func Test_Bits1(t *testing.T) {
	b := util.FillBits(31)
	assert.Assert(t, b.Bit(0))
	b.Set(0, false)
	assert.Assert(t, !b.Bit(0))
	assert.Assert(t, b.Len() == 31)
	// b => 0111....[i=31]000...
	assert.Assert(t, b.Repr() == "0"+strings.Repeat("1", 30))
	for i := 1; i < 31; i++ {
		assert.Assert(t, b.Bit(i))
	}
	assert.Assert(t, !b.Bit(31))
	z := util.FillBits(1)
	c := z.Append(b, 31)
	// c => 100....[31]0111...[i=62]000...
	assert.Assert(t, c.Repr() == "1"+strings.Repeat("0", 30)+"0"+strings.Repeat("1", 30))
	assert.Assert(t, c.Bit(0))
	for i := 1; i < 31; i++ {
		assert.Assert(t, !c.Bit(i))
	}
	assert.Assert(t, !c.Bit(31))
	for i := 32; i < 62; i++ {
		assert.Assert(t, c.Bit(i))
	}
	assert.Assert(t, !c.Bit(62))
	assert.Assert(t, !c.Bit(63))
	assert.Assert(t, !c.Bit(64))
	assert.Assert(t, !c.Bit(1064))
	c.Set(255, true)
	assert.Assert(t,
		c.Repr() == "1"+
			strings.Repeat("0", 30)+
			"0"+
			strings.Repeat("1", 30)+
			strings.Repeat("0", 255-62)+
			"1")
	assert.Assert(t, c.Len() == 256)
	assert.Assert(t, !c.Bit(62))
	assert.Assert(t, !c.Bit(63))
	assert.Assert(t, !c.Bit(64))
	assert.Assert(t, !c.Bit(1064))
}

func Test_Bits2(t *testing.T) {
	b := util.Bits{}
	assert.Assert(t, b.Len() == 0)
	b.Set(255, false)
	assert.Assert(t, b.Len() == 0)
	b.Set(128, true)
	assert.Assert(t, b.Len() == 129)
	b.Set(128, false)
	assert.Assert(t, b.Len() == 0)
	a := util.Bits{}
	a.Set(0, true)
	a.Set(256, true)
	a.Set(256, false)
	assert.Assert(t, a.Len() == 1)
}

func Test_Bits3(t *testing.T) {
	b := util.FillBits(128)
	assert.Assert(t, b.Len() == 128)
	b = b.Append(util.FillBits(120), 128)
	assert.Assert(t, b.Len() == 128+120)
	b = b.Append(util.FillBits(67), 128+120)
	assert.Assert(t, b.Len() == 128+120+67)
}

func Test_Bits4(t *testing.T) {
	b := util.Bits{}.Append(util.FillBits(3), 127)
	s := strings.Repeat("0", 127) + "111"
	r := []uint8{}
	for i := range s {
		if i != 0 && i%8 == 0 {
			r = append(r, uint8('.'))
		}
		r = append(r, s[i])
	}
	assert.Assert(t, string(r) == b.String())
}

func Test_Bits5(t *testing.T) {
	b := util.Bits{}.Append(util.FillBits(3), 127)
	b.Set(0, true)
	q := b.Slice(1, 127)
	assert.Assert(t, q.Len() == 0)
	q = b.Slice(1, 128)
	assert.Assert(t, q.Len() == 127)
	q = b.Slice(0, 127)
	assert.Assert(t, q.Len() == 1)
}

func Test_Convert(t *testing.T) {
	q := []int{1, 2, 3}
	assert.DeepEqual(t, util.Convert(reflect.ValueOf(q), internal.IntType), q)
	assert.Assert(t, cmp.Panics(func() {
		util.Convert(reflect.ValueOf(q), internal.TsType)
	}))
	assert.DeepEqual(t, util.Convert(reflect.ValueOf(int(1)), internal.StringType), "1")
}
