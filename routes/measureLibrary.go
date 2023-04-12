package server

import (
	controller "go-template/controllers"

	"github.com/gin-gonic/gin"
)

// SetApiMeasureLibraryRoutes returns 措施库相关接口
func SetApiMeasureLibraryRoutes(router *gin.RouterGroup) {
	routerGroup := router.Group("/v1/measure-libraries")
	routerGroup.GET("/", controller.GetMeasureLibraries)           // 获取措施库列表
	routerGroup.GET("/:id", controller.GetMeasureLibrary)          // 获取指定措施库
	routerGroup.POST("/", controller.AddMeasureLibrary)            // 新增措施库
	routerGroup.PUT("/:id", controller.UpdateMeasureLibrary)       // 更新指定措施库
	routerGroup.DELETE("/:ids", controller.DeleteMeasureLibraries) // 删除措施库
}
