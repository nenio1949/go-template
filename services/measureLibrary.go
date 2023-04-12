package service

import (
	"go-template/common"
	"go-template/models"
)

// 获取措施库列表
func GetMeasureLibraries(params common.PageSearchMeasureLibraryDto) ([]*models.MeasureLibrary, int64, error) {
	return models.GetMeasureLibraries(params)
}

// 根据id获取措施库信息
func GetMeasureLibrary(id int) (*models.MeasureLibrary, error) {
	return models.GetMeasureLibrary(id)
}

// 新增措施库
func AddMeasureLibrary(params common.MeasureLibraryCreateDto) (int, error) {
	return models.AddMeasureLibrary(params)
}

// 更新指定措施库
func UpdateMeasureLibrary(id int, params common.MeasureLibraryUpdateDto) (bool, error) {
	return models.UpdateMeasureLibrary(id, params)
}

// 删除措施库
func DeleteMeasureLibraries(ids []int) (int, error) {
	return models.DeleteMeasureLibraries(ids)
}
