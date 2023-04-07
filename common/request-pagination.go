package common

// 分页dto
type PaginationDto struct {
	// 当前页码
	Page int `form:"page" json:"page" default:"1"`
	// 每页数据量(默认10条/页)
	Size int `form:"size" json:"size" default:"10"`
	// 是否分页(默认分页)
	Pagination bool `form:"pagination" json:"pagination" default:"true"`
	// 排序
	Order string `form:"order" json:"order"`
}
