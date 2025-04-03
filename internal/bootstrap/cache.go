package bootstrap

import (
	"fmt"

	"github.com/bytedance/sonic"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"

	"go.dsig.cn/shortener/internal/cache"
	"go.dsig.cn/shortener/internal/dal/db/model"
	"go.dsig.cn/shortener/internal/shared"
	"go.dsig.cn/shortener/internal/types"
)

// initCache 初始化缓存
func initCache() {
	cacheType := viper.GetString("cache.type")
	if cacheType == "" {
		cacheType = "redis"
	}

	switch cacheType {
	case "redis":
		shared.GlobalCache = redisCache()
	default:
		panic("cache type not support: " + cacheType)
	}

	initCacheConfig()
	loadAllShorten()
}

// initCacheConfig 初始化缓存配置
func initCacheConfig() {
	shared.GlobalCacheConfig = &types.CfgCache{
		Expire: viper.GetInt("cache.expire"),
		Prefix: viper.GetString("cache.prefix"),
	}
}

// redisCache 初始化 redis
func redisCache() *redis.Client {
	host := viper.GetString("cache.redis.host")
	if host == "" {
		host = "localhost"
	}
	port := viper.GetInt("cache.redis.port")
	if port == 0 {
		port = 6379
	}
	addr := fmt.Sprintf("%s:%d", host, port)
	password := viper.GetString("cache.redis.password")
	db := viper.GetInt("cache.redis.db")
	if db == 0 {
		db = 0
	}

	// 创建 redis 客户端
	redisClient := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	return redisClient
}

// loadAllShorten 加载所有短链接
func loadAllShorten() {
	var shortens []model.Urls
	if err := shared.GlobalDB.Find(&shortens).Error; err != nil {
		panic("load all shorten failed: " + err.Error())
	}

	cache.ClearPrefix(cache.GetKey(""))

	items := make(map[string]string, len(shortens))
	for _, shorten := range shortens {
		item, _ := sonic.Marshal(shorten)
		items[cache.GetKey(shorten.ShortCode)] = string(item)
	}

	cache.BatchSet(items)
}
