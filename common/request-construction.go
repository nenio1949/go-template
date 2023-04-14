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

// 施工作业提交dto
type ConstructionSubmitDto struct {
	Files []FileDto `form:"files" json:"files" binding:"required"`
	// 是否确认宣读
	IsNotice bool `form:"is_notice" json:"is_notice" binding:"required"`
	// 宣读时间
	NoticeTime string `form:"notice_time" json:"notice_time" binding:"required"`
	// 停留时长(单位秒)
	StayTime int `form:"stay_time" json:"stay_time" binding:"required"`
	// 工具数量
	ToolNum int `form:"tool_num" json:"tool_num" binding:"required"`
	// 人员数量
	UserNum int `form:"user_num" json:"user_num" binding:"required"`
	// 打卡时间
	ClockTime string `form:"clock_time" json:"clock_time"`
	// 打卡人员id
	ClockUserID int `form:"clock_user_id" json:"clock_user_id"`
	// 打卡地点
	ClockLocation string `form:"clock_location" json:"clock_location"`
	// 工具和人员清点备注
	ToolRemark string `form:"tool_remark" json:"tool_remark,omitempty"`
	// 红闪灯备注
	LightRemark string `form:"light_remark" json:"light_remark,omitempty"`
	// 红闪灯类型
	LightType string `form:"light_type" json:"light_type"`
	// 安全防护员备注
	GuardRemark string `form:"guard_remark" json:"guard_remark,omitempty"`
	// 安全防护员类型
	GuardType string `form:"guard_type" json:"guard_type"`
	// 作业过程备注
	ProcessRemark string `form:"process_remark" json:"process_remark,omitempty"`
	// 出清工具数量
	QuitToolNum int `form:"quit_tool_num" json:"quit_tool_num"`
	// 出清人员数量
	QuitUserNum int `form:"quit_user_num" json:"quit_user_num"`
	// 出清工具备注
	QuitToolRemark string `form:"quit_tool_remark" json:"quit_tool_remark,omitempty"`
	// 出清人员备注
	QuitUserRemark string `form:"quit_user_remark" json:"quit_user_remark,omitempty"`
	// 出清打卡时间
	QuitClockTime string `form:"quit_clock_time" json:"quit_clock_time"`
	// 出清打卡地点
	QuitClockLocation string `form:"quit_clock_location" json:"quit_clock_location"`
	// 出清打卡人员id
	QuitClockUserID int `form:"quit_clock_user_id" json:"quit_clock_user_id"`
	// 交接内容
	Handover string `form:"handover" json:"handover,omitempty"`
	// 交接时间
	HandoverTime string `form:"handover_time" json:"handover_time"`
	// 交接类型
	HandoverType string `form:"handover_type" json:"handover_type"`
	// 作业登记类型
	WorkedType string `form:"worker_type" json:"worker_type"`
	// 作业登记备注
	WorkedRemark string `form:"worker_remark" json:"worker_remark,omitempty"`
	// 施工注销类型
	LogoutType string `form:"logout_type" json:"logout_type"`
	// 施工注销备注
	LogoutRemark string `form:"logout_remark" json:"logout_remark,omitempty"`
}

// 审批施工作业dto
type ConstructionApproveDto struct {
	ApproveStatus string `form:"approve_status" json:"approve_status"`
	ApproveRemark string `form:"approve_remark" json:"approve_remark,omitempty"`
}

// 终止施工作业dto
type ConstructionStopDto struct {
	StopReason string `form:"stop_reason" json:"stop_reason"`
}

// 提交复盘内容dto
type ConstructionSubmitReplayDto struct {
	ReplayContext string    `form:"replay_context" json:"replay_context"`
	Files         []FileDto `form:"files" json:"files"`
}

// 提交录音dto
type ConstructionSubmitSoundDto struct {
	Files       []FileDto `form:"files" json:"files"`
	SoundRemark string    `form:"sound_remark" json:"sound_remark"`
}
