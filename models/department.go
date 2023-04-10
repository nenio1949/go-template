package models

import (
	"go-template/common"

	"gorm.io/gorm"
)

// 部门model
type Department struct {
	Model
	Name     string `json:"name" gorm:"not null;unique;comment:名称"`
	ParentID int    `json:"parent_id" gorm:"not null;default:0;comment:父级部门ID"`
}

// 获取部门列表
func GetDepartments(params common.PageSearchDepartmentDto) ([]*Department, int64, error) {
	var departments []*Department
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
		err = tx.Offset(params.Page - 1).Limit(params.Size).Find(&departments).Error
	} else {
		err = tx.Find(&departments).Error
	}

	var total int64
	tx.Count(&total)

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0, err
	}

	return departments, total, nil
}

// 根据id获取部门信息
func GetDepartment(id int) (*Department, error) {
	var department Department
	err := db.Where("id = ? AND deleted = ? ", id, 0).First(&department).Error
	if err != nil {
		return nil, err
	}

	return &department, nil
}

// 新增部门
func AddDepartment(params common.DepartmentCreateDto) (int, error) {
	department := Department{
		Name:     params.Name,
		ParentID: params.ParentID,
	}

	if err := db.Create(&department).Error; err != nil {
		return 0, err
	}

	return department.Model.ID, nil
}

// 更新指定部门
func UpdateDepartment(id int, params common.DepartmentUpdateDto) (bool, error) {
	if hasDepartment, hasErr := GetDepartment(id); hasErr != nil {
		_ = hasDepartment
		return false, hasErr
	}

	department := Department{
		Name:     params.Name,
		ParentID: params.ParentID,
	}
	if r := db.Model(&Department{}).Where("id = ? AND deleted = ? ", id, 0).Updates(department); r.RowsAffected != 1 {
		return false, r.Error
	}

	return true, nil
}

// 删除部门
func DeleteDepartments(ids []int) (int, error) {
	r := db.Model(&Department{}).Where("id IN (?)", ids).Updates(map[string]interface{}{"deleted": 1})
	if r.Error != nil {
		return 0, r.Error
	}

	return int(r.RowsAffected), nil
}
