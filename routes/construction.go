package server

import (
	controller "go-template/controllers"

	"github.com/gin-gonic/gin"
)

// SetApiConstructionRoutes returns 施工作业相关接口
func SetApiConstructionRoutes(router *gin.RouterGroup) {
	routerGroup := router.Group("/v1/constructions")

	routerGroup.GET("/plans", controller.GetConstructionPlans)       // 获取施工作业计划列表
	routerGroup.GET("/plans/:id", controller.GetConstructionPlan)    // 获取指定施工作业计划
	routerGroup.POST("/plans", controller.AddConstructionPlan)       // 新增施工作业计划
	routerGroup.PUT("/plans/:id", controller.UpdateConstructionPlan) // 更新指定施工作业计划
	routerGroup.DELETE("/:ids", controller.DeleteConstructions)      // 删除施工作业
}
