package cache

import (
	"context"
	"time"

	"github.com/bytedance/sonic"

	"go.dsig.cn/shortener/internal/shared"
)

// GetKey 获取缓存key
func GetKey(key string) string {
	return shared.GlobalCacheConfig.Prefix + key
}

// Set 设置缓存
func Set(key string, value any, ttl ...time.Duration) error {
	jsonBytes, err := sonic.Marshal(value)
	if err != nil {
		return err
	}

	expire := 0
	if len(ttl) > 0 {
		expire = int(ttl[0])
	}
	return shared.GlobalCache.Set(context.Background(), key, string(jsonBytes), time.Duration(expire)*time.Second).Err()
}

// Get 获取缓存
func Get(key string) (string, error) {
	data, err := shared.GlobalCache.Get(context.Background(), key).Result()
	if err != nil {
		return "", err
	}
	return data, nil
}

// Delete 删除缓存
func Delete(key string) error {
	return shared.GlobalCache.Del(context.Background(), key).Err()
}

// ClearPrefix 清理缓存前缀
func ClearPrefix(prefix string) error {
	ctx := context.Background()
	iter := shared.GlobalCache.Scan(ctx, 0, prefix+"*", 0).Iterator()

	var keys []string
	for iter.Next(ctx) {
		keys = append(keys, iter.Val())
	}

	if err := iter.Err(); err != nil {
		return err
	}

	if len(keys) > 0 {
		if err := shared.GlobalCache.Del(ctx, keys...).Err(); err != nil {
			return err
		}
	}

	return nil
}

// 批量设置缓存
func BatchSet(values map[string]string, ttl ...time.Duration) error {
	ctx := context.Background()
	pipe := shared.GlobalCache.Pipeline()

	expire := 0
	if len(ttl) > 0 {
		expire = int(ttl[0])
	}

	for key, value := range values {
		// log.Printf("value: %v", value)
		pipe.Set(ctx, key, value, time.Duration(expire)*time.Second)
	}

	if _, err := pipe.Exec(ctx); err != nil {
		return err
	}

	return nil
}
