package models

import (
	"gorm.io/gorm"
)

// 日志model
type Log struct {
	Model

	Content string `json:"content" gorm:"not null;comment:日志内容"`
	User    User   `json:"user"`
	UserID  int    `json:"user_id" gorm:"comment:用户id;==default:'galeone'=="`
}

// 获取日志列表
func GetLogs(constructionId int) ([]*Log, int64, error) {
	var roles []*Log
	var err error
	tx := db.Where("deleted = 0")
	if constructionId > 0 {
		tx.Where("contruction_id = ?", constructionId)
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

// 根据id获取日志信息
func GetLog(id int) (*Log, error) {
	var role Log
	err := db.Where("id = ? AND deleted = ? ", id, 0).First(&role).Error
	if err != nil {
		return nil, err
	}

	return &role, nil
}

// 新增日志
func AddLog(content string, user User) (int, error) {
	role := Log{
		Content: content,
		UserID:  user.ID,
	}

	if err := db.Create(&role).Error; err != nil {
		return 0, err
	}

	return role.ID, nil
}
