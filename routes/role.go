package server

import (
	controller "go-template/controllers"

	"github.com/gin-gonic/gin"
)

// SetApiRoleRoutes returns 角色相关接口
func SetApiRoleRoutes(router *gin.RouterGroup) {
	routerGroup := router.Group("/v1/roles")
	routerGroup.GET("/", controller.GetRoles)           // 获取角色列表
	routerGroup.GET("/:id", controller.GetRole)         // 获取指定角色
	routerGroup.POST("/", controller.AddRole)           // 新增角色
	routerGroup.PUT("/:id", controller.UpdateRole)      // 更新指定角色
	routerGroup.DELETE("/:ids", controller.DeleteRoles) // 删除角色

}
