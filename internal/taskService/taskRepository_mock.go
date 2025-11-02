package taskService

import "github.com/stretchr/testify/mock"

// MockTaskRepository - поддельный репозиторий
type MockTaskRepository struct {
	mock.Mock
}

func (m *MockTaskRepository) CreateTask(task Task) error {
	args := m.Called(task)
	return args.Error(0)
}

func (m *MockTaskRepository) GetAllTask() ([]Task, error) {
	args := m.Called()
	var tasks []Task
	if res := args.Get(0); res != nil {
		tasks = res.([]Task)
	}
	return tasks, args.Error(1)
}

func (m *MockTaskRepository) GetTaskByID(id string) (Task, error) {
	args := m.Called(id)
	var t Task
	if res := args.Get(0); res != nil {
		t = res.(Task)
	}
	return t, args.Error(1)
}

func (m *MockTaskRepository) UpdateTaskBy(task Task) error {
	args := m.Called(task)
	return args.Error(0)
}

func (m *MockTaskRepository) DeleteTask(id string) error {
	args := m.Called(id)
	return args.Error(0)
}
