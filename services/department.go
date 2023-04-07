package service

import (
	"go-template/common"
	"go-template/models"
)

// 获取部门列表
func GetDepartments(params common.PageSearchDepartmentDto) ([]*models.Department, int64, error) {
	return models.GetDepartments(params)
}

// 根据id获取部门信息
func GetDepartment(id int) (*models.Department, error) {
	return models.GetDepartment(id)
}

// 新增部门
func AddDepartment(params common.DepartmentCreateDto) (int, error) {
	return models.AddDepartment(params)
}

// 更新指定部门
func UpdateDepartment(id int, params common.DepartmentUpdateDto) (bool, error) {
	return models.UpdateDepartment(id, params)
}

// 删除部门
func DeleteDepartments(ids []int) (int, error) {
	return models.DeleteDepartments(ids)
}
