package models

import (
	"errors"
	"go-template/common"

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

// 批量新增或更新临时人员
func AddOrUpdateTemporaryUsers(constructionId int, list []common.TemporaryUserDto) ([]int, error) {
	var ids []int
	for a := 0; a < len(list); a++ {
		if list[a].ID > 0 { // 更新
			user := TemporaryUser{
				Name:          list[a].Name,
				Mobile:        list[a].Mobile,
				Department:    list[a].Department,
				DockingUserID: list[a].DockingUserID,
			}
			if r := db.Model(&TemporaryUser{}).Where("id = ? AND deleted = 0", list[a].ID).Updates(user); r.RowsAffected == 1 {
				ids = append(ids, list[a].ID)
			}
		} else { // 新增
			user := TemporaryUser{
				Name:          list[a].Name,
				Mobile:        list[a].Mobile,
				Department:    list[a].Department,
				DockingUserID: list[a].DockingUserID,
			}
			if err := db.Create(&user).Error; err == nil {
				ids = append(ids, user.ID)
			}
		}
	}
	// 删除
	db.Where("construction_id = ? AND id NOT IN (?)", constructionId, ids).Delete(TemporaryUser{})
	return ids, nil
}
