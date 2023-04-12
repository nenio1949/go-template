package server

import (
	controller "go-template/controllers"

	"github.com/gin-gonic/gin"
)

// SetApiProjectRoutes returns 项目相关接口
func SetApiProjectRoutes(router *gin.RouterGroup) {
	routerGroup := router.Group("/v1/projects")
	routerGroup.GET("/", controller.GetProjects)           // 获取项目列表
	routerGroup.GET("/:id", controller.GetProject)         // 获取指定项目
	routerGroup.POST("/", controller.AddProject)           // 新增项目
	routerGroup.PUT("/:id", controller.UpdateProject)      // 更新指定项目
	routerGroup.DELETE("/:ids", controller.DeleteProjects) // 删除项目

}
