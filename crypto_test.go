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

var testExpandData = []struct{ a, b string }{
	{"AAA", "AXAXA"},
	{"AAAA", "AXAXAXA"},
	{"AAABRAACADAABRA", "AXAXABRAXACADAXABRA"},
	{"ARABESQUE", "ARABESQUE"},
	{"LANNONCE", "LANXNONCE"},
	{"PJRJJJJJJS", "PJRJJXJXJXJXJS"},
}

func TestExpand(t *testing.T) {
	for _, td := range testExpandData {
		assert.EqualValues(t, []byte(td.b), Expand([]byte(td.a)))
	}
}
