package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"io"
)

// Cipher a text with a key using AES-256 into an hexadecimal string.
func cipherA256S332(text string, key []byte) (string, error) {
	block, e := aes.NewCipher(key)
	if e != nil {
		return "", e
	}

	gcm, e := cipher.NewGCM(block)
	if e != nil {
		return "", e
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, e = io.ReadFull(rand.Reader, nonce); e != nil {
		return "", e
	}

	return hex.EncodeToString(gcm.Seal(nonce, nonce, []byte(text), nil)), nil
}

// Decipher an array of bytes with a key using AES-256.
func decipherA256S332(byteInput []byte, key []byte) (string, error) {
	block, e := aes.NewCipher(key)
	if e != nil {
		return "", e
	}

	gcm, e := cipher.NewGCM(block)
	if e != nil {
		return "", e
	}

	nonceLength := gcm.NonceSize()
	nonce, textCipher := byteInput[:nonceLength], byteInput[nonceLength:]
	text, e := gcm.Open(nil, nonce, textCipher, nil)
	if e != nil {
		return "", e
	}

	return string(text), nil
}
