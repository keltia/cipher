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
