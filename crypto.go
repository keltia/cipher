package crypto

import (
	"log"
	"strings"
)

type Bigram struct {
	r, c byte
}

type Key struct {
	value string
}

func (k *Key) String() string {
	return k.value
}

func Condense(str string) string {
	condensed := ""
	for _, ch := range str {
		if !strings.Contains(condensed, string(ch)) {
			condensed = condensed + string(ch)
		}
	}
	return condensed
}

func Expand(src []byte) []byte {
	var i int
	var dst []byte

	dst = append(dst, src[0])
	j := 0
	for i = 1; i <= len(src) - 1; {
		if src[i] == dst[j] {
			dst = append(dst, byte('X'))
			j += 1
			message("i=%d j=%d dst=%s", i, j, dst)
		} else {
			dst = append(dst, src[i])
			i += 1
			j += 1
			message("i=%d j=%d dst=%s", i, j, dst)
		}
	}
	message("---")
	return dst
}

// verbose displays only if fVerbose is set
func message(str string, a ...interface{}) {
	log.Printf(str, a...)
}
