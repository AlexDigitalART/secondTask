package userService

import (
	"firstTask/internal/taskService"

	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID          `json:"id"`
	Email    string             `json:"email"`
	Password string             `json:"password"`
	Tasks    []taskService.Task `json:"tasks,omitempty" gorm:"foreignKey:UserID"`
}

type CreateUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateUserRequest struct {
	Email    *string `json:"email,omitempty"`
	Password *string `json:"password,omitempty"`
}
