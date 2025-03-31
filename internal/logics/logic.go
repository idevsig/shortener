package logics

import (
	"time"

	"github.com/spf13/viper"
	"go.dsig.cn/shortener/internal/shared"
	"gorm.io/gorm"
)

// logic 逻辑层
type logic struct {
	db       *gorm.DB
	site_url string
}

// init 初始化
func (t *logic) init() {
	t.db = shared.GlobalDB
	t.site_url = viper.GetString("server.site_url")
	if t.site_url == "" {
		t.site_url = "http://" + viper.GetString("server.address")
	}
}

// GetSiteURL 获取短链接的完整URL
func (t *logic) GetSiteURL(code string) string {
	return t.site_url + "/" + code
}

// GetTimeFormat 获取时间格式
func (t *logic) GetTimeFormat(timestamp int64) string {
	return time.Unix(timestamp, 0).Format("2006-01-02 15:04:05")
}
