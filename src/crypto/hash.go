package crypto

import (
	"golang.org/x/crypto/sha3"
)

// hashSha332 hashes an array of bytes into a 32 bytes hash using SHA-3.
// https://en.wikipedia.org/wiki/SHA-3
func hashSha332(passphrase []byte) []byte {
	h := sha3.New256()
	h.Write(passphrase)

	// Set time to 1 in order to use the maximum available memory
	return h.Sum(nil)
}
