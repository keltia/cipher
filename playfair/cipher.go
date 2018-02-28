package playfair

import (
	"crypto/cipher"
	"github.com/keltia/cipher"
	"log"
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

type playfairCipher struct {
	key string
	i2c map[byte]couple
	c2i map[couple]byte
}

type couple struct {
	r, c byte
}

func (c *playfairCipher) transform(pt couple, opt byte) (ct couple) {

	bg1 := c.i2c[pt.r]
	bg2 := c.i2c[pt.c]
	if bg1.r == bg2.r {
		return couple{
			c.c2i[couple{bg1.r, (bg1.c + opt) % 5}],
			c.c2i[couple{bg2.r, (bg2.c + opt) % 5}]}
	}
	if bg1.c == bg2.c {
		return couple{
			c.c2i[couple{(bg1.r + opt) % 5, bg1.c}],
			c.c2i[couple{(bg2.r + opt) % 5, bg2.c}]}
	}
	return couple{c.c2i[couple{bg1.r, bg2.c}], c.c2i[couple{bg2.r, bg1.r}]}
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
	c := &playfairCipher{
		key: crypto.Condense(key + alphabet),
		i2c: map[byte]couple{},
		c2i: map[couple]byte{},
	}
	expandKey(c.key, c.i2c, c.c2i)
	return c, nil
}

func (c *playfairCipher) BlockSize() int {
	return 2
}

func (c *playfairCipher) Encrypt(dst, src []byte) {
	if (len(src) % 2) == 1 {
		src = append(src, 'X')
	}

	dst = make([]byte, len(src))

	for i := 0; i < len(src); i += 2 {
		bg := c.transform(couple{src[i], src[i+1]}, opEncrypt)
		dst[i] = bg.r
		dst[i+1] = bg.c
	}
}

func (c *playfairCipher) Decrypt(dst, src []byte) {
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
