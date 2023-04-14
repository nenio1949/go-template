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

	routerGroup.GET("/:id", controller.GetConstruction)                   // 获取指定施工作业
	routerGroup.PUT("/:id", controller.UpdateConstruction)                // 更新指定施工作业
	routerGroup.POST("/approves/:id", controller.ApproveConstruction)     // 审批指定施工作业
	routerGroup.POST("/receipts/:id", controller.ReceiveConstruction)     // 领取施工作业
	routerGroup.POST("/terminations/:id", controller.StopConstruction)    // 终止施工作业
	routerGroup.POST("/savings/:id", controller.SubmitConstruction)       // 提交施工作业
	routerGroup.POST("replaies/:id", controller.SubmitConstructionReplay) // 提交施工作业复盘
	routerGroup.POST("/sounds/:id", controller.SubmitConstructionSound)   // 提交施工作业录音
}
