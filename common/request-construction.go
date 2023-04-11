package common

// 施工作业新增dto
type ConstructionCreateDto struct {
	MeasureIds       []int  `form:"measure_ids" json:"measure_ids"`
	StartTtime       string `form:"start_time" json:"start_time"`
	EndTime          string `form:"end_time" json:"end_time"`
	ContinuedDays    int    `form:"continued_days" json:"continued_days"`
	LeaderID         int    `form:"leader_id" json:"leader_id"`
	ExecutiveUserIds []int  `form:"executive_user_ids" json:"executive_user_ids"`
	EquipmentType    string `form:"equipment_type" json:"equipment_type"`
	Location         string `form:"location" json:"location"`
	Remark           string `form:"remark" json:"remark"`
}

// 施工作业更新dto
type ConstructionUpdateDto struct {
	MeasureIds       []int  `form:"measure_id" json:"measure_id,omitempty"`
	StartTtime       string `form:"start_time" json:"start_time,omitempty"`
	EndTime          string `form:"end_time" json:"end_time,omitempty"`
	ActualTime       string `form:"actual_time" json:"actual_time,omitempty"`
	EquipmentType    string `form:"equipment_type" json:"equipment_type,omitempty"`
	Location         string `form:"location" json:"location,omitempty"`
	Status           string `form:"status" json:"status,omitempty"`
	ExecutiveUserIds []int  `form:"executive_user_ids" json:"executive_user_ids,omitempty"`
	Remark           string `form:"remark" json:"remark,omitempty"`
	// 施工任务相关
	ManagerID         int                `form:"manager_id" json:"manager_id,omitempty"`
	EngineerID        int                `form:"engineer_id" json:"engineer_id,omitempty"`
	LeaderID          int                `form:"leader_id" json:"leader_id,omitempty"`
	RecipientID       int                `form:"recipient_id" json:"recipient_id,omitempty"`
	PhoneID           string             `form:"phone_id" json:"phone_id,omitempty"`
	TerminationUserID int                `form:"termination_user_id" json:"termination_user_id,omitempty"`
	StopReason        string             `form:"stop_reason" json:"stop_reason,omitempty"`
	Process           string             `form:"process" json:"process,omitempty"`
	Content           string             `form:"content" json:"content,omitempty"`
	WorkScope         string             `form:"work_scope" json:"work_scope,omitempty"`
	Restrictions      string             `form:"restrictions" json:"restrictions,omitempty"`
	Matter            string             `form:"matter" json:"matter,omitempty"`
	IsRisk            bool               `form:"is_risk" json:"is_risk,omitempty"`
	TemporaryUser     []TemporaryUserDto `form:"temporary_user" json:"temporary_user,omitempty"`

	// 安全交底
	Explain    string `form:"explain" json:"explain,omitempty"`
	IsNotice   bool   `form:"is_notice" json:"is_notice,omitempty"`
	NoticeTime string `form:"notice_time" json:"notice_time,omitempty"`
	StayTime   int    `form:"stay_time" json:"stay_time,omitempty"`

	// 工具人员清点
	ToolNum     int    `form:"tool_num" json:"tool_num,omitempty"`
	UserNum     int    `form:"user_num" json:"user_num,omitempty"`
	ClockTime   string `form:"clock_time" json:"clock_time,omitempty"`
	ClockUserID int    `form:"clock_user_id" json:"clock_user_id,omitempty"`
	ToolRemark  string `form:"tool_remark" json:"tool_remark,omitempty"`

	// 作业边界
	LightRemark   string `form:"light_remark" json:"light_remark,omitempty"`
	LightType     string `form:"light_type" json:"light_type,omitempty"`
	GuardRemark   string `form:"guard_remark" json:"guard_remark,omitempty"`
	GuardType     string `form:"guard_type" json:"guard_type,omitempty"`
	NeedJob       bool   `form:"need_job" json:"need_job,omitempty"`
	ProcessRemark string `form:"process_remark" json:"process_remark,omitempty"`

	// 作业出清
	QuitToolNum       int    `form:"quit_tool_num" json:"quit_tool_num,omitempty"`
	QuitUserNum       int    `form:"quit_user_num" json:"quit_user_num,omitempty"`
	QuitToolRemark    string `form:"quit_tool_remark" json:"quit_tool_remark,omitempty"`
	QuitUserRemark    string `form:"quit_user_remark" json:"quit_user_remark,omitempty"`
	QuitClockTime     string `form:"quit_clock_time" json:"quit_clock_time,omitempty"`
	QuitClockLocation string `form:"quit_clock_location" json:"quit_clock_location,omitempty"`
	QuitClockUserID   int    `form:"quit_clock_user_id" json:"quit_clock_user_id,omitempty"`

	// 作业交接
	HaveHandover bool   `form:"have_handover" json:"have_handover,omitempty"`
	Handover     string `form:"handover" json:"handover,omitempty"`
	HandoverTime string `form:"handover_time" json:"handover_time,omitempty"`
	HandoverType string `form:"handover_type" json:"handover_type,omitempty"`

	// 每日复盘
	ReplayContext  string `form:"replay_context" json:"replay_context,omitempty"`
	ReplayTime     string `form:"replay_time" json:"replay_time,omitempty"`
	SoundRemark    string `form:"sound_remark" json:"sound_remark,omitempty"`
	MobileReceived bool   `form:"mobile_received" json:"mobile_received,omitempty"`
	WorkedType     string `form:"worker_type" json:"worker_type,omitempty"`
	WorkedRemark   string `form:"worker_remark" json:"worker_remark,omitempty"`
	LogoutJob      bool   `form:"logout_job" json:"logout_job,omitempty"`
	LogoutType     string `form:"logout_type" json:"logout_type,omitempty"`
	LogoutRemark   string `form:"logout_remark" json:"logout_remark,omitempty"`

	AuditStatus string `form:"audit_status" json:"audit_status,omitempty"`
}
