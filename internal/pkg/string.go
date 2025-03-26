package pkg

import (
	"math/rand"

	"go.dsig.cn/idev/shortener/internal/shared"
)

// GenerateCode 生成短码(6位)
func GenerateCode() string {
	randSeed()
	length := shared.GlobalShorten.Length
	charset := shared.GlobalShorten.Charset

	result := make([]byte, length)
	for i := 0; i < length; i++ {
		index := rand.Intn(len(charset))
		result[i] = charset[index]
	}
	return string(result)
}
