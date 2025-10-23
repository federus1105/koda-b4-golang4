package internals

import (
	"crypto/md5"
	"encoding/hex"
)

func GenerateMD5Hash(password string) string {
	hasher := md5.New()
	hasher.Write([]byte(password))
	return hex.EncodeToString(hasher.Sum(nil))
}
