package chaocipher

import (
	"bytes"
	"crypto/cipher"
	"fmt"
)

const (
	alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	zenith   = 0
	nadir    = 13 // 26/2 (+1 if one-based)
)

type chaocipher struct {
	pw, cw []byte
}

// NewCipher creates a new cipher with the provided keys
func NewCipher(pkey, ckey string) (cipher.Block, error) {
	if len(pkey) != len(alphabet) ||
		len(ckey) != len(alphabet) {
		return &chaocipher{}, fmt.Errorf("bad alphabet length")
	}

	c := &chaocipher{
		pw: bytes.NewBufferString(pkey).Bytes(),
		cw: bytes.NewBufferString(ckey).Bytes(),
	}
	return c, nil
}

func (c *chaocipher) BlockSize() int {
	return 1
}

func lshift(a []byte) {
	f := a[0]
	copy(a, a[1:])
	a[len(a)-1] = f
}

func rshift(a []byte) {
	f := a[len(a)-1]
	copy(a[1:], a[0:len(a)-1])
	a[0] = f
}

func lshiftN(a []byte, n int) {
	f := dup(a[0:n])
	copy(a, a[n:])
	copy(a[len(a)-n:], f)
}

func rshiftN(a []byte, n int) {
	f := dup(a[len(a)-n:])
	copy(a[n:], a[0:len(a)-n])
	copy(a, f)
}

func dup(a []byte) []byte {
	b := make([]byte, len(a))
	copy(b, a)
	return b
}

/*
Permute the two alphabets, first ciphertext then plaintext
We use the current plain & ciphertext characters (akin to autoclave)

Zenith is 0, Nadir is 13 (n/2 + 1 if 1-based)

Steps for left:
 1. shift from idx to Zenith
 2. take Zenith+1 out
 3. shift left one position and insert back the letter from step2

Steps for right
 1. shift everything from plain to Zenith
 2. shift one more entire string
 3. extract Zenith+2
 4. shift from Zenith+3 to Nadir left
 5. insert  letter from step 3 in place

XXX due to the way Go manages the array memory & copy() works with slices, there are
index & length differences with Ruby
*/
func (c *chaocipher) advance(idx int) {
	// First we shift the left alphabet (cw)
	lshiftN(c.cw, idx)
	l := c.cw[zenith+1]
	copy(c.cw[zenith+1:nadir], c.cw[zenith+2:nadir+1])
	c.cw[nadir] = l

	// Then we shift the right alphabet (pw)
	lshiftN(c.pw, idx+1)
	l = c.pw[zenith+2]
	copy(c.pw[zenith+2:nadir], c.pw[zenith+3:nadir+1])
	c.pw[nadir] = l
}

func (c *chaocipher) encodeBoth(r1, r2 []byte, ch byte) byte {
	idx := bytes.Index(r1, []byte{ch})
	pt := r2[idx]
	c.advance(idx)
	return pt
}

func (c *chaocipher) encode(ch byte) byte {
	return c.encodeBoth(c.pw, c.cw, ch)
}

func (c *chaocipher) decode(ch byte) byte {
	return c.encodeBoth(c.cw, c.pw, ch)
}

func (c *chaocipher) Encrypt(dst, src []byte) {
	for i, ch := range src {
		dst[i] = c.encode(ch)
	}
}

func (c *chaocipher) Decrypt(dst, src []byte) {
	for i, ch := range src {
		dst[i] = c.decode(ch)
	}
}
