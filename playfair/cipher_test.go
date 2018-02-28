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

}

func TestPlayfairCipher_Decrypt(t *testing.T) {
	c, _ := NewCipher("PLAYFAIREXAMPLE")

	ct := []byte("BMODZBXDNABEKUDMUIXMMOUVIF")
	pt := []byte("HIDETHEGOLDINTHETREXESTUMP")

	var dst []byte

	c.Decrypt(dst, ct)
	assert.EqualValues(t, pt, dst)
}
