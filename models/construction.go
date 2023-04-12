package models

import (
	"go-template/common"
	"time"

	"gorm.io/gorm"
)

// 施工作业model
type Construction struct {
	Model
	MeasureLibraries []MeasureLibrary `json:"measure_libraries" gorm:"many2many:construction_measure_library;"`
	StartTime        LocalTime        `json:"start_time" gorm:"comment:开始时间"`
	EndTime          LocalTime        `json:"end_time" gorm:"comment:结束时间"`
	ActualTime       LocalTime        `json:"actual_time" gorm:"comment:实际完成时间"`
	EquipmentType    Strs             `json:"equipment_type" gorm:"type:text;comment:设备类型"`
	Location         string           `json:"location" gorm:"comment:作业地点"`
	// 1: '未开始', 2: '进行中', 3: '已完成', 4: '已延期'
	Status string `json:"status" gorm:"comment:状态"`
	// 1: '待提交', 2: '待执行', 3: '审批中', 4: '执行中', 5: '复盘待上传', 6: '录音待上传', 7: '已完成', 8: '已终止', 9: '安全资质审批中'
	JobStatus      string `json:"job_status" gorm:"comment:作业状态"`
	ExecutiveUsers []User `json:"executive_users" gorm:"many2many:construction_user;"`
	Remark         string `json:"remark" gorm:"comment:备注"`
	// 施工任务相关
	ManagerID         int             `json:"manager_id" gorm:"comment:项目经理id"`
	EngineerID        int             `json:"engineer_id" gorm:"comment:项目总工id"`
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
	IsRisk            bool            `json:"is_risk" gorm:"comment:是否涉及高风险"`
	TemporaryUsers    []TemporaryUser `json:"temporary_users" gorm:"many2many:construction_temporary_user"`

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
func GetConstructionPlans(params common.PageSearchConstructionDto) ([]*common.ConstructionPlanDto, int64, error) {
	var constructions []*Construction
	var constructionPlans []*common.ConstructionPlanDto
	var err error
	tx := db.Where("deleted = 0")
	if len(params.Status) > 0 {
		tx.Where("status = ?", params.Status)
	}

	tx.Preload("MeasureLibraries").Preload("ExecutiveUsers").Preload("TemporaryUsers")

	if len(params.Order) > 0 {
		tx.Order(params.Order)
	} else {
		tx.Order("id DESC")
	}

	if params.Pagination {
		err = tx.Offset(params.Page - 1).Limit(params.Size).Find(&constructions).Error
	} else {
		err = tx.Find(&constructions).Error
	}

	var total int64
	tx.Count(&total)

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0, err
	}

	if len(constructions) > 0 {

		for a := 0; a < len(constructions); a++ {
			var measureLibraries []common.MeasureLibraryDto
			var executiveUsers []map[string]interface{}

			for b := 0; b < len(constructions[a].MeasureLibraries); b++ {
				measureLibraries = append(measureLibraries, common.MeasureLibraryDto{
					ID:       constructions[a].MeasureLibraries[b].ID,
					HomeWork: constructions[a].MeasureLibraries[b].HomeWork,
					RiskType: constructions[a].MeasureLibraries[b].RiskType,
					Name:     constructions[a].MeasureLibraries[b].Name,
					Risk:     constructions[a].MeasureLibraries[b].Risk,
					Measures: constructions[a].MeasureLibraries[b].Measures,
				})

			}
			for c := 0; c < len(constructions[a].ExecutiveUsers); c++ {
				executiveUsers = append(executiveUsers, map[string]interface{}{
					"id":   constructions[a].ExecutiveUsers[c].ID,
					"name": constructions[a].ExecutiveUsers[c].Name,
				})
			}

			leader, _ := GetUser(constructions[a].LeaderID)

			constructionPlans = append(constructionPlans, &common.ConstructionPlanDto{
				ID:               constructions[a].ID,
				MeasureLibraries: measureLibraries,
				StartTtime:       constructions[a].StartTime.String(),
				EndTime:          constructions[a].EndTime.String(),
				Location:         constructions[a].Location,
				Remark:           constructions[a].Remark,
				Status:           constructions[a].Status,
				EquipmentType:    constructions[a].EquipmentType,
				Leader:           map[string]interface{}{"id": constructions[a].LeaderID, "name": leader.Name},
				ExecutiveUsers:   executiveUsers,
			})
		}
	}

	return constructionPlans, total, nil
}

// 根据id获取施工作业信息
func GetConstructionPlan(id int) (*Construction, error) {
	var construction Construction
	err := db.Preload("MeasureLibraries").Preload("ExecutiveUsers").Preload("TemporaryUsers").Where("id = ? AND deleted = ? ", id, 0).First(&construction).Error
	if err != nil {
		return nil, err
	}

	return &construction, nil
}

// 新增施工作业计划
func AddConstructionPlan(params common.ConstructionPlanCreateDto) (int, error) {

	startTime, _ := time.ParseInLocation("2006-01-02 15:04:05", params.StartTtime, time.Local)
	endTime, _ := time.ParseInLocation("2006-01-02 15:04:05", params.EndTime, time.Local)
	measureLibraries, _ := GetMeasureLibrariesByIds(params.MeasureLibraryIds)
	executiveUsers, _ := GetUsersByIds(params.ExecutiveUserIds)

	construction := Construction{
		StartTime:        LocalTime{Time: startTime},
		EndTime:          LocalTime{Time: endTime},
		MeasureLibraries: measureLibraries,
		LeaderID:         params.LeaderID,
		ExecutiveUsers:   executiveUsers,
		EquipmentType:    params.EquipmentType,
		Location:         params.Location,
		Remark:           params.Remark,
	}

	if err := db.Create(&construction).Error; err != nil {
		return 0, err
	}

	return construction.ID, nil
}

// 更新施工作业计划
func UpdateConstructionPlan(id int, params common.ConstructionPlanUpdateDto) (bool, error) {
	var oldConstruction *Construction
	var err error

	if oldConstruction, err = GetConstruction(id); err != nil || oldConstruction == nil {
		return false, err
	}
	startTime, _ := time.ParseInLocation("2006-01-02 15:04:05", params.StartTtime, time.Local)
	endTime, _ := time.ParseInLocation("2006-01-02 15:04:05", params.EndTime, time.Local)
	measureLibraries, _ := GetMeasureLibrariesByIds(params.MeasureLibraryIds)
	executiveUsers, _ := GetUsersByIds(params.ExecutiveUserIds)

	construction := Construction{
		StartTime:        LocalTime{Time: startTime},
		EndTime:          LocalTime{Time: endTime},
		MeasureLibraries: measureLibraries,
		LeaderID:         params.LeaderID,
		ExecutiveUsers:   executiveUsers,
		EquipmentType:    params.EquipmentType,
		Location:         params.Location,
		Remark:           params.Remark,
	}

	if err := db.Create(&construction).Error; err != nil {
		return false, err
	}

	return true, nil
}

// 获取施工作业列表
func GetConstructions(params common.PageSearchConstructionDto) ([]*Construction, int64, error) {
	var constructions []*Construction
	var err error
	tx := db.Where("deleted = 0")
	if len(params.Status) > 0 {
		tx.Where("status = ?", params.Status)
	}

	tx.Preload("MeasureLibraries").Preload("ExecutiveUsers").Preload("TemporaryUsers")

	if len(params.Order) > 0 {
		tx.Order(params.Order)
	} else {
		tx.Order("id DESC")
	}

	if params.Pagination {
		err = tx.Offset(params.Page - 1).Limit(params.Size).Find(&constructions).Error
	} else {
		err = tx.Find(&constructions).Error
	}

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
	err := db.Preload("MeasureLibraries").Preload("ExecutiveUsers").Preload("TemporaryUsers").Where("id = ? AND deleted = ? ", id, 0).First(&construction).Error
	if err != nil {
		return nil, err
	}

	return &construction, nil
}

// 更新施工作业
func UpdateConstruction(id int, params common.ConstructionUpdateDto) (bool, error) {
	var oldConstruction *Construction
	var temporaryUsers []TemporaryUser
	var err error

	if oldConstruction, err = GetConstruction(id); err != nil || oldConstruction == nil {
		return false, err
	}

	startTime, _ := time.ParseInLocation("20060102150405", params.StartTtime, time.Local)
	endTime, _ := time.ParseInLocation("20060102150405", params.EndTime, time.Local)
	measureLibraries, _ := GetMeasureLibrariesByIds(params.MeasureLibraryIds)
	executiveUsers, _ := GetUsersByIds(params.ExecutiveUserIds)
	noticeTime, _ := time.ParseInLocation("20060102150405", params.NoticeTime, time.Local)
	clockTime, _ := time.ParseInLocation("20060102150405", params.ClockTime, time.Local)
	quitClockTime, _ := time.ParseInLocation("20060102150405", params.QuitClockTime, time.Local)
	handoverTime, _ := time.ParseInLocation("20060102150405", params.HandoverTime, time.Local)
	replayTime, _ := time.ParseInLocation("20060102150405", params.ReplayTime, time.Local)

	for a := 0; a < len(params.TemporaryUsers); a++ {
		temporaryUsers = append(temporaryUsers, TemporaryUser{
			Model:         Model{ID: params.TemporaryUsers[a].ID},
			Name:          params.TemporaryUsers[a].Name,
			Mobile:        params.TemporaryUsers[a].Mobile,
			Department:    params.TemporaryUsers[a].Department,
			DockingUserID: params.TemporaryUsers[a].DockingUserID,
		})
	}

	oldConstruction.StartTime = LocalTime{Time: startTime}
	oldConstruction.EndTime = LocalTime{Time: endTime}
	oldConstruction.MeasureLibraries = measureLibraries
	oldConstruction.ExecutiveUsers = executiveUsers
	oldConstruction.EquipmentType = params.EquipmentType
	oldConstruction.Location = params.Location
	oldConstruction.Remark = params.Remark
	oldConstruction.ManagerID = params.ManagerID
	oldConstruction.EngineerID = params.EngineerID
	oldConstruction.LeaderID = params.LeaderID
	oldConstruction.RecipientID = params.RecipientID
	oldConstruction.PhoneID = params.PhoneID
	oldConstruction.TerminationUserID = params.TerminationUserID
	oldConstruction.StopReason = params.StopReason
	oldConstruction.Process = params.Process
	oldConstruction.Content = params.Content
	oldConstruction.WorkScope = params.WorkScope
	oldConstruction.Restrictions = params.Restrictions
	oldConstruction.Matter = params.Matter
	oldConstruction.IsRisk = params.IsRisk
	oldConstruction.TemporaryUsers = temporaryUsers
	oldConstruction.Explain = params.Explain
	oldConstruction.IsNotice = params.IsNotice
	oldConstruction.NoticeTime = LocalTime{Time: noticeTime}
	oldConstruction.StayTime = params.StayTime
	oldConstruction.ToolNum = params.ToolNum
	oldConstruction.UserNum = params.UserNum
	oldConstruction.ClockTime = LocalTime{Time: clockTime}
	oldConstruction.ClockUserID = params.ClockUserID
	oldConstruction.ToolRemark = params.ToolRemark
	oldConstruction.LightRemark = params.LightRemark
	oldConstruction.LightType = params.LightType
	oldConstruction.GuardRemark = params.GuardRemark
	oldConstruction.GuardType = params.GuardType
	oldConstruction.NeedJob = params.NeedJob
	oldConstruction.ProcessRemark = params.ProcessRemark
	oldConstruction.QuitToolNum = params.QuitToolNum
	oldConstruction.QuitUserNum = params.QuitUserNum
	oldConstruction.QuitToolRemark = params.QuitToolRemark
	oldConstruction.QuitUserRemark = params.QuitUserRemark
	oldConstruction.QuitClockTime = LocalTime{Time: quitClockTime}
	oldConstruction.QuitClockLocation = params.QuitClockLocation
	oldConstruction.QuitClockUserID = params.QuitClockUserID
	oldConstruction.HaveHandover = params.HaveHandover
	oldConstruction.Handover = params.Handover
	oldConstruction.HandoverTime = LocalTime{Time: handoverTime}
	oldConstruction.HandoverType = params.HandoverType
	oldConstruction.ReplayContext = params.ReplayContext
	oldConstruction.ReplayTime = LocalTime{Time: replayTime}
	oldConstruction.SoundRemark = params.SoundRemark
	oldConstruction.MobileReceived = params.MobileReceived
	oldConstruction.WorkedType = params.WorkedType
	oldConstruction.WorkedRemark = params.WorkedRemark
	oldConstruction.LogoutJob = params.LogoutJob
	oldConstruction.LogoutType = params.LogoutType
	oldConstruction.LogoutRemark = params.LogoutRemark
	oldConstruction.AuditStatus = params.AuditStatus

	if r := db.Updates(&oldConstruction); r.RowsAffected != 1 {
		return false, r.Error
	}

	return true, nil
}

// 删除施工作业
func DeleteConstructions(ids []int) (int, error) {
	r := db.Model(&Construction{}).Where("id IN (?)", ids).Updates(map[string]interface{}{"deleted": 1})
	if r.Error != nil {
		return 0, r.Error
	}

	return int(r.RowsAffected), nil
}
