package common

// 项目创建dto
type ProjectCreateDto struct {
	Name   string `form:"name" json:"name" binding:"required"`
	Number string `form:"number" json:"number"`
	PinYin string `form:"pin-yin" json:"pin-y"`
	Region string `form:"region" json:"region"`
	Active bool   `form:"active" json:"active"`
}

// 项目更新dto
type ProjectUpdateDto struct {
	Name   string `form:"name" json:"name"`
	Number string `form:"number" json:"number"`
	PinYin string `form:"pin-yin" json:"pin-y"`
	Region string `form:"region" json:"region"`
	Active bool   `form:"active" json:"active"`
}

// 项目分页查询dto
type PageSearchProjectDto struct {
	PaginationDto
	Name string `form:"name" json:"name"`
}

func (params ProjectCreateDto) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"name.required": "项目名称不能为空",
	}
}
