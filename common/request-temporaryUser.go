package common

type TemporaryUserDto struct {
	ID            int    `form:"id" json:"id"`
	Name          string `form:"name" json:"name"`
	Mobile        string `form:"mobile" json:"mobile"`
	Department    string `form:"department" json:"department"`
	DockingUserID int    `form:"docking_user_id" json:"docking_user_id"`
}

type TemporaryUserCreateDto struct {
	Name          string `form:"name" json:"name"`
	Mobile        string `form:"mobile" json:"mobile"`
	Department    string `form:"department" json:"department"`
	DockingUserID int    `form:"docking_user_id" json:"docking_user_id"`
}

type TemporaryUserUpdateDto struct {
	ID            int    `form:"id" json:"id"`
	Name          string `form:"name" json:"name"`
	Mobile        string `form:"mobile" json:"mobile"`
	Department    string `form:"department" json:"department"`
	DockingUserID int    `form:"docking_user_id" json:"docking_user_id"`
}
