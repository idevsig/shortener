package types

// ReqCode URL Path
type ReqCode struct {
	Code string `uri:"code" binding:"required"`
}

// ReqQuery 请求参数结构体
type ReqQuery struct {
	Page     int64  `form:"page,default=1" binding:"min=1"`
	PageSize int64  `form:"page_size,default=10" binding:"min=1,max=100"`
	SortBy   string `form:"sort_by,default=created_at" binding:"oneof=created_at updated_at"`
	Order    string `form:"order,default=desc" binding:"oneof=asc desc"`
}

// ResShorten 短链接响应
type ResShorten struct {
	ID          int64  `json:"id"`
	Code        string `json:"code"`
	ShortURL    string `json:"short_url"`
	OriginalURL string `json:"original_url"`
	Describe    string `json:"describe"`
	Status      int8   `json:"status"`
	CreatedTime string `json:"created_time"`
	UpdatedTime string `json:"updated_time"`
}

// ResPage 分页响应
type ResPage struct {
	Page         int64 `json:"page"`          // 当前页码（从1开始）
	PageSize     int64 `json:"page_size"`     // 每页条数（可选返回，便于客户端验证）
	CurrentCount int64 `json:"current_count"` // 当前页实际条数
	TotalItems   int64 `json:"total_items"`   // 符合条件的总条数
	TotalPages   int64 `json:"total_pages"`   // 总页数
}

// ResErr 错误响应
type ResErr struct {
	ErrCode int    `json:"errcode"`
	ErrInfo string `json:"errinfo"`
}

// ResSuccess 成功响应
type ResSuccess struct {
	Data any     `json:"data"` // 数据
	Meta ResPage `json:"meta"` // 元数据
}

// CfgShorten 短链接配置
type CfgShorten struct {
	Length  int    `json:"length"`
	Charset string `json:"charset"`
}

// CfgCache 缓存配置
type CfgCache struct {
	Expire int    `json:"expire"`
	Prefix string `json:"prefix"`
}
