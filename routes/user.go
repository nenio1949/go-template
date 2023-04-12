package server

import (
	controller "go-template/controllers"

	"github.com/gin-gonic/gin"
)

// SetApiUserRoutes returns 用户相关接口
func SetApiUserRoutes(router *gin.RouterGroup) {
	routerGroup := router.Group("/v1/users")
	routerGroup.GET("/", controller.GetUsers)           // 获取用户列表
	routerGroup.GET("/:id", controller.GetUser)         // 获取指定用户
	routerGroup.POST("/", controller.AddUser)           // 新增用户
	routerGroup.PUT("/:id", controller.UpdateUser)      // 更新指定用户
	routerGroup.DELETE("/:ids", controller.DeleteUsers) // 删除用户

}
