package model

// TableNameUrlsModel 表名
const TableNameUrlsModel = "urls"

// Urls 短网址表
type Urls struct {
	ID          int64  `gorm:"column:id;primaryKey;autoIncrement;type:bigint" json:"id"`
	ShortCode   string `gorm:"column:short_code;type:varchar(16);uniqueIndex;not null" json:"short_code"`
	OriginalURL string `gorm:"column:original_url;type:varchar(2048);not null" json:"original_url"`
	Describe    string `gorm:"column:describe;type:varchar(255)" json:"describe"`
	Status      int8   `gorm:"column:status;type:smallint;default:0;index;not null" json:"status"`
	CreatedAt   int64  `gorm:"column:created_at;type:bigint;not null" json:"created_at"`
	UpdatedAt   int64  `gorm:"column:updated_at;type:bigint;not null" json:"updated_at"`
}

// TableName Urls的表名
func (u *Urls) TableName() string {
	return TableNameUrlsModel
}
