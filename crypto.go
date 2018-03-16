package crypto

import (
	"bytes"
	"io"
	"sort"
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

// Condense limits allocation with a bytes.Builder
func Condense4(str string) string {
	var condensed strings.Builder

	for _, ch := range str {
		if !strings.Contains(condensed.String(), string(ch)) {
			condensed.WriteByte(byte(ch))
		}
	}
	return condensed.String()
}

// Condense4 assembles the string within a byte buffer
func Condense(str string) string {
	var condensed bytes.Buffer

	for _, ch := range str {
		if !bytes.Contains(condensed.Bytes(), []byte{byte(ch)}) {
			condensed.WriteByte(byte(ch))
		}
	}
	return condensed.String()
}

// Condense5 is like Condense but work on the string as a bytes.Buffer
func Condense5(str string) string {
	var condensed bytes.Buffer
	var bstr = bytes.NewBufferString(str).Bytes()

	for _, ch := range bstr {
		if !bytes.Contains(condensed.Bytes(), []byte{byte(ch)}) {
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

/*
  # Form an alphabet formed with a keyword, re-shuffle everything to
  # make it less predictable (i.e. checkerboard effect)
  #
  # Shuffle the alphabet a bit to avoid sequential allocation of the
  # code numbers.  This is actually performing a transposition with the word
  # itself as key.
  #
  # Regular rectangle
  # -----------------
  # Key is ARABESQUE condensed into ARBESQU (len = 7) (height = 4)
  # Let word be ARBESQUCDFGHIJKLMNOPTVWXYZ/-
  #
  # First passes will generate
  #
  # A  RBESQUCDFGHIJKLMNOPTVWXYZ/-   c=0  0 x 6
  # AC  RBESQUDFGHIJKLMNOPTVWXYZ/-   c=6  1 x 6
  # ACK  RBESQUDFGHIJLMNOPTVWXYZ/-   c=12 2 x 6
  # ACKV  RBESQUDFGHIJLMNOPTWXYZ/-   c=18 3 x 6
  # ACKVR  BESQUDFGHIJLMNOPTWXYZ/-   c=0  0 x 5
  # ACKVRD  BESQUFGHIJLMNOPTWXYZ/-   c=5  1 x 5
  # ...
  # ACKVRDLWBFMXEGNYSHOZQIP/UJT-
  #
  # Irregular rectangle
  # -------------------
  # Key is SUBWAY condensed info SUBWAY (len = 6) (height = 5)
  #
  # S  UBWAYCDEFGHIJKLMNOPQRTVXZ/-   c=0  0 x 5
  # SC  UBWAYDEFGHIJKLMNOPQRTVXZ/-   c=5  1 x 5
  # SCI  UBWAYDEFGHJKLMNOPQRTVXZ/-   c=10 2 x 5
  # SCIO  UBWAYDEFGHJKLMNPQRTVXZ/-   c=15 3 x 5
  # SCIOX  UBWAYDEFGHJKLMNPQRTVZ/-   c=20 4 x 5
  # SCIOXU  BWAYDEFGHJKLMNPQRTVZ/-   c=0  0 x 4
  # ...
  # SCIOXUDJPZBEKQ/WFLR-AG  YHMNTV   c=1  1 x 1
  # SCIOXUDJPZBEKQ/WFLR-AGM  YHNTV   c=2  2 x 1
  # SCIOXUDJPZBEKQ/WFLR-AGMT  YHNV   c=3  3 x 1
  # SCIOXUDJPZBEKQ/WFLR-AGMTYHNV
*/
// Shuffle takes a word & alphabet and mixes them around - use strings.Builder
func Shuffle(key, alphabet string) string {
	word := bytes.NewBufferString(Condense(key + alphabet)).Bytes()
	length := len(Condense(key))

	height := len(alphabet) / length
	if (len(alphabet) % length) != 0 {
		height++
	}
	res := strings.Builder{}
	for i := length - 1; i >= 0; i-- {
		for j := 0; j <= height; j++ {
			if len(word) <= height-1 {
				res.Write(word)
				return res.String()
			} else {
				if i*j < len(word) {
					c := word[i*j]
					word = append(word[0:i*j], word[i*j+1:]...)
					res.WriteByte(c)
				}
			}
		}
	}
	return res.String()
}

// Shuffle takes a word & alphabet and mixes them around - word is a string
func Shuffle1(key, alphabet string) string {
	word := Condense(key + alphabet)
	length := len(Condense(key))

	height := len(alphabet) / length
	if (len(alphabet) % length) != 0 {
		height++
	}
	res := strings.Builder{}
	for i := length - 1; i >= 0; i-- {
		for j := 0; j <= height; j++ {
			if len(word) <= height-1 {
				res.WriteString(word)
				return res.String()
			} else {
				if i*j < len(word) {
					c := word[i*j]
					word = word[0:i*j] + word[i*j+1:]
					res.WriteByte(c)
				}
			}
		}
	}
	return res.String()
}

// Shuffle takes a word & alphabet and mixes them around - full string incl. res
func Shuffle2(key, alphabet string) string {
	word := Condense(key + alphabet)
	length := len(Condense(key))

	height := len(alphabet) / length
	if (len(alphabet) % length) != 0 {
		height++
	}
	res := ""
	for i := length - 1; i >= 0; i-- {
		for j := 0; j <= height; j++ {
			if len(word) <= height-1 {
				res = res + word
				return res
			} else {
				if i*j < len(word) {
					c := word[i*j]
					word = word[0:i*j] + word[i*j+1:]
					res = res + string(c)
				}
			}
		}
	}
	return res
}

func Dup(a []byte) []byte {
	b := make([]byte, len(a))
	copy(b, a)
	return b
}

func ToNumeric(key string) []byte {
	letters := bytes.NewBufferString(key).Bytes()
	sorted := Dup(letters)
	sort.Slice(sorted, func(i, j int) bool { return sorted[i] < sorted[j] })

	f := func(c rune) rune {
		k := bytes.Index(sorted, []byte{byte(c)})
		sorted[k] = 0
		return rune(k)
	}
	ar := bytes.Map(f, letters)
	return ar
}

func ByN(ct string, n int) string {
	var err error
	var out bytes.Buffer

	in := make([]byte, n)
	blank := strings.Repeat(" ", n)
	buf := bytes.NewBufferString(ct)

	for {
		if n, err = buf.Read(in); err == io.EOF {
			break
		}

		out.WriteString(string(in))
		out.WriteByte(byte(' '))
		copy(in, blank)
	}
	return strings.TrimRight(out.String(), " ")
}

func ByN1(ct string, n int) string {
	var err error
	var out strings.Builder

	in := make([]byte, n)
	blank := strings.Repeat(" ", n)
	buf := bytes.NewBufferString(ct)

	for {
		if n, err = buf.Read(in); err == io.EOF {
			break
		}

		out.Write(in)
		out.WriteByte(byte(' '))
		copy(in, blank)
	}
	// Skip the last space
	return strings.TrimRight(out.String(), " ")
}

// Replace all instance of NN with NQN
func FixDouble(str string, fill byte) string {
	fixed := bytes.Buffer{}

	p := rune(0)
	for _, ch := range str {
		if ch == p {
			fixed.WriteByte(fill)
		} else {
			p = ch
		}
		fixed.WriteByte(byte(ch))
	}
	return fixed.String()
}

/*
// verbose displays only if fVerbose is set
func message(str string, a ...interface{}) {
	log.Printf(str, a...)
}*/
