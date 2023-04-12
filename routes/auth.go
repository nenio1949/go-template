package server

import (
	controller "go-template/controllers"

	"github.com/gin-gonic/gin"
)

// SetApiAuthRoutes returns 授权相关接口
func SetApiAuthRoutes(router *gin.RouterGroup) {
	userRouter := router.Group("/v1")
	userRouter.POST("/login", controller.Login)   // 登录
	userRouter.POST("/logout", controller.Logout) // 登出

	// 版本2
	// userRouter2 := router.Group("/v2")
	// userRouter2.POST("/login", controller.Login)
}
