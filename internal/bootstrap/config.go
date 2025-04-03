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

// initDefaultConfig 初始化默认配置
func initDefaultConfig() {
	// 服务器配置
	viper.SetDefault("server.address", ":8080")
	viper.SetDefault("server.trusted-platform", "")
	viper.SetDefault("server.site_url", "http://localhost:8080")

	// API配置
	viper.SetDefault("api.url", "")
	viper.SetDefault("api.key", "")

	// 短链生成配置
	viper.SetDefault("shortener.code_length", 6)
	viper.SetDefault("shortener.code_charset", "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	// 数据库配置
	viper.SetDefault("database.type", "sqlite")
	viper.SetDefault("database.log_level", 0)

	// SQLite配置
	viper.SetDefault("database.sqlite.path", "data/shortener.db")

	// PostgreSQL配置
	viper.SetDefault("database.postgres.host", "localhost")
	viper.SetDefault("database.postgres.port", 5432)
	viper.SetDefault("database.postgres.user", "postgres")
	viper.SetDefault("database.postgres.password", "postgres")
	viper.SetDefault("database.postgres.database", "shortener")
	viper.SetDefault("database.postgres.sslmode", "disable")
	viper.SetDefault("database.postgres.timezone", "Asia/Shanghai")

	// MySQL配置
	viper.SetDefault("database.mysql.host", "localhost")
	viper.SetDefault("database.mysql.port", 3306)
	viper.SetDefault("database.mysql.user", "root")
	viper.SetDefault("database.mysql.password", "root")
	viper.SetDefault("database.mysql.database", "shortener")
	viper.SetDefault("database.mysql.charset", "utf8mb4")
	viper.SetDefault("database.mysql.parse_time", true)
	viper.SetDefault("database.mysql.loc", "Local")

	// 缓存配置
	viper.SetDefault("cache.enabled", false)
	viper.SetDefault("cache.type", "redis")
	viper.SetDefault("cache.expire", 3600)
	viper.SetDefault("cache.prefix", "shorten:")

	// Redis配置
	viper.SetDefault("cache.redis.host", "localhost")
	viper.SetDefault("cache.redis.port", 6379)
	viper.SetDefault("cache.redis.password", "")
	viper.SetDefault("cache.redis.db", 0)

	// Valkey配置
	viper.SetDefault("cache.valkey.host", "localhost")
	viper.SetDefault("cache.valkey.port", 6379)
	viper.SetDefault("cache.valkey.username", "")
	viper.SetDefault("cache.valkey.password", "")
	viper.SetDefault("cache.valkey.db", 0)
}
