package utils

import (
	"math/rand"

	"go.dsig.cn/shortener/internal/shared"
)

// GenerateCode 生成短码(6位)
func GenerateCode() string {
	length := shared.GlobalShorten.Length
	charset := shared.GlobalShorten.Charset

	result := make([]byte, length)
	for i := 0; i < length; i++ {
		index := rand.Intn(len(charset))
		result[i] = charset[index]
	}
	return string(result)
}
