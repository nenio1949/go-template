package models

import (
	"go-template/common"

	"gorm.io/gorm"
)

// 安全措施库model
type MeasureLibrary struct {
	Model
	HomeWork string      `json:"home_work" gorm:"comment:作业环节"`
	RiskType string      `json:"risk_type" gorm:"comment:风险类型"`
	Name     string      `json:"name" gorm:"comment:作业类型"`
	Risk     string      `json:"risk" gorm:"comment:潜在风险"`
	Measures common.Strs `json:"measures" gorm:"type:text;comment:安全措施"`
}

// 获取措施库列表
func GetMeasureLibraries(params common.PageSearchMeasureLibraryDto) ([]*MeasureLibrary, int64, error) {
	var measureLibrarys []*MeasureLibrary
	var err error
	tx := db.Where("deleted = 0")
	if len(params.Name) > 0 {
		tx.Where("name like ?", params.Name)
	}

	if len(params.Order) > 0 {
		tx.Order(params.Order)
	} else {
		tx.Order("id DESC")
	}

	if params.Pagination {
		err = tx.Offset(params.Page - 1).Limit(params.Size).Find(&measureLibrarys).Error
	} else {
		err = tx.Find(&measureLibrarys).Error
	}

	var total int64
	tx.Count(&total)

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0, err
	}

	return measureLibrarys, total, nil
}

// 根据ids获取措施库列表
func GetMeasureLibrariesByIds(ids []int) ([]MeasureLibrary, error) {
	var measureLibraries []MeasureLibrary
	var err error
	tx := db.Where("deleted = 0")
	tx.Where("id IN (?)", ids)
	tx.Order("id DESC")

	tx.Find(&measureLibraries)

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return measureLibraries, nil
}

// 根据id获取措施库信息
func GetMeasureLibrary(id int) (*MeasureLibrary, error) {
	var measureLibrary MeasureLibrary
	err := db.Where("id = ? AND deleted = ? ", id, 0).First(&measureLibrary).Error
	if err != nil {
		return nil, err
	}

	return &measureLibrary, nil
}

// 新增措施库
func AddMeasureLibrary(params common.MeasureLibraryCreateDto) (int, error) {
	measureLibrary := MeasureLibrary{
		HomeWork: params.HomeWork,
		RiskType: params.RiskType,
		Name:     params.Name,
		Risk:     params.Risk,
		Measures: params.Measures,
	}

	if err := db.Create(&measureLibrary).Error; err != nil {
		return 0, err
	}

	return measureLibrary.ID, nil
}

// 更新指定措施库
func UpdateMeasureLibrary(id int, params common.MeasureLibraryUpdateDto) (bool, error) {
	var oldMeasureLibrary *MeasureLibrary
	var err error
	if oldMeasureLibrary, err = GetMeasureLibrary(id); err != nil || oldMeasureLibrary == nil {
		return false, err
	}

	oldMeasureLibrary.Name = params.Name
	oldMeasureLibrary.HomeWork = params.HomeWork
	oldMeasureLibrary.Risk = params.Risk
	oldMeasureLibrary.RiskType = params.RiskType
	oldMeasureLibrary.Measures = params.Measures
	if r := db.Updates(&oldMeasureLibrary); r.RowsAffected != 1 {
		return false, r.Error
	}

	return true, nil
}

// 删除措施库
func DeleteMeasureLibraries(ids []int) (int, error) {
	if err := db.Model(&MeasureLibrary{}).Where("id IN (?)", ids).Update("deleted", 1).Error; err != nil {
		return 0, err
	}

	return len(ids), nil
}
