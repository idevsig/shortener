package logics

import (
	"errors"
	"fmt"
	"time"

	"go.dsig.cn/shortener/internal/dal/db/model"
	"go.dsig.cn/shortener/internal/ecodes"
	"go.dsig.cn/shortener/internal/types"
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

// ShortenAdd 添加短链接
func (t *ShortenLogic) ShortenAdd(code string, originalURL string, describe string) (int, types.ResShorten) {
	result := types.ResShorten{}
	existingURL := model.Urls{}

	// 1. 检查短码是否已存在（使用 GORM 的 Find 直接判断）
	if err := t.db.Where("short_code = ?", code).First(&existingURL).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return ecodes.ErrCodeDatabaseError, result // 数据库查询错误
		}
		// 短码不存在，继续流程
	} else {
		return ecodes.ErrCodeConflict, result // 短码已存在
	}

	// 2. 创建新记录
	nowTime := time.Now().Unix()
	newURL := model.Urls{
		ShortCode:   code,
		OriginalURL: originalURL,
		Describe:    describe,
		Status:      0,
		CreatedAt:   nowTime,
		UpdatedAt:   nowTime,
	}

	if err := t.db.Create(&newURL).Error; err != nil {
		return ecodes.ErrCodeDatabaseError, result // 创建失败
	}

	// 3. 构造返回结果
	result = types.ResShorten{
		ID:          newURL.ID,
		Code:        newURL.ShortCode,
		ShortURL:    t.GetSiteURL(newURL.ShortCode),
		OriginalURL: newURL.OriginalURL,
		Describe:    newURL.Describe,
		Status:      newURL.Status,
		CreatedTime: t.GetTimeFormat(nowTime),
		UpdatedTime: t.GetTimeFormat(nowTime),
	}

	return ecodes.ErrCodeSuccess, result
}

// ShortenDelete 删除短链接
func (t *ShortenLogic) ShortenDelete(code string) int {
	if res := t.db.Where("short_code = ?", code).Delete(&model.Urls{}); res.Error != nil {
		return ecodes.ErrCodeDatabaseError
	} else if res.RowsAffected == 0 {
		return ecodes.ErrCodeNotFound
	}
	return ecodes.ErrCodeSuccess
}

// ShortenUpdate 更新短链接
func (t *ShortenLogic) ShortenUpdate(code string, originalURL string, describe string) (int, types.ResShorten) {
	result := types.ResShorten{}

	var existingURL model.Urls
	if err := t.db.Where("short_code = ?", code).First(&existingURL).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ecodes.ErrCodeNotFound, result
		}
		return ecodes.ErrCodeDatabaseError, result
	}

	// 准备更新字段
	updates := make(map[string]any)
	updates["updated_at"] = time.Now().Unix()

	if originalURL != "" {
		updates["original_url"] = originalURL
	}
	if describe != "" {
		updates["describe"] = describe
	}

	if err := t.db.Model(&existingURL).Updates(updates).Error; err != nil {
		return ecodes.ErrCodeDatabaseError, result
	}

	result = types.ResShorten{
		ID:          existingURL.ID,
		Code:        existingURL.ShortCode,
		ShortURL:    t.GetSiteURL(existingURL.ShortCode),
		OriginalURL: existingURL.OriginalURL,
		Describe:    existingURL.Describe,
		Status:      existingURL.Status,
		CreatedTime: t.GetTimeFormat(existingURL.CreatedAt),
		UpdatedTime: t.GetTimeFormat(existingURL.UpdatedAt),
	}

	return ecodes.ErrCodeSuccess, result
}

// ShortenFind 获取短链接
func (t *ShortenLogic) ShortenFind(code string) (int, types.ResShorten) {
	var data model.Urls
	if err := t.db.Where("short_code = ?", code).First(&data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ecodes.ErrCodeNotFound, types.ResShorten{}
		}
		return ecodes.ErrCodeDatabaseError, types.ResShorten{}
	}

	result := types.ResShorten{
		ID:          data.ID,
		Code:        data.ShortCode,
		ShortURL:    t.GetSiteURL(data.ShortCode),
		OriginalURL: data.OriginalURL,
		Describe:    data.Describe,
		Status:      data.Status,
		CreatedTime: t.GetTimeFormat(data.CreatedAt),

		UpdatedTime: t.GetTimeFormat(data.UpdatedAt),
	}

	return ecodes.ErrCodeSuccess, result
}

// ShortenAll 获取所有短链接
func (t *ShortenLogic) ShortenAll(reqQuery types.ReqQuery) (int, []types.ResShorten, types.ResPage) {
	results := make([]types.ResShorten, 0)
	pageInfo := types.ResPage{}

	// 查询数据库
	query := t.db.Model(&model.Urls{}).
		Order(fmt.Sprintf("%s %s", reqQuery.SortBy, reqQuery.Order))

	// 计算总条数
	var total int64
	query = query.Count(&total)
	if query.Error != nil {
		return ecodes.ErrCodeDatabaseError, results, pageInfo
	}

	// 分页查询
	data := make([]model.Urls, 0)
	resDB := query.Offset(int((reqQuery.Page - 1) * reqQuery.PageSize)).
		Limit(int(reqQuery.PageSize)).
		Find(&data)
	if resDB.Error != nil {
		return ecodes.ErrCodeDatabaseError, results, pageInfo
	}

	// 页码信息
	pageInfo.Page = reqQuery.Page
	pageInfo.PageSize = reqQuery.PageSize
	pageInfo.CurrentCount = resDB.RowsAffected
	pageInfo.TotalItems = total
	pageInfo.TotalPages = total / int64(reqQuery.PageSize)
	if total%int64(reqQuery.PageSize) != 0 {
		pageInfo.TotalPages++
	}

	for _, item := range data {
		results = append(results, types.ResShorten{
			ID:          item.ID,
			Code:        item.ShortCode,
			ShortURL:    t.GetSiteURL(item.ShortCode),
			OriginalURL: item.OriginalURL,
			Describe:    item.Describe,
			Status:      item.Status,
			CreatedTime: t.GetTimeFormat(item.CreatedAt),
			UpdatedTime: t.GetTimeFormat(item.UpdatedAt),
		})
	}

	return ecodes.ErrCodeSuccess, results, pageInfo
}
