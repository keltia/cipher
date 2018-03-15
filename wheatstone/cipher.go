package wheatstone

import (
	"bytes"
	"crypto/cipher"
	"fmt"
	"github.com/keltia/cipher"
	"log"
)

const (
	alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lenPL    = len(alphabet) + 1
	lenCT    = len(alphabet)
)

type wheatstone struct {
	pkey, ckey string
	aplw, actw []byte
	start      byte
	curpos     int
	ctpos      int
}

// NewCipher creates a new cipher with the provided keys
func NewCipher(start byte, pkey, ckey string) (cipher.Block, error) {
	if pkey == "" ||
		ckey == "" {
		return &wheatstone{}, fmt.Errorf("keys can not be empty")
	}

	// Transform with key
	pkey = "+" + crypto.Shuffle(pkey, alphabet)
	ckey = crypto.Shuffle(ckey, alphabet)

	c := &wheatstone{
		start:  start,
		curpos: 0,
		pkey:   pkey,
		ckey:   ckey,
		aplw:   bytes.NewBufferString(pkey).Bytes(),
		actw:   bytes.NewBufferString(ckey).Bytes(),
	}
	c.ctpos = bytes.IndexByte(c.actw, start)

	//message("c=%#v", c)
	return c, nil
}

func (c *wheatstone) BlockSize() int {
	return 1
}

func (c *wheatstone) encode(ch byte) byte {
	var off int

	a := bytes.IndexByte(c.aplw, ch)
	if a <= c.curpos {
		off = (a + lenPL) - c.curpos
	} else {
		off = a - c.curpos
	}
	c.curpos = a
	c.ctpos = (c.ctpos + off) % lenCT
	return c.actw[c.ctpos]
}

func (c *wheatstone) decode(ch byte) byte {
	var off int

	a := bytes.IndexByte(c.actw, ch)
	if a <= c.ctpos {
		off = (a + lenCT) - c.curpos
	} else {
		off = a - c.curpos
	}
	c.ctpos = a
	c.curpos = (c.curpos + off) % lenPL
	return c.aplw[c.curpos]
}

func (c *wheatstone) Encrypt(dst, src []byte) {
	c.reset()
	for i, ch := range src {
		dst[i] = c.encode(ch)
	}
}

func (c *wheatstone) Decrypt(dst, src []byte) {
	c.reset()
	for i, ch := range src {
		dst[i] = c.decode(ch)
	}
}

/*
This is necessary because the wheatstone object retain state across calls
*/
// Reset state to the beginning.
func (c *wheatstone) reset() {
	// Transform with key
	pkey := "+" + crypto.Shuffle(c.pkey, alphabet)
	ckey := crypto.Shuffle(c.ckey, alphabet)

	c.aplw = bytes.NewBufferString(pkey).Bytes()
	c.actw = bytes.NewBufferString(ckey).Bytes()

	c.ctpos = bytes.IndexByte(c.actw, c.start)
}

// verbose displays only if fVerbose is set
func message(str string, a ...interface{}) {
	log.Printf(str, a...)
}
