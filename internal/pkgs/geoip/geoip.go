package geoip

import (
	"fmt"
	"strconv"
	"strings"
)

var shiftIndex = []int{24, 16, 8, 0}

// GeoIP 接口
type GeoIP interface {
	Search(ip uint32) (string, error)
	SearchByStr(ip string) (string, error)
	Parse(data string) *GeoIPData
}

// GeoIPData 地理IP数据
type GeoIPData struct {
	Country  string
	Region   string
	Province string
	City     string
	ISP      string
}

// GeoIPManager 缓存管理器
type GeoIPManager struct {
	Enabled bool
	Mode    string
	GeoIP   GeoIP
}

// NewGeoIPManager 创建 GeoIPManager
func NewGeoIPManager(enabled bool, mode string, geoip GeoIP) *GeoIPManager {
	return &GeoIPManager{Enabled: enabled, Mode: mode, GeoIP: geoip}
}

// Search 搜索IP
func (t *GeoIPManager) Search(ip uint32) (string, error) {
	return t.GeoIP.Search(ip)
}

// SearchByStr 搜索IP
func (t *GeoIPManager) SearchByStr(ip string) (string, error) {
	return t.GeoIP.SearchByStr(ip)
}

// Parse 解析IP数据
func (t *GeoIPManager) Parse(data string) *GeoIPData {
	return t.GeoIP.Parse(data)
}

// IP2Long 将IP转换为long
func (t *GeoIPManager) IP2Long(ip string) (uint32, error) {
	ps := strings.Split(strings.TrimSpace(ip), ".")
	if len(ps) != 4 {
		return 0, fmt.Errorf("invalid ip address `%s`", ip)
	}

	val := uint32(0)
	for i, s := range ps {
		d, err := strconv.Atoi(s)
		if err != nil {
			return 0, fmt.Errorf("the %dth part `%s` is not an integer", i, s)
		}

		if d < 0 || d > 255 {
			return 0, fmt.Errorf("the %dth part `%s` should be an integer bettween 0 and 255", i, s)
		}

		val |= uint32(d) << shiftIndex[i]
	}

	// convert the ip to integer
	return val, nil
}

// Long2IP 将long转换为IP
func (t *GeoIPManager) Long2IP(ip uint32) string {
	return fmt.Sprintf("%d.%d.%d.%d", (ip>>24)&0xFF, (ip>>16)&0xFF, (ip>>8)&0xFF, ip&0xFF)
}
