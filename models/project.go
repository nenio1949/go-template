package models

import (
	"go-template/common"

	"gorm.io/gorm"
)

// 项目model
type Project struct {
	Model
	Name   string `json:"content" gorm:"not null;comment:项目内容"`
	Number string `json:"number" gorm:"not null;comment:项目内容"`
	PinYin string `json:"pin_yin" gorm:"not null;comment:拼音"`
	Region string `json:"region" gorm:"not null;comment:所属区域"`
	Active bool   `json:"active" gorm:"not null;default:true;comment:是否激活"`
}

// 获取项目列表
func GetProjects(params common.PageSearchProjectDto) ([]*Project, int64, error) {
	var roles []*Project
	var err error
	tx := db.Where("deleted = 0")
	if len(params.Name) > 0 {
		tx.Where("contruction_id like ?", params.Name)
	}
	tx.Order("id DESC")

	tx.Find(&roles)

	var total int64
	tx.Count(&total)

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0, err
	}

	return roles, total, nil
}

// 根据id获取项目信息
func GetProject(id int) (*Project, error) {
	var role Project
	err := db.Where("id = ? AND deleted = ? ", id, 0).First(&role).Error
	if err != nil {
		return nil, err
	}

	return &role, nil
}

// 新增项目
func AddProject(params common.ProjectCreateDto) (int, error) {
	role := Project{
		Name:   params.Name,
		Number: params.Number,
		PinYin: params.PinYin,
		Region: params.Region,
		Active: params.Active,
	}

	if err := db.Create(&role).Error; err != nil {
		return 0, err
	}

	return role.ID, nil
}
