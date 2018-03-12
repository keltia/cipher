package crypto

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var testCondensedData = []struct{ a, b string }{
	{"ABCDE", "ABCDE"},
	{"AAAAA", "A"},
	{"ARABESQUE", "ARBESQU"},
	{"ARABESQUEABCDEFGHIKLMNOPQRSTUVWXYZ", "ARBESQUCDFGHIKLMNOPTVWXYZ"},
	{"PLAYFAIRABCDEFGHIKLMNOPQRSTUVWXYZ", "PLAYFIRBCDEGHKMNOQSTUVWXZ"},
	{"PLAYFAIREXMABCDEFGHIKLMNOPQRSTUVWXYZ", "PLAYFIREXMBCDGHKNOQSTUVWZ"},
}

func TestCondense(t *testing.T) {
	for _, td := range testCondensedData {
		assert.Equal(t, td.b, Condense(td.a))
	}
}

func TestCondense1(t *testing.T) {
	for _, td := range testCondensedData {
		assert.Equal(t, td.b, Condense1(td.a))
	}
}

func TestCondense2(t *testing.T) {
	for _, td := range testCondensedData {
		assert.Equal(t, td.b, Condense2(td.a))
	}
}

func TestCondense3(t *testing.T) {
	for _, td := range testCondensedData {
		assert.Equal(t, td.b, Condense3(td.a))
	}
}

func TestCondense4(t *testing.T) {
	for _, td := range testCondensedData {
		assert.Equal(t, td.b, Condense4(td.a))
	}
}

func TestCondense5(t *testing.T) {
	for _, td := range testCondensedData {
		assert.Equal(t, td.b, Condense5(td.a))
	}
}

var bar string

func BenchmarkCondense(b *testing.B) {
	var r string

	for _, td := range testCondensedData {
		for n := 0; n < b.N; n++ {
			r = Condense(td.a)
		}
	}
	bar = r
}

func BenchmarkCondense1(b *testing.B) {
	var r string

	for _, td := range testCondensedData {
		for n := 0; n < b.N; n++ {
			r = Condense1(td.a)
		}
	}
	bar = r
}

func BenchmarkCondense2(b *testing.B) {
	var r string

	for _, td := range testCondensedData {
		for n := 0; n < b.N; n++ {
			r = Condense2(td.a)
		}
	}
	bar = r
}

func BenchmarkCondense3(b *testing.B) {
	var r string

	for _, td := range testCondensedData {
		for n := 0; n < b.N; n++ {
			r = Condense3(td.a)
		}
	}
	bar = r
}

func BenchmarkCondense4(b *testing.B) {
	var r string

	for _, td := range testCondensedData {
		for n := 0; n < b.N; n++ {
			r = Condense4(td.a)
		}
	}
	bar = r
}

func BenchmarkCondense5(b *testing.B) {
	var r string

	for _, td := range testCondensedData {
		for n := 0; n < b.N; n++ {
			r = Condense5(td.a)
		}
	}
	bar = r
}

var testExpandData = []struct{ a, b string }{
	{"AAA", "AXAXA"},
	{"AAAA", "AXAXAXA"},
	{"AAABRAACADAABRA", "AXAXABRAXACADAXABRA"},
	{"ARABESQUE", "ARABESQUE"},
	{"LANNONCE", "LANXNONCE"},
	{"PJRJJJJJJS", "PJRJXJXJXJXJXJS"},
	{"ABCDEFGHJJKLM", "ABCDEFGHJXJKLM"},
}

var testExpandInsertData = []struct{ a, b string }{
	{"AAA", "AXAXA"},
	{"AAAA", "AXAXAXA"},
	{"AAABRAACADAABRA", "AXAXABRAACADAXABRA"},
	{"ARABESQUE", "ARABESQUE"},
	{"LANNONCE", "LANXNONCE"},
	{"PJRJJJJJJS", "PJRJJXJXJXJXJS"},
	{"ABCDEFGHJJKLM", "ABCDEFGHJXJKLM"},
}

func TestExpand(t *testing.T) {
	for _, td := range testExpandData {
		assert.EqualValues(t, []byte(td.b), ExpandBroken([]byte(td.a)))
	}
}

func TestExpandInsert(t *testing.T) {
	for _, td := range testExpandInsertData {
		assert.EqualValues(t, []byte(td.b), Expand([]byte(td.a)))
	}
}

func TestInsert(t *testing.T) {
	a := []byte{0, 1, 2, 3}
	b := []byte{0, 1, 42, 2, 3}
	assert.EqualValues(t, b, insert(a, 42, 2))
}

func TestShuffle(t *testing.T) {
	key := "ARABESQUE"
	alphabet := "ABCDEFGHIJKLMNOPQRSTUVWXYZ/-"

	res := Shuffle(key, alphabet)
	assert.Equal(t, "ACKVRDLWBFMXEGNYSHOZQIP/UJT-", res)
}

func TestShuffleOdd(t *testing.T) {
	key := "SUBWAY"
	alphabet := "ABCDEFGHIJKLMNOPQRSTUVWXYZ/-"

	res := Shuffle(key, alphabet)
	assert.Equal(t, "SCIOXUDJPZBEKQ/WFLR-AGMTYHNV", res)
}

func TestShuffleOdd1(t *testing.T) {
	key := "SUBWAY"
	alphabet := "ABCDEFGHIJKLMNOPQRSTUVWXYZ/-"

	res := Shuffle1(key, alphabet)
	assert.Equal(t, "SCIOXUDJPZBEKQ/WFLR-AGMTYHNV", res)
}

func TestShuffleOdd2(t *testing.T) {
	key := "SUBWAY"
	alphabet := "ABCDEFGHIJKLMNOPQRSTUVWXYZ/-"

	res := Shuffle2(key, alphabet)
	assert.Equal(t, "SCIOXUDJPZBEKQ/WFLR-AGMTYHNV", res)
}

var NumericData = []struct {
	str string
	key []byte
}{
	{"ARABESQUE", []byte{0, 6, 1, 2, 3, 7, 5, 8, 4}},
	{"PJRJJJJJJS", []byte{7, 0, 8, 1, 2, 3, 4, 5, 6, 9}},
	{"AAABRAACADAABRA", []byte{0, 1, 2, 9, 13, 3, 4, 11, 5, 12, 6, 7, 10, 14, 8}},
}

func TestToNumeric(t *testing.T) {
	for _, data := range NumericData {
		str := data.str
		key := ToNumeric(str)
		assert.EqualValues(t, data.key, key)
	}
}

var ByNData = []struct {
	n   int
	in  string
	out string
}{
	{5, "ARABESQUE", "ARABE SQUE"},
	{4, "PJRJJJJJJS", "PJRJ JJJJ JS"},
	{5, "AAABRAACADAABRA", "AAABR AACAD AABRA"},
}

func TestByN(t *testing.T) {
	for _, cp := range ByNData {
		assert.Equal(t, cp.out, ByN(cp.in, cp.n))
	}
}

func TestByN1(t *testing.T) {
	for _, cp := range ByNData {
		assert.Equal(t, cp.out, ByN1(cp.in, cp.n))
	}
}

// -- benchmarks

func BenchmarkShuffle(b *testing.B) {
	var res string

	key := "SUBWAY"
	alphabet := "ABCDEFGHIJKLMNOPQRSTUVWXYZ/-"

	for n := 0; n < b.N; n++ {
		res = Shuffle(key, alphabet)
	}
	bar = res
}

func BenchmarkShuffle1(b *testing.B) {
	var res string

	key := "SUBWAY"
	alphabet := "ABCDEFGHIJKLMNOPQRSTUVWXYZ/-"

	for n := 0; n < b.N; n++ {
		res = Shuffle1(key, alphabet)
	}
	bar = res
}

func BenchmarkShuffle2(b *testing.B) {
	var res string

	key := "SUBWAY"
	alphabet := "ABCDEFGHIJKLMNOPQRSTUVWXYZ/-"

	for n := 0; n < b.N; n++ {
		res = Shuffle2(key, alphabet)
	}
	bar = res
}

var foo []byte

func BenchmarkExpand(b *testing.B) {
	var r []byte

	for _, td := range testExpandData {
		for n := 0; n < b.N; n++ {
			r = ExpandBroken([]byte(td.a))
		}
	}
	foo = r
}

func BenchmarkExpandInsert(b *testing.B) {
	var r []byte

	for _, td := range testExpandData {
		for n := 0; n < b.N; n++ {
			r = Expand([]byte(td.a))
		}
	}
	foo = r
}

var gres []byte

func BenchmarkToNumeric(b *testing.B) {
	var res []byte

	str := "ANTICONSTITUTIONNELLEMENT"

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		res = ToNumeric(str)
	}
	gres = res
}

var gs string

func BenchmarkByN(b *testing.B) {
	var s string

	str := "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456"
	nb := 5
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		s = ByN(str, nb)
	}
	gs = s
}

func BenchmarkByN1(b *testing.B) {
	var s string

	str := "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456"
	nb := 5
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		s = ByN1(str, nb)
	}
	gs = s
}
