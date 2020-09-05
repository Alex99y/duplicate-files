package crypto

import (
	"crypto/sha256"
	"fmt"
)

// GetFileHash returns the sha256 of the file's content
func GetFileHash(file []byte) string {
	checksum := sha256.Sum256(file)
	return fmt.Sprintf("%x", checksum)
}
