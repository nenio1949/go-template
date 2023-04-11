package controller

import (
	"go-template/common"
	service "go-template/services"

	"github.com/gin-gonic/gin"
)

// 新增施工作业
func AddConstruction(c *gin.Context) {
	var form common.ConstructionCreateDto

	if err := c.ShouldBindJSON(&form); err != nil {
		common.ValidateFail(c, common.GetErrorMsg(form, err))
		return
	}

	id, err := service.AddConstruction(form)
	if err != nil {
		common.BusinessFail(c, err.Error())
		return
	}
	common.Success(c, id)
}
