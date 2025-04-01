package bootstrap

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigName("config")
	// viper.SetConfigType("toml")
	viper.AddConfigPath("./config/dev")
	viper.AddConfigPath("./config/prod")
	viper.AddConfigPath("./config")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	log.Printf("config: %+v\n", viper.AllSettings())

	bootstrap()
}

// 初始化
func bootstrap() {
	// init shorten
	initShortenConfig()

	// init db
	initDB()

	// init cache
	initCache()
}
