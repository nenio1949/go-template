package models

// 临时人员model
type TemporaryUser struct {
	Model
	Name          string `json:"name" gorm:"comment:姓名"`
	Mobile        string `json:"mobile" gorm:"comment:手机号"`
	Department    string `json:"department" gorm:"comment:所属部门"`
	User          User   `json:"user"`
	DockingUserID int    `json:"docking_user_id" gorm:"comment:对接人员id"`
}
