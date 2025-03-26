package model

// TableNameUrlsModel 表名
const TableNameUrlsModel = "urls"

// Urls 短网址表
type Urls struct {
	ID          int64  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	ShortCode   string `gorm:"column:short_code;not null; uniqueIndex" json:"short_code"`
	OriginalURL string `gorm:"column:original_url;not null" json:"original_url"`
	Describe    string `gorm:"column:describe;not null" json:"describe"`
	Status      int32  `gorm:"column:status;default:0;index;not null" json:"status"`
	CreatedAt   int64  `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt   int64  `gorm:"column:updated_at;not null" json:"updated_at"`
}

// TableName Urls的表名
func (u *Urls) TableName() string {
	return TableNameUrlsModel
}
