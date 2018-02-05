package xcrypto

import (
	"crypto/md5"
	"fmt"
)

// HexMd5 returns the MD5 string
func HexMd5(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}
