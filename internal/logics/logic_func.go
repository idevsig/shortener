package logics

import "strings"

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
