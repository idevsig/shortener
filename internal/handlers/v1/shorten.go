package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.dsig.cn/idev/shortener/internal/logics"
	"go.dsig.cn/idev/shortener/internal/pkg"
	"go.dsig.cn/idev/shortener/internal/types"
)

// ShortenHandler 短链接处理器
type ShortenHandler struct {
	handler
	logic *logics.ShortenLogic
}

// NewShortenHandler 创建短链接处理器
func NewShortenHandler() *ShortenHandler {
	t := &ShortenHandler{}
	t.logic = logics.NewShortenLogic()
	return t
}

// ShortenRedirect 短链接跳转
func (t *ShortenHandler) ShortenRedirect(c *gin.Context) {
	var reqUri types.ReqCode
	if err := c.ShouldBindUri(&reqUri); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 2001, "msg": err.Error()})
		return
	}

	err, data := t.logic.ShortenOne(reqUri.Code)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"code": 2002, "msg": err.Error()})
		return
	}

	c.Redirect(http.StatusMovedPermanently, data.OriginalURL)
}

// ShortenAdd 添加短链接
func (t *ShortenHandler) ShortenAdd(c *gin.Context) {
	var reqJson struct {
		Code        string `json:"code"`
		OriginalURL string `json:"original_url" binding:"required"`
		Describe    string `json:"describe"`
	}

	if err := c.ShouldBindJSON(&reqJson); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 2001, "msg": err.Error()})
		return
	}

	// 生成短码
	if reqJson.Code == "" {
		reqJson.Code = pkg.GenerateCode()
	}

	if len(reqJson.Code) > 16 {
		c.JSON(http.StatusBadRequest, gin.H{"code": 2001, "msg": "短码长度不能超过16个字符"})
		return
	}

	err, data := t.logic.ShortenAdd(reqJson.Code, reqJson.OriginalURL, reqJson.Describe)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 2002, "msg": err.Error()})
		return
	}

	item := types.ResShorten{
		Code:        reqJson.Code,
		ShortURL:    data.ShortURL,
		OriginalURL: data.OriginalURL,
		Describe:    data.Describe,
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "success", "data": item})
}

// ShortenList 获取短链接列表
func (t *ShortenHandler) ShortenList(c *gin.Context) {
	err, data := t.logic.ShortenAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 2002, "msg": err.Error()})
		return
	}

	items := make([]types.ResShorten, 0)

	for _, item := range data {
		items = append(items, types.ResShorten{
			Code:        item.ShortCode,
			ShortURL:    item.ShortURL,
			OriginalURL: item.OriginalURL,
			Describe:    item.Describe,
		})
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "success", "data": items})
}

// ShortenFind 获取短链接
func (t *ShortenHandler) ShortenFind(c *gin.Context) {
	var reqUri types.ReqCode
	if err := c.ShouldBindUri(&reqUri); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 2001, "msg": err.Error()})
		return
	}

	err, data := t.logic.ShortenOne(reqUri.Code)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"code": 2002, "msg": err.Error()})
		return
	}

	item := types.ResShorten{
		Code:        reqUri.Code,
		ShortURL:    data.ShortURL,
		OriginalURL: data.OriginalURL,
		Describe:    data.Describe,
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "success", "data": item})
}

// ShortenUpdate 更新短链接
func (t *ShortenHandler) ShortenUpdate(c *gin.Context) {
	var reqUri types.ReqCode
	if err := c.ShouldBindUri(&reqUri); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 2001, "msg": err.Error()})
		return
	}

	var reqJson struct {
		OriginalURL string `json:"original_url" binding:"required"`
		Describe    string `json:"describe"`
	}
	if err := c.ShouldBindJSON(&reqJson); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 2001, "msg": err.Error()})
		return
	}

	err, data := t.logic.ShortenUpdate(reqUri.Code, reqJson.OriginalURL, reqJson.Describe)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 2002, "msg": err.Error()})
		return
	}

	item := types.ResShorten{
		Code:        reqUri.Code,
		ShortURL:    data.ShortURL,
		OriginalURL: data.OriginalURL,
		Describe:    data.Describe,
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "success", "data": item})
}

// ShortenDelete 删除短链接
func (t *ShortenHandler) ShortenDelete(c *gin.Context) {
	var reqUri types.ReqCode
	if err := c.ShouldBindUri(&reqUri); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 2001, "msg": err.Error()})
		return
	}

	err := t.logic.ShortenDelete(reqUri.Code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 2002, "msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "success"})
}
