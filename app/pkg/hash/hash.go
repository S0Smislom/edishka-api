package hash

import (
	"crypto/sha1"
	"fmt"
)

func GenerateHash(s, salt string) string {
	hash := sha1.New()
	hash.Write([]byte(s))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
