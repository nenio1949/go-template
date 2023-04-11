package models

import (
	"go-template/common"

	"gorm.io/gorm"
)

// 安全措施库model
type MeasureLibrary struct {
	Model
	HomeWork string `json:"home_work" gorm:"comment:作业环节"`
	RiskType string `json:"risk_type" gorm:"comment:风险类型"`
	Name     string `json:"name" gorm:"comment:作业类型"`
	Risk     string `json:"risk" gorm:"comment:潜在风险"`
	Measures string `json:"measures" gorm:"type:text;comment:安全措施"`
}

// 根据ids获取措施库列表
func GetMeasureLibrarysByIds(ids []int) ([]MeasureLibrary, error) {
	var measureLibraries []MeasureLibrary
	var err error
	tx := db.Where("deleted = 0")
	tx.Where("construction_id IN (?)", ids)
	tx.Order("id DESC")

	tx.Find(&measureLibraries)

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return measureLibraries, nil
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
