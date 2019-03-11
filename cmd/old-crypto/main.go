package main

import (
	"bytes"
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
	"github.com/keltia/cipher/wheatstone"
	"strings"
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

	//plain
	// = "IFYOUCANREADTHISYOUEITHERDOWNLOADEDMYOWNIMPLEMENTATIONOFCHAOCIPHERORYOUWROTEONEOFYOUROWNINEITHERCASELETMEKNOWANDACCEPTMYCONGRATULATIONSX"
	plain = "CETOOTESTCHIFFREAVECADFGVXETLESCLESMASTODONETSOCIALX"
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

	c, _ = wheatstone.NewCipher('M', "CIPHER", "MACHINE")
	allciphers = append(allciphers, CPH{"Wheatstone", c, len(plain)})

	c, _ = adfgvx.NewCipher("MASTODON", "SOCIAL")
	allciphers = append(allciphers, CPH{"ADFGVX2", c, len(plain) * 2})

}

func main() {
	var fixpt string

	fmt.Printf("==> Plain = \n%s\n", plain)
	for _, cp := range allciphers {
		dst := make([]byte, cp.size)
		dst1 := make([]byte, len(plain))

		c := cp.c

		if cp.name == "Wheatstone" {
			fixpt = crypto.FixDouble(plain, 'Q')
			dst = make([]byte, len(fixpt))
			dst1 = make([]byte, len(fixpt))
		} else {
			fixpt = plain

		}
		c.Encrypt(dst, []byte(fixpt))
		fmt.Println("==> ", cp.name)
		fmt.Printf("%s\n", crypto.ByN(string(dst), 5))

		ncrypt := strings.TrimRight(string(dst), "\x00")
		src1 := bytes.NewBufferString(ncrypt).Bytes()

		c.Decrypt(dst1, src1)

		nplain := strings.TrimRight(string(dst1), "\x00")
		if nplain == fixpt {
			fmt.Printf("decrypt ok\n\n")
		} else {
			fmt.Printf("decrypt not ok\n%s\n%s\n\n", fixpt, nplain)
		}
	}
}
