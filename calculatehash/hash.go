package calculatehash

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"log"
	"os"
)

func CalculateSha256Checksum(fileName string) string {
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}
	hashInBytes := h.Sum(nil)
	sha256Hash := hex.EncodeToString(hashInBytes)
	return sha256Hash
}