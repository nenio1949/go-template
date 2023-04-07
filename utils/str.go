package utils

import (
	"math/rand"
	"time"
)

// RandString returns 指定长度随机字符串
// @len: 字符串长度
func RandString(len int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		b := r.Intn(26) + 65
		bytes[i] = byte(b)
	}
	return string(bytes)
}
