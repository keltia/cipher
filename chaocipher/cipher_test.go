package chaocipher

import (
	"bytes"
	"crypto/cipher"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	plainTxt  = "WELLDONEISBETTERTHANWELLSAID"
	cipherTxt = "OAHQHCNYNXTSZJRRHJBYHQKSOUJY"

	lplainTxt  = "IFYOUCANREADTHISYOUEITHERDOWNLOADEDMYOWNIMPLEMENTATIONOFCHAOCIPHERORYOUWROTEONEOFYOUROWNINEITHERCASELETMEKNOWANDACCEPTMYCONGRATULATIONS"
	lcipherTxt = "TLMAGOONSKJBJYBQVGDQCDUNWNMZPLOYCWPCWKWQRBOYADSLQBKYCDGXJOLONKTTLRUZZJQGJBQNRQHQRREUIYIDHZOMVWZMVYUFQOGSNNUVYTJGQPSQTBRWFHLTCLVVBPMYYQV"

	keyPlain  = "PTLNBQDEOYSFAVZKGJRIHWXUMC"
	keyCipher = "HXUCZVAMDSLKPEFJRIGTWOBNYQ"
)

func TestAdvance(t *testing.T) {
	c, _ := NewCipher(keyPlain, keyCipher)

	cc := c.(*chaocipher)
	idx := bytes.Index([]byte(keyPlain), []byte{'A'})
	assert.Equal(t, 12, idx)

	ct := cc.cw[idx]
	assert.Equal(t, byte('P'), ct)
	assert.Equal(t, 12, idx)
	cc.advance(idx)

	expcw := bytes.NewBufferString("PFJRIGTWOBNYQEHXUCZVAMDSLK").Bytes()
	assert.EqualValues(t, expcw, cc.cw)

	exppw := bytes.NewBufferString("VZGJRIHWXUMCPKTLNBQDEOYSFA").Bytes()
	assert.EqualValues(t, exppw, cc.pw)
}

func TestAdvanceReal(t *testing.T) {
	c, _ := NewCipher(keyPlain, keyCipher)

	cc := c.(*chaocipher)
	idx := bytes.Index([]byte(keyPlain), []byte{'W'})
	assert.Equal(t, 21, idx)

	ct := cc.cw[idx]
	assert.Equal(t, byte('O'), ct)
	assert.Equal(t, 21, idx)
	cc.advance(idx)

	expcw := bytes.NewBufferString("ONYQHXUCZVAMDBSLKPEFJRIGTW").Bytes()
	assert.EqualValues(t, expcw, cc.cw)

	exppw := bytes.NewBufferString("XUCPTLNBQDEOYMSFAVZKGJRIHW").Bytes()
	assert.EqualValues(t, exppw, cc.pw)
}

func TestMyPshift(t *testing.T) {
	acw := bytes.NewBufferString(keyCipher).Bytes()
	apw := bytes.NewBufferString(keyPlain).Bytes()

	fls := bytes.NewBufferString("PEFJRIGTWOBNYQHXUCZVAMDSLK").Bytes()

	idx := bytes.Index(apw, []byte{'A'})
	assert.Equal(t, 12, idx)

	lshiftN(acw, idx)
	assert.EqualValues(t, fls, acw)

	el := acw[1]
	assert.Equal(t, byte('E'), el)

	copy(acw[1:idx+1], acw[2:idx+2])
	sls := bytes.NewBufferString("PFJRIGTWOBNYQQHXUCZVAMDSLK").Bytes()
	assert.EqualValues(t, sls, acw)

	acw[nadir] = el
	final := bytes.NewBufferString("PFJRIGTWOBNYQEHXUCZVAMDSLK").Bytes()
	assert.EqualValues(t, final, acw)

	fls = bytes.NewBufferString("VZKGJRIHWXUMCPTLNBQDEOYSFA").Bytes()

	lshiftN(apw, idx+1)
	assert.EqualValues(t, fls, apw)

	el = apw[2]
	assert.Equal(t, byte('K'), el)

	sls = bytes.NewBufferString("VZGJRIHWXUMCPPTLNBQDEOYSFA").Bytes()
	copy(apw[2:idx+1], apw[3:idx+2])
	assert.EqualValues(t, sls, apw)

	apw[nadir] = el
	final = bytes.NewBufferString("VZGJRIHWXUMCPKTLNBQDEOYSFA").Bytes()
	assert.EqualValues(t, final, apw)

}

func TestMyPshiftReal(t *testing.T) {
	acw := bytes.NewBufferString(keyCipher).Bytes()
	apw := bytes.NewBufferString(keyPlain).Bytes()

	fls := bytes.NewBufferString("OBNYQHXUCZVAMDSLKPEFJRIGTW").Bytes()

	idx := bytes.Index(apw, []byte{'W'})
	assert.Equal(t, 21, idx)

	lshiftN(acw, idx)
	assert.EqualValues(t, fls, acw)

	el := acw[1]
	assert.Equal(t, byte('B'), el)

	copy(acw[1:nadir], acw[2:nadir+1])
	sls := bytes.NewBufferString("ONYQHXUCZVAMDDSLKPEFJRIGTW").Bytes()
	assert.EqualValues(t, sls, acw)

	acw[nadir] = el
	final := bytes.NewBufferString("ONYQHXUCZVAMDBSLKPEFJRIGTW").Bytes()
	assert.EqualValues(t, final, acw)

	// --
	fls = bytes.NewBufferString("XUMCPTLNBQDEOYSFAVZKGJRIHW").Bytes()

	lshiftN(apw, idx+1)
	assert.EqualValues(t, fls, apw)

	el = apw[2]
	assert.Equal(t, byte('M'), el)

	sls = bytes.NewBufferString("XUCPTLNBQDEOYYSFAVZKGJRIHW").Bytes()
	copy(apw[2:nadir], apw[3:nadir+1])
	assert.EqualValues(t, sls, apw)

	apw[nadir] = el
	final = bytes.NewBufferString("XUCPTLNBQDEOYMSFAVZKGJRIHW").Bytes()
	assert.EqualValues(t, final, apw)

}

func TestMyEncodeBothReal(t *testing.T) {
	c, err := NewCipher(keyPlain, keyCipher)

	assert.NoError(t, err)
	assert.NotNil(t, c)

	cc := c.(*chaocipher)

	cw := cc.encodeBoth(cc.pw, cc.cw, 'A')
	assert.Equal(t, byte('P'), cw)

	final := bytes.NewBufferString("VZGJRIHWXUMCPKTLNBQDEOYSFA").Bytes()
	assert.EqualValues(t, final, cc.pw)

	final = bytes.NewBufferString("PFJRIGTWOBNYQEHXUCZVAMDSLK").Bytes()
	assert.EqualValues(t, final, cc.cw)
}

func TestEncode(t *testing.T) {
	c, err := NewCipher(keyPlain, keyCipher)

	assert.NoError(t, err)
	assert.NotNil(t, c)

	cc := c.(*chaocipher)
	cw := cc.encode(byte('A'))

	assert.Equal(t, byte('P'), cw)

	final := bytes.NewBufferString("VZGJRIHWXUMCPKTLNBQDEOYSFA").Bytes()
	assert.EqualValues(t, final, cc.pw)

	final = bytes.NewBufferString("PFJRIGTWOBNYQEHXUCZVAMDSLK").Bytes()
	assert.EqualValues(t, final, cc.cw)
}

func TestEncode1(t *testing.T) {
	c, err := NewCipher(keyPlain, keyCipher)

	assert.NoError(t, err)
	assert.NotNil(t, c)

	cc := c.(*chaocipher)
	cw := cc.encode(byte('W'))

	assert.Equal(t, byte('O'), cw)

	final := bytes.NewBufferString("ONYQHXUCZVAMDBSLKPEFJRIGTW").Bytes()
	assert.EqualValues(t, final, cc.cw)

	final = bytes.NewBufferString("XUCPTLNBQDEOYMSFAVZKGJRIHW").Bytes()
	assert.EqualValues(t, final, cc.pw)
}

func TestDecode(t *testing.T) {
	c, err := NewCipher(keyPlain, keyCipher)

	assert.NoError(t, err)
	assert.NotNil(t, c)

	cc := c.(*chaocipher)
	cw := cc.decode(byte('P'))

	assert.Equal(t, byte('A'), cw)

	final := bytes.NewBufferString("VZGJRIHWXUMCPKTLNBQDEOYSFA").Bytes()
	assert.EqualValues(t, final, cc.pw)

	final = bytes.NewBufferString("PFJRIGTWOBNYQEHXUCZVAMDSLK").Bytes()
	assert.EqualValues(t, final, cc.cw)

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

func TestChaocipher_Encrypt(t *testing.T) {
	c, err := NewCipher(keyPlain, keyCipher)

	assert.NoError(t, err)
	assert.NotNil(t, c)

	src := bytes.NewBufferString(plainTxt).Bytes()
	enc := bytes.NewBufferString(cipherTxt).Bytes()
	dst := make([]byte, len(src))
	c.Encrypt(dst, src)
	assert.EqualValues(t, enc, dst)
	_ = enc
}

func TestChaocipher_EncryptLong(t *testing.T) {
	c, err := NewCipher(keyPlain, keyCipher)

	assert.NoError(t, err)
	assert.NotNil(t, c)

	src := bytes.NewBufferString(lplainTxt).Bytes()
	enc := bytes.NewBufferString(lcipherTxt).Bytes()
	dst := make([]byte, len(src))
	c.Encrypt(dst, src)
	assert.EqualValues(t, enc, dst)
	_ = enc
}

func TestChaocipher_Decrypt(t *testing.T) {
	c, err := NewCipher(keyPlain, keyCipher)

	assert.NoError(t, err)
	assert.NotNil(t, c)

	src := bytes.NewBufferString(plainTxt).Bytes()
	dec := bytes.NewBufferString(cipherTxt).Bytes()
	dst := make([]byte, len(dec))
	c.Decrypt(dst, dec)
	_ = src
	assert.EqualValues(t, src, dst)
}

func TestChaocipher_DecryptLong(t *testing.T) {
	c, err := NewCipher(keyPlain, keyCipher)

	assert.NoError(t, err)
	assert.NotNil(t, c)

	src := bytes.NewBufferString(lplainTxt).Bytes()
	dec := bytes.NewBufferString(lcipherTxt).Bytes()
	dst := make([]byte, len(dec))
	c.Decrypt(dst, dec)
	_ = src
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
	c, _ := NewCipher(keyPlain, keyCipher)

	src := bytes.NewBufferString(plainTxt).Bytes()
	dst := make([]byte, len(src))

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.Encrypt(dst, src)
	}
}

func BenchmarkChaocipher_Decrypt(b *testing.B) {
	c, _ := NewCipher(keyPlain, keyCipher)

	src := bytes.NewBufferString(cipherTxt).Bytes()
	dst := make([]byte, len(src))

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.Decrypt(dst, src)
	}
}

func BenchmarkChaocipher_EncryptLong(b *testing.B) {
	c, _ := NewCipher(keyPlain, keyCipher)

	src := bytes.NewBufferString(lplainTxt).Bytes()
	dst := make([]byte, len(src))

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.Encrypt(dst, src)
	}
}

func BenchmarkChaocipher_DecryptLong(b *testing.B) {
	c, _ := NewCipher(keyPlain, keyCipher)

	src := bytes.NewBufferString(lcipherTxt).Bytes()
	dst := make([]byte, len(src))

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.Decrypt(dst, src)
	}
}

// ---

func TestLshift(t *testing.T) {
	var a = []byte{0, 1, 2, 3, 4, 5}
	var b = []byte{1, 2, 3, 4, 5, 0}
	var c = []byte{2, 3, 4, 5, 0, 1}

	lshift(a)
	assert.EqualValues(t, b, a)
	lshift(a)
	assert.EqualValues(t, c, a)
}

func TestLshiftPartial(t *testing.T) {
	var a = []byte{0, 1, 2, 3, 4, 5}
	var c = []byte{0, 1, 3, 4, 5, 2}

	lshift(a[2:])
	assert.EqualValues(t, c, a)
}

func TestLshiftN(t *testing.T) {
	var a = []byte{0, 1, 2, 3, 4, 5}
	var c = []byte{2, 3, 4, 5, 0, 1}

	lshiftN(a, 2)
	assert.EqualValues(t, c, a)
}

func TestLshiftNPartial(t *testing.T) {
	var a = []byte{0, 1, 2, 3, 4, 5}
	var c = []byte{0, 1, 4, 5, 2, 3}

	lshiftN(a[2:], 2)
	assert.EqualValues(t, c, a)
}

func TestLshiftNCircle(t *testing.T) {
	var a = []byte{0, 1, 2, 3, 4, 5}
	var c = []byte{0, 1, 2, 3, 4, 5}

	lshiftN(a, len(a))
	assert.EqualValues(t, c, a)
}

func TestRshift(t *testing.T) {
	var a = []byte{0, 1, 2, 3, 4, 5}
	var b = []byte{5, 0, 1, 2, 3, 4}
	var c = []byte{4, 5, 0, 1, 2, 3}

	rshift(a)
	assert.EqualValues(t, b, a)
	rshift(a)
	assert.EqualValues(t, c, a)
}

func TestRshiftPartial(t *testing.T) {
	var a = []byte{0, 1, 2, 3, 4, 5}
	var c = []byte{0, 1, 5, 2, 3, 4}

	rshift(a[2:])
	assert.EqualValues(t, c, a)
}

func TestRshiftN(t *testing.T) {
	var a = []byte{0, 1, 2, 3, 4, 5}
	var c = []byte{4, 5, 0, 1, 2, 3}

	rshiftN(a, 2)
	assert.EqualValues(t, c, a)
}

func TestRshiftNPartial(t *testing.T) {
	var a = []byte{0, 1, 2, 3, 4, 5}
	var c = []byte{0, 1, 4, 5, 2, 3}

	rshiftN(a[2:], 2)
	assert.EqualValues(t, c, a)
}

func TestRshiftNCircle(t *testing.T) {
	var a = []byte{0, 1, 2, 3, 4, 5}
	var c = []byte{0, 1, 2, 3, 4, 5}

	rshiftN(a, len(a))
	assert.EqualValues(t, c, a)
}

func TestDup(t *testing.T) {
	var a = []byte{0, 1, 2, 3, 4, 5}

	b := dup(a)
	assert.EqualValues(t, b, a)
	assert.Equal(t, b, a)
	assert.True(t, bytes.Equal(a, b))
}

var gb []byte

func BenchmarkLshiftN(b *testing.B) {
	var aa = []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		lshiftN(aa, len(aa))
	}
}

func BenchmarkRshiftN(b *testing.B) {
	var aa = []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		rshiftN(aa, len(aa))
	}
}

func BenchmarkDup(b *testing.B) {
	var aa = []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	var bb []byte

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		bb = dup(aa)
	}
	gb = bb
}
