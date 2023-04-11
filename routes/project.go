package server

import (
	roleController "go-template/controllers"

	"github.com/gin-gonic/gin"
)

// SetApiProjectRoutes returns 项目相关接口
func SetApiProjectRoutes(router *gin.RouterGroup) {
	roleRouter := router.Group("/v1/projects")
	roleRouter.GET("/", roleController.GetProjects)           // 获取项目列表
	roleRouter.GET("/:id", roleController.GetProject)         // 获取指定项目
	roleRouter.POST("/", roleController.AddProject)           // 新增项目
	roleRouter.PUT("/:id", roleController.UpdateProject)      // 更新指定项目
	roleRouter.DELETE("/:ids", roleController.DeleteProjects) // 删除项目

}
