package null

import (
	"crypto/cipher"
	"log"
)

type nullCipher struct {
}

// NewCipher creates a new instance of cipher.Block
func NewCipher() (cipher.Block, error) {
	c := &nullCipher{}
	return c, nil
}

// BlockSize is part of the interface
func (c *nullCipher) BlockSize() int {
	return 1
}

// Encrypt is part of the interface
func (c *nullCipher) Encrypt(dst, src []byte) {
	for i, ch := range src {
		dst[i] = ch
	}
}

// Decrypt is part of the interface
func (c *nullCipher) Decrypt(dst, src []byte) {
	copy(dst, src)
	for i, ch := range src {
		dst[i] = ch
	}
}

// verbose displays only if fVerbose is set
func message(str string, a ...interface{}) {
	log.Printf(str, a...)
}
