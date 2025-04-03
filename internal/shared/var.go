package shared

import (
	"gorm.io/gorm"

	"go.dsig.cn/shortener/internal/cache"
	"go.dsig.cn/shortener/internal/pkgs/geoip"
	"go.dsig.cn/shortener/internal/types"
)

var (
	GlobalShorten *types.CfgShorten
	GlobalDB      *gorm.DB
	GlobalAPIKey  string
	GlobalCache   *cache.CacheManager
	GlobalGeoIP   *geoip.GeoIPManager
)
