package adfgvx

import (
	"crypto/cipher"
	"github.com/keltia/cipher/square"
	"github.com/keltia/cipher/transposition"
)

type adfgvxcipher struct {
	sqr    *cipher.Block
	transp *cipher.Block
}

func NewCipher(key1, key2 string) (cipher.Block, error) {
	sub, err := square.NewCipher(key1, "ADFGVX")
	if err != nil {
		return nil, err
	}

	transp, err := transposition.NewCipher(key2)
	if err != nil {
		return nil, err
	}

	c := &adfgvxcipher{
		sqr:    &sub,
		transp: &transp,
	}
	return c, nil
}

func (c *adfgvxcipher) BlockSize() int {
	return (*c.transp).BlockSize()
}

func (c *adfgvxcipher) Encrypt(dst, src []byte) {
	// We need to initialize that intermediary storage ourselves
	var buf = make([]byte, 2*len(src))

	(*c.sqr).Encrypt(buf, src)
	(*c.transp).Encrypt(dst, buf)
}

func (c *adfgvxcipher) Decrypt(dst, src []byte) {
	// We need to initialize that intermediary storage ourselves
	var buf = make([]byte, len(src))

	(*c.transp).Decrypt(buf, src)
	(*c.sqr).Decrypt(dst, buf)
}

/*
// verbose displays only if fVerbose is set
func message(str string, a ...interface{}) {
	log.Printf(str, a...)
}
*/
