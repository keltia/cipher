package chaocipher

import (
	"bytes"
	"crypto/cipher"
	"fmt"
)

const (
	alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	zenith   = 0
	nadir    = 13
)

type chaocipher struct {
	pw, cw []byte
}

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

func advance(idx int) {

}

func encodeBoth(r1, r2 []byte, ch byte) byte {
	idx := bytes.Index(r1, []byte{ch})
	pt := r2[idx]
	advance(idx)
	return pt
}

func (c *chaocipher) encode(ch byte) byte {
	return encodeBoth(c.pw, c.cw, ch)
}

func (c *chaocipher) decode(ch byte) byte {
	return encodeBoth(c.cw, c.pw, ch)
}

func (c *chaocipher) Encrypt(dst, src []byte) {

}

func (c *chaocipher) Decrypt(dst, src []byte) {

}
