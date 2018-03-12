package adfgvx

import (
	"bytes"
	"crypto/cipher"
	"github.com/stretchr/testify/assert"
	"testing"
)

var TestADFGVXData = []struct {
	key1, key2 string
	pt         string
	ct         string
}{
	{"PORTABLE", "SUBWAY", "ATTACKATDAWN", "AFDFADAGAAAAVVVVGFGVGGGX"},
}

func TestNewCipher(t *testing.T) {
	for _, cp := range TestADFGVXData {
		c, err := NewCipher(cp.key1, cp.key2)

		assert.NotNil(t, c)
		assert.NoError(t, err)
		assert.Implements(t, (*cipher.Block)(nil), c)
	}
}

func TestAdfgvxcipher_BlockSize(t *testing.T) {
	for _, cp := range TestADFGVXData {
		c, err := NewCipher(cp.key1, cp.key2)

		assert.NotNil(t, c)
		assert.NoError(t, err)
		assert.Equal(t, len(cp.key2), c.BlockSize())
	}
}

func TestAdfgvxcipher_Encrypt(t *testing.T) {
	c, _ := NewCipher(TestADFGVXData[0].key1, TestADFGVXData[0].key2)
	cc := c.(*adfgvxcipher)

	assert.NotNil(t, cc.sqr)
	assert.NotNil(t, cc.transp)

	pt := TestADFGVXData[0].pt
	ct := TestADFGVXData[0].ct

	dst := make([]byte, len(ct))
	c.Encrypt(dst, bytes.NewBufferString(pt).Bytes())

	assert.EqualValues(t, ct, string(dst))
}

func TestAdfgvxcipher_Decrypt(t *testing.T) {
	c, _ := NewCipher(TestADFGVXData[0].key1, TestADFGVXData[0].key2)
	cc := c.(*adfgvxcipher)

	assert.NotNil(t, cc.sqr)
	assert.NotNil(t, cc.transp)

	pt := TestADFGVXData[0].pt
	ct := TestADFGVXData[0].ct

	dst := make([]byte, len(pt))
	c.Decrypt(dst, bytes.NewBufferString(ct).Bytes())

	assert.EqualValues(t, pt, string(dst))
}

// - benchmarks

var gc cipher.Block

func BenchmarkNewCipher(b *testing.B) {
	var c cipher.Block

	key1 := "PORTABLE"
	key2 := "SUBWAY"

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c, _ = NewCipher(key1, key2)
	}
	gc = c
}

func BenchmarkAdfgvxcipher_Encrypt(b *testing.B) {
	c, _ := NewCipher(TestADFGVXData[0].key1, TestADFGVXData[0].key2)

	pt := TestADFGVXData[0].pt
	ct := TestADFGVXData[0].ct

	dst := make([]byte, len(ct))

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.Encrypt(dst, bytes.NewBufferString(pt).Bytes())
	}
}

func BenchmarkAdfgvxcipher_Decrypt(b *testing.B) {
	c, _ := NewCipher(TestADFGVXData[0].key1, TestADFGVXData[0].key2)

	pt := TestADFGVXData[0].pt
	ct := TestADFGVXData[0].ct

	dst := make([]byte, len(ct))
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.Encrypt(dst, bytes.NewBufferString(pt).Bytes())
	}
}
