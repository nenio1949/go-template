package service

import (
	"go-template/common"
	"go-template/models"
)

// 获取施工作业列表
func GetConstructionPlans(params common.PageSearchConstructionDto) ([]*common.ConstructionPlanDto, int64, error) {
	return models.GetConstructionPlans(params)
}

// 根据id获取施工作业计划信息
func GetConstructionPlan(id int) (*common.ConstructionPlanDto, error) {
	return models.GetConstructionPlan(id)
}

// 新增施工作业计划
func AddConstructionPlan(params common.ConstructionPlanCreateDto) (int, error) {
	return models.AddConstructionPlan(params)
}

// 更新指定施工作业计划
func UpdateConstructionPlan(id int, params common.ConstructionPlanUpdateDto) (bool, error) {
	return models.UpdateConstructionPlan(id, params)
}

// 删除施工作业
func DeleteConstructions(ids []int) (int, error) {
	return models.DeleteConstructions(ids)
}
