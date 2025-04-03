package main

import (
	_ "go.dsig.cn/shortener/internal/bootstrap"

	"github.com/spf13/viper"

	"go.dsig.cn/shortener/internal/routers"
)

func main() {
	addr := viper.GetString("server.address")
	if addr == "" {
		addr = ":8080"
	}

	r := routers.NewRouter()
	if err := r.Run(addr); err != nil {
		panic("run server failed: " + err.Error())
	}
}
