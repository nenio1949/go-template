package server

import (
	controller "go-template/controllers"

	"github.com/gin-gonic/gin"
)

// SetApiDepartmentRoutes returns 部门相关接口
func SetApiDepartmentRoutes(router *gin.RouterGroup) {
	routerGroup := router.Group("/v1/departments")
	routerGroup.GET("/", controller.GetDepartments)           // 获取部门列表
	routerGroup.GET("/:id", controller.GetDepartment)         // 获取指定部门
	routerGroup.POST("/", controller.AddDepartment)           // 新增部门
	routerGroup.PUT("/:id", controller.UpdateDepartment)      // 更新指定部门
	routerGroup.DELETE("/:ids", controller.DeleteDepartments) // 删除部门

}
