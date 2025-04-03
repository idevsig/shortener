package geoip

import (
	"fmt"
	"strings"

	"github.com/lionsoul2014/ip2region/binding/golang/xdb"
)

// IP2Region 结构体
type IP2Region struct {
	path     string
	mode     string
	searcher *xdb.Searcher
}

// NewIP2Region 创建 IP2Region 结构体
func NewIP2Region(dbPath string, loadMode string) (*IP2Region, error) {
	var err error
	var searcher *xdb.Searcher

	switch loadMode {
	case "vector":
		searcher, err = loadVectorIndex(dbPath)
	case "memory":
		searcher, err = loadMemoryIndex(dbPath)
	case "file":
		searcher, err = loadFileIndex(dbPath)
	default:
		return nil, fmt.Errorf("invalid load mode: %s", loadMode)
	}

	return &IP2Region{
		path:     dbPath,
		mode:     loadMode,
		searcher: searcher,
	}, err
}

// loadVectorIndex 加载向量索引
func loadVectorIndex(dbPath string) (*xdb.Searcher, error) {
	vIndex, err := xdb.LoadVectorIndexFromFile(dbPath)
	if err != nil {
		return nil, err
	}
	return xdb.NewWithVectorIndex(dbPath, vIndex)
}

// loadMemoryIndex 加载内存索引
func loadMemoryIndex(dbPath string) (*xdb.Searcher, error) {
	cBuff, err := xdb.LoadContentFromFile(dbPath)
	if err != nil {
		return nil, err
	}
	return xdb.NewWithBuffer(cBuff)
}

// loadFileIndex 加载文件索引
func loadFileIndex(dbPath string) (*xdb.Searcher, error) {
	return xdb.NewWithFileOnly(dbPath)
}

// Search 搜索IP
func (t *IP2Region) Search(ip uint32) (string, error) {
	return t.searcher.Search(ip)
}

// SearchByStr 搜索IP
func (t *IP2Region) SearchByStr(ip string) (string, error) {
	return t.searcher.SearchByStr(ip)
}

// Parse 解析IP数据
func (t *IP2Region) Parse(data string) *GeoIPData {
	parts := strings.Split(data, "|")
	var country, region, province, city, isp string
	if len(parts) > 0 {
		country = parts[0]
	}
	if len(parts) > 1 {
		region = parts[1]
	}
	if len(parts) > 2 {
		province = parts[2]
	}
	if len(parts) > 3 {
		city = parts[3]
	}
	if len(parts) > 4 {
		isp = parts[4]
	}

	return &GeoIPData{
		Country:  country,
		Region:   region,
		Province: province,
		City:     city,
		ISP:      isp,
	}
}
