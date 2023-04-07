package common

// 角色创建dto
type RoleCreateDto struct {
	Name       string `form:"name" binding:"required"`
	Permission string `form:"permission"`
}

// 角色更新dto
type RoleUpdateDto struct {
	ID         int    `json:"id"`
	Name       string `form:"name"`
	Permission string `form:"permission"`
}

// 角色分页查询dto
type PageSearchRoleDto struct {
	PaginationDto
	Name string `form:"name"`
}

func (params RoleCreateDto) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"name.required": "角色名称不能为空",
	}
}
