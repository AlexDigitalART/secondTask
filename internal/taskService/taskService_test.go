package taskService

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestGetTasksByUserID(t *testing.T) {
	userID := uuid.New()

	tests := []struct {
		name      string
		userID    uuid.UUID
		mockSetup func(m *MockTaskRepository, userID uuid.UUID)
		wantErr   bool
	}{
		{
			name:   "успешное получение задач пользователя",
			userID: userID,
			mockSetup: func(m *MockTaskRepository, userID uuid.UUID) {
				tasks := []Task{
					{ID: uuid.MustParse("uuid-1"), Task: "Task 1", UserID: userID},
					{ID: uuid.MustParse("uuid-2"), Task: "Task 2", UserID: userID},
				}
				m.On("GetTasksByUserID", userID).Return(tasks, nil)
			},
			wantErr: false,
		},
		{
			name:   "пользователь без задач",
			userID: userID,
			mockSetup: func(m *MockTaskRepository, userID uuid.UUID) {
				m.On("GetTasksByUserID", userID).Return([]Task{}, nil)
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(MockTaskRepository)
			tt.mockSetup(mockRepo, tt.userID)

			service := NewTaskService(mockRepo)
			tasks, err := service.GetTasksByUserID(tt.userID)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, tasks)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}
