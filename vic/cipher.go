package vic

import (
	"bytes"
	"crypto/cipher"
	"github.com/keltia/cipher"
	"github.com/keltia/cipher/transposition"
	"log"
	"sort"
)

/*
Full description & test vectors: http://www.quadibloc.com/crypto/pp1324.htm
*/

var (
	allcipher  = []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	enumDigits = []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}

	// English frequent letters
	freq = []byte{'A', 'T', 'O', 'N', 'E', 'S', 'I', 'R'}
)

type viccipher struct {
	ind    string
	phrase string
	persn  string

	imsg   []byte
	ikey5  []byte
	first  []byte
	second []byte
	third  []byte
	sckey  []byte
	tpkeys []byte

	// First transposition
	firsttp *cipher.Block
}

func NewCipher(persn, ind, phrase string, imsg string) (cipher.Block, error) {
	c := &viccipher{
		ind:    ind,
		persn:  persn,
		phrase: phrase,
		imsg:   str2int(imsg),
		ikey5:  str2int(ind[:5]),
	}
	c.expandKey()

	// We have two transpositions, first one is regular
	transp, err := transposition.NewCipher(string(c.second))
	if err != nil {
		return nil, err
	}

	c.firsttp = &transp
	return c, nil
}

func (c *viccipher) expandKey() {
	// First phase
	//message("ind=%s", c.ind)

	ph1 := toNumericOne(c.phrase[:10])
	ph2 := toNumericOne(c.phrase[10:])
	//message("ph1=%v ph2=%v", ph1, ph2)

	res := submod10(c.imsg, c.ikey5)
	c.first = expand5to10(res)
	//message("res=%v ikey5=%v first=%v", res, c.ikey5, c.first)

	// Second phase
	tmp := addmod10(c.first, ph1)
	//message("tmp=%v", tmp)
	c.second = firstEncode(tmp, ph2) // this will be the key for a transposition later
	//message("second=%v", c.second)

	var tptmp bytes.Buffer

	// Third phase
	r := crypto.Dup(c.second)
	for i := 0; i < 5; i++ {
		r = chainadd(r) // We store the intermediate results
		//message("r=%v", r)
		tptmp.Write(r)
	}
	tpkeys := tptmp.Bytes()
	//message("tpkeys=%v", tpkeys)
	c.tpkeys = tpkeys

	fourth := crypto.ToNumeric(string(r))
	//message("fourth=%v", fourth)
	c.third = r // Last one is stored

	// The last one is the Straddling Cherkerboard key
	c.sckey = fourth
}

// toNumericOne is ToNumeric interalized to return 1-based arrays
func toNumericOne(key string) []byte {
	letters := bytes.NewBufferString(key).Bytes()
	sorted := crypto.Dup(letters)
	sort.Slice(sorted, func(i, j int) bool { return sorted[i] < sorted[j] })

	f := func(c rune) rune {
		k := bytes.Index(sorted, []byte{byte(c)})
		sorted[k] = 0
		return rune((k + 1) % 10)
	}
	ar := bytes.Map(f, letters)
	return ar
}

func str2int(str string) []byte {
	var b bytes.Buffer

	for _, ch := range str {
		b.WriteByte(byte(ch - '0'))
	}
	return b.Bytes()
}

func addmod10(a, b []byte) []byte {
	var c bytes.Buffer
	for i, v := range a {
		c.WriteByte((v + b[i]) % 10)
	}
	return c.Bytes()
}

func submod10(a, b []byte) []byte {
	var c bytes.Buffer
	for i, v := range a {
		c.WriteByte((v - b[i] + 10) % 10)
	}
	return c.Bytes()
}

func chainadd(a []byte) []byte {

	b := crypto.Dup(a)
	l := len(a)
	for i, v := range b {
		b[i] = (v + b[(i+1)%l]) % 10
	}
	return b
}

func expand5to10(a []byte) []byte {
	b := chainadd(a)
	c := bytes.NewBuffer(a)
	c.Write(b)
	return c.Bytes()
}

func firstEncode(a, b []byte) []byte {
	var r bytes.Buffer

	for _, v := range a {
		r.WriteByte(b[(v+10)%10-1])
	}
	return r.Bytes()
}

func (c *viccipher) BlockSize() int {
	return 1
}

func (c *viccipher) Encrypt(dst, src []byte) {

}

func (c *viccipher) Decrypt(dst, src []byte) {

}

// verbose displays only if fVerbose is set
func message(str string, a ...interface{}) {
	log.Printf(str, a...)
}
