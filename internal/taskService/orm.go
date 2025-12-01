package taskService

import (
	"github.com/google/uuid"
)

type Task struct {
	ID     uuid.UUID `gorm:"primaryKey" json:"id"`
	Task   string    `json:"task"`
	IsDone bool      `json:"is_done"`
	UserID uuid.UUID `json:"user_id" gorm:"column:user_id"`
}

type TaskRequest struct {
	Task   *string `json:"task,omitempty"`
	IsDone *bool   `json:"is_done,omitempty"`
}

type UpdateTaskRequest struct {
	Task   *string `json:"task"`
	IsDone *bool   `json:"is_done"`
}
