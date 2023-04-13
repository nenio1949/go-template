package controller

import (
	"go-template/common"
	service "go-template/services"
	"go-template/utils"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// 获取措施库列表
func GetMeasureLibraries(c *gin.Context) {
	var form common.PageSearchMeasureLibraryDto

	form.Page = 1
	form.Size = 10
	form.Pagination = true

	if err := c.ShouldBindQuery(&form); err != nil {
		common.ValidateFail(c, common.GetErrorMsg(form, err))
		return
	}

	measureLibraries, total, err := service.GetMeasureLibraries(form)
	if err != nil {
		common.BusinessFail(c, err.Error())
		return
	}
	common.Success(c, map[string]interface{}{"measureLibraries": measureLibraries, "total": total})
}

// 根据id获取指定措施库
func GetMeasureLibrary(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		common.BusinessFail(c, err.Error())
		return
	}
	measureLibrary, err := service.GetMeasureLibrary(id)
	if err != nil {
		common.BusinessFail(c, err.Error())
		return
	}
	common.Success(c, measureLibrary)
}

// 新增措施库
func AddMeasureLibrary(c *gin.Context) {
	var form common.MeasureLibraryCreateDto

	if err := c.ShouldBindJSON(&form); err != nil {
		common.ValidateFail(c, common.GetErrorMsg(form, err))
		return
	}

	id, err := service.AddMeasureLibrary(form)
	if err != nil {
		common.BusinessFail(c, err.Error())
		return
	}
	common.Success(c, id)
}

// 更新指定项目
func UpdateMeasureLibrary(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		common.BusinessFail(c, err.Error())
		return
	}

	var form common.MeasureLibraryUpdateDto
	if err := c.ShouldBindJSON(&form); err != nil {
		common.ValidateFail(c, common.GetErrorMsg(form, err))
		return
	}

	success, err := service.UpdateMeasureLibrary(id, form)
	if err != nil {
		common.BusinessFail(c, err.Error())
		return
	}
	common.Success(c, success)
}

// 删除措施库
func DeleteMeasureLibraries(c *gin.Context) {
	ids := c.Param("ids")
	if len(ids) == 0 {
		common.BusinessFail(c, "参数非法")
	}

	idsArr := utils.StringToInt(strings.Split(ids, ","))

	number, err := service.DeleteMeasureLibraries(idsArr)
	if err != nil {
		common.BusinessFail(c, err.Error())
		return
	}
	common.Success(c, number)
}
