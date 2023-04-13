package controller

import (
	"go-template/common"
	service "go-template/services"
	"strconv"
	"strings"

	"go-template/utils"

	"github.com/gin-gonic/gin"
)

// 获取项目列表
func GetProjects(c *gin.Context) {
	var form common.PageSearchProjectDto

	form.Page = 1
	form.Size = 10
	form.Pagination = true

	if err := c.ShouldBindQuery(&form); err != nil {
		common.ValidateFail(c, common.GetErrorMsg(form, err))
		return
	}

	projects, total, err := service.GetProjects(form)
	if err != nil {
		common.BusinessFail(c, err.Error())
		return
	}
	common.Success(c, map[string]interface{}{"projects": projects, "total": total})
}

// 根据id获取指定项目
func GetProject(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		common.BusinessFail(c, err.Error())
		return
	}
	project, err := service.GetProject(id)
	if err != nil {
		common.BusinessFail(c, err.Error())
		return
	}
	common.Success(c, project)
}

// 新增项目
func AddProject(c *gin.Context) {
	var form common.ProjectCreateDto

	if err := c.ShouldBindJSON(&form); err != nil {
		common.ValidateFail(c, common.GetErrorMsg(form, err))
		return
	}

	id, err := service.AddProject(form)
	if err != nil {
		common.BusinessFail(c, err.Error())
		return
	}
	common.Success(c, id)
}

// 更新指定项目
func UpdateProject(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		common.BusinessFail(c, err.Error())
		return
	}

	var form common.ProjectUpdateDto
	if err := c.ShouldBindJSON(&form); err != nil {
		common.ValidateFail(c, common.GetErrorMsg(form, err))
		return
	}

	success, err := service.UpdateProject(id, form)
	if err != nil {
		common.BusinessFail(c, err.Error())
		return
	}
	common.Success(c, success)
}

// 删除项目
func DeleteProjects(c *gin.Context) {
	ids := c.Query("ids")
	if ids == "" {
		common.BusinessFail(c, "参数非法")
	}

	idsArr := utils.StringToInt(strings.Split(ids, ","))

	number, err := service.DeleteProjects(idsArr)
	if err != nil {
		common.BusinessFail(c, err.Error())
		return
	}
	common.Success(c, number)
}
