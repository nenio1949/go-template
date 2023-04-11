package server

import (
	roleController "go-template/controllers"

	"github.com/gin-gonic/gin"
)

// SetApiConstructionRoutes returns 施工作业相关接口
func SetApiConstructionRoutes(router *gin.RouterGroup) {
	roleRouter := router.Group("/v1/constructions")

	roleRouter.POST("/", roleController.AddConstruction) // 新增施工作业
}
