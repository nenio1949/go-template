package utils

import (
	"crypto/md5"
	"encoding/hex"
)

// EncodeMD5 md5 encryption
func MD5(str []byte, b ...byte) string {
	h := md5.New()
	h.Write(str)
	return hex.EncodeToString(h.Sum(b))
}
