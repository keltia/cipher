package playfair

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewCipher(t *testing.T) {
	c, err := NewCipher("ARABESQUE")

	assert.NotNil(t, c)
	assert.NoError(t, err)
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

func TestPlayfairCipher_Decrypt(t *testing.T) {
	c, _ := NewCipher("PLAYFAIREXAMPLE")

	ct := []byte("BMODZBXDNABEKUDMUIXMMOUVIF")
	pt := []byte("HIDETHEGOLDINTHETREXESTUMP")

	dst := make([]byte, len(ct))

	c.Decrypt(dst, ct)
	assert.EqualValues(t, pt, dst)
}

func BenchmarkCipher_Encrypt(b *testing.B) {
	c, _ := NewCipher("PLAYFAIREXAMPLE")

	pt := []byte("HIDETHEGOLDINTHETREXESTUMP")

	dst := make([]byte, len(pt))

	for n := 0; n < b.N; n++ {
		c.Encrypt(dst, pt)
	}
}

func BenchmarkCipher_Decrypt(b *testing.B) {
	c, _ := NewCipher("PLAYFAIREXAMPLE")

	ct := []byte("BMODZBXDNABEKUDMUIXMMOUVIF")

	dst := make([]byte, len(ct))

	for n := 0; n < b.N; n++ {
		c.Decrypt(dst, ct)
	}
}
