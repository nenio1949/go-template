package server

import (
	departmentController "go-template/controllers"

	"github.com/gin-gonic/gin"
)

// SetApiDepartmentRoutes returns 部门相关接口
func SetApiDepartmentRoutes(router *gin.RouterGroup) {
	departmentRouter := router.Group("/v1/departments")
	departmentRouter.GET("/", departmentController.GetDepartments)           // 获取部门列表
	departmentRouter.GET("/:id", departmentController.GetDepartment)         // 获取指定部门
	departmentRouter.POST("/", departmentController.AddDepartment)           // 新增部门
	departmentRouter.PUT("/:id", departmentController.UpdateDepartment)      // 更新指定部门
	departmentRouter.DELETE("/:ids", departmentController.DeleteDepartments) // 删除部门

}
