package crypto

import (
	"strings"
	"log"
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

// verbose displays only if fVerbose is set
func message(str string, a ...interface{}) {
	log.Printf(str, a...)
}
