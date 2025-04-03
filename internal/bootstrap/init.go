package bootstrap

import (
	"fmt"
	"path/filepath"

	"github.com/spf13/viper"
)

const (
	configName = "config"
	configType = "toml"
)

func init() {
	viper.SetConfigName(configName)
	// viper.SetConfigType("toml")
	viper.AddConfigPath("./config/dev")
	viper.AddConfigPath("./config/prod")
	viper.AddConfigPath("./config")
	viper.AddConfigPath(".")

	initDefaultConfig()

	if err := viper.ReadInConfig(); err != nil {
		// 处理配置文件不存在的情况
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// 创建并保存默认配置
			configFile := filepath.Join(configName + "." + configType)
			if err := viper.SafeWriteConfigAs(configFile); err != nil {
				panic(fmt.Errorf("write config failed: %s\n%v", configFile, err))
			}
		} else {
			// 其他类型的配置错误
			panic(
				fmt.Errorf("fatal error config file: %w", err),
			)
		}
	}

	// log.Printf("config: %+v\n", viper.AllSettings())
	bootstrap()
}

// 初始化
func bootstrap() {
	// init shared config
	initSharedConfig()

	// init db
	initDB()

	// init cache
	initCache()
}
