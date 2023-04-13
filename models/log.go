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
	var logs []*Log
	var err error
	tx := db.Where("deleted = 0")
	if constructionId > 0 {
		tx.Where("construction_id = ?", constructionId)
	}
	tx.Order("id DESC")

	tx.Find(&logs)

	var total int64
	tx.Count(&total)

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0, err
	}

	return logs, total, nil
}

// 根据id获取日志信息
func GetLog(id int) (*Log, error) {
	var log Log
	err := db.Where("id = ? AND deleted = ? ", id, 0).First(&log).Error
	if err != nil {
		return nil, err
	}

	return &log, nil
}

// 新增日志
func AddLog(content string, user User) (int, error) {
	log := Log{
		Content: content,
		UserID:  user.ID,
	}

	if err := db.Create(&log).Error; err != nil {
		return 0, err
	}

	return log.ID, nil
}
