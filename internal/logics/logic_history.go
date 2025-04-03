package logics

import (
	"strings"
	"time"

	"github.com/ua-parser/uap-go/uaparser"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	"go.dsig.cn/shortener/internal/dal/db/model"
	"go.dsig.cn/shortener/internal/pkgs/geoip"
	"go.dsig.cn/shortener/internal/shared"
	"go.dsig.cn/shortener/internal/types"
)

// HistoryLogic 历史记录逻辑层
type HistoryLogic struct {
	logic
	geoip *geoip.GeoIPManager
}

// NewHistoryLogic 创建历史记录逻辑层
func NewHistoryLogic() *HistoryLogic {
	t := &HistoryLogic{
		geoip: shared.GlobalGeoIP,
	}
	t.db = shared.GlobalDB
	return t
}

// HistoryAdd 添加历史记录
func (t *HistoryLogic) HistoryAdd(params types.HistoryParams) error {
	nowTime := time.Now().Local()

	// 解析用户代理
	uaParser := uaparser.NewFromSaved()
	client := uaParser.Parse(params.UserAgent)
	deviceType := cases.Title(language.English).String(
		simplifyDeviceType(client.Device.ToString()),
	)

	// 初始化地理位置信息
	geoInfo := struct {
		Country  string
		Region   string
		Province string
		City     string
		ISP      string
	}{}

	if t.geoip != nil && t.geoip.Enabled {
		if ipUint32, err := t.geoip.IP2Long(params.IPAddress); err == nil {
			if ipInfo, err := t.geoip.Search(ipUint32); err == nil {
				ipData := t.geoip.Parse(ipInfo)
				geoInfo.Country = ipData.Country
				geoInfo.Region = ipData.Region
				geoInfo.Province = ipData.Province
				geoInfo.City = ipData.City
				geoInfo.ISP = ipData.ISP
			}
		}
	}

	history := model.History{
		UrlID:      params.URLID,
		ShortCode:  params.ShortCode,
		IPAddress:  params.IPAddress,
		UserAgent:  params.UserAgent,
		Referer:    params.Referer,
		Country:    geoInfo.Country,
		Region:     geoInfo.Region,
		Province:   geoInfo.Province,
		City:       geoInfo.City,
		ISP:        geoInfo.ISP,
		DeviceType: deviceType,
		OS:         client.Os.ToString(),
		Browser:    client.UserAgent.ToString(),
		AccessedAt: nowTime,
		CreatedAt:  nowTime,
	}
	// log.Printf("history: %+v\n", history)

	return t.db.Create(&history).Error
}

// simplifyDeviceType 将具体设备型号转换为通用类型（mobile/pc/tablet）
func simplifyDeviceType(device string) string {
	device = strings.ToLower(device)
	switch {
	case strings.Contains(device, "iphone") ||
		strings.Contains(device, "android") ||
		strings.Contains(device, "mobile"):
		return "mobile"
	case strings.Contains(device, "ipad") ||
		strings.Contains(device, "tablet"):
		return "tablet"
	default:
		return "desktop"
	}
}
