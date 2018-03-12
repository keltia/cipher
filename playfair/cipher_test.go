package playfair

import (
	"crypto/cipher"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewCipher(t *testing.T) {
	c, err := NewCipher("ARABESQUE")

	assert.NotNil(t, c)
	assert.NoError(t, err)
	assert.Implements(t, (*cipher.Block)(nil), c)
}

func TestPlayfairCipher_BlockSize(t *testing.T) {
	c, err := NewCipher("ARABESQUE")

	assert.Equal(t, 2, c.BlockSize())
	assert.NoError(t, err)
}

func TestPlayfairCipher_Encrypt(t *testing.T) {
	c, _ := NewCipher("PLAYFAIREXAMPLE")

	ct := []byte("BMODZBXDNABEKUDMUIXMMOUVIF")
	pt := []byte("HIDETHEGOLDINTHETREXESTUMP")

	dst := make([]byte, len(ct))

	c.Encrypt(dst, pt)
	assert.EqualValues(t, ct, dst)
}

func TestPlayfairCipher_EncryptX(t *testing.T) {
	c, _ := NewCipher("PLAYFAIREXAMPLE")

	pt := []byte("HID")
	ct := []byte("BMGE")

	dst := make([]byte, len(pt)+1)

	c.Encrypt(dst, pt)
	assert.EqualValues(t, ct, dst)
}

func TestPlayfairCipher_Decrypt(t *testing.T) {
	c, _ := NewCipher("PLAYFAIREXAMPLE")

	ct := []byte("BMODZBXDNABEKUDMUIXMMOUVIF")
	pt := []byte("HIDETHEGOLDINTHETREXESTUMP")

	dst := make([]byte, len(ct))

	c.Decrypt(dst, ct)
	assert.EqualValues(t, pt, dst)
}

func TestPlayfairCipher_DecryptPanic(t *testing.T) {
	c, _ := NewCipher("PLAYFAIREXAMPLE")

	ct := []byte("BMO")

	dst := make([]byte, len(ct))

	assert.Panics(t, func() {
		c.Decrypt(dst, ct)
	})
}

var gc cipher.Block

func BenchmarkNewCipher(b *testing.B) {
	var c cipher.Block

	for n := 0; n < b.N; n++ {
		c, _ = NewCipher("PLAYFAIREXAMPLE")
	}
	gc = c
}

func BenchmarkCipher_Encrypt(b *testing.B) {
	c, _ := NewCipher("PLAYFAIREXAMPLE")

	pt := []byte("HIDETHEGOLDINTHETREXESTUMP")

	dst := make([]byte, len(pt))

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.Encrypt(dst, pt)
	}
}

func BenchmarkCipher_Decrypt(b *testing.B) {
	c, _ := NewCipher("PLAYFAIREXAMPLE")

	ct := []byte("BMODZBXDNABEKUDMUIXMMOUVIF")

	dst := make([]byte, len(ct))

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.Decrypt(dst, ct)
	}
}
