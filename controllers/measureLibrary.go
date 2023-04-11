package controller

import (
	"go-template/common"
	service "go-template/services"

	"github.com/gin-gonic/gin"
)

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
