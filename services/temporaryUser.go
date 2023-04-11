package service

import (
	"go-template/common"
	"go-template/models"
)

// 根据ids获取临时人员列表
func GetTemporaryUsersByIds(ids []int) ([]models.TemporaryUser, error) {
	return models.GetTemporaryUsersByIds(ids)
}

// 批量新增或更新临时人员
func AddOrUpdateTemporaryUsers(constructionId int, list []common.TemporaryUserDto) ([]int, error) {
	return models.AddOrUpdateTemporaryUsers(constructionId, list)
}
