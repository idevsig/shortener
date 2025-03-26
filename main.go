package main

import (
	_ "go.dsig.cn/idev/shortener/internal/bootstrap"

	"github.com/spf13/viper"
	"go.dsig.cn/idev/shortener/internal/routers"
)

func main() {
	addr := viper.GetString("server.address")
	if addr == "" {
		addr = ":8080"
	}

	r := routers.NewRouter()
	r.Run(addr)
}
