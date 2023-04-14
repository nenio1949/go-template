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
func AddConstructionPlan(params common.ConstructionPlanCreateDto, currentUser *models.User) (int, error) {
	return models.AddConstructionPlan(params, *currentUser)
}

// 更新指定施工作业计划
func UpdateConstructionPlan(id int, params common.ConstructionPlanUpdateDto) (bool, error) {
	return models.UpdateConstructionPlan(id, params)
}

// 删除施工作业
func DeleteConstructions(ids []int) (int, error) {
	return models.DeleteConstructions(ids)
}

// 获取指定施工作业信息
func GetConstruction(id int) (*models.Construction, error) {
	return models.GetConstruction(id)
}

// 更新施工作业
func UpdateConstruction(id int, params common.ConstructionUpdateDto, currentUser *models.User) (bool, error) {
	return models.UpdateConstruction(id, params, *currentUser)
}

// 审批指定施工作业
func ApproveConstruction(id int, params common.ConstructionApproveDto, currentUser *models.User) (bool, error) {
	return models.ApproveConstruction(id, params, *currentUser)
}

// 领取施工作业(移动端)
func ReceiveConstruction(id int, currentUser *models.User) (bool, error) {
	return models.ReceiveConstruction(id, *currentUser)
}

// 终止施工作业
func StopConstruction(id int, params common.ConstructionStopDto, currentUser *models.User) (bool, error) {
	return models.StopConstruction(id, params, *currentUser)
}

// 提交施工作业(移动端)
func SubmitConstruction(id int, params common.ConstructionSubmitDto) (bool, error) {
	return models.SubmitConstruction(id, params)
}

// 提交施工作业复盘
func SubmitConstructionReplay(id int, params common.ConstructionSubmitReplayDto) (bool, error) {
	return models.SubmitConstructionReplay(id, params)
}

// 提交施工作业录音
func SubmitConstructionSound(id int, params common.ConstructionSubmitSoundDto) (bool, error) {
	return models.SubmitConstructionSound(id, params)
}
