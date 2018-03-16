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
	assert.Equal(t, len(alphabet)+1, len(cc.aplw))
	assert.Equal(t, len(alphabet), len(cc.actw))
	assert.Equal(t, 0, cc.curpos)
	assert.Equal(t, 0, cc.ctpos)
	assert.Equal(t, byte('M'), cc.start)
}

func TestWheatstone_BlockSize(t *testing.T) {
	c, _ := NewCipher('M', key1, key2)
	assert.NotNil(t, c)
	assert.Equal(t, 1, c.BlockSize())
}

func TestEncode(t *testing.T) {
	c, err := NewCipher('M', key1, key2)

	assert.NoError(t, err)
	assert.NotNil(t, c)

	cc := c.(*wheatstone)

	assert.Equal(t, 0, cc.curpos)
	assert.Equal(t, 0, cc.ctpos)

	// Round 1
	cw := cc.encode(byte('C'))

	assert.Equal(t, 1, cc.curpos)
	assert.Equal(t, 1, cc.ctpos)

	assert.Equal(t, byte('B'), cw)

	// Round 2
	cw = cc.encode(byte('H'))

	assert.Equal(t, 15, cc.curpos)
	assert.Equal(t, 15, cc.ctpos)

	assert.Equal(t, byte('Y'), cw)

	// Round 3
	cw = cc.encode(byte('A'))

	assert.Equal(t, 2, cc.curpos)
	assert.Equal(t, 3, cc.ctpos)

	assert.Equal(t, byte('V'), cw)

	// Round 4
	cw = cc.encode(byte('R'))

	assert.Equal(t, 23, cc.curpos)
	assert.Equal(t, 24, cc.ctpos)

	assert.Equal(t, byte('L'), cw)

}

func TestEncode1(t *testing.T) {
	c, err := NewCipher('M', key1, key2)

	assert.NoError(t, err)
	assert.NotNil(t, c)

	cc := c.(*wheatstone)

	cw := cc.encode(byte('W'))
	assert.Equal(t, byte('T'), cw)

}

func TestDecode(t *testing.T) {
	c, err := NewCipher('M', key1, key2)

	assert.NoError(t, err)
	assert.NotNil(t, c)

	cc := c.(*wheatstone)
	assert.Equal(t, 0, cc.curpos)
	assert.Equal(t, 0, cc.ctpos)

	cw := cc.decode(byte('B'))

	assert.Equal(t, byte('C'), cw)
	assert.Equal(t, 1, cc.curpos)
	assert.Equal(t, 1, cc.ctpos)

	cw = cc.decode(byte('Y'))
	assert.Equal(t, byte('H'), cw)
	assert.Equal(t, 15, cc.curpos)
	assert.Equal(t, 15, cc.ctpos)

	cw = cc.decode(byte('V'))
	assert.Equal(t, byte('A'), cw)
	assert.Equal(t, 2, cc.curpos)
	assert.Equal(t, 3, cc.ctpos)

	cw = cc.decode(byte('L'))
	assert.Equal(t, byte('R'), cw)
	assert.Equal(t, 23, cc.curpos)
	assert.Equal(t, 24, cc.ctpos)

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
}

func TestWheatstone_Decrypt(t *testing.T) {
	c, err := NewCipher('M', key1, key2)

	assert.NoError(t, err)
	assert.NotNil(t, c)

	src := bytes.NewBufferString(plainTxt).Bytes()
	dec := bytes.NewBufferString(cipherTxt).Bytes()
	dst := make([]byte, len(dec))
	c.Decrypt(dst, dec)
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
	assert.EqualValues(t, src, dst)
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

func BenchmarkWheatstone_Encrypt(b *testing.B) {
	c, _ := NewCipher('M', key1, key2)

	src := bytes.NewBufferString(plainTxt).Bytes()
	dst := make([]byte, len(src))

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.Encrypt(dst, src)
	}
}

func BenchmarkWheatstone_Decrypt(b *testing.B) {
	c, _ := NewCipher('M', key1, key2)

	src := bytes.NewBufferString(cipherTxt).Bytes()
	dst := make([]byte, len(src))

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.Decrypt(dst, src)
	}
}

func BenchmarkWheatstone_EncryptLong(b *testing.B) {
	c, _ := NewCipher('M', key1, key2)

	src := bytes.NewBufferString(lplainTxt).Bytes()
	dst := make([]byte, len(src))

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.Encrypt(dst, src)
	}
}

func BenchmarkWheatstone_DecryptLong(b *testing.B) {
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
