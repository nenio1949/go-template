package common

// 角色创建dto
type RoleCreateDto struct {
	Name       string `form:"name" json:"name" binding:"required"`
	Permission string `form:"permission" json:"permission"`
}

// 角色更新dto
type RoleUpdateDto struct {
	Name       string `form:"name" json:"name"`
	Permission string `form:"permission" json:"permission"`
}

// 角色分页查询dto
type PageSearchRoleDto struct {
	PaginationDto
	Name string `form:"name" json:"name"`
}

func (params RoleCreateDto) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"name.required": "角色名称不能为空",
	}
}
