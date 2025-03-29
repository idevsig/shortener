package bootstrap

import (
	"log"

	"github.com/spf13/viper"
	"go.dsig.cn/shortener/internal/shared"
	"go.dsig.cn/shortener/internal/types"
)

// shorten 初始化短链接
func shorten() {
	log.Println("shorten init")
	length := viper.GetInt("shortener.code_length")
	charset := viper.GetString("shortener.code_charset")

	if length <= 0 {
		length = 6
	}

	if charset == "" {
		charset = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	}

	shared.GlobalShorten = &types.Shorten{
		Length:  length,
		Charset: charset,
	}
}
