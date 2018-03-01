package playfair

import (
	"crypto/cipher"
	"github.com/keltia/cipher"
	"log"
	"fmt"
)

const (
	alphabet     = "ABCDEFGHIKLMNOPQRSTUVWXYZ"
	alphabetSize = len(alphabet)
	alphabetBase = 'A'

	opEncrypt = 1
	opDecrypt = 4
)

var (
	codeWord = "01234"
)

type Cipher struct {
	key string
	i2c map[byte]couple
	c2i map[couple]byte
}

type couple struct {
	r, c byte
}

func (c *Cipher) Debug() {
	fmt.Printf("key=%s\n", c.key)
	fmt.Printf("i2c=%v\n", c.i2c)
	fmt.Printf("c2i=%v\n", c.c2i)
}

func (c *Cipher) transform(pt couple, opt byte) (ct couple) {

	bg1 := c.i2c[pt.r]
	message("line/bg1=%v", bg1)
	bg2 := c.i2c[pt.c]
	message("bg2=%v", bg2)
	if bg1.r == bg2.r {
		ct1 := couple{bg1.r, (bg1.c + opt) % 5}
		ct2 := couple{bg2.r, (bg2.c + opt) % 5}
		return couple{c.c2i[ct1],c.c2i[ct2]}
	}
	if bg1.c == bg2.c {
		ct1 := couple{(bg1.r + opt) % 5, bg1.c}
		ct2 := couple{(bg2.r + opt) % 5, bg2.c}
		message("col/ct1=%v", ct1)
		message("ct2=%v", ct2)
		return couple{c.c2i[ct1],c.c2i[ct2]}
	}
	ct1 := couple{bg1.r, bg2.c}
	ct2 := couple{bg2.r, bg1.c}
	message("sq/ct1=%v", ct1)
	message("ct2=%v", ct2)
	return couple{c.c2i[ct1], c.c2i[ct2]}
}

func expandKey(key string, i2c map[byte]couple, c2i map[couple]byte) {
	ind := 0
	for i := range codeWord {
		for j := range codeWord {
			c := key[ind]
			i2c[byte(c)] = couple{byte(i), byte(j)}
			c2i[couple{byte(i), byte(j)}] = byte(c)
			ind++
		}
	}
}

func NewCipher(key string) (cipher.Block, error) {
	c := &Cipher{
		key: crypto.Condense(key + alphabet),
		i2c: map[byte]couple{},
		c2i: map[couple]byte{},
	}
	expandKey(c.key, c.i2c, c.c2i)
	c.Debug()
	return c, nil
}

func (c *Cipher) BlockSize() int {
	return 2
}

func (c *Cipher) Encrypt(dst, src []byte) {
	if (len(src) % 2) == 1 {
		src = append(src, 'X')
	}

	dst = make([]byte, len(src))

	for i := 0; i < len(src); i += 2 {
		bg := c.transform(couple{src[i], src[i+1]}, opEncrypt)
		message("bg=%v", bg)
		dst[i] = bg.r
		dst[i+1] = bg.c
		message("dst=%s", string(dst))
	}
}

func (c *Cipher) Decrypt(dst, src []byte) {
	dst = make([]byte, len(src))

	for i := 0; i < len(src); i += 2 {
		bg := c.transform(couple{src[i], src[i+1]}, opDecrypt)
		dst[i] = bg.r
		dst[i+1] = bg.c
	}
}

// verbose displays only if fVerbose is set
func message(str string, a ...interface{}) {
	log.Printf(str, a...)
}
