package models

import (
	"errors"
	"fmt"
	"go-template/common"
	"go-template/utils"

	"github.com/jinzhu/gorm"
)

// 用户model
type User struct {
	Model
	Name         string     `json:"name" gorm:"not null;comment:姓名"`
	Account      string     `json:"account" gorm:"not null;unique;comment:账户"`
	Password     string     `json:"-" gorm:"not null;comment:密码"`
	NickName     string     `json:"nick_name" gorm:"comment:昵称"`
	Gender       string     `json:"gender" gorm:"not null;default:'unknow';comment:性别"`
	Mobile       string     `json:"mobile" gorm:"index;not null;comment:手机号"`
	Email        string     `json:"email" gorm:"comment:邮箱"`
	Status       string     `json:"status" gorm:"default:'normal';comment:状态"`
	Role         Role       `json:"role"`
	RoleID       int        `json:"role_id" gorm:"comment:角色id;==default:'galeone'=="`
	Department   Department `json:"department"`
	DepartmentID int        `json:"department_id" gorm:"comment:部门id;==default:'galeone'=="`
	Projects     []Project  `json:"projects" gorm:"many2many:user_project;"`
}

// 获取用户列表
func GetUsers(params common.PageSearchUserDto) ([]*User, int64, error) {
	var users []*User
	var err error

	tx := db.Where("deleted = 0")

	if len(params.Name) > 0 {
		tx.Where("name like ?", params.Name)
	}
	if len(params.Account) > 0 {
		tx.Where("account like ?", params.Account)
	}
	if len(params.NickName) > 0 {
		tx.Where("nick_name like ?", params.NickName)
	}
	if len(params.Status) > 0 {
		tx.Where("status = ?", params.Status)
	}
	if len(params.Mobile) > 0 {
		tx.Where("mobile like ?", params.Mobile)
	}
	if params.RoleID > 0 {
		tx.Where("role_id = ?", params.RoleID)
	}
	if params.DepartmentID > 0 {
		tx.Where("department_id = ?", params.DepartmentID)
	}
	tx.Preload("Role").Preload("Department")

	if len(params.Order) > 0 {
		tx.Order(params.Order)
	} else {
		tx.Order("id DESC")
	}

	if params.Pagination {
		err = tx.Offset(params.Page - 1).Limit(params.Size).Find(&users).Error
	} else {
		err = tx.Find(&users).Error
	}

	var total int64
	tx.Count(&total)

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0, err
	}

	return users, total, nil
}

// 根据ids获取用户列表
func GetUsersByIds(ids []int) ([]User, error) {
	if len(ids) == 0 {
		return nil, errors.New("参数非法")
	}
	var users []User

	err := db.Where("deleted = 0 AND id IN (?)", ids).Find(&users).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return users, nil
}

// 根据id获取用户信息
func GetUser(id int) (*User, error) {
	var user User
	err := db.Preload("Role").Preload("Department").Where("id = ? AND deleted = ? ", id, 0).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// 新增用户
func AddUser(params common.UserCreateDto) (int, error) {
	projects, _ := GetProjectsByIds(params.ProjectIds)
	user := User{
		Name:         params.Name,
		Account:      params.Account,
		Password:     utils.MD5(params.Password),
		NickName:     params.NickName,
		Gender:       params.Gender,
		Mobile:       params.Mobile,
		Email:        params.Email,
		RoleID:       params.RoleID,
		DepartmentID: params.DepartmentID,
		Projects:     projects,
	}

	if err := db.Create(&user).Error; err != nil {
		return 0, err
	}

	return user.ID, nil
}

// 更新指定用户
func UpdateUser(id int, params common.UserUpdateDto) (bool, error) {
	var oldUser *User
	var err error

	if oldUser, err = GetUser(id); err != nil {
		return false, err
	}

	projects, _ := GetProjectsByIds(params.ProjectIds)
	user := User{
		Name:         params.Name,
		Account:      params.Account,
		Password:     params.Password,
		NickName:     params.NickName,
		Gender:       params.Gender,
		Mobile:       params.Mobile,
		Email:        params.Email,
		Status:       params.Status,
		RoleID:       params.RoleID,
		DepartmentID: params.DepartmentID,
		Projects:     projects,
	}
	if r := db.Model(&oldUser).Updates(user); r.RowsAffected != 1 {
		fmt.Printf("ssss,%s", r.Error)
		return false, r.Error
	}

	return true, nil
}

// 删除用户
func DeleteUsers(ids []int) (int, error) {
	if err := db.Model(&User{}).Where("id IN (?)", ids).Update("deleted", 1).Error; err != nil {
		return 0, err
	}

	return int(db.RowsAffected), nil
}

// 根据账号密码获取用户
func GetUserByLogin(account, password string) (*User, error) {
	var user User
	err := db.Preload("Role").Preload("Department").Where(User{Account: account, Password: password}).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &user, nil
}
