package controller

import (
	"go-template/common"
	service "go-template/services"
	"strconv"
	"strings"

	"go-template/utils"

	"github.com/gin-gonic/gin"
)

// 获取部门列表
func GetDepartments(c *gin.Context) {
	var form common.PageSearchDepartmentDto

	departments, total, err := service.GetDepartments(form)
	if err != nil {
		common.BusinessFail(c, err.Error())
		return
	}
	common.Success(c, map[string]interface{}{"departments": departments, "total": total})
}

// 根据id获取指定部门
func GetDepartment(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		common.BusinessFail(c, err.Error())
		return
	}
	department, err := service.GetDepartment(id)
	if err != nil {
		common.BusinessFail(c, err.Error())
		return
	}
	common.Success(c, department)
}

// 新增部门
func AddDepartment(c *gin.Context) {
	var form common.DepartmentCreateDto

	if err := c.ShouldBindJSON(&form); err != nil {
		common.ValidateFail(c, common.GetErrorMsg(form, err))
		return
	}

	id, err := service.AddDepartment(form)
	if err != nil {
		common.BusinessFail(c, err.Error())
		return
	}
	common.Success(c, id)
}

// 更新指定部门
func UpdateDepartment(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		common.BusinessFail(c, err.Error())
		return
	}

	var form common.DepartmentUpdateDto
	if err := c.ShouldBindJSON(&form); err != nil {
		common.ValidateFail(c, common.GetErrorMsg(form, err))
		return
	}

	number, err := service.UpdateDepartment(id, form)
	if err != nil {
		common.BusinessFail(c, err.Error())
		return
	}
	common.Success(c, number)
}

// 删除部门
func DeleteDepartments(c *gin.Context) {
	ids := c.Query("ids")
	if ids == "" {
		common.BusinessFail(c, "参数非法")
	}

	idsArr := utils.StringToInt(strings.Split(ids, ","))

	number, err := service.DeleteDepartments(idsArr)
	if err != nil {
		common.BusinessFail(c, err.Error())
		return
	}
	common.Success(c, number)
}
