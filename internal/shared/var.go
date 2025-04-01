package shared

import (
	"github.com/redis/go-redis/v9"
	"go.dsig.cn/shortener/internal/types"
	"gorm.io/gorm"
)

var (
	GlobalShorten     *types.CfgShorten
	GlobalDB          *gorm.DB
	GlobalCache       *redis.Client
	GlobalCacheConfig *types.CfgCache
)
