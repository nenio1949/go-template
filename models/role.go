package models

import (
	"go-template/common"

	"gorm.io/gorm"
)

// 角色model
type Role struct {
	Model
	Name       string `json:"name" gorm:"not null;unique;comment:名称"`
	Permission string `json:"permission" gorm:"default:'';comment:权限"`
}

// 获取角色列表
func GetRoles(params common.PageSearchRoleDto) ([]*Role, int64, error) {
	var roles []*Role
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
		err = tx.Offset(params.Page - 1).Limit(params.Size).Find(&roles).Error
	} else {
		err = tx.Find(&roles).Error
	}

	var total int64
	tx.Count(&total)

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0, err
	}

	return roles, total, nil
}

// 根据id获取角色信息
func GetRole(id int) (*Role, error) {
	var role Role
	err := db.Where("id = ? AND deleted = ? ", id, 0).First(&role).Error
	if err != nil {
		return nil, err
	}

	return &role, nil
}

// 新增角色
func AddRole(params common.RoleCreateDto) (int, error) {
	role := Role{
		Name:       params.Name,
		Permission: params.Permission,
	}

	if err := db.Create(&role).Error; err != nil {
		return 0, err
	}

	return role.ID, nil
}

// 更新指定角色
func UpdateRole(id int, params common.RoleUpdateDto) (bool, error) {
	if hasRole, hasErr := GetRole(id); hasErr != nil {
		_ = hasRole
		return false, hasErr
	}

	role := Role{
		Name:       params.Name,
		Permission: params.Permission,
	}
	if r := db.Model(&Role{}).Where("id = ? AND deleted = ? ", id, 0).Updates(role); r.RowsAffected != 1 {
		return false, r.Error
	}

	return true, nil
}

// 删除角色
func DeleteRoles(ids []int) (int, error) {
	if err := db.Model(&Role{}).Where("id IN (?)", ids).Update("deleted", 1).Error; err != nil {
		return 0, err
	}

	return len(ids), nil
}
