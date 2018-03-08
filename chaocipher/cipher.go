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

/*
   # Permute the two alphabets, first ciphertext then plaintext
   # We use the current plain & ciphertext characters (akin to autoclave)
   #
   # Zenith is 0, Nadir is 13 (n/2 + 1 if 1-based)
   # Steps for left:
   # 1. shift from idx to Zenith
   # 2. take Zenith+1 out
   # 3. shift left one position and insert back the letter from step2
   #
   # Steps for right
   # 1. shift everything from plain to Zenith
   # 2. shift one more entire string
   # 3. extract Zenith+2
   # 4. shift from Zenith+3 to Nadir left
   # 5. insert  letter from step 3 in place
   #
   def advance(idx)
     if idx != 0 then
       cw = @cipher[idx..-1] + @cipher[ZENITH..(idx - 1)]
       pw = @plain[idx..-1] + @plain[ZENITH..(idx - 1)]
     else
       cw = @cipher
       pw = @plain
     end
     @cipher = cw[ZENITH].chr + cw[(ZENITH + 2)..NADIR] + \
             cw[ZENITH + 1].chr + cw[(NADIR + 1)..-1]
     raise DataError, 'cw length bad' if cw.length != BASE.length

     pw = pw[(ZENITH + 1)..-1] + pw[ZENITH].chr
     @plain = pw[ZENITH..(ZENITH + 1)] + pw[(ZENITH + 3)..NADIR] + \
            pw[ZENITH + 2].chr + pw[(NADIR + 1)..-1]
     raise DataError, 'pw length bad' if pw.length != BASE.length
   end # -- advance
*/
func (c *chaocipher) advance(idx int) {
	var pw, cw []byte

	if idx != 0 {
		pw = append(c.pw[idx:], c.pw[zenith:idx]...)
		cw = append(c.cw[idx:], c.cw[zenith:idx]...)
	} else {
		pw = c.pw
		cw = c.cw
	}

	// shift ciphertext wheel
	tcw := []byte{cw[zenith]}
	tcw = append(tcw, cw[zenith+2:nadir]...)
	tcw = append(tcw, cw[zenith+1])
	tcw = append(tcw, cw[nadir+1:]...)
	c.cw = tcw

	// shift plaintext wheel
	tpw := cw[zenith+1:]
	tpw = append(tpw, pw[zenith])
	tpw = append(tpw, pw[zenith:zenith+1]...)
	tpw = append(tpw, pw[zenith+3:nadir]...)
	tpw = append(tpw, pw[zenith+2])
	tpw = append(tpw, pw[nadir+1:]...)
	c.pw = tpw
}

func (c *chaocipher) encodeBoth(r1, r2 []byte, ch byte) byte {
	idx := bytes.Index(r1, []byte{ch})
	pt := r2[idx]
	c.advance(idx)
	return pt
}

func (c *chaocipher) encode(ch byte) byte {
	return c.encodeBoth(c.pw, c.cw, ch)
}

func (c *chaocipher) decode(ch byte) byte {
	return c.encodeBoth(c.cw, c.pw, ch)
}

func (c *chaocipher) Encrypt(dst, src []byte) {

}

func (c *chaocipher) Decrypt(dst, src []byte) {

}
