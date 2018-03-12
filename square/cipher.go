package square

import (
	"bytes"
	"crypto/cipher"
	"github.com/keltia/cipher"
)

const (
	Base36 = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

type squarecipher struct {
	key   string
	chrs  string
	alpha []byte
	enc   map[byte]string
	dec   map[string]byte
}

func NewCipher(key string, chrs string) (cipher.Block, error) {
	alpha := bytes.NewBufferString(crypto.Condense(key + Base36)).Bytes()

	c := &squarecipher{
		key:   key,
		chrs:  chrs,
		alpha: alpha,
		enc:   make(map[byte]string, len(alpha)),
		dec:   make(map[string]byte, len(alpha)),
	}
	c.expandKey()
	return c, nil
}

// First version of expandKey
func (c *squarecipher) expandKey2() {
	for i := range c.chrs {
		for j := range c.chrs {
			ind := i*len(c.chrs) + j
			c.enc[c.alpha[ind]] = string(c.chrs[i]) + string(c.chrs[j])
			c.dec[string(c.chrs[i])+string(c.chrs[j])] = c.alpha[ind]
		}
	}
}

// Let's try with a Buffer
func (c *squarecipher) expandKey1() {
	var bigr = bytes.Buffer{}

	for i := range c.chrs {
		for j := range c.chrs {

			bigr.Write([]byte{c.chrs[i], c.chrs[j]})

			ind := i*len(c.chrs) + j
			c.enc[c.alpha[ind]] = bigr.String()
			c.dec[bigr.String()] = c.alpha[ind]

			bigr.Reset()
		}
	}
}

// Fixed []byte maybe?
func (c *squarecipher) expandKey3() {
	var bigr = []byte{0, 0}

	klen := len(c.chrs)
	for i := range c.chrs {
		for j := range c.chrs {

			copy(bigr, []byte{c.chrs[i], c.chrs[j]})

			ind := i*klen + j
			c.enc[c.alpha[ind]] = string(bigr)
			c.dec[string(bigr)] = c.alpha[ind]
		}
	}
}

// copy() for two bytes is overrated
func (c *squarecipher) expandKey() {
	var bigr = []byte{0, 0}

	klen := len(c.chrs)
	for i := range c.chrs {
		for j := range c.chrs {

			bigr[0] = c.chrs[i]
			bigr[1] = c.chrs[j]

			ind := i*klen + j
			c.enc[c.alpha[ind]] = string(bigr)
			c.dec[string(bigr)] = c.alpha[ind]
		}
	}
}

func (c *squarecipher) BlockSize() int {
	return len(c.key)
}

func (c *squarecipher) Encrypt(dst, src []byte) {
	plen := len(src)
	for i := 0; i < plen; i++ {
		ct := c.enc[src[i]]
		dst[i*2] = ct[0]
		dst[i*2+1] = ct[1]
	}
}

func (c *squarecipher) Decrypt(dst, src []byte) {
	clen := len(src)
	for i := 0; i < clen; i += 2 {
		pt := string([]byte{src[i], src[i+1]})
		dst[i/2] = c.dec[pt]
	}
}

/*
// verbose displays only if fVerbose is set
func message(str string, a ...interface{}) {
	log.Printf(str, a...)
}
*/
