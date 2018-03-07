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

// Condense is ported from Ruby
func Condense(str string) string {
	var condensed []byte

	for _, ch := range str {
		if !strings.Contains(string(condensed), string(ch)) {
			condensed = append(condensed, byte(ch))
		}
	}
	return string(condensed)
}

// Condense1 is an alternate version using a map to weed out dup letters
func Condense1(str string) string {
	var s = make(map[rune]bool)
	var r string

	for _, ch := range str {
		if _, ok := s[ch]; !ok {
			r = r + string(ch)
		}
	}
	return r
}

func insert(src []byte, obj byte, ind int) []byte {
	dst := make([]byte, 2*len(src))
	copy(dst, src)
	dst = append(dst[0:ind], obj)
	dst = append(dst, src[ind:]...)
	return dst
}

func ExpandBroken(src []byte) []byte {
	var i int
	var dst []byte

	dst = append(dst, src[0])
	j := 0
	for i = 1; i <= len(src)-1; {
		if src[i] == dst[j] {
			dst = append(dst, byte('X'))
			//message("<i=%d j=%d dst=%s", i, j, dst)
		} else {
			dst = append(dst, src[i])
			i += 1
			//message(">i=%d j=%d dst=%s", i, j, dst)
		}
		j += 1
	}
	//message("---")
	return dst
}

func Expand(src []byte) []byte {
	//dst = append(dst, src[len(src) - 1])
	for i := 0; i < len(src)-1; {
		if src[i] == src[i+1] {
			src = insert(src, 'X', i+1)
			//message("src=%v", src)
		}
		//message("i=%d src=%s", i, src)
		i += 2
	}
	//message("--->")
	return src
}

// verbose displays only if fVerbose is set
func message(str string, a ...interface{}) {
	log.Printf(str, a...)
}
