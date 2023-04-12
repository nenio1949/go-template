package models

import (
	"errors"
	"go-template/common"

	"gorm.io/gorm"
)

// 临时人员model
type TemporaryUser struct {
	Model
	Name          string `json:"name" gorm:"comment:姓名"`
	Mobile        string `json:"mobile" gorm:"comment:手机号"`
	Department    string `json:"department" gorm:"comment:所属部门"`
	DockingUser   User   `json:"user"`
	DockingUserID int    `json:"docking_user_id" gorm:"comment:对接人员id"`
}

// 获取临时人员列表
func GetTemporaryUsers(params common.PageSearchTemporaryUserDto) ([]*TemporaryUser, int64, error) {
	var temporaryUsers []*TemporaryUser
	var err error
	tx := db.Where("deleted = 0")
	if len(params.Name) > 0 {
		tx.Where("name like ?", params.Name)
	}
	if len(params.Mobile) > 0 {
		tx.Where("mobile like ?", params.Mobile)
	}

	tx.Preload("DockingUser")

	if len(params.Order) > 0 {
		tx.Order(params.Order)
	} else {
		tx.Order("id DESC")
	}

	if params.Pagination {
		err = tx.Offset(params.Page - 1).Limit(params.Size).Find(&temporaryUsers).Error
	} else {
		err = tx.Find(&temporaryUsers).Error
	}

	var total int64
	tx.Count(&total)

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0, err
	}

	return temporaryUsers, total, nil
}

// 根据ids获取用户列表
func GetTemporaryUsersByIds(ids []int) ([]TemporaryUser, error) {
	if len(ids) == 0 {
		return nil, errors.New("参数非法")
	}
	var users []TemporaryUser

	err := db.Preload("DockingUser").Where("deleted = 0 AND id IN (?)", ids).Find(&users).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return users, nil
}
