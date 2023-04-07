package service

import (
	"go-template/common"
	"go-template/global"
	"go-template/models"
	"go-template/utils"
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

// 登录
func Login(params common.UserLoginDto) (*models.User, TokenOutPut, error) {
	user, err := models.GetUserByLogin(params.Account, utils.MD5(params.Password))
	if err != nil || user.ID <= 0 {
		return nil, TokenOutPut{}, err
	}

	tokenData, err := CreateToken(global.AppGuardName, user)
	if err != nil {
		return nil, tokenData, err
	}
	return user, tokenData, nil
}
