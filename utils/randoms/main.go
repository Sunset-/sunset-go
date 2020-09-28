package randoms

import (
	"bytes"
	"fmt"
	"math/rand"
	"time"
)

//随机x位字符串
func RandomStr(len int) string {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		b := r.Intn(26) + 65
		bytes[i] = byte(b)
	}
	return string(bytes)
}

//随机整数字符串
func RandomIntStr(len int) string {
	times := len / 8
	if len%8 > 0 {
		times += 1
	}
	var buff bytes.Buffer
	for i := 0; i < times; i++ {
		buff.WriteString(fmt.Sprintf("%08v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(100000000)))
	}
	return buff.String()[:len]
}
