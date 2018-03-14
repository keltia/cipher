package nihilist

import (
	"bytes"
	"crypto/cipher"
	"github.com/stretchr/testify/assert"
	"testing"
)

var TestNihilistData = []struct {
	key1, key2, chrs string
	pt               string
	ct               string
}{
	{"ARABESQUE", "SUBWAY", "37", "IFYOUCANREADTHIS", "1037306631738227035749"},
}

func TestNewCipher(t *testing.T) {
	for _, cp := range TestNihilistData {
		c, err := NewCipher(cp.key1, cp.key2, cp.chrs)

		assert.NotNil(t, c)
		assert.NoError(t, err)
		assert.Implements(t, (*cipher.Block)(nil), c)
	}
}

func TestNewCipher2(t *testing.T) {
	c, err := NewCipher("PORTABLE", "", "89")

	assert.Empty(t, c)
	assert.Error(t, err)
}

func TestNewCipher3(t *testing.T) {
	c, err := NewCipher("", "SUBWAY", "62")

	assert.Empty(t, c)
	assert.Error(t, err)
}

func TestNihilistcipher_BlockSize(t *testing.T) {
	for _, cp := range TestNihilistData {
		c, err := NewCipher(cp.key1, cp.key2, cp.chrs)

		assert.NotNil(t, c)
		assert.NoError(t, err)
		assert.Equal(t, len(cp.key2), c.BlockSize())
	}
}

func TestNihilistcipher_Encrypt(t *testing.T) {
	c, _ := NewCipher(TestNihilistData[0].key1, TestNihilistData[0].key2, TestNihilistData[0].chrs)
	cc := c.(*nihilistcipher)

	assert.NotNil(t, cc.sc)
	assert.NotNil(t, cc.transp)

	pt := TestNihilistData[0].pt
	ct := TestNihilistData[0].ct

	dst := make([]byte, len(ct))
	c.Encrypt(dst, bytes.NewBufferString(pt).Bytes())

	assert.EqualValues(t, ct, string(dst))
}

func TestNihilistcipher_Decrypt(t *testing.T) {
	c, _ := NewCipher(TestNihilistData[0].key1, TestNihilistData[0].key2, TestNihilistData[0].chrs)
	cc := c.(*nihilistcipher)

	assert.NotNil(t, cc.sc)
	assert.NotNil(t, cc.transp)

	pt := TestNihilistData[0].pt
	ct := TestNihilistData[0].ct

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
	chrs := "83"

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c, _ = NewCipher(key1, key2, chrs)
	}
	gc = c
}

func BenchmarkNihilistcipher_Encrypt(b *testing.B) {
	c, _ := NewCipher(TestNihilistData[0].key1, TestNihilistData[0].key2, TestNihilistData[0].chrs)

	pt := TestNihilistData[0].pt
	ct := TestNihilistData[0].ct

	dst := make([]byte, len(ct))

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.Encrypt(dst, bytes.NewBufferString(pt).Bytes())
	}
}

func BenchmarkNihilistcipher_Decrypt(b *testing.B) {
	c, _ := NewCipher(TestNihilistData[0].key1, TestNihilistData[0].key2, TestNihilistData[0].chrs)

	pt := TestNihilistData[0].pt
	ct := TestNihilistData[0].ct

	dst := make([]byte, len(ct))
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.Encrypt(dst, bytes.NewBufferString(pt).Bytes())
	}
}
