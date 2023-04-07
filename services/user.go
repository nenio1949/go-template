package service

import (
	"fmt"
	"go-template/common"
	"go-template/models"
)

// 获取用户列表
func GetUsers(params common.PageSearchUserDto) ([]*models.User, int64, error) {
	return models.GetUsers(params)
}

// 根据id获取用户信息
func GetUser(id int) (*models.User, error) {
	return models.GetUser(id)
}

// 新增用户
func AddUser(params common.UserCreateDto) (int, error) {
	fmt.Printf("sssddd=%+v\n", params)

	return models.AddUser(params)
}

// 更新指定用户
func UpdateUser(id int, params common.UserUpdateDto) (bool, error) {
	return models.UpdateUser(id, params)
}

// 删除用户
func DeleteUsers(ids []int) (int, error) {
	return models.DeleteUsers(ids)
}
