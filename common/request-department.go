package common

// 部门创建dto
type DepartmentCreateDto struct {
	Name     string `form:"name" json:"name" binding:"required"`
	ParentID int    `form:"parent_id" json:"parent_id"`
}

// 部门更新dto
type DepartmentUpdateDto struct {
	Name     string `form:"name" json:"name"`
	ParentID int    `form:"parent_id" json:"parent_id"`
}

// 部门分页查询dto
type PageSearchDepartmentDto struct {
	PaginationDto
	Name string `form:"name" json:"name"`
}

func (params DepartmentCreateDto) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"name.required": "部门名称不能为空",
	}
}
