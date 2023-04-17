package models

import (
	"errors"
	"go-template/common"
	"time"

	"github.com/jasonlvhit/gocron"
	"gorm.io/gorm"
)

// 施工作业model
type Construction struct {
	Model
	MeasureLibraries []MeasureLibrary `json:"measure_libraries" gorm:"many2many:construction_measure_library;"`
	StartTime        common.LocalTime `json:"start_time" gorm:"comment:开始时间"`
	EndTime          common.LocalTime `json:"end_time" gorm:"comment:结束时间"`
	ActualTime       common.LocalTime `json:"actual_time" gorm:"comment:实际完成时间"`
	EquipmentType    common.Strs      `json:"equipment_type" gorm:"type:text;comment:设备类型"`
	Location         string           `json:"location" gorm:"comment:作业地点"`
	// 1: '未开始', 2: '进行中', 3: '已完成', 4: '已延期'
	Status string `json:"status" gorm:"default:'1';comment:状态"`
	// 1: '待提交', 2: '待执行', 3: '审批中', 4: '执行中', 5: '复盘待上传', 6: '录音待上传', 7: '已完成', 8: '已终止', 9: '安全资质审批中'
	JobStatus      string `json:"job_status" gorm:"default:'1';comment:作业状态"`
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
	Explain    string           `json:"explain" gorm:"comment:补充说明"`
	IsNotice   bool             `json:"is_notice" gorm:"comment:是否确认宣读"`
	NoticeTime common.LocalTime `json:"notice_time" gorm:"comment:宣读时间"`
	StayTime   int              `json:"stay_time" gorm:"comment:停留时间(单位秒)"`

	// 工具人员清点
	ToolNum       int              `json:"tool_num" gorm:"comment:工具数量"`
	UserNum       int              `json:"user_num" gorm:"comment:人员数量"`
	ClockTime     common.LocalTime `json:"clock_time" gorm:"comment:打卡时间"`
	ClockUserID   int              `json:"clock_user_id" gorm:"comment:打卡人id"`
	ClockLocation string           `json:"clock_location" gorm:"comment:打卡地点"`
	ToolRemark    string           `json:"tool_remark" gorm:"type:text;comment:工具清点备注"`

	// 作业边界
	LightRemark string `json:"light_remark" gorm:"type:text;comment:红闪灯备注"`
	LightType   string `json:"light_type" gorm:"comment:红闪灯类型"`
	GuardRemark string `json:"guard_remark" gorm:"type:text;comment:防护员备注"`
	GuardType   string `json:"guard_type" gorm:"comment:防护员类型"`
	// NeedJob       bool   `json:"need_job" gorm:"comment:是否作业边界"`
	ProcessRemark string `json:"process_remark" gorm:"type:text;comment:作业过程备注"`

	// 作业出清
	QuitToolNum       int              `json:"quit_tool_num" gorm:"comment:出清工具数量"`
	QuitUserNum       int              `json:"quit_user_num" gorm:"comment:出清人员数量"`
	QuitToolRemark    string           `json:"quit_tool_remark" gorm:"type:text;comment:工具备注"`
	QuitUserRemark    string           `json:"quit_user_remark" gorm:"type:text;comment:人员备注"`
	QuitClockTime     common.LocalTime `json:"quit_clock_time" gorm:"comment:出清打卡时间"`
	QuitClockLocation string           `json:"quit_clock_location" gorm:"comment:出清打卡地点"`
	QuitClockUserID   int              `json:"quit_clock_user_id" gorm:"comment:出清打卡人id"`

	// 作业交接
	// HaveHandover bool             `json:"have_handover" gorm:"comment:是否交接"`
	Handover     string           `json:"handover" gorm:"type:text;comment:交接内容"`
	HandoverTime common.LocalTime `json:"handover_time" gorm:"comment:交接时间"`
	HandoverType string           `json:"handover_type" gorm:"comment:交接类型"`

	// 每日复盘
	ReplayContext  string           `json:"replay_context" gorm:"type:text;comment:复盘内容"`
	ReplayTime     common.LocalTime `json:"replay_time" gorm:"comment:复盘时间"`
	SoundRemark    string           `json:"sound_remark" gorm:"type:text;comment:录音备注"`
	MobileReceived bool             `json:"mobile_received" gorm:"comment:手机是否已领取"`
	WorkedType     string           `json:"worker_type" gorm:"comment:作业登记类型"`
	WorkedRemark   string           `json:"worker_remark" gorm:"type:text;comment:作业登记备注"`
	// LogoutJob      bool             `json:"logout_job" gorm:"comment:是否注销作业令"`
	LogoutType    string `json:"logout_type" gorm:"comment:注销类型"`
	LogoutRemark  string `json:"logout_remark" gorm:"type:text;comment:作业注销备注"`
	ApproveStatus string `json:"approve_status" gorm:"comment:审批状态(pass表示通过,refuse表示驳回)"`
	ApproveRemark string `json:"approve_remark" gorm:"type:text;comment:审批备注"`

	AuditStatus string `json:"audit_status" gorm:"comment:审计状态"`
	Logs        []Log  `json:"logs" gorm:"foreignKey:ConstructionID"`
	Files       []File `json:"files" gorm:"foreignKey:ConstructionID"`
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

		for _, a := range constructions {
			var measureLibraries []common.MeasureLibraryDto
			var executiveUsers []map[string]interface{}

			for _, b := range a.MeasureLibraries {
				measureLibraries = append(measureLibraries, common.MeasureLibraryDto{
					ID:       b.ID,
					HomeWork: b.HomeWork,
					RiskType: b.RiskType,
					Name:     b.Name,
					Risk:     b.Risk,
					Measures: b.Measures,
				})
			}

			for _, c := range a.ExecutiveUsers {
				executiveUsers = append(executiveUsers, map[string]interface{}{
					"id":   c.ID,
					"name": c.Name,
				})
			}

			leader, _ := GetUser(a.LeaderID)
			status, _ := GetConstructionStatus("status", a.Status)
			constructionPlans = append(constructionPlans, &common.ConstructionPlanDto{
				ID:               a.ID,
				MeasureLibraries: measureLibraries,
				StartTtime:       a.StartTime,
				EndTime:          a.EndTime,
				Location:         a.Location,
				Remark:           a.Remark,
				Status:           status,
				EquipmentType:    a.EquipmentType,
				Leader:           map[string]interface{}{"id": a.LeaderID, "name": leader.Name},
				ExecutiveUsers:   executiveUsers,
			})
		}
	}

	return constructionPlans, total, nil
}

// 根据id获取施工作业信息
func GetConstructionPlan(id int) (*common.ConstructionPlanDto, error) {
	var construction Construction
	var constructionPlan common.ConstructionPlanDto
	var measureLibraries []common.MeasureLibraryDto
	var executiveUsers []map[string]interface{}
	err := db.Preload("MeasureLibraries").Preload("ExecutiveUsers").Preload("TemporaryUsers").Where("id = ? AND deleted = ? ", id, 0).First(&construction).Error
	if err != nil {
		return nil, err
	}

	for _, b := range construction.MeasureLibraries {
		measureLibraries = append(measureLibraries, common.MeasureLibraryDto{
			ID:       b.ID,
			HomeWork: b.HomeWork,
			RiskType: b.RiskType,
			Name:     b.Name,
			Risk:     b.Risk,
			Measures: b.Measures,
		})

	}

	for _, c := range construction.ExecutiveUsers {
		executiveUsers = append(executiveUsers, map[string]interface{}{
			"id":   c.ID,
			"name": c.Name,
		})
	}
	leader, _ := GetUser(construction.LeaderID)
	status, _ := GetConstructionStatus("status", construction.Status)
	constructionPlan = common.ConstructionPlanDto{
		ID:               construction.ID,
		MeasureLibraries: measureLibraries,
		StartTtime:       construction.StartTime,
		EndTime:          construction.EndTime,
		Location:         construction.Location,
		Remark:           construction.Remark,
		Status:           status,
		EquipmentType:    construction.EquipmentType,
		Leader:           map[string]interface{}{"id": construction.LeaderID, "name": leader.Name},
		ExecutiveUsers:   executiveUsers,
	}

	return &constructionPlan, nil
}

// 根据状态获取状态对象
func GetConstructionStatus(statusType string, status string) (map[string]interface{}, error) {
	if statusType == "status" {
		switch status {
		case "1":
			return map[string]interface{}{"id": "1", "name": "未开始"}, nil
		case "2":
			return map[string]interface{}{"id": "2", "name": "进行中"}, nil
		case "3":
			return map[string]interface{}{"id": "3", "name": "已完成"}, nil
		case "4":
			return map[string]interface{}{"id": "4", "name": "已延期"}, nil
		}
	} else if statusType == "jobStatus" {
		switch status {
		case "1":
			return map[string]interface{}{"id": "1", "name": "待提交"}, nil
		case "2":
			return map[string]interface{}{"id": "2", "name": "待执行"}, nil
		case "3":
			return map[string]interface{}{"id": "3", "name": "审批中"}, nil
		case "4":
			return map[string]interface{}{"id": "4", "name": "执行中"}, nil
		case "5":
			return map[string]interface{}{"id": "5", "name": "复盘待上传"}, nil
		case "6":
			return map[string]interface{}{"id": "6", "name": "录音待上传"}, nil
		case "7":
			return map[string]interface{}{"id": "7", "name": "已完成"}, nil
		case "8":
			return map[string]interface{}{"id": "8", "name": "已终止"}, nil
		case "9":
			return map[string]interface{}{"id": "9", "name": "安全资质审批中"}, nil
		}
	}
	return nil, errors.New("参数错误")
}

// 新增施工作业计划
func AddConstructionPlan(params common.ConstructionPlanCreateDto, currentUser User) (int, error) {

	startTime, _ := time.ParseInLocation("2006-01-02 15:04:05", params.StartTtime, time.Local)
	endTime, _ := time.ParseInLocation("2006-01-02 15:04:05", params.EndTime, time.Local)
	measureLibraries, _ := GetMeasureLibrariesByIds(params.MeasureLibraryIds)
	executiveUsers, _ := GetUsersByIds(params.ExecutiveUserIds)

	var logs []Log
	logs = append(logs, Log{Content: "创建施工作业计划", UserID: currentUser.ID})

	construction := Construction{
		StartTime:        common.LocalTime{Time: startTime},
		EndTime:          common.LocalTime{Time: endTime},
		MeasureLibraries: measureLibraries,
		LeaderID:         params.LeaderID,
		ExecutiveUsers:   executiveUsers,
		EquipmentType:    params.EquipmentType,
		Location:         params.Location,
		Remark:           params.Remark,
		Logs:             logs,
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

	oldConstruction.StartTime = common.LocalTime{Time: startTime}
	oldConstruction.EndTime = common.LocalTime{Time: endTime}
	oldConstruction.MeasureLibraries = measureLibraries
	oldConstruction.LeaderID = params.LeaderID
	oldConstruction.ExecutiveUsers = executiveUsers
	oldConstruction.EquipmentType = params.EquipmentType
	oldConstruction.Location = params.Location
	oldConstruction.Remark = params.Remark

	if err := db.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&oldConstruction).Error; err != nil {
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
func UpdateConstruction(id int, params common.ConstructionUpdateDto, currentUser User) (bool, error) {
	var oldConstruction *Construction
	var temporaryUsers []TemporaryUser
	var err error

	if oldConstruction, err = GetConstruction(id); err != nil || oldConstruction == nil {
		return false, err
	}

	startTime, _ := time.ParseInLocation("20060102150405", params.StartTtime, time.Local)
	endTime, _ := time.ParseInLocation("20060102150405", params.EndTime, time.Local)
	executiveUsers, _ := GetUsersByIds(params.ExecutiveUserIds)

	for _, a := range params.TemporaryUsers {
		temporaryUsers = append(temporaryUsers, TemporaryUser{
			Model:         Model{ID: a.ID},
			Name:          a.Name,
			Mobile:        a.Mobile,
			Department:    a.Department,
			DockingUserID: a.DockingUserID,
		})
	}

	oldConstruction.StartTime = common.LocalTime{Time: startTime}
	oldConstruction.EndTime = common.LocalTime{Time: endTime}
	oldConstruction.ExecutiveUsers = executiveUsers
	oldConstruction.EquipmentType = params.EquipmentType
	oldConstruction.Location = params.Location
	oldConstruction.Remark = params.Remark
	oldConstruction.ManagerID = params.ManagerID
	oldConstruction.EngineerID = params.EngineerID
	oldConstruction.LeaderID = params.LeaderID
	oldConstruction.Process = params.Process
	oldConstruction.WorkScope = params.WorkScope
	oldConstruction.Restrictions = params.Restrictions
	oldConstruction.Matter = params.Matter
	oldConstruction.IsRisk = params.IsRisk
	oldConstruction.TemporaryUsers = temporaryUsers
	if params.IsSubmit {
		oldConstruction.JobStatus = "3"
		oldConstruction.Logs = append(oldConstruction.Logs, Log{Content: "提交作业", UserID: currentUser.ID, ConstructionID: oldConstruction.ID})
	}

	if r := db.Updates(&oldConstruction); r.RowsAffected != 1 {
		return false, r.Error
	}

	return true, nil
}

// 审批指定施工作业
func ApproveConstruction(id int, params common.ConstructionApproveDto, currentUser User) (bool, error) {
	var oldConstruction *Construction
	var err error

	if oldConstruction, err = GetConstruction(id); err != nil || oldConstruction == nil {
		return false, err
	}
	if params.Status != "pass" && params.Status != "refuse" {
		return false, errors.New("参数错误！")
	}

	if params.Status == "refuse" && len(params.Remark) == 0 {
		return false, errors.New("未填写驳回备注！")
	}

	var ExecutiveUserNoQualificationCount int
	for _, e := range oldConstruction.ExecutiveUsers {
		if !e.HasQualification {
			ExecutiveUserNoQualificationCount += 1
		}
	}
	if params.Status == "pass" {
		if ExecutiveUserNoQualificationCount > 0 {
			// 存在执行人员无资质则需要人员资质审核通过后才能领取执行
			oldConstruction.JobStatus = "9"
		} else {
			// 更新状态为待执行
			oldConstruction.JobStatus = "2"
		}
	} else {
		// 驳回则更新状态为待提交
		oldConstruction.JobStatus = "1"
	}
	oldConstruction.Status = params.Status
	oldConstruction.Remark = params.Remark
	oldConstruction.Logs = append(oldConstruction.Logs, Log{Content: "审核作业", UserID: currentUser.ID, ConstructionID: oldConstruction.ID})
	if r := db.Updates(&oldConstruction); r.RowsAffected != 1 {
		return false, r.Error
	}

	return true, nil
}

// 领取施工作业
func ReceiveConstruction(id int, currentUser User) (bool, error) {
	var oldConstruction *Construction
	var err error

	if oldConstruction, err = GetConstruction(id); err != nil || oldConstruction == nil {
		return false, err
	}
	// 更新状态为执行中
	oldConstruction.JobStatus = "4"
	oldConstruction.Logs = append(oldConstruction.Logs, Log{Content: "领取任务", UserID: currentUser.ID, ConstructionID: oldConstruction.ID})
	if r := db.Updates(&oldConstruction); r.RowsAffected != 1 {
		return false, r.Error
	}

	return true, nil
}

// 终止施工作业
func StopConstruction(id int, params common.ConstructionStopDto, currentUser User) (bool, error) {
	var oldConstruction *Construction
	var err error

	if oldConstruction, err = GetConstruction(id); err != nil || oldConstruction == nil {
		return false, err
	}
	// 更新状态为已终止
	oldConstruction.JobStatus = "8"
	oldConstruction.StopReason = params.StopReason
	oldConstruction.Logs = append(oldConstruction.Logs, Log{Content: "终止任务", UserID: currentUser.ID, ConstructionID: oldConstruction.ID})
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

// 获取施工作业数量
func GetConstructionCount() (int64, error) {
	var total int64
	if err := db.Model(&Construction{}).Where("deleted = 0").Count(&total).Error; err != nil {
		return 0, err
	}
	return total, nil
}

// 提交施工作业
func SubmitConstruction(id int, params common.ConstructionSubmitDto) (bool, error) {
	var oldConstruction *Construction
	var err error

	if oldConstruction, err = GetConstruction(id); err != nil || oldConstruction == nil {
		return false, err
	}

	if oldConstruction.Status == "7" || oldConstruction.Status == "8" {
		return false, errors.New("任务已完结，不允许再提交！")
	}

	var files []File = oldConstruction.Files
	for _, f := range params.Files {
		if f.ID == 0 {
			files = append(files, File{Name: f.Name, Type: f.Type, Url: f.Url, ConstructionID: oldConstruction.ID})
		}
	}

	noticeTime, _ := time.ParseInLocation("2006-01-02 15:04:05", params.NoticeTime, time.Local)
	clockTime, _ := time.ParseInLocation("2006-01-02 15:04:05", params.ClockTime, time.Local)
	quitClockTime, _ := time.ParseInLocation("2006-01-02 15:04:05", params.QuitClockTime, time.Local)
	handoverTime, _ := time.ParseInLocation("2006-01-02 15:04:05", params.HandoverTime, time.Local)

	oldConstruction.Files = files
	oldConstruction.IsNotice = params.IsNotice
	oldConstruction.NoticeTime = common.LocalTime{Time: noticeTime}
	oldConstruction.StayTime = params.StayTime
	oldConstruction.ToolNum = params.ToolNum
	oldConstruction.UserNum = params.UserNum
	oldConstruction.ClockTime = common.LocalTime{Time: clockTime}
	oldConstruction.ClockUserID = params.ClockUserID
	oldConstruction.ClockLocation = params.ClockLocation
	oldConstruction.ToolRemark = params.ToolRemark
	oldConstruction.LightRemark = params.LightRemark
	oldConstruction.LightType = params.LightType
	oldConstruction.GuardRemark = params.GuardRemark
	oldConstruction.GuardType = params.GuardType
	oldConstruction.ProcessRemark = params.ProcessRemark
	oldConstruction.QuitToolNum = params.QuitToolNum
	oldConstruction.QuitUserNum = params.QuitUserNum
	oldConstruction.QuitToolRemark = params.QuitToolRemark
	oldConstruction.QuitUserRemark = params.QuitUserRemark
	oldConstruction.QuitClockTime = common.LocalTime{Time: quitClockTime}
	oldConstruction.QuitClockLocation = params.QuitClockLocation
	oldConstruction.QuitClockUserID = params.QuitClockUserID
	oldConstruction.Handover = params.Handover
	oldConstruction.HandoverTime = common.LocalTime{Time: handoverTime}
	oldConstruction.HandoverType = params.HandoverType
	oldConstruction.WorkedType = params.WorkedType
	oldConstruction.WorkedRemark = params.WorkedRemark
	oldConstruction.LogoutType = params.LogoutType
	oldConstruction.LogoutRemark = params.LogoutRemark
	// 状态更改为复盘待上传
	oldConstruction.Status = "5"

	if r := db.Updates(&oldConstruction); r.RowsAffected != 1 {
		return false, r.Error
	}

	return true, nil
}

// 提交复盘
func SubmitConstructionReplay(id int, params common.ConstructionSubmitReplayDto) (bool, error) {
	var oldConstruction *Construction
	var err error

	if oldConstruction, err = GetConstruction(id); err != nil || oldConstruction == nil {
		return false, err
	}

	if oldConstruction.Status == "7" || oldConstruction.Status == "8" {
		return false, errors.New("任务已完结，不允许再提交！")
	}
	var newFiles []File = oldConstruction.Files
	for _, f := range params.Files {
		if f.ID == 0 {
			newFiles = append(newFiles, File{Name: f.Name, Type: "replay", Url: f.Name, ConstructionID: oldConstruction.ID})
		}
	}
	oldConstruction.Files = newFiles
	oldConstruction.ReplayContext = params.ReplayContext
	oldConstruction.ReplayTime = common.LocalTime{Time: time.Now()}
	// 状态更改为录音待上传
	oldConstruction.Status = "6"

	if r := db.Updates(&oldConstruction); r.RowsAffected != 1 {
		return false, r.Error
	}

	return true, nil
}

// 提交录音文件
func SubmitConstructionSound(id int, params common.ConstructionSubmitSoundDto) (bool, error) {
	var oldConstruction *Construction
	var err error

	if oldConstruction, err = GetConstruction(id); err != nil || oldConstruction == nil {
		return false, err
	}

	if oldConstruction.Status == "7" || oldConstruction.Status == "8" {
		return false, errors.New("任务已完结，不允许再提交！")
	}
	var newFiles []File = oldConstruction.Files
	for _, f := range params.Files {
		if f.ID == 0 {
			newFiles = append(newFiles, File{Name: f.Name, Type: f.Type, Url: f.Name, ConstructionID: oldConstruction.ID})
		}
	}
	oldConstruction.Files = newFiles
	oldConstruction.SoundRemark = params.SoundRemark
	// 状态更改为已完成
	oldConstruction.Status = "7"

	if r := db.Updates(&oldConstruction); r.RowsAffected != 1 {
		return false, r.Error
	}

	return true, nil
}

// 同步施工作业状态(每隔5s执行一次)
func SyncConstructionStatus() {
	s := gocron.NewScheduler()
	s.Every(5).Seconds().Do(func() {
		var unStartCount int64
		db.Model(&Construction{}).Where("deleted=0 AND status='1' AND start_time < now() AND end_time > now()").Count(&unStartCount)
		if unStartCount > 0 {
			// 当前时间处于开始与结束时间范围内且状态为未开始，自动改为进行中
			db.Model(&Construction{}).Where("deleted=0 AND status='1' AND start_time < now() AND end_time > now()").Updates(map[string]interface{}{"status": "2"})
		}

		var expiredCount int64
		db.Model(&Construction{}).Where("deleted=0 AND status!='3' AND end_time < now()").Count(&expiredCount)
		if expiredCount > 0 {
			// 当前时间大于结束时间且状态不为已完成，自动改为已延期
			db.Model(&Construction{}).Where("deleted=0 AND status!='3' AND end_time < now()").Updates(map[string]interface{}{"status": "4"})
		}
	})
	<-s.Start()
}
