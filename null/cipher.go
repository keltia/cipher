package null

import (
	"crypto/cipher"
	"log"
)

type nullCipher struct {
}

func NewCipher() (cipher.Block, error) {
	c := &nullCipher{}
	return c, nil
}

func (c *nullCipher) BlockSize() int {
	return 1
}

func (c *nullCipher) Encrypt(dst, src []byte) {
	for i, ch := range src {
		dst[i] = ch
	}
}

func (c *nullCipher) Decrypt(dst, src []byte) {
	for i, ch := range src {
		dst[i] = ch
	}
}

// verbose displays only if fVerbose is set
func message(str string, a ...interface{}) {
	log.Printf(str, a...)
}
