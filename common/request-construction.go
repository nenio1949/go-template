package common

// 施工作业查询dto
type PageSearchConstructionDto struct {
	PaginationDto
	Status string `form:"status" json:"status,omitempty"`
}

// 施工作业计划dto
type ConstructionPlanDto struct {
	ID               int                      `json:"id"`
	MeasureLibraries []MeasureLibraryDto      `json:"measure_libraries"`
	StartTtime       LocalTime                `json:"start_time"`
	EndTime          LocalTime                `json:"end_time"`
	Leader           map[string]interface{}   `json:"leader"`
	Location         string                   `json:"location"`
	Remark           string                   `json:"remark"`
	Status           map[string]interface{}   `json:"status"`
	EquipmentType    []string                 `json:"equipment_type"`
	ExecutiveUsers   []map[string]interface{} `json:"executive_users"`
}

// 施工作业计划新增dto
type ConstructionPlanCreateDto struct {
	MeasureLibraryIds []int    `form:"measure_library_ids" json:"measure_library_ids" binding:"required"`
	StartTtime        string   `form:"start_time" json:"start_time" binding:"required"`
	EndTime           string   `form:"end_time" json:"end_time" binding:"required"`
	ContinuedDays     int      `form:"continued_days" json:"continued_days" binding:"required"`
	LeaderID          int      `form:"leader_id" json:"leader_id" binding:"required"`
	ExecutiveUserIds  []int    `form:"executive_user_ids" json:"executive_user_ids" binding:"required"`
	EquipmentType     []string `form:"equipment_type" json:"equipment_type" binding:"required"`
	Location          string   `form:"location" json:"location" binding:"required"`
	Remark            string   `form:"remark" json:"remark"`
}

func (params ConstructionPlanCreateDto) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"measure_library_ids.required": "作业类型不能为空",
		"start_time.required":          "开始时间不能为空",
		"end_time.required":            "结束时间不能为空",
		"continued_days.required":      "持续天数不能为空",
		"leader_id.required":           "作业组长不能为空",
		"executive_user_ids.required":  "执行人员不能为空",
		"equipment_type.required":      "设备类型不能为空",
		"location.required":            "作业地点不能为空",
	}
}

// 施工作业计划更新dto
type ConstructionPlanUpdateDto struct {
	MeasureLibraryIds []int    `form:"measure_library_ids" json:"measure_library_ids,omitempty"`
	StartTtime        string   `form:"start_time" json:"start_time,omitempty"`
	EndTime           string   `form:"end_time" json:"end_time,omitempty"`
	ContinuedDays     int      `form:"continued_days" json:"continued_days,omitempty"`
	LeaderID          int      `form:"leader_id" json:"leader_id,omitempty"`
	ExecutiveUserIds  []int    `form:"executive_user_ids" json:"executive_user_ids,omitempty"`
	EquipmentType     []string `form:"equipment_type" json:"equipment_type,omitempty"`
	Location          string   `form:"location" json:"location,omitempty"`
	Remark            string   `form:"remark" json:"remark,omitempty"`
}

// 施工作业更新dto
type ConstructionUpdateDto struct {
	StartTtime       string   `form:"start_time" json:"start_time,omitempty"`
	EndTime          string   `form:"end_time" json:"end_time,omitempty"`
	EquipmentType    []string `form:"equipment_type" json:"equipment_type,omitempty"`
	Location         string   `form:"location" json:"location,omitempty"`
	ExecutiveUserIds []int    `form:"executive_user_ids" json:"executive_user_ids,omitempty"`
	Remark           string   `form:"remark" json:"remark,omitempty"`
	// 施工任务相关
	ManagerID  int `form:"manager_id" json:"manager_id,omitempty"`
	EngineerID int `form:"engineer_id" json:"engineer_id,omitempty"`
	LeaderID   int `form:"leader_id" json:"leader_id,omitempty"`

	Process        string             `form:"process" json:"process,omitempty"`
	WorkScope      string             `form:"work_scope" json:"work_scope,omitempty"`
	Restrictions   string             `form:"restrictions" json:"restrictions,omitempty"`
	Matter         string             `form:"matter" json:"matter,omitempty"`
	IsRisk         bool               `form:"is_risk" json:"is_risk,omitempty"`
	TemporaryUsers []TemporaryUserDto `form:"temporary_users" json:"temporary_users,omitempty"`

	// 是否提交
	IsSubmit bool `form:"is_submit" json:"is_submit,omitempty"`
}
