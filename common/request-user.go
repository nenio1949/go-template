package common

// 用户分页查询dto
type PageSearchUserDto struct {
	PaginationDto
	Name         string `form:"name" json:"name"`
	Account      string `form:"account" json:"account"`
	NickName     string `form:"nick_name" json:"nick_name"`
	Mobile       string `form:"mobile" json:"mobile"`
	Status       string `form:"status" json:"status"`
	RoleID       int    `form:"role_id" json:"role_id"`
	DepartmentID int    `form:"department_id" json:"department_id"`
}

// 用户创建dto
type UserCreateDto struct {
	Name         string `form:"name" json:"name" binding:"required"`
	Account      string `form:"account" json:"account" binding:"required"`
	Password     string `form:"password" json:"password" binding:"required"`
	NickName     string `form:"nick_name" json:"nick_name"`
	Gender       string `form:"gender" json:"gender" gorm:"default:'unknow'"`
	Mobile       string `form:"mobile" json:"mobile"  binding:"required"`
	Email        string `form:"email" josn:"email"`
	RoleID       int    `form:"role_id" json:"role_id" validate:"gt=0"`
	DepartmentID int    `form:"department_id" json:"department_id" validate:"gt=0"`
}

func (params UserCreateDto) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"name.required":     "用户名称不能为空",
		"mobile.required":   "手机号码不能为空",
		"mobile.mobile":     "手机号码格式不正确",
		"password.required": "用户密码不能为空",
		"account.required":  "登录账号不能为空",
	}
}

// 用户更新dto
type UserUpdateDto struct {
	Name         string `form:"name" json:"name,omitempty"`
	Account      string `form:"account" json:"account,omitempty" `
	Password     string `form:"password" json:"password,omitempty" `
	NickName     string `form:"nick_name" json:"nick_name,omitempty" `
	Gender       string `form:"gender" json:"gender,omitempty" `
	Mobile       string `form:"mobile" json:"mobile,omitempty" `
	Email        string `form:"email" json:"email,omitempty" `
	Status       string `form:"status" json:"status,omitempty" `
	RoleID       int    `form:"role_id" json:"role_id,omitempty" `
	DepartmentID int    `form:"department_id" json:"department_id,omitempty" `
}

// 用户登录dto
type UserLoginDto struct {
	Account  string `form:"account" json:"account" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}
