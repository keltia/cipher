package wheatstone

import (
	"bytes"
	"crypto/cipher"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	plainTxt  = "CHARLES+WHEATSTONE+HAD+A+REMARKABLY+FERTILE+MIND"
	cipherTxt = "BYVLQKWAMNLCYXIOUBFLHTXGHFPBJHZZLUEZFHIVBVRTFVRQ"

	key1 = "CIPHER"
	key2 = "MACHINE"

	pkey = "+CAKSYIBLTZPDMUHFNVEGOWRJQX"
	ckey = "MBOVADPWCFQXHGRYIJSZNKTELU"

	lplainTxt  = "IFYOUCANREADTHISYOUEITHERDOWNLOADEDMYOWNIMPLEMENTATIONOFwheatstoneORYOUWROTEONEOFYOUROWNINEITHERCASELETMEKNOWANDACCEPTMYCONGRATULATIONSX"
	lcipherTxt = "TLMAGOONSKJBJYBQVGDQCDUNWNMZPLOYCWPCWKWQRBOYADSLQBKYCDGXJOLONKTTLRUZZJQGJBQNRQHQRREUIYIDHZOMVWZMVYUFQOGSNNUVYTJGQPSQTBRWFHLTCLVVBPMYYQVC"

	aplw = []byte{'+', 'C', 'A', 'K', 'S', 'Y', 'I', 'B', 'L', 'T', 'Z', 'P', 'D', 'M', 'U', 'H', 'F', 'N', 'V', 'E', 'G', 'O', 'W', 'R', 'J', 'Q', 'X'}
	actw = []byte{'M', 'B', 'O', 'V', 'A', 'D', 'P', 'W', 'C', 'F', 'Q', 'X', 'H', 'G', 'R', 'Y', 'I', 'J', 'S', 'Z', 'N', 'K', 'T', 'E', 'L', 'U'}
)

var TestWheatstoneData = []struct {
	pkey  string
	ckey  string
	start byte
	aplw  []byte
	actw  []byte
	pt    string
	ct    string
}{}

func TestNewCipher(t *testing.T) {
	c, err := NewCipher('M', key1, key2)
	assert.NoError(t, err)
	assert.NotNil(t, c)
	assert.Implements(t, (*cipher.Block)(nil), c)

	cc := c.(*wheatstone)
	assert.Equal(t, pkey, cc.pkey)
	assert.Equal(t, ckey, cc.ckey)
	assert.EqualValues(t, aplw, cc.aplw)
	assert.EqualValues(t, actw, cc.actw)
}

func TestWheatstone_BlockSize(t *testing.T) {
	c, _ := NewCipher('M', key1, key2)
	assert.NotNil(t, c)
	assert.Equal(t, 1, c.BlockSize())
}

func TestWheatstone_Encrypt(t *testing.T) {
	c, err := NewCipher('M', key1, key2)

	assert.NoError(t, err)
	assert.NotNil(t, c)

	src := bytes.NewBufferString(plainTxt).Bytes()
	enc := bytes.NewBufferString(cipherTxt).Bytes()
	dst := make([]byte, len(src))
	c.Encrypt(dst, src)
	assert.EqualValues(t, enc, dst)
	gb = enc
}

func Testwheatstone_EncryptLong(t *testing.T) {
	c, err := NewCipher('M', key1, key2)

	assert.NoError(t, err)
	assert.NotNil(t, c)

	src := bytes.NewBufferString(lplainTxt).Bytes()
	enc := bytes.NewBufferString(lcipherTxt).Bytes()
	dst := make([]byte, len(src))
	c.Encrypt(dst, src)
	assert.EqualValues(t, enc, dst)
	gb = enc
}

func Testwheatstone_Decrypt(t *testing.T) {
	c, err := NewCipher('M', key1, key2)

	assert.NoError(t, err)
	assert.NotNil(t, c)

	src := bytes.NewBufferString(plainTxt).Bytes()
	dec := bytes.NewBufferString(cipherTxt).Bytes()
	dst := make([]byte, len(dec))
	c.Decrypt(dst, dec)
	gb = src
	assert.EqualValues(t, src, dst)
}

func Testwheatstone_DecryptLong(t *testing.T) {
	c, err := NewCipher('M', key1, key2)

	assert.NoError(t, err)
	assert.NotNil(t, c)

	src := bytes.NewBufferString(lplainTxt).Bytes()
	dec := bytes.NewBufferString(lcipherTxt).Bytes()
	dst := make([]byte, len(dec))
	c.Decrypt(dst, dec)
	gb = src
	assert.EqualValues(t, src, dst)
}

func TestEncode(t *testing.T) {
	c, err := NewCipher('M', key1, key2)

	assert.NoError(t, err)
	assert.NotNil(t, c)

	cc := c.(*wheatstone)
	cw := cc.encode(byte('C'))

	assert.Equal(t, 2, cc.curpos)
	assert.Equal(t, 2, cc.ctpos)
	assert.Equal(t, byte('B'), cw)
}

func TestEncode1(t *testing.T) {
	c, err := NewCipher('M', key1, key2)

	assert.NoError(t, err)
	assert.NotNil(t, c)

	cc := c.(*wheatstone)
	cw := cc.encode(byte('W'))

	assert.Equal(t, byte('O'), cw)
}

func TestDecode(t *testing.T) {
	c, err := NewCipher('M', key1, key2)

	assert.NoError(t, err)
	assert.NotNil(t, c)

	cc := c.(*wheatstone)
	cw := cc.decode(byte('P'))

	assert.Equal(t, byte('A'), cw)
}

// -- benchmarks

var gcw byte

func BenchmarkEncode(b *testing.B) {
	var cw byte
	c, _ := NewCipher('M', key1, key2)

	cc := c.(*wheatstone)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		cw = cc.encode(byte('A'))
	}
	gcw = cw
}

func BenchmarkDecode(b *testing.B) {
	var pw byte
	c, _ := NewCipher('M', key1, key2)

	cc := c.(*wheatstone)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		pw = cc.decode(byte('P'))
	}
	gcw = pw
}

var gc cipher.Block

func BenchmarkNewCipher(b *testing.B) {
	var c cipher.Block

	for n := 0; n < b.N; n++ {
		c, _ = NewCipher('M', key1, key2)
	}
	gc = c
}

func Benchmarkwheatstone_Encrypt(b *testing.B) {
	c, _ := NewCipher('M', key1, key2)

	src := bytes.NewBufferString(plainTxt).Bytes()
	dst := make([]byte, len(src))

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.Encrypt(dst, src)
	}
}

func Benchmarkwheatstone_Decrypt(b *testing.B) {
	c, _ := NewCipher('M', key1, key2)

	src := bytes.NewBufferString(cipherTxt).Bytes()
	dst := make([]byte, len(src))

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.Decrypt(dst, src)
	}
}

func Benchmarkwheatstone_EncryptLong(b *testing.B) {
	c, _ := NewCipher('M', key1, key2)

	src := bytes.NewBufferString(lplainTxt).Bytes()
	dst := make([]byte, len(src))

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.Encrypt(dst, src)
	}
}

func Benchmarkwheatstone_DecryptLong(b *testing.B) {
	c, _ := NewCipher('M', key1, key2)

	src := bytes.NewBufferString(lcipherTxt).Bytes()
	dst := make([]byte, len(src))

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.Decrypt(dst, src)
	}
}

// ---

var gb []byte
