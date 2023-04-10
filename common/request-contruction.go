package common

type TemporaryUserDto struct {
	ID            int    `form:"id" json:"id"`
	Name          string `form:"name" json:"name"`
	Mobile        string `form:"mobile" json:"mobile"`
	Department    string `form:"department" json:"department"`
	DockingUserID int    `form:"docking_user_id" json:"docking_user_id"`
}

// 施工作业新增dto
type ContructionCreateDto struct {
	MeasureIds       []int  `form:"measure_id" json:"measure_id"`
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
type ContructionUpdateDto struct {
	StartTtime      string `form:"start_time" json:"start_time"`
	EndTime         string `form:"end_time" json:"end_time"`
	ActualTime      string `form:"actual_time" json:"actual_time"`
	EquipmentType   string `form:"equipment_type" json:"equipment_type"`
	Location        string `form:"location" json:"location"`
	Status          string `form:"status" json:"status"`
	ExecutiveUserID int    `form:"executive_user_id" json:"executive_user_id"`
	Remark          string `form:"remark" json:"remark"`
	// 施工任务相关
	ManagerID         int                `form:"manager_id" json:"manager_id"`
	EngineerID        int                `form:"engineer_id" json:"engineer_id"`
	LeaderID          int                `form:"leader_id" json:"leader_id"`
	RecipientID       int                `form:"recipient_id" json:"recipient_id"`
	PhoneID           string             `form:"phone_id" json:"phone_id"`
	TerminationUserID int                `form:"termination_user_id" json:"termination_user_id"`
	StopReason        string             `form:"stop_reason" json:"stop_reason"`
	Process           string             `form:"process" json:"process"`
	Content           string             `form:"content" json:"content"`
	WorkScope         string             `form:"work_scope" json:"work_scope"`
	Restrictions      string             `form:"restrictions" json:"restrictions"`
	Matter            string             `form:"matter" json:"matter"`
	IsRisk            bool               `form:"is_risk" json:"is_risk"`
	TemporaryUser     []TemporaryUserDto `form:"temporary_user" json:"temporary_user"`
	TemporaryUserIDs  []int              `form:"temporary_user_id" json:"temporary_user_id"`

	// 安全交底
	Explain    string `form:"explain" json:"explain"`
	IsNotice   bool   `form:"is_notice" json:"is_notice"`
	NoticeTime string `form:"notice_time" json:"notice_time"`
	StayTime   string `form:"stay_time" json:"stay_time"`

	// 工具人员清点
	ToolNum     int    `form:"tool_num" json:"tool_num"`
	UserNum     int    `form:"user_num" json:"user_num"`
	ClockTime   string `form:"clock_time" json:"clock_time"`
	ClockUserID int    `form:"clock_user_id" json:"clock_user_id"`
	ToolRemark  string `form:"tool_remark" json:"tool_remark"`

	// 作业边界
	LightRemark   string `form:"light_remark" json:"light_remark"`
	LightType     string `form:"light_type" json:"light_type"`
	GuardRemark   string `form:"guard_remark" json:"guard_remark"`
	GuardType     string `form:"guard_type" json:"guard_type"`
	NeedJob       bool   `form:"need_job" json:"need_job"`
	ProcessRemark string `form:"process_remark" json:"process_remark"`

	// 作业出清
	QuitToolNum       int    `form:"quit_tool_num" json:"quit_tool_num"`
	QuitUserNum       int    `form:"quit_user_num" json:"quit_user_num"`
	QuitToolRemark    string `form:"quit_tool_remark" json:"quit_tool_remark"`
	QuitUserRemark    string `form:"quit_user_remark" json:"quit_user_remark"`
	QuitClockTime     string `form:"quit_clock_time" json:"quit_clock_time"`
	QuitClockLocation string `form:"quit_clock_location" json:"quit_clock_location"`
	QuitClockUserID   int    `form:"quit_clock_user_id" json:"quit_clock_user_id"`

	// 作业交接
	HaveHandover bool   `form:"have_handover" json:"have_handover"`
	Handover     string `form:"handover" json:"handover"`
	HandoverTime string `form:"handover_time" json:"handover_time"`
	HandoverType string `form:"handover_type" json:"handover_type"`

	// 每日复盘
	ReplayContext  string `form:"replay_context" json:"replay_context"`
	ReplayTime     string `form:"replay_time" json:"replay_time"`
	SoundRemark    string `form:"sound_remark" json:"sound_remark"`
	MobileReceived bool   `form:"mobile_received" json:"mobile_received"`
	WorkedType     string `form:"worker_type" json:"worker_type"`
	WorkedRemark   string `form:"worker_remark" json:"worker_remark"`
	LogoutJob      bool   `form:"logout_job" json:"logout_job"`
	LogoutType     string `form:"logout_type" json:"logout_type"`
	LogoutRemark   string `form:"logout_remark" json:"logout_remark"`

	AuditStatus string `form:"audit_status" json:"audit_status"`
}
