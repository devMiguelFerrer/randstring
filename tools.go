package randstring

import "crypto/rand"

const randomStringSource = "abcdefghijklmnñopqrstuvwxyzABCDEFGHIJKLMNÑOPQRSTUVWXYZ0123456789_+"

// Tools is the type used to instantiate this module, Any variable ofthis will access
// to all the methods with the receiver *Tools
type Tools struct{}

// RandomString returns a string of random characters of length n.
func (t *Tools) RandomString(n int) string {
	s, r := make([]rune, n), []rune(randomStringSource)
	for i := range s {
		p, _ := rand.Prime(rand.Reader, len(r))
		x, y := p.Uint64(), uint64(len(r))
		s[i] = r[x%y]
	}

	return string(s)
}
