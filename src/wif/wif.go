package wif

import (
	"encoding/hex"
	"errors"

	"github.com/btcsuite/btcutil/base58"
)

// PrivateKeyChecksumByteLength is the length of a WIF checksum.
const PrivateKeyChecksumByteLength int = 4

// PrivateKeyDecodedByteLength is the length of an uncompressed base58 decoded
// WIF.
const PrivateKeyDecodedByteLength int = 33

// PrivateKeyDecodedCompressedByteLength is the length of a compressed base58
// decoded WIF.
const PrivateKeyDecodedCompressedByteLength int = 34

// ErrorPrivateKeyWifInvalid describes an error where a WIF-encoded private
// key cannot be decoded due to being improperly formatted. This may occur if
// the byte length is incorrect or an unexpected magic number was encountered.
var ErrorPrivateKeyWifInvalid = errors.New("invalid WIF private key format")

// ErrorPrivateKeyChecksumInvalid describes an error where decoding failed due
// to a bad checksum.
var ErrorPrivateKeyChecksumInvalid = errors.New("invalid private key checksum")

// PrivateKey comment
type PrivateKey struct {
	hexString    string
	isCompressed bool
}

// Decode comment
//
// https://en.bitcoin.it/wiki/Wallet_import_format#WIF_to_private_key
func Decode(privateKey string) (*PrivateKey, error) {
	var isCompressed bool

	privateKeyDecoded := base58.Decode(privateKey)
	privateKeyDecodedLength := len(privateKeyDecoded)

	switch privateKeyDecodedLength {
	case PrivateKeyDecodedCompressedByteLength + PrivateKeyChecksumByteLength:
		println("is compressed")
		isCompressed = true
	case PrivateKeyDecodedByteLength + PrivateKeyChecksumByteLength:
		println("is uncompressed")
		isCompressed = false
	default:
		println("is invalid")
		println(privateKeyDecodedLength)
		println(privateKeyDecodedLength)
		return nil, ErrorPrivateKeyWifInvalid
	}

	privKeyBytes := privateKeyDecoded[1 : privateKeyDecodedLength-PrivateKeyChecksumByteLength]

	// println(privateKey)
	// println(hex.EncodeToString(privateKeyDecoded))
	// println(privateKeyDecodedLength)
	// println(hex.EncodeToString(privKeyBytes))

	hexString := hex.EncodeToString(privKeyBytes)

	return &PrivateKey{
		hexString,
		isCompressed,
	}, nil
}
