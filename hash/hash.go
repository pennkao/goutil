package hash

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
)

// Md5Byte 获取字节数组md5值
func Md5Byte(s []byte) string {
	h := md5.New()
	h.Write(s)
	return hex.EncodeToString(h.Sum(nil))
}

// Sha1Byte 获取节数组sha1值
func Sha1Byte(s []byte) string {
	h := sha1.New()
	h.Write(s)
	return hex.EncodeToString(h.Sum(nil))
}

// Sha256Byte 获取节数组sha256值
func Sha256Byte(s []byte) string {
	h := sha256.New()
	h.Write(s)
	return hex.EncodeToString(h.Sum(nil))
}

// Sha512Byte 获取节数组sha512值
func Sha512Byte(s []byte) string {
	h := sha512.New()
	h.Write(s)
	return hex.EncodeToString(h.Sum(nil))
}

// Md5String 获取字符串md5值
func Md5String(s string) string {
	return Md5Byte([]byte(s))
}

// Sha1String 获取字符串sha1值
func Sha1String(s string) string {
	return Sha1Byte([]byte(s))
}

// Sha256String 获取字符串sha256值
func Sha256String(s string) string {
	return Sha256Byte([]byte(s))
}

// Sha512String 获取字符串sha512值
func Sha512String(s string) string {
	return Sha512Byte([]byte(s))
}
