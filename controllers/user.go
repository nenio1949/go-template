package controller

import (
	"fmt"
	"go-template/common"
	service "go-template/services"
	"strconv"
	"strings"

	"go-template/utils"

	"github.com/gin-gonic/gin"
)

// 获取用户列表
func GetUsers(c *gin.Context) {
	var form common.PageSearchUserDto
	form.Page = 1
	form.Size = 10
	form.Pagination = true

	if err := c.ShouldBindQuery(&form); err != nil {
		common.ValidateFail(c, common.GetErrorMsg(form, err))
		return
	}

	users, total, err := service.GetUsers(form)
	if err != nil {
		common.BusinessFail(c, err.Error())
		return
	}
	common.Success(c, map[string]interface{}{"users": users, "total": total})
}

// 根据id获取指定用户
func GetUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		common.BusinessFail(c, err.Error())
		return
	}
	user, err := service.GetUser(id)
	if err != nil {
		common.BusinessFail(c, err.Error())
		return
	}
	common.Success(c, user)
}

// 新增用户
func AddUser(c *gin.Context) {
	var form common.UserCreateDto

	if err := c.ShouldBindJSON(&form); err != nil {
		common.ValidateFail(c, common.GetErrorMsg(form, err))
		return
	}

	id, err := service.AddUser(form)
	if err != nil {
		common.BusinessFail(c, err.Error())
		return
	}
	common.Success(c, id)
}

// 更新指定用户
func UpdateUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		common.BusinessFail(c, err.Error())
		return
	}

	var form common.UserUpdateDto
	if err := c.ShouldBindJSON(&form); err != nil {
		common.ValidateFail(c, common.GetErrorMsg(form, err))
		return
	}
	fmt.Printf("sssddd=%+v\n", form)

	number, err := service.UpdateUser(id, form)
	if err != nil {
		common.BusinessFail(c, err.Error())
		return
	}
	common.Success(c, number)
}

// 删除用户
func DeleteUsers(c *gin.Context) {
	ids := c.Query("ids")
	if ids == "" {
		common.BusinessFail(c, "参数非法")
	}

	idsArr := utils.StringToInt(strings.Split(ids, ","))

	number, err := service.DeleteUsers(idsArr)
	if err != nil {
		common.BusinessFail(c, err.Error())
		return
	}
	common.Success(c, number)
}
