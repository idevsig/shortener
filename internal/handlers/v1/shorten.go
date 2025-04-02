package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.dsig.cn/shortener/internal/ecodes"
	"go.dsig.cn/shortener/internal/logics"
	"go.dsig.cn/shortener/internal/pkg"
	"go.dsig.cn/shortener/internal/types"
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
		c.JSON(http.StatusBadRequest, t.JsonRespErr(ecodes.ErrCodeInvalidParam))
		return
	}

	errCode, data := t.logic.ShortenFind(reqUri.Code)
	if errCode != ecodes.ErrCodeSuccess {
		errInfo := t.JsonRespErr(errCode)
		if errCode == ecodes.ErrCodeNotFound {
			c.JSON(http.StatusNotFound, errInfo)
		} else {
			c.JSON(http.StatusInternalServerError, errInfo)
		}
		return
	}

	c.Redirect(http.StatusFound, data.OriginalURL)
}

// ShortenAdd 添加短链接
func (t *ShortenHandler) ShortenAdd(c *gin.Context) {
	var reqJson struct {
		Code        string `json:"code,omitempty"`
		OriginalURL string `json:"original_url" binding:"required,url"`
		Describe    string `json:"describe,omitempty"`
	}

	if err := c.ShouldBindJSON(&reqJson); err != nil {
		c.JSON(http.StatusBadRequest, t.JsonRespErr(ecodes.ErrCodeInvalidParam))
		return
	}

	if reqJson.OriginalURL != "" && !t.IsURL(reqJson.OriginalURL) {
		c.JSON(http.StatusBadRequest, t.JsonRespErr(ecodes.ErrCodeInvalidParam))
		return
	}

	// 生成短码
	if reqJson.Code == "" {
		reqJson.Code = pkg.GenerateCode()
	}

	if len(reqJson.Code) > 16 {
		c.JSON(http.StatusBadRequest, t.JsonRespErr(ecodes.ErrCodeBadRequest))
		return
	}

	errCode, data := t.logic.ShortenAdd(reqJson.Code, reqJson.OriginalURL, reqJson.Describe)
	if errCode != 0 {
		errInfo := t.JsonRespErr(errCode)
		if errCode == ecodes.ErrCodeConflict {
			c.JSON(http.StatusConflict, errInfo)
		} else {
			c.JSON(http.StatusInternalServerError, errInfo)
		}
		return
	}

	c.Header("Location", c.Request.RequestURI+"/"+data.Code)
	c.JSON(http.StatusCreated, data)
}

// ShortenDelete 删除短链接
func (t *ShortenHandler) ShortenDelete(c *gin.Context) {
	var reqUri types.ReqCode
	if err := c.ShouldBindUri(&reqUri); err != nil {
		c.JSON(http.StatusBadRequest, t.JsonRespErr(ecodes.ErrCodeInvalidParam))
		return
	}

	errCode := t.logic.ShortenDelete(reqUri.Code)
	if errCode != ecodes.ErrCodeSuccess {
		errInfo := t.JsonRespErr(errCode)
		if errCode == ecodes.ErrCodeNotFound {
			c.JSON(http.StatusNotFound, errInfo)
		} else {
			c.JSON(http.StatusInternalServerError, errInfo)
		}
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

// ShortenUpdate 更新短链接
func (t *ShortenHandler) ShortenUpdate(c *gin.Context) {
	var reqUri types.ReqCode
	if err := c.ShouldBindUri(&reqUri); err != nil {
		c.JSON(http.StatusBadRequest, t.JsonRespErr(ecodes.ErrCodeInvalidParam))
		return
	}

	var reqJson struct {
		OriginalURL string `json:"original_url,omitempty" binding:"omitempty,url"`
		Describe    string `json:"describe,omitempty"`
	}
	if err := c.ShouldBindJSON(&reqJson); err != nil {
		c.JSON(http.StatusBadRequest, t.JsonRespErr(ecodes.ErrCodeInvalidParam))
		return
	}

	if reqJson.OriginalURL != "" && !t.IsURL(reqJson.OriginalURL) {
		c.JSON(http.StatusBadRequest, t.JsonRespErr(ecodes.ErrCodeInvalidParam))
		return
	}

	errCode, data := t.logic.ShortenUpdate(reqUri.Code, reqJson.OriginalURL, reqJson.Describe)
	if errCode != ecodes.ErrCodeSuccess {
		errInfo := t.JsonRespErr(errCode)
		if errCode == ecodes.ErrCodeNotFound {
			c.JSON(http.StatusNotFound, errInfo)
		} else {
			c.JSON(http.StatusInternalServerError, errInfo)
		}
		return
	}

	c.JSON(http.StatusOK, data)
}

// ShortenFind 获取短链接
func (t *ShortenHandler) ShortenFind(c *gin.Context) {
	var reqUri types.ReqCode
	if err := c.ShouldBindUri(&reqUri); err != nil {
		c.JSON(http.StatusBadRequest, t.JsonRespErr(ecodes.ErrCodeInvalidParam))
		return
	}

	errCode, data := t.logic.ShortenFind(reqUri.Code)
	if errCode != ecodes.ErrCodeSuccess {
		errInfo := t.JsonRespErr(errCode)
		if errCode == ecodes.ErrCodeNotFound {
			c.JSON(http.StatusNotFound, errInfo)
		} else {
			c.JSON(http.StatusInternalServerError, errInfo)
		}
		return
	}

	c.JSON(http.StatusOK, data)
}

// ShortenList 获取短链接列表
func (t *ShortenHandler) ShortenList(c *gin.Context) {
	var reqQuery types.ReqQuery
	if err := c.ShouldBindQuery(&reqQuery); err != nil {
		c.JSON(http.StatusBadRequest, t.JsonRespErr(ecodes.ErrCodeInvalidParam))
		return
	}

	errCode, data, pageInfo := t.logic.ShortenAll(reqQuery)
	if errCode != ecodes.ErrCodeSuccess {
		errInfo := t.JsonRespErr(errCode)
		if errCode == ecodes.ErrCodeDatabaseError {
			c.JSON(http.StatusInternalServerError, errInfo)
		} else {
			c.JSON(http.StatusBadRequest, errInfo)
		}
		return
	}

	result := types.ResSuccess[[]types.ResShorten]{
		Data: data,
		Meta: pageInfo,
	}

	c.JSON(http.StatusOK, result)
}
