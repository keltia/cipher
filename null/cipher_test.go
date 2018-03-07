package null

import (
	"crypto/cipher"
	"github.com/stretchr/testify/assert"
	"testing"
)

type NullTest struct {
	pt, ct string
}

var encryptCaesarTests = []NullTest{
	{"ABCDE", "ABCDE"},
	{"COUCOU", "COUCOU"},
}

func TestNewCipher(t *testing.T) {
	c, _ := NewCipher()
	assert.NotNil(t, c)
	assert.EqualValues(t, 1, c.BlockSize())
}

func TestNullCipher_Encrypt(t *testing.T) {
	c, _ := NewCipher()
	assert.NotNil(t, c)
	assert.EqualValues(t, 1, c.BlockSize())

	for _, pair := range encryptCaesarTests {
		plain := []byte(pair.pt)
		cipher := make([]byte, len(plain))
		c.Encrypt(cipher, plain)
		assert.Equal(t, []byte(pair.ct), cipher)
	}
}

func TestNullCipher_Decrypt(t *testing.T) {
	c, _ := NewCipher()
	assert.NotNil(t, c)
	assert.EqualValues(t, 1, c.BlockSize())

	for _, pair := range encryptCaesarTests {
		plain := []byte(pair.pt)
		cipher := []byte(pair.ct)
		nplain := make([]byte, len(plain))
		c.Decrypt(nplain, cipher)
		assert.Equal(t, []byte(pair.pt), nplain)
	}
}

var gc cipher.Block

func BenchmarkNewCipher(b *testing.B) {
	var c cipher.Block

	for n := 0; n < b.N; n++ {
		c, _ = NewCipher()
	}
	gc = c
}

func BenchmarkNullCipher_Encrypt(b *testing.B) {
	c, _ := NewCipher()

	plain := []byte("ABCDE")
	cipher := make([]byte, len(plain))

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.Encrypt(cipher, plain)
	}
}

func BenchmarkNullCipher_Decrypt(b *testing.B) {
	c, _ := NewCipher()
	cipher := []byte("ABCDE")
	nplain := make([]byte, len(cipher))

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.Decrypt(nplain, cipher)
	}
}
