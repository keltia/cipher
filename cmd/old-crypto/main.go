package main

import (
	"crypto/cipher"
	"flag"
	"fmt"
	"github.com/keltia/cipher"
	"github.com/keltia/cipher/adfgvx"
	"github.com/keltia/cipher/caesar"
	"github.com/keltia/cipher/chaocipher"
	"github.com/keltia/cipher/nihilist"
	"github.com/keltia/cipher/playfair"
	"github.com/keltia/cipher/square"
	"github.com/keltia/cipher/straddling"
	"github.com/keltia/cipher/transposition"
)

var (
	fDebug bool
)

func init() {
	flag.BoolVar(&fDebug, "D", false, "debug mode.")
}

const (
	keyPlain  = "PTLNBQDEOYSFAVZKGJRIHWXUMC"
	keyCipher = "HXUCZVAMDSLKPEFJRIGTWOBNYQ"

	plain = "IFYOUCANREADTHISYOUEITHERDOWNLOADEDMYOWNIMPLEMENTATIONOFCHAOCIPHERORYOUWROTEONEOFYOUROWNINEITHERCASELETMEKNOWANDACCEPTMYCONGRATULATIONSX"
)

var allciphers []CPH

type CPH struct {
	name string
	c    cipher.Block
	size int
}

func init() {
	var c cipher.Block

	c, _ = caesar.NewCipher(3)
	allciphers = append(allciphers, CPH{"Caesar", c, len(plain)})

	c, _ = square.NewCipher("ARABESQUE", "012345")
	allciphers = append(allciphers, CPH{"Square", c, len(plain) * 2})

	c, _ = transposition.NewCipher("SUBWAY")
	allciphers = append(allciphers, CPH{"Transp", c, len(plain)})

	c, _ = chaocipher.NewCipher(keyPlain, keyCipher)
	allciphers = append(allciphers, CPH{"Chaocipher", c, len(plain)})

	c, _ = playfair.NewCipher("ARABESQUE")
	allciphers = append(allciphers, CPH{"Playfair", c, len(plain)})

	c, _ = adfgvx.NewCipher("ARABESQUE", "SUBWAY")
	allciphers = append(allciphers, CPH{"ADFGVX", c, len(plain) * 2})

	c, _ = straddling.NewCipher("ARABESQUE", "37")
	allciphers = append(allciphers, CPH{"Straddling", c, len(plain) * 2})

	c, _ = nihilist.NewCipher("ARABESQUE", "SUBWAY", "37")
	allciphers = append(allciphers, CPH{"Nihilist", c, len(plain) * 2})
}

func main() {

	fmt.Printf("==> Plain = \n%s\n", plain)
	for _, cp := range allciphers {
		dst := make([]byte, cp.size)

		c := cp.c
		c.Encrypt(dst, []byte(plain))
		fmt.Println("==> ", cp.name)
		fmt.Printf("%s\n", crypto.ByN(string(dst), 5))
	}
}
