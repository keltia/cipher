package caesar

import (
	"crypto/cipher"
	"log"
)

const (
	alphabet     = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	alphabetSize = len(alphabet)
)

type caesarCipher struct {
	key byte
	enc map[byte]byte
	dec map[byte]byte
}

func encrypt(pt byte, in map[byte]byte) byte {
	return in[pt]
}

func decrypt(ct byte, out map[byte]byte) byte {
	return out[ct]
}

func expandKey(key byte, in, out map[byte]byte) {
	for i, ch := range alphabet {
		transform := byte((i + int(key)) % alphabetSize)
		in[byte(ch)] = alphabet[transform]
		out[alphabet[transform]] = byte(ch)
	}
}

// NewCipher creates a new instance of cipher.Block
func NewCipher(key int) (cipher.Block, error) {
	c := &caesarCipher{
		key: byte(key),
		enc: map[byte]byte{},
		dec: map[byte]byte{},
	}
	expandKey(c.key, c.enc, c.dec)
	return c, nil
}

// BlockSize is part of the interface
func (c *caesarCipher) BlockSize() int {
	return 1
}

// Encrypt is part of the interface
func (c *caesarCipher) Encrypt(dst, src []byte) {
	for i, ch := range src {
		dst[i] = encrypt(ch, c.enc)
	}
}

// Decrypt is part of the interface
func (c *caesarCipher) Decrypt(dst, src []byte) {
	for i, ch := range src {
		dst[i] = decrypt(ch, c.dec)
	}
}

// verbose displays only if fVerbose is set
func message(str string, a ...interface{}) {
	log.Printf(str, a...)
}
