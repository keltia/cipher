package caesar

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type CaesarTest struct {
	key    int
	pt, ct string
}

var encryptCaesarTests = []CaesarTest{
	{3, "ABCDE", "DEFGH"},
	{4, "COUCOU", "GSYGSY"},
	{13, "COUCOU", "PBHPBH"},
}

func TestExpandKey(t *testing.T) {
	enc := map[byte]byte{}
	dec := map[byte]byte{}

	myenc := map[byte]byte{
		'A': 'D', 'B': 'E', 'C': 'F', 'D': 'G', 'E': 'H', 'F': 'I',
		'G': 'J', 'H': 'K', 'I': 'L', 'J': 'M', 'K': 'N', 'L': 'O',
		'M': 'P', 'N': 'Q', 'O': 'R', 'P': 'S', 'Q': 'T', 'R': 'U',
		'S': 'V', 'T': 'W', 'U': 'X', 'V': 'Y', 'W': 'Z', 'X': 'A',
		'Y': 'B', 'Z': 'C',
	}

	mydec := map[byte]byte{
		'D': 'A', 'E': 'B', 'F': 'C', 'G': 'D', 'H': 'E', 'I': 'F',
		'J': 'G', 'K': 'H', 'L': 'I', 'M': 'J', 'N': 'K', 'O': 'L',
		'P': 'M', 'Q': 'N', 'R': 'O', 'S': 'P', 'T': 'Q', 'U': 'R',
		'V': 'S', 'W': 'T', 'X': 'U', 'Y': 'V', 'Z': 'W', 'A': 'X',
		'B': 'Y', 'C': 'Z',
	}

	expandKey(3, enc, dec)
	assert.EqualValues(t, myenc, enc)
	assert.EqualValues(t, mydec, dec)
}

func TestNewCipher(t *testing.T) {
	for _, pair := range encryptCaesarTests {
		c, _ := NewCipher(pair.key)

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
