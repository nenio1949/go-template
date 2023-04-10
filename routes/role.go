package server

import (
	roleController "go-template/controllers"

	"github.com/gin-gonic/gin"
)

// SetApiRoleRoutes returns 角色相关接口
func SetApiRoleRoutes(router *gin.RouterGroup) {
	roleRouter := router.Group("/v1/roles")
	roleRouter.GET("/", roleController.GetRoles)           // 获取角色列表
	roleRouter.GET("/:id", roleController.GetRole)         // 获取指定角色
	roleRouter.POST("/", roleController.AddRole)           // 新增角色
	roleRouter.PUT("/:id", roleController.UpdateRole)      // 更新指定角色
	roleRouter.DELETE("/:ids", roleController.DeleteRoles) // 删除角色

}
