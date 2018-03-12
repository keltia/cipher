package transposition

import (
	"bytes"
	"crypto/cipher"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewCipher(t *testing.T) {
	c, err := NewCipher("ABCDE")

	assert.NotNil(t, c)
	assert.NoError(t, err)
	assert.Implements(t, (*cipher.Block)(nil), c)

	cc := c.(*transp)

	assert.Equal(t, "ABCDE", cc.key)
	assert.EqualValues(t, []byte{0, 1, 2, 3, 4}, cc.tkey)
}

func TestNewCipher2(t *testing.T) {
	_, err := NewCipher("")

	assert.Error(t, err)
}

func TestTransp_BlockSize(t *testing.T) {
	c, err := NewCipher("ABCDE")

	assert.NotNil(t, c)
	assert.NoError(t, err)
	assert.Equal(t, 5, c.BlockSize())
}

func TestTransp_Encrypt(t *testing.T) {
	pt := bytes.NewBufferString("ATTACKATDAWNATPOINT42X23XSENDMOREMUNITIONSBYNIGHTX123")

	cts := []struct{ k, c string }{
		{"ARABESQUE", "AATNIITN2MIHAAXOOTCT2RNXDNENNAOXMB2TW4DTGKP3ES1TISUY3"},
		{"SUBWAY", "CWI2DUNG3TDP2EEIN1AAATXOIBTTTT4SRTYXAAOXNMOI2KNN3MNSH"},
		{"PORTABLE", "CA2DIN3KTXMTITO3ROHAP2OIGTANSMSXADIXENTTWTEUB1AN4NNY2"},
	}

	for _, cti := range cts {
		ct := bytes.NewBufferString(cti.c)
		key := cti.k

		c, err := NewCipher(key)
		assert.NotNil(t, c)
		assert.NoError(t, err)

		dst := make([]byte, pt.Len())
		c.Encrypt(dst, pt.Bytes())

		assert.EqualValues(t, ct.Bytes(), dst)
	}
}

func TestTransp_Encrypt1(t *testing.T) {
	pt := bytes.NewBufferString("AVAGAGAVDFFGAVAGDGAVGVFX")

	cts := []struct{ k, c string }{
		{"SUBWAY", "AFDFADAGAAAAVVVVGFGVGGGX"},
	}

	for _, cti := range cts {
		ct := bytes.NewBufferString(cti.c)
		key := cti.k

		c, err := NewCipher(key)
		assert.NotNil(t, c)
		assert.NoError(t, err)

		dst := make([]byte, pt.Len())
		c.Encrypt(dst, pt.Bytes())

		assert.EqualValues(t, ct.Bytes(), dst)
	}
}

func TestTransp_Decrypt(t *testing.T) {
	pt := bytes.NewBufferString("ATTACKATDAWNATPOINT42X23XSENDMOREMUNITIONSBYNIGHTX123")

	cts := []struct{ k, c string }{
		{"ARABESQUE", "AATNIITN2MIHAAXOOTCT2RNXDNENNAOXMB2TW4DTGKP3ES1TISUY3"},
		{"SUBWAY", "CWI2DUNG3TDP2EEIN1AAATXOIBTTTT4SRTYXAAOXNMOI2KNN3MNSH"},
		{"PORTABLE", "CA2DIN3KTXMTITO3ROHAP2OIGTANSMSXADIXENTTWTEUB1AN4NNY2"},
	}

	for _, cti := range cts {
		ct := bytes.NewBufferString(cti.c)
		key := cti.k

		c, err := NewCipher(key)
		assert.NotNil(t, c)
		assert.NoError(t, err)

		dst := make([]byte, pt.Len())
		c.Decrypt(dst, ct.Bytes())

		assert.EqualValues(t, pt.Bytes(), dst)
	}

}

func TestTransp_Decrypt1(t *testing.T) {
	pt := bytes.NewBufferString("AVAGAGAVDFFGAVAGDGAVGVFX")

	cts := []struct{ k, c string }{
		{"SUBWAY", "AFDFADAGAAAAVVVVGFGVGGGX"},
	}

	for _, cti := range cts {
		ct := bytes.NewBufferString(cti.c)
		key := cti.k

		c, err := NewCipher(key)
		assert.NotNil(t, c)
		assert.NoError(t, err)

		dst := make([]byte, pt.Len())
		c.Decrypt(dst, ct.Bytes())

		assert.EqualValues(t, pt.Bytes(), dst)
	}

}

// ----- benchmark

var gc cipher.Block

func BenchmarkNewCipher(b *testing.B) {
	var c cipher.Block

	for n := 0; n < b.N; n++ {
		c, _ = NewCipher("ABCDE")
	}
	gc = c
}

func BenchmarkTransp_Encrypt(b *testing.B) {
	pt := bytes.NewBufferString("ATTACKATDAWNATPOINT42X23XSENDMOREMUNITIONSBYNIGHTX123")
	key := "ARABESQUE"

	c, _ := NewCipher(key)
	dst := make([]byte, pt.Len())
	for n := 0; n < b.N; n++ {
		c.Encrypt(dst, pt.Bytes())
	}
}

func BenchmarkTransp_Decrypt(b *testing.B) {
	ct := bytes.NewBufferString("AATNIITN2MIHAAXOOTCT2RNXDNENNAOXMB2TW4DTGKP3ES1TISUY3")
	key := "ARABESQUE"
	c, _ := NewCipher(key)
	dst := make([]byte, ct.Len())
	for n := 0; n < b.N; n++ {
		c.Encrypt(dst, ct.Bytes())
	}
}
