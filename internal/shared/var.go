package shared

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"

	"go.dsig.cn/shortener/internal/types"
)

var (
	GlobalShorten     *types.CfgShorten
	GlobalDB          *gorm.DB
	GlobalCache       *redis.Client
	GlobalCacheConfig *types.CfgCache
	GlobalAPIKey      string
)
