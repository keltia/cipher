package transposition

import (
	"bytes"
	"crypto/cipher"
	"fmt"
	"github.com/keltia/cipher"
	"log"
)

type transp struct {
	key  string
	tkey []byte
}

func NewCipher(key string) (cipher.Block, error) {
	if key == "" {
		return &transp{}, fmt.Errorf("key can not be empty")
	}

	c := &transp{
		key:  key,
		tkey: crypto.ToNumeric(key),
	}
	return c, nil
}

func (c *transp) BlockSize() int {
	return len(c.tkey)
}

func (c *transp) Encrypt(dst, src []byte) {
	klen := len(c.tkey)
	table := make([]bytes.Buffer, klen)

	// Fill-in the table
	for i, ch := range src {
		table[i%klen].WriteByte(ch)
	}

	var res bytes.Buffer

	// Extract each column in order
	for i := 0; i < klen; i++ {
		j := bytes.IndexByte(c.tkey, byte(i))
		res.Write(table[j].Bytes())
	}
	copy(dst, res.Bytes())
}

func (c *transp) Decrypt(dst, src []byte) {
	klen := len(c.tkey)
	table := make([]bytes.Buffer, klen)
	scol := len(src)/klen + 1

	col := make([]byte, scol)
	pcol := make([]byte, scol-1)
	var pt = bytes.NewBuffer(src)

	// Find how many columns are not filled in (irregular table)
	pad := len(src) % klen // col 0..pad-1 are complete

	j := 0
	for i := 0; i < len(src); {

		howMany := scol
		ind := bytes.IndexByte(c.tkey, byte(j))

		// if on a non-complete column
		if ind >= pad {
			howMany--

			pt.Read(pcol)
			table[ind].Write(pcol)
		} else {
			// Complete column
			pt.Read(col)
			table[ind].Write(col)
		}

		i += howMany
		j++
	}

	// Now get all text
	for i := range src {
		pt, _ := table[i%klen].ReadByte()
		dst[i] = pt
	}
}

// verbose displays only if fVerbose is set
func message(str string, a ...interface{}) {
	log.Printf(str, a...)
}
