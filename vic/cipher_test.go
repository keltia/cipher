package vic

import (
	"crypto/cipher"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewCipher(t *testing.T) {
	c, err := NewCipher("741776", "IDREAMOFJEANNIEWITHT", "77651")

	assert.NotNil(t, c)
	assert.NoError(t, err)
	assert.Implements(t, (*cipher.Block)(nil), c)
}

var TestToNumericOneData = []struct {
	s string
	r []byte
}{
	{"IDREAMOFJE", []byte{6, 2, 0, 3, 1, 8, 9, 5, 7, 4}},
	{"ANNIEWITHT", []byte{1, 6, 7, 4, 2, 0, 5, 8, 3, 9}},
}

func TestToNumericOne(t *testing.T) {
	for _, cp := range TestToNumericOneData {
		assert.EqualValues(t, cp.r, toNumericOne(cp.s))
	}
}

var TestAddmod10Data = []struct {
	a, b, c []byte
}{
	{[]byte{8, 6, 1, 5, 4}, []byte{2, 0, 9, 5, 2}, []byte{0, 6, 0, 0, 6}},
	{[]byte{7, 7, 6, 5, 1}, []byte{7, 4, 1, 7, 7}, []byte{4, 1, 7, 2, 8}},
}

func TestAddmod10(t *testing.T) {
	for _, cp := range TestAddmod10Data {
		r := addmod10(cp.a, cp.b)
		assert.EqualValues(t, cp.c, r)
	}
}

var TestSubmod10Data = []struct {
	a, b, c []byte
}{
	{[]byte{8, 6, 1, 5, 4}, []byte{2, 0, 9, 5, 2}, []byte{6, 6, 2, 0, 2}},
	{[]byte{7, 7, 6, 5, 1}, []byte{7, 4, 1, 7, 7}, []byte{0, 3, 5, 8, 4}},
}

func TestSubmod10(t *testing.T) {
	for _, cp := range TestSubmod10Data {
		r := submod10(cp.a, cp.b)
		assert.EqualValues(t, cp.c, r)
	}
}

var TestChainaddData = []struct {
	a, b []byte
}{
	{[]byte{8, 6, 1, 5, 4}, []byte{4, 7, 6, 9, 8}},
	{[]byte{7, 7, 6, 5, 1}, []byte{4, 3, 1, 6, 5}},
}

func TestChainadd(t *testing.T) {
	for _, cp := range TestChainaddData {
		r := chainadd(cp.a)
		assert.EqualValues(t, cp.b, r)
	}
}

var TestExpand5To10Data = []struct {
	a, b []byte
}{
	{[]byte{8, 6, 1, 5, 4}, []byte{8, 6, 1, 5, 4, 4, 7, 6, 9, 8}},
	{[]byte{7, 7, 6, 5, 1}, []byte{7, 7, 6, 5, 1, 4, 3, 1, 6, 5}},
	{[]byte{0, 3, 5, 8, 4}, []byte{0, 3, 5, 8, 4, 3, 8, 3, 2, 7}},
}

func TestExpand5To10(t *testing.T) {
	for _, cp := range TestExpand5To10Data {
		r := expand5to10(cp.a)
		assert.EqualValues(t, cp.b, r)
	}
}

var TestFirstEncodeData = []struct {
	r1, r2, r []byte
}{
	{[]byte{6, 5, 5, 1, 5, 1, 7, 8, 9, 1}, []byte{1, 6, 7, 4, 2, 0, 5, 8, 3, 9}, []byte{0, 2, 2, 1, 2, 1, 5, 8, 3, 1}},
}

func TestFirstEncode(t *testing.T) {
	for _, cp := range TestFirstEncodeData {
		r := firstEncode(cp.r1, cp.r2)
		assert.EqualValues(t, cp.r, r)
	}
}

func TestViccipher_BlockSize(t *testing.T) {
	c, _ := NewCipher("741776", "IDREAMOFJEANNIEWITHT", "77651")

	assert.NotNil(t, c)
	assert.Equal(t, 1, c.BlockSize())
}

// -- benchmarks

var gc cipher.Block

func BenchmarkNewCipher(b *testing.B) {
	var c cipher.Block

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c, _ = NewCipher("741776", "IDREAMOFJEANNIEWITHT", "77651")
	}
	gc = c
}

func BenchmarkToNumericOne(b *testing.B) {
	var c []byte

	cp := TestToNumericOneData[0]

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c = toNumericOne(cp.s)
	}
	gb = c

}

var gb []byte

func BenchmarkAddmod10(b *testing.B) {
	var c []byte

	cp := TestAddmod10Data[0]

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c = addmod10(cp.a, cp.b)
	}
	gb = c
}

func BenchmarkSubmod10(b *testing.B) {
	var c []byte

	cp := TestSubmod10Data[0]

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c = submod10(cp.a, cp.b)
	}
	gb = c
}

func BenchmarkExpand5To10(b *testing.B) {
	var c []byte

	cp := TestExpand5To10Data[0]

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c = expand5to10(cp.a)
	}
	gb = c
}

func BenchmarkFirstEncode(b *testing.B) {
	var r []byte

	cp := TestFirstEncodeData[0]

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		r = firstEncode(cp.r1, cp.r2)
	}
	gb = r
}
