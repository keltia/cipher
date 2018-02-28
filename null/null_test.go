package null

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

type NullTest struct {
    pt, ct string
}

var encryptCaesarTests = []NullTest{
    { "ABCDE", "ABCDE"},
    { "COUCOU", "COUCOU"},
}

func TestNewCipher(t *testing.T) {
    for _, pair := range encryptCaesarTests {
        c, _ := NewCipher()

        assert.EqualValues(t, 1, c.BlockSize())
        plain := []byte(pair.pt)
        cipher := make([]byte, len(plain))
        c.Encrypt(cipher, plain)
        assert.Equal(t, []byte(pair.ct), cipher)

        nplain := make([]byte, len(plain))
        c.Decrypt(nplain, cipher)
        assert.Equal(t, []byte(pair.pt), nplain)
    }
}
