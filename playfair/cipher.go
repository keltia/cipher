package playfair

import (
	"crypto/cipher"
	"github.com/keltia/cipher"
)

const (
	alphabet = "ABCDEFGHIKLMNOPQRSTUVWXYZ"

	opEncrypt = 1
	opDecrypt = 4
)

var (
	codeWord = "01234"
)

// Cipher holds the key and transformation maps
type Cipher struct {
	key string
	i2c map[byte]couple
	c2i map[couple]byte
}

type couple struct {
	r, c byte
}

// transform is the cipher itself
func (c *Cipher) transform(pt couple, opt byte) (ct couple) {

	bg1 := c.i2c[pt.r]
	bg2 := c.i2c[pt.c]
	if bg1.r == bg2.r {
		ct1 := couple{bg1.r, (bg1.c + opt) % 5}
		ct2 := couple{bg2.r, (bg2.c + opt) % 5}
		return couple{c.c2i[ct1], c.c2i[ct2]}
	}
	if bg1.c == bg2.c {
		ct1 := couple{(bg1.r + opt) % 5, bg1.c}
		ct2 := couple{(bg2.r + opt) % 5, bg2.c}
		return couple{c.c2i[ct1], c.c2i[ct2]}
	}
	ct1 := couple{bg1.r, bg2.c}
	ct2 := couple{bg2.r, bg1.c}
	return couple{c.c2i[ct1], c.c2i[ct2]}
}

// expandKey create the two transformation maps
func expandKey(key string, i2c map[byte]couple, c2i map[couple]byte) {
	ind := 0
	for i := range codeWord {
		for j := range codeWord {
			c := key[ind]
			i2c[c] = couple{byte(i), byte(j)}
			c2i[couple{byte(i), byte(j)}] = c
			ind++
		}
	}
}

// NewCipher is part of the interface
func NewCipher(key string) (cipher.Block, error) {
	c := &Cipher{
		key: crypto.Condense(key + alphabet),
		i2c: map[byte]couple{},
		c2i: map[couple]byte{},
	}
	expandKey(c.key, c.i2c, c.c2i)
	return c, nil
}

// BlockSize is part of the interface
func (c *Cipher) BlockSize() int {
	return 2
}

// Encrypt is part of the interface
func (c *Cipher) Encrypt(dst, src []byte) {
	if (len(src) % 2) == 1 {
		src = append(src, 'X')
	}

	for i := 0; i < len(src); i += 2 {
		bg := c.transform(couple{src[i], src[i+1]}, opEncrypt)
		dst[i] = bg.r
		dst[i+1] = bg.c
	}
}

// Decrypt is part of the interface
func (c *Cipher) Decrypt(dst, src []byte) {
	if (len(src) % 2) == 1 {
		panic("odd number of elements")
	}

	for i := 0; i < len(src); i += 2 {
		bg := c.transform(couple{src[i], src[i+1]}, opDecrypt)
		dst[i] = bg.r
		dst[i+1] = bg.c
	}
}

/*
// verbose displays only if fVerbose is set
func message(str string, a ...interface{}) {
	log.Printf(str, a...)
}*/
