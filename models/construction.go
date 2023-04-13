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
	ToolNum     int              `json:"tool_num" gorm:"comment:工具数量"`
	UserNum     int              `json:"user_num" gorm:"comment:人员数量"`
	ClockTime   common.LocalTime `json:"clock_time" gorm:"comment:打卡时间"`
	ClockUserID int              `json:"clock_user_id" gorm:"comment:打卡人id"`
	ToolRemark  string           `json:"tool_remark" gorm:"type:text;comment:工具清点备注"`

	// 作业边界
	LightRemark   string `json:"light_remark" gorm:"type:text;comment:红闪灯备注"`
	LightType     string `json:"light_type" gorm:"comment:红闪灯类型"`
	GuardRemark   string `json:"guard_remark" gorm:"type:text;comment:防护员备注"`
	GuardType     string `json:"guard_type" gorm:"comment:防护员类型"`
	NeedJob       bool   `json:"need_job" gorm:"comment:是否作业边界"`
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
	HaveHandover bool             `json:"have_handover" gorm:"comment:是否交接"`
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
	LogoutJob      bool             `json:"logout_job" gorm:"comment:是否注销作业令"`
	LogoutType     string           `json:"logout_type" gorm:"comment:注销类型"`
	LogoutRemark   string           `json:"logout_remark" gorm:"type:text;comment:作业注销备注"`

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
			status, _ := GetConstructionStatus("status", constructions[a].Status)
			constructionPlans = append(constructionPlans, &common.ConstructionPlanDto{
				ID:               constructions[a].ID,
				MeasureLibraries: measureLibraries,
				StartTtime:       constructions[a].StartTime,
				EndTime:          constructions[a].EndTime,
				Location:         constructions[a].Location,
				Remark:           constructions[a].Remark,
				Status:           status,
				EquipmentType:    constructions[a].EquipmentType,
				Leader:           map[string]interface{}{"id": constructions[a].LeaderID, "name": leader.Name},
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

	for b := 0; b < len(construction.MeasureLibraries); b++ {
		measureLibraries = append(measureLibraries, common.MeasureLibraryDto{
			ID:       construction.MeasureLibraries[b].ID,
			HomeWork: construction.MeasureLibraries[b].HomeWork,
			RiskType: construction.MeasureLibraries[b].RiskType,
			Name:     construction.MeasureLibraries[b].Name,
			Risk:     construction.MeasureLibraries[b].Risk,
			Measures: construction.MeasureLibraries[b].Measures,
		})

	}

	for c := 0; c < len(construction.ExecutiveUsers); c++ {
		executiveUsers = append(executiveUsers, map[string]interface{}{
			"id":   construction.ExecutiveUsers[c].ID,
			"name": construction.ExecutiveUsers[c].Name,
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
func AddConstructionPlan(params common.ConstructionPlanCreateDto) (int, error) {

	startTime, _ := time.ParseInLocation("2006-01-02 15:04:05", params.StartTtime, time.Local)
	endTime, _ := time.ParseInLocation("2006-01-02 15:04:05", params.EndTime, time.Local)
	measureLibraries, _ := GetMeasureLibrariesByIds(params.MeasureLibraryIds)
	executiveUsers, _ := GetUsersByIds(params.ExecutiveUserIds)

	construction := Construction{
		StartTime:        common.LocalTime{Time: startTime},
		EndTime:          common.LocalTime{Time: endTime},
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
func UpdateConstruction(id int, params common.ConstructionUpdateDto) (bool, error) {
	var oldConstruction *Construction
	var temporaryUsers []TemporaryUser
	var err error

	if oldConstruction, err = GetConstruction(id); err != nil || oldConstruction == nil {
		return false, err
	}

	startTime, _ := time.ParseInLocation("20060102150405", params.StartTtime, time.Local)
	endTime, _ := time.ParseInLocation("20060102150405", params.EndTime, time.Local)
	executiveUsers, _ := GetUsersByIds(params.ExecutiveUserIds)

	for a := 0; a < len(params.TemporaryUsers); a++ {
		temporaryUsers = append(temporaryUsers, TemporaryUser{
			Model:         Model{ID: params.TemporaryUsers[a].ID},
			Name:          params.TemporaryUsers[a].Name,
			Mobile:        params.TemporaryUsers[a].Mobile,
			Department:    params.TemporaryUsers[a].Department,
			DockingUserID: params.TemporaryUsers[a].DockingUserID,
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
	}

	if r := db.Updates(&oldConstruction); r.RowsAffected != 1 {
		return false, r.Error
	}

	return true, nil
}

// 审批指定施工作业
func ApproveConstruction(id int) (bool, error) {
	var oldConstruction *Construction
	var err error

	if oldConstruction, err = GetConstruction(id); err != nil || oldConstruction == nil {
		return false, err
	}

	var ExecutiveUserNoQualificationCount int
	for _, e := range oldConstruction.ExecutiveUsers {
		if !e.HasQualification {
			ExecutiveUserNoQualificationCount += 1
		}
	}
	if ExecutiveUserNoQualificationCount > 0 {
		// 存在执行人员无资质则需要人员资质审核通过后才能领取执行
		oldConstruction.JobStatus = "9"
	} else {
		// 更新状态为待执行
		oldConstruction.JobStatus = "2"
	}

	if r := db.Updates(&oldConstruction); r.RowsAffected != 1 {
		return false, r.Error
	}

	return true, nil
}

// 领取施工作业
func ReceiveConstruction(id int) (bool, error) {
	var oldConstruction *Construction
	var err error

	if oldConstruction, err = GetConstruction(id); err != nil || oldConstruction == nil {
		return false, err
	}
	// 更新状态为执行中
	oldConstruction.JobStatus = "4"
	if r := db.Updates(&oldConstruction); r.RowsAffected != 1 {
		return false, r.Error
	}

	return true, nil
}

// 终止施工作业
func StopConstruction(id int) (bool, error) {
	var oldConstruction *Construction
	var err error

	if oldConstruction, err = GetConstruction(id); err != nil || oldConstruction == nil {
		return false, err
	}
	// 更新状态为已终止
	oldConstruction.JobStatus = "8"
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
