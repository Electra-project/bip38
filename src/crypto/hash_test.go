package crypto

import (
	"encoding/hex"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestCryptoUtilities(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GitHub utilities")
}

var _ = Describe("crypto/hash", func() {
	var (
		passphraseOne     []byte
		passphraseOneHash []byte
		passphraseTwo     []byte
		passphraseTwoHash []byte
	)

	BeforeEach(func() {
		passphraseOne = []byte("test")
		passphraseOneHash, _ = hex.DecodeString("36f028580bb02cc8272a9a020f4200e346e276ae664e45ee80745574e2f5ab80")
		passphraseTwo = []byte(`nGj.c;8ENo7q,CH|L"!=z)HNFsNR`)
		passphraseTwoHash, _ = hex.DecodeString("bd25027aa445d74986d3a661e23a23147077006847d618c715338ebc06f95d4d")
	})

	Describe("hashSha332()", func() {
		Context("WITH predefined passphrases", func() {
			It("SHOULD return the expected hashes", func() {
				Expect(hashSha332(passphraseOne)).To(Equal(passphraseOneHash))
				Expect(hashSha332(passphraseTwo)).To(Equal(passphraseTwoHash))
			})
		})
	})
})
