package types

// ReqCode URL Path
type ReqCode struct {
	Code string `uri:"code" binding:"required"`
}

// ResShorten 短链接响应
type ResShorten struct {
	Code        string `json:"code"`
	ShortURL    string `json:"short_url"`
	OriginalURL string `json:"original_url"`
	Describe    string `json:"describe"`
}

// Shorten 短链接配置
type Shorten struct {
	Length  int    `json:"length"`
	Charset string `json:"charset"`
}
