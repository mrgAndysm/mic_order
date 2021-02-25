package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/mrgAndysm/mic_order/com/validata"
)

// InitRouters 初始化路由
func InitRouters() *gin.Engine {
	ginRouter := gin.Default()
	ginRouter.GET("/order", func(context *gin.Context) {
		// context.String(200, "get orderinfos")
		var json validata.TestData

		context.ShouldBindJSON(&json)

		context.String(200, "get orderinfos"+json.Name)
	})

	return ginRouter
}
