package taskService

import "time"

type Task struct {
	ID        string    `gorm:"primaryKey" json:"id"`
	Task      string    `json:"task"`
	IsDone    bool      `json:"is_done"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

type TaskRequest struct {
	Task   *string `json:"task,omitempty"`
	IsDone *bool   `json:"is_done,omitempty"`
}

type UpdateTaskRequest struct {
	Task   *string `json:"task"`
	IsDone *bool   `json:"is_done"`
}
