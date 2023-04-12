package controller

import (
	"go-template/common"
	service "go-template/services"
	"go-template/utils"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// 获取施工作业列表
func GetConstructionPlans(c *gin.Context) {
	var form common.PageSearchConstructionDto

	form.Page = 1
	form.Size = 10
	form.Pagination = true

	if err := c.ShouldBindQuery(&form); err != nil {
		common.ValidateFail(c, common.GetErrorMsg(form, err))
		return
	}

	constructions, total, err := service.GetConstructionPlans(form)
	if err != nil {
		common.BusinessFail(c, err.Error())
		return
	}
	common.Success(c, map[string]interface{}{"constructions": constructions, "total": total})
}

// 根据id获取指定施工作业
func GetConstructionPlan(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		common.BusinessFail(c, err.Error())
		return
	}
	construction, err := service.GetConstructionPlan(id)
	if err != nil {
		common.BusinessFail(c, err.Error())
		return
	}
	common.Success(c, construction)
}

// 新增施工作业计划
func AddConstructionPlan(c *gin.Context) {
	var form common.ConstructionPlanCreateDto

	if err := c.ShouldBindJSON(&form); err != nil {
		common.ValidateFail(c, common.GetErrorMsg(form, err))
		return
	}

	id, err := service.AddConstructionPlan(form)
	if err != nil {
		common.BusinessFail(c, err.Error())
		return
	}
	common.Success(c, id)
}

// 更新指定施工作业计划
func UpdateConstructionPlan(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		common.BusinessFail(c, err.Error())
		return
	}

	var form common.ConstructionPlanUpdateDto
	if err := c.ShouldBindJSON(&form); err != nil {
		common.ValidateFail(c, common.GetErrorMsg(form, err))
		return
	}

	number, err := service.UpdateConstructionPlan(id, form)
	if err != nil {
		common.BusinessFail(c, err.Error())
		return
	}
	common.Success(c, number)
}

// 删除施工作业
func DeleteConstructions(c *gin.Context) {
	ids := c.Param("ids")
	if len(ids) == 0 {
		common.BusinessFail(c, "参数非法")
	}

	idsArr := utils.StringToInt(strings.Split(ids, ","))

	number, err := service.DeleteConstructions(idsArr)
	if err != nil {
		common.BusinessFail(c, err.Error())
		return
	}
	common.Success(c, number)
}
