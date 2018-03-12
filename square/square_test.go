package square

import (
	"bytes"
	"crypto/cipher"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

var TestSQData = []struct {
	key  string
	chrs string
	enc  map[byte]string
	dec  map[string]byte
}{
	{
		"PORTABLE",
		"ADFGVX",
		map[byte]string{
			byte('6'): "XF",
			byte('V'): "GG",
			byte('K'): "FG",
			byte('7'): "XG",
			byte('W'): "GV",
			byte('L'): "DA",
			byte('A'): "AV",
			byte('8'): "XV",
			byte('X'): "GX",
			byte('M'): "FV",
			byte('B'): "AX",
			byte('9'): "XX",
			byte('Y'): "VA",
			byte('N'): "FX",
			byte('C'): "DF",
			byte('Z'): "VD",
			byte('D'): "DG",
			byte('O'): "AD",
			byte('0'): "VF",
			byte('E'): "DD",
			byte('P'): "AA",
			byte('1'): "VG",
			byte('Q'): "GA",
			byte('F'): "DV",
			byte('2'): "VV",
			byte('G'): "DX",
			byte('R'): "AF",
			byte('3'): "VX",
			byte('S'): "GD",
			byte('H'): "FA",
			byte('4'): "XA",
			byte('I'): "FD",
			byte('T'): "AG",
			byte('5'): "XD",
			byte('U'): "GF",
			byte('J'): "FF",
		},
		map[string]byte{
			"XF": byte('6'),
			"GG": byte('V'),
			"FG": byte('K'),
			"XG": byte('7'),
			"GV": byte('W'),
			"DA": byte('L'),
			"AV": byte('A'),
			"XV": byte('8'),
			"GX": byte('X'),
			"FV": byte('M'),
			"AX": byte('B'),
			"XX": byte('9'),
			"VA": byte('Y'),
			"FX": byte('N'),
			"DF": byte('C'),
			"VD": byte('Z'),
			"DG": byte('D'),
			"AD": byte('O'),
			"VF": byte('0'),
			"DD": byte('E'),
			"AA": byte('P'),
			"VG": byte('1'),
			"GA": byte('Q'),
			"DV": byte('F'),
			"VV": byte('2'),
			"DX": byte('G'),
			"AF": byte('R'),
			"VX": byte('3'),
			"GD": byte('S'),
			"FA": byte('H'),
			"XA": byte('4'),
			"FD": byte('I'),
			"AG": byte('T'),
			"XD": byte('5'),
			"GF": byte('U'),
			"FF": byte('J'),
		},
	},
	{
		"ARABESQUE",
		"012345",
		map[byte]string{
			byte('6'): "52",
			byte('V'): "33",
			byte('K'): "22",
			byte('7'): "53",
			byte('W'): "34",
			byte('L'): "23",
			byte('A'): "00",
			byte('8'): "54",
			byte('X'): "35",
			byte('M'): "24",
			byte('B'): "02",
			byte('9'): "55",
			byte('Y'): "40",
			byte('N'): "25",
			byte('C'): "11",
			byte('Z'): "41",
			byte('O'): "30",
			byte('D'): "12",
			byte('0'): "42",
			byte('P'): "31",
			byte('E'): "03",
			byte('1'): "43",
			byte('F'): "13",
			byte('Q'): "05",
			byte('2'): "44",
			byte('G'): "14",
			byte('R'): "01",
			byte('3'): "45",
			byte('H'): "15",
			byte('S'): "04",
			byte('4'): "50",
			byte('T'): "32",
			byte('I'): "20",
			byte('5'): "51",
			byte('J'): "21",
			byte('U'): "10",
		},
		map[string]byte{
			"52": byte('6'),
			"33": byte('V'),
			"22": byte('K'),
			"53": byte('7'),
			"34": byte('W'),
			"23": byte('L'),
			"00": byte('A'),
			"54": byte('8'),
			"35": byte('X'),
			"24": byte('M'),
			"02": byte('B'),
			"55": byte('9'),
			"40": byte('Y'),
			"25": byte('N'),
			"11": byte('C'),
			"41": byte('Z'),
			"30": byte('O'),
			"12": byte('D'),
			"42": byte('0'),
			"31": byte('P'),
			"03": byte('E'),
			"43": byte('1'),
			"13": byte('F'),
			"05": byte('Q'),
			"44": byte('2'),
			"14": byte('G'),
			"01": byte('R'),
			"45": byte('3'),
			"15": byte('H'),
			"04": byte('S'),
			"50": byte('4'),
			"32": byte('T'),
			"20": byte('I'),
			"51": byte('5'),
			"21": byte('J'),
			"10": byte('U'),
		},
	},
}

func TestExpandKey(t *testing.T) {
	for _, cp := range TestSQData {

		c, _ := NewCipher(cp.key, cp.chrs)

		cc := c.(*squarecipher)
		cc.expandKey()

		if !reflect.DeepEqual(cp.enc, cc.enc) {
			t.Errorf("%v is different from %v", cp.enc, cc.enc)
		}
		if !reflect.DeepEqual(cp.dec, cc.dec) {
			t.Errorf("%v is different from %v", cp.dec, cc.dec)
		}
	}
}

func TestExpandKey1(t *testing.T) {
	for _, cp := range TestSQData {

		c, _ := NewCipher(cp.key, cp.chrs)

		cc := c.(*squarecipher)
		cc.expandKey1()

		if !reflect.DeepEqual(cp.enc, cc.enc) {
			t.Errorf("%v is different from %v", cp.enc, cc.enc)
		}
		if !reflect.DeepEqual(cp.dec, cc.dec) {
			t.Errorf("%v is different from %v", cp.dec, cc.dec)
		}
	}
}

func TestExpandKey2(t *testing.T) {
	for _, cp := range TestSQData {

		c, _ := NewCipher(cp.key, cp.chrs)

		cc := c.(*squarecipher)
		cc.expandKey2()

		if !reflect.DeepEqual(cp.enc, cc.enc) {
			t.Errorf("%v is different from %v", cp.enc, cc.enc)
		}
		if !reflect.DeepEqual(cp.dec, cc.dec) {
			t.Errorf("%v is different from %v", cp.dec, cc.dec)
		}
	}
}

func TestExpandKey3(t *testing.T) {
	for _, cp := range TestSQData {

		c, _ := NewCipher(cp.key, cp.chrs)

		cc := c.(*squarecipher)
		cc.expandKey3()

		if !reflect.DeepEqual(cp.enc, cc.enc) {
			t.Errorf("%v is different from %v", cp.enc, cc.enc)
		}
		if !reflect.DeepEqual(cp.dec, cc.dec) {
			t.Errorf("%v is different from %v", cp.dec, cc.dec)
		}
	}
}

func TestNewCipher(t *testing.T) {
	for _, cp := range TestSQData {
		c, err := NewCipher(cp.key, cp.chrs)

		assert.NotNil(t, c)
		assert.NoError(t, err)
	}
}

func TestNewCipher2(t *testing.T) {
	c, err := NewCipher("", "012345")

	assert.Empty(t, c)
	assert.Error(t, err)
}

func TestNewCipher3(t *testing.T) {
	c, err := NewCipher("SUBWAY", "")

	assert.Empty(t, c)
	assert.Error(t, err)
}

func TestSquarecipher_BlockSize(t *testing.T) {
	for _, cp := range TestSQData {
		c, err := NewCipher(cp.key, cp.chrs)
		assert.NotNil(t, c)
		assert.NoError(t, err)

		assert.Equal(t, len(cp.key), c.BlockSize())
	}
}

var TestSQDataED = []struct {
	key  string
	chrs string
	pt   string
	ct   string
}{
	{"PORTABLE", "ADFGVX", "ATTACKATDAWN", "AVAGAGAVDFFGAVAGDGAVGVFX"},
	{"ARABESQUE", "012345", "ATTACKATDAWN", "003232001122003212003425"},
}

func TestSquarecipher_Encrypt(t *testing.T) {
	for _, cp := range TestSQDataED {
		c, err := NewCipher(cp.key, cp.chrs)
		assert.NotNil(t, c)
		assert.NoError(t, err)

		src := bytes.NewBufferString(cp.pt).Bytes()
		dst := make([]byte, 2*len(cp.pt))
		c.Encrypt(dst, src)
		assert.EqualValues(t, cp.ct, string(dst))
	}
}

func TestSquarecipher_Decrypt(t *testing.T) {
	for _, cp := range TestSQDataED {
		c, err := NewCipher(cp.key, cp.chrs)
		assert.NotNil(t, c)
		assert.NoError(t, err)

		src := bytes.NewBufferString(cp.ct).Bytes()
		dst := make([]byte, len(cp.ct)/2)
		c.Decrypt(dst, src)
		assert.EqualValues(t, cp.pt, string(dst))
	}

}

// -- benchmarks

func BenchmarkExpandKey(b *testing.B) {
	key := "PORTABLE"
	chrs := "ADFGVX"

	c, _ := NewCipher(key, chrs)
	cc := c.(*squarecipher)

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		cc.expandKey()
	}
}

func BenchmarkExpandKey1(b *testing.B) {
	key := "PORTABLE"
	chrs := "ADFGVX"

	c, _ := NewCipher(key, chrs)
	cc := c.(*squarecipher)

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		cc.expandKey1()
	}
}

func BenchmarkExpandKey2(b *testing.B) {
	key := "PORTABLE"
	chrs := "ADFGVX"

	c, _ := NewCipher(key, chrs)
	cc := c.(*squarecipher)

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		cc.expandKey2()
	}
}

func BenchmarkExpandKey3(b *testing.B) {
	key := "PORTABLE"
	chrs := "ADFGVX"

	c, _ := NewCipher(key, chrs)
	cc := c.(*squarecipher)

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		cc.expandKey3()
	}
}

var gc cipher.Block

func BenchmarkNewCipher(b *testing.B) {
	var c cipher.Block

	key := "PORTABLE"
	chrs := "ADFGVX"

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c, _ = NewCipher(key, chrs)
	}
	gc = c
}

func BenchmarkSquarecipher_Encrypt(b *testing.B) {
	key := "PORTABLE"
	chrs := "ADFGVX"

	c, _ := NewCipher(key, chrs)

	cp := TestSQDataED[0]
	src := bytes.NewBufferString(cp.pt).Bytes()
	dst := make([]byte, len(cp.pt)*2)

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.Encrypt(dst, src)
	}
}

func BenchmarkSquarecipher_Decrypt(b *testing.B) {
	key := "PORTABLE"
	chrs := "ADFGVX"

	c, _ := NewCipher(key, chrs)

	cp := TestSQDataED[0]
	src := bytes.NewBufferString(cp.ct).Bytes()
	dst := make([]byte, len(cp.ct)/2)

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.Decrypt(dst, src)
	}
}
