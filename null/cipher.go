package null

import (
	"crypto/cipher"
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
	copy(dst, src)
}

// Decrypt is part of the interface
func (c *nullCipher) Decrypt(dst, src []byte) {
	copy(dst, src)
}
