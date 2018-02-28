package main

import (
	"fmt"
	"github.com/keltia/cipher/null"
	"bytes"
	"github.com/keltia/cipher/caesar"
)

func main() {
	var plain  = []byte("ABCDE")
	var cipher = make([]byte, len(plain))

	c, _ := null.NewCipher()
	c.Encrypt(cipher, plain)
	if !bytes.Equal(cipher, plain) {
		fmt.Printf("plain: %s cipher: %s\n", string(plain), string(cipher))
	}


	cipher = make([]byte, len(plain))

	var mycipher = []byte("DEFGH")

	c, _ = caesar.NewCipher(3)
	c.Encrypt(cipher, plain)
	if !bytes.Equal(mycipher, cipher) {
		fmt.Printf("cipher: %s real: %s\n", string(mycipher), string(cipher))
	}

	myplain := make([]byte, len(plain))

	c.Decrypt(myplain, cipher)
	if !bytes.Equal(myplain, plain) {
		fmt.Printf("plain: %s real: %s\n", string(myplain), string(plain))
	}
}
