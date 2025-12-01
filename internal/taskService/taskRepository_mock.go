package taskService

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type MockTaskRepository struct {
	mock.Mock
}

func (m *MockTaskRepository) CreateTask(task *Task) error {
	args := m.Called(task)
	return args.Error(0)
}

func (m *MockTaskRepository) GetTasksByUserID(userID uuid.UUID) ([]Task, error) {
	args := m.Called(userID)
	var tasks []Task
	if res := args.Get(0); res != nil {
		tasks = res.([]Task)
	}
	return tasks, args.Error(1)
}

func (m *MockTaskRepository) GetAllTask() ([]Task, error) {
	args := m.Called()
	var tasks []Task
	if res := args.Get(0); res != nil {
		tasks = res.([]Task)
	}
	return tasks, args.Error(1)
}

func (m *MockTaskRepository) GetTaskByID(id uuid.UUID) (Task, error) {
	args := m.Called(id)
	var task Task
	if res := args.Get(0); res != nil {
		task = res.(Task)
	}
	return task, args.Error(1)
}

func (m *MockTaskRepository) UpdateTaskBy(task *Task) error {
	args := m.Called(task)
	return args.Error(0)
}

func (m *MockTaskRepository) DeleteTask(id uuid.UUID) error {
	args := m.Called(id)
	return args.Error(0)
}
