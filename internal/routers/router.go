package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.dsig.cn/shortener/internal/handlers"
)

func NewRouter() *gin.Engine {
	g := gin.Default()

	// swagger api docs
	// g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// pprof router 性能分析路由
	// 默认关闭，开发环境下可以打开
	// 访问方式: HOST/debug/pprof
	// 通过 HOST/debug/pprof/profile 生成profile
	// 查看分析图 go tool pprof -http=:5000 profile
	// see: https://github.com/gin-contrib/pprof
	// pprof.Register(g)

	// PING
	g.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// favicon.ico
	g.GET("/favicon.ico", func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Not Found",
		})
	})

	shortener := handlers.Handle.ShortenHandler

	apiV1 := g.Group("/api/v1")
	apiV1.Use()
	{
		apiV1.POST("/shorten", shortener.ShortenAdd)
		apiV1.GET("/shorten", shortener.ShortenList)
		apiV1.GET("/shorten/:code", shortener.ShortenFind)
		apiV1.PUT("/shorten/:code", shortener.ShortenUpdate)
		apiV1.DELETE("/shorten/:code", shortener.ShortenDelete)
	}

	// 短链接跳转路由
	g.GET("/:code", shortener.ShortenRedirect)
	g.HEAD("/:code", shortener.ShortenRedirect)

	return g
}
