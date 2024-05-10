package hello

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
)

func GetVersion() string {
    return "v1.0.1"
}

func GetInfo() string {
	return "this is test demo info"
}

func GetMD5Sum(in string) string {
	sum := md5.Sum([]byte(in))

	return hex.EncodeToString(sum[:])
}

func GetSha256(in string) string {
	sh := sha256.New().Sum([]byte(in))

	return hex.EncodeToString(sh)
}
