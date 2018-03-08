package crypto

import (
	"strings"
)

// Condense3 is ported from Ruby
func Condense3(str string) string {
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
			s[ch] = true
		}
	}
	return r
}

// Condense2 is an alternate version using a map to weed out dup letters
func Condense2(str string) string {
	var s = make(map[rune]bool)
	var r strings.Builder

	for _, ch := range str {
		if _, ok := s[ch]; !ok {
			r.WriteByte(byte(ch))
			s[ch] = true
		}
	}
	return r.String()
}

// Condense is ported from Ruby
func Condense(str string) string {
	var condensed strings.Builder

	for _, ch := range str {
		if !strings.Contains(condensed.String(), string(ch)) {
			condensed.WriteByte(byte(ch))
		}
	}
	return condensed.String()
}

// insert one character inside the array
func insert(src []byte, obj byte, ind int) []byte {
	dst := make([]byte, 2*len(src))
	copy(dst, src)
	dst = append(dst[0:ind], obj)
	dst = append(dst, src[ind:]...)
	return dst
}

// ExpandBroken is a rewrite of Expand
func ExpandBroken(src []byte) []byte {
	var i int
	var dst []byte

	dst = append(dst, src[0])
	j := 0
	for i = 1; i <= len(src)-1; {
		if src[i] == dst[j] {
			dst = append(dst, byte('X'))
		} else {
			dst = append(dst, src[i])
			i++
		}
		j++
	}
	return dst
}

// Expand is a port of the Ruby version.
func Expand(src []byte) []byte {
	for i := 0; i < len(src)-1; {
		if src[i] == src[i+1] {
			src = insert(src, 'X', i+1)
		}
		i += 2
	}
	return src
}
