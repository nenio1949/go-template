package service

import (
	"go-template/common"
	"go-template/models"
)

// 获取角色列表
func GetRoles(params common.PageSearchRoleDto) ([]*models.Role, int64, error) {
	return models.GetRoles(params)
}

// 根据id获取角色信息
func GetRole(id int) (*models.Role, error) {
	return models.GetRole(id)
}

// 新增角色
func AddRole(params common.RoleCreateDto) (int, error) {
	return models.AddRole(params)
}

// 更新指定角色
func UpdateRole(id int, params common.RoleUpdateDto) (bool, error) {
	return models.UpdateRole(id, params)
}

// 删除角色
func DeleteRoles(ids []int) (int, error) {
	return models.DeleteRoles(ids)
}
