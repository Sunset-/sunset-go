package signs

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
)

//base64编码
func Base64Encode(bytes []byte) string {
	return base64.StdEncoding.EncodeToString(bytes)
}

//base64解码
func Base64Decode(base64Str string) (bytes []byte, err error) {
	return base64.StdEncoding.DecodeString(base64Str)
}

//base64解码Must
func Base64DecodeMust(base64Str string) (bytes []byte) {
	res, _ := Base64Decode(base64Str)
	return res
}

//生成32位MD5
func MD5(text string) string {
	ctx := md5.New()
	ctx.Write([]byte(text))
	return hex.EncodeToString(ctx.Sum(nil))
}

//SHA256
func Sha256(src []byte) []byte {
	a := sha256.Sum256(src)
	return a[:]
}
