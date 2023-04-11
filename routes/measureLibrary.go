package server

import (
	roleController "go-template/controllers"

	"github.com/gin-gonic/gin"
)

// SetApiMeasureLibraryRoutes returns 措施库相关接口
func SetApiMeasureLibraryRoutes(router *gin.RouterGroup) {
	roleRouter := router.Group("/v1/measure-libraries")

	roleRouter.POST("/", roleController.AddMeasureLibrary) // 新增措施库
}
