package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"go.dsig.cn/shortener/internal/ecodes"
	"go.dsig.cn/shortener/internal/shared"
	"go.dsig.cn/shortener/internal/types"
)

// ApiKeyAuth 检查请求头中的 API Key
func ApiKeyAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKey := c.GetHeader("X-API-KEY")
		if shared.GlobalAPIKey != "" && apiKey != shared.GlobalAPIKey {
			errCode := ecodes.ErrCodeUnauthorized
			c.JSON(http.StatusUnauthorized, types.ResErr{
				ErrCode: errCode,
				ErrInfo: ecodes.GetErrCodeMessage(errCode),
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
