package main

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/keltia/cipher/caesar"
	"github.com/keltia/cipher/playfair"
	"github.com/keltia/cipher"
)

var (
	fDebug bool
)

func init() {
	flag.BoolVar(&fDebug, "D", false, "debug mode.")
}

func main() {
	var plain = []byte("ABCDE")
	var cipher = make([]byte, len(plain))

	flag.Parse()

	fmt.Println("test starting.")
	cipher = make([]byte, len(plain))

	var mycipher = []byte("DEFGH")

	fmt.Printf("pt=%s\n", string(plain))

	c, _ := caesar.NewCipher(3)
	c.Encrypt(cipher, plain)
	if !bytes.Equal(mycipher, cipher) {
		fmt.Printf("cipher: %s real: %s\n", string(mycipher), string(cipher))
	}
	fmt.Printf("ct=%s\n", string(cipher))

	myplain := make([]byte, len(plain))

	c.Decrypt(myplain, cipher)
	if !bytes.Equal(myplain, plain) {
		fmt.Printf("plain: %s real: %s\n", string(myplain), string(plain))
	}

	fmt.Println("-----------------")

	plain = []byte("HIDETHEGOLDINTHETREESTUMP")
	plain = crypto.Expand(plain)

	mycipher = []byte("BMODZBXDNABEKUDMUIXMMOUVIF")
	cipher = make([]byte, len(plain))

	fmt.Printf("pt=%s\n", string(plain))

	c, _ = playfair.NewCipher("PLAYFAIREXAMPLE")

	c.Encrypt(cipher, plain)
	if !bytes.Equal(mycipher, cipher) {
		fmt.Printf("cipher: %s real: %s\n", string(mycipher), string(cipher))
	}
	fmt.Printf("ct=%s\n", string(cipher))

	myplain = make([]byte, len(plain))

	c.Decrypt(myplain, cipher)
	if !bytes.Equal(myplain, plain) {
		fmt.Printf("plain: %s real: %s\n", string(myplain), string(plain))
	}
}
