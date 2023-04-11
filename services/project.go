package service

import (
	"go-template/common"
	"go-template/models"
)

// 获取项目列表
func GetProjects(params common.PageSearchProjectDto) ([]*models.Project, int64, error) {
	return models.GetProjects(params)
}

// 根据id获取项目信息
func GetProject(id int) (*models.Project, error) {
	return models.GetProject(id)
}

// 新增项目
func AddProject(params common.ProjectCreateDto) (int, error) {
	return models.AddProject(params)
}

// 更新指定项目
func UpdateProject(id int, params common.ProjectUpdateDto) (bool, error) {
	return models.UpdateProject(id, params)
}

// 删除项目
func DeleteProjects(ids []int) (int, error) {
	return models.DeleteProjects(ids)
}
