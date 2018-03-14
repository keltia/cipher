package straddling

import (
	"bytes"
	"crypto/cipher"
	"github.com/stretchr/testify/assert"
	"reflect"
	"strings"
	"testing"
)

func TestNewCipher(t *testing.T) {
	c, err := NewCipher("ARABESQUE", "89")

	assert.NotNil(t, c)
	assert.NoError(t, err)
	assert.Implements(t, (*cipher.Block)(nil), c)

	cc := c.(*straddlingcheckerboard)
	assert.Equal(t, "ACKVRDLWBFMXEGNYSHOZQIP/UJT-", cc.full)
	assert.EqualValues(t, []byte{'0', '1', '2', '3', '4', '5', '6', '7'}, cc.shortc)
	assert.EqualValues(t, []byte{'8', '9'}, cc.longc)
}

func TestNewCipher2(t *testing.T) {
	c, err := NewCipher("ARABESQUE", "")

	assert.Nil(t, c)
	assert.Error(t, err)
}

func TestNewCipher3(t *testing.T) {
	c, err := NewCipher("", "89")

	assert.Nil(t, c)
	assert.Error(t, err)
}

var TestExpandKeyData = []struct {
	enc map[byte]string
	dec map[string]byte
}{
	{map[byte]string{
		byte('V'): "82",
		byte('K'): "81",
		byte('W'): "85",
		byte('L'): "84",
		byte('A'): "0",
		byte('-'): "99",
		byte('X'): "89",
		byte('M'): "88",
		byte('B'): "86",
		byte('Y'): "91",
		byte('N'): "3",
		byte('C'): "80",
		byte('/'): "97",
		byte('Z'): "94",
		byte('O'): "93",
		byte('D'): "83",
		byte('P'): "96",
		byte('E'): "2",
		byte('Q'): "95",
		byte('F'): "87",
		byte('G'): "90",
		byte('R'): "1",
		byte('H'): "92",
		byte('S'): "4",
		byte('T'): "7",
		byte('I'): "5",
		byte('J'): "98",
		byte('U'): "6",
	},
		map[string]byte{
			"5":  byte('I'),
			"93": byte('O'),
			"82": byte('V'),
			"99": byte('-'),
			"88": byte('M'),
			"0":  byte('A'),
			"6":  byte('U'),
			"94": byte('Z'),
			"83": byte('D'),
			"89": byte('X'),
			"1":  byte('R'),
			"7":  byte('T'),
			"95": byte('Q'),
			"84": byte('L'),
			"90": byte('G'),
			"2":  byte('E'),
			"96": byte('P'),
			"85": byte('W'),
			"91": byte('Y'),
			"3":  byte('N'),
			"80": byte('C'),
			"97": byte('/'),
			"86": byte('B'),
			"92": byte('H'),
			"4":  byte('S'),
			"81": byte('K'),
			"98": byte('J'),
			"87": byte('F'),
		}},
}

func TestExpandKey(t *testing.T) {
	c, err := NewCipher("ARABESQUE", "89")

	assert.NotNil(t, c)
	assert.NoError(t, err)
	assert.Implements(t, (*cipher.Block)(nil), c)

	cc := c.(*straddlingcheckerboard)
	cp := TestExpandKeyData[0]

	assert.EqualValues(t, []byte{'0', '1', '2', '3', '4', '5', '6', '7'}, cc.shortc)
	assert.EqualValues(t, []byte{'8', '9'}, cc.longc)

	if !reflect.DeepEqual(cp.enc, cc.enc) {
		t.Errorf("%v is different from %v", cc.enc, cp.enc)
	}
	if !reflect.DeepEqual(cp.dec, cc.dec) {
		t.Errorf("%v is different from %v", cc.dec, cp.dec)
	}
}

func TestExpandKey1(t *testing.T) {
	c, err := NewCipher("ARABESQUE", "36")

	assert.NotNil(t, c)
	assert.NoError(t, err)
	assert.Implements(t, (*cipher.Block)(nil), c)

	cc := c.(*straddlingcheckerboard)

	assert.EqualValues(t, []byte{'0', '1', '2', '4', '5', '7', '8', '9'}, cc.shortc)
	assert.EqualValues(t, []byte{'3', '6'}, cc.longc)

	cp := TestExpandKeyData[0]
	if !reflect.DeepEqual(cp.enc, cc.enc) {
		t.Errorf("%v is different from %v", cc.enc, cp.enc)
	}
}

func TestTimes10(t *testing.T) {
	trente := times10('3')
	assert.EqualValues(t, []string{"30", "31", "32", "33", "34", "35", "36", "37", "38", "39"}, trente)

	dix := times10('1')
	assert.EqualValues(t, []string{"10", "11", "12", "13", "14", "15", "16", "17", "18", "19"}, dix)
}

func TestTimes11(t *testing.T) {
	trente := times11('3')
	assert.EqualValues(t, []string{"30", "31", "32", "33", "34", "35", "36", "37", "38", "39"}, trente)

	dix := times11('1')
	assert.EqualValues(t, []string{"10", "11", "12", "13", "14", "15", "16", "17", "18", "19"}, dix)
}

func TestSettimes10(t *testing.T) {
	threeone := settimes10([]byte{'3', '1'})
	assert.EqualValues(t,
		[]string{"30", "31", "32", "33", "34", "35", "36", "37", "38", "39",
			"10", "11", "12", "13", "14", "15", "16", "17", "18", "19"}, threeone)

	huitneuf := settimes10([]byte{'8', '9'})
	assert.EqualValues(t,
		[]string{"80", "81", "82", "83", "84", "85", "86", "87", "88", "89",
			"90", "91", "92", "93", "94", "95", "96", "97", "98", "99"}, huitneuf)
}

func TestSettimes11(t *testing.T) {
	threeone := settimes11([]byte{'3', '1'})
	assert.EqualValues(t,
		[]string{"30", "31", "32", "33", "34", "35", "36", "37", "38", "39",
			"10", "11", "12", "13", "14", "15", "16", "17", "18", "19"}, threeone)

	huitneuf := settimes11([]byte{'8', '9'})
	assert.EqualValues(t,
		[]string{"80", "81", "82", "83", "84", "85", "86", "87", "88", "89",
			"90", "91", "92", "93", "94", "95", "96", "97", "98", "99"}, huitneuf)
}

func TestExtract(t *testing.T) {
	shortc := extract(allcipher, []byte{'2', '5'})
	assert.EqualValues(t, []byte{'0', '1', '3', '4', '6', '7', '8', '9'}, shortc)

	shortc = extract(allcipher, []byte{'1', '6'})
	assert.EqualValues(t, []byte{'0', '2', '3', '4', '5', '7', '8', '9'}, shortc)

	shortc = extract(allcipher, []byte{'4', '2'})
	assert.EqualValues(t, []byte{'0', '1', '3', '5', '6', '7', '8', '9'}, shortc)
}

func TestStraddlingcheckerboard_BlockSize(t *testing.T) {
	c, _ := NewCipher("ARABESQUE", "89")
	assert.Equal(t, len("ARABESQUE"), c.BlockSize())
}

var TestSCEncryptData = []struct {
	key  string
	chrs string
	pt   string
	ct   string
}{
	{"ARABESQUE", "89", "ATTACKAT2AM", "0770808107972297088"},
	{"ARABESQUE", "36", "ATTACKAT2AM", "0990303109672267038"},
	{"ARABESQUE", "37", "IFYOUCANREADTHIS", "6377173830041203397265"},
	{"ARABESQUE", "89", "ATTACK", "07708081"},
	{"SUBWAY", "89", "TOLKIEN", "6819388137"},
	{"PORTABLE", "89", "RETRIBUTION", "1721693526840"},
}

func TestStraddlingcheckerboard_Encrypt(t *testing.T) {
	for _, cp := range TestSCEncryptData {
		key := cp.key
		plain := cp.pt
		chrs := cp.chrs

		dst := make([]byte, len(plain)*2)
		c, _ := NewCipher(key, chrs)
		c.Encrypt(dst, bytes.NewBufferString(plain).Bytes())

		// Have to remove right-hand \x00
		sct := strings.TrimRight(string(dst), "\x00")
		assert.Equal(t, cp.ct, sct)
	}
}

func TestStraddlingcheckerboard_Decrypt(t *testing.T) {
	for _, cp := range TestSCEncryptData {
		key := cp.key
		ct := cp.ct
		chrs := cp.chrs

		dst := make([]byte, len(ct))
		c, _ := NewCipher(key, chrs)
		c.Decrypt(dst, bytes.NewBufferString(ct).Bytes())

		// Have to remove right-hand \x00
		spt := strings.TrimRight(string(dst), "\x00")
		assert.Equal(t, cp.pt, spt)
	}

}

// -- benchmarks

var gc cipher.Block

func BenchmarkNewCipher(b *testing.B) {
	var c cipher.Block

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c, _ = NewCipher("ARABESQUE", "89")
	}
	gc = c
}

var gb []string

func BenchmarkTimes10(b *testing.B) {
	var zero []string

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		zero = times10(0)
	}
	gb = zero
}

func BenchmarkSettimes10(b *testing.B) {
	var zero []string

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		zero = settimes10([]byte{8, 9})
	}
	gb = zero
}

func BenchmarkTimes11(b *testing.B) {
	var zero []string

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		zero = times11(0)
	}
	gb = zero
}

func BenchmarkSettimes11(b *testing.B) {
	var zero []string

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		zero = settimes11([]byte{8, 9})
	}
	gb = zero
}

var gsh []byte

func BenchmarkExtract(b *testing.B) {
	var shortc []byte

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		shortc = extract(allcipher, []byte{'4', '2'})
	}
	gsh = shortc
}

func BenchmarkStraddlingcheckerboard_Encrypt(b *testing.B) {
	key := TestSCEncryptData[0].key
	chrs := "89"

	c, _ := NewCipher(key, chrs)
	pt := TestSCEncryptData[0].pt
	dst := make([]byte, 2*len(pt))
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.Encrypt(dst, bytes.NewBufferString(pt).Bytes())
	}
}

func BenchmarkStraddlingcheckerboard_Decrypt(b *testing.B) {
	key := TestSCEncryptData[0].key
	chrs := "89"

	c, _ := NewCipher(key, chrs)
	ct := TestSCEncryptData[0].ct
	dst := make([]byte, len(ct))
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.Encrypt(dst, bytes.NewBufferString(ct).Bytes())
	}
}
