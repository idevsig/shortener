package logics

import (
	"errors"
	"time"

	"go.dsig.cn/idev/shortener/internal/dal/db/model"
	"go.dsig.cn/idev/shortener/internal/types"
	"gorm.io/gorm"
)

// ShortenLogic 短链接逻辑层
type ShortenLogic struct {
	logic
}

// NewShortenLogic 创建短链接逻辑层
func NewShortenLogic() *ShortenLogic {
	t := &ShortenLogic{}
	t.init()
	return t
}

// ShortenFind 获取短链接
func (t *ShortenLogic) ShortenOne(code string) (err error, url types.Url) {
	var data model.Urls
	resDB := t.db.Where("short_code = ?", code).First(&data)
	if resDB.Error != nil {
		return resDB.Error, url
	}

	url.ID = data.ID
	url.ShortCode = data.ShortCode
	url.ShortURL = t.GetSiteURL(data.ShortCode)
	url.OriginalURL = data.OriginalURL
	url.Describe = data.Describe
	url.Status = data.Status
	return nil, url
}

// ShortenAdd 添加短链接
func (t *ShortenLogic) ShortenAdd(code string, originalURL string, describe string) (err error, url types.Url) {
	var data model.Urls
	resDB := t.db.Where("short_code = ?", code).First(&data)
	if data.ID > 0 {
		return errors.New("code already exists"), url
	}

	if resDB.Error != nil && resDB.Error != gorm.ErrRecordNotFound {
		return resDB.Error, url
	}

	nowTime := time.Now().Unix()
	data.ShortCode = code
	data.OriginalURL = originalURL
	data.Describe = describe
	data.Status = 0
	data.CreatedAt = nowTime
	data.UpdatedAt = nowTime
	resDB = t.db.Create(&data)
	if resDB.Error != nil {
		return resDB.Error, url
	}

	url.ID = data.ID
	url.ShortCode = data.ShortCode
	url.ShortURL = t.GetSiteURL(data.ShortCode)
	url.OriginalURL = data.OriginalURL
	url.Describe = data.Describe
	url.Status = data.Status
	url.CreatedAt = nowTime
	url.UpdatedAt = nowTime
	return nil, url
}

// ShortenAll 获取所有短链接
func (t *ShortenLogic) ShortenAll() (err error, urls []types.Url) {
	var data []model.Urls
	resDB := t.db.Find(&data)
	if resDB.Error != nil {
		return resDB.Error, urls
	}

	for _, item := range data {
		urls = append(urls, types.Url{
			ID:          item.ID,
			ShortCode:   item.ShortCode,
			ShortURL:    t.GetSiteURL(item.ShortCode),
			OriginalURL: item.OriginalURL,
			Describe:    item.Describe,
			Status:      item.Status,
			CreatedAt:   item.CreatedAt,
			UpdatedAt:   item.UpdatedAt,
		})
	}
	return nil, urls
}

// ShortenUpdate 更新短链接
func (t *ShortenLogic) ShortenUpdate(code string, originalURL string, describe string) (err error, url types.Url) {
	var data model.Urls
	resDB := t.db.Where("short_code = ?", code).First(&data)
	if resDB.Error != nil {
		return resDB.Error, url
	}

	if originalURL != "" {
		data.OriginalURL = originalURL
	}
	if describe != "" {
		data.Describe = describe
	}
	data.UpdatedAt = time.Now().Unix()

	resDB = t.db.Save(&data)
	if resDB.Error != nil {
		return resDB.Error, url
	}

	url.ID = data.ID
	url.ShortCode = data.ShortCode
	url.ShortURL = t.GetSiteURL(data.ShortCode)
	url.OriginalURL = data.OriginalURL
	url.Describe = data.Describe
	url.Status = data.Status
	url.CreatedAt = data.CreatedAt
	url.UpdatedAt = data.UpdatedAt
	return nil, url
}

// ShortenDelete 删除短链接
func (t *ShortenLogic) ShortenDelete(code string) (err error) {
	resDB := t.db.Where("short_code = ?", code).Delete(&model.Urls{})
	if resDB.Error != nil {
		return resDB.Error
	}
	return nil
}
