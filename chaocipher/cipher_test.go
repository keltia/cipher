package chaocipher

import (
	"bytes"
	"crypto/cipher"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAdvance(t *testing.T) {

}

func TestEncodeBoth(t *testing.T) {
	c, _ := NewCipher(alphabet, alphabet)
	gc = c
}

func TestEncode(t *testing.T) {

}

func TestDecode(t *testing.T) {

}

func TestNewCipher(t *testing.T) {
	c, err := NewCipher(alphabet, alphabet)
	assert.NoError(t, err)
	assert.NotNil(t, c)
	assert.Implements(t, (*cipher.Block)(nil), c)
}

func TestNewCipher2(t *testing.T) {
	c, err := NewCipher("AB", "CD")
	assert.Error(t, err)
	assert.EqualValues(t, &chaocipher{}, c)
}

func TestChaocipher_BlockSize(t *testing.T) {
	c, _ := NewCipher(alphabet, alphabet)
	assert.NotNil(t, c)
	assert.Equal(t, 1, c.BlockSize())
}

var (
	plainTxt  = "WELLDONEISBETTERTHANWELLSAID"
	cipherTxt = "OAHQHCNYNXTSZJRRHJBYHQKSOUJY"

	keyPlain  = "PTLNBQDEOYSFAVZKGJRIHWXUMC"
	KeyCipher = "HXUCZVAMDSLKPEFJRIGTWOBNYQ"
)

func TestChaocipher_Encrypt(t *testing.T) {
	c, err := NewCipher(keyPlain, KeyCipher)

	assert.NoError(t, err)
	assert.NotNil(t, c)

	src := bytes.NewBufferString(plainTxt).Bytes()
	enc := bytes.NewBufferString(cipherTxt).Bytes()
	dst := make([]byte, len(src))
	c.Encrypt(dst, src)
	assert.EqualValues(t, enc, dst)
}

func TestChaocipher_Decrypt(t *testing.T) {
	c, err := NewCipher(keyPlain, KeyCipher)

	assert.NoError(t, err)
	assert.NotNil(t, c)

	src := bytes.NewBufferString(plainTxt).Bytes()
	dec := bytes.NewBufferString(cipherTxt).Bytes()
	dst := make([]byte, len(dec))
	c.Decrypt(dst, dec)
	assert.EqualValues(t, src, dst)
}

var gc cipher.Block

func BenchmarkNewCipher(b *testing.B) {
	var c cipher.Block

	for n := 0; n < b.N; n++ {
		c, _ = NewCipher(alphabet, alphabet)
	}
	gc = c
}

func BenchmarkChaocipher_Encrypt(b *testing.B) {

}

func BenchmarkChaocipher_Decrypt(b *testing.B) {

}
