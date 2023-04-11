package models

import (
	"errors"
	"go-template/common"

	"gorm.io/gorm"
)

// 项目model
type Project struct {
	Model
	Name   string `json:"content" gorm:"not null;comment:项目名称"`
	Number string `json:"number" gorm:"not null;comment:项目代号"`
	PinYin string `json:"pin_yin" gorm:"not null;comment:拼音"`
	Region string `json:"region" gorm:"not null;comment:所属区域"`
	Active bool   `json:"active" gorm:"not null;default:true;comment:是否激活"`
}

// 获取项目列表
func GetProjects(params common.PageSearchProjectDto) ([]*Project, int64, error) {
	var projects []*Project
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
		err = tx.Offset(params.Page - 1).Limit(params.Size).Find(&projects).Error
	} else {
		err = tx.Find(&projects).Error
	}

	var total int64
	tx.Count(&total)

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0, err
	}

	return projects, total, nil
}

// 根据ids获取项目列表
func GetProjectsByIds(ids []int) ([]Project, error) {
	if len(ids) == 0 {
		return nil, errors.New("参数非法")
	}
	var projects []Project

	err := db.Where("deleted = 0 AND id IN (?)", ids).Find(&projects).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return projects, nil
}

// 根据id获取项目信息
func GetProject(id int) (*Project, error) {
	var project Project
	err := db.Where("id = ? AND deleted = ? ", id, 0).First(&project).Error
	if err != nil {
		return nil, err
	}

	return &project, nil
}

// 新增项目
func AddProject(params common.ProjectCreateDto) (int, error) {
	project := Project{
		Name:   params.Name,
		Number: params.Number,
		PinYin: params.PinYin,
		Region: params.Region,
	}

	if err := db.Create(&project).Error; err != nil {
		return 0, err
	}

	return project.ID, nil
}

// 更新指定项目
func UpdateProject(id int, params common.ProjectUpdateDto) (bool, error) {
	if hasProject, hasErr := GetProject(id); hasErr != nil {
		_ = hasProject
		return false, hasErr
	}

	project := Project{
		Name:   params.Name,
		Number: params.Number,
		PinYin: params.PinYin,
		Region: params.Region,
		Active: params.Active,
	}
	if r := db.Model(&Project{}).Where("id = ? AND deleted = ? ", id, 0).Updates(project); r.RowsAffected != 1 {
		return false, r.Error
	}

	return true, nil
}

// 删除项目
func DeleteProjects(ids []int) (int, error) {
	if err := db.Model(&Project{}).Where("id IN (?)", ids).Update("deleted", 1).Error; err != nil {
		return 0, err
	}

	return len(ids), nil
}
