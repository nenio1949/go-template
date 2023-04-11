package models

import (
	"fmt"
	"go-template/common"
	"time"

	"gorm.io/gorm"
)

// 施工作业model
type Construction struct {
	Model
	MeasureLibrarys []MeasureLibrary `json:"measure_libraries" gorm:"many2many:construction_measure;"`
	StartTime       LocalTime        `json:"start_time" gorm:"comment:开始时间"`
	EndTime         LocalTime        `json:"end_time" gorm:"comment:结束时间"`
	ActualTime      LocalTime        `json:"actual_time" gorm:"comment:实际完成时间"`
	EquipmentType   string           `json:"equipment_type" gorm:"comment:设备类型"`
	Location        string           `json:"location" gorm:"comment:作业地点"`
	Status          string           `json:"status" gorm:"comment:状态"`
	ExecutiveUsers  []User           `json:"executive_users" gorm:"many2many:construction_user;"`
	Remark          string           `json:"remark" gorm:"comment:备注"`
	// 施工任务相关
	ManagerID         int             `json:"manager_id" gorm:"comment:项目经理id"`
	EngineerID        int             `json:"engineer_id" gorm:"comment:项目做过id"`
	LeaderID          int             `json:"leader_id" gorm:"comment:组长id"`
	RecipientID       int             `json:"recipient_id" gorm:"comment:领取人id"`
	PhoneID           string          `json:"phone_id" gorm:"comment:领取设备id"`
	TerminationUserID int             `json:"termination_user_id" gorm:"comment:终止人id"`
	StopReason        string          `json:"stop_reason" gorm:"comment:终止原因"`
	Process           string          `json:"process" gorm:"type:text;comment:作业流程"`
	Content           string          `json:"content" gorm:"type:text;comment:工作内容"`
	WorkScope         string          `json:"work_scope" gorm:"type:text;comment:工作范围"`
	Restrictions      string          `json:"restrictions" gorm:"type:text;comment:安全限制条件"`
	Matter            string          `json:"matter" gorm:"type:text;comment:注意事项"`
	IsRisk            bool            `json:"is_risk" gorm:"comment:是否设计高风险"`
	TemporaryUsers    []TemporaryUser `json:"temporary_users" gorm:"foreignKey:ConstructionId"`

	// 安全交底
	Explain    string    `json:"explain" gorm:"comment:补充说明"`
	IsNotice   bool      `json:"is_notice" gorm:"comment:是否确认宣读"`
	NoticeTime LocalTime `json:"notice_time" gorm:"comment:宣读时间"`
	StayTime   int       `json:"stay_time" gorm:"comment:停留时间(单位秒)"`

	// 工具人员清点
	ToolNum     int       `json:"tool_num" gorm:"comment:工具数量"`
	UserNum     int       `json:"user_num" gorm:"comment:人员数量"`
	ClockTime   LocalTime `json:"clock_time" gorm:"comment:打卡时间"`
	ClockUserID int       `json:"clock_user_id" gorm:"comment:打卡人id"`
	ToolRemark  string    `json:"tool_remark" gorm:"type:text;comment:工具清点备注"`

	// 作业边界
	LightRemark   string `json:"light_remark" gorm:"type:text;comment:红闪灯备注"`
	LightType     string `json:"light_type" gorm:"comment:红闪灯类型"`
	GuardRemark   string `json:"guard_remark" gorm:"type:text;comment:防护员备注"`
	GuardType     string `json:"guard_type" gorm:"comment:防护员类型"`
	NeedJob       bool   `json:"need_job" gorm:"comment:是否作业边界"`
	ProcessRemark string `json:"process_remark" gorm:"type:text;comment:作业过程备注"`

	// 作业出清
	QuitToolNum       int       `json:"quit_tool_num" gorm:"comment:出清工具数量"`
	QuitUserNum       int       `json:"quit_user_num" gorm:"comment:出清人员数量"`
	QuitToolRemark    string    `json:"quit_tool_remark" gorm:"type:text;comment:工具备注"`
	QuitUserRemark    string    `json:"quit_user_remark" gorm:"type:text;comment:人员备注"`
	QuitClockTime     LocalTime `json:"quit_clock_time" gorm:"comment:出清打卡时间"`
	QuitClockLocation string    `json:"quit_clock_location" gorm:"comment:出清打卡地点"`
	QuitClockUserID   int       `json:"quit_clock_user_id" gorm:"comment:出清打卡人id"`

	// 作业交接
	HaveHandover bool      `json:"have_handover" gorm:"comment:是否交接"`
	Handover     string    `json:"handover" gorm:"type:text;comment:交接内容"`
	HandoverTime LocalTime `json:"handover_time" gorm:"comment:交接时间"`
	HandoverType string    `json:"handover_type" gorm:"comment:交接类型"`

	// 每日复盘
	ReplayContext  string    `json:"replay_context" gorm:"type:text;comment:复盘内容"`
	ReplayTime     LocalTime `json:"replay_time" gorm:"comment:复盘时间"`
	SoundRemark    string    `json:"sound_remark" gorm:"type:text;comment:录音备注"`
	MobileReceived bool      `json:"mobile_received" gorm:"comment:手机是否已领取"`
	WorkedType     string    `json:"worker_type" gorm:"comment:作业登记类型"`
	WorkedRemark   string    `json:"worker_remark" gorm:"type:text;comment:作业登记备注"`
	LogoutJob      bool      `json:"logout_job" gorm:"comment:是否注销作业令"`
	LogoutType     string    `json:"logout_type" gorm:"comment:注销类型"`
	LogoutRemark   string    `json:"logout_remark" gorm:"type:text;comment:作业注销备注"`

	AuditStatus string `json:"audit_status" gorm:"comment:审计状态"`
}

// 获取施工作业列表
func GetConstructions(constructionId int) ([]*Construction, int64, error) {
	var constructions []*Construction
	var err error
	tx := db.Where("deleted = 0")
	if constructionId > 0 {
		tx.Where("construction_id = ?", constructionId)
	}
	tx.Order("id DESC")

	tx.Find(&constructions)

	var total int64
	tx.Count(&total)

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0, err
	}

	return constructions, total, nil
}

// 根据id获取施工作业信息
func GetConstruction(id int) (*Construction, error) {
	var construction Construction
	err := db.Where("id = ? AND deleted = ? ", id, 0).First(&construction).Error
	if err != nil {
		return nil, err
	}

	return &construction, nil
}

// 新增施工作业
func AddConstruction(params common.ConstructionCreateDto) (int, error) {
	startTime, _ := time.ParseInLocation("2006-01-02 15:04:05", params.StartTtime, time.Local)
	endTime, _ := time.ParseInLocation("2006-01-02 15:04:05", params.EndTime, time.Local)
	measureLibraries, _ := GetMeasureLibrarysByIds(params.MeasureIds)
	executiveUsers, _ := GetUsersByIds(params.ExecutiveUserIds)

	construction := Construction{
		StartTime:       LocalTime{Time: startTime},
		EndTime:         LocalTime{Time: endTime},
		MeasureLibrarys: measureLibraries,
		LeaderID:        params.LeaderID,
		ExecutiveUsers:  executiveUsers,
		EquipmentType:   params.EquipmentType,
		Location:        params.Location,
		Remark:          params.Remark,
	}

	fmt.Printf("construction: %v\n", construction)

	if err := db.Create(&construction).Error; err != nil {
		return 0, err
	}

	return construction.ID, nil
}

// 更新施工作业
func UpdateConstruction(id int, params common.ConstructionUpdateDto) (bool, error) {
	startTime, _ := time.ParseInLocation("20060102150405", params.StartTtime, time.Local)
	endTime, _ := time.ParseInLocation("20060102150405", params.EndTime, time.Local)
	measureLibraries, _ := GetMeasureLibrarysByIds(params.MeasureIds)
	executiveUsers, _ := GetUsersByIds(params.ExecutiveUserIds)
	noticeTime, _ := time.ParseInLocation("20060102150405", params.NoticeTime, time.Local)
	clockTime, _ := time.ParseInLocation("20060102150405", params.ClockTime, time.Local)
	quitClockTime, _ := time.ParseInLocation("20060102150405", params.QuitClockTime, time.Local)
	handoverTime, _ := time.ParseInLocation("20060102150405", params.HandoverTime, time.Local)
	replayTime, _ := time.ParseInLocation("20060102150405", params.ReplayTime, time.Local)
	ids, _ := AddOrUpdateTemporaryUsers(id, params.TemporaryUser)
	temporaryUsers, _ := GetTemporaryUsersByIds(ids)

	construction := Construction{
		StartTime:         LocalTime{Time: startTime},
		EndTime:           LocalTime{Time: endTime},
		MeasureLibrarys:   measureLibraries,
		ExecutiveUsers:    executiveUsers,
		EquipmentType:     params.EquipmentType,
		Location:          params.Location,
		Remark:            params.Remark,
		ManagerID:         params.ManagerID,
		EngineerID:        params.EngineerID,
		LeaderID:          params.LeaderID,
		RecipientID:       params.RecipientID,
		PhoneID:           params.PhoneID,
		TerminationUserID: params.TerminationUserID,
		StopReason:        params.StopReason,
		Process:           params.Process,
		Content:           params.Content,
		WorkScope:         params.WorkScope,
		Restrictions:      params.Restrictions,
		Matter:            params.Matter,
		IsRisk:            params.IsRisk,
		TemporaryUsers:    temporaryUsers,
		Explain:           params.Explain,
		IsNotice:          params.IsNotice,
		NoticeTime:        LocalTime{Time: noticeTime},
		StayTime:          params.StayTime,
		ToolNum:           params.ToolNum,
		UserNum:           params.UserNum,
		ClockTime:         LocalTime{Time: clockTime},
		ClockUserID:       params.ClockUserID,
		ToolRemark:        params.ToolRemark,
		LightRemark:       params.LightRemark,
		LightType:         params.LightType,
		GuardRemark:       params.GuardRemark,
		GuardType:         params.GuardType,
		NeedJob:           params.NeedJob,
		ProcessRemark:     params.ProcessRemark,
		QuitToolNum:       params.QuitToolNum,
		QuitUserNum:       params.QuitUserNum,
		QuitToolRemark:    params.QuitToolRemark,
		QuitUserRemark:    params.QuitUserRemark,
		QuitClockTime:     LocalTime{Time: quitClockTime},
		QuitClockLocation: params.QuitClockLocation,
		QuitClockUserID:   params.QuitClockUserID,
		HaveHandover:      params.HaveHandover,
		Handover:          params.Handover,
		HandoverTime:      LocalTime{Time: handoverTime},
		HandoverType:      params.HandoverType,
		ReplayContext:     params.ReplayContext,
		ReplayTime:        LocalTime{Time: replayTime},
		SoundRemark:       params.SoundRemark,
		MobileReceived:    params.MobileReceived,
		WorkedType:        params.WorkedType,
		WorkedRemark:      params.WorkedRemark,
		LogoutJob:         params.LogoutJob,
		LogoutType:        params.LogoutType,
		LogoutRemark:      params.LogoutRemark,
		AuditStatus:       params.AuditStatus,
	}

	if r := db.Model(&Construction{}).Where("id = ? AND deleted = ? ", id, 0).Updates(construction); r.RowsAffected != 1 {
		return false, r.Error
	}

	return true, nil
}
