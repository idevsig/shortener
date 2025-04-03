package utils

import "strings"

// IsURL 判断是否为URL
func IsURL(url string) bool {
	return strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://")
}
