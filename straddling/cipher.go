package straddling

import (
	"bytes"
	"crypto/cipher"
	"fmt"
	"github.com/keltia/cipher"
	"log"
)

const (
	alphabetTxt = "ABCDEFGHIJKLMNOPQRSTUVWXYZ/-"
)

var (
	allcipher = []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	freq      = []byte{'E', 'S', 'A', 'N', 'T', 'I', 'R', 'U'}
)

type straddlingcheckerboard struct {
	key    string
	longc  []byte
	shortc []byte
	full   string
	enc    map[byte]string
	dec    map[string]byte
}

func NewCipher(key string, chrs string) (cipher.Block, error) {
	if key == "" || chrs == "" {
		return nil, fmt.Errorf("neither key nor long can be empty")
	}

	longc := []byte{chrs[0], chrs[1]}
	c := &straddlingcheckerboard{
		key:    key,
		full:   crypto.Shuffle(key, alphabetTxt),
		longc:  longc,
		shortc: extract(allcipher, longc),
		enc:    make(map[byte]string),
		dec:    make(map[string]byte),
	}
	c.expandKey()
	return c, nil
}

// times11 generates the set of c[0..9] aka "00"-"09" or "30"-"39" by appending into tmp
func times11(c byte) []string {
	var tmp []string

	if c == '0' {
		for _, b := range allcipher {
			tmp = append(tmp, string(b))
		}
	} else {
		for _, b := range allcipher {
			tmp = append(tmp, string(c)+string(b))
		}
	}
	return tmp
}

// times10 generates the set of c[0..9] aka "00"-"09" or "30"-"39", tmp is pre-allocated
func times10(c byte) []string {
	var tmp = make([]string, 10)

	if c == '0' {
		for i, b := range allcipher {
			tmp[i] = string(b)
		}
	} else {
		for i, b := range allcipher {
			tmp[i] = string(c) + string(b)
		}
	}
	return tmp
}

// extract returns all cipher numbers not in the "two" set
func extract(set, two []byte) []byte {
	f := func(r rune) rune {
		if bytes.ContainsRune(two, r) {
			return -1
		}
		return r
	}
	return bytes.Map(f, set)
}

func settimes11(set []byte) []string {

	longc := []string{}

	// Generate all double digits ciphertext bigrams
	for _, v := range set {
		tmpc := times10(byte(v))
		longc = append(longc, tmpc...)
	}
	return longc
}

func settimes10(set []byte) []string {

	longc := make([]string, 20)

	// Generate all double digits ciphertext bigrams

	copy(longc, times10(set[0]))
	copy(longc[10:], times10(set[1]))
	return longc
}

func (c *straddlingcheckerboard) expandKey() {
	shortc := times10('0')
	longc := settimes10(c.longc)

	// Assign a mono/bigram to each letter in the shuffled key
	i := 0
	j := 0
	bfull := bytes.NewBufferString(c.full).Bytes()
	for _, ch := range bfull {
		if bytes.Contains(freq, []byte{ch}) {
			c.enc[ch] = shortc[i]
			c.dec[shortc[i]] = ch
			i++
		} else {
			c.enc[ch] = longc[j]
			c.dec[longc[j]] = ch
			j++
		}
	}
}

func (c *straddlingcheckerboard) BlockSize() int {
	return len(c.key)
}

func (c *straddlingcheckerboard) Encrypt(dst, src []byte) {
	var ct bytes.Buffer

	plen := len(src)
	for i := 0; i < plen; i++ {
		if src[i] >= '0' && src[i] <= '9' {
			ct.WriteString(c.enc['/'])
			ct.WriteString(string(src[i])) // yeah, this is plaintext
			ct.WriteString(string(src[i]))
			ct.WriteString(c.enc['/'])
		} else {
			ct.WriteString(c.enc[src[i]])
		}
	}
	copy(dst, ct.Bytes())
}

func (c *straddlingcheckerboard) Decrypt(dst, src []byte) {
	var (
		pt  bytes.Buffer
		ptc byte
	)

	ct := bytes.NewBuffer(src)
	plen := len(src)
	j := 0
	for i := 0; i < plen; {
		var (
			db = []byte{0, 0}
		)

		// Check whether we have a short or long codegroup
		ch, _ := ct.ReadByte()
		if ch == c.longc[0] || ch == c.longc[1] {
			// Long
			db[0] = ch
			db[1], _ = ct.ReadByte()
			ptc = c.dec[string(db)]
			i += 2
		} else {
			// Short
			ptc = c.dec[string(ch)]
			i++
		}

		/*
			Check for the '/' code group
			followed by <pt><pt><b0><b1> where "b0b1" is the encoding group for /

			Can be a regular / so check afterward for matching /. If not, it gets complicated.

			If I am to support this, I will have to rollback up to 4 bytes and managing all the
			corner cases there, not sure it is worth it.

			Encrypt() can not generate such degenerate cases so I might just leave it at that.
		*/
		if ptc == '/' {
			/*
				Read the next four bytes representing <n0><n0><b0><b1> with <b0><b1> being the
				code word for / (it will never be one of the short codes.
			*/
			var numb = []byte{0, 0, 0, 0}

			n, err := ct.Read(numb)
			if err != nil || n != 4 {
				ptc = 'X'
				break
			}
			if numb[0] == numb[1] && bytes.Equal(db, []byte{numb[2], numb[3]}) {
				// We have a number
				ptc = numb[0]
				i += 4
			}
		}
		pt.WriteByte(ptc)
		j++
	}
	copy(dst, pt.Bytes())
}

// verbose displays only if fVerbose is set
func message(str string, a ...interface{}) {
	log.Printf(str, a...)
}
