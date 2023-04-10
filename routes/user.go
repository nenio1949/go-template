package server

import (
	userController "go-template/controllers"

	"github.com/gin-gonic/gin"
)

// SetApiUserRoutes returns 用户相关接口
func SetApiUserRoutes(router *gin.RouterGroup) {
	userRouter := router.Group("/v1/users")
	userRouter.GET("/", userController.GetUsers)           // 获取用户列表
	userRouter.GET("/:id", userController.GetUser)         // 获取指定用户
	userRouter.POST("/", userController.AddUser)           // 新增用户
	userRouter.PUT("/:id", userController.UpdateUser)      // 更新指定用户
	userRouter.DELETE("/:ids", userController.DeleteUsers) // 删除用户

}
