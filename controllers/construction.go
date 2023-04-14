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

	constructionPlans, total, err := service.GetConstructionPlans(form)
	if err != nil {
		common.BusinessFail(c, err.Error())
		return
	}
	common.Success(c, map[string]interface{}{"constructions": constructionPlans, "total": total})
}

// 根据id获取指定施工作业计划
func GetConstructionPlan(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		common.BusinessFail(c, err.Error())
		return
	}
	constructionPlan, err := service.GetConstructionPlan(id)
	if err != nil {
		common.BusinessFail(c, err.Error())
		return
	}
	common.Success(c, constructionPlan)
}

// 新增施工作业计划
func AddConstructionPlan(c *gin.Context) {
	var form common.ConstructionPlanCreateDto

	if err := c.ShouldBindJSON(&form); err != nil {
		common.ValidateFail(c, common.GetErrorMsg(form, err))
		return
	}
	currentUser, _ := service.GetUserInfoByRequest(c)
	id, err := service.AddConstructionPlan(form, currentUser)
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

	success, err := service.UpdateConstructionPlan(id, form)
	if err != nil {
		common.BusinessFail(c, err.Error())
		return
	}
	common.Success(c, success)
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

// 根据id获取指定施工作业
func GetConstruction(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		common.BusinessFail(c, err.Error())
		return
	}
	construction, err := service.GetConstruction(id)
	if err != nil {
		common.BusinessFail(c, err.Error())
		return
	}
	common.Success(c, construction)
}

// 更新指定施工作业
func UpdateConstruction(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		common.BusinessFail(c, err.Error())
		return
	}

	var form common.ConstructionUpdateDto
	if err := c.ShouldBindJSON(&form); err != nil {
		common.ValidateFail(c, common.GetErrorMsg(form, err))
		return
	}

	currentUser, _ := service.GetUserInfoByRequest(c)
	success, err := service.UpdateConstruction(id, form, currentUser)
	if err != nil {
		common.BusinessFail(c, err.Error())
		return
	}
	common.Success(c, success)
}

// 审批指定施工作业
func ApproveConstruction(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		common.BusinessFail(c, err.Error())
		return
	}

	var form common.ConstructionApproveDto
	if err := c.ShouldBindJSON(&form); err != nil {
		common.ValidateFail(c, common.GetErrorMsg(form, err))
		return
	}

	currentUser, _ := service.GetUserInfoByRequest(c)
	success, err := service.ApproveConstruction(id, form, currentUser)
	if err != nil {
		common.BusinessFail(c, err.Error())
		return
	}
	common.Success(c, success)
}

// 领取施工作业
func ReceiveConstruction(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		common.BusinessFail(c, err.Error())
		return
	}

	var form common.ConstructionApproveDto
	if err := c.ShouldBindJSON(&form); err != nil {
		common.ValidateFail(c, common.GetErrorMsg(form, err))
		return
	}

	currentUser, _ := service.GetUserInfoByRequest(c)
	success, err := service.ReceiveConstruction(id, currentUser)
	if err != nil {
		common.BusinessFail(c, err.Error())
		return
	}
	common.Success(c, success)
}

// 终止施工作业
func StopConstruction(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		common.BusinessFail(c, err.Error())
		return
	}
	var form common.ConstructionStopDto
	if err := c.ShouldBindJSON(&form); err != nil {
		common.ValidateFail(c, common.GetErrorMsg(form, err))
		return
	}

	currentUser, _ := service.GetUserInfoByRequest(c)
	success, err := service.StopConstruction(id, form, currentUser)
	if err != nil {
		common.BusinessFail(c, err.Error())
		return
	}
	common.Success(c, success)
}

// 提交施工作业
func SubmitConstruction(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		common.BusinessFail(c, err.Error())
		return
	}
	var form common.ConstructionSubmitDto
	if err := c.ShouldBindJSON(&form); err != nil {
		common.ValidateFail(c, common.GetErrorMsg(form, err))
		return
	}

	success, err := service.SubmitConstruction(id, form)
	if err != nil {
		common.BusinessFail(c, err.Error())
		return
	}
	common.Success(c, success)
}

// 提交施工作业复盘
func SubmitConstructionReplay(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		common.BusinessFail(c, err.Error())
		return
	}
	var form common.ConstructionSubmitReplayDto
	if err := c.ShouldBindJSON(&form); err != nil {
		common.ValidateFail(c, common.GetErrorMsg(form, err))
		return
	}

	success, err := service.SubmitConstructionReplay(id, form)
	if err != nil {
		common.BusinessFail(c, err.Error())
		return
	}
	common.Success(c, success)
}

// 提交施工作业录音文件
func SubmitConstructionSound(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		common.BusinessFail(c, err.Error())
		return
	}
	var form common.ConstructionSubmitSoundDto
	if err := c.ShouldBindJSON(&form); err != nil {
		common.ValidateFail(c, common.GetErrorMsg(form, err))
		return
	}

	success, err := service.SubmitConstructionSound(id, form)
	if err != nil {
		common.BusinessFail(c, err.Error())
		return
	}
	common.Success(c, success)
}
