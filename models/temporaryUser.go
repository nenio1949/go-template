package models

import (
	"errors"

	"gorm.io/gorm"
)

// 临时人员model
type TemporaryUser struct {
	Model
	Name           string `json:"name" gorm:"comment:姓名"`
	Mobile         string `json:"mobile" gorm:"comment:手机号"`
	Department     string `json:"department" gorm:"comment:所属部门"`
	DockingUser    User   `json:"user"`
	DockingUserID  int    `json:"docking_user_id" gorm:"comment:对接人员id"`
	ConstructionId int    `json:"construction_id" gorm:"comment:施工作业id"`
}

// 根据ids获取用户列表
func GetTemporaryUsersByIds(ids []int) ([]TemporaryUser, error) {
	if len(ids) == 0 {
		return nil, errors.New("参数非法")
	}
	var users []TemporaryUser

	err := db.Where("deleted = 0 AND id IN (?)", ids).Find(&users).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return users, nil
}
