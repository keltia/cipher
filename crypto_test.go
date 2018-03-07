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

var bar string

func BenchmarkCondense(b *testing.B) {
	var r string

	for n := 0; n < b.N; n++ {
		for _, td := range testCondensedData {
			r = Condense(td.a)
		}
	}
	bar = r
}

func BenchmarkCondense1(b *testing.B) {
	var r string

	for n := 0; n < b.N; n++ {
		for _, td := range testCondensedData {
			r = Condense1(td.a)
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
		assert.EqualValues(t, []byte(td.b), Expand([]byte(td.a)))
	}
}

func TestExpandInsert(t *testing.T) {
	for _, td := range testExpandInsertData {
		assert.EqualValues(t, []byte(td.b), ExpandInsert([]byte(td.a)))
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

	for n := 0; n < b.N; n++ {
		for _, td := range testExpandData {
			r = Expand([]byte(td.a))
		}
	}
	foo = r
}

func BenchmarkExpandInsert(b *testing.B) {
	var r []byte

	for n := 0; n < b.N; n++ {
		for _, td := range testExpandData {
			r = ExpandInsert([]byte(td.a))
		}
	}
	foo = r
}