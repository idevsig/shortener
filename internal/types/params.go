package types

// HistoryParams 历史记录的参数
type HistoryParams struct {
	URLID     int64
	ShortCode string
	IPAddress string
	UserAgent string
	Referer   string
}

// User 用户信息
type User struct {
	Username string `json:"Username"`
	Password string `json:"password,omitempty"`
}
