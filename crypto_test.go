package crypto

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

var testCondensedData = []struct{a, b string}{
	{"ABCDE", "ABCDE"},
	{"AAAAA", "A"},
	{"ARABESQUE", "ARBESQU"},
}

func TestCondense(t *testing.T) {
	for _, td := range testCondensedData {
		assert.Equal(t, td.b, Condense(td.a))
	}
}
