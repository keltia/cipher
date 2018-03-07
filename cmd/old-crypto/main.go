package main

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/keltia/cipher"
	"github.com/keltia/cipher/caesar"
	"github.com/keltia/cipher/playfair"
)

var (
	fDebug bool
)

func init() {
	flag.BoolVar(&fDebug, "D", false, "debug mode.")
}

func main() {
	var plain = []byte("ABCDE")

	var ncipher []byte

	flag.Parse()

	fmt.Println("test starting.")
	ncipher = make([]byte, len(plain))

	var mycipher = []byte("DEFGH")

	fmt.Printf("pt=%s\n", string(plain))

	c, _ := caesar.NewCipher(3)
	c.Encrypt(ncipher, plain)
	if !bytes.Equal(mycipher, ncipher) {
		fmt.Printf("ncipher: %s real: %s\n", string(mycipher), string(ncipher))
	}
	fmt.Printf("ct=%s\n", string(ncipher))

	myplain := make([]byte, len(plain))

	c.Decrypt(myplain, ncipher)
	if !bytes.Equal(myplain, plain) {
		fmt.Printf("plain: %s real: %s\n", string(myplain), string(plain))
	}

	fmt.Println("-----------------")

	plain = []byte("HIDETHEGOLDINTHETREESTUMP")
	plain = crypto.Expand(plain)

	mycipher = []byte("BMODZBXDNABEKUDMUIXMMOUVIF")
	ncipher = make([]byte, len(plain))

	fmt.Printf("pt=%s\n", string(plain))

	c, _ = playfair.NewCipher("PLAYFAIREXAMPLE")

	c.Encrypt(ncipher, plain)
	if !bytes.Equal(mycipher, ncipher) {
		fmt.Printf("ncipher: %s real: %s\n", string(mycipher), string(ncipher))
	}
	fmt.Printf("ct=%s\n", string(ncipher))

	myplain = make([]byte, len(plain))

	c.Decrypt(myplain, ncipher)
	if !bytes.Equal(myplain, plain) {
		fmt.Printf("plain: %s real: %s\n", string(myplain), string(plain))
	}
}
