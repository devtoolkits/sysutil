package sysutil

import (
	"crypto/md5"
	"encoding/hex"
)

// Md5 get md5sum of raw
// faster and less memory overhead
func Md5(raw string) string {
	h := md5.Sum([]byte(raw))
	return hex.EncodeToString(h[:])
}
