package taskService

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateTask(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		mockSetup func(m *MockTaskRepository, input string)
		wantErr   bool
	}{
		{
			name:  "успешное создание задачи",
			input: "Test task",
			mockSetup: func(m *MockTaskRepository, input string) {
				m.On("CreateTask", mock.AnythingOfType("taskService.Task")).Return(nil)
			},
			wantErr: false,
		},
		{
			name:  "ошибка при создании задачи",
			input: "Bad task",
			mockSetup: func(m *MockTaskRepository, input string) {
				m.On("CreateTask", mock.AnythingOfType("taskService.Task")).Return(errors.New("db error"))
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(MockTaskRepository)
			tt.mockSetup(mockRepo, tt.input)

			service := NewTaskService(mockRepo)
			_, err := service.CreateTask(tt.input)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}

func TestGetAllTask(t *testing.T) {
	tests := []struct {
		name      string
		mockSetup func(m *MockTaskRepository)
		wantErr   bool
	}{
		{
			name: "успешное получение списка задач",
			mockSetup: func(m *MockTaskRepository) {
				tasks := []Task{
					{ID: "1", Task: "Test1", IsDone: false},
					{ID: "2", Task: "Test2", IsDone: true},
				}
				m.On("GetAllTask").Return(tasks, nil)
			},
			wantErr: false,
		},
		{
			name: "ошибка при получении списка",
			mockSetup: func(m *MockTaskRepository) {
				m.On("GetAllTask").Return([]Task{}, errors.New("db error"))
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(MockTaskRepository)
			tt.mockSetup(mockRepo)

			service := NewTaskService(mockRepo)
			_, err := service.GetAllTask()

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}

func TestGetTaskByID(t *testing.T) {
	tests := []struct {
		name      string
		id        string
		mockSetup func(m *MockTaskRepository, id string)
		wantErr   bool
	}{
		{
			name: "успешное получение задачи",
			id:   "123",
			mockSetup: func(m *MockTaskRepository, id string) {
				task := Task{ID: id, Task: "Test task", IsDone: false}
				m.On("GetTaskByID", id).Return(task, nil)
			},
			wantErr: false,
		},
		{
			name: "ошибка при получении задачи",
			id:   "404",
			mockSetup: func(m *MockTaskRepository, id string) {
				m.On("GetTaskByID", id).Return(Task{}, errors.New("not found"))
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(MockTaskRepository)
			tt.mockSetup(mockRepo, tt.id)

			service := NewTaskService(mockRepo)
			_, err := service.GetTaskByID(tt.id)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}

func TestUpdateTask(t *testing.T) {
	tests := []struct {
		name      string
		id        string
		req       UpdateTaskRequest
		mockSetup func(m *MockTaskRepository, id string, req UpdateTaskRequest)
		wantErr   bool
	}{
		{
			name: "успешное обновление задачи",
			id:   "123",
			req: func() UpdateTaskRequest {
				text := "Updated"
				done := true
				return UpdateTaskRequest{Task: &text, IsDone: &done}
			}(),
			mockSetup: func(m *MockTaskRepository, id string, req UpdateTaskRequest) {
				existing := Task{ID: id, Task: "Old", IsDone: false}
				updated := Task{ID: id, Task: *req.Task, IsDone: *req.IsDone}

				m.On("GetTaskByID", id).Return(existing, nil)
				m.On("UpdateTaskBy", updated).Return(nil)
			},
			wantErr: false,
		},
		{
			name: "ошибка при обновлении (задача не найдена)",
			id:   "404",
			req: func() UpdateTaskRequest {
				text := "Doesn't matter"
				done := false
				return UpdateTaskRequest{Task: &text, IsDone: &done}
			}(),
			mockSetup: func(m *MockTaskRepository, id string, req UpdateTaskRequest) {
				m.On("GetTaskByID", id).Return(Task{}, errors.New("not found"))
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(MockTaskRepository)
			tt.mockSetup(mockRepo, tt.id, tt.req)

			service := NewTaskService(mockRepo)
			_, err := service.UpdateTask(tt.id, tt.req)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}

func TestDeleteTask(t *testing.T) {
	tests := []struct {
		name      string
		id        string
		mockSetup func(m *MockTaskRepository, id string)
		wantErr   bool
	}{
		{
			name: "успешное удаление задачи",
			id:   "1",
			mockSetup: func(m *MockTaskRepository, id string) {
				m.On("DeleteTask", id).Return(nil)
			},
			wantErr: false,
		},
		{
			name: "ошибка при удалении задачи",
			id:   "2",
			mockSetup: func(m *MockTaskRepository, id string) {
				m.On("DeleteTask", id).Return(errors.New("db error"))
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(MockTaskRepository)
			tt.mockSetup(mockRepo, tt.id)

			service := NewTaskService(mockRepo)
			err := service.DeleteTask(tt.id)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}
