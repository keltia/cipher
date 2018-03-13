package nihilist

import (
	"bytes"
	"crypto/cipher"
	"github.com/keltia/cipher/straddling"
	"github.com/keltia/cipher/transposition"
	"log"
	"strings"
)

type nihilistcipher struct {
	sc     *cipher.Block
	transp *cipher.Block
}

func NewCipher(key1, key2 string, chrs string) (cipher.Block, error) {
	sub, err := straddling.NewCipher(key1, chrs)
	if err != nil {
		return nil, err
	}

	transp, err := transposition.NewCipher(key2)
	if err != nil {
		return nil, err
	}

	c := &nihilistcipher{
		sc:     &sub,
		transp: &transp,
	}
	return c, nil

}

func (c *nihilistcipher) BlockSize() int {
	return (*c.transp).BlockSize()
}

func (c *nihilistcipher) Encrypt(dst, src []byte) {
	// We need to initialize that intermediary storage ourselves
	var buf = make([]byte, 2*len(src))

	(*c.sc).Encrypt(buf, src)
	tmp := strings.TrimRight(string(buf), "\x00")
	(*c.transp).Encrypt(dst, bytes.NewBufferString(tmp).Bytes())
}

func (c *nihilistcipher) Decrypt(dst, src []byte) {
	// We need to initialize that intermediary storage ourselves
	var buf = make([]byte, len(src))

	(*c.transp).Decrypt(buf, src)
	message("buf=%s", string(buf))
	(*c.sc).Decrypt(dst, buf)
}

// verbose displays only if fVerbose is set
func message(str string, a ...interface{}) {
	log.Printf(str, a...)
}
