package routers

import "github.com/gin-gonic/gin"

// InitRouters 初始化路由
func InitRouters() *gin.Engine {
	ginRouter := gin.Default()
	ginRouter.GET("/order", func(context *gin.Context) {
		context.String(200, "get orderinfos")
	})

	return ginRouter
}
