package models

import (
	"gorm.io/gorm"
)

// 文件model
type File struct {
	Model

	// 注销作业令,quit_tool:作业出清工具,quit_tool_user:作业出清人员,guard:防护员,light:红闪灯,register: 作业登记,tool_user:人员清点,tool: 工具清点,user: 安全交底作业人员,user_other: 安全交底外来人员,replay: 每日复盘,sound: 录音(安全交底),sound_center: 录音(中心),sound_car: 录音(车载),order:作业令,mobileSound:移动端录音
	Type    string `json:"type" gorm:"not null;comment:cancel:文件类型"`
	Content string `json:"content" gorm:"not null;comment:文件内容"`
	User    User   `json:"user"`
	UserID  int    `json:"user_id" gorm:"comment:用户id;==default:'galeone'=="`
}

// 获取文件列表
func GetFiles(constructionId int) ([]*File, int64, error) {
	var roles []*File
	var err error
	tx := db.Where("deleted = 0")
	if constructionId > 0 {
		tx.Where("construction_id = ?", constructionId)
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

// 根据id获取文件信息
func GetFile(id int) (*File, error) {
	var role File
	err := db.Where("id = ? AND deleted = ? ", id, 0).First(&role).Error
	if err != nil {
		return nil, err
	}

	return &role, nil
}

// 新增文件
func AddFile(content string, user User) (int, error) {
	role := File{
		Content: content,
		UserID:  user.ID,
	}

	if err := db.Create(&role).Error; err != nil {
		return 0, err
	}

	return role.ID, nil
}
