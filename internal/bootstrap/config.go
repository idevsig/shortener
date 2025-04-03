package bootstrap

import (
	"github.com/spf13/viper"

	"go.dsig.cn/shortener/internal/shared"
	"go.dsig.cn/shortener/internal/types"
)

// initSharedConfig 初始化共享配置
func initSharedConfig() {
	// log.Println("shorten init")
	length := viper.GetInt("shortener.code_length")
	charset := viper.GetString("shortener.code_charset")

	if length <= 0 {
		length = 6
	}

	if charset == "" {
		charset = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	}

	shared.GlobalShorten = &types.CfgShorten{
		Length:  length,
		Charset: charset,
	}

	shared.GlobalAPIKey = viper.GetString("api.key")
}
