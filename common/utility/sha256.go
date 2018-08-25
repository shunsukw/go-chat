package utility

import (
	"crypto/sha256"
	"fmt"
)

// SHA256OfString ...
func SHA256OfString(input string) string {
	sum := sha256.Sum256([]byte(input))
	return fmt.Sprintf("%x", sum)
}
