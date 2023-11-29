package model

import "time"

type JSONMap map[string]interface{}

type TaskStatus string

const (
	TaskStatusPending TaskStatus = "pending"
	TaskStatusRunning TaskStatus = "running"
	TaskStatusSuccess TaskStatus = "success"
	TaskStatusFailed  TaskStatus = "failed"
	TaskStatusCancel  TaskStatus = "cancel"
)

type TaskRecord struct {
	ID        uint       `gorm:"primary_key" form:"id" json:"id" binding:"omitempty,gte=1" label:"taskID"`
	TaskID    string     `gorm:"type:varchar(255)" form:"task_id" json:"task_id" label:"task UUID" `
	Name      string     `form:"name" json:"name" binding:"required" label:"task name"`
	Type      string     `gorm:"type:varchar(255)" form:"type" json:"type" binding:"required" label:"task type"`
	Params    JSONMap    `form:"params" json:"params" binding:"required" label:"task params"`
	Status    TaskStatus `gorm:"type:varchar(255)" form:"status" json:"status"  label:"task status"`
	Result    string     `gorm:"result" form:"result" json:"result"  label:"task result"`
	CreatedAt time.Time  `json:"created_at,omitempty"`
	UpdatedAt time.Time  `json:"updated_at,omitempty"`
	GroupID   string     `gorm:"type:varchar(255)" form:"group_id" json:"group_id" binding:"required" label:"group_id"`
	Priority  uint       `form:"priority" json:"priority" binding:"required" label:"priority"`
	UserId    uint       `form:"user_id" json:"user_id" binding:"required" label:"user_id"`
}

type TaskResult struct {
	TaskID    string  `json:"task_id"`
	ErrorCode int     `json:"error_code"`
	ErrMsg    string  `json:"err_msg"`
	Data      JSONMap `json:"data"`
}
