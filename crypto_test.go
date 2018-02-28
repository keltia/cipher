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
}

func TestCondense(t *testing.T) {
	for _, td := range testCondensedData {
		assert.Equal(t, td.b, Condense(td.a))
	}
}
