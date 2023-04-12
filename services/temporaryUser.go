package service

import (
	"go-template/models"
)

// 根据ids获取临时人员列表
func GetTemporaryUsersByIds(ids []int) ([]models.TemporaryUser, error) {
	return models.GetTemporaryUsersByIds(ids)
}
