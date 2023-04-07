package server

import (
	userController "go-template/controllers"

	"github.com/gin-gonic/gin"
)

// SetApiAuthRoutes returns 授权相关接口
func SetApiAuthRoutes(router *gin.RouterGroup) {
	userRouter := router.Group("/v1")
	userRouter.POST("/login", userController.Login)   // 登录
	userRouter.POST("/logout", userController.Logout) // 登出
}
